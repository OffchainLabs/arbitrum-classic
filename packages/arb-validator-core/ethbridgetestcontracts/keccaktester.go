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

// KeccakTesterABI is the input ABI used to generate the binding from.
const KeccakTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"A\",\"type\":\"uint256[25]\"}],\"name\":\"keccak_f\",\"outputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"\",\"type\":\"uint256[25]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// KeccakTesterFuncSigs maps the 4-byte function signature to its string representation.
var KeccakTesterFuncSigs = map[string]string{
	"d7533595": "keccak_f(uint256[25])",
}

// KeccakTesterBin is the compiled bytecode used for deploying new contracts.
var KeccakTesterBin = "0x610ed1610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063d75335951461003a575b600080fd5b61008e600480360361032081101561005157600080fd5b8101908080610320019060198060200260405190810160405280929190826019602002808284376000920191909152509194506100c79350505050565b604051808261032080838360005b838110156100b457818101518382015260200161009c565b5050505090500191505060405180910390f35b6100cf610e40565b6100d8826100de565b92915050565b6100e6610e40565b6100ee610e5f565b6100f6610e5f565b600080610101610e40565b610109610e7d565b60405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060008090505b6018811015610e3357886004602002015189600360200201518a600260200201518b600160200201518c60006020020151181818188760006020020152886009602002015189600860200201518a600760200201518b600660200201518c6005602002015118181818876001602002015288600e602002015189600d60200201518a600c60200201518b600b60200201518c600a6020020151181818188760026020020152886013602002015189601260200201518a601160200201518b601060200201518c600f6020020151181818188760036020020152886018602002015189601760200201518a601660200201518b601560200201518c601460200201511818181887600460200201526001603f1b87600160200201518161035157fe5b0487600160200201516002026001600160401b03161787600460200201511886600060200201526001603f1b87600260200201518161038c57fe5b0487600260200201516002026001600160401b03161787600060200201511886600160200201526001603f1b8760036020020151816103c757fe5b0487600360200201516002026001600160401b03161787600160200201511886600260200201526001603f1b87600460200201518161040257fe5b0487600460200201516002026001600160401b0316178760026005811061042557fe5b602002015118606087015286516001603f1b900487600060200201516002026001600160401b031617876003602002015118866004602002015285600060200201518960006020020151188960006020020152856000602002015189600160200201511889600160200201528560006020020151896002602002015118896002602002015285600060200201518960036020020151188960036020020152856000602002015189600460200201511889600460200201528560016020020151896005602002015118896005602002015285600160200201518960066020020151188960066020020152856001602002015189600760200201511889600760200201528560016020020151896008602002015118896008602002015285600160200201518960096020020151188960096020020152856002602002015189600a60200201511889600a6020020152856002602002015189600b60200201511889600b6020020152856002602002015189600c60200201511889600c6020020152856002602002015189600d60200201511889600d6020020152856002602002015189600e60200201511889600e6020020152856003602002015189600f60200201511889600f60200201528560036020020151896010602002015118896010602002015285600360200201518960116020020151188960116020020152856003602002015189601260200201511889601260200201528560036020020151896013602002015118896013602002015285600460200201518960146020020151188960146020020152856004602002015189601560200201511889601560200201528560046020020151896016602002015118896016602002015285600460200201518960176020020151188960176020020152856004602002015189601860200201511889601860200201528860006020020151836000602002015263100000008960016020020151816106f457fe5b60208b01519190046410000000009091026001600160401b039081169190911761010085015260408a01516001603d1b8104600890910282161761016085015260608a01516280000081046502000000000090910282161761026085015260808a0151654000000000008104620400009091028216176102c085015260a08a015160028082029092166001603f1b90910417908490602002015260c0890151621000008104651000000000009091026001600160401b039081169190911760a085015260e08a0151664000000000000081046104009091028216176101a08501526101008a0151620800008104652000000000009091028216176102008501526101208a015160048082029092166001603e1b909104176103008501526101408a01516101408b01516001600160401b036001603e1b90910216919004176080840152610160890151670400000000000000906101608b01516001600160401b036040909102169190041760e084015261018089015162200000906101808b01516001600160401b036508000000000090910216919004176101408401526101a08901516602000000000000906101a08b01516001600160401b0361800090910216919004176102408401526101c08901516008906101c08b01516001600160401b036001603d1b90910216919004176102a08401526101e0890151641000000000906101e08b01516001600160401b03631000000090910216919004176020840152610200808a01516102008b01516001600160401b0366800000000000009091021691900417610120840152610220890151648000000000906102208b01516001600160401b036302000000909102169190041761018084015261024089015165080000000000906102408b01516001600160401b036220000090910216919004176101e0840152610260890151610100906102608b01516001600160401b03600160381b90910216919004176102e0840152610280890151642000000000906102808b01516001600160401b036308000000909102169190041760608401526102a089015165100000000000906102a08b01516001600160401b0362100000909102169190041760c08401526102c08901516302000000906102c08b01516001600160401b0364800000000090910216919004176101c08401526102e0890151600160381b906102e08b01516001600160401b036101009091021691900417610220840152610300890151660400000000000090048960186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118896000602002015282600b602002015183600660200201511916836001602002015118896001602002015282600c602002015183600760200201511916836002602002015118896002602002015282600d602002015183600860200201511916836003602002015118896003602002015282600e602002015183600960200201511916836004602002015118896004602002015282600f602002015183600a602002015119168360056020020151188960056020020152826010602002015183600b602002015119168360066020020151188960066020020152826011602002015183600c602002015119168360076020020151188960076020020152826012602002015183600d602002015119168360086020020151188960086020020152826013602002015183600e602002015119168360096020020151188960096020020152826014602002015183600f6020020151191683600a60200201511889600a602002015282601560200201518360106020020151191683600b60200201511889600b602002015282601660200201518360116020020151191683600c60200201511889600c602002015282601760200201518360126020020151191683600d60200201511889600d602002015282601860200201518360136020020151191683600e60200201511889600e602002015282600060200201518360146020020151191683600f60200201511889600f6020020152826001602002015183601560200201511916836010602002015118896010602002015282600260200201518360166020020151191683601160200201511889601160200201528260036020020151836017602002015119168360126020020151188960126020020152826004602002015183601860200201511916836013602002015118896013602002015282600560200201518360006020020151191683601460200201511889601460200201528260066020020151836001602002015119168360156020020151188960156020020152826007602002015183600260200201511916836016602002015118896016602002015282600860200201518360036020020151191683601760200201511889601760200201528260096020020151836004602002015119168360186020020151188960186020020152818160188110610e2157fe5b60200201518951188952600101610230565b5096979650505050505050565b6040518061032001604052806019906020820280388339509192915050565b6040518060a001604052806005906020820280388339509192915050565b604051806103000160405280601890602082028038833950919291505056fea265627a7a72315820a3af6ce233acd6b934841c9e7f73acad2b8b9e7915943f9972641a83dee991a364736f6c63430005110032"

