// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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

// GasRefunderABI is the input ABI used to generate the binding from.
const GasRefunderABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ContractAllowedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"parameter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ParameterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reason\",\"type\":\"uint256\"}],\"name\":\"RefundGasCostsDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"RefundedGasCosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calldataCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraGasMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContractRefund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasTip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxRefundeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSingleGasUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"refundee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"onGasSpent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setCalldataCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setContractAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setExtraGasMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxGasCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxGasTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxRefundeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxSingleGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// GasRefunderBin is the compiled bytecode used for deploying new contracts.
var GasRefunderBin = "0x608060405234801561001057600080fd5b5061001a33610042565b610fa0600455600c6005556377359400600655641bf08eb000600755621e8480600855610092565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610ce1806100a16000396000f3fe6080604052600436106101235760003560e01c806395554f71116100a0578063c88bee1711610064578063c88bee1714610359578063cd2c4e841461036f578063e3db8a491461038f578063f2fde38b146103af578063f3fef3a3146103cf57600080fd5b806395554f71146102c0578063a715dd59146102e0578063b029df96146102f6578063bddaf01d1461030c578063c53d26a01461033957600080fd5b8063715018a6116100e7578063715018a6146102285780637965d76b1461023d57806380599c3d1461025d5780638da5cb5b1461027d5780638e0bc1b6146102aa57600080fd5b80630119b029146101675780633a16db731461019057806351e0e26b146101b2578063549e89bb146101f2578063696e930c1461021257600080fd5b3661016257604080513381523460208201527f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4910160405180910390a1005b600080fd5b34801561017357600080fd5b5061017d60065481565b6040519081526020015b60405180910390f35b34801561019c57600080fd5b506101b06101ab366004610bc4565b6103ef565b005b3480156101be57600080fd5b506101e26101cd366004610b01565b60016020526000908152604090205460ff1681565b6040519015158152602001610187565b3480156101fe57600080fd5b506101b061020d366004610bc4565b610454565b34801561021e57600080fd5b5061017d60085481565b34801561023457600080fd5b506101b06104a9565b34801561024957600080fd5b506101b0610258366004610b86565b6104e4565b34801561026957600080fd5b506101b0610278366004610bc4565b610572565b34801561028957600080fd5b506102926105c7565b6040516001600160a01b039091168152602001610187565b3480156102b657600080fd5b5061017d60045481565b3480156102cc57600080fd5b506101b06102db366004610bc4565b6105d6565b3480156102ec57600080fd5b5061017d60075481565b34801561030257600080fd5b5061017d60055481565b34801561031857600080fd5b5061017d610327366004610b01565b60026020526000908152604090205481565b34801561034557600080fd5b506101b0610354366004610bc4565b61062b565b34801561036557600080fd5b5061017d60035481565b34801561037b57600080fd5b506101b061038a366004610bc4565b610680565b34801561039b57600080fd5b506101b06103aa366004610b51565b6106d5565b3480156103bb57600080fd5b506101b06103ca366004610b01565b61095e565b3480156103db57600080fd5b506101b06103ea366004610b25565b6109fe565b336103f86105c7565b6001600160a01b0316146104275760405162461bcd60e51b815260040161041e90610bdd565b60405180910390fd5b6003819055604051818152600090600080516020610c8c833981519152906020015b60405180910390a250565b3361045d6105c7565b6001600160a01b0316146104835760405162461bcd60e51b815260040161041e90610bdd565b6005819055604051818152600290600080516020610c8c83398151915290602001610449565b336104b26105c7565b6001600160a01b0316146104d85760405162461bcd60e51b815260040161041e90610bdd565b6104e26000610ab1565b565b336104ed6105c7565b6001600160a01b0316146105135760405162461bcd60e51b815260040161041e90610bdd565b6001600160a01b038216600081815260016020908152604091829020805460ff191685151590811790915591519182527fb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c01910160405180910390a25050565b3361057b6105c7565b6001600160a01b0316146105a15760405162461bcd60e51b815260040161041e90610bdd565b6008819055604051818152600590600080516020610c8c83398151915290602001610449565b6000546001600160a01b031690565b336105df6105c7565b6001600160a01b0316146106055760405162461bcd60e51b815260040161041e90610bdd565b6007819055604051818152600490600080516020610c8c83398151915290602001610449565b336106346105c7565b6001600160a01b03161461065a5760405162461bcd60e51b815260040161041e90610bdd565b6006819055604051818152600390600080516020610c8c83398151915290602001610449565b336106896105c7565b6001600160a01b0316146106af5760405162461bcd60e51b815260040161041e90610bdd565b6004819055604051818152600190600080516020610c8c83398151915290602001610449565b60005a3360009081526001602052604090205490915060ff166107315760405162461bcd60e51b81526020600482015260146024820152731393d517d0531313d5d15117d0d3d395149050d560621b604482015260640161041e565b3360009081526002602052604090205443141561079457604080518481526000602082015233916001600160a01b038716917f5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc910160405180910390a350505050565b3360009081526002602052604081204390556006546107b39048610c12565b9050803a10156107c057503a5b600754158015906107d2575060075481115b156107dc57506007545b6003546008546005546001600160a01b038816319291906107fd9087610c2a565b60045461080a9087610c12565b6108149190610c12565b61081e9088610c12565b96505a61082b9088610c49565b9650801580159061083d575060085487115b15610846578096505b60006108528886610c2a565b9050821580159061086b5750826108698286610c12565b115b156108d657828411156108c957604080518981526001602082015233916001600160a01b038c16917f5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc910160405180910390a3505050505050505050565b6108d38484610c49565b90505b6040516000906001600160a01b038b169083156108fc0290849084818181858888f1604080518f815260208101899052821515918101919091529095503394506001600160a01b038f1693507f7cd15c0b7405ca49bc503d8c005f66660c34917e5d7e24e50deb585da34d98539250606001905060405180910390a350505050505050505050565b336109676105c7565b6001600160a01b03161461098d5760405162461bcd60e51b815260040161041e90610bdd565b6001600160a01b0381166109f25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161041e565b6109fb81610ab1565b50565b33610a076105c7565b6001600160a01b031614610a2d5760405162461bcd60e51b815260040161041e90610bdd565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610a63573d6000803e3d6000fd5b50604080513381526001600160a01b03841660208201529081018290527fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060600160405180910390a15050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610b1357600080fd5b8135610b1e81610c76565b9392505050565b60008060408385031215610b3857600080fd5b8235610b4381610c76565b946020939093013593505050565b600080600060608486031215610b6657600080fd5b8335610b7181610c76565b95602085013595506040909401359392505050565b60008060408385031215610b9957600080fd5b8235610ba481610c76565b915060208301358015158114610bb957600080fd5b809150509250929050565b600060208284031215610bd657600080fd5b5035919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60008219821115610c2557610c25610c60565b500190565b6000816000190483118215151615610c4457610c44610c60565b500290565b600082821015610c5b57610c5b610c60565b500390565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b03811681146109fb57600080fdfe0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170ca2646970667358221220b05a80702dd14dda747f22a23b74911248ef79fc5b8f0467814293189069b10164736f6c63430008070033"

