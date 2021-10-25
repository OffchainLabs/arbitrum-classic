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
	Bin: "0x608060405234801561001057600080fd5b506116db806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635270f3e91461003b5780635f098d7f146100f8575b600080fd5b6100df6004803603602081101561005157600080fd5b810190602081018135600160201b81111561006b57600080fd5b82018360208201111561007d57600080fd5b803590602001918460018302840111600160201b8311171561009e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610233945050505050565b6040805192835260208301919091528051918290030190f35b6102216004803603604081101561010e57600080fd5b810190602081018135600160201b81111561012857600080fd5b82018360208201111561013a57600080fd5b803590602001918460018302840111600160201b8311171561015b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101ad57600080fd5b8201836020820111156101bf57600080fd5b803590602001918460018302840111600160201b831117156101e057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610265945050505050565b60408051918252519081900360200190f35b60008060006102406115b1565b61024b8560006102b4565b90925090508161025a82610359565b935093505050915091565b60008061027061160f565b61027861160f565b61028386600061041e565b909350915061029385600061041e565b90935090506102aa6102a583836105f8565b610676565b9695505050505050565b60006102be6115b1565b6102c66115b1565b600060e08201819052806102da8787610809565b90965091506102e9878761087d565b602085015295506102fa878761087d565b6040850152955061030b878761041e565b6060850152955061031c878761041e565b6080850152955061032d8787610809565b60a0850152955061033e8787610809565b92845260c084019290925250935083925090505b9250929050565b600060028260e00151141561037057506000610419565b60018260e00151141561038557506001610419565b8151602083015161039590610676565b6103a28460400151610676565b6103af8560600151610676565b6103bc8660800151610676565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090505b919050565b600061042861160f565b8351831061046e576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b60008061047b8686610912565b91509150610487610939565b60ff168160ff1614156104bb5760006104a08784610809565b9093509050826104af8261093e565b94509450505050610352565b6104c3610a0b565b60ff168160ff1614156104e5576104da8683610a10565b935093505050610352565b6104ed610ab2565b60ff168160ff161415610504576104da8683610ab7565b61050c610b32565b60ff168160ff1614156105345760006105258784610809565b9093509050826104af82610b37565b61053c610c2f565b60ff168160ff161415610553576104da868361087d565b61055b610c34565b60ff168160ff161015801561057c5750610573610c39565b60ff168160ff16105b156105b857600061058b610c34565b82039050606061059c828986610c3e565b9094509050836105ab82610ce7565b9550955050505050610352565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b61060061160f565b6040805160028082526060828101909352816020015b61061e61160f565b815260200190600190039081610616579050509050828160008151811061064157fe5b6020026020010181905250838160018151811061065a57fe5b602002602001018190525061066e81610e38565b949350505050565b6000610680610939565b60ff168260a0015160ff1614156106a357815161069c90610fb4565b9050610419565b6106ab610a0b565b60ff168260a0015160ff1614156106c95761069c8260200151610fd8565b6106d1610c2f565b60ff168260a0015160ff1614156106f357815160c083015161069c91906110c8565b6106fb610c34565b60ff168260a0015160ff1614156107345761071461160f565b6107218360400151610e38565b905061072c81610676565b915050610419565b61073c611119565b60ff168260a0015160ff16141561075557508051610419565b61075d610b32565b60ff168260a0015160ff1614156107a2575060608082015160408051607b60208083019190915281830193909352815180820383018152930190528151910120610419565b6107aa610ab2565b60ff168260a0015160ff1614156107c85761069c826080015161111e565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b60008082845110158015610821575060208385510310155b61085e576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301610872858563ffffffff61118e16565b915091509250929050565b600061088761160f565b8284511015801561089c575060408385510310155b6108d9576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6000806108e686866111e7565b90945091506108f58685610809565b90945090508361090583836111fe565b9350935050509250929050565b6000808260010184848151811061092557fe5b016020015190925060f81c90509250929050565b600090565b61094661160f565b6040805160e081018252838152815160608101835260008082526020808301829052845182815280820186529394908501939083019161099c565b61098961160f565b8152602001906001900390816109815790505b509052815260200160006040519080825280602002602001820160405280156109df57816020015b6109cc61160f565b8152602001906001900390816109c45790505b508152600060208201526040016109f46112ca565b815260006020820152600160409091015292915050565b600190565b6000610a1a61160f565b82600080610a2661160f565b6000610a328986610912565b9095509350610a418986610912565b9095509250600160ff85161415610a6257610a5c898661041e565b90955091505b610a6c89866111e7565b9095509050600160ff85161415610a975784610a89848385611301565b965096505050505050610352565b84610aa28483611385565b9650965050505050509250929050565b600e90565b6000610ac161160f565b82600080808080610ad28a876111e7565b9096509450610ae18a876111e7565b9096509350610af08a87610809565b9096509250610aff8a876111e7565b9096509150610b0e8a87610809565b909650905085610b2186868686866113e7565b975097505050505050509250929050565b600c90565b610b3f61160f565b6040518060e00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610b8757600080fd5b50604051908082528060200260200182016040528015610bc157816020015b610bae61160f565b815260200190600190039081610ba65790505b50905281526020016000604051908082528060200260200182016040528015610c0457816020015b610bf161160f565b815260200190600190039081610be95790505b508152602001838152602001610c186112ca565b8152600c6020820152600160409091015292915050565b600290565b600390565b600f90565b60006060828160ff871667ffffffffffffffff81118015610c5e57600080fd5b50604051908082528060200260200182016040528015610c9857816020015b610c8561160f565b815260200190600190039081610c7d5790505b50905060005b8760ff168160ff161015610cda57610cb6878461041e565b838360ff1681518110610cc557fe5b60209081029190910101529250600101610c9e565b5090969095509350505050565b610cef61160f565b610cf9825161150c565b610d4a576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610d8157838181518110610d6457fe5b602002602001015160c00151820191508080600101915050610d4f565b506040518060e00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff81118015610dca57600080fd5b50604051908082528060200260200182016040528015610e0457816020015b610df161160f565b815260200190600190039081610de95790505b50905281526020810185905260006040820152606001610e226112ca565b8152600360208201526040019190915292915050565b610e4061160f565b600882511115610e8e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825167ffffffffffffffff81118015610ea857600080fd5b50604051908082528060200260200182016040528015610ed2578160200160208202803683370190505b508051909150600160005b82811015610f4257610f01868281518110610ef457fe5b6020026020010151610676565b848281518110610f0d57fe5b602002602001018181525050858181518110610f2557fe5b602002602001015160c00151820191508080600101915050610edd565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610f87578181015183820152602001610f6f565b50505050905001925050506040516020818303038152906040528051906020012090506102aa81836111fe565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610fe957fe5b60408201515161104e57610ffb610a0b565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b909316602185015260228085019190915282518085039091018152604290930190915281519101209050610419565b611056610a0b565b826000015161106f8460400151600081518110610ef457fe5b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60006110d2610c34565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b606490565b6000611128610ab2565b8260000151836020015184606001518560800151604051602001808660ff1660ff1660f81b815260010185815260200184815260200183815260200182815260200195505050505050604051602081830303815290604052805190602001209050919050565b600081602001835110156111de576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008060208301610872858563ffffffff61118e16565b61120661160f565b6040805160e081018252848152815160608101835260008082526020808301829052845182815280820186529394908501939083019161125c565b61124961160f565b8152602001906001900390816112415790505b5090528152602001600060405190808252806020026020018201604052801561129f57816020015b61128c61160f565b8152602001906001900390816112845790505b508152600060208201526040016112b46112ca565b8152600260208201526040019290925250919050565b6112d2611658565b506040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b61130961160f565b604080516001808252818301909252606091816020015b61132861160f565b815260200190600190039081611320579050509050828160008151811061134b57fe5b602002602001018190525061137c60405180606001604052808760ff16815260200186815260200183815250611513565b95945050505050565b61138d61160f565b6040805160608101825260ff8516815260208082018590528251600080825291810184526113e0938301916113d8565b6113c561160f565b8152602001906001900390816113bd5790505b509052611513565b9392505050565b6113ef61160f565b6040518060e00160405280600081526020016040518060600160405280600060ff1681526020016000801b8152602001600067ffffffffffffffff8111801561143757600080fd5b5060405190808252806020026020018201604052801561147157816020015b61145e61160f565b8152602001906001900390816114565790505b509052815260200160006040519080825280602002602001820160405280156114b457816020015b6114a161160f565b8152602001906001900390816114995790505b5081526020016000801b81526020016040518060a001604052808981526020018881526020018781526020018681526020018581525081526020016003600b0160ff1681526020016001815250905095945050505050565b6008101590565b61151b61160f565b6040518060e0016040528060008152602001838152602001600067ffffffffffffffff8111801561154b57600080fd5b5060405190808252806020026020018201604052801561158557816020015b61157261160f565b81526020019060019003908161156a5790505b5081526000602082015260400161159a6112ca565b815260016020820181905260409091015292915050565b60408051610100810190915260008152602081016115cd61160f565b81526020016115da61160f565b81526020016115e761160f565b81526020016115f461160f565b81526000602082018190526040820181905260609091015290565b6040518060e0016040528060008152602001611629611686565b81526060602082018190526000604083015201611644611658565b815260006020820181905260409091015290565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b604080516060808201835260008083526020830152918101919091529056fea264697066735822122059d1504a1f5682bf2aa331de5ce7e1efff671045233387dc2f24d6250a45b0a564736f6c634300060b0033",
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
