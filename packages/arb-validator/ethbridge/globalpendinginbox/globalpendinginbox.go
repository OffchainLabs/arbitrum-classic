// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalpendinginbox

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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204a4406891b86a3dacfb6ea735b92a605972908aec5c506cc187bb6f01ef6a24064736f6c634300050f0032"

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

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// GlobalPendingInboxABI is the input ABI used to generate the binding from.
const GlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"AssertionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ERC20DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ERC721DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EthDepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionMessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"depositERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getNFTTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getPending\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getTokenBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var GlobalPendingInboxFuncSigs = map[string]string{
	"bca22b76": "depositERC20Message(address,address,address,uint256)",
	"8b7010aa": "depositERC721Message(address,address,address,uint256)",
	"5bd21290": "depositEthMessage(address,address)",
	"8bef8df0": "forwardTransactionMessage(address,address,uint256,uint256,bytes,bytes)",
	"578c049a": "getNFTTokens(address)",
	"11ae9ed2": "getPending()",
	"c489744b": "getTokenBalance(address,address)",
	"764f3aa8": "getTokenBalances(address)",
	"e318b003": "hasNFT(address,address,uint256)",
	"f3972383": "registerForInbox()",
	"e4eb8c63": "sendMessages(bytes)",
	"8f5ed73e": "sendTransactionMessage(address,address,uint256,uint256,bytes)",
	"f4f3b200": "withdrawERC20(address)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"a0ef91df": "withdrawEth()",
}

