// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengemanager

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
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"beforeBalancesValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"af17d922": "beforeBalancesValid(bytes21[],uint256[])",
	"0f89fbff": "calculateBeforeValues(bytes21[],uint16[],uint256[])",
	"d931f370": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32,bytes21[],uint256[])",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"85ecb92a": "generatePreconditionHash(bytes32,uint64[2],bytes32)",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x61112d610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100605760003560e01c80624c28f6146100655780630f89fbff146100bd57806385ecb92a146102b2578063af17d92214610307578063d931f3701461043e575b600080fd5b6100ab6004803603608081101561007b57600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b031661058c565b60408051918252519081900360200190f35b610262600480360360608110156100d357600080fd5b810190602081018135600160201b8111156100ed57600080fd5b8201836020820111156100ff57600080fd5b803590602001918460208302840111600160201b8311171561012057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460208302840111600160201b831117156101a257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101f157600080fd5b82018360208201111561020357600080fd5b803590602001918460208302840111600160201b8311171561022457600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061067e945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561029e578181015183820152602001610286565b505050509050019250505060405180910390f35b6100ab600480360360808110156102c857600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091945050903591506108639050565b61042a6004803603604081101561031d57600080fd5b810190602081018135600160201b81111561033757600080fd5b82018360208201111561034957600080fd5b803590602001918460208302840111600160201b8311171561036a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103b957600080fd5b8201836020820111156103cb57600080fd5b803590602001918460208302840111600160201b831117156103ec57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108b7945050505050565b604080519115158252519081900360200190f35b6100ab600480360361010081101561045557600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561049957600080fd5b8201836020820111156104ab57600080fd5b803590602001918460208302840111600160201b831117156104cc57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561051b57600080fd5b82018360208201111561052d57600080fd5b803590602001918460208302840111600160201b8311171561054e57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610abc945050505050565b60408051600480825260a0820190925260009160609190816020015b6105b0611091565b8152602001906001900390816105a85790505090506105ce86610b8e565b816000815181106105db57fe5b60200260200101819052506105f8836001600160a01b0316610c0e565b8160018151811061060557fe5b602002602001018190525061061984610c0e565b8160028151811061062657fe5b60209081029190910101526106486affffffffffffffffffffff198616610c0e565b8160038151811061065557fe5b602002602001018190525061067161066c82610c8c565b610d3c565b519150505b949350505050565b6060600083519050606085516040519080825280602002602001820160405280156106b3578160200160208202803883390190505b50905060005b828110156108595760008682815181106106cf57fe5b60200260200101519050878161ffff16815181106106e957fe5b60200260200101516014601581106106fd57fe5b1a60f81b6001600160f81b03191661074a5785828151811061071b57fe5b6020026020010151838261ffff168151811061073357fe5b602002602001018181510191508181525050610850565b828161ffff168151811061075a57fe5b60200260200101516000146107b6576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420696e636c756465204e465420746f6b656e207477696365000000604482015290519081900360640190fd5b8582815181106107c257fe5b60200260200101516000141561081f576040805162461bcd60e51b815260206004820152601f60248201527f4e465420746f6b656e206d7573742068617665206e6f6e2d7a65726f20696400604482015290519081900360640190fd5b85828151811061082b57fe5b6020026020010151838261ffff168151811061084357fe5b6020026020010181815250505b506001016106b9565b5095945050505050565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b81516000908015806108c95750806001145b156108d8576001915050610ab6565b60005b60018203811015610a695760008582815181106108f457fe5b602002602001015160146015811061090857fe5b1a60f81b90506001600160f81b031981166109725785828151811061092957fe5b60200260200101516001600160581b03191686836001018151811061094a57fe5b60200260200101516001600160581b0319161161096d5760009350505050610ab6565b610a60565b600160f81b6001600160f81b031982161415610a545785828151811061099457fe5b60200260200101516001600160581b0319168683600101815181106109b557fe5b60200260200101516001600160581b0319161080610a4357508582815181106109da57fe5b60200260200101516001600160581b0319168683600101815181106109fb57fe5b60200260200101516001600160581b031916148015610a435750848281518110610a2157fe5b6020026020010151858360010181518110610a3857fe5b602002602001015111155b1561096d5760009350505050610ab6565b60009350505050610ab6565b506001016108db565b50600160f81b846001830381518110610a7e57fe5b6020026020010151601460158110610a9257fe5b1a60f81b6001600160f81b0319161115610ab0576000915050610ab6565b60019150505b92915050565b60008888888888888888604051602001808981526020018863ffffffff1663ffffffff1660e01b8152600401878152602001868152602001858152602001848152602001838051906020019060200280838360005b83811015610b29578181015183820152602001610b11565b50505050905001828051906020019060200280838360005b83811015610b59578181015183820152602001610b41565b505050509050019850505050505050505060405160208183030381529060405280519060200120905098975050505050505050565b610b96611091565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610bfb565b610be8611091565b815260200190600190039081610be05790505b508152600260209091015290505b919050565b610c16611091565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610c7b565b610c68611091565b815260200190600190039081610c605790505b508152600060209091015292915050565b610c94611091565b610c9e8251610e72565b610cef576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b610d446110bf565b6060820151600c60ff90911610610d96576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610dc3576040518060200160405280610dba8460000151610e79565b90529050610c09565b606082015160ff1660011415610e0a576040518060200160405280610dba846020015160000151856020015160400151866020015160600151876020015160200151610e9d565b606082015160ff1660021415610e2f5750604080516020810190915281518152610c09565b600360ff16826060015160ff1610158015610e5357506060820151600c60ff909116105b15610e70576040518060200160405280610dba8460400151610f45565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315610ef7575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120610676565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115610f95576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610fc2578160200160208202803883390190505b50805190915060005b8181101561101e57610fdb6110bf565b610ff7868381518110610fea57fe5b6020026020010151610d3c565b9050806000015184838151811061100a57fe5b602090810291909101015250600101610fcb565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561106757818101518382015260200161104f565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6040518060800160405280600081526020016110ab6110d1565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a7231582089220dee014023455ed2587707ef30e51da144565f29fb18448ce540c9f3fa1564736f6c634300050c0032"

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

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCaller) BeforeBalancesValid(opts *bind.CallOpts, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "beforeBalancesValid", _tokenTypes, _beforeBalances)
	return *ret0, err
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCallerSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
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

// GenerateAssertionHash is a free data retrieval call binding the contract method 0xd931f370.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, bytes21[] _tokenTypes, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _tokenTypes [][21]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _tokenTypes, _totalMessageValueAmounts)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0xd931f370.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, bytes21[] _tokenTypes, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _tokenTypes [][21]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _tokenTypes, _totalMessageValueAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0xd931f370.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, bytes21[] _tokenTypes, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _tokenTypes [][21]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _tokenTypes, _totalMessageValueAmounts)
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

// BisectionABI is the input ABI used to generate the binding from.
const BisectionABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"totalMessageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionIndex\",\"type\":\"uint256\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"}]"

// BisectionFuncSigs maps the 4-byte function signature to its string representation.
var BisectionFuncSigs = map[string]string{
	"d0fef902": "bisectAssertion(Challenge.Data storage,bytes32,bytes32[],uint256[],uint32,uint64[2],bytes21[])",
	"110112ae": "continueChallenge(Challenge.Data storage,uint256,bytes,bytes32,bytes32)",
}

