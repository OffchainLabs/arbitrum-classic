// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MachineTesterMetaData contains all meta data concerning the MachineTester contract.
var MachineTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data2\",\"type\":\"bytes\"}],\"name\":\"addStackVal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeMachine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506113cf806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635270f3e91461003b5780635f098d7f146100f8575b600080fd5b6100df6004803603602081101561005157600080fd5b810190602081018135600160201b81111561006b57600080fd5b82018360208201111561007d57600080fd5b803590602001918460018302840111600160201b8311171561009e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610233945050505050565b6040805192835260208301919091528051918290030190f35b6102216004803603604081101561010e57600080fd5b810190602081018135600160201b81111561012857600080fd5b82018360208201111561013a57600080fd5b803590602001918460018302840111600160201b8311171561015b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101ad57600080fd5b8201836020820111156101bf57600080fd5b803590602001918460018302840111600160201b831117156101e057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610265945050505050565b60408051918252519081900360200190f35b60008060006102406112df565b61024b8560006102b4565b90925090508161025a82610359565b935093505050915091565b60008061027061133d565b61027861133d565b61028386600061041e565b909350915061029385600061041e565b90935090506102aa6102a583836105d9565b610657565b9695505050505050565b60006102be6112df565b6102c66112df565b600060e08201819052806102da87876107c4565b90965091506102e98787610838565b602085015295506102fa8787610838565b6040850152955061030b878761041e565b6060850152955061031c878761041e565b6080850152955061032d87876107c4565b60a0850152955061033e87876107c4565b92845260c084019290925250935083925090505b9250929050565b600060028260e00151141561037057506000610419565b60018260e00151141561038557506001610419565b8151602083015161039590610657565b6103a28460400151610657565b6103af8560600151610657565b6103bc8660800151610657565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090505b919050565b600061042861133d565b8351831061046e576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b60008061047b86866108cd565b915091506104876108f4565b60ff168160ff1614156104bb5760006104a087846107c4565b9093509050826104af826108f9565b94509450505050610352565b6104c36109b9565b60ff168160ff1614156104e5576104da86836109be565b935093505050610352565b6104ed610a60565b60ff168160ff16141561051557600061050687846107c4565b9093509050826104af82610a65565b61051d610b52565b60ff168160ff161415610534576104da8683610838565b61053c610b57565b60ff168160ff161015801561055d5750610554610b5c565b60ff168160ff16105b1561059957600061056c610b57565b82039050606061057d828986610b61565b90945090508361058c82610c0a565b9550955050505050610352565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b6105e161133d565b6040805160028082526060828101909352816020015b6105ff61133d565b8152602001906001900390816105f7579050509050828160008151811061062257fe5b6020026020010181905250838160018151811061063b57fe5b602002602001018190525061064f81610d4c565b949350505050565b60006106616108f4565b60ff16826080015160ff16141561068457815161067d90610ec8565b9050610419565b61068c6109b9565b60ff16826080015160ff1614156106aa5761067d8260200151610eec565b6106b2610b52565b60ff16826080015160ff1614156106d457815160a083015161067d9190610fdc565b6106dc610b57565b60ff16826080015160ff161415610715576106f561133d565b6107028360400151610d4c565b905061070d81610657565b915050610419565b61071d61102d565b60ff16826080015160ff16141561073657508051610419565b61073e610a60565b60ff16826080015160ff161415610783575060608082015160408051607b60208083019190915281830193909352815180820383018152930190528151910120610419565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b600080828451101580156107dc575060208385510310155b610819576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6020830161082d858563ffffffff61103216565b915091509250929050565b600061084261133d565b82845110158015610857575060408385510310155b610894576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6000806108a1868661108b565b90945091506108b086856107c4565b9094509050836108c083836110a2565b9350935050509250929050565b600080826001018484815181106108e057fe5b016020015190925060f81c90509250929050565b600090565b61090161133d565b6040805160c0810182528381528151606081018352600080825260208083018290528451828152808201865293949085019390830191610957565b61094461133d565b81526020019060019003908161093c5790505b5090528152602001600060405190808252806020026020018201604052801561099a57816020015b61098761133d565b81526020019060019003908161097f5790505b5081526000602082018190526040820152600160609091015292915050565b600190565b60006109c861133d565b826000806109d461133d565b60006109e089866108cd565b90955093506109ef89866108cd565b9095509250600160ff85161415610a1057610a0a898661041e565b90955091505b610a1a898661108b565b9095509050600160ff85161415610a455784610a37848385611161565b965096505050505050610352565b84610a5084836111e5565b9650965050505050509250929050565b600c90565b610a6d61133d565b6040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610ab557600080fd5b50604051908082528060200260200182016040528015610aef57816020015b610adc61133d565b815260200190600190039081610ad45790505b50905281526020016000604051908082528060200260200182016040528015610b3257816020015b610b1f61133d565b815260200190600190039081610b175790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b600390565b600d90565b60006060828160ff871667ffffffffffffffff81118015610b8157600080fd5b50604051908082528060200260200182016040528015610bbb57816020015b610ba861133d565b815260200190600190039081610ba05790505b50905060005b8760ff168160ff161015610bfd57610bd9878461041e565b838360ff1681518110610be857fe5b60209081029190910101529250600101610bc1565b5090969095509350505050565b610c1261133d565b610c1c8251611247565b610c6d576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610ca457838181518110610c8757fe5b602002602001015160a00151820191508080600101915050610c72565b506040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610ced57600080fd5b50604051908082528060200260200182016040528015610d2757816020015b610d1461133d565b815260200190600190039081610d0c5790505b5090528152602081019490945260006040850152600360608501526080909301525090565b610d5461133d565b600882511115610da2576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825167ffffffffffffffff81118015610dbc57600080fd5b50604051908082528060200260200182016040528015610de6578160200160208202803683370190505b508051909150600160005b82811015610e5657610e15868281518110610e0857fe5b6020026020010151610657565b848281518110610e2157fe5b602002602001018181525050858181518110610e3957fe5b602002602001015160a00151820191508080600101915050610df1565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610e9b578181015183820152602001610e83565b50505050905001925050506040516020818303038152906040528051906020012090506102aa81836110a2565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610efd57fe5b604082015151610f6257610f0f6109b9565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b909316602185015260228085019190915282518085039091018152604290930190915281519101209050610419565b610f6a6109b9565b8260000151610f838460400151600081518110610e0857fe5b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b6000610fe6610b57565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b606490565b60008160200183511015611082576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000806020830161082d858563ffffffff61103216565b6110aa61133d565b6040805160c0810182528481528151606081018352600080825260208083018290528451828152808201865293949085019390830191611100565b6110ed61133d565b8152602001906001900390816110e55790505b5090528152602001600060405190808252806020026020018201604052801561114357816020015b61113061133d565b8152602001906001900390816111285790505b50815260006020820152600260408201526060019290925250919050565b61116961133d565b604080516001808252818301909252606091816020015b61118861133d565b81526020019060019003908161118057905050905082816000815181106111ab57fe5b60200260200101819052506111dc60405180606001604052808760ff1681526020018681526020018381525061124e565b95945050505050565b6111ed61133d565b6040805160608101825260ff85168152602080820185905282516000808252918101845261124093830191611238565b61122561133d565b81526020019060019003908161121d5790505b50905261124e565b9392505050565b6008101590565b61125661133d565b6040518060c0016040528060008152602001838152602001600067ffffffffffffffff8111801561128657600080fd5b506040519080825280602002602001820160405280156112c057816020015b6112ad61133d565b8152602001906001900390816112a55790505b5081526000602082015260016040820181905260609091015292915050565b60408051610100810190915260008152602081016112fb61133d565b815260200161130861133d565b815260200161131561133d565b815260200161132261133d565b81526000602082018190526040820181905260609091015290565b6040518060c001604052806000815260200161135761137a565b815260606020820181905260006040830181905290820181905260809091015290565b604080516060808201835260008083526020830152918101919091529056fea2646970667358221220f4be8829982394560bb6d19d80558c026bd2d911b3a026d4bd1cb1e65a30c97164736f6c634300060b0033",
}

// MachineTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use MachineTesterMetaData.ABI instead.
var MachineTesterABI = MachineTesterMetaData.ABI

// MachineTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MachineTesterMetaData.Bin instead.
var MachineTesterBin = MachineTesterMetaData.Bin

// DeployMachineTester deploys a new Ethereum contract, binding an instance of MachineTester to it.
func DeployMachineTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MachineTester, error) {
	parsed, err := MachineTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MachineTesterBin), backend)
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
