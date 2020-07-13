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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582087e25d1b3581edfbcc93fab1b50976d754d7ebc01e825ad79f9dee7c91ee0bf164736f6c63430005110032"

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
var ValueTesterBin = "0x608060405234801561001057600080fd5b506116ac806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063364df277146100675780639820679214610081578063b325b7d014610149578063c6d25c8e146101f4578063e4d476f414610217578063fd5d0c8b14610349575b600080fd5b61006f610351565b60408051918252519081900360200190f35b6101296004803603604081101561009757600080fd5b8101906020810181356401000000008111156100b257600080fd5b8201836020820111156100c457600080fd5b803590602001918460018302840111640100000000831117156100e657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610360915050565b604080519315158452602084019290925282820152519081900360600190f35b61006f6004803603606081101561015f57600080fd5b81019060208101813564010000000081111561017a57600080fd5b82018360208201111561018c57600080fd5b803590602001918460018302840111640100000000831117156101ae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561039c565b61006f6004803603604081101561020a57600080fd5b50803590602001356103b9565b6102bf6004803603604081101561022d57600080fd5b81019060208101813564010000000081111561024857600080fd5b82018360208201111561025a57600080fd5b8035906020019184600183028401116401000000008311171561027c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506103cc915050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561030c5781810151838201526020016102f4565b50505050905090810190601f1680156103395780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61006f6103e8565b600061035b6104c1565b905090565b6000806000806000610370611615565b61037a88886104e2565b925092509250828261038b836105ff565b9550955095505050505b9250925092565b60006103b16103ac8585856106fc565b6105ff565b949350505050565b60006103c5838361088d565b9392505050565b60008060606103db85856108c7565b9250925092509250925092565b60408051600280825260608281019093526000929190816020015b61040b611615565b81526020019060019003908161040357905050905061042a606f610b8a565b8160008151811061043757fe5b6020026020010181905250610486600060405190808252806020026020018201604052801561048057816020015b61046d611615565b8152602001906001900390816104655790505b50610c16565b8160018151811061049357fe5b60200260200101819052506104a6611615565b6104af82610c16565b90506104ba816105ff565b9250505090565b604080516000808252602082019092526104dc816001610d0d565b91505090565b6000806104ed611615565b8451841061050d576000846105026000610b8a565b925092509250610395565b600080859050600087828151811061052157fe5b016020015160019092019160f81c9050600081610563576105428984610d2c565b91955093509050838361055483610b8a565b96509650965050505050610395565b60ff821660011415610579576105548984610d7f565b60ff82166002141561058f576105548984610e6f565b600360ff8316108015906105a65750600c60ff8316105b156105e057600219820160606105bd828c87610f14565b9197509550905085856105cf83610c16565b985098509850505050505050610395565b6000806105ed6000610b8a565b91985096509450505050509250925092565b606081015160009060ff1661062057815161061990610fd2565b90506106f7565b606082015160ff166001141561065357602080830151805160408201516060830151929093015161061993919290610ff6565b606082015160ff166002141561066c576106198261109e565b600360ff16826060015160ff161015801561069057506060820151600c60ff909116105b1561069e576106198261110a565b606082015160ff16606414156106b6575080516106f7565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b919050565b610704611615565b6020820460006107126104c1565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b848110156107b0576107636103ac61075e8b602085028c0163ffffffff61112816565b610b8a565b8260008151811061077057fe5b602002602001018181525050838260018151811061078a57fe5b6020026020010181815250506002830192506107a68284610d0d565b935060010161073b565b5060208606156108325760006107d289601f198a8a010163ffffffff61112816565b9050602087066020036008021b6107eb6103ac82610b8a565b826000815181106107f857fe5b602002602001018181525050838260018151811061081257fe5b60200260200101818152505060028301925061082e8284610d0d565b9350505b61083e6103ac87610b8a565b8160008151811061084b57fe5b602002602001018181525050828160018151811061086557fe5b6020026020010181815250506002820191506108818183611144565b98975050505050505050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b600080606083915060008583815181106108dd57fe5b016020015160019093019260f81c9050600581146108ff575060009250610395565b600061090b8785611163565b91965094509050846109235750600093506103959050565b60208104601f821660008161093957600061093c565b60015b60ff168301905060608360405190808252806020026020018201604052801561096f578160200160208202803883390190505b5090506060836040519080825280601f01601f19166020018201604052801561099f576020820181803883390190505b5090506000805b84811015610aa2578d8b815181106109ba57fe5b01602001516001909b019a60f81c9850600589146109e5575060009a50610395975050505050505050565b60006109f18f8d611163565b919e509c5090508c610a11575060009b5061039598505050505050505050565b81158015610a1f5750600087115b15610a72578060005b88811015610a6b57818160208110610a3c57fe5b1a60f81b868281518110610a4c57fe5b60200101906001600160f81b031916908160001a905350600101610a28565b5050610a99565b8060001b858460018b030381518110610a8757fe5b60209081029190910101526001909201915b506001016109a6565b508c8a81518110610aaf57fe5b01602001516001909a019960f81c975060038814610ad95750600099506103959650505050505050565b60018a848460405160200180838051906020019060200280838360005b83811015610b0e578181015183820152602001610af6565b5050505090500182805190602001908083835b60208310610b405780518252601f199092019160209182019101610b21565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529a509a509a5050505050505050509250925092565b610b92611615565b6040805160a080820183528482528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191610bfe565b610beb611615565b815260200190600190039081610be35790505b50815260006020820152600160409091015292915050565b610c1e611615565b610c2882516111cc565b610c79576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610cb057838181518110610c9357fe5b602002602001015160800151820191508080600101915050610c7e565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b6000610d17611615565b610d218484611144565b90506103b18161109e565b6000806000808551905084811080610d4657506020858203105b15610d5b575060009250839150829050610395565b600160208601610d71888863ffffffff61112816565b935093509350509250925092565b600080610d8a611615565b60008490506000868281518110610d9d57fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110610dc357fe5b016020015160019093019260f81c9050610ddb611615565b8260ff1660011415610e1c576000610df38a866104e2565b9096509250905080610e1a57600089610e0a6111d3565b9750975097505050505050610395565b505b6000610e2e8a8663ffffffff61112816565b90506020850194508360ff1660011415610e5057600185610e0a85848661125e565b600185610e5d85846112a7565b97509750975050505050509250925092565b600080610e7a611615565b610e82611615565b8551600090819087811080610e9957506040888203105b15610eb1576000888596509650965050505050610395565b6000610ec38a8a63ffffffff61112816565b9050602089019850610ed58a8a610d2c565b909a50945092508215610f0057610eec81856112e8565b600198508997509550610395945050505050565b600089869750975097505050505050610395565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610f5f57816020015b610f4c611615565b815260200190600190039081610f445790505b50905060005b8960ff168160ff161015610fbc57610f7d89856104e2565b8451859060ff8616908110610f8e57fe5b60209081029190910101529450925082610fb45750600095508694509250610fc9915050565b600101610f65565b5060019550919350909150505b93509350939050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611050575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206103b1565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff166002146110f3576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b81516080830151611104919061088d565b92915050565b6000611114611615565b61111d83611373565b90506103c58161109e565b6000816020018351101561113b57600080fd5b50016020015190565b61114c611615565b6000611157846113e9565b90506103b181846112e8565b600080600080855190508481108061117d57506021858203105b8061119f5750600060ff1686868151811061119457fe5b016020015160f81c14155b156111b4575060009250839150829050610395565b600160218601610d718888840163ffffffff61112816565b6008101590565b6111db611615565b6040805160a080820183526000808352835191820184528082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191611247565b611234611615565b81526020019060019003908161122c5790505b508152600360208201526001604090910152905090565b611266611615565b6103b16040518060a001604052808660ff168152602001858152602001600115158152602001611295856105ff565b815260200184608001518152506114a9565b6112af611615565b6103c56040518060a001604052808560ff1681526020018481526020016000151581526020016000801b815260200160008152506114a9565b6112f0611615565b6040805160a08082018352858252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161135c565b611349611615565b8152602001906001900390816113415790505b508152600260208201526040019290925250919050565b61137b611615565b61138482611510565b6113ca576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60606113d9836040015161151f565b90506103c5818460800151611144565b6000600882511115611439576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561147d578181015183820152602001611465565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6114b1611615565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906114f8565b6114e5611615565b8152602001906001900390816114dd5790505b50815260016020820181905260409091015292915050565b600061110482606001516115f7565b606060088251111561156f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825160405190808252806020026020018201604052801561159c578160200160208202803883390190505b50805190915060005b818110156115ee5760006115cb8683815181106115be57fe5b60200260200101516105ff565b9050808483815181106115da57fe5b6020908102919091010152506001016115a5565b50909392505050565b6000600c60ff8316108015611104575050600360ff91909116101590565b6040518060a001604052806000815260200161162f611649565b815260606020820181905260006040830181905291015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a7231582008c11069191c78fceeedf11a898cafcf45909c857acf2e3c5cd225550e96e33964736f6c63430005110032"

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