// DeployKeccakTester deploys a new Ethereum contract, binding an instance of KeccakTester to it.
func DeployKeccakTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeccakTester, error) {
	parsed, err := abi.JSON(strings.NewReader(KeccakTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KeccakTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeccakTester{KeccakTesterCaller: KeccakTesterCaller{contract: contract}, KeccakTesterTransactor: KeccakTesterTransactor{contract: contract}, KeccakTesterFilterer: KeccakTesterFilterer{contract: contract}}, nil
}

// KeccakTester is an auto generated Go binding around an Ethereum contract.
type KeccakTester struct {
	KeccakTesterCaller     // Read-only binding to the contract
	KeccakTesterTransactor // Write-only binding to the contract
	KeccakTesterFilterer   // Log filterer for contract events
}

// KeccakTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeccakTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeccakTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeccakTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeccakTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeccakTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeccakTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeccakTesterSession struct {
	Contract     *KeccakTester     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeccakTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeccakTesterCallerSession struct {
	Contract *KeccakTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KeccakTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeccakTesterTransactorSession struct {
	Contract     *KeccakTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KeccakTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeccakTesterRaw struct {
	Contract *KeccakTester // Generic contract binding to access the raw methods on
}

// KeccakTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeccakTesterCallerRaw struct {
	Contract *KeccakTesterCaller // Generic read-only contract binding to access the raw methods on
}

// KeccakTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeccakTesterTransactorRaw struct {
	Contract *KeccakTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeccakTester creates a new instance of KeccakTester, bound to a specific deployed contract.
func NewKeccakTester(address common.Address, backend bind.ContractBackend) (*KeccakTester, error) {
	contract, err := bindKeccakTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeccakTester{KeccakTesterCaller: KeccakTesterCaller{contract: contract}, KeccakTesterTransactor: KeccakTesterTransactor{contract: contract}, KeccakTesterFilterer: KeccakTesterFilterer{contract: contract}}, nil
}

// NewKeccakTesterCaller creates a new read-only instance of KeccakTester, bound to a specific deployed contract.
func NewKeccakTesterCaller(address common.Address, caller bind.ContractCaller) (*KeccakTesterCaller, error) {
	contract, err := bindKeccakTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeccakTesterCaller{contract: contract}, nil
}

// NewKeccakTesterTransactor creates a new write-only instance of KeccakTester, bound to a specific deployed contract.
func NewKeccakTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*KeccakTesterTransactor, error) {
	contract, err := bindKeccakTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeccakTesterTransactor{contract: contract}, nil
}

// NewKeccakTesterFilterer creates a new log filterer instance of KeccakTester, bound to a specific deployed contract.
func NewKeccakTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*KeccakTesterFilterer, error) {
	contract, err := bindKeccakTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeccakTesterFilterer{contract: contract}, nil
}

// bindKeccakTester binds a generic wrapper to an already deployed contract.
func bindKeccakTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeccakTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeccakTester *KeccakTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeccakTester.Contract.KeccakTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeccakTester *KeccakTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeccakTester.Contract.KeccakTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeccakTester *KeccakTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeccakTester.Contract.KeccakTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeccakTester *KeccakTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeccakTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeccakTester *KeccakTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeccakTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeccakTester *KeccakTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeccakTester.Contract.contract.Transact(opts, method, params...)
}

