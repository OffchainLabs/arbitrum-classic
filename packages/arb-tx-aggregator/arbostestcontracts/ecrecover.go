// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostestcontracts

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

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[]"

// ECDSABin is the compiled bytecode used for deploying new contracts.
var ECDSABin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201a42dda541a13dcdfde6ca2062a26afabf562fd137853297be56a8f148e3cc6064736f6c63430005110032"

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// ECRecoverTestABI is the input ABI used to generate the binding from.
const ECRecoverTestABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"getByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"parseSig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ECRecoverTestFuncSigs maps the 4-byte function signature to its string representation.
var ECRecoverTestFuncSigs = map[string]string{
	"0f71e434": "getByte(uint256,bytes32)",
	"f42e5f25": "parseSig(bytes)",
	"2e295ec9": "recoverSigner(bytes,bytes)",
}

// ECRecoverTestBin is the compiled bytecode used for deploying new contracts.
var ECRecoverTestBin = "0x608060405234801561001057600080fd5b50610583806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630f71e434146100465780632e295ec91461007b578063f42e5f25146101c4575b600080fd5b6100696004803603604081101561005c57600080fd5b508035906020013561028b565b60408051918252519081900360200190f35b6101a86004803603604081101561009157600080fd5b8101906020810181356401000000008111156100ac57600080fd5b8201836020820111156100be57600080fd5b803590602001918460018302840111640100000000831117156100e057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561013357600080fd5b82018360208201111561014557600080fd5b8035906020019184600183028401116401000000008311171561016757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610290945050505050565b604080516001600160a01b039092168252519081900360200190f35b61026a600480360360208110156101da57600080fd5b8101906020810181356401000000008111156101f557600080fd5b82018360208201111561020757600080fd5b8035906020019184600183028401116401000000008311171561022957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102b9945050505050565b60408051938452602084019290925260ff1682820152519081900360600190f35b901a90565b81516020830120600090816102a4826102d2565b90506102b08185610323565b95945050505050565b60208101516040820151606083015160001a9193909250565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b6000815160411461037b576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156103ec5760405162461bcd60e51b815260040180806020018281038252602281526020018061050b6022913960400191505060405180910390fd5b8060ff16601b1415801561040457508060ff16601c14155b156104405760405162461bcd60e51b815260040180806020018281038252602281526020018061052d6022913960400191505060405180910390fd5b60408051600080825260208083018085528a905260ff85168385015260608301879052608083018690529251909260019260a080820193601f1981019281900390910190855afa158015610498573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610500576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b969550505050505056fe45434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c7565a265627a7a72315820996825ca912c37c2486d2f7088d270ba002e10666414cf37884379e337d4682564736f6c63430005110032"

// DeployECRecoverTest deploys a new Ethereum contract, binding an instance of ECRecoverTest to it.
func DeployECRecoverTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecoverTest, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoverTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoverTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecoverTest{ECRecoverTestCaller: ECRecoverTestCaller{contract: contract}, ECRecoverTestTransactor: ECRecoverTestTransactor{contract: contract}, ECRecoverTestFilterer: ECRecoverTestFilterer{contract: contract}}, nil
}

// ECRecoverTest is an auto generated Go binding around an Ethereum contract.
type ECRecoverTest struct {
	ECRecoverTestCaller     // Read-only binding to the contract
	ECRecoverTestTransactor // Write-only binding to the contract
	ECRecoverTestFilterer   // Log filterer for contract events
}

// ECRecoverTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoverTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoverTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoverTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverTestSession struct {
	Contract     *ECRecoverTest    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoverTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoverTestCallerSession struct {
	Contract *ECRecoverTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ECRecoverTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoverTestTransactorSession struct {
	Contract     *ECRecoverTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ECRecoverTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoverTestRaw struct {
	Contract *ECRecoverTest // Generic contract binding to access the raw methods on
}

// ECRecoverTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoverTestCallerRaw struct {
	Contract *ECRecoverTestCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoverTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoverTestTransactorRaw struct {
	Contract *ECRecoverTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecoverTest creates a new instance of ECRecoverTest, bound to a specific deployed contract.
func NewECRecoverTest(address common.Address, backend bind.ContractBackend) (*ECRecoverTest, error) {
	contract, err := bindECRecoverTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecoverTest{ECRecoverTestCaller: ECRecoverTestCaller{contract: contract}, ECRecoverTestTransactor: ECRecoverTestTransactor{contract: contract}, ECRecoverTestFilterer: ECRecoverTestFilterer{contract: contract}}, nil
}

// NewECRecoverTestCaller creates a new read-only instance of ECRecoverTest, bound to a specific deployed contract.
func NewECRecoverTestCaller(address common.Address, caller bind.ContractCaller) (*ECRecoverTestCaller, error) {
	contract, err := bindECRecoverTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoverTestCaller{contract: contract}, nil
}

// NewECRecoverTestTransactor creates a new write-only instance of ECRecoverTest, bound to a specific deployed contract.
func NewECRecoverTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoverTestTransactor, error) {
	contract, err := bindECRecoverTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoverTestTransactor{contract: contract}, nil
}

// NewECRecoverTestFilterer creates a new log filterer instance of ECRecoverTest, bound to a specific deployed contract.
func NewECRecoverTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoverTestFilterer, error) {
	contract, err := bindECRecoverTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoverTestFilterer{contract: contract}, nil
}

// bindECRecoverTest binds a generic wrapper to an already deployed contract.
func bindECRecoverTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoverTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecoverTest *ECRecoverTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecoverTest.Contract.ECRecoverTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecoverTest *ECRecoverTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecoverTest.Contract.ECRecoverTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecoverTest *ECRecoverTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecoverTest.Contract.ECRecoverTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecoverTest *ECRecoverTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecoverTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecoverTest *ECRecoverTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecoverTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecoverTest *ECRecoverTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecoverTest.Contract.contract.Transact(opts, method, params...)
}

