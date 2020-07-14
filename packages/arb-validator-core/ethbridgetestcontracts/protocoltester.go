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

// ProtocolTesterABI is the input ABI used to generate the binding from.
const ProtocolTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInboxHash\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProtocolTesterFuncSigs maps the 4-byte function signature to its string representation.
var ProtocolTesterFuncSigs = map[string]string{
	"c21ef0fa": "generateAssertionHash(bytes32,bool,uint64,bytes32,bytes32,bytes32,bytes32)",
	"9353b9b4": "generatePreconditionHash(bytes32,bytes32)",
}

// ProtocolTesterBin is the compiled bytecode used for deploying new contracts.
var ProtocolTesterBin = "0x608060405234801561001057600080fd5b506101b1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80639353b9b41461003b578063c21ef0fa14610070575b600080fd5b61005e6004803603604081101561005157600080fd5b50803590602001356100bd565b60408051918252519081900360200190f35b61005e600480360360e081101561008657600080fd5b50803590602081013515159067ffffffffffffffff6040820135169060608101359060808101359060a08101359060c001356100d0565b60006100c983836100ed565b9392505050565b60006100e188888888888888610119565b98975050505050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6040805160208082019990995296151560f81b8782015260c09590951b6001600160c01b031916604187015260498601939093526069850191909152608984015260a9808401919091528151808403909101815260c9909201905280519101209056fea265627a7a72315820725e874cbcfa1fa7428f9d950b9e680c9c13fb514e1092c90c18afb44535826b64736f6c63430005110032"

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
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) pure returns(bytes32)
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
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) pure returns(bytes32)
func (_ProtocolTester *ProtocolTesterSession) GenerateAssertionHash(_afterHash [32]byte, _didInboxInsn bool, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GenerateAssertionHash(&_ProtocolTester.CallOpts, _afterHash, _didInboxInsn, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0xc21ef0fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) pure returns(bytes32)
func (_ProtocolTester *ProtocolTesterCallerSession) GenerateAssertionHash(_afterHash [32]byte, _didInboxInsn bool, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GenerateAssertionHash(&_ProtocolTester.CallOpts, _afterHash, _didInboxInsn, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x9353b9b4.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, bytes32 _beforeInboxHash) pure returns(bytes32)
func (_ProtocolTester *ProtocolTesterCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _beforeInboxHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ProtocolTester.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _beforeInboxHash)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x9353b9b4.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, bytes32 _beforeInboxHash) pure returns(bytes32)
func (_ProtocolTester *ProtocolTesterSession) GeneratePreconditionHash(_beforeHash [32]byte, _beforeInboxHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GeneratePreconditionHash(&_ProtocolTester.CallOpts, _beforeHash, _beforeInboxHash)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x9353b9b4.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, bytes32 _beforeInboxHash) pure returns(bytes32)
func (_ProtocolTester *ProtocolTesterCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _beforeInboxHash [32]byte) ([32]byte, error) {
	return _ProtocolTester.Contract.GeneratePreconditionHash(&_ProtocolTester.CallOpts, _beforeHash, _beforeInboxHash)
}
