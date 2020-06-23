// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package protocoltester

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

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208216d4682014e8cd8a0899329bca2ffb7fd563e94ecf7caa96ce6756cf76d5d864736f6c634300050f0032"

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

// ProtocolTesterABI is the input ABI used to generate the binding from.
const ProtocolTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInboxHash\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProtocolTesterFuncSigs maps the 4-byte function signature to its string representation.
var ProtocolTesterFuncSigs = map[string]string{
	"c21ef0fa": "generateAssertionHash(bytes32,bool,uint64,bytes32,bytes32,bytes32,bytes32)",
	"02be0bd0": "generateLastMessageHash(bytes,uint256,uint256)",
	"fcc6af6e": "generatePreconditionHash(bytes32,uint128[4],bytes32)",
}

// ProtocolTesterBin is the compiled bytecode used for deploying new contracts.
var ProtocolTesterBin = "0x608060405234801561001057600080fd5b506104f1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806302be0bd014610046578063c21ef0fa1461010a578063fcc6af6e14610169575b600080fd5b6100f16004803603606081101561005c57600080fd5b81019060208101813564010000000081111561007757600080fd5b82018360208201111561008957600080fd5b803590602001918460018302840111640100000000831117156100ab57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356101c1565b6040805192835260208301919091528051918290030190f35b610157600480360360e081101561012057600080fd5b50803590602081013515159067ffffffffffffffff6040820135169060608101359060808101359060a08101359060c001356101db565b60408051918252519081900360200190f35b610157600480360360c081101561017f57600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091945050903591506101f89050565b6000806101cf85858561020d565b91509150935093915050565b60006101ec88888888888888610368565b98975050505050505050565b60006102058484846103cb565b949350505050565b8251600090819081908190819087870190811115610263576040805162461bcd60e51b815260206004820152600e60248201526d0d2dcecc2d8d2c840d8cadccee8d60931b604482015290519081900360640190fd5b8760005b82821015610301576102798b83610445565b9197509094509150856102cc576040805162461bcd60e51b8152602060048201526016602482015275496e76616c6964206f7574707574206d65737361676560501b604482015290519081900360640190fd5b604080516020808201979097528082018690528151808203830181526060909101909152805195019490942093600101610267565b888a018214610357576040805162461bcd60e51b815260206004820152601a60248201527f76616c756520657874656e6465642070617374206c656e677468000000000000604482015290519081900360640190fd5b939a93995092975050505050505050565b6040805160208082019990995296151560f81b8782015260c09590951b6001600160c01b031916604187015260498601939093526069850191909152608984015260a9808401919091528151808403909101815260c99092019052805191012090565b81516020808401516040808601516060968701518251808601999099526fffffffffffffffffffffffffffffffff19608096871b81168a85015293861b841660508a015290851b83169688019690965294831b166070860152818501929092528251808503909101815260a0909301909152815191012090565b6000806000806000865190508581108061046157506020868203105b156104755750600093508492509050610499565b610485878763ffffffff6104a016565b600195506020870194509250610499915050565b9250925092565b600081602001835110156104b357600080fd5b5001602001519056fea265627a7a723158200f25665ee2462dbef4f344eff71ba91d467b0a1750a49e5c801eb5c4abb6205b64736f6c634300050f0032"

// DeployProtocolTester deploys a new Ethereum contract, binding an instance of ProtocolTester to it.
func DeployProtocolTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProtocolTester, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtocolTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProtocolTester{ProtocolTesterCaller: ProtocolTesterCaller{contract: contract}, ProtocolTesterTransactor: ProtocolTesterTransactor{contract: contract}, ProtocolTesterFilterer: ProtocolTesterFilterer{contract: contract}}, nil
}

// ProtocolTester is an auto generated Go binding around an Ethereum contract.
type ProtocolTester struct {
	ProtocolTesterCaller     // Read-only binding to the contract
	ProtocolTesterTransactor // Write-only binding to the contract
	ProtocolTesterFilterer   // Log filterer for contract events
}

// ProtocolTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolTesterSession struct {
	Contract     *ProtocolTester   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolTesterCallerSession struct {
	Contract *ProtocolTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ProtocolTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTesterTransactorSession struct {
	Contract     *ProtocolTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ProtocolTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolTesterRaw struct {
	Contract *ProtocolTester // Generic contract binding to access the raw methods on
}

// ProtocolTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolTesterCallerRaw struct {
	Contract *ProtocolTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTesterTransactorRaw struct {
	Contract *ProtocolTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolTester creates a new instance of ProtocolTester, bound to a specific deployed contract.
func NewProtocolTester(address common.Address, backend bind.ContractBackend) (*ProtocolTester, error) {
	contract, err := bindProtocolTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolTester{ProtocolTesterCaller: ProtocolTesterCaller{contract: contract}, ProtocolTesterTransactor: ProtocolTesterTransactor{contract: contract}, ProtocolTesterFilterer: ProtocolTesterFilterer{contract: contract}}, nil
}

// NewProtocolTesterCaller creates a new read-only instance of ProtocolTester, bound to a specific deployed contract.
func NewProtocolTesterCaller(address common.Address, caller bind.ContractCaller) (*ProtocolTesterCaller, error) {
	contract, err := bindProtocolTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTesterCaller{contract: contract}, nil
}

// NewProtocolTesterTransactor creates a new write-only instance of ProtocolTester, bound to a specific deployed contract.
func NewProtocolTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTesterTransactor, error) {
	contract, err := bindProtocolTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTesterTransactor{contract: contract}, nil
}

// NewProtocolTesterFilterer creates a new log filterer instance of ProtocolTester, bound to a specific deployed contract.
func NewProtocolTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolTesterFilterer, error) {
	contract, err := bindProtocolTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolTesterFilterer{contract: contract}, nil
}

// bindProtocolTester binds a generic wrapper to an already deployed contract.
func bindProtocolTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolTester *ProtocolTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProtocolTester.Contract.ProtocolTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolTester *ProtocolTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolTester.Contract.ProtocolTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolTester *ProtocolTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolTester.Contract.ProtocolTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolTester *ProtocolTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProtocolTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolTester *ProtocolTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolTester *ProtocolTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolTester.Contract.contract.Transact(opts, method, params...)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0xc21ef0fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ProtocolTester *ProtocolTesterCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _didInboxInsn bool, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ProtocolTester.contract.Call(opts, out, "generateAssertionHash", _afterHash, _didInboxInsn, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0xc21ef0fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ProtocolTester *ProtocolTesterSession) GenerateAssertionHash(_afterHash [32]byte, _didInboxInsn bool, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GenerateAssertionHash(&_ProtocolTester.CallOpts, _afterHash, _didInboxInsn, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0xc21ef0fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ProtocolTester *ProtocolTesterCallerSession) GenerateAssertionHash(_afterHash [32]byte, _didInboxInsn bool, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GenerateAssertionHash(&_ProtocolTester.CallOpts, _afterHash, _didInboxInsn, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) constant returns(bytes32, uint256)
func (_ProtocolTester *ProtocolTesterCaller) GenerateLastMessageHash(opts *bind.CallOpts, messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ProtocolTester.contract.Call(opts, out, "generateLastMessageHash", messages, startOffset, length)
	return *ret0, *ret1, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) constant returns(bytes32, uint256)
func (_ProtocolTester *ProtocolTesterSession) GenerateLastMessageHash(messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	return _ProtocolTester.Contract.GenerateLastMessageHash(&_ProtocolTester.CallOpts, messages, startOffset, length)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) constant returns(bytes32, uint256)
func (_ProtocolTester *ProtocolTesterCallerSession) GenerateLastMessageHash(messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	return _ProtocolTester.Contract.GenerateLastMessageHash(&_ProtocolTester.CallOpts, messages, startOffset, length)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0xfcc6af6e.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint128[4] _timeBounds, bytes32 _beforeInboxHash) constant returns(bytes32)
func (_ProtocolTester *ProtocolTesterCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [4]*big.Int, _beforeInboxHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ProtocolTester.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInboxHash)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0xfcc6af6e.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint128[4] _timeBounds, bytes32 _beforeInboxHash) constant returns(bytes32)
func (_ProtocolTester *ProtocolTesterSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [4]*big.Int, _beforeInboxHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GeneratePreconditionHash(&_ProtocolTester.CallOpts, _beforeHash, _timeBounds, _beforeInboxHash)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0xfcc6af6e.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint128[4] _timeBounds, bytes32 _beforeInboxHash) constant returns(bytes32)
func (_ProtocolTester *ProtocolTesterCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [4]*big.Int, _beforeInboxHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GeneratePreconditionHash(&_ProtocolTester.CallOpts, _beforeHash, _timeBounds, _beforeInboxHash)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a1a1c0c60e169357bb80e951ac1cb57368ef18899e17b41422edaf115ce97ae764736f6c634300050f0032"

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
