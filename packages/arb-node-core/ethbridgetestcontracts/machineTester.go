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

// MachineTesterABI is the input ABI used to generate the binding from.
const MachineTesterABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data2\",\"type\":\"bytes\"}],\"name\":\"addStackVal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeMachine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MachineTesterFuncSigs maps the 4-byte function signature to its string representation.
var MachineTesterFuncSigs = map[string]string{
	"5f098d7f": "addStackVal(bytes,bytes)",
	"5270f3e9": "deserializeMachine(bytes)",
}

// MachineTesterBin is the compiled bytecode used for deploying new contracts.
var MachineTesterBin = "0x608060405234801561001057600080fd5b506113ef806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635270f3e91461003b5780635f098d7f146100fa575b600080fd5b6100e16004803603602081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610239945050505050565b6040805192835260208301919091528051918290030190f35b6102276004803603604081101561011057600080fd5b81019060208101813564010000000081111561012b57600080fd5b82018360208201111561013d57600080fd5b8035906020019184600183028401116401000000008311171561015f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101b257600080fd5b8201836020820111156101c457600080fd5b803590602001918460018302840111640100000000831117156101e657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061026b945050505050565b60408051918252519081900360200190f35b60008060006102466112f2565b6102518560006102ba565b90925090508161026082610370565b935093505050915091565b60008061027661135d565b61027e61135d565b61028986600061044b565b909350915061029985600061044b565b90935090506102b06102ab8383610606565b610684565b9695505050505050565b60006102c46112f2565b6102cc6112f2565b60006101008201819052806102e187876107f1565b90965091506102f0878761085f565b60208501529550610301878761085f565b60408501529550610312878761044b565b60608501529550610323878761044b565b6080850152955061033487876107f1565b60a0850152955061034587876107f1565b9096509050610354878761044b565b60e085015291835260c0830152935083925090505b9250929050565b60006002826101000151141561038857506000610446565b6001826101000151141561039e57506001610446565b815160208301516103ae90610684565b6103bb8460400151610684565b6103c88560600151610684565b6103d58660800151610684565b8660a001518760c001516103ec8960e00151610684565b60405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012090505b919050565b600061045561135d565b8351831061049b576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806104a886866108f4565b915091506104b461091b565b60ff168160ff1614156104e85760006104cd87846107f1565b9093509050826104dc82610920565b94509450505050610369565b6104f06109e0565b60ff168160ff1614156105125761050786836109e5565b935093505050610369565b61051a610a87565b60ff168160ff16141561054257600061053387846107f1565b9093509050826104dc82610a8c565b61054a610b79565b60ff168160ff16141561056157610507868361085f565b610569610b7e565b60ff168160ff161015801561058a5750610581610b83565b60ff168160ff16105b156105c6576000610599610b7e565b8203905060606105aa828986610b88565b9094509050836105b982610c31565b9550955050505050610369565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b61060e61135d565b6040805160028082526060828101909352816020015b61062c61135d565b815260200190600190039081610624579050509050828160008151811061064f57fe5b6020026020010181905250838160018151811061066857fe5b602002602001018190525061067c81610d73565b949350505050565b600061068e61091b565b60ff16826080015160ff1614156106b15781516106aa90610eec565b9050610446565b6106b96109e0565b60ff16826080015160ff1614156106d7576106aa8260200151610f10565b6106df610b79565b60ff16826080015160ff16141561070157815160a08301516106aa9190610ff8565b610709610b7e565b60ff16826080015160ff1614156107425761072261135d565b61072f8360400151610d73565b905061073a81610684565b915050610446565b61074a611046565b60ff16826080015160ff16141561076357508051610446565b61076b610a87565b60ff16826080015160ff1614156107b0575060608082015160408051607b60208083019190915281830193909352815180820383018152930190528151910120610446565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b60008082845110158015610809575060208385510310155b610846576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301610854858561104b565b915091509250929050565b600061086961135d565b8284511015801561087e575060408385510310155b6108bb576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6000806108c886866110a4565b90945091506108d786856107f1565b9094509050836108e783836110b5565b9350935050509250929050565b6000808260010184848151811061090757fe5b016020015190925060f81c90509250929050565b600090565b61092861135d565b6040805160c081018252838152815160608101835260008082526020808301829052845182815280820186529394908501939083019161097e565b61096b61135d565b8152602001906001900390816109635790505b509052815260200160006040519080825280602002602001820160405280156109c157816020015b6109ae61135d565b8152602001906001900390816109a65790505b5081526000602082018190526040820152600160609091015292915050565b600190565b60006109ef61135d565b826000806109fb61135d565b6000610a0789866108f4565b9095509350610a1689866108f4565b9095509250600160ff85161415610a3757610a31898661044b565b90955091505b610a4189866110a4565b9095509050600160ff85161415610a6c5784610a5e848385611174565b965096505050505050610369565b84610a7784836111f8565b9650965050505050509250929050565b600c90565b610a9461135d565b6040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610adc57600080fd5b50604051908082528060200260200182016040528015610b1657816020015b610b0361135d565b815260200190600190039081610afb5790505b50905281526020016000604051908082528060200260200182016040528015610b5957816020015b610b4661135d565b815260200190600190039081610b3e5790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b600390565b600d90565b60006060828160ff871667ffffffffffffffff81118015610ba857600080fd5b50604051908082528060200260200182016040528015610be257816020015b610bcf61135d565b815260200190600190039081610bc75790505b50905060005b8760ff168160ff161015610c2457610c00878461044b565b838360ff1681518110610c0f57fe5b60209081029190910101529250600101610be8565b5090969095509350505050565b610c3961135d565b610c43825161125a565b610c94576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610ccb57838181518110610cae57fe5b602002602001015160a00151820191508080600101915050610c99565b506040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610d1457600080fd5b50604051908082528060200260200182016040528015610d4e57816020015b610d3b61135d565b815260200190600190039081610d335790505b5090528152602081019490945260006040850152600360608501526080909301525090565b610d7b61135d565b600882511115610dc9576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825167ffffffffffffffff81118015610de357600080fd5b50604051908082528060200260200182016040528015610e0d578160200160208202803683370190505b508051909150600160005b82811015610e7d57610e3c868281518110610e2f57fe5b6020026020010151610684565b848281518110610e4857fe5b602002602001018181525050858181518110610e6057fe5b602002602001015160a00151820191508080600101915050610e18565b506000835184604051602001808360ff1660f81b8152600101828051906020019060200280838360005b83811015610ebf578181015183820152602001610ea7565b50505050905001925050506040516020818303038152906040528051906020012090506102b081836110b5565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610f2157fe5b604082015151610f8457610f336109e0565b82600001518360200151604051602001808460ff1660f81b81526001018360ff1660f81b81526001018281526020019350505050604051602081830303815290604052805190602001209050610446565b610f8c6109e0565b8260000151610fa58460400151600081518110610e2f57fe5b8460200151604051602001808560ff1660f81b81526001018460ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b6000611002610b7e565b8383604051602001808460ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b606490565b6000816020018351101561109b576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008060208301610854858561104b565b6110bd61135d565b6040805160c0810182528481528151606081018352600080825260208083018290528451828152808201865293949085019390830191611113565b61110061135d565b8152602001906001900390816110f85790505b5090528152602001600060405190808252806020026020018201604052801561115657816020015b61114361135d565b81526020019060019003908161113b5790505b50815260006020820152600260408201526060019290925250919050565b61117c61135d565b604080516001808252818301909252606091816020015b61119b61135d565b81526020019060019003908161119357905050905082816000815181106111be57fe5b60200260200101819052506111ef60405180606001604052808760ff16815260200186815260200183815250611261565b95945050505050565b61120061135d565b6040805160608101825260ff8516815260208082018590528251600080825291810184526112539383019161124b565b61123861135d565b8152602001906001900390816112305790505b509052611261565b9392505050565b6008101590565b61126961135d565b6040518060c0016040528060008152602001838152602001600067ffffffffffffffff8111801561129957600080fd5b506040519080825280602002602001820160405280156112d357816020015b6112c061135d565b8152602001906001900390816112b85790505b5081526000602082015260016040820181905260609091015292915050565b604080516101208101909152600081526020810161130e61135d565b815260200161131b61135d565b815260200161132861135d565b815260200161133561135d565b8152600060208201819052604082015260600161135061135d565b8152602001600081525090565b6040518060c001604052806000815260200161137761139a565b815260606020820181905260006040830181905290820181905260809091015290565b604080516060808201835260008083526020830152918101919091529056fea264697066735822122085c548d1362d901107f85662343e616362098f4424cc30c97292bf807a057d3564736f6c634300060c0033"

