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
var KeccakTesterBin = "0x610d60610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063d75335951461003a575b600080fd5b61008e600480360361032081101561005157600080fd5b8101908080610320019060198060200260405190810160405280929190826019602002808284376000920191909152509194506100c79350505050565b604051808261032080838360005b838110156100b457818101518382015260200161009c565b5050505090500191505060405180910390f35b6100cf610ccf565b6100d8826100de565b92915050565b6100e6610ccf565b6100ee610cee565b6100f6610cee565b6100fe610ccf565b610106610d0c565b60405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060008090505b6018811015610cc4576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b85600360200201518161035257fe5b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b85600460200201518161038d57fe5b0485600460200201516002026001600160401b031617856002600581106103b057fe5b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b0316178360026019811061061a57fe5b602002015260c0870151621000008104651000000000009091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b90910216919004176080840152610160870151670400000000000000906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036508000000000090910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b036302000000909102169190041761018084015261024087015165080000000000906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a087015165100000000000906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f6020020152826001602002015183601560200201511916836010602002015118876010602002015282600260200201518360166020020151191683601160200201511887601160200201528260036020020151836017602002015119168360126020020151188760126020020152826004602002015183601860200201511916836013602002015118876013602002015282600560200201518360006020020151191683601460200201511887601460200201528260066020020151836001602002015119168360156020020151188760156020020152826007602002015183600260200201511916836016602002015118876016602002015282600860200201518360036020020151191683601760200201511887601760200201528260096020020151836004602002015119168360186020020151188760186020020152818160188110610cb257fe5b6020020151875118875260010161022d565b509495945050505050565b6040518061032001604052806019906020820280388339509192915050565b6040518060a001604052806005906020820280388339509192915050565b604051806103000160405280601890602082028038833950919291505056fea265627a7a7231582018a9b6b5e872bb4e9030d8cf30a1ffef2c07abb87403012e95c65d1cfd6d241564736f6c634300050f0032"

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