// GlobalPendingInboxBin is the compiled bytecode used for deploying new contracts.
var GlobalPendingInboxBin = "0x6080604052600280546001600160a01b03191673caad408788c192979384768dd5be04ec1b3787da17905534801561003657600080fd5b5061395b806100466000396000f3fe6080604052600436106100e85760003560e01c8063a0ef91df1161008a578063e4eb8c6311610059578063e4eb8c6314610525578063f3972383146105a0578063f3e414f8146105b5578063f4f3b200146105ee576100e8565b8063a0ef91df14610423578063bca22b7614610438578063c489744b14610481578063e318b003146104ce576100e8565b8063764f3aa8116100c6578063764f3aa8146102175780638b7010aa1461024a5780638bef8df0146102935780638f5ed73e14610383576100e8565b806311ae9ed2146100ed578063578c049a1461011b5780635bd21290146101e7575b600080fd5b3480156100f957600080fd5b50610102610621565b6040805192835260208301919091528051918290030190f35b34801561012757600080fd5b5061014e6004803603602081101561013e57600080fd5b50356001600160a01b031661063c565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561019257818101518382015260200161017a565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156101d15781810151838201526020016101b9565b5050505090500194505050505060405180910390f35b610215600480360360408110156101fd57600080fd5b506001600160a01b03813581169160200135166107d4565b005b34801561022357600080fd5b5061014e6004803603602081101561023a57600080fd5b50356001600160a01b03166107ed565b34801561025657600080fd5b506102156004803603608081101561026d57600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610921565b34801561029f57600080fd5b50610215600480360360c08110156102b657600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359181019060a081016080820135600160201b8111156102f557600080fd5b82018360208201111561030757600080fd5b803590602001918460018302840111600160201b8311171561032857600080fd5b919390929091602081019035600160201b81111561034557600080fd5b82018360208201111561035757600080fd5b803590602001918460018302840111600160201b8311171561037857600080fd5b50909250905061093f565b34801561038f57600080fd5b50610215600480360360a08110156103a657600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359181019060a081016080820135600160201b8111156103e557600080fd5b8201836020820111156103f757600080fd5b803590602001918460018302840111600160201b8311171561041857600080fd5b509092509050610a80565b34801561042f57600080fd5b50610215610acc565b34801561044457600080fd5b506102156004803603608081101561045b57600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610b4f565b34801561048d57600080fd5b506104bc600480360360408110156104a457600080fd5b506001600160a01b0381358116916020013516610b67565b60408051918252519081900360200190f35b3480156104da57600080fd5b50610511600480360360608110156104f157600080fd5b506001600160a01b03813581169160208101359091169060400135610bcf565b604080519115158252519081900360200190f35b34801561053157600080fd5b506102156004803603602081101561054857600080fd5b810190602081018135600160201b81111561056257600080fd5b82018360208201111561057457600080fd5b803590602001918460018302840111600160201b8311171561059557600080fd5b509092509050610c50565b3480156105ac57600080fd5b50610215610d77565b3480156105c157600080fd5b50610215600480360360408110156105d857600080fd5b506001600160a01b038135169060200135610df3565b3480156105fa57600080fd5b506102156004803603602081101561061157600080fd5b50356001600160a01b0316610eb7565b33600090815260016020819052604090912080549101549091565b6001600160a01b0381166000908152602081905260408120606091829190805b60038301548110156106985782600301818154811061067757fe5b6000918252602090912060026003909202010154919091019060010161065c565b6060826040519080825280602002602001820160405280156106c4578160200160208202803883390190505b5090506060836040519080825280602002602001820160405280156106f3578160200160208202803883390190505b5060038601546000945090915083905b808510156107c557600087600301868154811061071c57fe5b600091825260208220600260039092020190810154909250905b818110156107b757825487516001600160a01b039091169088908790811061075a57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505082600201818154811061078957fe5b90600052602060002001548686815181106107a057fe5b602090810291909101015260019485019401610736565b505060019095019450610703565b50919650945050505050915091565b6107dd82610f84565b6107e982338334610f93565b5050565b6001600160a01b03811660009081526020818152604091829020600181015483518181528184028101909301909352606092839283918015610839578160200160208202803883390190505b50905060608151604051908082528060200260200182016040528015610869578160200160208202803883390190505b50825190915060005b818110156109145784600101818154811061088957fe5b600091825260209091206002909102015484516001600160a01b03909116908590839081106108b457fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508460010181815481106108e357fe5b90600052602060002090600202016001015483828151811061090157fe5b6020908102919091010152600101610872565b5091945092505050915091565b61092c828583611241565b61093984338585856112bd565b50505050565b6000610a2f898989896109878a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061137092505050565b60405160200180866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b8152601401848152602001838152602001828152602001955050505050506040516020818303038152906040528051906020012084848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113f692505050565b9050610a758989838a8a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061152992505050565b505050505050505050565b610ac4868633878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061152992505050565b505050505050565b6000610ad9600033610b67565b9050610ae7336000836118a6565b610b225760405162461bcd60e51b815260040180806020018281038252602e8152602001806138d9602e913960400191505060405180910390fd5b604051339082156108fc029083906000818181858888f193505050501580156107e9573d6000803e3d6000fd5b610b5a828583611a46565b6109398433858585611ad3565b6001600160a01b0380821660009081526020818152604080832093861683529083905281205490919080610ba057600092505050610bc9565b816001016001820381548110610bb257fe5b906000526020600020906002020160010154925050505b92915050565b6001600160a01b0380831660009081526020818152604080832093871683526002840190915281205490919080610c0b57600092505050610c49565b816003016001820381548110610c1d57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b60025460408051600080825260208201819052338284015260806060838101829052600691840191909152656576656e743160d01b60a084015292519093849384938493849389926001600160a01b03909116916000805160206139078339815191529181900360c00190a2604080516000808252602082015233818301819052608060608301819052600690830152656576656e743160d01b60a083015291516000805160206139078339815191529181900360c00190a25b80861015610a7557610d5389898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250611b85915050565b949b509299509097509550935091508615610d7257610d728285611cd2565b610d0a565b3360009081526001602052604090205415610dd9576040805162461bcd60e51b815260206004820152601d60248201527f50656e64696e67206d75737420626520756e696e697469616c697a6564000000604482015290519081900360640190fd5b610de1612218565b33600090815260016020526040902055565b610dfe33838361228c565b610e4f576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b158015610ea357600080fd5b505af1158015610ac4573d6000803e3d6000fd5b6000610ec38233610b67565b9050610ed03383836118a6565b610f0b5760405162461bcd60e51b815260040180806020018281038252602e8152602001806138d9602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b158015610f5a57600080fd5b505af1158015610f6e573d6000803e3d6000fd5b505050506040513d602081101561093957600080fd5b610f90816000346124fb565b50565b6001600160a01b038416600090815260016020526040902054156109395760408051600160f81b602080830191909152606087811b6001600160601b0319908116602185015286821b8116603585015287821b166049840152605d80840186905284518085039091018152607d840180865281519190930120600380845260fd8501909552939092609d015b6110276137ae565b81526020019060019003908161101f57905050905061104e846001600160a01b03166125e4565b8160008151811061105b57fe5b602002602001018190525061106f836125e4565b8160018151811061107c57fe5b602090810291909101015260408051600380825260808201909252606091816020015b6110a76137ae565b81526020019060019003908161109f5790505090506110c660016125e4565b816000815181106110d357fe5b60200260200101819052506110f0866001600160a01b03166125e4565b816001815181106110fd57fe5b602002602001018190525061111182612662565b8160028151811061111e57fe5b602090810291909101015260408051600380825260808201909252606091816020015b6111496137ae565b815260200190600190039081611141579050509050611167436125e4565b8160008151811061117457fe5b6020908102919091010152611188846125e4565b8160018151811061119557fe5b60200260200101819052506111a982612662565b816002815181106111b657fe5b602002602001018190525060006111d46111cf83612662565b612712565b5190506111e18982612848565b336001600160a01b0316876001600160a01b03168a6001600160a01b03167f4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40346040518082815260200191505060405180910390a4505050505050505050565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd91606480830192600092919082900301818387803b15801561129557600080fd5b505af11580156112a9573d6000803e3d6000fd5b505050506112b882848361287e565b505050565b6001600160a01b038516600090815260016020526040902054156113695760006112ec60038787878787612a16565b90506112f88682612848565b846001600160a01b0316846001600160a01b0316876001600160a01b03167fb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39868660405180836001600160a01b03166001600160a01b031681526020018281526020019250505060405180910390a4505b5050505050565b6000808061137c6137ae565b611387856000612c88565b9194509250905082156113e1576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b6113ea81612712565b5193505050505b919050565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081886040516020018083805190602001908083835b6020831061146c5780518252601f19909201916020918201910161144d565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381529382019052825192019190912092506114b391508890506000612e12565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa158015611512573d6000803e3d6000fd5b5050604051601f1901519998505050505050505050565b6001600160a01b038616600090815260016020526040902080541561189d57600061155383611370565b60408051600060208083019190915260608c811b6001600160601b031990811660218501528c821b811660358501528b821b166049840152605d83018a9052607d8301899052609d8084018690528451808503909101815260bd840180865281519190930120600480845261015d850190955294955092909160dd015b6115d86137ae565b8152602001906001900390816115d05790505090506115ff896001600160a01b03166125e4565b8160008151811061160c57fe5b6020026020010181905250611620876125e4565b8160018151811061162d57fe5b6020026020010181905250611641866125e4565b8160028151811061164e57fe5b602002602001018190525061166283612ea0565b8160038151811061166f57fe5b602090810291909101015260408051600380825260808201909252606091816020015b61169a6137ae565b8152602001906001900390816116925790505090506116b960006125e4565b816000815181106116c657fe5b60200260200101819052506116e3896001600160a01b03166125e4565b816001815181106116f057fe5b602002602001018190525061170482612662565b8160028151811061171157fe5b602090810291909101015260408051600380825260808201909252606091816020015b61173c6137ae565b81526020019060019003908161173457905050905061175a436125e4565b8160008151811061176757fe5b602090810291909101015261177b846125e4565b8160018151811061178857fe5b602002602001018190525061179c82612662565b816002815181106117a957fe5b602002602001018190525060006117c26111cf83612662565b5190506117cf8d82612848565b8a6001600160a01b03168c6001600160a01b03168e6001600160a01b03167fcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b38d8d8d6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561185a578181015183820152602001611842565b50505050905090810190601f1680156118875780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a45050505050505b50505050505050565b6000816118b557506001610c49565b6001600160a01b03808516600090815260208181526040808320938716835290839052902054806118eb57600092505050610c49565b60008260010160018303815481106118ff57fe5b9060005260206000209060020201905080600101548511156119275760009350505050610c49565b600181015461193c908663ffffffff612f1e16565b60018201819055611a39576001830180548391859160009190600019810190811061196357fe5b600091825260208083206002909202909101546001600160a01b0316835282019290925260400190205560018301805460001981019081106119a157fe5b90600052602060002090600202018360010160018403815481106119c157fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b0393841617815560019485015490850155908916825285905260408120558301805480611a0f57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd9160648083019260209291908290030181600087803b158015611a9b57600080fd5b505af1158015611aaf573d6000803e3d6000fd5b505050506040513d6020811015611ac557600080fd5b506112b890508284836124fb565b6001600160a01b03851660009081526001602052604090205415611369576000611b0260028787878787612a16565b9050611b0e8682612848565b846001600160a01b0316846001600160a01b0316876001600160a01b03167fb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355868660405180836001600160a01b03166001600160a01b031681526020018281526020019250505060405180910390a4505050505050565b600080600080600060608694506000888681518110611ba057fe5b016020015160019096019560f81c905060068114611bbe5750611cc8565b611bc88987612f60565b8097508195505050888681518110611bdc57fe5b016020015160019096019560f81c9050611bf68987612f60565b8097508194505050888681518110611c0a57fe5b0160200151600187019660f89190911c9150600219820160606000611c30838e8c612f89565b909b5092509050611c4a8d85808d0363ffffffff61304416565b6040805160028082526060808301845260019f509399509091602083019080388339019050509050611c7b896130c4565b81600081518110611c8857fe5b602002602001018181525050611c9d886130c4565b81600181518110611caa57fe5b602002602001018181525050611cbf816130e8565b99505050505050505b9295509295509295565b6001811415611ea857600080600073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63ac715d70866040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611d47578181015183820152602001611d2f565b50505050905090810190601f168015611d745780820380516001836020036101000a031916815260200191505b509250505060606040518083038186803b158015611d9157600080fd5b505af4158015611da5573d6000803e3d6000fd5b505050506040513d6060811015611dbb57600080fd5b508051602080830151604093840151600254855185151581529384018a9052606083811c8588015260809085018190526006908501526532bb32b73a1960d11b60a0850152945193975090955093506001600160a01b0390921691600080516020613907833981519152919081900360c00190a260408051841515815260208101869052606084811c8284015260809082018190526006908201526532bb32b73a1960d11b60a082015290513391600080516020613907833981519152919081900360c00190a28215611ea057611e9733606084901c836131a8565b611ea057600080fd5b5050506107e9565b600281141561208e5760008060008073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63bf532221876040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611f1e578181015183820152602001611f06565b50505050905090810190601f168015611f4b5780820380516001836020036101000a031916815260200191505b509250505060806040518083038186803b158015611f6857600080fd5b505af4158015611f7c573d6000803e3d6000fd5b505050506040513d6080811015611f9257600080fd5b508051602080830151604080850151606095860151600254835187151581529586018d905282881c868501526080978601889052600697860197909752656576656e743360d01b60a0860152915194995091975090955093506001600160a01b0390921691600080516020613907833981519152919081900360c00190a260408051851515815260208101879052606084811c828401526080908201819052600690820152656576656e743360d01b60a082015290513391600080516020613907833981519152919081900360c00190a283156120855761207c33606084811c9086901c846131d8565b61208557600080fd5b505050506107e9565b60038114156107e95760008060008073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63bf532221876040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156121045781810151838201526020016120ec565b50505050905090810190601f1680156121315780820380516001836020036101000a031916815260200191505b509250505060806040518083038186803b15801561214e57600080fd5b505af4158015612162573d6000803e3d6000fd5b505050506040513d608081101561217857600080fd5b508051602080830151604080850151606095860151600254835187151581529586018d905282881c86850152608097860188905260069786019790975265195d995b9d0d60d21b60a0860152915194995091975090955093506001600160a01b0390921691600080516020613907833981519152919081900360c00190a28315610ac45761220f33606084811c9086901c84613208565b610ac457600080fd5b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b8381101561226457818101518382015260200161224c565b5050505090500192505050604051602081830303815290604052805190602001209150505b90565b6001600160a01b03808416600090815260208181526040808320938616835260028401909152812054909190806122c857600092505050610c49565b60008260030160018303815481106122dc57fe5b600091825260208083208884526001600390930201918201905260409091205490915080612311576000945050505050610c49565b6002820180548291600185019160009190600019810190811061233057fe5b60009182526020808320909101548352820192909252604001902055600282018054600019810190811061236057fe5b906000526020600020015482600201600183038154811061237d57fe5b6000918252602080832090910192909255878152600184019091526040812055600282018054806123aa57fe5b60008281526020812082016000199081019190915501905560028201546124ed57600384018054849160028701916000919060001981019081106123ea57fe5b60009182526020808320600392830201546001600160a01b03168452830193909352604090910190209190915584018054600019810190811061242957fe5b906000526020600020906003020184600301600185038154811061244957fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b039092169190911781556002808301805461248c92840191906137dc565b5050506001600160a01b0387166000908152600285016020526040812055600384018054806124b757fe5b60008281526020812060036000199093019283020180546001600160a01b0319168155906124e8600283018261382c565b505090555b506001979650505050505050565b80612505576112b8565b6001600160a01b038084166000908152602081815260408083209386168352908390529020548061259c57506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b60008260010160018303815481106125b057fe5b906000526020600020906002020190506125d784826001015461322c90919063ffffffff16565b6001909101555050505050565b6125ec6137ae565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612651565b61263e6137ae565b8152602001906001900390816126365790505b508152600060209091015292915050565b61266a6137ae565b6126748251613286565b6126c5576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b61271a61384a565b6060820151600c60ff9091161061276c576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661279957604051806020016040528061279084600001516130c4565b905290506113f1565b606082015160ff16600114156127e057604051806020016040528061279084602001516000015185602001516040015186602001516060015187602001516020015161328d565b606082015160ff166002141561280557506040805160208101909152815181526113f1565b600360ff16826060015160ff161015801561282957506060820151600c60ff909116105b156128465760405180602001604052806127908460400151613335565bfe5b6001600160a01b0382166000908152600160205260409020805461286c9083613481565b81556001908101805490910190555050565b6001600160a01b0380841660009081526020818152604080832093861683526002840190915290205480612952576040805180820182526001600160a01b03861681528151600080825260208281019094526003860193830191905090528154600181018084556000938452602093849020835160039093020180546001600160a01b0319166001600160a01b03909316929092178255828401518051919461292f9260028501929091019061385c565b5050506001600160a01b0385166000908152600284016020526040902081905590505b600082600301600183038154811061296657fe5b90600052602060002090600302019050806001016000858152602001908152602001600020546000146129e0576040805162461bcd60e51b815260206004820152601d60248201527f63616e27742061646420616c7265616479206f776e656420746f6b656e000000604482015290519081900360640190fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b6040805160f888901b6001600160f81b031916602080830191909152606088811b6001600160601b0319908116602185015288821b8116603585015287821b8116604985015286821b16605d840152607180840186905284518085039091018152609184018086528151919093012060038084526101118501909552600094909391929160b1015b612aa66137ae565b815260200190600190039081612a9e579050509050612acd856001600160a01b03166125e4565b81600081518110612ada57fe5b6020026020010181905250612af7866001600160a01b03166125e4565b81600181518110612b0457fe5b6020026020010181905250612b18846125e4565b81600281518110612b2557fe5b602090810291909101015260408051600380825260808201909252606091816020015b612b506137ae565b815260200190600190039081612b48579050509050612b718a60ff166125e4565b81600081518110612b7e57fe5b6020026020010181905250612b9b886001600160a01b03166125e4565b81600181518110612ba857fe5b6020026020010181905250612bbc82612662565b81600281518110612bc957fe5b602090810291909101015260408051600380825260808201909252606091816020015b612bf46137ae565b815260200190600190039081612bec579050509050612c12436125e4565b81600081518110612c1f57fe5b6020908102919091010152612c33846125e4565b81600181518110612c4057fe5b6020026020010181905250612c5482612662565b81600281518110612c6157fe5b6020026020010181905250612c786111cf82612662565b519b9a5050505050505050505050565b600080612c936137ae565b84518410612ce8576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110612cfb57fe5b016020015160019092019160f81c90506000612d15613897565b60ff8316612d4957612d278985612f60565b9094509150600084612d38846125e4565b91985096509450612e0b9350505050565b60ff831660011415612d7057612d5f89856134af565b9094509050600084612d388361361c565b60ff831660021415612d9757612d868985612f60565b9094509150600084612d3884612ea0565b600360ff841610801590612dae5750600c60ff8416105b15612deb57600219830160606000612dc7838d89612f89565b909850925090508087612dd984612662565b99509950995050505050505050612e0b565b8260ff16612710016000612dff60006125e4565b91985096509450505050505b9250925092565b604180820283810160208101516040820151919093015160ff169291601b841015612e3e57601b840193505b8360ff16601b1480612e5357508360ff16601c145b612e98576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b612ea86137ae565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612f0d565b612efa6137ae565b815260200190600190039081612ef25790505b508152600260209091015292915050565b6000610c4983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061367c565b6000808281612f75868363ffffffff61371316565b6020929092019350909150505b9250929050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015612fd457816020015b612fc16137ae565b815260200190600190039081612fb95790505b50905060005b8960ff168160ff16101561302e57612ff28985612c88565b8451859060ff861690811061300357fe5b60209081029190910101529450925082156130265750909450909250905061303b565b600101612fda565b5060009550919350909150505b93509350939050565b60608183018451101561305657600080fd5b606082158015613071576040519150602082016040526130bb565b6040519150601f8416801560200281840101858101878315602002848b0101015b818310156130aa578051835260209283019201613092565b5050858452601f01601f1916604052505b50949350505050565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115613138576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8151600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613180578181015183820152602001613168565b5050505090500192505050604051602081830303815290604052805190602001209050919050565b60006131b6846000846118a6565b6131c257506000610c49565b6131ce836000846124fb565b5060019392505050565b60006131e58584846118a6565b6131f157506000613200565b6131fc8484846124fb565b5060015b949350505050565b600061321585848461228c565b61322157506000613200565b6131fc84848461287e565b600082820183811015610c49576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6008101590565b600083156132e7575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120613200565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115613385576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156133b2578160200160208202803883390190505b50805190915060005b8181101561340e576133cb61384a565b6133e78683815181106133da57fe5b6020026020010151612712565b905080600001518483815181106133fa57fe5b6020908102919091010152506001016133bb565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561345757818101518382015260200161343f565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000610c49604051806040016040528061349a86612ea0565b81526020016134a885612ea0565b905261372f565b60006134b9613897565b600083905060008582815181106134cc57fe5b602001015160f81c60f81b60f81c9050818060010192505060008683815181106134f257fe5b016020015160019384019360f89190911c915060009060ff8416141561359057600061351c6137ae565b6135268a87612c88565b90975090925090508115613581576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b61358a81612712565b51925050505b60006135a2898663ffffffff61371316565b90506020850194508360ff16600114156135e7576040805160808101825260ff909416845260208401919091526001908301526060820152919350909150612f829050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6136246137ae565b60408051608081018252600080825260208083018690528351828152908101845291928301919061366b565b6136586137ae565b8152602001906001900390816136505790505b508152600160209091015292915050565b6000818484111561370b5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156136d05781810151838201526020016136b8565b50505050905090810190601f1680156136fd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000816020018351101561372657600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b6137526137ae565b81526020019060019003908161374a575050805190915060005b818110156137a45784816002811061378057fe5b602002015183828151811061379157fe5b602090810291909101015260010161376c565b5061320082613335565b6040518060800160405280600081526020016137c8613897565b815260606020820152600060409091015290565b82805482825590600052602060002090810192821561381c5760005260206000209182015b8281111561381c578254825591600101919060010190613801565b506138289291506138be565b5090565b5080546000825590600052602060002090810190610f9091906138be565b60408051602081019091526000815290565b82805482825590600052602060002090810192821561381c579160200282015b8281111561381c57825182559160200191906001019061387c565b60408051608081018252600080825260208201819052918101829052606081019190915290565b61228991905b8082111561382857600081556001016138c456fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656e1d05879e5c1de5d46d38225763e79fd4094972abb444ec84ce9bffa0da68079da265627a7a72315820f13f975f98eeb60931ed720b436ef02fc7a62652acaf7d2cc3cbff3e869f9bea64736f6c634300050f0032"

