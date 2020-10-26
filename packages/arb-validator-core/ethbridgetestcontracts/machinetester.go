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
var MachineTesterBin = "0x608060405234801561001057600080fd5b506111f2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635270f3e91461003b5780635f098d7f146100fa575b600080fd5b6100e16004803603602081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610239945050505050565b6040805192835260208301919091528051918290030190f35b6102276004803603604081101561011057600080fd5b81019060208101813564010000000081111561012b57600080fd5b82018360208201111561013d57600080fd5b8035906020019184600183028401116401000000008311171561015f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101b257600080fd5b8201836020820111156101c457600080fd5b803590602001918460018302840111640100000000831117156101e657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061026b945050505050565b60408051918252519081900360200190f35b60008060006102466110ff565b6102518560006102ba565b90925090508161026082610370565b935093505050915091565b60008061027661116a565b61027e61116a565b61028986600061044b565b909350915061029985600061044b565b90935090506102b06102ab83836105d6565b610654565b9695505050505050565b60006102c46110ff565b6102cc6110ff565b60006101008201819052806102e18787610774565b90965091506102f087876107e8565b6020850152955061030187876107e8565b60408501529550610312878761044b565b60608501529550610323878761044b565b608085015295506103348787610774565b60a085015295506103458787610774565b9096509050610354878761044b565b60e085015291835260c0830152935083925090505b9250929050565b60006002826101000151141561038857506000610446565b6001826101000151141561039e57506001610446565b815160208301516103ae90610654565b6103bb8460400151610654565b6103c88560600151610654565b6103d58660800151610654565b8660a001518760c001516103ec8960e00151610654565b60405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012090505b919050565b600061045561116a565b8351831061049b576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806104a8868661087c565b915091506104b46108a3565b60ff168160ff1614156104e85760006104cd8784610774565b9093509050826104dc826108a8565b94509450505050610369565b6104f061095a565b60ff168160ff16141561051257610507868361095f565b935093505050610369565b61051a610a01565b60ff168160ff1614156105315761050786836107e8565b610539610a06565b60ff168160ff161015801561055a5750610551610a0b565b60ff168160ff16105b15610596576000610569610a06565b82039050606061057a828986610a10565b90945090508361058982610aa9565b9550955050505050610369565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b6105de61116a565b6040805160028082526060828101909352816020015b6105fc61116a565b8152602001906001900390816105f4579050509050828160008151811061061f57fe5b6020026020010181905250838160018151811061063857fe5b602002602001018190525061064c81610bbb565b949350505050565b600061065e6108a3565b60ff16826060015160ff16141561068157815161067a90610d20565b9050610446565b61068961095a565b60ff16826060015160ff1614156106a75761067a8260200151610d44565b6106af610a01565b60ff16826060015160ff1614156106d1578151608083015161067a9190610e34565b6106d9610a06565b60ff16826060015160ff161415610712576106f261116a565b6106ff8360400151610bbb565b905061070a81610654565b915050610446565b61071a610e85565b60ff16826060015160ff16141561073357508051610446565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6000808284511015801561078c575060208385510310155b6107c9576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b602083016107dd858563ffffffff610e8a16565b915091509250929050565b60006107f261116a565b82845110158015610807575060408385510310155b610843576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000806108508686610ee3565b909450915061085f8685610774565b90945090508361086f8383610efa565b9350935050509250929050565b6000808260010184848151811061088f57fe5b016020015190925060f81c90509250929050565b600090565b6108b061116a565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610906565b6108f361116a565b8152602001906001900390816108eb5790505b50905281526040805160008082526020828101909352919092019190610942565b61092f61116a565b8152602001906001900390816109275790505b50815260006020820152600160409091015292915050565b600190565b600061096961116a565b8260008061097561116a565b6000610981898661087c565b9095509350610990898661087c565b9095509250600160ff851614156109b1576109ab898661044b565b90955091505b6109bb8986610ee3565b9095509050600160ff851614156109e657846109d8848385610fab565b965096505050505050610369565b846109f1848361102f565b9650965050505050509250929050565b600290565b600390565b600c90565b60006060600083905060608660ff16604051908082528060200260200182016040528015610a5857816020015b610a4561116a565b815260200190600190039081610a3d5790505b50905060005b8760ff168160ff161015610a9c57610a76878461044b565b8351849060ff8516908110610a8757fe5b60209081029190910101529250600101610a5e565b5090969095509350505050565b610ab161116a565b610abb8251611091565b610b0c576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610b4357838181518110610b2657fe5b602002602001015160800151820191508080600101915050610b11565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610b9d565b610b8a61116a565b815260200190600190039081610b825790505b50905281526020810194909452600360408501526060909301525090565b610bc361116a565b600882511115610c11576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610c3e578160200160208202803883390190505b508051909150600160005b82811015610cae57610c6d868281518110610c6057fe5b6020026020010151610654565b848281518110610c7957fe5b602002602001018181525050858181518110610c9157fe5b602002602001015160800151820191508080600101915050610c49565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610cf3578181015183820152602001610cdb565b50505050905001925050506040516020818303038152906040528051906020012090506102b08183610efa565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610d5557fe5b604082015151610dba57610d6761095a565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b909316602185015260228085019190915282518085039091018152604290930190915281519101209050610446565b610dc261095a565b8260000151610ddb8460400151600081518110610c6057fe5b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b6000610e3e610a06565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b606490565b60008160200183511015610eda576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b600080602083016107dd858563ffffffff610e8a16565b610f0261116a565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191610f58565b610f4561116a565b815260200190600190039081610f3d5790505b50905281526040805160008082526020828101909352919092019190610f94565b610f8161116a565b815260200190600190039081610f795790505b508152600260208201526040019290925250919050565b610fb361116a565b604080516001808252818301909252606091816020015b610fd261116a565b815260200190600190039081610fca5790505090508281600081518110610ff557fe5b602002602001018190525061102660405180606001604052808760ff16815260200186815260200183815250611098565b95945050505050565b61103761116a565b6040805160608101825260ff85168152602080820185905282516000808252918101845261108a93830191611082565b61106f61116a565b8152602001906001900390816110675790505b509052611098565b9392505050565b6008101590565b6110a061116a565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906110e7565b6110d461116a565b8152602001906001900390816110cc5790505b50815260016020820181905260409091015292915050565b604080516101208101909152600081526020810161111b61116a565b815260200161112861116a565b815260200161113561116a565b815260200161114261116a565b8152600060208201819052604082015260600161115d61116a565b8152602001600081525090565b6040518060a001604052806000815260200161118461119e565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fea265627a7a723158207728a9c8f671c6e42acd10afec3b8e612663f0d3329fd4940e4f38635e1265a164736f6c63430005110032"

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
