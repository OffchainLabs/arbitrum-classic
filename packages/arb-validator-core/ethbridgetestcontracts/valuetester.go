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
const ValueTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"}],\"name\":\"bytesToBytestackHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"bytestackToBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashTestTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"innerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueSize\",\"type\":\"uint256\"}],\"name\":\"hashTuplePreImage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueTesterFuncSigs maps the 4-byte function signature to its string representation.
var ValueTesterFuncSigs = map[string]string{
	"b325b7d0": "bytesToBytestackHash(bytes,uint256,uint256)",
	"e4d476f4": "bytestackToBytes(bytes,uint256)",
	"98206792": "deserializeHash(bytes,uint256)",
	"fd5d0c8b": "hashTestTuple()",
	"c6d25c8e": "hashTuplePreImage(bytes32,uint256)",
}

// ValueTesterBin is the compiled bytecode used for deploying new contracts.
var ValueTesterBin = "0x608060405234801561001057600080fd5b50611638806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063982067921461005c578063b325b7d01461011d578063c6d25c8e146101da578063e4d476f4146101fd578063fd5d0c8b1461032f575b600080fd5b6101046004803603604081101561007257600080fd5b81019060208101813564010000000081111561008d57600080fd5b82018360208201111561009f57600080fd5b803590602001918460018302840111640100000000831117156100c157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610337915050565b6040805192835260208301919091528051918290030190f35b6101c86004803603606081101561013357600080fd5b81019060208101813564010000000081111561014e57600080fd5b82018360208201111561016057600080fd5b8035906020019184600183028401116401000000008311171561018257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561036a565b60408051918252519081900360200190f35b6101c8600480360360408110156101f057600080fd5b5080359060200135610387565b6102a56004803603604081101561021357600080fd5b81019060208101813564010000000081111561022e57600080fd5b82018360208201111561024057600080fd5b8035906020019184600183028401116401000000008311171561026257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061039a915050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102f25781810151838201526020016102da565b50505050905090810190601f16801561031f5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6101c86103b7565b60008060006103446115b0565b61034e868661047f565b915091508161035c8261060a565b9350935050505b9250929050565b600061037f61037a85858561072f565b61060a565b949350505050565b600061039383836108b3565b9392505050565b60008060606103a98585610904565b9250925092505b9250925092565b60408051600280825260608281019093526000929190816020015b6103da6115b0565b8152602001906001900390816103d25790505090506103f9606f610b5b565b8160008151811061040657fe5b6020026020010181905250610455600060405190808252806020026020018201604052801561044f57816020015b61043c6115b0565b8152602001906001900390816104345790505b50610c0d565b8160018151811061046257fe5b602002602001018190525061047961037a82610c0d565b91505090565b60006104896115b0565b835183106104cf576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806104dc8686610d1f565b915091506104e8610d46565b60ff168160ff16141561051c5760006105018784610d4b565b90935090508261051082610b5b565b94509450505050610363565b610524610dbf565b60ff168160ff1614156105465761053b8683610dc4565b935093505050610363565b61054e610e66565b60ff168160ff1614156105655761053b8683610e6b565b61056d610ef2565b60ff168160ff161015801561058e5750610585610ef7565b60ff168160ff16105b156105ca57600061059d610ef2565b8203905060606105ae828986610efc565b9094509050836105bd82610c0d565b9550955050505050610363565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b6000610614610d46565b60ff16826060015160ff16141561063757815161063090610f95565b905061072a565b61063f610dbf565b60ff16826060015160ff16141561065d576106308260200151610fb9565b610665610e66565b60ff16826060015160ff161415610687578151608083015161063091906108b3565b61068f610ef2565b60ff16826060015160ff1614156106c8576106a86115b0565b6106b583604001516110b6565b90506106c08161060a565b91505061072a565b6106d0611218565b60ff16826060015160ff1614156106e95750805161072a565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b919050565b6107376115b0565b602082046107436115b0565b61074b61121d565b60408051600280825260608281019093529293509091816020015b61076e6115b0565b81526020019060019003908161076657905050905060005b838110156107ef576107ab6107a689602084028a0163ffffffff61126316565b610b5b565b826000815181106107b857fe5b602002602001018190525082826001815181106107d157fe5b60200260200101819052506107e5826110b6565b9250600101610786565b50602085061561086557600061081188601f198989010163ffffffff61126316565b9050602086066020036008021b61082781610b5b565b8260008151811061083457fe5b6020026020010181905250828260018151811061084d57fe5b6020026020010181905250610861826110b6565b9250505b61086e85610b5b565b8160008151811061087b57fe5b6020026020010181905250818160018151811061089457fe5b60200260200101819052506108a8816110b6565b979650505050505050565b60006108bd610ef2565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b6000806060600061091586866112bc565b919550935090508361092b5750600092506103b0565b60208104601f8216600081610941576000610944565b60015b60ff1683019050606083604051908082528060200260200182016040528015610977578160200160208202803883390190505b5090506060836040519080825280601f01601f1916602001820160405280156109a7576020820181803883390190505b5090506000805b84811015610a725760006109c28e8c6112bc565b919d509b5090508b6109e1575060009a506103b0975050505050505050565b811580156109ef5750600087115b15610a42578060005b88811015610a3b57818160208110610a0c57fe5b1a60f81b868281518110610a1c57fe5b60200101906001600160f81b031916908160001a9053506001016109f8565b5050610a69565b8060001b858460018b030381518110610a5757fe5b60209081029190910101526001909201915b506001016109ae565b506000610a7f8d8b610d1f565b909a509050610a8c610ef2565b60ff168160ff1614610aaa5750600099506103b09650505050505050565b60018a858560405160200180838051906020019060200280838360005b83811015610adf578181015183820152602001610ac7565b5050505090500182805190602001908083835b60208310610b115780518252601f199092019160209182019101610af2565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529a509a509a5050505050505050509250925092565b610b636115b0565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610bb9565b610ba66115b0565b815260200190600190039081610b9e5790505b50905281526040805160008082526020828101909352919092019190610bf5565b610be26115b0565b815260200190600190039081610bda5790505b50815260006020820152600160409091015292915050565b610c156115b0565b610c1f825161131e565b610c70576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610ca757838181518110610c8a57fe5b602002602001015160800151820191508080600101915050610c75565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610d01565b610cee6115b0565b815260200190600190039081610ce65790505b50905281526020810194909452600360408501526060909301525090565b60008082600101848481518110610d3257fe5b016020015190925060f81c90509250929050565b600090565b60008082845110158015610d63575060208385510310155b610da0576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301610db4858563ffffffff61126316565b915091509250929050565b600190565b6000610dce6115b0565b82600080610dda6115b0565b6000610de68986610d1f565b9095509350610df58986610d1f565b9095509250600160ff85161415610e1657610e10898661047f565b90955091505b610e208986611325565b9095509050600160ff85161415610e4b5784610e3d84838561133c565b965096505050505050610363565b84610e5684836113c0565b9650965050505050509250929050565b600290565b6000610e756115b0565b82845110158015610e8a575060408385510310155b610ec6576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b600080610ed38686611325565b9094509150610ee28685610d4b565b90945090508361035c838361141b565b600390565b600c90565b60006060600083905060608660ff16604051908082528060200260200182016040528015610f4457816020015b610f316115b0565b815260200190600190039081610f295790505b50905060005b8760ff168160ff161015610f8857610f62878461047f565b8351849060ff8516908110610f7357fe5b60209081029190910101529250600101610f4a565b5090969095509350505050565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610fca57fe5b60408201515161102f57610fdc610dbf565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b90931660218501526022808501919091528251808503909101815260429093019091528151910120905061072a565b611037610dbf565b826000015161105d846040015160008151811061105057fe5b602002602001015161060a565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b6110be6115b0565b60088251111561110c576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611139578160200160208202803883390190505b508051909150600160005b8281101561119c5761115b86828151811061105057fe5b84828151811061116757fe5b60200260200101818152505085818151811061117f57fe5b602002602001015160800151820191508080600101915050611144565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156111e15781810151838201526020016111c9565b505050509050019250505060405160208183030381529060405280519060200120905061120e818361141b565b9695505050505050565b606490565b6112256115b0565b6040805160008082526020820190925261125e9161044f565b6112466115b0565b81526020019060019003908161123e57905050610c0d565b905090565b600081602001835110156112b3576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000806000806112cc8686610d1f565b90935090506112d9610ef2565b60020160ff168160ff16146112f25750600092506103b0565b6112fc86846114cc565b91955093509150836113125750600092506103b0565b50600192509250925092565b6008101590565b60008060208301610db4858563ffffffff61126316565b6113446115b0565b604080516001808252818301909252606091816020015b6113636115b0565b81526020019060019003908161135b579050509050828160008151811061138657fe5b60200260200101819052506113b760405180606001604052808760ff16815260200186815260200183815250611549565b95945050505050565b6113c86115b0565b6040805160608101825260ff85168152602080820185905282516000808252918101845261039393830191611413565b6114006115b0565b8152602001906001900390816113f85790505b509052611549565b6114236115b0565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191611479565b6114666115b0565b81526020019060019003908161145e5790505b509052815260408051600080825260208281019093529190920191906114b5565b6114a26115b0565b81526020019060019003908161149a5790505b508152600260208201526040019290925250919050565b60008060008085519050848110806114e657506021858203105b8061150e57506114f4610d46565b60ff1686868151811061150357fe5b016020015160f81c14155b156115235750600092508391508290506103b0565b60016021860161153b8888840163ffffffff61126316565b935093509350509250925092565b6115516115b0565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190611598565b6115856115b0565b81526020019060019003908161157d5790505b50815260016020820181905260409091015292915050565b6040518060a00160405280600081526020016115ca6115e4565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fea265627a7a723158205298f74ecc04943e4e8c7fc06ab64b78785aa76db574504e39ac2e12dc151cf264736f6c63430005110032"

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
func (_ValueTester *ValueTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ValueTester *ValueTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// BytesToBytestackHash is a free data retrieval call binding the contract method 0xb325b7d0.
//
// Solidity: function bytesToBytestackHash(bytes data, uint256 startOffset, uint256 dataLength) pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) BytesToBytestackHash(opts *bind.CallOpts, data []byte, startOffset *big.Int, dataLength *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "bytesToBytestackHash", data, startOffset, dataLength)
	return *ret0, err
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0xb325b7d0.
//
// Solidity: function bytesToBytestackHash(bytes data, uint256 startOffset, uint256 dataLength) pure returns(bytes32)
func (_ValueTester *ValueTesterSession) BytesToBytestackHash(data []byte, startOffset *big.Int, dataLength *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.BytesToBytestackHash(&_ValueTester.CallOpts, data, startOffset, dataLength)
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0xb325b7d0.
//
// Solidity: function bytesToBytestackHash(bytes data, uint256 startOffset, uint256 dataLength) pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) BytesToBytestackHash(data []byte, startOffset *big.Int, dataLength *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.BytesToBytestackHash(&_ValueTester.CallOpts, data, startOffset, dataLength)
}

// BytestackToBytes is a free data retrieval call binding the contract method 0xe4d476f4.
//
// Solidity: function bytestackToBytes(bytes data, uint256 offset) pure returns(bool, uint256, bytes)
func (_ValueTester *ValueTesterCaller) BytestackToBytes(opts *bind.CallOpts, data []byte, offset *big.Int) (bool, *big.Int, []byte, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ValueTester.contract.Call(opts, out, "bytestackToBytes", data, offset)
	return *ret0, *ret1, *ret2, err
}

// BytestackToBytes is a free data retrieval call binding the contract method 0xe4d476f4.
//
// Solidity: function bytestackToBytes(bytes data, uint256 offset) pure returns(bool, uint256, bytes)
func (_ValueTester *ValueTesterSession) BytestackToBytes(data []byte, offset *big.Int) (bool, *big.Int, []byte, error) {
	return _ValueTester.Contract.BytestackToBytes(&_ValueTester.CallOpts, data, offset)
}

// BytestackToBytes is a free data retrieval call binding the contract method 0xe4d476f4.
//
// Solidity: function bytestackToBytes(bytes data, uint256 offset) pure returns(bool, uint256, bytes)
func (_ValueTester *ValueTesterCallerSession) BytestackToBytes(data []byte, offset *big.Int) (bool, *big.Int, []byte, error) {
	return _ValueTester.Contract.BytestackToBytes(&_ValueTester.CallOpts, data, offset)
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(uint256, bytes32)
func (_ValueTester *ValueTesterCaller) DeserializeHash(opts *bind.CallOpts, data []byte, startOffset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ValueTester.contract.Call(opts, out, "deserializeHash", data, startOffset)
	return *ret0, *ret1, err
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
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "hashTestTuple")
	return *ret0, err
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
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "hashTuplePreImage", innerHash, valueSize)
	return *ret0, err
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
