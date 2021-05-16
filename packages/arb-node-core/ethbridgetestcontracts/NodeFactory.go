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

// NodeFactoryABI is the input ABI used to generate the binding from.
const NodeFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"createNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"templateContract\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeFactoryBin is the compiled bytecode used for deploying new contracts.
var NodeFactoryBin = "0x608060405234801561001057600080fd5b5060405161001d9061005f565b604051809103906000f080158015610039573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b039290921691909117905561006c565b6108e78061046283390190565b6103e78061007b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806372be06d81461003b578063d45ab2b51461005f575b600080fd5b610043610094565b604080516001600160a01b039092168252519081900360200190f35b610043600480360360a081101561007557600080fd5b50803590602081013590604081013590606081013590608001356100a3565b6000546001600160a01b031681565b6000805481906100bb906001600160a01b031661014c565b60408051632901acdd60e21b8152336004820152602481018a905260448101899052606481018890526084810187905260a4810186905290519192506001600160a01b0383169163a406b3749160c48082019260009290919082900301818387803b15801561012957600080fd5b505af115801561013d573d6000803e3d6000fd5b50929998505050505050505050565b600081610161816001600160a01b031661030a565b604051806040016040528060188152602001772727afa1a7a72a2920a1aa2fa1a627a722afa6a0a9aa22a960411b8152509061021b5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156101e05781810151838201526020016101c8565b50505050905090810190601f16801561020d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561025557600080fd5b505afa158015610269573d6000803e3d6000fd5b505050506040513d602081101561027f57600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b6020820152906102f05760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101e05781810151838201526020016101c8565b50610303836001600160a01b0316610314565b9392505050565b803b15155b919050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b03811661030f576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fdfea26469706673582212206f9f1b28b843fd4a9581cd9f13557e7ec6aeb1fbc84bd049ffa90bea6a073b4364736f6c634300060b0033608060405234801561001057600080fd5b506000805460ff191660011790556108ba8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806388d221c6116100b8578063a406b3741161007c578063a406b37414610260578063cb23bcb5146102a4578063d7ff5e35146102c8578063da83b48e146102d0578063dff69787146102d8578063f0dd77ff146102e057610137565b806388d221c6146101fc5780639168ae721461020457806396a9fdc01461022a57806397bdc51014610250578063a0369c141461025857610137565b80635b8b2280116100ff5780635b8b2280146101ab5780636971dfe5146101b35780636f791d29146101d0578063701da98e146101ec57806383197ef0146101f457610137565b80631bc09d0a1461013c5780632466696e1461015b5780632edfb42a146101935780633aa192741461019b578063479c9254146101a3575b600080fd5b6101596004803603602081101561015257600080fd5b50356102e8565b005b6101816004803603602081101561017157600080fd5b50356001600160a01b0316610346565b60408051918252519081900360200190f35b610181610429565b61015961042f565b61018161047b565b610181610481565b610159600480360360208110156101c957600080fd5b5035610487565b6101d86104d9565b604080519115158252519081900360200190f35b6101816104e2565b6101596104e8565b61015961053e565b6101d86004803603602081101561021a57600080fd5b50356001600160a01b0316610587565b6101596004803603602081101561024057600080fd5b50356001600160a01b031661059c565b61018161066e565b610181610674565b610159600480360360c081101561027657600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a0013561067a565b6102ac61074d565b604080516001600160a01b039092168252519081900360200190f35b61018161075c565b610159610762565b6101816107bb565b6101816107c1565b6009546001600160a01b03163314610335576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b600a546103415743600a555b600b55565b6009546000906001600160a01b03163314610396576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b6001600160a01b03821660009081526008602052604090205460ff16156103f5576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b506001600160a01b03166000908152600860205260409020805460ff19166001908117909155600780549091019081905590565b60055481565b600654431015610479576040805162461bcd60e51b815260206004820152601060248201526f10d212531117d513d3d7d49150d1539560821b604482015290519081900360640190fd5b565b60045481565b60025481565b6009546001600160a01b031633146104d4576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b600655565b60005460ff1690565b60015481565b6009546001600160a01b03163314610535576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b610479336107c7565b600554431015610479576040805162461bcd60e51b815260206004820152600f60248201526e4245464f52455f444541444c494e4560881b604482015290519081900360640190fd5b60086020526000908152604090205460ff1681565b6009546001600160a01b031633146105e9576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b6001600160a01b03811660009081526008602052604090205460ff16610643576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6001600160a01b03166000908152600860205260409020805460ff1916905560078054600019019055565b60035481565b60065481565b6009546001600160a01b0316156106c7576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b6001600160a01b038616610710576040805162461bcd60e51b815260206004820152600b60248201526a2927a6262aa82fa0a2222960a91b604482015290519081900360640190fd5b600980546001600160a01b0319166001600160a01b0397909716969096179095556001939093556002919091556003556004556005819055600655565b6009546001600160a01b031681565b600a5481565b6009546001600160a01b031633146107af576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b6000600a819055600b55565b60075481565b600b5481565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff16156108775760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561083c578181015183820152602001610824565b50505050905090810190601f1680156108695780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea2646970667358221220e106d6e7a49097e59c544d9a17fbe7fab611f4f565f1bc2466291e7b7f9ec60d64736f6c634300060b0033"