// DeployGlobalPendingInbox deploys a new Ethereum contract, binding an instance of GlobalPendingInbox to it.
func DeployGlobalPendingInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalPendingInbox, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalPendingInboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	valueAddr, _, _, _ := DeployValue(auth, backend)
	GlobalPendingInboxBin = strings.Replace(GlobalPendingInboxBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalPendingInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalPendingInbox{GlobalPendingInboxCaller: GlobalPendingInboxCaller{contract: contract}, GlobalPendingInboxTransactor: GlobalPendingInboxTransactor{contract: contract}, GlobalPendingInboxFilterer: GlobalPendingInboxFilterer{contract: contract}}, nil
}

// GlobalPendingInbox is an auto generated Go binding around an Ethereum contract.
type GlobalPendingInbox struct {
	GlobalPendingInboxCaller     // Read-only binding to the contract
	GlobalPendingInboxTransactor // Write-only binding to the contract
	GlobalPendingInboxFilterer   // Log filterer for contract events
}

// GlobalPendingInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalPendingInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalPendingInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalPendingInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalPendingInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalPendingInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalPendingInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalPendingInboxSession struct {
	Contract     *GlobalPendingInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GlobalPendingInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalPendingInboxCallerSession struct {
	Contract *GlobalPendingInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GlobalPendingInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalPendingInboxTransactorSession struct {
	Contract     *GlobalPendingInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GlobalPendingInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalPendingInboxRaw struct {
	Contract *GlobalPendingInbox // Generic contract binding to access the raw methods on
}

// GlobalPendingInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalPendingInboxCallerRaw struct {
	Contract *GlobalPendingInboxCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalPendingInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalPendingInboxTransactorRaw struct {
	Contract *GlobalPendingInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalPendingInbox creates a new instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInbox(address common.Address, backend bind.ContractBackend) (*GlobalPendingInbox, error) {
	contract, err := bindGlobalPendingInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInbox{GlobalPendingInboxCaller: GlobalPendingInboxCaller{contract: contract}, GlobalPendingInboxTransactor: GlobalPendingInboxTransactor{contract: contract}, GlobalPendingInboxFilterer: GlobalPendingInboxFilterer{contract: contract}}, nil
}

// NewGlobalPendingInboxCaller creates a new read-only instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInboxCaller(address common.Address, caller bind.ContractCaller) (*GlobalPendingInboxCaller, error) {
	contract, err := bindGlobalPendingInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxCaller{contract: contract}, nil
}

// NewGlobalPendingInboxTransactor creates a new write-only instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalPendingInboxTransactor, error) {
	contract, err := bindGlobalPendingInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxTransactor{contract: contract}, nil
}

// NewGlobalPendingInboxFilterer creates a new log filterer instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalPendingInboxFilterer, error) {
	contract, err := bindGlobalPendingInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxFilterer{contract: contract}, nil
}