// DeployMachineTester deploys a new Ethereum contract, binding an instance of MachineTester to it.
func DeployMachineTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MachineTester, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MachineTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MachineTester{MachineTesterCaller: MachineTesterCaller{contract: contract}, MachineTesterTransactor: MachineTesterTransactor{contract: contract}, MachineTesterFilterer: MachineTesterFilterer{contract: contract}}, nil
}

// MachineTester is an auto generated Go binding around an Ethereum contract.
type MachineTester struct {
	MachineTesterCaller     // Read-only binding to the contract
	MachineTesterTransactor // Write-only binding to the contract
	MachineTesterFilterer   // Log filterer for contract events
}

// MachineTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineTesterSession struct {
	Contract     *MachineTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineTesterCallerSession struct {
	Contract *MachineTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MachineTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineTesterTransactorSession struct {
	Contract     *MachineTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MachineTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineTesterRaw struct {
	Contract *MachineTester // Generic contract binding to access the raw methods on
}

// MachineTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineTesterCallerRaw struct {
	Contract *MachineTesterCaller // Generic read-only contract binding to access the raw methods on
}

// MachineTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineTesterTransactorRaw struct {
	Contract *MachineTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachineTester creates a new instance of MachineTester, bound to a specific deployed contract.
func NewMachineTester(address common.Address, backend bind.ContractBackend) (*MachineTester, error) {
	contract, err := bindMachineTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MachineTester{MachineTesterCaller: MachineTesterCaller{contract: contract}, MachineTesterTransactor: MachineTesterTransactor{contract: contract}, MachineTesterFilterer: MachineTesterFilterer{contract: contract}}, nil
}

// NewMachineTesterCaller creates a new read-only instance of MachineTester, bound to a specific deployed contract.
func NewMachineTesterCaller(address common.Address, caller bind.ContractCaller) (*MachineTesterCaller, error) {
	contract, err := bindMachineTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTesterCaller{contract: contract}, nil
}

// NewMachineTesterTransactor creates a new write-only instance of MachineTester, bound to a specific deployed contract.
func NewMachineTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineTesterTransactor, error) {
	contract, err := bindMachineTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTesterTransactor{contract: contract}, nil
}

