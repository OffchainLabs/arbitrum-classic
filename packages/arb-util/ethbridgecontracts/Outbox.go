// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OutboxMetaData contains all meta data concerning the Outbox contract.
var OutboxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outboxEntryIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outboxEntryIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numInBatch\",\"type\":\"uint256\"}],\"name\":\"OutboxEntryCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OUTBOX_VERSION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"calculateItemHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"item\",\"type\":\"bytes32\"}],\"name\":\"calculateMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isNitroReady\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1BatchNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1OutputId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outboxEntries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"}],\"name\":\"outboxEntryExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"processOutgoingMessages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"newBridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506000805460ff1916600117905561145a8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106101105760003560e01c80638dd14802116100ad578063b0f3053711610071578063b0f3053714610501578063c75184df14610509578063cb23bcb51461052d578063e78cea9214610535578063f1fd3a391461053d57610110565b80638dd148021461031b5780639229bab6146103415780639c5cfe0b1461035e5780639f0c04bf1461045a578063a8929e0b146104f957610110565b80627436d3146101155780630c726847146101cd578063119852711461028d5780634654779014610295578063485cc9551461029d5780636f791d29146102cb57806372f2a8c7146102e757806380648b02146102ef5780638515bc6a14610313575b600080fd5b6101bb6004803603606081101561012b57600080fd5b810190602081018135600160201b81111561014557600080fd5b82018360208201111561015757600080fd5b803590602001918460208302840111600160201b8311171561017857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550508235935050506020013561055a565b60408051918252519081900360200190f35b61028b600480360360408110156101e357600080fd5b810190602081018135600160201b8111156101fd57600080fd5b82018360208201111561020f57600080fd5b803590602001918460018302840111600160201b8311171561023057600080fd5b919390929091602081019035600160201b81111561024d57600080fd5b82018360208201111561025f57600080fd5b803590602001918460208302840111600160201b8311171561028057600080fd5b509092509050610595565b005b6101bb610681565b6101bb610697565b61028b600480360360408110156102b357600080fd5b506001600160a01b03813581169160200135166106a6565b6102d3610730565b604080519115158252519081900360200190f35b6101bb610739565b6102f761073f565b604080516001600160a01b039092168252519081900360200190f35b6101bb61074e565b61028b6004803603602081101561033157600080fd5b50356001600160a01b0316610764565b6101bb6004803603602081101561035757600080fd5b503561084f565b61028b600480360361014081101561037557600080fd5b81359190810190604081016020820135600160201b81111561039657600080fd5b8201836020820111156103a857600080fd5b803590602001918460208302840111600160201b831117156103c957600080fd5b919390928235926001600160a01b03602082013581169360408301359091169260608301359260808101359260a08201359260c08301359261010081019060e00135600160201b81111561041c57600080fd5b82018360208201111561042e57600080fd5b803590602001918460018302840111600160201b8311171561044f57600080fd5b509092509050610861565b6101bb600480360360e081101561047057600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b8111156104bb57600080fd5b8201836020820111156104cd57600080fd5b803590602001918460018302840111600160201b831117156104ee57600080fd5b509092509050610c39565b6101bb610cd9565b6101bb610cdf565b610511610cee565b604080516001600160801b039092168252519081900360200190f35b6102f7610cf3565b6102f7610d07565b6102d36004803603602081101561055357600080fd5b5035610d16565b600061058d8484846040516020018082815260200191505060405160208183030381529060405280519060200120610d2a565b949350505050565b60005461010090046001600160a01b031633146105e7576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b806000805b828110156106785761065687838888888681811061060657fe5b9050602002013586019261061c939291906113fc565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610df892505050565b84848281811061066257fe5b60200291909101359290920191506001016105ec565b50505050505050565b600454600160801b90046001600160801b031690565b6003546001600160801b031690565b60005461010090046001600160a01b0316156106f8576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b60008054610100600160a81b0319166101006001600160a01b0394851602179055600180546001600160a01b03191691909216179055565b60005460ff1690565b60055490565b6006546001600160a01b031690565b600354600160801b90046001600160801b031690565b600160009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107b257600080fd5b505afa1580156107c6573d6000803e3d6000fd5b505050506040513d60208110156107dc57600080fd5b50516001600160a01b0316331461082d576040805162461bcd60e51b815260206004820152601060248201526f2727aa2fa12924a223a2afa7aba722a960811b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b60026020526000908152604090205481565b6000806108748a8a8a8a8a8a8a8a610c39565b90506108c38e8e8e80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508d84610f60565b91508d8a6001600160a01b03168a6001600160a01b03167f20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab189648e6040518082815260200191505060405180910390a45061091a6113b5565b60036040518060c00160405290816000820160009054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016000820160109054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016001820160009054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016001820160109054906101000a90046001600160801b03166001600160801b03166001600160801b03168152602001600282015481526020016003820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152505090506040518060c00160405280896001600160801b03168152602001886001600160801b03168152602001876001600160801b031681526020018f6001600160801b031681526020018381526020018b6001600160a01b0316815250600360008201518160000160006101000a8154816001600160801b0302191690836001600160801b0316021790555060208201518160000160106101000a8154816001600160801b0302191690836001600160801b0316021790555060408201518160010160006101000a8154816001600160801b0302191690836001600160801b0316021790555060608201518160010160106101000a8154816001600160801b0302191690836001600160801b031602179055506080820151816002015560a08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550905050610ba9898686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061113f92505050565b80516003805460208401516001600160801b03199182166001600160801b03948516178416600160801b9185168202179092556040840151600480546060870151931691851691909117841691909316909102179055608081015160055560a00151600680546001600160a01b0319166001600160a01b0390921691909117905550505050505050505050505050565b600060038960601b60601c6001600160a01b03168960601b60601c6001600160a01b0316898989898989604051602001808a60ff1660ff1660f81b815260010189815260200188815260200187815260200186815260200185815260200184815260200183838082843780830192505050995050505050505050505060405160208183030381529060405280519060200120905098975050505050505050565b61a4b490565b6004546001600160801b031690565b600181565b60005461010090046001600160a01b031681565b6001546001600160a01b031681565b600090815260026020526040902054151590565b8251600090610100811115610d3e57600080fd5b8260005b82811015610dee5760028606610d9b57868181518110610d5e57fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150610de0565b81878281518110610da857fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101610d42565b5095945050505050565b805160009082908290610e0757fe5b01602001516001600160f81b0319161415610f5d578051606114610e5f576040805162461bcd60e51b815260206004820152600a6024820152690848288be988a9c8ea8960b31b604482015290519081900360640190fd5b6000610e7282600163ffffffff61135c16565b9050610e7d81610d16565b15610ec6576040805162461bcd60e51b8152602060048201526014602482015273454e5452595f414c52454144595f45584953545360601b604482015290519081900360640190fd5b6000610ed983602163ffffffff61135c16565b90506000610eee84604163ffffffff61135c16565b9050610ef86113ea565b5060408051602080820183528382526000868152600282528390208251905582518681529081018490528083018590529151909185917fe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c1319181900360600190a2505050505b50565b6000610100845110610faa576040805162461bcd60e51b815260206004820152600e60248201526d50524f4f465f544f4f5f4c4f4e4760901b604482015290519081900360640190fd5b835160020a8310610ff5576040805162461bcd60e51b815260206004820152601060248201526f1410551217d393d517d352539253505360821b604482015290519081900360640190fd5b600061100285858561055a565b6000878152600260205260409020805491925090611059576040805162461bcd60e51b815260206004820152600f60248201526e4e4f5f4f5554424f585f454e54525960881b604482015290519081900360640190fd5b8551604080516020808201899052818301939093528151808203830181526060909101825280519083012060008181526001850190935291205460ff16156110d8576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b81548314611118576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b6000818152600192830160205260409020805460ff19169092179091559695505050505050565b600154604051639e5d4c4960e01b81526001600160a01b03858116600483019081526024830186905260606044840181815286516064860152865160009692959490921693639e5d4c49938a938a938a93909160849091019060208501908083838e5b838110156111ba5781810151838201526020016111a2565b50505050905090810190601f1680156111e75780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561120857600080fd5b505af115801561121c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561124557600080fd5b815160208301805160405192949293830192919084600160201b82111561126b57600080fd5b90830190602082018581111561128057600080fd5b8251600160201b81118282018810171561129957600080fd5b82525081516020918201929091019080838360005b838110156112c65781810151838201526020016112ae565b50505050905090810190601f1680156112f35780820380516001836020036101000a031916815260200191505b506040525050509150915081611355578051156113135780518082602001fd5b6040805162461bcd60e51b81526020600482015260126024820152711094925111d157d0d0531317d1905253115160721b604482015290519081900360640190fd5b5050505050565b600081602001835110156113ac576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b60408051602081019091526000815290565b6000808585111561140b578182fd5b83861115611417578182fd5b505082019391909203915056fea2646970667358221220b8a07fdd4002ff00965732e20a6a2b65fe9c9f128e8325843424ab200c3aff1a64736f6c634300060b0033",
}

