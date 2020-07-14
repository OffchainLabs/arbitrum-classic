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
const MachineTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data2\",\"type\":\"bytes\"}],\"name\":\"addStackVal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeMachine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MachineTesterFuncSigs maps the 4-byte function signature to its string representation.
var MachineTesterFuncSigs = map[string]string{
	"5f098d7f": "addStackVal(bytes,bytes)",
	"5270f3e9": "deserializeMachine(bytes)",
}

// MachineTesterBin is the compiled bytecode used for deploying new contracts.
var MachineTesterBin = "0x608060405234801561001057600080fd5b50611181806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635270f3e91461003b5780635f098d7f146100fa575b600080fd5b6100e16004803603602081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610239945050505050565b6040805192835260208301919091528051918290030190f35b6102276004803603604081101561011057600080fd5b81019060208101813564010000000081111561012b57600080fd5b82018360208201111561013d57600080fd5b8035906020019184600183028401116401000000008311171561015f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101b257600080fd5b8201836020820111156101c457600080fd5b803590602001918460018302840111640100000000831117156101e657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061026b945050505050565b60408051918252519081900360200190f35b600080600061024661109b565b6102518560006102ba565b9092509050816102608261035f565b935093505050915091565b6000806102766110f9565b61027e6110f9565b610289866000610424565b9093509150610299856000610424565b90935090506102b06102ab83836105af565b61062d565b9695505050505050565b60006102c461109b565b6102cc61109b565b600060e08201819052806102e0878761074d565b90965091506102ef87876107c1565b6020850152955061030087876107c1565b604085015295506103118787610424565b606085015295506103228787610424565b60808501529550610333878761074d565b60a08501529550610344878761074d565b92845260c084019290925250935083925090505b9250929050565b600060028260e0015114156103765750600061041f565b60018260e00151141561038b5750600161041f565b8151602083015161039b9061062d565b6103a8846040015161062d565b6103b5856060015161062d565b6103c2866080015161062d565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090505b919050565b600061042e6110f9565b83518310610474576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806104818686610855565b9150915061048d61087c565b60ff168160ff1614156104c15760006104a6878461074d565b9093509050826104b582610881565b94509450505050610358565b6104c9610933565b60ff168160ff1614156104eb576104e08683610938565b935093505050610358565b6104f36109da565b60ff168160ff16141561050a576104e086836107c1565b6105126109df565b60ff168160ff1610158015610533575061052a6109e4565b60ff168160ff16105b1561056f5760006105426109df565b8203905060606105538289866109e9565b90945090508361056282610a82565b9550955050505050610358565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b6105b76110f9565b6040805160028082526060828101909352816020015b6105d56110f9565b8152602001906001900390816105cd57905050905082816000815181106105f857fe5b6020026020010181905250838160018151811061061157fe5b602002602001018190525061062581610b94565b949350505050565b600061063761087c565b60ff16826060015160ff16141561065a57815161065390610cf9565b905061041f565b610662610933565b60ff16826060015160ff161415610680576106538260200151610d1d565b6106886109da565b60ff16826060015160ff1614156106aa57815160808301516106539190610e0d565b6106b26109df565b60ff16826060015160ff1614156106eb576106cb6110f9565b6106d88360400151610b94565b90506106e38161062d565b91505061041f565b6106f3610e5e565b60ff16826060015160ff16141561070c5750805161041f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b60008082845110158015610765575060208385510310155b6107a2576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b602083016107b6858563ffffffff610e6316565b915091509250929050565b60006107cb6110f9565b828451101580156107e0575060408385510310155b61081c576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000806108298686610e7f565b9094509150610838868561074d565b9094509050836108488383610e96565b9350935050509250929050565b6000808260010184848151811061086857fe5b016020015190925060f81c90509250929050565b600090565b6108896110f9565b6040805160a08101825283815281516060810183526000808252602082810182905284518281528082018652939490850193908301916108df565b6108cc6110f9565b8152602001906001900390816108c45790505b5090528152604080516000808252602082810190935291909201919061091b565b6109086110f9565b8152602001906001900390816109005790505b50815260006020820152600160409091015292915050565b600190565b60006109426110f9565b8260008061094e6110f9565b600061095a8986610855565b90955093506109698986610855565b9095509250600160ff8516141561098a576109848986610424565b90955091505b6109948986610e7f565b9095509050600160ff851614156109bf57846109b1848385610f47565b965096505050505050610358565b846109ca8483610fcb565b9650965050505050509250929050565b600290565b600390565b600c90565b60006060600083905060608660ff16604051908082528060200260200182016040528015610a3157816020015b610a1e6110f9565b815260200190600190039081610a165790505b50905060005b8760ff168160ff161015610a7557610a4f8784610424565b8351849060ff8516908110610a6057fe5b60209081029190910101529250600101610a37565b5090969095509350505050565b610a8a6110f9565b610a94825161102d565b610ae5576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610b1c57838181518110610aff57fe5b602002602001015160800151820191508080600101915050610aea565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610b76565b610b636110f9565b815260200190600190039081610b5b5790505b50905281526020810194909452600360408501526060909301525090565b610b9c6110f9565b600882511115610bea576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610c17578160200160208202803883390190505b508051909150600160005b82811015610c8757610c46868281518110610c3957fe5b602002602001015161062d565b848281518110610c5257fe5b602002602001018181525050858181518110610c6a57fe5b602002602001015160800151820191508080600101915050610c22565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610ccc578181015183820152602001610cb4565b50505050905001925050506040516020818303038152906040528051906020012090506102b08183610e96565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610d2e57fe5b604082015151610d9357610d40610933565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b90931660218501526022808501919091528251808503909101815260429093019091528151910120905061041f565b610d9b610933565b8260000151610db48460400151600081518110610c3957fe5b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b6000610e176109df565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b606490565b60008160200183511015610e7657600080fd5b50016020015190565b600080602083016107b6858563ffffffff610e6316565b610e9e6110f9565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191610ef4565b610ee16110f9565b815260200190600190039081610ed95790505b50905281526040805160008082526020828101909352919092019190610f30565b610f1d6110f9565b815260200190600190039081610f155790505b508152600260208201526040019290925250919050565b610f4f6110f9565b604080516001808252818301909252606091816020015b610f6e6110f9565b815260200190600190039081610f665790505090508281600081518110610f9157fe5b6020026020010181905250610fc260405180606001604052808760ff16815260200186815260200183815250611034565b95945050505050565b610fd36110f9565b6040805160608101825260ff8516815260208082018590528251600080825291810184526110269383019161101e565b61100b6110f9565b8152602001906001900390816110035790505b509052611034565b9392505050565b6008101590565b61103c6110f9565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190611083565b6110706110f9565b8152602001906001900390816110685790505b50815260016020820181905260409091015292915050565b60408051610100810190915260008152602081016110b76110f9565b81526020016110c46110f9565b81526020016110d16110f9565b81526020016110de6110f9565b81526000602082018190526040820181905260609091015290565b6040518060a001604052806000815260200161111361112d565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fea265627a7a72315820feaae597db68c6c2784b1e069fe8c289f958b30da4c24ff44546a53eadb08d5764736f6c63430005110032"

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
func (_MachineTester *MachineTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_MachineTester *MachineTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MachineTester.contract.Call(opts, out, "addStackVal", data1, data2)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MachineTester.contract.Call(opts, out, "deserializeMachine", data)
	return *ret0, *ret1, err
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
