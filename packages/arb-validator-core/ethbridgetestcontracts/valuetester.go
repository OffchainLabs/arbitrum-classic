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

// ValueTesterABI is the input ABI used to generate the binding from.
const ValueTesterABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hashTestTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"innerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueSize\",\"type\":\"uint256\"}],\"name\":\"hashTuplePreImage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueTesterFuncSigs maps the 4-byte function signature to its string representation.
var ValueTesterFuncSigs = map[string]string{
	"98206792": "deserializeHash(bytes,uint256)",
	"fd5d0c8b": "hashTestTuple()",
	"c6d25c8e": "hashTuplePreImage(bytes32,uint256)",
}

// ValueTesterBin is the compiled bytecode used for deploying new contracts.
var ValueTesterBin = "0x608060405234801561001057600080fd5b5061110d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80639820679214610046578063c6d25c8e14610107578063fd5d0c8b1461013c575b600080fd5b6100ee6004803603604081101561005c57600080fd5b81019060208101813564010000000081111561007757600080fd5b82018360208201111561008957600080fd5b803590602001918460018302840111640100000000831117156100ab57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610144915050565b6040805192835260208301919091528051918290030190f35b61012a6004803603604081101561011d57600080fd5b5080359060200135610177565b60408051918252519081900360200190f35b61012a61018a565b600080600061015161107b565b61015b8686610257565b915091508161016982610412565b9350935050505b9250929050565b60006101838383610584565b9392505050565b60408051600280825260608281019093526000929190816020015b6101ad61107b565b8152602001906001900390816101a55790505090506101cc606f6105d2565b816000815181106101d957fe5b6020908102919091010152610228600060405190808252806020026020018201604052801561022257816020015b61020f61107b565b8152602001906001900390816102075790505b50610692565b8160018151811061023557fe5b602002602001018190525061025161024c82610692565b610412565b91505090565b600061026161107b565b835183106102a7576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806102b486866107d4565b915091506102c06107fb565b60ff168160ff1614156102f45760006102d98784610800565b9093509050826102e8826105d2565b94509450505050610170565b6102fc61086e565b60ff168160ff16141561031e576103138683610873565b935093505050610170565b610326610915565b60ff168160ff16141561034e57600061033f8784610800565b9093509050826102e88261091a565b610356610a07565b60ff168160ff16141561036d576103138683610a0c565b610375610a94565b60ff168160ff1610158015610396575061038d610a99565b60ff168160ff16105b156103d25760006103a5610a94565b8203905060606103b6828986610a9e565b9094509050836103c582610692565b9550955050505050610170565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b600061041c6107fb565b60ff16826080015160ff16141561043f57815161043890610b47565b905061057f565b61044761086e565b60ff16826080015160ff161415610465576104388260200151610b6b565b61046d610a07565b60ff16826080015160ff16141561048f57815160a08301516104389190610584565b610497610a94565b60ff16826080015160ff1614156104d0576104b061107b565b6104bd8360400151610c60565b90506104c881610412565b91505061057f565b6104d8610dd6565b60ff16826080015160ff1614156104f15750805161057f565b6104f9610915565b60ff16826080015160ff16141561053e575060608082015160408051607b6020808301919091528183019390935281518082038301815293019052815191012061057f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b919050565b600061058e610a94565b8383604051602001808460ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b6105da61107b565b6040805160c0810182528381528151606081018352600080825260208083018290528451828152808201865293949085019390830191610630565b61061d61107b565b8152602001906001900390816106155790505b5090528152602001600060405190808252806020026020018201604052801561067357816020015b61066061107b565b8152602001906001900390816106585790505b5081526000602082018190526040820152600160609091015292915050565b61069a61107b565b6106a48251610ddb565b6106f5576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561072c5783818151811061070f57fe5b602002602001015160a001518201915080806001019150506106fa565b506040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff8111801561077557600080fd5b506040519080825280602002602001820160405280156107af57816020015b61079c61107b565b8152602001906001900390816107945790505b5090528152602081019490945260006040850152600360608501526080909301525090565b600080826001018484815181106107e757fe5b016020015190925060f81c90509250929050565b600090565b60008082845110158015610818575060208385510310155b610855576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b602083016108638585610de2565b915091509250929050565b600190565b600061087d61107b565b8260008061088961107b565b600061089589866107d4565b90955093506108a489866107d4565b9095509250600160ff851614156108c5576108bf8986610257565b90955091505b6108cf8986610e3b565b9095509050600160ff851614156108fa57846108ec848385610e4c565b965096505050505050610170565b846109058483610ed0565b9650965050505050509250929050565b600c90565b61092261107b565b6040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff8111801561096a57600080fd5b506040519080825280602002602001820160405280156109a457816020015b61099161107b565b8152602001906001900390816109895790505b509052815260200160006040519080825280602002602001820160405280156109e757816020015b6109d461107b565b8152602001906001900390816109cc5790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b6000610a1661107b565b82845110158015610a2b575060408385510310155b610a68576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b600080610a758686610e3b565b9094509150610a848685610800565b9094509050836101698383610f2b565b600390565b600d90565b60006060828160ff871667ffffffffffffffff81118015610abe57600080fd5b50604051908082528060200260200182016040528015610af857816020015b610ae561107b565b815260200190600190039081610add5790505b50905060005b8760ff168160ff161015610b3a57610b168784610257565b838360ff1681518110610b2557fe5b60209081029190910101529250600101610afe565b5090969095509350505050565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610b7c57fe5b604082015151610bdf57610b8e61086e565b82600001518360200151604051602001808460ff1660f81b81526001018360ff1660f81b8152600101828152602001935050505060405160208183030381529060405280519060200120905061057f565b610be761086e565b8260000151610c0d8460400151600081518110610c0057fe5b6020026020010151610412565b8460200151604051602001808560ff1660f81b81526001018460ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b610c6861107b565b600882511115610cb6576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825167ffffffffffffffff81118015610cd057600080fd5b50604051908082528060200260200182016040528015610cfa578160200160208202803683370190505b508051909150600160005b82811015610d5d57610d1c868281518110610c0057fe5b848281518110610d2857fe5b602002602001018181525050858181518110610d4057fe5b602002602001015160a00151820191508080600101915050610d05565b506000835184604051602001808360ff1660f81b8152600101828051906020019060200280838360005b83811015610d9f578181015183820152602001610d87565b5050505090500192505050604051602081830303815290604052805190602001209050610dcc8183610f2b565b9695505050505050565b606490565b6008101590565b60008160200183511015610e32576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b600080602083016108638585610de2565b610e5461107b565b604080516001808252818301909252606091816020015b610e7361107b565b815260200190600190039081610e6b5790505090508281600081518110610e9657fe5b6020026020010181905250610ec760405180606001604052808760ff16815260200186815260200183815250610fea565b95945050505050565b610ed861107b565b6040805160608101825260ff85168152602080820185905282516000808252918101845261018393830191610f23565b610f1061107b565b815260200190600190039081610f085790505b509052610fea565b610f3361107b565b6040805160c0810182528481528151606081018352600080825260208083018290528451828152808201865293949085019390830191610f89565b610f7661107b565b815260200190600190039081610f6e5790505b50905281526020016000604051908082528060200260200182016040528015610fcc57816020015b610fb961107b565b815260200190600190039081610fb15790505b50815260006020820152600260408201526060019290925250919050565b610ff261107b565b6040518060c0016040528060008152602001838152602001600067ffffffffffffffff8111801561102257600080fd5b5060405190808252806020026020018201604052801561105c57816020015b61104961107b565b8152602001906001900390816110415790505b5081526000602082015260016040820181905260609091015292915050565b6040518060c00160405280600081526020016110956110b8565b815260606020820181905260006040830181905290820181905260809091015290565b604080516060808201835260008083526020830152918101919091529056fea2646970667358221220b22d37cd0ab39bf0eb03ad2beb71763a790ebe2ce8c063a21e9a7c1ec597a98a64736f6c634300060c0033"

