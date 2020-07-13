// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package valuetester

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

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205130ee8069d4d45b671687a9bd9471ae9b0f29d9a9560a567b2a039c58a80e1a64736f6c63430005110032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820142891281d8f03ec873930437a8ca5e0c61de65eedd6218afc3087c76cc8fbbd64736f6c63430005110032"

// DeployValue deploys a new Ethereum contract, binding an instance of Value to it.
func DeployValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Value, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// Value is an auto generated Go binding around an Ethereum contract.
type Value struct {
	ValueCaller     // Read-only binding to the contract
	ValueTransactor // Write-only binding to the contract
	ValueFilterer   // Log filterer for contract events
}

// ValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueSession struct {
	Contract     *Value            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueCallerSession struct {
	Contract *ValueCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTransactorSession struct {
	Contract     *ValueTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueRaw struct {
	Contract *Value // Generic contract binding to access the raw methods on
}

// ValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueCallerRaw struct {
	Contract *ValueCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTransactorRaw struct {
	Contract *ValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValue creates a new instance of Value, bound to a specific deployed contract.
func NewValue(address common.Address, backend bind.ContractBackend) (*Value, error) {
	contract, err := bindValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// NewValueCaller creates a new read-only instance of Value, bound to a specific deployed contract.
func NewValueCaller(address common.Address, caller bind.ContractCaller) (*ValueCaller, error) {
	contract, err := bindValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueCaller{contract: contract}, nil
}

// NewValueTransactor creates a new write-only instance of Value, bound to a specific deployed contract.
func NewValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTransactor, error) {
	contract, err := bindValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTransactor{contract: contract}, nil
}

// NewValueFilterer creates a new log filterer instance of Value, bound to a specific deployed contract.
func NewValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueFilterer, error) {
	contract, err := bindValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueFilterer{contract: contract}, nil
}

// bindValue binds a generic wrapper to an already deployed contract.
func bindValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.ValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.contract.Transact(opts, method, params...)
}

// ValueTesterABI is the input ABI used to generate the binding from.
const ValueTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"}],\"name\":\"bytesToBytestackHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"bytestackToBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashTestTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"innerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueSize\",\"type\":\"uint256\"}],\"name\":\"hashTuplePreImage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueTesterFuncSigs maps the 4-byte function signature to its string representation.
var ValueTesterFuncSigs = map[string]string{
	"b325b7d0": "bytesToBytestackHash(bytes,uint256,uint256)",
	"e4d476f4": "bytestackToBytes(bytes,uint256)",
	"98206792": "deserializeHash(bytes,uint256)",
	"364df277": "hashEmptyTuple()",
	"fd5d0c8b": "hashTestTuple()",
	"c6d25c8e": "hashTuplePreImage(bytes32,uint256)",
}

// ValueTesterBin is the compiled bytecode used for deploying new contracts.
var ValueTesterBin = "0x608060405234801561001057600080fd5b50611728806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063364df277146100675780639820679214610081578063b325b7d014610149578063c6d25c8e146101f4578063e4d476f414610217578063fd5d0c8b14610349575b600080fd5b61006f610351565b60408051918252519081900360200190f35b6101296004803603604081101561009757600080fd5b8101906020810181356401000000008111156100b257600080fd5b8201836020820111156100c457600080fd5b803590602001918460018302840111640100000000831117156100e657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610360915050565b604080519315158452602084019290925282820152519081900360600190f35b61006f6004803603606081101561015f57600080fd5b81019060208101813564010000000081111561017a57600080fd5b82018360208201111561018c57600080fd5b803590602001918460018302840111640100000000831117156101ae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561039c565b61006f6004803603604081101561020a57600080fd5b50803590602001356103b9565b6102bf6004803603604081101561022d57600080fd5b81019060208101813564010000000081111561024857600080fd5b82018360208201111561025a57600080fd5b8035906020019184600183028401116401000000008311171561027c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506103cc915050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561030c5781810151838201526020016102f4565b50505050905090810190601f1680156103395780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61006f6103e8565b600061035b6104c1565b905090565b60008060008060006103706116a0565b61037a88886104e2565b925092509250828261038b836105ff565b9550955095505050505b9250925092565b60006103b16103ac858585610709565b6105ff565b949350505050565b60006103c5838361089a565b9392505050565b60008060606103db85856108d4565b9250925092509250925092565b60408051600280825260608281019093526000929190816020015b61040b6116a0565b81526020019060019003908161040357905050905061042a606f610b97565b8160008151811061043757fe5b6020026020010181905250610486600060405190808252806020026020018201604052801561048057816020015b61046d6116a0565b8152602001906001900390816104655790505b50610c49565b8160018151811061049357fe5b60200260200101819052506104a66116a0565b6104af82610c49565b90506104ba816105ff565b9250505090565b604080516000808252602082019092526104dc816001610d60565b91505090565b6000806104ed6116a0565b8451841061050d576000846105026000610b97565b925092509250610395565b600080859050600087828151811061052157fe5b016020015160019092019160f81c9050600081610563576105428984610d6f565b91955093509050838361055483610b97565b96509650965050505050610395565b60ff821660011415610579576105548984610dc2565b60ff82166002141561058f576105548984610eb2565b600360ff8316108015906105a65750600c60ff8316105b156105e057600219820160606105bd828c87610f57565b9197509550905085856105cf83610c49565b985098509850505050505050610395565b6000806105ed6000610b97565b91985096509450505050509250925092565b606081015160009060ff1661062057815161061990611015565b9050610704565b606082015160ff166001141561063d576106198260200151611039565b606082015160ff166002141561065e5781516080830151610619919061089a565b600360ff16826060015160ff161015801561068257506060820151600c60ff909116105b156106ab5761068f6116a0565b61069883611125565b90506106a3816105ff565b915050610704565b606082015160ff16606414156106c357508051610704565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b919050565b6107116116a0565b60208204600061071f6104c1565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b848110156107bd576107706103ac61076b8b602085028c0163ffffffff61125716565b610b97565b8260008151811061077d57fe5b602002602001018181525050838260018151811061079757fe5b6020026020010181815250506002830192506107b38284610d60565b9350600101610748565b50602086061561083f5760006107df89601f198a8a010163ffffffff61125716565b9050602087066020036008021b6107f86103ac82610b97565b8260008151811061080557fe5b602002602001018181525050838260018151811061081f57fe5b60200260200101818152505060028301925061083b8284610d60565b9350505b61084b6103ac87610b97565b8160008151811061085857fe5b602002602001018181525050828160018151811061087257fe5b60200260200101818152505060028201915061088e8183611273565b98975050505050505050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b600080606083915060008583815181106108ea57fe5b016020015160019093019260f81c90506005811461090c575060009250610395565b60006109188785611292565b91965094509050846109305750600093506103959050565b60208104601f8216600081610946576000610949565b60015b60ff168301905060608360405190808252806020026020018201604052801561097c578160200160208202803883390190505b5090506060836040519080825280601f01601f1916602001820160405280156109ac576020820181803883390190505b5090506000805b84811015610aaf578d8b815181106109c757fe5b01602001516001909b019a60f81c9850600589146109f2575060009a50610395975050505050505050565b60006109fe8f8d611292565b919e509c5090508c610a1e575060009b5061039598505050505050505050565b81158015610a2c5750600087115b15610a7f578060005b88811015610a7857818160208110610a4957fe5b1a60f81b868281518110610a5957fe5b60200101906001600160f81b031916908160001a905350600101610a35565b5050610aa6565b8060001b858460018b030381518110610a9457fe5b60209081029190910101526001909201915b506001016109b3565b508c8a81518110610abc57fe5b01602001516001909a019960f81c975060038814610ae65750600099506103959650505050505050565b60018a848460405160200180838051906020019060200280838360005b83811015610b1b578181015183820152602001610b03565b5050505090500182805190602001908083835b60208310610b4d5780518252601f199092019160209182019101610b2e565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529a509a509a5050505050505050509250925092565b610b9f6116a0565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610bf5565b610be26116a0565b815260200190600190039081610bda5790505b50905281526040805160008082526020828101909352919092019190610c31565b610c1e6116a0565b815260200190600190039081610c165790505b50815260006020820152600160409091015292915050565b610c516116a0565b610c5b82516112fb565b610cac576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610ce357838181518110610cc657fe5b602002602001015160800151820191508080600101915050610cb1565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610d3d565b610d2a6116a0565b815260200190600190039081610d225790505b509052815260208101859052935160030160ff1660408501526060909301525090565b60006103c56103ac8484611273565b6000806000808551905084811080610d8957506020858203105b15610d9e575060009250839150829050610395565b600160208601610db4888863ffffffff61125716565b935093509350509250925092565b600080610dcd6116a0565b60008490506000868281518110610de057fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110610e0657fe5b016020015160019093019260f81c9050610e1e6116a0565b8260ff1660011415610e5f576000610e368a866104e2565b9096509250905080610e5d57600089610e4d611302565b9750975097505050505050610395565b505b6000610e718a8663ffffffff61125716565b90506020850194508360ff1660011415610e9357600185610e4d8584866113b6565b600185610ea0858461143a565b97509750975050505050509250925092565b600080610ebd6116a0565b610ec56116a0565b8551600090819087811080610edc57506040888203105b15610ef4576000888596509650965050505050610395565b6000610f068a8a63ffffffff61125716565b9050602089019850610f188a8a610d6f565b909a50945092508215610f4357610f2f8185611495565b600198508997509550610395945050505050565b600089869750975097505050505050610395565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610fa257816020015b610f8f6116a0565b815260200190600190039081610f875790505b50905060005b8960ff168160ff161015610fff57610fc089856104e2565b8451859060ff8616908110610fd157fe5b60209081029190910101529450925082610ff7575060009550869450925061100c915050565b600101610fa8565b5060019550919350909150505b93509350939050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061104a57fe5b6040820151516110a45750805160208083015160408051600160f81b8185015260f89490941b6001600160f81b03191660218501526022808501929092528051808503909201825260429093019092528151910120610704565b600182600001516110cc84604001516000815181106110bf57fe5b60200260200101516105ff565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b61112d6116a0565b61113682611546565b61117c576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b600882604001515111156111ce576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608260400151516040519080825280602002602001820160405280156111ff578160200160208202803883390190505b50805190915060005b81811015611248576000611225866040015183815181106110bf57fe5b90508084838151811061123457fe5b602090810291909101015250600101611208565b506103b1828560800151611273565b6000816020018351101561126a57600080fd5b50016020015190565b61127b6116a0565b60006112868461155b565b90506103b18184611495565b60008060008085519050848110806112ac57506021858203105b806112ce5750600060ff168686815181106112c357fe5b016020015160f81c14155b156112e3575060009250839150829050610395565b600160218601610db48888840163ffffffff61125716565b6008101590565b61130a6116a0565b6040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190611363565b6113506116a0565b8152602001906001900390816113485790505b5090528152604080516000808252602082810190935291909201919061139f565b61138c6116a0565b8152602001906001900390816113845790505b508152600360208201526001604090910152905090565b6113be6116a0565b604080516001808252818301909252606091816020015b6113dd6116a0565b8152602001906001900390816113d5579050509050828160008151811061140057fe5b602002602001018190525061143160405180606001604052808760ff1681526020018681526020018381525061161b565b95945050505050565b6114426116a0565b6040805160608101825260ff8516815260208082018590528251600080825291810184526103c59383019161148d565b61147a6116a0565b8152602001906001900390816114725790505b50905261161b565b61149d6116a0565b6040805160a08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916114f3565b6114e06116a0565b8152602001906001900390816114d85790505b5090528152604080516000808252602082810190935291909201919061152f565b61151c6116a0565b8152602001906001900390816115145790505b508152600260208201526040019290925250919050565b60006115558260600151611682565b92915050565b60006008825111156115ab576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156115ef5781810151838201526020016115d7565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6116236116a0565b6040805160a081018252600080825260208083018690528351828152908101845291928301919061166a565b6116576116a0565b81526020019060019003908161164f5790505b50815260016020820181905260409091015292915050565b6000600c60ff8316108015611555575050600360ff91909116101590565b6040518060a00160405280600081526020016116ba6116d4565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fea265627a7a72315820f1e2c38080a66e39579bf0bfbfc9a720f344cf2303a7aa24ee74e607daeedcb164736f6c63430005110032"

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
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(bool, uint256, bytes32)
func (_ValueTester *ValueTesterCaller) DeserializeHash(opts *bind.CallOpts, data []byte, startOffset *big.Int) (bool, *big.Int, [32]byte, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ValueTester.contract.Call(opts, out, "deserializeHash", data, startOffset)
	return *ret0, *ret1, *ret2, err
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(bool, uint256, bytes32)
func (_ValueTester *ValueTesterSession) DeserializeHash(data []byte, startOffset *big.Int) (bool, *big.Int, [32]byte, error) {
	return _ValueTester.Contract.DeserializeHash(&_ValueTester.CallOpts, data, startOffset)
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(bool, uint256, bytes32)
func (_ValueTester *ValueTesterCallerSession) DeserializeHash(data []byte, startOffset *big.Int) (bool, *big.Int, [32]byte, error) {
	return _ValueTester.Contract.DeserializeHash(&_ValueTester.CallOpts, data, startOffset)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterSession) HashEmptyTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashEmptyTuple(&_ValueTester.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashEmptyTuple(&_ValueTester.CallOpts)
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