// DeployNodeFactory deploys a new Ethereum contract, binding an instance of NodeFactory to it.
func DeployNodeFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeFactory{NodeFactoryCaller: NodeFactoryCaller{contract: contract}, NodeFactoryTransactor: NodeFactoryTransactor{contract: contract}, NodeFactoryFilterer: NodeFactoryFilterer{contract: contract}}, nil
}

// NodeFactory is an auto generated Go binding around an Ethereum contract.
type NodeFactory struct {
	NodeFactoryCaller     // Read-only binding to the contract
	NodeFactoryTransactor // Write-only binding to the contract
	NodeFactoryFilterer   // Log filterer for contract events
}

// NodeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeFactorySession struct {
	Contract     *NodeFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeFactoryCallerSession struct {
	Contract *NodeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeFactoryTransactorSession struct {
	Contract     *NodeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeFactoryRaw struct {
	Contract *NodeFactory // Generic contract binding to access the raw methods on
}

// NodeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeFactoryCallerRaw struct {
	Contract *NodeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// NodeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeFactoryTransactorRaw struct {
	Contract *NodeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeFactory creates a new instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactory(address common.Address, backend bind.ContractBackend) (*NodeFactory, error) {
	contract, err := bindNodeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeFactory{NodeFactoryCaller: NodeFactoryCaller{contract: contract}, NodeFactoryTransactor: NodeFactoryTransactor{contract: contract}, NodeFactoryFilterer: NodeFactoryFilterer{contract: contract}}, nil
}

// NewNodeFactoryCaller creates a new read-only instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryCaller(address common.Address, caller bind.ContractCaller) (*NodeFactoryCaller, error) {
	contract, err := bindNodeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryCaller{contract: contract}, nil
}

// NewNodeFactoryTransactor creates a new write-only instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeFactoryTransactor, error) {
	contract, err := bindNodeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryTransactor{contract: contract}, nil
}

// NewNodeFactoryFilterer creates a new log filterer instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFactoryFilterer, error) {
	contract, err := bindNodeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryFilterer{contract: contract}, nil
}

// bindNodeFactory binds a generic wrapper to an already deployed contract.
func bindNodeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeFactory *NodeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeFactory.Contract.NodeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeFactory *NodeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeFactory.Contract.NodeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeFactory *NodeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeFactory.Contract.NodeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeFactory *NodeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeFactory *NodeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeFactory *NodeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeFactory.Contract.contract.Transact(opts, method, params...)
}

// TemplateContract is a free data retrieval call binding the contract method 0x72be06d8.
//
// Solidity: function templateContract() view returns(address)
func (_NodeFactory *NodeFactoryCaller) TemplateContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeFactory.contract.Call(opts, &out, "templateContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TemplateContract is a free data retrieval call binding the contract method 0x72be06d8.
//
// Solidity: function templateContract() view returns(address)
func (_NodeFactory *NodeFactorySession) TemplateContract() (common.Address, error) {
	return _NodeFactory.Contract.TemplateContract(&_NodeFactory.CallOpts)
}

// TemplateContract is a free data retrieval call binding the contract method 0x72be06d8.
//
// Solidity: function templateContract() view returns(address)
func (_NodeFactory *NodeFactoryCallerSession) TemplateContract() (common.Address, error) {
	return _NodeFactory.Contract.TemplateContract(&_NodeFactory.CallOpts)
}

// CreateNode is a paid mutator transaction binding the contract method 0xd45ab2b5.
//
// Solidity: function createNode(bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns(address)
func (_NodeFactory *NodeFactoryTransactor) CreateNode(opts *bind.TransactOpts, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _NodeFactory.contract.Transact(opts, "createNode", _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// CreateNode is a paid mutator transaction binding the contract method 0xd45ab2b5.
//
// Solidity: function createNode(bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns(address)
func (_NodeFactory *NodeFactorySession) CreateNode(_stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _NodeFactory.Contract.CreateNode(&_NodeFactory.TransactOpts, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// CreateNode is a paid mutator transaction binding the contract method 0xd45ab2b5.
//
// Solidity: function createNode(bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns(address)
func (_NodeFactory *NodeFactoryTransactorSession) CreateNode(_stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _NodeFactory.Contract.CreateNode(&_NodeFactory.TransactOpts, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}
