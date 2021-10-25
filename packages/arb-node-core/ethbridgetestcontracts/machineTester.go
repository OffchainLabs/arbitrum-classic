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
var MachineTesterBin = "0x608060405234801561001057600080fd5b506116f8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635270f3e91461003b5780635f098d7f146100fa575b600080fd5b6100e16004803603602081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610239945050505050565b6040805192835260208301919091528051918290030190f35b6102276004803603604081101561011057600080fd5b81019060208101813564010000000081111561012b57600080fd5b82018360208201111561013d57600080fd5b8035906020019184600183028401116401000000008311171561015f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101b257600080fd5b8201836020820111156101c457600080fd5b803590602001918460018302840111640100000000831117156101e657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061026b945050505050565b60408051918252519081900360200190f35b60008060006102466115c1565b6102518560006102ba565b90925090508161026082610370565b935093505050915091565b60008061027661162c565b61027e61162c565b61028986600061044b565b909350915061029985600061044b565b90935090506102b06102ab8383610625565b6106a3565b9695505050505050565b60006102c46115c1565b6102cc6115c1565b60006101008201819052806102e18787610836565b90965091506102f087876108a4565b6020850152955061030187876108a4565b60408501529550610312878761044b565b60608501529550610323878761044b565b608085015295506103348787610836565b60a085015295506103458787610836565b9096509050610354878761044b565b60e085015291835260c0830152935083925090505b9250929050565b60006002826101000151141561038857506000610446565b6001826101000151141561039e57506001610446565b815160208301516103ae906106a3565b6103bb84604001516106a3565b6103c885606001516106a3565b6103d586608001516106a3565b8660a001518760c001516103ec8960e001516106a3565b60405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012090505b919050565b600061045561162c565b8351831061049b576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806104a88686610939565b915091506104b4610960565b60ff168160ff1614156104e85760006104cd8784610836565b9093509050826104dc82610965565b94509450505050610369565b6104f0610a32565b60ff168160ff161415610512576105078683610a37565b935093505050610369565b61051a610ad9565b60ff168160ff161415610531576105078683610ade565b610539610b59565b60ff168160ff1614156105615760006105528784610836565b9093509050826104dc82610b5e565b610569610c56565b60ff168160ff1614156105805761050786836108a4565b610588610c5b565b60ff168160ff16101580156105a957506105a0610c60565b60ff168160ff16105b156105e55760006105b8610c5b565b8203905060606105c9828986610c65565b9094509050836105d882610d0e565b9550955050505050610369565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b61062d61162c565b6040805160028082526060828101909352816020015b61064b61162c565b815260200190600190039081610643579050509050828160008151811061066e57fe5b6020026020010181905250838160018151811061068757fe5b602002602001018190525061069b81610e5f565b949350505050565b60006106ad610960565b60ff168260a0015160ff1614156106d05781516106c990610fd8565b9050610446565b6106d8610a32565b60ff168260a0015160ff1614156106f6576106c98260200151610ffc565b6106fe610c56565b60ff168260a0015160ff16141561072057815160c08301516106c991906110e4565b610728610c5b565b60ff168260a0015160ff1614156107615761074161162c565b61074e8360400151610e5f565b9050610759816106a3565b915050610446565b610769611132565b60ff168260a0015160ff16141561078257508051610446565b61078a610b59565b60ff168260a0015160ff1614156107cf575060608082015160408051607b60208083019190915281830193909352815180820383018152930190528151910120610446565b6107d7610ad9565b60ff168260a0015160ff1614156107f5576106c98260800151611137565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6000808284511015801561084e575060208385510310155b61088b576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6020830161089985856111a4565b915091509250929050565b60006108ae61162c565b828451101580156108c3575060408385510310155b610900576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60008061090d86866111fd565b909450915061091c8685610836565b90945090508361092c838361120e565b9350935050509250929050565b6000808260010184848151811061094c57fe5b016020015190925060f81c90509250929050565b600090565b61096d61162c565b6040805160e08101825283815281516060810183526000808252602080830182905284518281528082018652939490850193908301916109c3565b6109b061162c565b8152602001906001900390816109a85790505b50905281526020016000604051908082528060200260200182016040528015610a0657816020015b6109f361162c565b8152602001906001900390816109eb5790505b50815260006020820152604001610a1b6112da565b815260006020820152600160409091015292915050565b600190565b6000610a4161162c565b82600080610a4d61162c565b6000610a598986610939565b9095509350610a688986610939565b9095509250600160ff85161415610a8957610a83898661044b565b90955091505b610a9389866111fd565b9095509050600160ff85161415610abe5784610ab0848385611311565b965096505050505050610369565b84610ac98483611395565b9650965050505050509250929050565b600e90565b6000610ae861162c565b82600080808080610af98a876111fd565b9096509450610b088a876111fd565b9096509350610b178a87610836565b9096509250610b268a876111fd565b9096509150610b358a87610836565b909650905085610b4886868686866113f7565b975097505050505050509250929050565b600c90565b610b6661162c565b6040518060e00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610bae57600080fd5b50604051908082528060200260200182016040528015610be857816020015b610bd561162c565b815260200190600190039081610bcd5790505b50905281526020016000604051908082528060200260200182016040528015610c2b57816020015b610c1861162c565b815260200190600190039081610c105790505b508152602001838152602001610c3f6112da565b8152600c6020820152600160409091015292915050565b600290565b600390565b600f90565b60006060828160ff871667ffffffffffffffff81118015610c8557600080fd5b50604051908082528060200260200182016040528015610cbf57816020015b610cac61162c565b815260200190600190039081610ca45790505b50905060005b8760ff168160ff161015610d0157610cdd878461044b565b838360ff1681518110610cec57fe5b60209081029190910101529250600101610cc5565b5090969095509350505050565b610d1661162c565b610d20825161151c565b610d71576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610da857838181518110610d8b57fe5b602002602001015160c00151820191508080600101915050610d76565b506040518060e00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610df157600080fd5b50604051908082528060200260200182016040528015610e2b57816020015b610e1861162c565b815260200190600190039081610e105790505b50905281526020810185905260006040820152606001610e496112da565b8152600360208201526040019190915292915050565b610e6761162c565b600882511115610eb5576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825167ffffffffffffffff81118015610ecf57600080fd5b50604051908082528060200260200182016040528015610ef9578160200160208202803683370190505b508051909150600160005b82811015610f6957610f28868281518110610f1b57fe5b60200260200101516106a3565b848281518110610f3457fe5b602002602001018181525050858181518110610f4c57fe5b602002602001015160c00151820191508080600101915050610f04565b506000835184604051602001808360ff1660f81b8152600101828051906020019060200280838360005b83811015610fab578181015183820152602001610f93565b50505050905001925050506040516020818303038152906040528051906020012090506102b0818361120e565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061100d57fe5b6040820151516110705761101f610a32565b82600001518360200151604051602001808460ff1660f81b81526001018360ff1660f81b81526001018281526020019350505050604051602081830303815290604052805190602001209050610446565b611078610a32565b82600001516110918460400151600081518110610f1b57fe5b8460200151604051602001808560ff1660f81b81526001018460ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60006110ee610c5b565b8383604051602001808460ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b606490565b6000611141610ad9565b8260000151836020015184606001518560800151604051602001808660ff1660f81b815260010185815260200184815260200183815260200182815260200195505050505050604051602081830303815290604052805190602001209050919050565b600081602001835110156111f4576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000806020830161089985856111a4565b61121661162c565b6040805160e081018252848152815160608101835260008082526020808301829052845182815280820186529394908501939083019161126c565b61125961162c565b8152602001906001900390816112515790505b509052815260200160006040519080825280602002602001820160405280156112af57816020015b61129c61162c565b8152602001906001900390816112945790505b508152600060208201526040016112c46112da565b8152600260208201526040019290925250919050565b6112e2611675565b506040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b61131961162c565b604080516001808252818301909252606091816020015b61133861162c565b815260200190600190039081611330579050509050828160008151811061135b57fe5b602002602001018190525061138c60405180606001604052808760ff16815260200186815260200183815250611523565b95945050505050565b61139d61162c565b6040805160608101825260ff8516815260208082018590528251600080825291810184526113f0938301916113e8565b6113d561162c565b8152602001906001900390816113cd5790505b509052611523565b9392505050565b6113ff61162c565b6040518060e00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff8111801561144757600080fd5b5060405190808252806020026020018201604052801561148157816020015b61146e61162c565b8152602001906001900390816114665790505b509052815260200160006040519080825280602002602001820160405280156114c457816020015b6114b161162c565b8152602001906001900390816114a95790505b5081526020016000801b81526020016040518060a001604052808981526020018881526020018781526020018681526020018581525081526020016003600b0160ff1681526020016001815250905095945050505050565b6008101590565b61152b61162c565b6040518060e0016040528060008152602001838152602001600067ffffffffffffffff8111801561155b57600080fd5b5060405190808252806020026020018201604052801561159557816020015b61158261162c565b81526020019060019003908161157a5790505b508152600060208201526040016115aa6112da565b815260016020820181905260409091015292915050565b60408051610120810190915260008152602081016115dd61162c565b81526020016115ea61162c565b81526020016115f761162c565b815260200161160461162c565b8152600060208201819052604082015260600161161f61162c565b8152602001600081525090565b6040518060e00160405280600081526020016116466116a3565b81526060602082018190526000604083015201611661611675565b815260006020820181905260409091015290565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b604080516060808201835260008083526020830152918101919091529056fea2646970667358221220528824eab1d69c35cdb06bcb973b8453e7de38c9968aac105b647a9adb4f2c1364736f6c634300060c0033"

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