// bindGlobalPendingInbox binds a generic wrapper to an already deployed contract.
func bindGlobalPendingInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalPendingInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalPendingInbox *GlobalPendingInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalPendingInbox.Contract.GlobalPendingInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalPendingInbox *GlobalPendingInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GlobalPendingInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalPendingInbox *GlobalPendingInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GlobalPendingInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalPendingInbox *GlobalPendingInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalPendingInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalPendingInbox *GlobalPendingInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalPendingInbox *GlobalPendingInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.contract.Transact(opts, method, params...)
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x578c049a.
//
// Solidity: function getNFTTokens(address _owner) constant returns(address[], uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxCaller) GetNFTTokens(opts *bind.CallOpts, _owner common.Address) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GlobalPendingInbox.contract.Call(opts, out, "getNFTTokens", _owner)
	return *ret0, *ret1, err
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x578c049a.
//
// Solidity: function getNFTTokens(address _owner) constant returns(address[], uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetNFTTokens(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetNFTTokens(&_GlobalPendingInbox.CallOpts, _owner)
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x578c049a.
//
// Solidity: function getNFTTokens(address _owner) constant returns(address[], uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) GetNFTTokens(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetNFTTokens(&_GlobalPendingInbox.CallOpts, _owner)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxCaller) GetTokenBalance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "getTokenBalance", _tokenContract, _owner)
	return *ret0, err
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetTokenBalance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetTokenBalance(&_GlobalPendingInbox.CallOpts, _tokenContract, _owner)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) GetTokenBalance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetTokenBalance(&_GlobalPendingInbox.CallOpts, _tokenContract, _owner)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x764f3aa8.
//
// Solidity: function getTokenBalances(address _owner) constant returns(address[], uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxCaller) GetTokenBalances(opts *bind.CallOpts, _owner common.Address) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GlobalPendingInbox.contract.Call(opts, out, "getTokenBalances", _owner)
	return *ret0, *ret1, err
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x764f3aa8.
//
// Solidity: function getTokenBalances(address _owner) constant returns(address[], uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetTokenBalances(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetTokenBalances(&_GlobalPendingInbox.CallOpts, _owner)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x764f3aa8.
//
// Solidity: function getTokenBalances(address _owner) constant returns(address[], uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) GetTokenBalances(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetTokenBalances(&_GlobalPendingInbox.CallOpts, _owner)
}

// HasNFT is a free data retrieval call binding the contract method 0xe318b003.
//
// Solidity: function hasNFT(address _tokenContract, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalPendingInbox *GlobalPendingInboxCaller) HasNFT(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "hasNFT", _tokenContract, _owner, _tokenId)
	return *ret0, err
}

// HasNFT is a free data retrieval call binding the contract method 0xe318b003.
//
// Solidity: function hasNFT(address _tokenContract, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalPendingInbox *GlobalPendingInboxSession) HasNFT(_tokenContract common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalPendingInbox.Contract.HasNFT(&_GlobalPendingInbox.CallOpts, _tokenContract, _owner, _tokenId)
}

// HasNFT is a free data retrieval call binding the contract method 0xe318b003.
//
// Solidity: function hasNFT(address _tokenContract, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) HasNFT(_tokenContract common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalPendingInbox.Contract.HasNFT(&_GlobalPendingInbox.CallOpts, _tokenContract, _owner, _tokenId)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) DepositERC20Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "depositERC20Message", _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC20Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC20Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _id) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) DepositERC721Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc721 common.Address, _id *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "depositERC721Message", _chain, _to, _erc721, _id)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _id) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _id *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC721Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _id)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _id) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _id *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC721Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _id)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) DepositEthMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "depositEthMessage", _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositEthMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositEthMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) ForwardTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "forwardTransactionMessage", _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.ForwardTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.ForwardTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) GetPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "getPending")
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetPending() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GetPending(&_GlobalPendingInbox.TransactOpts)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) GetPending() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GetPending(&_GlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) RegisterForInbox(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "registerForInbox")
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) RegisterForInbox() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.RegisterForInbox(&_GlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) RegisterForInbox() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.RegisterForInbox(&_GlobalPendingInbox.TransactOpts)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _messages []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "sendMessages", _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendMessages(&_GlobalPendingInbox.TransactOpts, _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendMessages(&_GlobalPendingInbox.TransactOpts, _messages)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) SendTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "sendTransactionMessage", _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) WithdrawERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "withdrawERC20", _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC20(&_GlobalPendingInbox.TransactOpts, _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC20(&_GlobalPendingInbox.TransactOpts, _tokenContract)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) WithdrawERC721(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "withdrawERC721", _tokenContract, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) WithdrawERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC721(&_GlobalPendingInbox.TransactOpts, _tokenContract, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) WithdrawERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC721(&_GlobalPendingInbox.TransactOpts, _tokenContract, _tokenId)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawEth(&_GlobalPendingInbox.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawEth(&_GlobalPendingInbox.TransactOpts)
}