// KeccakF is a free data retrieval call binding the contract method 0xd7533595.
//
// Solidity: function keccak_f(uint256[25] A) pure returns(uint256[25])
func (_KeccakTester *KeccakTesterCaller) KeccakF(opts *bind.CallOpts, A [25]*big.Int) ([25]*big.Int, error) {
	var (
		ret0 = new([25]*big.Int)
	)
	out := ret0
	err := _KeccakTester.contract.Call(opts, out, "keccak_f", A)
	return *ret0, err
}

// KeccakF is a free data retrieval call binding the contract method 0xd7533595.
//
// Solidity: function keccak_f(uint256[25] A) pure returns(uint256[25])
func (_KeccakTester *KeccakTesterSession) KeccakF(A [25]*big.Int) ([25]*big.Int, error) {
	return _KeccakTester.Contract.KeccakF(&_KeccakTester.CallOpts, A)
}

// KeccakF is a free data retrieval call binding the contract method 0xd7533595.
//
// Solidity: function keccak_f(uint256[25] A) pure returns(uint256[25])
func (_KeccakTester *KeccakTesterCallerSession) KeccakF(A [25]*big.Int) ([25]*big.Int, error) {
	return _KeccakTester.Contract.KeccakF(&_KeccakTester.CallOpts, A)
}