// DeployGasRefunder deploys a new Ethereum contract, binding an instance of GasRefunder to it.
func DeployGasRefunder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasRefunder, error) {
	parsed, err := abi.JSON(strings.NewReader(GasRefunderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GasRefunderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasRefunder{GasRefunderCaller: GasRefunderCaller{contract: contract}, GasRefunderTransactor: GasRefunderTransactor{contract: contract}, GasRefunderFilterer: GasRefunderFilterer{contract: contract}}, nil
}

// GasRefunder is an auto generated Go binding around an Ethereum contract.
type GasRefunder struct {
	GasRefunderCaller     // Read-only binding to the contract
	GasRefunderTransactor // Write-only binding to the contract
	GasRefunderFilterer   // Log filterer for contract events
}

// GasRefunderCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasRefunderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasRefunderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasRefunderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasRefunderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasRefunderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasRefunderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasRefunderSession struct {
	Contract     *GasRefunder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasRefunderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasRefunderCallerSession struct {
	Contract *GasRefunderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GasRefunderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasRefunderTransactorSession struct {
	Contract     *GasRefunderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GasRefunderRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasRefunderRaw struct {
	Contract *GasRefunder // Generic contract binding to access the raw methods on
}

// GasRefunderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasRefunderCallerRaw struct {
	Contract *GasRefunderCaller // Generic read-only contract binding to access the raw methods on
}

// GasRefunderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasRefunderTransactorRaw struct {
	Contract *GasRefunderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasRefunder creates a new instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunder(address common.Address, backend bind.ContractBackend) (*GasRefunder, error) {
	contract, err := bindGasRefunder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasRefunder{GasRefunderCaller: GasRefunderCaller{contract: contract}, GasRefunderTransactor: GasRefunderTransactor{contract: contract}, GasRefunderFilterer: GasRefunderFilterer{contract: contract}}, nil
}

// NewGasRefunderCaller creates a new read-only instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunderCaller(address common.Address, caller bind.ContractCaller) (*GasRefunderCaller, error) {
	contract, err := bindGasRefunder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasRefunderCaller{contract: contract}, nil
}

// NewGasRefunderTransactor creates a new write-only instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunderTransactor(address common.Address, transactor bind.ContractTransactor) (*GasRefunderTransactor, error) {
	contract, err := bindGasRefunder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasRefunderTransactor{contract: contract}, nil
}

// NewGasRefunderFilterer creates a new log filterer instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunderFilterer(address common.Address, filterer bind.ContractFilterer) (*GasRefunderFilterer, error) {
	contract, err := bindGasRefunder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasRefunderFilterer{contract: contract}, nil
}