// NewMachineTesterFilterer creates a new log filterer instance of MachineTester, bound to a specific deployed contract.
func NewMachineTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineTesterFilterer, error) {
	contract, err := bindMachineTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineTesterFilterer{contract: contract}, nil
}

// bindMachineTester binds a generic wrapper to an already deployed contract.
func bindMachineTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineTester *MachineTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineTester.Contract.MachineTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineTester *MachineTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineTester.Contract.MachineTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineTester *MachineTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineTester.Contract.MachineTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineTester *MachineTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineTester *MachineTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineTester *MachineTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineTester.Contract.contract.Transact(opts, method, params...)
}

// AddStackVal is a free data retrieval call binding the contract method 0x5f098d7f.
//
// Solidity: function addStackVal(bytes data1, bytes data2) pure returns(bytes32)
func (_MachineTester *MachineTesterCaller) AddStackVal(opts *bind.CallOpts, data1 []byte, data2 []byte) ([32]byte, error) {
	var out []interface{}
	err := _MachineTester.contract.Call(opts, &out, "addStackVal", data1, data2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddStackVal is a free data retrieval call binding the contract method 0x5f098d7f.
//
// Solidity: function addStackVal(bytes data1, bytes data2) pure returns(bytes32)
func (_MachineTester *MachineTesterSession) AddStackVal(data1 []byte, data2 []byte) ([32]byte, error) {
	return _MachineTester.Contract.AddStackVal(&_MachineTester.CallOpts, data1, data2)
}

// AddStackVal is a free data retrieval call binding the contract method 0x5f098d7f.
//
// Solidity: function addStackVal(bytes data1, bytes data2) pure returns(bytes32)
func (_MachineTester *MachineTesterCallerSession) AddStackVal(data1 []byte, data2 []byte) ([32]byte, error) {
	return _MachineTester.Contract.AddStackVal(&_MachineTester.CallOpts, data1, data2)
}

// DeserializeMachine is a free data retrieval call binding the contract method 0x5270f3e9.
//
// Solidity: function deserializeMachine(bytes data) pure returns(uint256, bytes32)
func (_MachineTester *MachineTesterCaller) DeserializeMachine(opts *bind.CallOpts, data []byte) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _MachineTester.contract.Call(opts, &out, "deserializeMachine", data)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// DeserializeMachine is a free data retrieval call binding the contract method 0x5270f3e9.
//
// Solidity: function deserializeMachine(bytes data) pure returns(uint256, bytes32)
func (_MachineTester *MachineTesterSession) DeserializeMachine(data []byte) (*big.Int, [32]byte, error) {
	return _MachineTester.Contract.DeserializeMachine(&_MachineTester.CallOpts, data)
}

// DeserializeMachine is a free data retrieval call binding the contract method 0x5270f3e9.
//
// Solidity: function deserializeMachine(bytes data) pure returns(uint256, bytes32)
func (_MachineTester *MachineTesterCallerSession) DeserializeMachine(data []byte) (*big.Int, [32]byte, error) {
	return _MachineTester.Contract.DeserializeMachine(&_MachineTester.CallOpts, data)
}