// OutboxABI is the input ABI used to generate the binding from.
// Deprecated: Use OutboxMetaData.ABI instead.
var OutboxABI = OutboxMetaData.ABI

// OutboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutboxMetaData.Bin instead.
var OutboxBin = OutboxMetaData.Bin

// DeployOutbox deploys a new Ethereum contract, binding an instance of Outbox to it.
func DeployOutbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Outbox, error) {
	parsed, err := OutboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// Outbox is an auto generated Go binding around an Ethereum contract.
type Outbox struct {
	OutboxCaller     // Read-only binding to the contract
	OutboxTransactor // Write-only binding to the contract
	OutboxFilterer   // Log filterer for contract events
}

// OutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxSession struct {
	Contract     *Outbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxCallerSession struct {
	Contract *OutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxTransactorSession struct {
	Contract     *OutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxRaw struct {
	Contract *Outbox // Generic contract binding to access the raw methods on
}

// OutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxCallerRaw struct {
	Contract *OutboxCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxTransactorRaw struct {
	Contract *OutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutbox creates a new instance of Outbox, bound to a specific deployed contract.
func NewOutbox(address common.Address, backend bind.ContractBackend) (*Outbox, error) {
	contract, err := bindOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// NewOutboxCaller creates a new read-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxCaller(address common.Address, caller bind.ContractCaller) (*OutboxCaller, error) {
	contract, err := bindOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxCaller{contract: contract}, nil
}

// NewOutboxTransactor creates a new write-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxTransactor, error) {
	contract, err := bindOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxTransactor{contract: contract}, nil
}

// NewOutboxFilterer creates a new log filterer instance of Outbox, bound to a specific deployed contract.
func NewOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxFilterer, error) {
	contract, err := bindOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxFilterer{contract: contract}, nil
}

// bindOutbox binds a generic wrapper to an already deployed contract.
func bindOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.OutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transact(opts, method, params...)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_Outbox *OutboxCaller) OUTBOXVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "OUTBOX_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_Outbox *OutboxSession) OUTBOXVERSION() (*big.Int, error) {
	return _Outbox.Contract.OUTBOXVERSION(&_Outbox.CallOpts)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_Outbox *OutboxCallerSession) OUTBOXVERSION() (*big.Int, error) {
	return _Outbox.Contract.OUTBOXVERSION(&_Outbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Outbox *OutboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Outbox *OutboxSession) Bridge() (common.Address, error) {
	return _Outbox.Contract.Bridge(&_Outbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Outbox *OutboxCallerSession) Bridge() (common.Address, error) {
	return _Outbox.Contract.Bridge(&_Outbox.CallOpts)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) pure returns(bytes32)
func (_Outbox *OutboxCaller) CalculateItemHash(opts *bind.CallOpts, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "calculateItemHash", l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) pure returns(bytes32)
func (_Outbox *OutboxSession) CalculateItemHash(l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateItemHash(&_Outbox.CallOpts, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) pure returns(bytes32)
func (_Outbox *OutboxCallerSession) CalculateItemHash(l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateItemHash(&_Outbox.CallOpts, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxCaller) CalculateMerkleRoot(opts *bind.CallOpts, proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "calculateMerkleRoot", proof, path, item)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateMerkleRoot(&_Outbox.CallOpts, proof, path, item)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxCallerSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateMerkleRoot(&_Outbox.CallOpts, proof, path, item)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Outbox *OutboxCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Outbox *OutboxSession) IsMaster() (bool, error) {
	return _Outbox.Contract.IsMaster(&_Outbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Outbox *OutboxCallerSession) IsMaster() (bool, error) {
	return _Outbox.Contract.IsMaster(&_Outbox.CallOpts)
}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_Outbox *OutboxCaller) IsNitroReady(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "isNitroReady")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_Outbox *OutboxSession) IsNitroReady() (*big.Int, error) {
	return _Outbox.Contract.IsNitroReady(&_Outbox.CallOpts)
}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_Outbox *OutboxCallerSession) IsNitroReady() (*big.Int, error) {
	return _Outbox.Contract.IsNitroReady(&_Outbox.CallOpts)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1BatchNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1BatchNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1BatchNum() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1BatchNum(&_Outbox.CallOpts)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1BatchNum() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1BatchNum(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1EthBlock() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1EthBlock(&_Outbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1EthBlock(&_Outbox.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_Outbox *OutboxCaller) L2ToL1OutputId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1OutputId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_Outbox *OutboxSession) L2ToL1OutputId() ([32]byte, error) {
	return _Outbox.Contract.L2ToL1OutputId(&_Outbox.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_Outbox *OutboxCallerSession) L2ToL1OutputId() ([32]byte, error) {
	return _Outbox.Contract.L2ToL1OutputId(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCallerSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// OutboxEntries is a free data retrieval call binding the contract method 0x9229bab6.
//
// Solidity: function outboxEntries(uint256 ) view returns(bytes32 root)
func (_Outbox *OutboxCaller) OutboxEntries(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "outboxEntries", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OutboxEntries is a free data retrieval call binding the contract method 0x9229bab6.
//
// Solidity: function outboxEntries(uint256 ) view returns(bytes32 root)
func (_Outbox *OutboxSession) OutboxEntries(arg0 *big.Int) ([32]byte, error) {
	return _Outbox.Contract.OutboxEntries(&_Outbox.CallOpts, arg0)
}

// OutboxEntries is a free data retrieval call binding the contract method 0x9229bab6.
//
// Solidity: function outboxEntries(uint256 ) view returns(bytes32 root)
func (_Outbox *OutboxCallerSession) OutboxEntries(arg0 *big.Int) ([32]byte, error) {
	return _Outbox.Contract.OutboxEntries(&_Outbox.CallOpts, arg0)
}

// OutboxEntryExists is a free data retrieval call binding the contract method 0xf1fd3a39.
//
// Solidity: function outboxEntryExists(uint256 batchNum) view returns(bool)
func (_Outbox *OutboxCaller) OutboxEntryExists(opts *bind.CallOpts, batchNum *big.Int) (bool, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "outboxEntryExists", batchNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OutboxEntryExists is a free data retrieval call binding the contract method 0xf1fd3a39.
//
// Solidity: function outboxEntryExists(uint256 batchNum) view returns(bool)
func (_Outbox *OutboxSession) OutboxEntryExists(batchNum *big.Int) (bool, error) {
	return _Outbox.Contract.OutboxEntryExists(&_Outbox.CallOpts, batchNum)
}

// OutboxEntryExists is a free data retrieval call binding the contract method 0xf1fd3a39.
//
// Solidity: function outboxEntryExists(uint256 batchNum) view returns(bool)
func (_Outbox *OutboxCallerSession) OutboxEntryExists(batchNum *big.Int) (bool, error) {
	return _Outbox.Contract.OutboxEntryExists(&_Outbox.CallOpts, batchNum)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Outbox *OutboxCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Outbox *OutboxSession) Rollup() (common.Address, error) {
	return _Outbox.Contract.Rollup(&_Outbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Outbox *OutboxCallerSession) Rollup() (common.Address, error) {
	return _Outbox.Contract.Rollup(&_Outbox.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x9c5cfe0b.
//
// Solidity: function executeTransaction(uint256 batchNum, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, batchNum *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransaction", batchNum, proof, index, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x9c5cfe0b.
//
// Solidity: function executeTransaction(uint256 batchNum, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxSession) ExecuteTransaction(batchNum *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, batchNum, proof, index, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x9c5cfe0b.
//
// Solidity: function executeTransaction(uint256 batchNum, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransaction(batchNum *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, batchNum, proof, index, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rollup, address _bridge) returns()
func (_Outbox *OutboxTransactor) Initialize(opts *bind.TransactOpts, _rollup common.Address, _bridge common.Address) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "initialize", _rollup, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rollup, address _bridge) returns()
func (_Outbox *OutboxSession) Initialize(_rollup common.Address, _bridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.Initialize(&_Outbox.TransactOpts, _rollup, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rollup, address _bridge) returns()
func (_Outbox *OutboxTransactorSession) Initialize(_rollup common.Address, _bridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.Initialize(&_Outbox.TransactOpts, _rollup, _bridge)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactor) ProcessOutgoingMessages(opts *bind.TransactOpts, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "processOutgoingMessages", sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactorSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address newBridge) returns()
func (_Outbox *OutboxTransactor) SetBridge(opts *bind.TransactOpts, newBridge common.Address) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "setBridge", newBridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address newBridge) returns()
func (_Outbox *OutboxSession) SetBridge(newBridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.SetBridge(&_Outbox.TransactOpts, newBridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address newBridge) returns()
func (_Outbox *OutboxTransactorSession) SetBridge(newBridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.SetBridge(&_Outbox.TransactOpts, newBridge)
}

// OutboxOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the Outbox contract.
type OutboxOutBoxTransactionExecutedIterator struct {
	Event *OutboxOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *OutboxOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxOutBoxTransactionExecuted)
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
		it.Event = new(OutboxOutBoxTransactionExecuted)
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
func (it *OutboxOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the Outbox contract.
type OutboxOutBoxTransactionExecuted struct {
	DestAddr         common.Address
	L2Sender         common.Address
	OutboxEntryIndex *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxEntryIndex, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, destAddr []common.Address, l2Sender []common.Address, outboxEntryIndex []*big.Int) (*OutboxOutBoxTransactionExecutedIterator, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var outboxEntryIndexRule []interface{}
	for _, outboxEntryIndexItem := range outboxEntryIndex {
		outboxEntryIndexRule = append(outboxEntryIndexRule, outboxEntryIndexItem)
	}

	logs, sub, err := _Outbox.contract.FilterLogs(opts, "OutBoxTransactionExecuted", destAddrRule, l2SenderRule, outboxEntryIndexRule)
	if err != nil {
		return nil, err
	}
	return &OutboxOutBoxTransactionExecutedIterator{contract: _Outbox.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxEntryIndex, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *OutboxOutBoxTransactionExecuted, destAddr []common.Address, l2Sender []common.Address, outboxEntryIndex []*big.Int) (event.Subscription, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var outboxEntryIndexRule []interface{}
	for _, outboxEntryIndexItem := range outboxEntryIndex {
		outboxEntryIndexRule = append(outboxEntryIndexRule, outboxEntryIndexItem)
	}

	logs, sub, err := _Outbox.contract.WatchLogs(opts, "OutBoxTransactionExecuted", destAddrRule, l2SenderRule, outboxEntryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxOutBoxTransactionExecuted)
				if err := _Outbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxEntryIndex, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*OutboxOutBoxTransactionExecuted, error) {
	event := new(OutboxOutBoxTransactionExecuted)
	if err := _Outbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxOutboxEntryCreatedIterator is returned from FilterOutboxEntryCreated and is used to iterate over the raw logs and unpacked data for OutboxEntryCreated events raised by the Outbox contract.
type OutboxOutboxEntryCreatedIterator struct {
	Event *OutboxOutboxEntryCreated // Event containing the contract specifics and raw log

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
func (it *OutboxOutboxEntryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxOutboxEntryCreated)
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
		it.Event = new(OutboxOutboxEntryCreated)
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
func (it *OutboxOutboxEntryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxOutboxEntryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxOutboxEntryCreated represents a OutboxEntryCreated event raised by the Outbox contract.
type OutboxOutboxEntryCreated struct {
	BatchNum         *big.Int
	OutboxEntryIndex *big.Int
	OutputRoot       [32]byte
	NumInBatch       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutboxEntryCreated is a free log retrieval operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxEntryIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) FilterOutboxEntryCreated(opts *bind.FilterOpts, batchNum []*big.Int) (*OutboxOutboxEntryCreatedIterator, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Outbox.contract.FilterLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return &OutboxOutboxEntryCreatedIterator{contract: _Outbox.contract, event: "OutboxEntryCreated", logs: logs, sub: sub}, nil
}

// WatchOutboxEntryCreated is a free log subscription operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxEntryIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) WatchOutboxEntryCreated(opts *bind.WatchOpts, sink chan<- *OutboxOutboxEntryCreated, batchNum []*big.Int) (event.Subscription, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Outbox.contract.WatchLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxOutboxEntryCreated)
				if err := _Outbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
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

// ParseOutboxEntryCreated is a log parse operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxEntryIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) ParseOutboxEntryCreated(log types.Log) (*OutboxOutboxEntryCreated, error) {
	event := new(OutboxOutboxEntryCreated)
	if err := _Outbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