// bindGasRefunder binds a generic wrapper to an already deployed contract.
func bindGasRefunder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasRefunderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasRefunder *GasRefunderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasRefunder.Contract.GasRefunderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasRefunder *GasRefunderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.Contract.GasRefunderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasRefunder *GasRefunderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasRefunder.Contract.GasRefunderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasRefunder *GasRefunderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasRefunder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasRefunder *GasRefunderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasRefunder *GasRefunderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasRefunder.Contract.contract.Transact(opts, method, params...)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_GasRefunder *GasRefunderCaller) AllowedContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "allowedContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_GasRefunder *GasRefunderSession) AllowedContracts(arg0 common.Address) (bool, error) {
	return _GasRefunder.Contract.AllowedContracts(&_GasRefunder.CallOpts, arg0)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_GasRefunder *GasRefunderCallerSession) AllowedContracts(arg0 common.Address) (bool, error) {
	return _GasRefunder.Contract.AllowedContracts(&_GasRefunder.CallOpts, arg0)
}

// CalldataCost is a free data retrieval call binding the contract method 0xb029df96.
//
// Solidity: function calldataCost() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) CalldataCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "calldataCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalldataCost is a free data retrieval call binding the contract method 0xb029df96.
//
// Solidity: function calldataCost() view returns(uint256)
func (_GasRefunder *GasRefunderSession) CalldataCost() (*big.Int, error) {
	return _GasRefunder.Contract.CalldataCost(&_GasRefunder.CallOpts)
}

// CalldataCost is a free data retrieval call binding the contract method 0xb029df96.
//
// Solidity: function calldataCost() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) CalldataCost() (*big.Int, error) {
	return _GasRefunder.Contract.CalldataCost(&_GasRefunder.CallOpts)
}

// ExtraGasMargin is a free data retrieval call binding the contract method 0x8e0bc1b6.
//
// Solidity: function extraGasMargin() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) ExtraGasMargin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "extraGasMargin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraGasMargin is a free data retrieval call binding the contract method 0x8e0bc1b6.
//
// Solidity: function extraGasMargin() view returns(uint256)
func (_GasRefunder *GasRefunderSession) ExtraGasMargin() (*big.Int, error) {
	return _GasRefunder.Contract.ExtraGasMargin(&_GasRefunder.CallOpts)
}

// ExtraGasMargin is a free data retrieval call binding the contract method 0x8e0bc1b6.
//
// Solidity: function extraGasMargin() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) ExtraGasMargin() (*big.Int, error) {
	return _GasRefunder.Contract.ExtraGasMargin(&_GasRefunder.CallOpts)
}