// GlobalPendingInboxAssertionEventIterator is returned from FilterAssertionEvent and is used to iterate over the raw logs and unpacked data for AssertionEvent events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxAssertionEventIterator struct {
	Event *GlobalPendingInboxAssertionEvent // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxAssertionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxAssertionEvent)
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
		it.Event = new(GlobalPendingInboxAssertionEvent)
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
func (it *GlobalPendingInboxAssertionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxAssertionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxAssertionEvent represents a AssertionEvent event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxAssertionEvent struct {
	Chain       common.Address
	Valid       bool
	MessageType *big.Int
	Destination common.Address
	Value       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionEvent is a free log retrieval operation binding the contract event 0x1d05879e5c1de5d46d38225763e79fd4094972abb444ec84ce9bffa0da68079d.
//
// Solidity: event AssertionEvent(address indexed chain, bool valid, uint256 messageType, address destination, string value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterAssertionEvent(opts *bind.FilterOpts, chain []common.Address) (*GlobalPendingInboxAssertionEventIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "AssertionEvent", chainRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxAssertionEventIterator{contract: _GlobalPendingInbox.contract, event: "AssertionEvent", logs: logs, sub: sub}, nil
}

// WatchAssertionEvent is a free log subscription operation binding the contract event 0x1d05879e5c1de5d46d38225763e79fd4094972abb444ec84ce9bffa0da68079d.
//
// Solidity: event AssertionEvent(address indexed chain, bool valid, uint256 messageType, address destination, string value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchAssertionEvent(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxAssertionEvent, chain []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "AssertionEvent", chainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxAssertionEvent)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "AssertionEvent", log); err != nil {
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

// ParseAssertionEvent is a log parse operation binding the contract event 0x1d05879e5c1de5d46d38225763e79fd4094972abb444ec84ce9bffa0da68079d.
//
// Solidity: event AssertionEvent(address indexed chain, bool valid, uint256 messageType, address destination, string value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseAssertionEvent(log types.Log) (*GlobalPendingInboxAssertionEvent, error) {
	event := new(GlobalPendingInboxAssertionEvent)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "AssertionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalPendingInboxERC20DepositMessageDeliveredIterator is returned from FilterERC20DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC20DepositMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC20DepositMessageDeliveredIterator struct {
	Event *GlobalPendingInboxERC20DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxERC20DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxERC20DepositMessageDelivered)
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
		it.Event = new(GlobalPendingInboxERC20DepositMessageDelivered)
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
func (it *GlobalPendingInboxERC20DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxERC20DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxERC20DepositMessageDelivered represents a ERC20DepositMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC20DepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Erc20 common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterERC20DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxERC20DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxERC20DepositMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "ERC20DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC20DepositMessageDelivered is a free log subscription operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchERC20DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxERC20DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxERC20DepositMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
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

// ParseERC20DepositMessageDelivered is a log parse operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseERC20DepositMessageDelivered(log types.Log) (*GlobalPendingInboxERC20DepositMessageDelivered, error) {
	event := new(GlobalPendingInboxERC20DepositMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalPendingInboxERC721DepositMessageDeliveredIterator is returned from FilterERC721DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC721DepositMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC721DepositMessageDeliveredIterator struct {
	Event *GlobalPendingInboxERC721DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxERC721DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxERC721DepositMessageDelivered)
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
		it.Event = new(GlobalPendingInboxERC721DepositMessageDelivered)
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
func (it *GlobalPendingInboxERC721DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxERC721DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxERC721DepositMessageDelivered represents a ERC721DepositMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC721DepositMessageDelivered struct {
	Chain  common.Address
	To     common.Address
	From   common.Address
	Erc721 common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC721DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterERC721DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxERC721DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxERC721DepositMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "ERC721DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC721DepositMessageDelivered is a free log subscription operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchERC721DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxERC721DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxERC721DepositMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
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

// ParseERC721DepositMessageDelivered is a log parse operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseERC721DepositMessageDelivered(log types.Log) (*GlobalPendingInboxERC721DepositMessageDelivered, error) {
	event := new(GlobalPendingInboxERC721DepositMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalPendingInboxEthDepositMessageDeliveredIterator is returned from FilterEthDepositMessageDelivered and is used to iterate over the raw logs and unpacked data for EthDepositMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxEthDepositMessageDeliveredIterator struct {
	Event *GlobalPendingInboxEthDepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxEthDepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxEthDepositMessageDelivered)
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
		it.Event = new(GlobalPendingInboxEthDepositMessageDelivered)
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
func (it *GlobalPendingInboxEthDepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxEthDepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxEthDepositMessageDelivered represents a EthDepositMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxEthDepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthDepositMessageDelivered is a free log retrieval operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterEthDepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxEthDepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxEthDepositMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "EthDepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchEthDepositMessageDelivered is a free log subscription operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchEthDepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxEthDepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxEthDepositMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
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

// ParseEthDepositMessageDelivered is a log parse operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseEthDepositMessageDelivered(log types.Log) (*GlobalPendingInboxEthDepositMessageDelivered, error) {
	event := new(GlobalPendingInboxEthDepositMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalPendingInboxTransactionMessageDeliveredIterator is returned from FilterTransactionMessageDelivered and is used to iterate over the raw logs and unpacked data for TransactionMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxTransactionMessageDeliveredIterator struct {
	Event *GlobalPendingInboxTransactionMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxTransactionMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxTransactionMessageDelivered)
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
		it.Event = new(GlobalPendingInboxTransactionMessageDelivered)
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
func (it *GlobalPendingInboxTransactionMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxTransactionMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxTransactionMessageDelivered represents a TransactionMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxTransactionMessageDelivered struct {
	Chain     common.Address
	To        common.Address
	From      common.Address
	SeqNumber *big.Int
	Value     *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransactionMessageDelivered is a free log retrieval operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterTransactionMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxTransactionMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxTransactionMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "TransactionMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchTransactionMessageDelivered is a free log subscription operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchTransactionMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxTransactionMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxTransactionMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
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

// ParseTransactionMessageDelivered is a log parse operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseTransactionMessageDelivered(log types.Log) (*GlobalPendingInboxTransactionMessageDelivered, error) {
	event := new(GlobalPendingInboxTransactionMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalWalletABI is the input ABI used to generate the binding from.
const GlobalWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getNFTTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getTokenBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalWalletFuncSigs = map[string]string{
	"578c049a": "getNFTTokens(address)",
	"c489744b": "getTokenBalance(address,address)",
	"764f3aa8": "getTokenBalances(address)",
	"e318b003": "hasNFT(address,address,uint256)",
	"f4f3b200": "withdrawERC20(address)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"a0ef91df": "withdrawEth()",
}

// GlobalWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalWalletBin = "0x608060405234801561001057600080fd5b50610dfc806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063c489744b1161005b578063c489744b14610171578063e318b003146101b1578063f3e414f8146101fb578063f4f3b200146102275761007d565b8063578c049a14610082578063764f3aa814610141578063a0ef91df14610167575b600080fd5b6100a86004803603602081101561009857600080fd5b50356001600160a01b031661024d565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156100ec5781810151838201526020016100d4565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561012b578181015183820152602001610113565b5050505090500194505050505060405180910390f35b6100a86004803603602081101561015757600080fd5b50356001600160a01b03166103e5565b61016f610519565b005b61019f6004803603604081101561018757600080fd5b506001600160a01b03813581169160200135166105a0565b60408051918252519081900360200190f35b6101e7600480360360608110156101c757600080fd5b506001600160a01b03813581169160208101359091169060400135610608565b604080519115158252519081900360200190f35b61016f6004803603604081101561021157600080fd5b506001600160a01b038135169060200135610689565b61016f6004803603602081101561023d57600080fd5b50356001600160a01b0316610755565b6001600160a01b0381166000908152602081905260408120606091829190805b60038301548110156102a95782600301818154811061028857fe5b6000918252602090912060026003909202010154919091019060010161026d565b6060826040519080825280602002602001820160405280156102d5578160200160208202803883390190505b509050606083604051908082528060200260200182016040528015610304578160200160208202803883390190505b5060038601546000945090915083905b808510156103d657600087600301868154811061032d57fe5b600091825260208220600260039092020190810154909250905b818110156103c857825487516001600160a01b039091169088908790811061036b57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505082600201818154811061039a57fe5b90600052602060002001548686815181106103b157fe5b602090810291909101015260019485019401610347565b505060019095019450610314565b50919650945050505050915091565b6001600160a01b03811660009081526020818152604091829020600181015483518181528184028101909301909352606092839283918015610431578160200160208202803883390190505b50905060608151604051908082528060200260200182016040528015610461578160200160208202803883390190505b50825190915060005b8181101561050c5784600101818154811061048157fe5b600091825260209091206002909102015484516001600160a01b03909116908590839081106104ac57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508460010181815481106104db57fe5b9060005260206000209060020201600101548382815181106104f957fe5b602090810291909101015260010161046a565b5091945092505050915091565b60006105266000336105a0565b905061053433600083610828565b61056f5760405162461bcd60e51b815260040180806020018281038252602e815260200180610d9a602e913960400191505060405180910390fd5b604051339082156108fc029083906000818181858888f1935050505015801561059c573d6000803e3d6000fd5b5050565b6001600160a01b03808216600090815260208181526040808320938616835290839052812054909190806105d957600092505050610602565b8160010160018203815481106105eb57fe5b906000526020600020906002020160010154925050505b92915050565b6001600160a01b038083166000908152602081815260408083209387168352600284019091528120549091908061064457600092505050610682565b81600301600182038154811061065657fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b6106943383836109c8565b6106e5576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b15801561073957600080fd5b505af115801561074d573d6000803e3d6000fd5b505050505050565b600061076182336105a0565b905061076e338383610828565b6107a95760405162461bcd60e51b815260040180806020018281038252602e815260200180610d9a602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b1580156107f857600080fd5b505af115801561080c573d6000803e3d6000fd5b505050506040513d602081101561082257600080fd5b50505050565b60008161083757506001610682565b6001600160a01b038085166000908152602081815260408083209387168352908390529020548061086d57600092505050610682565b600082600101600183038154811061088157fe5b9060005260206000209060020201905080600101548511156108a95760009350505050610682565b60018101546108be908663ffffffff610c3716565b600182018190556109bb57600183018054839185916000919060001981019081106108e557fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061092357fe5b906000526020600020906002020183600101600184038154811061094357fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061099157fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b6001600160a01b0380841660009081526020818152604080832093861683526002840190915281205490919080610a0457600092505050610682565b6000826003016001830381548110610a1857fe5b600091825260208083208884526001600390930201918201905260409091205490915080610a4d576000945050505050610682565b60028201805482916001850191600091906000198101908110610a6c57fe5b600091825260208083209091015483528201929092526040019020556002820180546000198101908110610a9c57fe5b9060005260206000200154826002016001830381548110610ab957fe5b600091825260208083209091019290925587815260018401909152604081205560028201805480610ae657fe5b6000828152602081208201600019908101919091550190556002820154610c295760038401805484916002870191600091906000198101908110610b2657fe5b60009182526020808320600392830201546001600160a01b031684528301939093526040909101902091909155840180546000198101908110610b6557fe5b9060005260206000209060030201846003016001850381548110610b8557fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b0390921691909117815560028083018054610bc89284019190610d0b565b5050506001600160a01b038716600090815260028501602052604081205560038401805480610bf357fe5b60008281526020812060036000199093019283020180546001600160a01b031916815590610c246002830182610d5b565b505090555b506001979650505050505050565b600061068283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060008184841115610d035760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610cc8578181015183820152602001610cb0565b50505050905090810190601f168015610cf55780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b828054828255906000526020600020908101928215610d4b5760005260206000209182015b82811115610d4b578254825591600101919060010190610d30565b50610d57929150610d7c565b5090565b5080546000825590600052602060002090810190610d799190610d7c565b50565b610d9691905b80821115610d575760008155600101610d82565b9056fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a723158202b8a733814556181bcad510f4c2339fa94782407ed3bfedc55ac775cdc8101e764736f6c634300050f0032"

// DeployGlobalWallet deploys a new Ethereum contract, binding an instance of GlobalWallet to it.
func DeployGlobalWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalWallet{GlobalWalletCaller: GlobalWalletCaller{contract: contract}, GlobalWalletTransactor: GlobalWalletTransactor{contract: contract}, GlobalWalletFilterer: GlobalWalletFilterer{contract: contract}}, nil
}

// GlobalWallet is an auto generated Go binding around an Ethereum contract.
type GlobalWallet struct {
	GlobalWalletCaller     // Read-only binding to the contract
	GlobalWalletTransactor // Write-only binding to the contract
	GlobalWalletFilterer   // Log filterer for contract events
}

// GlobalWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalWalletSession struct {
	Contract     *GlobalWallet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalWalletCallerSession struct {
	Contract *GlobalWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GlobalWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalWalletTransactorSession struct {
	Contract     *GlobalWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GlobalWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalWalletRaw struct {
	Contract *GlobalWallet // Generic contract binding to access the raw methods on
}

// GlobalWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalWalletCallerRaw struct {
	Contract *GlobalWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalWalletTransactorRaw struct {
	Contract *GlobalWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalWallet creates a new instance of GlobalWallet, bound to a specific deployed contract.
func NewGlobalWallet(address common.Address, backend bind.ContractBackend) (*GlobalWallet, error) {
	contract, err := bindGlobalWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalWallet{GlobalWalletCaller: GlobalWalletCaller{contract: contract}, GlobalWalletTransactor: GlobalWalletTransactor{contract: contract}, GlobalWalletFilterer: GlobalWalletFilterer{contract: contract}}, nil
}

// NewGlobalWalletCaller creates a new read-only instance of GlobalWallet, bound to a specific deployed contract.
func NewGlobalWalletCaller(address common.Address, caller bind.ContractCaller) (*GlobalWalletCaller, error) {
	contract, err := bindGlobalWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalWalletCaller{contract: contract}, nil
}

// NewGlobalWalletTransactor creates a new write-only instance of GlobalWallet, bound to a specific deployed contract.
func NewGlobalWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalWalletTransactor, error) {
	contract, err := bindGlobalWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalWalletTransactor{contract: contract}, nil
}

// NewGlobalWalletFilterer creates a new log filterer instance of GlobalWallet, bound to a specific deployed contract.
func NewGlobalWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalWalletFilterer, error) {
	contract, err := bindGlobalWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalWalletFilterer{contract: contract}, nil
}

// bindGlobalWallet binds a generic wrapper to an already deployed contract.
func bindGlobalWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalWallet *GlobalWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalWallet.Contract.GlobalWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalWallet *GlobalWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalWallet.Contract.GlobalWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalWallet *GlobalWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalWallet.Contract.GlobalWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalWallet *GlobalWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalWallet *GlobalWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalWallet *GlobalWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalWallet.Contract.contract.Transact(opts, method, params...)
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x578c049a.
//
// Solidity: function getNFTTokens(address _owner) constant returns(address[], uint256[])
func (_GlobalWallet *GlobalWalletCaller) GetNFTTokens(opts *bind.CallOpts, _owner common.Address) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GlobalWallet.contract.Call(opts, out, "getNFTTokens", _owner)
	return *ret0, *ret1, err
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x578c049a.
//
// Solidity: function getNFTTokens(address _owner) constant returns(address[], uint256[])
func (_GlobalWallet *GlobalWalletSession) GetNFTTokens(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalWallet.Contract.GetNFTTokens(&_GlobalWallet.CallOpts, _owner)
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x578c049a.
//
// Solidity: function getNFTTokens(address _owner) constant returns(address[], uint256[])
func (_GlobalWallet *GlobalWalletCallerSession) GetNFTTokens(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalWallet.Contract.GetNFTTokens(&_GlobalWallet.CallOpts, _owner)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalWallet *GlobalWalletCaller) GetTokenBalance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalWallet.contract.Call(opts, out, "getTokenBalance", _tokenContract, _owner)
	return *ret0, err
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalWallet *GlobalWalletSession) GetTokenBalance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalWallet.Contract.GetTokenBalance(&_GlobalWallet.CallOpts, _tokenContract, _owner)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalWallet *GlobalWalletCallerSession) GetTokenBalance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalWallet.Contract.GetTokenBalance(&_GlobalWallet.CallOpts, _tokenContract, _owner)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x764f3aa8.
//
// Solidity: function getTokenBalances(address _owner) constant returns(address[], uint256[])
func (_GlobalWallet *GlobalWalletCaller) GetTokenBalances(opts *bind.CallOpts, _owner common.Address) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GlobalWallet.contract.Call(opts, out, "getTokenBalances", _owner)
	return *ret0, *ret1, err
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x764f3aa8.
//
// Solidity: function getTokenBalances(address _owner) constant returns(address[], uint256[])
func (_GlobalWallet *GlobalWalletSession) GetTokenBalances(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalWallet.Contract.GetTokenBalances(&_GlobalWallet.CallOpts, _owner)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x764f3aa8.
//
// Solidity: function getTokenBalances(address _owner) constant returns(address[], uint256[])
func (_GlobalWallet *GlobalWalletCallerSession) GetTokenBalances(_owner common.Address) ([]common.Address, []*big.Int, error) {
	return _GlobalWallet.Contract.GetTokenBalances(&_GlobalWallet.CallOpts, _owner)
}

// HasNFT is a free data retrieval call binding the contract method 0xe318b003.
//
// Solidity: function hasNFT(address _tokenContract, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalWallet *GlobalWalletCaller) HasNFT(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalWallet.contract.Call(opts, out, "hasNFT", _tokenContract, _owner, _tokenId)
	return *ret0, err
}

// HasNFT is a free data retrieval call binding the contract method 0xe318b003.
//
// Solidity: function hasNFT(address _tokenContract, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalWallet *GlobalWalletSession) HasNFT(_tokenContract common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalWallet.Contract.HasNFT(&_GlobalWallet.CallOpts, _tokenContract, _owner, _tokenId)
}

// HasNFT is a free data retrieval call binding the contract method 0xe318b003.
//
// Solidity: function hasNFT(address _tokenContract, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalWallet *GlobalWalletCallerSession) HasNFT(_tokenContract common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalWallet.Contract.HasNFT(&_GlobalWallet.CallOpts, _tokenContract, _owner, _tokenId)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalWallet *GlobalWalletTransactor) WithdrawERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalWallet.contract.Transact(opts, "withdrawERC20", _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalWallet *GlobalWalletSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalWallet.Contract.WithdrawERC20(&_GlobalWallet.TransactOpts, _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalWallet *GlobalWalletTransactorSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalWallet.Contract.WithdrawERC20(&_GlobalWallet.TransactOpts, _tokenContract)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_GlobalWallet *GlobalWalletTransactor) WithdrawERC721(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalWallet.contract.Transact(opts, "withdrawERC721", _tokenContract, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_GlobalWallet *GlobalWalletSession) WithdrawERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalWallet.Contract.WithdrawERC721(&_GlobalWallet.TransactOpts, _tokenContract, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_GlobalWallet *GlobalWalletTransactorSession) WithdrawERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalWallet.Contract.WithdrawERC721(&_GlobalWallet.TransactOpts, _tokenContract, _tokenId)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalWallet *GlobalWalletTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalWallet.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalWallet *GlobalWalletSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalWallet.Contract.WithdrawEth(&_GlobalWallet.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalWallet *GlobalWalletTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalWallet.Contract.WithdrawEth(&_GlobalWallet.TransactOpts)
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ABI is the input ABI used to generate the binding from.
const IERC721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721FuncSigs maps the 4-byte function signature to its string representation.
var IERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxABI is the input ABI used to generate the binding from.
const IGlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"AssertionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ERC20DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ERC721DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EthDepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionMessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getPending\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var IGlobalPendingInboxFuncSigs = map[string]string{
	"bca22b76": "depositERC20Message(address,address,address,uint256)",
	"8b7010aa": "depositERC721Message(address,address,address,uint256)",
	"5bd21290": "depositEthMessage(address,address)",
	"8bef8df0": "forwardTransactionMessage(address,address,uint256,uint256,bytes,bytes)",
	"11ae9ed2": "getPending()",
	"f3972383": "registerForInbox()",
	"e4eb8c63": "sendMessages(bytes)",
	"8f5ed73e": "sendTransactionMessage(address,address,uint256,uint256,bytes)",
}

// IGlobalPendingInbox is an auto generated Go binding around an Ethereum contract.
type IGlobalPendingInbox struct {
	IGlobalPendingInboxCaller     // Read-only binding to the contract
	IGlobalPendingInboxTransactor // Write-only binding to the contract
	IGlobalPendingInboxFilterer   // Log filterer for contract events
}

// IGlobalPendingInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGlobalPendingInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGlobalPendingInboxSession struct {
	Contract     *IGlobalPendingInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGlobalPendingInboxCallerSession struct {
	Contract *IGlobalPendingInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IGlobalPendingInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGlobalPendingInboxTransactorSession struct {
	Contract     *IGlobalPendingInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGlobalPendingInboxRaw struct {
	Contract *IGlobalPendingInbox // Generic contract binding to access the raw methods on
}

// IGlobalPendingInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCallerRaw struct {
	Contract *IGlobalPendingInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IGlobalPendingInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactorRaw struct {
	Contract *IGlobalPendingInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGlobalPendingInbox creates a new instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInbox(address common.Address, backend bind.ContractBackend) (*IGlobalPendingInbox, error) {
	contract, err := bindIGlobalPendingInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInbox{IGlobalPendingInboxCaller: IGlobalPendingInboxCaller{contract: contract}, IGlobalPendingInboxTransactor: IGlobalPendingInboxTransactor{contract: contract}, IGlobalPendingInboxFilterer: IGlobalPendingInboxFilterer{contract: contract}}, nil
}

// NewIGlobalPendingInboxCaller creates a new read-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxCaller(address common.Address, caller bind.ContractCaller) (*IGlobalPendingInboxCaller, error) {
	contract, err := bindIGlobalPendingInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxCaller{contract: contract}, nil
}

// NewIGlobalPendingInboxTransactor creates a new write-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IGlobalPendingInboxTransactor, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxTransactor{contract: contract}, nil
}

// NewIGlobalPendingInboxFilterer creates a new log filterer instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IGlobalPendingInboxFilterer, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxFilterer{contract: contract}, nil
}

// bindIGlobalPendingInbox binds a generic wrapper to an already deployed contract.
func bindIGlobalPendingInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGlobalPendingInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transact(opts, method, params...)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) DepositERC20Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "depositERC20Message", _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC20Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC20Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) DepositERC721Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc721 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "depositERC721Message", _chain, _to, _erc721, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC721Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC721Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _value)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) DepositEthMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "depositEthMessage", _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositEthMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositEthMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) ForwardTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "forwardTransactionMessage", _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) GetPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "getPending")
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) GetPending() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.GetPending(&_IGlobalPendingInbox.TransactOpts)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) GetPending() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.GetPending(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) RegisterForInbox(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "registerForInbox")
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessages", _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendTransactionMessage", _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// IGlobalPendingInboxAssertionEventIterator is returned from FilterAssertionEvent and is used to iterate over the raw logs and unpacked data for AssertionEvent events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxAssertionEventIterator struct {
	Event *IGlobalPendingInboxAssertionEvent // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxAssertionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxAssertionEvent)
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
		it.Event = new(IGlobalPendingInboxAssertionEvent)
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
func (it *IGlobalPendingInboxAssertionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxAssertionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxAssertionEvent represents a AssertionEvent event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxAssertionEvent struct {
	Chain       common.Address
	Valid       bool
	MessageType *big.Int
	Destination common.Address
	Value       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionEvent is a free log retrieval operation binding the contract event 0x1d05879e5c1de5d46d38225763e79fd4094972abb444ec84ce9bffa0da68079d.
//
// Solidity: event AssertionEvent(address indexed chain, bool valid, uint256 messageType, address destination, string value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterAssertionEvent(opts *bind.FilterOpts, chain []common.Address) (*IGlobalPendingInboxAssertionEventIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "AssertionEvent", chainRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxAssertionEventIterator{contract: _IGlobalPendingInbox.contract, event: "AssertionEvent", logs: logs, sub: sub}, nil
}

// WatchAssertionEvent is a free log subscription operation binding the contract event 0x1d05879e5c1de5d46d38225763e79fd4094972abb444ec84ce9bffa0da68079d.
//
// Solidity: event AssertionEvent(address indexed chain, bool valid, uint256 messageType, address destination, string value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchAssertionEvent(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxAssertionEvent, chain []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "AssertionEvent", chainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxAssertionEvent)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "AssertionEvent", log); err != nil {
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

// ParseAssertionEvent is a log parse operation binding the contract event 0x1d05879e5c1de5d46d38225763e79fd4094972abb444ec84ce9bffa0da68079d.
//
// Solidity: event AssertionEvent(address indexed chain, bool valid, uint256 messageType, address destination, string value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseAssertionEvent(log types.Log) (*IGlobalPendingInboxAssertionEvent, error) {
	event := new(IGlobalPendingInboxAssertionEvent)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "AssertionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxERC20DepositMessageDeliveredIterator is returned from FilterERC20DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC20DepositMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC20DepositMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxERC20DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxERC20DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxERC20DepositMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxERC20DepositMessageDelivered)
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
func (it *IGlobalPendingInboxERC20DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxERC20DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxERC20DepositMessageDelivered represents a ERC20DepositMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC20DepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Erc20 common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterERC20DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxERC20DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxERC20DepositMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "ERC20DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC20DepositMessageDelivered is a free log subscription operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchERC20DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxERC20DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxERC20DepositMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
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

// ParseERC20DepositMessageDelivered is a log parse operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseERC20DepositMessageDelivered(log types.Log) (*IGlobalPendingInboxERC20DepositMessageDelivered, error) {
	event := new(IGlobalPendingInboxERC20DepositMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxERC721DepositMessageDeliveredIterator is returned from FilterERC721DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC721DepositMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC721DepositMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxERC721DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxERC721DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxERC721DepositMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxERC721DepositMessageDelivered)
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
func (it *IGlobalPendingInboxERC721DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxERC721DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxERC721DepositMessageDelivered represents a ERC721DepositMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC721DepositMessageDelivered struct {
	Chain  common.Address
	To     common.Address
	From   common.Address
	Erc721 common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC721DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterERC721DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxERC721DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxERC721DepositMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "ERC721DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC721DepositMessageDelivered is a free log subscription operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchERC721DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxERC721DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxERC721DepositMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
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

// ParseERC721DepositMessageDelivered is a log parse operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseERC721DepositMessageDelivered(log types.Log) (*IGlobalPendingInboxERC721DepositMessageDelivered, error) {
	event := new(IGlobalPendingInboxERC721DepositMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxEthDepositMessageDeliveredIterator is returned from FilterEthDepositMessageDelivered and is used to iterate over the raw logs and unpacked data for EthDepositMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxEthDepositMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxEthDepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxEthDepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxEthDepositMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxEthDepositMessageDelivered)
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
func (it *IGlobalPendingInboxEthDepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxEthDepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxEthDepositMessageDelivered represents a EthDepositMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxEthDepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthDepositMessageDelivered is a free log retrieval operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterEthDepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxEthDepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxEthDepositMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "EthDepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchEthDepositMessageDelivered is a free log subscription operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchEthDepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxEthDepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxEthDepositMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
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

// ParseEthDepositMessageDelivered is a log parse operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseEthDepositMessageDelivered(log types.Log) (*IGlobalPendingInboxEthDepositMessageDelivered, error) {
	event := new(IGlobalPendingInboxEthDepositMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxTransactionMessageDeliveredIterator is returned from FilterTransactionMessageDelivered and is used to iterate over the raw logs and unpacked data for TransactionMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxTransactionMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxTransactionMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxTransactionMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxTransactionMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxTransactionMessageDelivered)
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
func (it *IGlobalPendingInboxTransactionMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxTransactionMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxTransactionMessageDelivered represents a TransactionMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxTransactionMessageDelivered struct {
	Chain     common.Address
	To        common.Address
	From      common.Address
	SeqNumber *big.Int
	Value     *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransactionMessageDelivered is a free log retrieval operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterTransactionMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxTransactionMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxTransactionMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "TransactionMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchTransactionMessageDelivered is a free log subscription operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchTransactionMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxTransactionMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxTransactionMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
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

// ParseTransactionMessageDelivered is a log parse operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseTransactionMessageDelivered(log types.Log) (*IGlobalPendingInboxTransactionMessageDelivered, error) {
	event := new(IGlobalPendingInboxTransactionMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a6aef2a5c100dbc457a0564b840718120d6130fa3222baa93903a30cc7e6c0b564736f6c634300050f0032"

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201ca52c8cfd94848edb89c2f757752003702d51eace6d9e6123cdcc328cbf2aca64736f6c634300050f0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[]"

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820f1dd3f86360fcfbddb558ec0503d02c14ac6271a08bdee1156a0317e67af40fb64736f6c634300050f0032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getERCTokenMsgData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tokenAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getEthMsgData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getTransactionMsgData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vmAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueFuncSigs maps the 4-byte function signature to its string representation.
var ValueFuncSigs = map[string]string{
	"bf532221": "getERCTokenMsgData(bytes)",
	"ac715d70": "getEthMsgData(bytes)",
	"fef067dc": "getTransactionMsgData(bytes)",
}

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x611163610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063ac715d7014610050578063bf53222114610116578063fef067dc146101e4575b600080fd5b6100f66004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460018302840111640100000000831117156100b557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610329945050505050565b604080519315158452602084019290925282820152519081900360600190f35b6101bc6004803603602081101561012c57600080fd5b81019060208101813564010000000081111561014757600080fd5b82018360208201111561015957600080fd5b8035906020019184600183028401116401000000008311171561017b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610426945050505050565b6040805194151585526020850193909352838301919091526060830152519081900360800190f35b61028a600480360360208110156101fa57600080fd5b81019060208101813564010000000081111561021557600080fd5b82018360208201111561022757600080fd5b8035906020019184600183028401116401000000008311171561024957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610569945050505050565b604051808715151515815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102e95781810151838201526020016102d1565b50505050905090810190601f1680156103165780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b60008060008060009050600085828151811061034157fe5b016020015160019092019160f81c905060058114156104185760021981016060600061036e838a8761065e565b90965092509050801580156103a15750600060ff168260008151811061039057fe5b60200260200101516060015160ff16145b80156103cb5750600060ff16826001815181106103ba57fe5b60200260200101516060015160ff16145b15610414576001826000815181106103df57fe5b602002602001015160000151836001815181106103f857fe5b602002602001015160000151975097509750505050505061041f565b5050505b5060009350505b9193909250565b6000806000806000809050600086828151811061043f57fe5b016020015160019092019160f81c9050600681141561055b5760021981016060600061046c838b8761065e565b909650925090508015801561049f5750600060ff168260008151811061048e57fe5b60200260200101516060015160ff16145b80156104c95750600060ff16826001815181106104b857fe5b60200260200101516060015160ff16145b80156104f35750600060ff16826002815181106104e257fe5b60200260200101516060015160ff16145b156105575760018260008151811061050757fe5b6020026020010151600001518360018151811061052057fe5b6020026020010151600001518460028151811061053957fe5b60200260200101516000015198509850985098505050505050610562565b5050505b5060009450505b9193509193565b600080600080600060606000809050600088828151811061058657fe5b016020015160019092019160f81c90506007811415610653578882815181106105ab57fe5b016020015160019092019160f81c90506105c58983610719565b8a5191975092508990839081106105d857fe5b016020015160019092019160f81c90506105f28983610719565b8a51919650925089908390811061060557fe5b016020015160019092019160f81c905061061f8983610719565b909450915060006106308a84610742565b909350905061064b8a6001600019860163ffffffff6107cd16565b935060019850505b505091939550919395565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156106a957816020015b6106966110c7565b81526020019060019003908161068e5790505b50905060005b8960ff168160ff161015610703576106c7898561084d565b8451859060ff86169081106106d857fe5b60209081029190910101529450925082156106fb57509094509092509050610710565b6001016106af565b5060009550919350909150505b93509350939050565b600080828161072e868363ffffffff6109d716565b6020929092019350909150505b9250929050565b6000806000806107506110c7565b61075a878761084d565b9194509250905082156107b4576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b816107be826109f3565b51909890975095505050505050565b6060818301845110156107df57600080fd5b6060821580156107fa57604051915060208201604052610844565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561083357805183526020928301920161081b565b5050858452601f01601f1916604052505b50949350505050565b6000806108586110c7565b845184106108ad576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b600084905060008682815181106108c057fe5b016020015160019092019160f81c905060006108da6110f5565b60ff831661090e576108ec8985610719565b90945091506000846108fd84610b2e565b919850965094506109d09350505050565b60ff831660011415610935576109248985610bac565b90945090506000846108fd83610d19565b60ff83166002141561095c5761094b8985610719565b90945091506000846108fd84610d79565b600360ff8416108015906109735750600c60ff8416105b156109b05760021983016060600061098c838d8961065e565b90985092509050808761099e84610df7565b995099509950505050505050506109d0565b8260ff166127100160006109c46000610b2e565b91985096509450505050505b9250925092565b600081602001835110156109ea57600080fd5b50016020015190565b6109fb61111c565b6060820151600c60ff90911610610a4d576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610a7a576040518060200160405280610a718460000151610ea7565b90529050610b29565b606082015160ff1660011415610ac1576040518060200160405280610a71846020015160000151856020015160400151866020015160600151876020015160200151610ecb565b606082015160ff1660021415610ae65750604080516020810190915281518152610b29565b600360ff16826060015160ff1610158015610b0a57506060820151600c60ff909116105b15610b27576040518060200160405280610a718460400151610f74565bfe5b919050565b610b366110c7565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610b9b565b610b886110c7565b815260200190600190039081610b805790505b508152600060209091015292915050565b6000610bb66110f5565b60008390506000858281518110610bc957fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610bef57fe5b016020015160019384019360f89190911c915060009060ff84161415610c8d576000610c196110c7565b610c238a8761084d565b90975090925090508115610c7e576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b610c87816109f3565b51925050505b6000610c9f898663ffffffff6109d716565b90506020850194508360ff1660011415610ce4576040805160808101825260ff90941684526020840191909152600190830152606082015291935090915061073b9050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b610d216110c7565b604080516080810182526000808252602080830186905283518281529081018452919283019190610d68565b610d556110c7565b815260200190600190039081610d4d5790505b508152600160209091015292915050565b610d816110c7565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610de6565b610dd36110c7565b815260200190600190039081610dcb5790505b508152600260209091015292915050565b610dff6110c7565b610e0982516110c0565b610e5a576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315610f25575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120610f6c565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6000600882511115610fc4576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610ff1578160200160208202803883390190505b50805190915060005b8181101561104d5761100a61111c565b61102686838151811061101957fe5b60200260200101516109f3565b9050806000015184838151811061103957fe5b602090810291909101015250600101610ffa565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561109657818101518382015260200161107e565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6008101590565b6040518060800160405280600081526020016110e16110f5565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fea265627a7a723158209e0428e75c462832bc8b3930744d3fbe41a3f33ae2126a915014ae5fa6e7cab564736f6c634300050f0032"

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

// GetERCTokenMsgData is a free data retrieval call binding the contract method 0xbf532221.
//
// Solidity: function getERCTokenMsgData(bytes data) constant returns(bool valid, uint256 tokenAddress, uint256 destination, uint256 value)
func (_Value *ValueCaller) GetERCTokenMsgData(opts *bind.CallOpts, data []byte) (struct {
	Valid        bool
	TokenAddress *big.Int
	Destination  *big.Int
	Value        *big.Int
}, error) {
	ret := new(struct {
		Valid        bool
		TokenAddress *big.Int
		Destination  *big.Int
		Value        *big.Int
	})
	out := ret
	err := _Value.contract.Call(opts, out, "getERCTokenMsgData", data)
	return *ret, err
}

// GetERCTokenMsgData is a free data retrieval call binding the contract method 0xbf532221.
//
// Solidity: function getERCTokenMsgData(bytes data) constant returns(bool valid, uint256 tokenAddress, uint256 destination, uint256 value)
func (_Value *ValueSession) GetERCTokenMsgData(data []byte) (struct {
	Valid        bool
	TokenAddress *big.Int
	Destination  *big.Int
	Value        *big.Int
}, error) {
	return _Value.Contract.GetERCTokenMsgData(&_Value.CallOpts, data)
}

// GetERCTokenMsgData is a free data retrieval call binding the contract method 0xbf532221.
//
// Solidity: function getERCTokenMsgData(bytes data) constant returns(bool valid, uint256 tokenAddress, uint256 destination, uint256 value)
func (_Value *ValueCallerSession) GetERCTokenMsgData(data []byte) (struct {
	Valid        bool
	TokenAddress *big.Int
	Destination  *big.Int
	Value        *big.Int
}, error) {
	return _Value.Contract.GetERCTokenMsgData(&_Value.CallOpts, data)
}

// GetEthMsgData is a free data retrieval call binding the contract method 0xac715d70.
//
// Solidity: function getEthMsgData(bytes data) constant returns(bool valid, uint256 destination, uint256 value)
func (_Value *ValueCaller) GetEthMsgData(opts *bind.CallOpts, data []byte) (struct {
	Valid       bool
	Destination *big.Int
	Value       *big.Int
}, error) {
	ret := new(struct {
		Valid       bool
		Destination *big.Int
		Value       *big.Int
	})
	out := ret
	err := _Value.contract.Call(opts, out, "getEthMsgData", data)
	return *ret, err
}

// GetEthMsgData is a free data retrieval call binding the contract method 0xac715d70.
//
// Solidity: function getEthMsgData(bytes data) constant returns(bool valid, uint256 destination, uint256 value)
func (_Value *ValueSession) GetEthMsgData(data []byte) (struct {
	Valid       bool
	Destination *big.Int
	Value       *big.Int
}, error) {
	return _Value.Contract.GetEthMsgData(&_Value.CallOpts, data)
}

// GetEthMsgData is a free data retrieval call binding the contract method 0xac715d70.
//
// Solidity: function getEthMsgData(bytes data) constant returns(bool valid, uint256 destination, uint256 value)
func (_Value *ValueCallerSession) GetEthMsgData(data []byte) (struct {
	Valid       bool
	Destination *big.Int
	Value       *big.Int
}, error) {
	return _Value.Contract.GetEthMsgData(&_Value.CallOpts, data)
}

// GetTransactionMsgData is a free data retrieval call binding the contract method 0xfef067dc.
//
// Solidity: function getTransactionMsgData(bytes data) constant returns(bool valid, uint256 vmAddress, uint256 destination, uint256 seqNumber, uint256 value, bytes messageData)
func (_Value *ValueCaller) GetTransactionMsgData(opts *bind.CallOpts, data []byte) (struct {
	Valid       bool
	VmAddress   *big.Int
	Destination *big.Int
	SeqNumber   *big.Int
	Value       *big.Int
	MessageData []byte
}, error) {
	ret := new(struct {
		Valid       bool
		VmAddress   *big.Int
		Destination *big.Int
		SeqNumber   *big.Int
		Value       *big.Int
		MessageData []byte
	})
	out := ret
	err := _Value.contract.Call(opts, out, "getTransactionMsgData", data)
	return *ret, err
}

// GetTransactionMsgData is a free data retrieval call binding the contract method 0xfef067dc.
//
// Solidity: function getTransactionMsgData(bytes data) constant returns(bool valid, uint256 vmAddress, uint256 destination, uint256 seqNumber, uint256 value, bytes messageData)
func (_Value *ValueSession) GetTransactionMsgData(data []byte) (struct {
	Valid       bool
	VmAddress   *big.Int
	Destination *big.Int
	SeqNumber   *big.Int
	Value       *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.GetTransactionMsgData(&_Value.CallOpts, data)
}

// GetTransactionMsgData is a free data retrieval call binding the contract method 0xfef067dc.
//
// Solidity: function getTransactionMsgData(bytes data) constant returns(bool valid, uint256 vmAddress, uint256 destination, uint256 seqNumber, uint256 value, bytes messageData)
func (_Value *ValueCallerSession) GetTransactionMsgData(data []byte) (struct {
	Valid       bool
	VmAddress   *big.Int
	Destination *big.Int
	SeqNumber   *big.Int
	Value       *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.GetTransactionMsgData(&_Value.CallOpts, data)
}