// BisectionBin is the compiled bytecode used for deploying new contracts.
var BisectionBin = "0x6112b0610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063110112ae14610045578063d0fef90214610109575b600080fd5b81801561005157600080fd5b50610107600480360360a081101561006857600080fd5b813591602081013591810190606081016040820135600160201b81111561008e57600080fd5b8201836020820111156100a057600080fd5b803590602001918460018302840111600160201b831117156100c157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610300565b005b81801561011557600080fd5b50610107600480360361010081101561012d57600080fd5b813591602081013591810190606081016040820135600160201b81111561015357600080fd5b82018360208201111561016557600080fd5b803590602001918460208302840111600160201b8311171561018657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101d557600080fd5b8201836020820111156101e757600080fd5b803590602001918460208302840111600160201b8311171561020857600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080518082018252939663ffffffff8635169690959094606082019450925060200190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561028f57600080fd5b8201836020820111156102a157600080fd5b803590602001918460208302840111600160201b831117156102c257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506105db945050505050565b846001015482146103425760405162461bcd60e51b815260040180806020018281038252602b815260200180611222602b913960400191505060405180910390fd5b600585015467ffffffffffffffff164311156103a5576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038501600101546001600160a01b031633146103f35760405162461bcd60e51b815260040180806020018281038252602f81526020018061124d602f913960400191505060405180910390fd5b73__$800fcb2f4a98daa165a5cdb21a355d7a15$__63b792d767848484886001016040518563ffffffff1660e01b81526004018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b8381101561047257818101518382015260200161045a565b50505050905090810190601f16801561049f5780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b1580156104bf57600080fd5b505af41580156104d3573d6000803e3d6000fd5b505050506040513d60208110156104e957600080fd5b505161053c576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420617373657274696f6e2073656c6563746564000000000000604482015290519081900360640190fd5b60058501805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160601b1793840416011667ffffffffffffffff19919091161790556001850181905584546004860154604080516001600160a01b03928316815260208101889052815192909316927f18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00929081900390910190a25050505050565b60016005880154600160601b900460ff1660028111156105f757fe5b146106335760405162461bcd60e51b81526004018080602001828103825260348152602001806111ee6034913960400191505060405180910390fd5b8051158061064a5750805184518161064757fe5b06155b610694576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b805115806106ba575080518451816106a857fe5b04600160038751816106b657fe5b0403145b610704576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b600587015467ffffffffffffffff16431115610767576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038701600001546001600160a01b031633146107cb576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f7269676e616c2061737365727465722063616e20626973656374604482015290519081900360640190fd5b600060606108246040518060e00160405280600160038b51816107ea57fe5b040363ffffffff1681526020018981526020018881526020018763ffffffff1681526020018681526020018581526020018a815250610a95565b60018b015491935091508214610881576040805162461bcd60e51b815260206004820152601960248201527f446f6573206e6f74206d61746368207072657620737461746500000000000000604482015290519081900360640190fd5b60058901805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160611b1793840416011667ffffffffffffffff19919091161790556040516309898dc160e41b815260206004820181815283516024840152835173__$800fcb2f4a98daa165a5cdb21a355d7a15$__93639898dc1093869392839260440191858101910280838360005b83811015610928578181015183820152602001610910565b505050509050019250505060206040518083038186803b15801561094b57600080fd5b505af415801561095f573d6000803e3d6000fd5b505050506040513d602081101561097557600080fd5b505160018a015588546001600160a01b03167f9d5d1d0657f25018347f45be267e99ba4b45456b86a2c4b40a9660f71a564c1e60038b0160000160009054906101000a90046001600160a01b031689888a60405180856001600160a01b03166001600160a01b03168152602001806020018463ffffffff1663ffffffff16815260200180602001838103835286818151815260200191508051906020019060200280838360005b83811015610a34578181015183820152602001610a1c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610a73578181015183820152602001610a5b565b50505050905001965050505050505060405180910390a2505050505050505050565b60006060610aa16111bd565b836000015163ffffffff16604051908082528060200260200182016040528015610ad5578160200160208202803883390190505b5060408201528351606085015163ffffffff918216911681610af357fe5b0460010163ffffffff1660808201526000805b855163ffffffff168110156111a657856000015163ffffffff16866060015163ffffffff1681610b3257fe5b0663ffffffff16811415610b53576080830180516000190163ffffffff1690525b8560a0015151604051908082528060200260200182016040528015610b82578160200160208202803883390190505b506060840152600091505b8560a0015151821015610be7578560400151828760a001515183020181518110610bb357fe5b602002602001015183606001518381518110610bcb57fe5b6020908102919091010180519091019052600190910190610b8d565b73__$9836fa7140e5a33041d4b827682e675a30$__6385ecb92a87602001518360030281518110610c1457fe5b602002602001015188608001518960c001516040518463ffffffff1660e01b81526004018084815260200183600260200280838360005b83811015610c63578181015183820152602001610c4b565b50505050905001828152602001935050505060206040518083038186803b158015610c8d57600080fd5b505af4158015610ca1573d6000803e3d6000fd5b505050506040513d6020811015610cb757600080fd5b50518084526020870151805173__$9836fa7140e5a33041d4b827682e675a30$__9163d931f3709160036001870102908110610cef57fe5b602002602001015186608001518a602001518660030260010181518110610d1257fe5b60200260200101518b602001518760010160030260010181518110610d3357fe5b60200260200101518c602001518860030260020181518110610d5157fe5b60200260200101518d602001518960010160030260020181518110610d7257fe5b60200260200101518e60a001518c606001516040518963ffffffff1660e01b8152600401808981526020018863ffffffff1663ffffffff1681526020018781526020018681526020018581526020018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610e07578181015183820152602001610def565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610e46578181015183820152602001610e2e565b505050509050019a505050505050505050505060206040518083038186803b158015610e7157600080fd5b505af4158015610e85573d6000803e3d6000fd5b505050506040513d6020811015610e9b57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920181528151919092012090840151805183908110610ed857fe5b60209081029190910101528061119e578560a0015151604051908082528060200260200182016040528015610f17578160200160208202803883390190505b506060840152600091505b856040015151821015610f815785604001518281518110610f3f57fe5b602002602001015183606001518760a00151518481610f5a57fe5b0681518110610f6557fe5b6020908102919091010180519091019052600190910190610f22565b826000015173__$9836fa7140e5a33041d4b827682e675a30$__63d931f3708860200151896000015160030263ffffffff1681518110610fbd57fe5b602002602001015189606001518a60200151600181518110610fdb57fe5b60200260200101518b602001518c6000015160030260010163ffffffff168151811061100357fe5b60200260200101518c6020015160028151811061101c57fe5b60200260200101518d602001518e6000015160030260020163ffffffff168151811061104457fe5b60200260200101518e60a001518c606001516040518963ffffffff1660e01b8152600401808981526020018863ffffffff1663ffffffff1681526020018781526020018681526020018581526020018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156110d95781810151838201526020016110c1565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611118578181015183820152602001611100565b505050509050019a505050505050505050505060206040518083038186803b15801561114357600080fd5b505af4158015611157573d6000803e3d6000fd5b505050506040513d602081101561116d57600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019052805190820120908401525b600101610b06565b505060208101516040909101519092509050915091565b6040805160a08101825260008082526020820181905260609282018390528282019290925260808101919091529056fe43616e206f6e6c792062697365637420617373657274696f6e20696e20726573706f6e736520746f2061206368616c6c656e6765636f6e74696e75654368616c6c656e67653a20496e636f72726563742070726576696f75732073746174654f6e6c79206f726967696e616c206368616c6c656e6765722063616e20636f6e74696e7565206368616c6c656e6765a265627a7a7231582027c802df6e67446581f9a92d6903a80eace9da2f326d3e7011fafb6fd0f0ff3d64736f6c634300050c0032"

// DeployBisection deploys a new Ethereum contract, binding an instance of Bisection to it.
func DeployBisection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bisection, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	merkleLibAddr, _, _, _ := DeployMerkleLib(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$800fcb2f4a98daa165a5cdb21a355d7a15$__", merkleLibAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BisectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// Bisection is an auto generated Go binding around an Ethereum contract.
type Bisection struct {
	BisectionCaller     // Read-only binding to the contract
	BisectionTransactor // Write-only binding to the contract
	BisectionFilterer   // Log filterer for contract events
}

// BisectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BisectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BisectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BisectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BisectionSession struct {
	Contract     *Bisection        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BisectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BisectionCallerSession struct {
	Contract *BisectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BisectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BisectionTransactorSession struct {
	Contract     *BisectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BisectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BisectionRaw struct {
	Contract *Bisection // Generic contract binding to access the raw methods on
}

// BisectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BisectionCallerRaw struct {
	Contract *BisectionCaller // Generic read-only contract binding to access the raw methods on
}

// BisectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BisectionTransactorRaw struct {
	Contract *BisectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBisection creates a new instance of Bisection, bound to a specific deployed contract.
func NewBisection(address common.Address, backend bind.ContractBackend) (*Bisection, error) {
	contract, err := bindBisection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// NewBisectionCaller creates a new read-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionCaller(address common.Address, caller bind.ContractCaller) (*BisectionCaller, error) {
	contract, err := bindBisection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionCaller{contract: contract}, nil
}

// NewBisectionTransactor creates a new write-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionTransactor(address common.Address, transactor bind.ContractTransactor) (*BisectionTransactor, error) {
	contract, err := bindBisection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionTransactor{contract: contract}, nil
}

// NewBisectionFilterer creates a new log filterer instance of Bisection, bound to a specific deployed contract.
func NewBisectionFilterer(address common.Address, filterer bind.ContractFilterer) (*BisectionFilterer, error) {
	contract, err := bindBisection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BisectionFilterer{contract: contract}, nil
}

// bindBisection binds a generic wrapper to an already deployed contract.
func bindBisection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.BisectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transact(opts, method, params...)
}

// BisectionBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the Bisection contract.
type BisectionBisectedAssertionIterator struct {
	Event *BisectionBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *BisectionBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionBisectedAssertion)
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
		it.Event = new(BisectionBisectedAssertion)
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
func (it *BisectionBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionBisectedAssertion represents a BisectedAssertion event raised by the Bisection contract.
type BisectionBisectedAssertion struct {
	VmAddress                            common.Address
	Bisecter                             common.Address
	AfterHashAndMessageAndLogsBisections [][32]byte
	TotalSteps                           uint32
	TotalMessageAmounts                  []*big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0x9d5d1d0657f25018347f45be267e99ba4b45456b86a2c4b40a9660f71a564c1e.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_Bisection *BisectionFilterer) FilterBisectedAssertion(opts *bind.FilterOpts, vmAddress []common.Address) (*BisectionBisectedAssertionIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "BisectedAssertion", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &BisectionBisectedAssertionIterator{contract: _Bisection.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0x9d5d1d0657f25018347f45be267e99ba4b45456b86a2c4b40a9660f71a564c1e.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_Bisection *BisectionFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *BisectionBisectedAssertion, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "BisectedAssertion", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionBisectedAssertion)
				if err := _Bisection.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0x9d5d1d0657f25018347f45be267e99ba4b45456b86a2c4b40a9660f71a564c1e.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_Bisection *BisectionFilterer) ParseBisectedAssertion(log types.Log) (*BisectionBisectedAssertion, error) {
	event := new(BisectionBisectedAssertion)
	if err := _Bisection.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the Bisection contract.
type BisectionContinuedChallengeIterator struct {
	Event *BisectionContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *BisectionContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionContinuedChallenge)
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
		it.Event = new(BisectionContinuedChallenge)
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
func (it *BisectionContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionContinuedChallenge represents a ContinuedChallenge event raised by the Bisection contract.
type BisectionContinuedChallenge struct {
	VmAddress      common.Address
	Challenger     common.Address
	AssertionIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) FilterContinuedChallenge(opts *bind.FilterOpts, vmAddress []common.Address) (*BisectionContinuedChallengeIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "ContinuedChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &BisectionContinuedChallengeIterator{contract: _Bisection.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *BisectionContinuedChallenge, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "ContinuedChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionContinuedChallenge)
				if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) ParseContinuedChallenge(log types.Log) (*BisectionContinuedChallenge, error) {
	event := new(BisectionContinuedChallenge)
	if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ChallengeManagerABI is the input ABI used to generate the binding from.
const ChallengeManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"totalMessageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionIndex\",\"type\":\"uint256\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"challengerWrong\",\"type\":\"bool\"}],\"name\":\"TimedOutChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"}],\"name\":\"asserterTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challengeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalMessageAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32\",\"name\":\"_totalSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"}],\"name\":\"challengerTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assertionToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"continueChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[2]\",\"name\":\"_beforeHashAndInbox\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"_afterHashAndMessages\",\"type\":\"bytes32[5]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeManagerFuncSigs = map[string]string{
	"36ddd35d": "asserterTimedOut(address)",
	"20f1dc5b": "bisectAssertion(address,bytes32,bytes32[],uint256[],uint32,uint64[2],bytes21[])",
	"bf06db66": "challengerTimedOut(address)",
	"eb94f27b": "continueChallenge(address,uint256,bytes,bytes32,bytes32)",
	"208e04d4": "initiateChallenge(address[2],uint128[2],uint32,bytes32)",
	"99fe4e8b": "oneStepProof(address,bytes32[2],uint64[2],bytes21[],bytes32[5],uint256[],bytes)",
}

// ChallengeManagerBin is the compiled bytecode used for deploying new contracts.
var ChallengeManagerBin = "0x608060405234801561001057600080fd5b506112de806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063208e04d41461006757806320f1dc5b1461009957806336ddd35d1461028c57806399fe4e8b146102b2578063bf06db66146104ed578063eb94f27b14610513575b600080fd5b610097600480360360c081101561007d57600080fd5b506040810163ffffffff60808301351660a08301356105d1565b005b61009760048036036101008110156100b057600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156100df57600080fd5b8201836020820111156100f157600080fd5b803590602001918460208302840111600160201b8311171561011257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561016157600080fd5b82018360208201111561017357600080fd5b803590602001918460208302840111600160201b8311171561019457600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080518082018252939663ffffffff8635169690959094606082019450925060200190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561021b57600080fd5b82018360208201111561022d57600080fd5b803590602001918460208302840111600160201b8311171561024e57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061078f945050505050565b610097600480360360208110156102a257600080fd5b50356001600160a01b0316610943565b61009760048036036101a08110156102c957600080fd5b6040805180820182526001600160a01b0384351693928301929160608301919060208401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561034957600080fd5b82018360208201111561035b57600080fd5b803590602001918460208302840111600160201b8311171561037c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160a0818101909252939695948181019493509150600590839083908082843760009201919091525091949392602081019250359050600160201b8111156103f757600080fd5b82018360208201111561040957600080fd5b803590602001918460208302840111600160201b8311171561042a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561047957600080fd5b82018360208201111561048b57600080fd5b803590602001918460018302840111600160201b831117156104ac57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a5c945050505050565b6100976004803603602081101561050357600080fd5b50356001600160a01b0316610d36565b610097600480360360a081101561052957600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561055857600080fd5b82018360208201111561056a57600080fd5b803590602001918460018302840111600160201b8311171561058b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610e4f565b33600090815260208190526040902060010154156106205760405162461bcd60e51b81526004018080602001828103825260238152602001806112876023913960400191505060405180910390fd5b6040805160e081018252338152602081018390528151808301835290918281019190869060029083908390808284376000920191909152505050815260408051808201825260209092019190879060029083908390808284376000920191909152505050815263ffffffff841643810167ffffffffffffffff1660208301526040820152606001600190523360009081526020818152604091829020835181546001600160a01b0319166001600160a01b03909116178155908301516001820155908201516106f590600280840191906110db565b50606082015161070b9060038301906002611180565b50608082015160058201805460a085015163ffffffff1668010000000000000000026bffffffff00000000000000001967ffffffffffffffff90941667ffffffffffffffff1990921691909117929092169190911780825560c0840151919060ff60601b1916600160601b83600281111561078257fe5b0217905550505050505050565b6000806000896001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f5eea941ded5358daea4da7ea13a2128fd$__63d0fef902828989898989896040518863ffffffff1660e01b81526004018088815260200187815260200180602001806020018663ffffffff1663ffffffff16815260200185600260200280838360005b8381101561083557818101518382015260200161081d565b5050505090500180602001848103845289818151815260200191508051906020019060200280838360005b83811015610878578181015183820152602001610860565b50505050905001848103835288818151815260200191508051906020019060200280838360005b838110156108b757818101518382015260200161089f565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156108f65781810151838201526020016108de565b505050509050019a505050505050505050505060006040518083038186803b15801561092157600080fd5b505af4158015610935573d6000803e3d6000fd5b505050505050505050505050565b6001600160a01b038116600090815260208190526040902060016005820154600160601b900460ff16600281111561097757fe5b146109b35760405162461bcd60e51b815260040180806020018281038252602e815260200180611229602e913960400191505060405180910390fd5b600581015467ffffffffffffffff164311610a0f576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b610a1881610f68565b604080516001815290516001600160a01b038416917f2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35919081900360200190a25050565b6000806000896001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f55f7f918072b72dcc999cdc8e581605f5$__631af59c1a828989898989896040518863ffffffff1660e01b81526004018088815260200187600260200280838360005b83811015610ae2578181015183820152602001610aca565b5050505090500186600260200280838360005b83811015610b0d578181015183820152602001610af5565b5050509201915050602081018560a080838360005b83811015610b3a578181015183820152602001610b22565b505050509050018060200180602001848103845288818151815260200191508051906020019060200280838360005b83811015610b81578181015183820152602001610b69565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015610bc0578181015183820152602001610ba8565b50505050905001848103825285818151815260200191508051906020019080838360005b83811015610bfc578181015183820152602001610be4565b50505050905090810190601f168015610c295780820380516001836020036101000a031916815260200191505b509a505050505050505050505060006040518083038186803b158015610c4e57600080fd5b505af4158015610c62573d6000803e3d6000fd5b50505050610c6f8161104c565b876001600160a01b03167ffd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572338460405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610cf1578181015183820152602001610cd9565b50505050905090810190601f168015610d1e5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050505050505050565b6001600160a01b038116600090815260208190526040902060026005820154600160601b900460ff166002811115610d6a57fe5b14610da65760405162461bcd60e51b81526004018080602001828103825260308152602001806112576030913960400191505060405180910390fd5b600581015467ffffffffffffffff164311610e02576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b610e0b8161104c565b604080516000815290516001600160a01b038416917f2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35919081900360200190a25050565b6000806000876001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f5eea941ded5358daea4da7ea13a2128fd$__63110112ae82878787876040518663ffffffff1660e01b81526004018086815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b83811015610efa578181015183820152602001610ee2565b50505050905090810190601f168015610f275780820380516001836020036101000a031916815260200191505b50965050505050505060006040518083038186803b158015610f4857600080fd5b505af4158015610f5c573d6000803e3d6000fd5b50505050505050505050565b805460408051808201825260008152600280850154600160801b81046001600160801b03908116918116929092040116602082015290516308b0246f60e21b81526001600160a01b03909216916322c091bc916003850191600481019060440183825b81546001600160a01b03168152600190910190602001808311610fcb5750839050604080838360005b8381101561100c578181015183820152602001610ff4565b5050505090500192505050600060405180830381600087803b15801561103157600080fd5b505af1158015611045573d6000803e3d6000fd5b5050505050565b80546040805180820182526002808501546001600160801b03808216600160801b909204811692909204011681526000602082015290516308b0246f60e21b81526003840180546001600160a01b03908116600480850191825291909516946322c091bc9492939091604482019190880190602401808311610fcb575050825181528260408083836020610ff4565b6001830191839082156111705791602002820160005b8382111561113b57835183826101000a8154816001600160801b0302191690836001600160801b031602179055509260200192601001602081600f010492830192600103026110f1565b801561116e5782816101000a8154906001600160801b030219169055601001602081600f0104928301926001030261113b565b505b5061117c9291506111d4565b5090565b82600281019282156111c8579160200282015b828111156111c857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611193565b5061117c929150611204565b61120191905b8082111561117c5780546fffffffffffffffffffffffffffffffff191681556001016111da565b90565b61120191905b8082111561117c5780546001600160a01b031916815560010161120a56fe43616e206f6e6c792074696d65206f7574206173736572746572206966206974206973207468656972207475726e43616e206f6e6c792074696d65206f7574206368616c6c656e676572206966206974206973207468656972207475726e5468657265206d757374206265206e6f206578697374696e67206368616c6c656e6765a265627a7a72315820d52de8279a86a400a576d26cf59e4e33e6fc4707b31e4298beecdf8980173a9964736f6c634300050c0032"