// LastContractRefund is a free data retrieval call binding the contract method 0xbddaf01d.
//
// Solidity: function lastContractRefund(address ) view returns(uint256)
func (_GasRefunder *GasRefunderCaller) LastContractRefund(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "lastContractRefund", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastContractRefund is a free data retrieval call binding the contract method 0xbddaf01d.
//
// Solidity: function lastContractRefund(address ) view returns(uint256)
func (_GasRefunder *GasRefunderSession) LastContractRefund(arg0 common.Address) (*big.Int, error) {
	return _GasRefunder.Contract.LastContractRefund(&_GasRefunder.CallOpts, arg0)
}

// LastContractRefund is a free data retrieval call binding the contract method 0xbddaf01d.
//
// Solidity: function lastContractRefund(address ) view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) LastContractRefund(arg0 common.Address) (*big.Int, error) {
	return _GasRefunder.Contract.LastContractRefund(&_GasRefunder.CallOpts, arg0)
}

// MaxGasCost is a free data retrieval call binding the contract method 0xa715dd59.
//
// Solidity: function maxGasCost() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxGasCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxGasCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasCost is a free data retrieval call binding the contract method 0xa715dd59.
//
// Solidity: function maxGasCost() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxGasCost() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasCost(&_GasRefunder.CallOpts)
}

// MaxGasCost is a free data retrieval call binding the contract method 0xa715dd59.
//
// Solidity: function maxGasCost() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxGasCost() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasCost(&_GasRefunder.CallOpts)
}

// MaxGasTip is a free data retrieval call binding the contract method 0x0119b029.
//
// Solidity: function maxGasTip() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxGasTip(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxGasTip")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasTip is a free data retrieval call binding the contract method 0x0119b029.
//
// Solidity: function maxGasTip() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxGasTip() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasTip(&_GasRefunder.CallOpts)
}

// MaxGasTip is a free data retrieval call binding the contract method 0x0119b029.
//
// Solidity: function maxGasTip() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxGasTip() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasTip(&_GasRefunder.CallOpts)
}

// MaxRefundeeBalance is a free data retrieval call binding the contract method 0xc88bee17.
//
// Solidity: function maxRefundeeBalance() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxRefundeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxRefundeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRefundeeBalance is a free data retrieval call binding the contract method 0xc88bee17.
//
// Solidity: function maxRefundeeBalance() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxRefundeeBalance() (*big.Int, error) {
	return _GasRefunder.Contract.MaxRefundeeBalance(&_GasRefunder.CallOpts)
}

// MaxRefundeeBalance is a free data retrieval call binding the contract method 0xc88bee17.
//
// Solidity: function maxRefundeeBalance() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxRefundeeBalance() (*big.Int, error) {
	return _GasRefunder.Contract.MaxRefundeeBalance(&_GasRefunder.CallOpts)
}

// MaxSingleGasUsage is a free data retrieval call binding the contract method 0x696e930c.
//
// Solidity: function maxSingleGasUsage() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxSingleGasUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxSingleGasUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSingleGasUsage is a free data retrieval call binding the contract method 0x696e930c.
//
// Solidity: function maxSingleGasUsage() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxSingleGasUsage() (*big.Int, error) {
	return _GasRefunder.Contract.MaxSingleGasUsage(&_GasRefunder.CallOpts)
}

// MaxSingleGasUsage is a free data retrieval call binding the contract method 0x696e930c.
//
// Solidity: function maxSingleGasUsage() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxSingleGasUsage() (*big.Int, error) {
	return _GasRefunder.Contract.MaxSingleGasUsage(&_GasRefunder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasRefunder *GasRefunderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasRefunder *GasRefunderSession) Owner() (common.Address, error) {
	return _GasRefunder.Contract.Owner(&_GasRefunder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasRefunder *GasRefunderCallerSession) Owner() (common.Address, error) {
	return _GasRefunder.Contract.Owner(&_GasRefunder.CallOpts)
}

// OnGasSpent is a paid mutator transaction binding the contract method 0xe3db8a49.
//
// Solidity: function onGasSpent(address refundee, uint256 gasUsed, uint256 calldataSize) returns()
func (_GasRefunder *GasRefunderTransactor) OnGasSpent(opts *bind.TransactOpts, refundee common.Address, gasUsed *big.Int, calldataSize *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "onGasSpent", refundee, gasUsed, calldataSize)
}