// DeployValueTester deploys a new Ethereum contract, binding an instance of ValueTester to it.
func DeployValueTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueTester, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueTester{ValueTesterCaller: ValueTesterCaller{contract: contract}, ValueTesterTransactor: ValueTesterTransactor{contract: contract}, ValueTesterFilterer: ValueTesterFilterer{contract: contract}}, nil
}

// ValueTester is an auto generated Go binding around an Ethereum contract.
type ValueTester struct {
	ValueTesterCaller     // Read-only binding to the contract
	ValueTesterTransactor // Write-only binding to the contract
	ValueTesterFilterer   // Log filterer for contract events
}

// ValueTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueTesterSession struct {
	Contract     *ValueTester      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueTesterCallerSession struct {
	Contract *ValueTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ValueTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTesterTransactorSession struct {
	Contract     *ValueTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ValueTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueTesterRaw struct {
	Contract *ValueTester // Generic contract binding to access the raw methods on
}

// ValueTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueTesterCallerRaw struct {
	Contract *ValueTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTesterTransactorRaw struct {
	Contract *ValueTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueTester creates a new instance of ValueTester, bound to a specific deployed contract.
func NewValueTester(address common.Address, backend bind.ContractBackend) (*ValueTester, error) {
	contract, err := bindValueTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueTester{ValueTesterCaller: ValueTesterCaller{contract: contract}, ValueTesterTransactor: ValueTesterTransactor{contract: contract}, ValueTesterFilterer: ValueTesterFilterer{contract: contract}}, nil
}

// NewValueTesterCaller creates a new read-only instance of ValueTester, bound to a specific deployed contract.
func NewValueTesterCaller(address common.Address, caller bind.ContractCaller) (*ValueTesterCaller, error) {
	contract, err := bindValueTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTesterCaller{contract: contract}, nil
}

// NewValueTesterTransactor creates a new write-only instance of ValueTester, bound to a specific deployed contract.
func NewValueTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTesterTransactor, error) {
	contract, err := bindValueTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTesterTransactor{contract: contract}, nil
}

// NewValueTesterFilterer creates a new log filterer instance of ValueTester, bound to a specific deployed contract.
func NewValueTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueTesterFilterer, error) {
	contract, err := bindValueTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueTesterFilterer{contract: contract}, nil
}

// bindValueTester binds a generic wrapper to an already deployed contract.
func bindValueTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueTester *ValueTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueTester.Contract.ValueTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueTester *ValueTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueTester.Contract.ValueTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueTester *ValueTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueTester.Contract.ValueTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueTester *ValueTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueTester *ValueTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueTester *ValueTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueTester.Contract.contract.Transact(opts, method, params...)
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(uint256, bytes32)
func (_ValueTester *ValueTesterCaller) DeserializeHash(opts *bind.CallOpts, data []byte, startOffset *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _ValueTester.contract.Call(opts, &out, "deserializeHash", data, startOffset)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(uint256, bytes32)
func (_ValueTester *ValueTesterSession) DeserializeHash(data []byte, startOffset *big.Int) (*big.Int, [32]byte, error) {
	return _ValueTester.Contract.DeserializeHash(&_ValueTester.CallOpts, data, startOffset)
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(uint256, bytes32)
func (_ValueTester *ValueTesterCallerSession) DeserializeHash(data []byte, startOffset *big.Int) (*big.Int, [32]byte, error) {
	return _ValueTester.Contract.DeserializeHash(&_ValueTester.CallOpts, data, startOffset)
}

// HashTestTuple is a free data retrieval call binding the contract method 0xfd5d0c8b.
//
// Solidity: function hashTestTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) HashTestTuple(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValueTester.contract.Call(opts, &out, "hashTestTuple")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTestTuple is a free data retrieval call binding the contract method 0xfd5d0c8b.
//
// Solidity: function hashTestTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterSession) HashTestTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashTestTuple(&_ValueTester.CallOpts)
}

// HashTestTuple is a free data retrieval call binding the contract method 0xfd5d0c8b.
//
// Solidity: function hashTestTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) HashTestTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashTestTuple(&_ValueTester.CallOpts)
}

// HashTuplePreImage is a free data retrieval call binding the contract method 0xc6d25c8e.
//
// Solidity: function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) HashTuplePreImage(opts *bind.CallOpts, innerHash [32]byte, valueSize *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ValueTester.contract.Call(opts, &out, "hashTuplePreImage", innerHash, valueSize)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTuplePreImage is a free data retrieval call binding the contract method 0xc6d25c8e.
//
// Solidity: function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) pure returns(bytes32)
func (_ValueTester *ValueTesterSession) HashTuplePreImage(innerHash [32]byte, valueSize *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.HashTuplePreImage(&_ValueTester.CallOpts, innerHash, valueSize)
}

// HashTuplePreImage is a free data retrieval call binding the contract method 0xc6d25c8e.
//
// Solidity: function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) HashTuplePreImage(innerHash [32]byte, valueSize *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.HashTuplePreImage(&_ValueTester.CallOpts, innerHash, valueSize)
}