// DeployChallengeManager deploys a new Ethereum contract, binding an instance of ChallengeManager to it.
func DeployChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	oneStepProofAddr, _, _, _ := DeployOneStepProof(auth, backend)
	ChallengeManagerBin = strings.Replace(ChallengeManagerBin, "__$f55f7f918072b72dcc999cdc8e581605f5$__", oneStepProofAddr.String()[2:], -1)

	bisectionAddr, _, _, _ := DeployBisection(auth, backend)
	ChallengeManagerBin = strings.Replace(ChallengeManagerBin, "__$f5eea941ded5358daea4da7ea13a2128fd$__", bisectionAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeManagerBin), backend)
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

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x36ddd35d.
//
// Solidity: function asserterTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactor) AsserterTimedOut(opts *bind.TransactOpts, _vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "asserterTimedOut", _vmAddress)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x36ddd35d.
//
// Solidity: function asserterTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerSession) AsserterTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x36ddd35d.
//
// Solidity: function asserterTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) AsserterTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x20f1dc5b.
//
// Solidity: function bisectAssertion(address _challengeId, bytes32 _beforeInbox, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes) returns()
func (_ChallengeManager *ChallengeManagerTransactor) BisectAssertion(opts *bind.TransactOpts, _challengeId common.Address, _beforeInbox [32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "bisectAssertion", _challengeId, _beforeInbox, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x20f1dc5b.
//
// Solidity: function bisectAssertion(address _challengeId, bytes32 _beforeInbox, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes) returns()
func (_ChallengeManager *ChallengeManagerSession) BisectAssertion(_challengeId common.Address, _beforeInbox [32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertion(&_ChallengeManager.TransactOpts, _challengeId, _beforeInbox, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x20f1dc5b.
//
// Solidity: function bisectAssertion(address _challengeId, bytes32 _beforeInbox, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) BisectAssertion(_challengeId common.Address, _beforeInbox [32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertion(&_ChallengeManager.TransactOpts, _challengeId, _beforeInbox, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xbf06db66.
//
// Solidity: function challengerTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ChallengerTimedOut(opts *bind.TransactOpts, _vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "challengerTimedOut", _vmAddress)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xbf06db66.
//
// Solidity: function challengerTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerSession) ChallengerTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xbf06db66.
//
// Solidity: function challengerTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ChallengerTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xeb94f27b.
//
// Solidity: function continueChallenge(address _vmAddress, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ContinueChallenge(opts *bind.TransactOpts, _vmAddress common.Address, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "continueChallenge", _vmAddress, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xeb94f27b.
//
// Solidity: function continueChallenge(address _vmAddress, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerSession) ContinueChallenge(_vmAddress common.Address, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmAddress, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xeb94f27b.
//
// Solidity: function continueChallenge(address _vmAddress, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ContinueChallenge(_vmAddress common.Address, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmAddress, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "initiateChallenge", _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerSession) InitiateChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) InitiateChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _players, _escrows, _challengePeriod, _challengeRoot)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x99fe4e8b.
//
// Solidity: function oneStepProof(address _vmAddress, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactor) OneStepProof(opts *bind.TransactOpts, _vmAddress common.Address, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "oneStepProof", _vmAddress, _beforeHashAndInbox, _timeBounds, _tokenTypes, _afterHashAndMessages, _amounts, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x99fe4e8b.
//
// Solidity: function oneStepProof(address _vmAddress, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerSession) OneStepProof(_vmAddress common.Address, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProof(&_ChallengeManager.TransactOpts, _vmAddress, _beforeHashAndInbox, _timeBounds, _tokenTypes, _afterHashAndMessages, _amounts, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x99fe4e8b.
//
// Solidity: function oneStepProof(address _vmAddress, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) OneStepProof(_vmAddress common.Address, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProof(&_ChallengeManager.TransactOpts, _vmAddress, _beforeHashAndInbox, _timeBounds, _tokenTypes, _afterHashAndMessages, _amounts, _proof)
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
	VmAddress                            common.Address
	Bisecter                             common.Address
	AfterHashAndMessageAndLogsBisections [][32]byte
	TotalSteps                           uint32
	TotalMessageAmounts                  []*big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0x9d5d1d0657f25018347f45be267e99ba4b45456b86a2c4b40a9660f71a564c1e.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) FilterBisectedAssertion(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerBisectedAssertionIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "BisectedAssertion", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerBisectedAssertionIterator{contract: _ChallengeManager.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0x9d5d1d0657f25018347f45be267e99ba4b45456b86a2c4b40a9660f71a564c1e.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ChallengeManagerBisectedAssertion, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "BisectedAssertion", vmAddressRule)
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0x9d5d1d0657f25018347f45be267e99ba4b45456b86a2c4b40a9660f71a564c1e.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) ParseBisectedAssertion(log types.Log) (*ChallengeManagerBisectedAssertion, error) {
	event := new(ChallengeManagerBisectedAssertion)
	if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
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
	VmAddress      common.Address
	Challenger     common.Address
	AssertionIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) FilterContinuedChallenge(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerContinuedChallengeIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ContinuedChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerContinuedChallengeIterator{contract: _ChallengeManager.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerContinuedChallenge, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ContinuedChallenge", vmAddressRule)
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) ParseContinuedChallenge(log types.Log) (*ChallengeManagerContinuedChallenge, error) {
	event := new(ChallengeManagerContinuedChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
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
	VmAddress common.Address
	Asserter  common.Address
	Proof     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xfd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572.
//
// Solidity: event OneStepProofCompleted(address indexed vmAddress, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerOneStepProofCompletedIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerOneStepProofCompletedIterator{contract: _ChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xfd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572.
//
// Solidity: event OneStepProofCompleted(address indexed vmAddress, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeManagerOneStepProofCompleted, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", vmAddressRule)
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xfd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572.
//
// Solidity: event OneStepProofCompleted(address indexed vmAddress, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*ChallengeManagerOneStepProofCompleted, error) {
	event := new(ChallengeManagerOneStepProofCompleted)
	if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
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
	VmAddress       common.Address
	ChallengerWrong bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedOutChallenge is a free log retrieval operation binding the contract event 0x2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35.
//
// Solidity: event TimedOutChallenge(address indexed vmAddress, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) FilterTimedOutChallenge(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerTimedOutChallengeIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "TimedOutChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTimedOutChallengeIterator{contract: _ChallengeManager.contract, event: "TimedOutChallenge", logs: logs, sub: sub}, nil
}

// WatchTimedOutChallenge is a free log subscription operation binding the contract event 0x2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35.
//
// Solidity: event TimedOutChallenge(address indexed vmAddress, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) WatchTimedOutChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerTimedOutChallenge, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "TimedOutChallenge", vmAddressRule)
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

// ParseTimedOutChallenge is a log parse operation binding the contract event 0x2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35.
//
// Solidity: event TimedOutChallenge(address indexed vmAddress, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) ParseTimedOutChallenge(log types.Log) (*ChallengeManagerTimedOutChallenge, error) {
	event := new(ChallengeManagerTimedOutChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
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

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeManagerFuncSigs = map[string]string{
	"208e04d4": "initiateChallenge(address[2],uint128[2],uint32,bytes32)",
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

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initiateChallenge", players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// IVMTrackerABI is the input ABI used to generate the binding from.
const IVMTrackerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVMTrackerFuncSigs maps the 4-byte function signature to its string representation.
var IVMTrackerFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
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

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _players, _rewards)
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
var MerkleLibBin = "0x610575610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043f565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b6000815b600181511115610422576060600282516001018161031f57fe5b04604051908082528060200260200182016040528015610349578160200160208202803883390190505b50905060005b815181101561041a5782518160020260010110156103e25782816002028151811061037657fe5b602002602001015183826002026001018151811061039057fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d157fe5b602002602001018181525050610412565b8281600202815181106103f157fe5b602002602001015182828151811061040557fe5b6020026020010181815250505b60010161034f565b509050610305565b8060008151811061042f57fe5b6020026020010151915050919050565b600080838160205b88518111610532578089015193506020818a51036020018161046557fe5b0491505b60008211801561047c5750600286066001145b801561048a57508160020a86115b1561049d57600286046001019550610469565b600286066104e85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104e057fe5b04955061052a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052357fe5b0460010195505b602001610447565b50509094149594505050505056fea265627a7a723158203fb57bd4323cd277c34e2a619fb9b56d2a507018bfb9267f361b0788f05a1d2964736f6c634300050c0032"

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
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[7]\",\"name\":\"fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes21[]\",\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageValue\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"1af59c1a": "oneStepProof(Challenge.Data storage,bytes32[2],uint64[2],bytes21[],bytes32[5],uint256[],bytes)",
	"97ce6508": "validateProof(bytes32[7],uint64[2],bytes21[],uint256[],bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x614172610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c80631af59c1a1461004557806397ce650814610279575b600080fd5b61027760048036036101a081101561005c57600080fd5b604080518082018252833593928301929160608301919060208401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b8111156100d357600080fd5b8201836020820111156100e557600080fd5b803590602001918460208302840111600160201b8311171561010657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160a0818101909252939695948181019493509150600590839083908082843760009201919091525091949392602081019250359050600160201b81111561018157600080fd5b82018360208201111561019357600080fd5b803590602001918460208302840111600160201b831117156101b457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561020357600080fd5b82018360208201111561021557600080fd5b803590602001918460018302840111600160201b8311171561023657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610490945050505050565b005b61047e600480360361018081101561029057600080fd5b810190808060e001906007806020026040519081016040528092919082600760200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561030657600080fd5b82018360208201111561031857600080fd5b803590602001918460208302840111600160201b8311171561033957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561038857600080fd5b82018360208201111561039a57600080fd5b803590602001918460208302840111600160201b831117156103bb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561040a57600080fd5b82018360208201111561041c57600080fd5b803590602001918460018302840111600160201b8311171561043d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108d2945050505050565b60408051918252519081900360200190f35b60016005880154600160601b900460ff1660028111156104ac57fe5b146104e85760405162461bcd60e51b81526004018080602001828103825260398152602001806141056039913960400191505060405180910390fd5b600587015467ffffffffffffffff1643111561054b576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6520737465702070726f6f66206d697373656420646561646c696e650000604482015290519081900360640190fd5b600187015486516020880151604080516342f65c9560e11b81526004810184815273__$9836fa7140e5a33041d4b827682e675a30$__946385ecb92a9490938c9391929160240190849080838360005b838110156105b357818101518382015260200161059b565b50505050905001828152602001935050505060206040518083038186803b1580156105dd57600080fd5b505af41580156105f1573d6000803e3d6000fd5b505050506040513d602081101561060757600080fd5b505173__$9836fa7140e5a33041d4b827682e675a30$__63d931f370866000602002015160018881602002015189600260200201518a600360200201518b600460200201518d8c6040518963ffffffff1660e01b8152600401808981526020018863ffffffff1681526020018781526020018681526020018581526020018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156106cb5781810151838201526020016106b3565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561070a5781810151838201526020016106f2565b505050509050019a505050505050505050505060206040518083038186803b15801561073557600080fd5b505af4158015610749573d6000803e3d6000fd5b505050506040513d602081101561075f57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920190528051910120146107c65760405162461bcd60e51b8152600401808060200182810382526026815260200180613fe16026913960400191505060405180910390fd5b600061087d6040518060e00160405280896000600281106107e357fe5b60200201518152602001896001600281106107fa57fe5b602002015181526020018660006005811061081157fe5b602002015181526020018660016005811061082857fe5b602002015181526020018660026005811061083f57fe5b602002015181526020018660036005811061085657fe5b602002015181526020018660046005811061086d57fe5b60200201519052878786866108d2565b905080156108c8576040805162461bcd60e51b8152602060048201526013602482015272141c9bdbd9881dd85cc81a5b98dbdc9c9958dd606a1b604482015290519081900360640190fd5b5050505050505050565b6080850151606086015160009182918291829114801590600019906109d25760005b88518167ffffffffffffffff16101561098957888167ffffffffffffffff168151811061091d57fe5b6020026020010151600014610981578160070b6000191461097d576040805162461bcd60e51b81526020600482015260156024820152746d756c7469706c65206f7574206d6573736167657360581b604482015290519081900360640190fd5b8091505b6001016108f4565b508060070b600019146109cd57878160070b815181106109a557fe5b60200260200101519350888160070b815181106109be57fe5b60200260200101519450600192505b610a64565b60005b88518167ffffffffffffffff161015610a6257888167ffffffffffffffff16815181106109fe57fe5b6020026020010151600014610a5a576040805162461bcd60e51b815260206004820152601b60248201527f4d7573742068617665206e6f206d6573736167652076616c7565730000000000604482015290519081900360640190fd5b6001016109d5565b505b610b416040518061018001604052808d600060078110610a8057fe5b602002015181526020018c81526020018d600160078110610a9d57fe5b602002015181526020018d600260078110610ab457fe5b602002015181526020018d600360078110610acb57fe5b602002015181526020018d600460078110610ae257fe5b602002015181526020018d600560078110610af957fe5b602002015181526020018d600660078110610b1057fe5b60200201518152602001876001600160581b0319168152602001868152602001851515815260200189815250610b50565b9b9a5050505050505050505050565b60008080806060610b5f613f18565b610b67613f18565b610b7088611b0e565b939950929650909450925090506001600060ff8816821415610bc657610bbf8386600081518110610b9d57fe5b602002602001015187600181518110610bb257fe5b6020026020010151611f4e565b9150611962565b60ff881660021415610c0557610bbf8386600081518110610be357fe5b602002602001015187600181518110610bf857fe5b6020026020010151611f9e565b60ff881660031415610c4457610bbf8386600081518110610c2257fe5b602002602001015187600181518110610c3757fe5b6020026020010151611fdf565b60ff881660041415610c8357610bbf8386600081518110610c6157fe5b602002602001015187600181518110610c7657fe5b6020026020010151612020565b60ff881660051415610cc257610bbf8386600081518110610ca057fe5b602002602001015187600181518110610cb557fe5b6020026020010151612071565b60ff881660061415610d0157610bbf8386600081518110610cdf57fe5b602002602001015187600181518110610cf457fe5b60200260200101516120c2565b60ff881660071415610d4057610bbf8386600081518110610d1e57fe5b602002602001015187600181518110610d3357fe5b6020026020010151612113565b60ff881660081415610d9457610bbf8386600081518110610d5d57fe5b602002602001015187600181518110610d7257fe5b602002602001015188600281518110610d8757fe5b6020026020010151612164565b60ff881660091415610de857610bbf8386600081518110610db157fe5b602002602001015187600181518110610dc657fe5b602002602001015188600281518110610ddb57fe5b60200260200101516121ce565b60ff8816600a1415610e2757610bbf8386600081518110610e0557fe5b602002602001015187600181518110610e1a57fe5b6020026020010151612227565b60ff881660101415610e6657610bbf8386600081518110610e4457fe5b602002602001015187600181518110610e5957fe5b6020026020010151612268565b60ff881660111415610ea557610bbf8386600081518110610e8357fe5b602002602001015187600181518110610e9857fe5b60200260200101516122a9565b60ff881660121415610ee457610bbf8386600081518110610ec257fe5b602002602001015187600181518110610ed757fe5b60200260200101516122ea565b60ff881660131415610f2357610bbf8386600081518110610f0157fe5b602002602001015187600181518110610f1657fe5b602002602001015161232b565b60ff881660141415610f6257610bbf8386600081518110610f4057fe5b602002602001015187600181518110610f5557fe5b602002602001015161236c565b60ff881660151415610f8c57610bbf8386600081518110610f7f57fe5b6020026020010151612398565b60ff881660161415610fcb57610bbf8386600081518110610fa957fe5b602002602001015187600181518110610fbe57fe5b60200260200101516123de565b60ff88166017141561100a57610bbf8386600081518110610fe857fe5b602002602001015187600181518110610ffd57fe5b602002602001015161241f565b60ff88166018141561104957610bbf838660008151811061102757fe5b60200260200101518760018151811061103c57fe5b6020026020010151612460565b60ff88166019141561107357610bbf838660008151811061106657fe5b60200260200101516124a1565b60ff8816601a14156110b257610bbf838660008151811061109057fe5b6020026020010151876001815181106110a557fe5b60200260200101516124d7565b60ff8816601b14156110f157610bbf83866000815181106110cf57fe5b6020026020010151876001815181106110e457fe5b6020026020010151612518565b60ff88166020141561111b57610bbf838660008151811061110e57fe5b6020026020010151612559565b60ff88166021141561114557610bbf838660008151811061113857fe5b6020026020010151612575565b60ff88166030141561116f57610bbf838660008151811061116257fe5b6020026020010151612590565b60ff88166031141561118457610bbf83612598565b60ff88166032141561119957610bbf836125b9565b60ff8816603314156111c357610bbf83866000815181106111b657fe5b60200260200101516125d2565b60ff8816603414156111ed57610bbf83866000815181106111e057fe5b60200260200101516125eb565b60ff88166035141561122c57610bbf838660008151811061120a57fe5b60200260200101518760018151811061121f57fe5b6020026020010151612601565b60ff88166036141561124157610bbf83612649565b60ff88166037141561125b57610bbf83856000015161267b565b60ff88166038141561128557610bbf838660008151811061127857fe5b602002602001015161268d565b60ff88166039141561131257611299613f79565b6112a88b61016001518861269f565b9199509750905087156112ec5760405162461bcd60e51b81526004018080602001828103825260218152602001806140e46021913960400191505060405180910390fd5b6112fc858263ffffffff61282916565b61130c848263ffffffff61284b16565b50611962565b60ff8816603a141561132757610bbf83612868565b60ff8816603b141561133857611962565b60ff8816603c141561134d57610bbf83612888565b60ff8816603d141561137757610bbf838660008151811061136a57fe5b60200260200101516128a1565b60ff8816604014156113a157610bbf838660008151811061139457fe5b60200260200101516128cf565b60ff8816604114156113e057610bbf83866000815181106113be57fe5b6020026020010151876001815181106113d357fe5b60200260200101516128f1565b60ff88166042141561143457610bbf83866000815181106113fd57fe5b60200260200101518760018151811061141257fe5b60200260200101518860028151811061142757fe5b6020026020010151612923565b60ff88166043141561147357610bbf838660008151811061145157fe5b60200260200101518760018151811061146657fe5b6020026020010151612965565b60ff8816604414156114c757610bbf838660008151811061149057fe5b6020026020010151876001815181106114a557fe5b6020026020010151886002815181106114ba57fe5b6020026020010151612977565b60ff88166050141561150657610bbf83866000815181106114e457fe5b6020026020010151876001815181106114f957fe5b6020026020010151612999565b60ff88166051141561155a57610bbf838660008151811061152357fe5b60200260200101518760018151811061153857fe5b60200260200101518860028151811061154d57fe5b6020026020010151612a0f565b60ff88166052141561158457610bbf838660008151811061157757fe5b6020026020010151612a87565b60ff88166060141561159957610bbf83612aba565b60ff881660611415611683576115c383866000815181106115b657fe5b6020026020010151612ac0565b60e08c015160c08d01516040805160208082019390935280820185905281518082038301815260609091019091528051910120929450909250146116385760405162461bcd60e51b81526004018080602001828103825260258152602001806140706025913960400191505060405180910390fd5b8960a001518a608001511461167e5760405162461bcd60e51b81526004018080602001828103825260278152602001806140956027913960400191505060405180910390fd5b611962565b60ff88166070141561183f576000806116b085886000815181106116a357fe5b6020026020010151612ae4565b809450819550829650839750505050508b61010001516001600160581b031916826001600160581b0319161461172d576040805162461bcd60e51b815260206004820152601960248201527f546f6b656e207479706520646f6573206e6f74206d6174636800000000000000604482015290519081900360640190fd5b8b6101200151811461177e576040805162461bcd60e51b8152602060048201526015602482015274082dadeeadce840c8decae640dcdee840dac2e8c6d605b1b604482015290519081900360640190fd5b8b60a001518c6080015184604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146117f25760405162461bcd60e51b81526004018080602001828103825260288152602001806140bc6028913960400191505060405180910390fd5b8b60e001518c60c00151146118385760405162461bcd60e51b81526004018080602001828103825260268152602001806140296026913960400191505060405180910390fd5b5050611962565b60ff8816607214156118fb576040805160028082526060828101909352816020015b611869613f79565b81526020019060019003908161186157505060208c015190915061189e9060005b602002015167ffffffffffffffff16612cc8565b816000815181106118ab57fe5b60200260200101819052506118ca8b6020015160016002811061188a57fe5b816001815181106118d757fe5b602002602001018190525061130c6118ee82612d46565b859063ffffffff61284b16565b60ff88166073141561193857610bbf838660008151811061191857fe5b602002602001015160405180602001604052808e60400151815250612df6565b60ff88166074141561194d5760009150611962565b60ff8816607514156119625761196283612e68565b806119f3578960a001518a60800151146119ad5760405162461bcd60e51b81526004018080602001828103825260278152602001806140956027913960400191505060405180910390fd5b8960e001518a60c00151146119f35760405162461bcd60e51b81526004018080602001828103825260268152602001806140296026913960400191505060405180910390fd5b81611a555760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a0840151511415611a4d57611a4883612e72565b611a55565b60a083015183525b611a5e84612e7c565b8a5114611a9c5760405162461bcd60e51b81526004018080602001828103825260228152602001806140076022913960400191505060405180910390fd5b611aa583612e7c565b8a6060015114611afc576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b60006060611b1a613f18565b611b22613f18565b60008080611b2e613f18565b611b3781612f11565b611b4689610160015184612f1b565b9094509092509050611b56613f18565b611b5f82613020565b905060008a61016001518581518110611b7457fe5b602001015160f81c60f81b60f81c905060008b61016001518660010181518110611b9a57fe5b016020015160f81c90506000611baf8261307e565b9050606081604051908082528060200260200182016040528015611bed57816020015b611bda613f79565b815260200190600190039081611bd25790505b5090506002880197508360ff1660001480611c0b57508360ff166001145b611c5c576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff8416611cff576040805160208082018084528951516353409fab60e01b90915260ff87166024840152604483015291519091829173__$d969135829891f807aa9c34494da4ecd99$__916353409fab916064808601929190818703018186803b158015611cca57600080fd5b505af4158015611cde573d6000803e3d6000fd5b505050506040513d6020811015611cf457600080fd5b505190528652611e56565b611d07613f79565b611d168f61016001518a61269f565b909a5090985090508715611d71576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b8215611d95578082600081518110611d8557fe5b6020026020010181905250611da5565b611da5868263ffffffff61284b16565b604051806020016040528073__$d969135829891f807aa9c34494da4ecd99$__63264f384b87611dd486613098565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b158015611e2457600080fd5b505af4158015611e38573d6000803e3d6000fd5b505050506040513d6020811015611e4e57600080fd5b505190528752505b60ff84165b82811015611eea57611e728f61016001518a61269f565b8451859085908110611e8057fe5b6020908102919091010152995097508715611ee2576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611e5b565b815115611f37575060005b8460ff16825103811015611f3757611f2f828260018551030381518110611f1857fe5b60200260200101518861284b90919063ffffffff16565b600101611ef5565b50919d919c50939a50919850939650945050505050565b6000611f59836131ce565b1580611f6b5750611f69826131ce565b155b15611f7857506000611f97565b82518251808201611f8f878263ffffffff6131d916565b600193505050505b9392505050565b6000611fa9836131ce565b1580611fbb5750611fb9826131ce565b155b15611fc857506000611f97565b82518251808202611f8f878263ffffffff6131d916565b6000611fea836131ce565b1580611ffc5750611ffa826131ce565b155b1561200957506000611f97565b82518251808203611f8f878263ffffffff6131d916565b600061202b836131ce565b158061203d575061203b826131ce565b155b1561204a57506000611f97565b825182518061205e57600092505050611f97565b808204611f8f878263ffffffff6131d916565b600061207c836131ce565b158061208e575061208c826131ce565b155b1561209b57506000611f97565b82518251806120af57600092505050611f97565b808205611f8f878263ffffffff6131d916565b60006120cd836131ce565b15806120df57506120dd826131ce565b155b156120ec57506000611f97565b825182518061210057600092505050611f97565b808206611f8f878263ffffffff6131d916565b600061211e836131ce565b1580612130575061212e826131ce565b155b1561213d57506000611f97565b825182518061215157600092505050611f97565b808207611f8f878263ffffffff6131d916565b600061216f846131ce565b1580612181575061217f836131ce565b155b1561218e575060006121c6565b835183518351806121a557600093505050506121c6565b60008183850890506121bd898263ffffffff6131d916565b60019450505050505b949350505050565b60006121d9846131ce565b15806121eb57506121e9836131ce565b155b156121f8575060006121c6565b8351835183518061220f57600093505050506121c6565b60008183850990506121bd898263ffffffff6131d916565b6000612232836131ce565b15806122445750612242826131ce565b155b1561225157506000611f97565b8251825180820a611f8f878263ffffffff6131d916565b6000612273836131ce565b15806122855750612283826131ce565b155b1561229257506000611f97565b82518251808210611f8f878263ffffffff6131d916565b60006122b4836131ce565b15806122c657506122c4826131ce565b155b156122d357506000611f97565b82518251808211611f8f878263ffffffff6131d916565b60006122f5836131ce565b15806123075750612305826131ce565b155b1561231457506000611f97565b82518251808212611f8f878263ffffffff6131d916565b6000612336836131ce565b15806123485750612346826131ce565b155b1561235557506000611f97565b82518251808213611f8f878263ffffffff6131d916565b600061238e6118ee61237d84613098565b5161238786613098565b51146131ed565b5060019392505050565b60006123a3826131ce565b6123bd576123b883600063ffffffff6131d916565b6123d4565b815180156123d1858263ffffffff6131d916565b50505b5060015b92915050565b60006123e9836131ce565b15806123fb57506123f9826131ce565b155b1561240857506000611f97565b82518251808216611f8f878263ffffffff6131d916565b600061242a836131ce565b158061243c575061243a826131ce565b155b1561244957506000611f97565b82518251808217611f8f878263ffffffff6131d916565b600061246b836131ce565b158061247d575061247b826131ce565b155b1561248a57506000611f97565b82518251808218611f8f878263ffffffff6131d916565b60006124ac826131ce565b6124b8575060006123d8565b815180196124cc858263ffffffff6131d916565b506001949350505050565b60006124e2836131ce565b15806124f457506124f2826131ce565b155b1561250157506000611f97565b8251825181811a611f8f878263ffffffff6131d916565b6000612523836131ce565b15806125355750612533826131ce565b155b1561254257506000611f97565b8251825181810b611f8f878263ffffffff6131d916565b60006123d461256783613098565b51849063ffffffff6131d916565b60006123d461258383613216565b849063ffffffff61284b16565b600192915050565b60006125b182608001518361329f90919063ffffffff16565b506001919050565b60006125b182606001518361329f90919063ffffffff16565b60006125dd82613098565b606084015250600192915050565b60006125f682613098565b835250600192915050565b600061260c836132ad565b61261857506000611f97565b612621826131ce565b61262d57506000611f97565b81511561238e5761263d83613098565b84525060019392505050565b60006125b161266e61266161265c6132ba565b613098565b51602085015151146131ed565b839063ffffffff61284b16565b60006123d4838363ffffffff61329f16565b60006123d4838363ffffffff61282916565b6000806126aa613f79565b845184106126ff576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b6000849050600086828151811061271257fe5b016020015160019092019160f81c9050600061272c613fa7565b60ff83166127605761273e8985613337565b909450915060008461274f84612cc8565b919850965094506128229350505050565b60ff83166001141561278757612776898561335e565b909450905060008461274f836134cb565b60ff8316600214156127ae5761279d8985613337565b909450915060008461274f8461352b565b600360ff8416108015906127c55750600c60ff8416105b15612802576002198301606060006127de838d896135a9565b9098509250905080876127f084612d46565b99509950995050505050505050612822565b8260ff166127100160006128166000612cc8565b91985096509450505050505b9250925092565b61283f826040015161283a83613098565b613664565b82604001819052505050565b61285c826020015161283a83613098565b82602001819052505050565b60006125b161266e61287b61265c6132ba565b51604085015151146131ed565b60006125b18260a001518361329f90919063ffffffff16565b60006128ac826132ad565b6128b8575060006123d8565b6128c182613098565b60a084015250600192915050565b60006128e1838363ffffffff61284b16565b6123d4838363ffffffff61284b16565b6000612903848363ffffffff61284b16565b612913848463ffffffff61284b16565b61238e848363ffffffff61284b16565b6000612935858363ffffffff61284b16565b612945858463ffffffff61284b16565b612955858563ffffffff61284b16565b6124cc858363ffffffff61284b16565b6000612913848463ffffffff61284b16565b6000612989858563ffffffff61284b16565b612955858463ffffffff61284b16565b60006129a4836131ce565b15806129b657506129b48261371a565b155b156129c357506000611f97565b6129cc82613729565b60ff168360000151106129e157506000611f97565b61238e82604001518460000151815181106129f857fe5b60200260200101518561284b90919063ffffffff16565b6000612a1a8361371a565b1580612a2c5750612a2a846131ce565b155b15612a39575060006121c6565b612a4283613729565b60ff16846000015110612a57575060006121c6565b818360400151856000015181518110612a6c57fe5b60209081029190910101526124cc858463ffffffff61284b16565b6000612a928261371a565b612a9e575060006123d8565b6123d4612aaa83613729565b849060ff1663ffffffff6131d916565b50600190565b600080612acb613fce565b612ad484613098565b51600193509150505b9250929050565b6000806000806000806000612af88861371a565b612b0c576000965094509092509050612cbf565b612b2d8860400151600181518110612b2057fe5b60200260200101516131ce565b612b41576000965094509092509050612cbf565b612b558860400151600281518110612b2057fe5b612b69576000965094509092509050612cbf565b612b7d8860400151600381518110612b2057fe5b612b91576000965094509092509050612cbf565b8760400151600381518110612ba257fe5b60200260200101516000015160001b92508760400151600281518110612bc457fe5b602002602001015160000151915073__$9836fa7140e5a33041d4b827682e675a30$__624c28f6612bf48a613098565b6000015185858c60400151600181518110612c0b57fe5b6020026020010151600001516040518563ffffffff1660e01b815260040180858152602001846001600160581b0319166001600160581b0319168152602001838152602001826001600160a01b03166001600160a01b0316815260200194505050505060206040518083038186803b158015612c8657600080fd5b505af4158015612c9a573d6000803e3d6000fd5b505050506040513d6020811015612cb057600080fd5b50516001975095509193509150505b92959194509250565b612cd0613f79565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612d35565b612d22613f79565b815260200190600190039081612d1a5790505b508152600060209091015292915050565b612d4e613f79565b612d588251613738565b612da9576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b8051600090612e0484613098565b511415612e58576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b61238e848363ffffffff61329f16565b600260c090910152565b600160c090910152565b600060028260c001511415612e9357506000611b09565b60018260c001511415612ea857506001611b09565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e09093019092528151910120611b09565b600060c090910152565b600080612f26613f18565b612f2e613f18565b600060c08201819052612f41878761373f565b8452965090508015612f595793508492509050612822565b612f63878761373f565b6020850152965090508015612f7e5793508492509050612822565b612f88878761373f565b6040850152965090508015612fa35793508492509050612822565b612fad878761373f565b6060850152965090508015612fc85793508492509050612822565b612fd2878761373f565b6080850152965090508015612fed5793508492509050612822565b612ff7878761373f565b60a08501529650905080156130125793508492509050612822565b506000969495509392505050565b613028613f18565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b600080600061308f8460ff1661377d565b50949350505050565b6130a0613fce565b6060820151600c60ff909116106130f2576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661311f5760405180602001604052806131168460000151613c24565b90529050611b09565b606082015160ff1660011415613166576040518060200160405280613116846020015160000151856020015160400151866020015160600151876020015160200151613c48565b606082015160ff166002141561318b5750604080516020810190915281518152611b09565b600360ff16826060015160ff16101580156131af57506060820151600c60ff909116105b156131cc5760405180602001604052806131168460400151613cf0565bfe5b6060015160ff161590565b61285c826020015161283a61265c84612cc8565b6131f5613f79565b811561320c576132056001612cc8565b9050611b09565b6132056000612cc8565b61321e613f79565b816060015160ff16600214156132655760405162461bcd60e51b815260040180806020018281038252602181526020018061404f6021913960400191505060405180910390fd5b606082015160ff1661327b576132056000612cc8565b816060015160ff1660011415613295576132056001612cc8565b6132056003612cc8565b61285c826020015182613664565b6060015160ff1660011490565b6132c2613f79565b604080516080808201835260008083528351918201845280825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191613327565b613314613f79565b81526020019060019003908161330c5790505b5081526003602090910152905090565b600080828161334c868363ffffffff613e3c16565b60209290920196919550909350505050565b6000613368613fa7565b6000839050600085828151811061337b57fe5b602001015160f81c60f81b60f81c9050818060010192505060008683815181106133a157fe5b016020015160019384019360f89190911c915060009060ff8416141561343f5760006133cb613f79565b6133d58a8761269f565b90975090925090508115613430576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b61343981613098565b51925050505b6000613451898663ffffffff613e3c16565b90506020850194508360ff1660011415613496576040805160808101825260ff909416845260208401919091526001908301526060820152919350909150612add9050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6134d3613f79565b60408051608081018252600080825260208083018690528351828152908101845291928301919061351a565b613507613f79565b8152602001906001900390816134ff5790505b508152600160209091015292915050565b613533613f79565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191613598565b613585613f79565b81526020019060019003908161357d5790505b508152600260209091015292915050565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156135f457816020015b6135e1613f79565b8152602001906001900390816135d95790505b50905060005b8960ff168160ff16101561364e57613612898561269f565b8451859060ff861690811061362357fe5b60209081029190910101529450925082156136465750909450909250905061365b565b6001016135fa565b5060009550919350909150505b93509350939050565b61366c613fce565b6040805160028082526060828101909352816020015b61368a613fce565b81526020019060019003908161368257905050905082816000815181106136ad57fe5b602002602001018190525083816001815181106136c657fe5b6020026020010181905250604051806020016040528061371060405180604001604052806136f7886000015161352b565b8152602001613709896000015161352b565b9052613e58565b9052949350505050565b60006123d88260600151613ed7565b60006123d88260600151613ef5565b6008101590565b60008061374a613fce565b83600061375d878363ffffffff613e3c16565b604080516020808201909252918252600099930197509550909350505050565b60008060018314156137955750600290506001613c1f565b60028314156137aa5750600290506001613c1f565b60038314156137bf5750600290506001613c1f565b60048314156137d45750600290506001613c1f565b60058314156137e95750600290506001613c1f565b60068314156137fe5750600290506001613c1f565b60078314156138135750600290506001613c1f565b60088314156138285750600390506001613c1f565b600983141561383d5750600390506001613c1f565b600a8314156138525750600290506001613c1f565b60108314156138675750600290506001613c1f565b601183141561387c5750600290506001613c1f565b60128314156138915750600290506001613c1f565b60138314156138a65750600290506001613c1f565b60148314156138bb5750600290506001613c1f565b60158314156138cf57506001905080613c1f565b60168314156138e45750600290506001613c1f565b60178314156138f95750600290506001613c1f565b601883141561390e5750600290506001613c1f565b601983141561392257506001905080613c1f565b601a8314156139375750600290506001613c1f565b601b83141561394c5750600290506001613c1f565b602083141561396057506001905080613c1f565b602183141561397457506001905080613c1f565b60308314156139895750600190506000613c1f565b603183141561399e5750600090506001613c1f565b60328314156139b35750600090506001613c1f565b60338314156139c85750600190506000613c1f565b60348314156139dd5750600190506000613c1f565b60358314156139f25750600290506000613c1f565b6036831415613a075750600090506001613c1f565b6037831415613a1c5750600090506001613c1f565b6038831415613a315750600190506000613c1f565b6039831415613a465750600090506001613c1f565b603a831415613a5b5750600090506001613c1f565b603b831415613a6f57506000905080613c1f565b603c831415613a845750600090506001613c1f565b603d831415613a995750600190506000613c1f565b6040831415613aae5750600190506002613c1f565b6041831415613ac35750600290506003613c1f565b6042831415613ad85750600390506004613c1f565b6043831415613aec57506002905080613c1f565b6044831415613b0057506003905080613c1f565b6050831415613b155750600290506001613c1f565b6051831415613b2a5750600390506001613c1f565b6052831415613b3e57506001905080613c1f565b6060831415613b5257506000905080613c1f565b6061831415613b675750600190506000613c1f565b6070831415613b7c5750600190506000613c1f565b6071831415613b9057506001905080613c1f565b6072831415613ba55750600090506001613c1f565b6073831415613bb957506001905080613c1f565b6074831415613bcd57506000905080613c1f565b6075831415613be157506000905080613c1f565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613ca2575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206121c6565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115613d40576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613d6d578160200160208202803883390190505b50805190915060005b81811015613dc957613d86613fce565b613da2868381518110613d9557fe5b6020026020010151613098565b90508060000151848381518110613db557fe5b602090810291909101015250600101613d76565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613e12578181015183820152602001613dfa565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60008160200183511015613e4f57600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b613e7b613f79565b815260200190600190039081613e73575050805190915060005b81811015613ecd57848160028110613ea957fe5b6020020151838281518110613eba57fe5b6020908102919091010152600101613e95565b506121c682613cf0565b6000600c60ff83161080156123d8575050600360ff91909116101590565b6000613f0082613ed7565b15613f1057506002198101611b09565b506001611b09565b6040518060e00160405280613f2b613fce565b8152602001613f38613fce565b8152602001613f45613fce565b8152602001613f52613fce565b8152602001613f5f613fce565b8152602001613f6c613fce565b8152602001600081525090565b604051806080016040528060008152602001613f93613fa7565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4f6e6520737465702070726f6f66207769746820696e76616c6964207072657620737461746550726f6f6620686164206e6f6e206d61746368696e672073746172742073746174654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f73656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d657361676550726f6f66206f6620617578706f702068616420626164206175782076616c756543616e206f6e6c79206f6e6520737465702070726f6f6620666f6c6c6f77696e6720612073696e676c652073746570206368616c6c656e6765a265627a7a723158208f852fed63860a148d382fa1d309236d27e3d085ecf45f08857bb35f4225ae2164736f6c634300050c0032"

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

// ValidateProof is a free data retrieval call binding the contract method 0x97ce6508.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", fields, timeBounds, tokenTypes, messageValue, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0x97ce6508.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, tokenTypes, messageValue, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0x97ce6508.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, tokenTypes, messageValue, proof)
}