// OnGasSpent is a paid mutator transaction binding the contract method 0xe3db8a49.
//
// Solidity: function onGasSpent(address refundee, uint256 gasUsed, uint256 calldataSize) returns()
func (_GasRefunder *GasRefunderSession) OnGasSpent(refundee common.Address, gasUsed *big.Int, calldataSize *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.OnGasSpent(&_GasRefunder.TransactOpts, refundee, gasUsed, calldataSize)
}

// OnGasSpent is a paid mutator transaction binding the contract method 0xe3db8a49.
//
// Solidity: function onGasSpent(address refundee, uint256 gasUsed, uint256 calldataSize) returns()
func (_GasRefunder *GasRefunderTransactorSession) OnGasSpent(refundee common.Address, gasUsed *big.Int, calldataSize *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.OnGasSpent(&_GasRefunder.TransactOpts, refundee, gasUsed, calldataSize)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasRefunder *GasRefunderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasRefunder *GasRefunderSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasRefunder.Contract.RenounceOwnership(&_GasRefunder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasRefunder *GasRefunderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasRefunder.Contract.RenounceOwnership(&_GasRefunder.TransactOpts)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0x549e89bb.
//
// Solidity: function setCalldataCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetCalldataCost(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setCalldataCost", newValue)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0x549e89bb.
//
// Solidity: function setCalldataCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetCalldataCost(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetCalldataCost(&_GasRefunder.TransactOpts, newValue)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0x549e89bb.
//
// Solidity: function setCalldataCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetCalldataCost(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetCalldataCost(&_GasRefunder.TransactOpts, newValue)
}

// SetContractAllowed is a paid mutator transaction binding the contract method 0x7965d76b.
//
// Solidity: function setContractAllowed(address contractAddress, bool allowed) returns()
func (_GasRefunder *GasRefunderTransactor) SetContractAllowed(opts *bind.TransactOpts, contractAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setContractAllowed", contractAddress, allowed)
}

// SetContractAllowed is a paid mutator transaction binding the contract method 0x7965d76b.
//
// Solidity: function setContractAllowed(address contractAddress, bool allowed) returns()
func (_GasRefunder *GasRefunderSession) SetContractAllowed(contractAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetContractAllowed(&_GasRefunder.TransactOpts, contractAddress, allowed)
}

// SetContractAllowed is a paid mutator transaction binding the contract method 0x7965d76b.
//
// Solidity: function setContractAllowed(address contractAddress, bool allowed) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetContractAllowed(contractAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetContractAllowed(&_GasRefunder.TransactOpts, contractAddress, allowed)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0xcd2c4e84.
//
// Solidity: function setExtraGasMargin(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetExtraGasMargin(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setExtraGasMargin", newValue)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0xcd2c4e84.
//
// Solidity: function setExtraGasMargin(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetExtraGasMargin(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetExtraGasMargin(&_GasRefunder.TransactOpts, newValue)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0xcd2c4e84.
//
// Solidity: function setExtraGasMargin(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetExtraGasMargin(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetExtraGasMargin(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x95554f71.
//
// Solidity: function setMaxGasCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxGasCost(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxGasCost", newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x95554f71.
//
// Solidity: function setMaxGasCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxGasCost(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasCost(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x95554f71.
//
// Solidity: function setMaxGasCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxGasCost(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasCost(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0xc53d26a0.
//
// Solidity: function setMaxGasTip(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxGasTip(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxGasTip", newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0xc53d26a0.
//
// Solidity: function setMaxGasTip(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxGasTip(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasTip(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0xc53d26a0.
//
// Solidity: function setMaxGasTip(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxGasTip(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasTip(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0x3a16db73.
//
// Solidity: function setMaxRefundeeBalance(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxRefundeeBalance(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxRefundeeBalance", newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0x3a16db73.
//
// Solidity: function setMaxRefundeeBalance(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxRefundeeBalance(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxRefundeeBalance(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0x3a16db73.
//
// Solidity: function setMaxRefundeeBalance(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxRefundeeBalance(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxRefundeeBalance(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x80599c3d.
//
// Solidity: function setMaxSingleGasUsage(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxSingleGasUsage(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxSingleGasUsage", newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x80599c3d.
//
// Solidity: function setMaxSingleGasUsage(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxSingleGasUsage(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxSingleGasUsage(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x80599c3d.
//
// Solidity: function setMaxSingleGasUsage(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxSingleGasUsage(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxSingleGasUsage(&_GasRefunder.TransactOpts, newValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasRefunder *GasRefunderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasRefunder *GasRefunderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.TransferOwnership(&_GasRefunder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasRefunder *GasRefunderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.TransferOwnership(&_GasRefunder.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address destination, uint256 amount) returns()
func (_GasRefunder *GasRefunderTransactor) Withdraw(opts *bind.TransactOpts, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "withdraw", destination, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address destination, uint256 amount) returns()
func (_GasRefunder *GasRefunderSession) Withdraw(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.Withdraw(&_GasRefunder.TransactOpts, destination, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address destination, uint256 amount) returns()
func (_GasRefunder *GasRefunderTransactorSession) Withdraw(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.Withdraw(&_GasRefunder.TransactOpts, destination, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasRefunder *GasRefunderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasRefunder *GasRefunderSession) Receive() (*types.Transaction, error) {
	return _GasRefunder.Contract.Receive(&_GasRefunder.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasRefunder *GasRefunderTransactorSession) Receive() (*types.Transaction, error) {
	return _GasRefunder.Contract.Receive(&_GasRefunder.TransactOpts)
}

// GasRefunderContractAllowedSetIterator is returned from FilterContractAllowedSet and is used to iterate over the raw logs and unpacked data for ContractAllowedSet events raised by the GasRefunder contract.
type GasRefunderContractAllowedSetIterator struct {
	Event *GasRefunderContractAllowedSet // Event containing the contract specifics and raw log

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
func (it *GasRefunderContractAllowedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderContractAllowedSet)
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
		it.Event = new(GasRefunderContractAllowedSet)
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
func (it *GasRefunderContractAllowedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderContractAllowedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderContractAllowedSet represents a ContractAllowedSet event raised by the GasRefunder contract.
type GasRefunderContractAllowedSet struct {
	ContractAddress common.Address
	Allowed         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractAllowedSet is a free log retrieval operation binding the contract event 0xb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c01.
//
// Solidity: event ContractAllowedSet(address indexed contractAddress, bool allowed)
func (_GasRefunder *GasRefunderFilterer) FilterContractAllowedSet(opts *bind.FilterOpts, contractAddress []common.Address) (*GasRefunderContractAllowedSetIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "ContractAllowedSet", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderContractAllowedSetIterator{contract: _GasRefunder.contract, event: "ContractAllowedSet", logs: logs, sub: sub}, nil
}

// WatchContractAllowedSet is a free log subscription operation binding the contract event 0xb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c01.
//
// Solidity: event ContractAllowedSet(address indexed contractAddress, bool allowed)
func (_GasRefunder *GasRefunderFilterer) WatchContractAllowedSet(opts *bind.WatchOpts, sink chan<- *GasRefunderContractAllowedSet, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "ContractAllowedSet", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderContractAllowedSet)
				if err := _GasRefunder.contract.UnpackLog(event, "ContractAllowedSet", log); err != nil {
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

// ParseContractAllowedSet is a log parse operation binding the contract event 0xb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c01.
//
// Solidity: event ContractAllowedSet(address indexed contractAddress, bool allowed)
func (_GasRefunder *GasRefunderFilterer) ParseContractAllowedSet(log types.Log) (*GasRefunderContractAllowedSet, error) {
	event := new(GasRefunderContractAllowedSet)
	if err := _GasRefunder.contract.UnpackLog(event, "ContractAllowedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GasRefunder contract.
type GasRefunderDepositedIterator struct {
	Event *GasRefunderDeposited // Event containing the contract specifics and raw log

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
func (it *GasRefunderDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderDeposited)
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
		it.Event = new(GasRefunderDeposited)
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
func (it *GasRefunderDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderDeposited represents a Deposited event raised by the GasRefunder contract.
type GasRefunderDeposited struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) FilterDeposited(opts *bind.FilterOpts) (*GasRefunderDepositedIterator, error) {

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &GasRefunderDepositedIterator{contract: _GasRefunder.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GasRefunderDeposited) (event.Subscription, error) {

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderDeposited)
				if err := _GasRefunder.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) ParseDeposited(log types.Log) (*GasRefunderDeposited, error) {
	event := new(GasRefunderDeposited)
	if err := _GasRefunder.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasRefunder contract.
type GasRefunderOwnershipTransferredIterator struct {
	Event *GasRefunderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GasRefunderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderOwnershipTransferred)
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
		it.Event = new(GasRefunderOwnershipTransferred)
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
func (it *GasRefunderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderOwnershipTransferred represents a OwnershipTransferred event raised by the GasRefunder contract.
type GasRefunderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasRefunder *GasRefunderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasRefunderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderOwnershipTransferredIterator{contract: _GasRefunder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasRefunder *GasRefunderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasRefunderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderOwnershipTransferred)
				if err := _GasRefunder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GasRefunder *GasRefunderFilterer) ParseOwnershipTransferred(log types.Log) (*GasRefunderOwnershipTransferred, error) {
	event := new(GasRefunderOwnershipTransferred)
	if err := _GasRefunder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderParameterSetIterator is returned from FilterParameterSet and is used to iterate over the raw logs and unpacked data for ParameterSet events raised by the GasRefunder contract.
type GasRefunderParameterSetIterator struct {
	Event *GasRefunderParameterSet // Event containing the contract specifics and raw log

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
func (it *GasRefunderParameterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderParameterSet)
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
		it.Event = new(GasRefunderParameterSet)
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
func (it *GasRefunderParameterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderParameterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderParameterSet represents a ParameterSet event raised by the GasRefunder contract.
type GasRefunderParameterSet struct {
	Parameter *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterParameterSet is a free log retrieval operation binding the contract event 0x0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170c.
//
// Solidity: event ParameterSet(uint256 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) FilterParameterSet(opts *bind.FilterOpts, parameter []*big.Int) (*GasRefunderParameterSetIterator, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "ParameterSet", parameterRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderParameterSetIterator{contract: _GasRefunder.contract, event: "ParameterSet", logs: logs, sub: sub}, nil
}

// WatchParameterSet is a free log subscription operation binding the contract event 0x0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170c.
//
// Solidity: event ParameterSet(uint256 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) WatchParameterSet(opts *bind.WatchOpts, sink chan<- *GasRefunderParameterSet, parameter []*big.Int) (event.Subscription, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "ParameterSet", parameterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderParameterSet)
				if err := _GasRefunder.contract.UnpackLog(event, "ParameterSet", log); err != nil {
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

// ParseParameterSet is a log parse operation binding the contract event 0x0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170c.
//
// Solidity: event ParameterSet(uint256 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) ParseParameterSet(log types.Log) (*GasRefunderParameterSet, error) {
	event := new(GasRefunderParameterSet)
	if err := _GasRefunder.contract.UnpackLog(event, "ParameterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderRefundGasCostsDeniedIterator is returned from FilterRefundGasCostsDenied and is used to iterate over the raw logs and unpacked data for RefundGasCostsDenied events raised by the GasRefunder contract.
type GasRefunderRefundGasCostsDeniedIterator struct {
	Event *GasRefunderRefundGasCostsDenied // Event containing the contract specifics and raw log

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
func (it *GasRefunderRefundGasCostsDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderRefundGasCostsDenied)
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
		it.Event = new(GasRefunderRefundGasCostsDenied)
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
func (it *GasRefunderRefundGasCostsDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderRefundGasCostsDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderRefundGasCostsDenied represents a RefundGasCostsDenied event raised by the GasRefunder contract.
type GasRefunderRefundGasCostsDenied struct {
	Refundee        common.Address
	ContractAddress common.Address
	Gas             *big.Int
	Reason          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundGasCostsDenied is a free log retrieval operation binding the contract event 0x5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 reason)
func (_GasRefunder *GasRefunderFilterer) FilterRefundGasCostsDenied(opts *bind.FilterOpts, refundee []common.Address, contractAddress []common.Address) (*GasRefunderRefundGasCostsDeniedIterator, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "RefundGasCostsDenied", refundeeRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderRefundGasCostsDeniedIterator{contract: _GasRefunder.contract, event: "RefundGasCostsDenied", logs: logs, sub: sub}, nil
}

// WatchRefundGasCostsDenied is a free log subscription operation binding the contract event 0x5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 reason)
func (_GasRefunder *GasRefunderFilterer) WatchRefundGasCostsDenied(opts *bind.WatchOpts, sink chan<- *GasRefunderRefundGasCostsDenied, refundee []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "RefundGasCostsDenied", refundeeRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderRefundGasCostsDenied)
				if err := _GasRefunder.contract.UnpackLog(event, "RefundGasCostsDenied", log); err != nil {
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

// ParseRefundGasCostsDenied is a log parse operation binding the contract event 0x5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 reason)
func (_GasRefunder *GasRefunderFilterer) ParseRefundGasCostsDenied(log types.Log) (*GasRefunderRefundGasCostsDenied, error) {
	event := new(GasRefunderRefundGasCostsDenied)
	if err := _GasRefunder.contract.UnpackLog(event, "RefundGasCostsDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderRefundedGasCostsIterator is returned from FilterRefundedGasCosts and is used to iterate over the raw logs and unpacked data for RefundedGasCosts events raised by the GasRefunder contract.
type GasRefunderRefundedGasCostsIterator struct {
	Event *GasRefunderRefundedGasCosts // Event containing the contract specifics and raw log

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
func (it *GasRefunderRefundedGasCostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderRefundedGasCosts)
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
		it.Event = new(GasRefunderRefundedGasCosts)
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
func (it *GasRefunderRefundedGasCostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderRefundedGasCostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderRefundedGasCosts represents a RefundedGasCosts event raised by the GasRefunder contract.
type GasRefunderRefundedGasCosts struct {
	Refundee        common.Address
	ContractAddress common.Address
	Gas             *big.Int
	AmountPaid      *big.Int
	Success         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundedGasCosts is a free log retrieval operation binding the contract event 0x7cd15c0b7405ca49bc503d8c005f66660c34917e5d7e24e50deb585da34d9853.
//
// Solidity: event RefundedGasCosts(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 amountPaid, bool success)
func (_GasRefunder *GasRefunderFilterer) FilterRefundedGasCosts(opts *bind.FilterOpts, refundee []common.Address, contractAddress []common.Address) (*GasRefunderRefundedGasCostsIterator, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "RefundedGasCosts", refundeeRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderRefundedGasCostsIterator{contract: _GasRefunder.contract, event: "RefundedGasCosts", logs: logs, sub: sub}, nil
}

// WatchRefundedGasCosts is a free log subscription operation binding the contract event 0x7cd15c0b7405ca49bc503d8c005f66660c34917e5d7e24e50deb585da34d9853.
//
// Solidity: event RefundedGasCosts(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 amountPaid, bool success)
func (_GasRefunder *GasRefunderFilterer) WatchRefundedGasCosts(opts *bind.WatchOpts, sink chan<- *GasRefunderRefundedGasCosts, refundee []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "RefundedGasCosts", refundeeRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderRefundedGasCosts)
				if err := _GasRefunder.contract.UnpackLog(event, "RefundedGasCosts", log); err != nil {
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

// ParseRefundedGasCosts is a log parse operation binding the contract event 0x7cd15c0b7405ca49bc503d8c005f66660c34917e5d7e24e50deb585da34d9853.
//
// Solidity: event RefundedGasCosts(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 amountPaid, bool success)
func (_GasRefunder *GasRefunderFilterer) ParseRefundedGasCosts(log types.Log) (*GasRefunderRefundedGasCosts, error) {
	event := new(GasRefunderRefundedGasCosts)
	if err := _GasRefunder.contract.UnpackLog(event, "RefundedGasCosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GasRefunder contract.
type GasRefunderWithdrawnIterator struct {
	Event *GasRefunderWithdrawn // Event containing the contract specifics and raw log

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
func (it *GasRefunderWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderWithdrawn)
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
		it.Event = new(GasRefunderWithdrawn)
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
func (it *GasRefunderWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderWithdrawn represents a Withdrawn event raised by the GasRefunder contract.
type GasRefunderWithdrawn struct {
	Initiator   common.Address
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address initiator, address destination, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*GasRefunderWithdrawnIterator, error) {

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &GasRefunderWithdrawnIterator{contract: _GasRefunder.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address initiator, address destination, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GasRefunderWithdrawn) (event.Subscription, error) {

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderWithdrawn)
				if err := _GasRefunder.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address initiator, address destination, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) ParseWithdrawn(log types.Log) (*GasRefunderWithdrawn, error) {
	event := new(GasRefunderWithdrawn)
	if err := _GasRefunder.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