// GetByte is a free data retrieval call binding the contract method 0x0f71e434.
//
// Solidity: function getByte(uint256 index, bytes32 val) pure returns(uint256)
func (_ECRecoverTest *ECRecoverTestCaller) GetByte(opts *bind.CallOpts, index *big.Int, val [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ECRecoverTest.contract.Call(opts, out, "getByte", index, val)
	return *ret0, err
}

// GetByte is a free data retrieval call binding the contract method 0x0f71e434.
//
// Solidity: function getByte(uint256 index, bytes32 val) pure returns(uint256)
func (_ECRecoverTest *ECRecoverTestSession) GetByte(index *big.Int, val [32]byte) (*big.Int, error) {
	return _ECRecoverTest.Contract.GetByte(&_ECRecoverTest.CallOpts, index, val)
}

// GetByte is a free data retrieval call binding the contract method 0x0f71e434.
//
// Solidity: function getByte(uint256 index, bytes32 val) pure returns(uint256)
func (_ECRecoverTest *ECRecoverTestCallerSession) GetByte(index *big.Int, val [32]byte) (*big.Int, error) {
	return _ECRecoverTest.Contract.GetByte(&_ECRecoverTest.CallOpts, index, val)
}

// ParseSig is a free data retrieval call binding the contract method 0xf42e5f25.
//
// Solidity: function parseSig(bytes signature) pure returns(bytes32, bytes32, uint8)
func (_ECRecoverTest *ECRecoverTestCaller) ParseSig(opts *bind.CallOpts, signature []byte) ([32]byte, [32]byte, uint8, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ECRecoverTest.contract.Call(opts, out, "parseSig", signature)
	return *ret0, *ret1, *ret2, err
}

// ParseSig is a free data retrieval call binding the contract method 0xf42e5f25.
//
// Solidity: function parseSig(bytes signature) pure returns(bytes32, bytes32, uint8)
func (_ECRecoverTest *ECRecoverTestSession) ParseSig(signature []byte) ([32]byte, [32]byte, uint8, error) {
	return _ECRecoverTest.Contract.ParseSig(&_ECRecoverTest.CallOpts, signature)
}

// ParseSig is a free data retrieval call binding the contract method 0xf42e5f25.
//
// Solidity: function parseSig(bytes signature) pure returns(bytes32, bytes32, uint8)
func (_ECRecoverTest *ECRecoverTestCallerSession) ParseSig(signature []byte) ([32]byte, [32]byte, uint8, error) {
	return _ECRecoverTest.Contract.ParseSig(&_ECRecoverTest.CallOpts, signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x2e295ec9.
//
// Solidity: function recoverSigner(bytes message, bytes signature) pure returns(address)
func (_ECRecoverTest *ECRecoverTestCaller) RecoverSigner(opts *bind.CallOpts, message []byte, signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ECRecoverTest.contract.Call(opts, out, "recoverSigner", message, signature)
	return *ret0, err
}

// RecoverSigner is a free data retrieval call binding the contract method 0x2e295ec9.
//
// Solidity: function recoverSigner(bytes message, bytes signature) pure returns(address)
func (_ECRecoverTest *ECRecoverTestSession) RecoverSigner(message []byte, signature []byte) (common.Address, error) {
	return _ECRecoverTest.Contract.RecoverSigner(&_ECRecoverTest.CallOpts, message, signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x2e295ec9.
//
// Solidity: function recoverSigner(bytes message, bytes signature) pure returns(address)
func (_ECRecoverTest *ECRecoverTestCallerSession) RecoverSigner(message []byte, signature []byte) (common.Address, error) {
	return _ECRecoverTest.Contract.RecoverSigner(&_ECRecoverTest.CallOpts, message, signature)
}
