// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package messagetester

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204a4406891b86a3dacfb6ea735b92a605972908aec5c506cc187bb6f01ef6a24064736f6c634300050f0032"

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

// MessageTesterABI is the input ABI used to generate the binding from.
const MessageTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc20Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc20MessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc721Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc721MessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"ethHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"ethMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"transactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"transactionMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MessageTesterFuncSigs maps the 4-byte function signature to its string representation.
var MessageTesterFuncSigs = map[string]string{
	"79cf8c0d": "erc20Hash(address,address,address,uint256,uint256,uint256)",
	"8346d164": "erc20MessageHash(address,address,address,uint256,uint256,uint256)",
	"cfe44216": "erc721Hash(address,address,address,uint256,uint256,uint256)",
	"af015670": "erc721MessageHash(address,address,address,uint256,uint256,uint256)",
	"fcbf9131": "ethHash(address,address,uint256,uint256,uint256)",
	"74faf594": "ethMessageHash(address,address,uint256,uint256,uint256)",
	"e8151c85": "transactionHash(address,address,address,uint256,uint256,bytes,uint256)",
	"31be08f9": "transactionMessageHash(address,address,address,uint256,uint256,bytes,uint256)",
}

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
var MessageTesterBin = "0x608060405234801561001057600080fd5b506117a8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063af0156701161005b578063af01567014610246578063cfe442161461028e578063e8151c85146102d6578063fcbf9131146103ab57610088565b806331be08f91461008d57806374faf5941461017457806379cf8c0d146101b65780638346d164146101fe575b600080fd5b610162600480360360e08110156100a357600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160808101359181019060c0810160a08201356401000000008111156100eb57600080fd5b8201836020820111156100fd57600080fd5b8035906020019184600183028401116401000000008311171561011f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506103ed915050565b60408051918252519081900360200190f35b610162600480360360a081101561018a57600080fd5b506001600160a01b0381358116916020810135909116906040810135906060810135906080013561040a565b610162600480360360c08110156101cc57600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a00135610423565b610162600480360360c081101561021457600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a0013561043e565b610162600480360360c081101561025c57600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a00135610462565b610162600480360360c08110156102a457600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a00135610475565b610162600480360360e08110156102ec57600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160808101359181019060c0810160a082013564010000000081111561033457600080fd5b82018360208201111561034657600080fd5b8035906020019184600183028401116401000000008311171561036857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610485915050565b610162600480360360a08110156103c157600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060800135610496565b60006103fe888888888888886104a5565b98975050505050505050565b600061041986868686866104d6565b9695505050505050565b60006104338787878787876104f8565b979650505050505050565b600061045661045188888888888861050a565b610522565b51979650505050505050565b6000610456610451888888888888610659565b6000610433878787878787610671565b60006103fe88888888888888610683565b60006104198686868686610778565b60006104af611706565b6104be898989898989896107e3565b90506104c981610af5565b9998505050505050505050565b60006104e0611706565b6104ed8787878787610cac565b905061043381610af5565b60006104336002888888888888610e75565b610512611706565b6104336002888888888888610ef3565b61052a61173a565b6060820151600c60ff9091161061057c576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166105a95760405180602001604052806105a084600001516110f9565b90529050610654565b606082015160ff16600114156105f05760405180602001604052806105a084602001516000015185602001516040015186602001516060015187602001516020015161111d565b606082015160ff16600214156106155750604080516020810190915281518152610654565b600360ff16826060015160ff161015801561063957506060820151600c60ff909116105b156106525760405180602001604052806105a084610af5565bfe5b919050565b610661611706565b6104336003888888888888610ef3565b60006104336003888888888888610e75565b60008088888888888888604051602001808960ff1660ff1660f81b8152600101886001600160a01b03166001600160a01b031660601b8152601401876001600160a01b03166001600160a01b031660601b8152601401866001600160a01b03166001600160a01b031660601b815260140185815260200184815260200183805190602001908083835b6020831061072b5780518252601f19909201916020918201910161070c565b51815160001960209485036101000a019081169019919091161790529201938452506040805180850381529382019052825192019190912098505050505050505050979650505050505050565b60408051600160f81b6020808301919091526bffffffffffffffffffffffff19606089811b8216602185015288901b166035830152604982018690526069820185905260898083018590528351808403909101815260a9909201909252805191012095945050505050565b6107eb611706565b600080898989898989604051602001808860ff1660ff1660f81b8152600101876001600160a01b03166001600160a01b031660601b8152601401866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b815260140184815260200183815260200182805190602001908083835b602083106108925780518252601f199092019160209182019101610873565b6001836020036101000a0380198251168184511680821785525050505050509050019750505050505050506040516020818303038152906040528051906020012090506060600460405190808252806020026020018201604052801561091257816020015b6108ff611706565b8152602001906001900390816108f75790505b509050610927896001600160a01b03166111c6565b8160008151811061093457fe5b6020026020010181905250610948876111c6565b8160018151811061095557fe5b6020026020010181905250610969866111c6565b8160028151811061097657fe5b602002602001018190525061098a8561124b565b8160038151811061099757fe5b602090810291909101015260408051600380825260808201909252606091816020015b6109c2611706565b8152602001906001900390816109ba5790505090506109e160006111c6565b816000815181106109ee57fe5b6020026020010181905250610a0b896001600160a01b03166111c6565b81600181518110610a1857fe5b6020026020010181905250610a2c826113f2565b81600281518110610a3957fe5b602090810291909101015260408051600380825260808201909252606091816020015b610a64611706565b815260200190600190039081610a5c579050509050610a82866111c6565b81600081518110610a8f57fe5b6020908102919091010152610aa3846111c6565b81600181518110610ab057fe5b6020026020010181905250610ac4826113f2565b81600281518110610ad157fe5b6020026020010181905250610ae5816113f2565b9c9b505050505050505050505050565b6000610b00826114e1565b610b46576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60088260400151511115610b98576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060826040015151604051908082528060200260200182016040528015610bc9578160200160208202803883390190505b50805190915060005b81811015610c2957610be261173a565b610c0286604001518381518110610bf557fe5b6020026020010151610522565b90508060000151848381518110610c1557fe5b602090810291909101015250600101610bd2565b50836040015151600360ff1601846080015183604051602001808460ff1660ff1660f81b8152600101838152602001828051906020019060200280838360005b83811015610c81578181015183820152602001610c69565b5050505090500193505050506040516020818303038152906040528051906020012092505050919050565b610cb4611706565b6040805160028082526060828101909352816020015b610cd2611706565b815260200190600190039081610cca579050509050610cf9876001600160a01b03166111c6565b81600081518110610d0657fe5b6020026020010181905250610d1a856111c6565b81600181518110610d2757fe5b602090810291909101015260408051600380825260808201909252606091816020015b610d52611706565b815260200190600190039081610d4a579050509050610d7160016111c6565b81600081518110610d7e57fe5b6020026020010181905250610d9b876001600160a01b03166111c6565b81600181518110610da857fe5b6020026020010181905250610dbc826113f2565b81600281518110610dc957fe5b602090810291909101015260408051600380825260808201909252606091816020015b610df4611706565b815260200190600190039081610dec579050509050610e12866111c6565b81600081518110610e1f57fe5b6020026020010181905250610e33856111c6565b81600181518110610e4057fe5b6020026020010181905250610e54826113f2565b81600281518110610e6157fe5b60200260200101819052506104c9816113f2565b6040805160f89890981b6001600160f81b0319166020808a0191909152606097881b6bffffffffffffffffffffffff1990811660218b015296881b871660358a01529490961b9094166049870152605d860191909152607d850152609d808501929092528251808503909201825260bd909301909152805191012090565b610efb611706565b60408051600380825260808201909252606091816020015b610f1b611706565b815260200190600190039081610f13579050509050610f42866001600160a01b03166111c6565b81600081518110610f4f57fe5b6020026020010181905250610f6c886001600160a01b03166111c6565b81600181518110610f7957fe5b6020026020010181905250610f8d856111c6565b81600281518110610f9a57fe5b602090810291909101015260408051600380825260808201909252606091816020015b610fc5611706565b815260200190600190039081610fbd579050509050610fe68a60ff166111c6565b81600081518110610ff357fe5b6020026020010181905250611010886001600160a01b03166111c6565b8160018151811061101d57fe5b6020026020010181905250611031826113f2565b8160028151811061103e57fe5b602090810291909101015260408051600380825260808201909252606091816020015b611069611706565b815260200190600190039081611061579050509050611087866111c6565b8160008151811061109457fe5b60200260200101819052506110a8856111c6565b816001815181106110b557fe5b60200260200101819052506110c9826113f2565b816002815181106110d657fe5b60200260200101819052506110ea816113f2565b9b9a5050505050505050505050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611177575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206111be565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6111ce611706565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611233565b611220611706565b8152602001906001900390816112185790505b50815260006020820152600160409091015292915050565b611253611706565b8151602080820490601f830104600061126a6114f6565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b8581101561130d5783826000815181106112a957fe5b6020026020010181815250506112d66104516112d1836020028c61157890919063ffffffff16565b6111c6565b60000151826001815181106112e757fe5b6020026020010181815250506002830192506113038284611594565b9350600101611293565b508385101561139657600061132c89601f19890163ffffffff61157816565b905085602002876020030360080281901b9050838260008151811061134d57fe5b602002602001018181525050611365610451826111c6565b600001518260018151811061137657fe5b6020026020010181815250506002830192506113928284611594565b9350505b6113a2610451876111c6565b60000151816000815181106113b357fe5b60200260200101818152505082816001815181106113cd57fe5b6020026020010181815250506002820191506103fe6113ec8284611594565b8361165d565b6113fa611706565b61140482516116e1565b611455576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561148c5783818151811061146f57fe5b60200260200101516080015182019150808060010191505061145a565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b60006114f082606001516116e8565b92915050565b604080516000808252602082019092526003600182604051602001808460ff1660ff1660f81b8152600101838152602001828051906020019060200280838360005b83811015611550578181015183820152602001611538565b5050505090500193505050506040516020818303038152906040528051906020012091505090565b6000816020018351101561158b57600080fd5b50016020015190565b60006008835111156115e4576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8251600360ff16018284604051602001808460ff1660ff1660f81b8152600101838152602001828051906020019060200280838360005b8381101561163357818101518382015260200161161b565b50505050905001935050505060405160208183030381529060405280519060200120905092915050565b611665611706565b6040805160a08101825284815281516080810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916116ca565b6116b7611706565b8152602001906001900390816116af5790505b508152600260208201526040019290925250919050565b6008101590565b6000600c60ff83161080156114f0575050600360ff91909116101590565b6040518060a001604052806000815260200161172061174c565b815260606020820181905260006040830181905291015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a72315820b436a46244e6c4ea39387cc838a116e496a1c97f639ad7757b52156be8d0e35f64736f6c634300050f0032"

// DeployMessageTester deploys a new Ethereum contract, binding an instance of MessageTester to it.
func DeployMessageTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageTester, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessageTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// MessageTester is an auto generated Go binding around an Ethereum contract.
type MessageTester struct {
	MessageTesterCaller     // Read-only binding to the contract
	MessageTesterTransactor // Write-only binding to the contract
	MessageTesterFilterer   // Log filterer for contract events
}

// MessageTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTesterSession struct {
	Contract     *MessageTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTesterCallerSession struct {
	Contract *MessageTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessageTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTesterTransactorSession struct {
	Contract     *MessageTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessageTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTesterRaw struct {
	Contract *MessageTester // Generic contract binding to access the raw methods on
}

// MessageTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTesterCallerRaw struct {
	Contract *MessageTesterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTesterTransactorRaw struct {
	Contract *MessageTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTester creates a new instance of MessageTester, bound to a specific deployed contract.
func NewMessageTester(address common.Address, backend bind.ContractBackend) (*MessageTester, error) {
	contract, err := bindMessageTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// NewMessageTesterCaller creates a new read-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterCaller(address common.Address, caller bind.ContractCaller) (*MessageTesterCaller, error) {
	contract, err := bindMessageTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterCaller{contract: contract}, nil
}

// NewMessageTesterTransactor creates a new write-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTesterTransactor, error) {
	contract, err := bindMessageTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterTransactor{contract: contract}, nil
}

// NewMessageTesterFilterer creates a new log filterer instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTesterFilterer, error) {
	contract, err := bindMessageTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTesterFilterer{contract: contract}, nil
}

// bindMessageTester binds a generic wrapper to an already deployed contract.
func bindMessageTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.MessageTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transact(opts, method, params...)
}

// Erc20Hash is a free data retrieval call binding the contract method 0x79cf8c0d.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc20Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc20Hash", to, from, erc20, value, blockNumber, messageNum)
	return *ret0, err
}

// Erc20Hash is a free data retrieval call binding the contract method 0x79cf8c0d.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc20Hash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20Hash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, messageNum)
}

// Erc20Hash is a free data retrieval call binding the contract method 0x79cf8c0d.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc20Hash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20Hash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, messageNum)
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0x8346d164.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc20MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc20MessageHash", to, from, erc20, value, blockNumber, messageNum)
	return *ret0, err
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0x8346d164.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc20MessageHash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20MessageHash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, messageNum)
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0x8346d164.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc20MessageHash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20MessageHash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, messageNum)
}

// Erc721Hash is a free data retrieval call binding the contract method 0xcfe44216.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc721Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc721Hash", to, from, erc721, id, blockNumber, messageNum)
	return *ret0, err
}

// Erc721Hash is a free data retrieval call binding the contract method 0xcfe44216.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc721Hash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721Hash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, messageNum)
}

// Erc721Hash is a free data retrieval call binding the contract method 0xcfe44216.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc721Hash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721Hash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, messageNum)
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xaf015670.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc721MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc721MessageHash", to, from, erc721, id, blockNumber, messageNum)
	return *ret0, err
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xaf015670.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc721MessageHash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721MessageHash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, messageNum)
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xaf015670.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc721MessageHash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721MessageHash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, messageNum)
}

// EthHash is a free data retrieval call binding the contract method 0xfcbf9131.
//
// Solidity: function ethHash(address to, address from, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) EthHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "ethHash", to, from, value, blockNumber, messageNum)
	return *ret0, err
}

// EthHash is a free data retrieval call binding the contract method 0xfcbf9131.
//
// Solidity: function ethHash(address to, address from, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) EthHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthHash(&_MessageTester.CallOpts, to, from, value, blockNumber, messageNum)
}

// EthHash is a free data retrieval call binding the contract method 0xfcbf9131.
//
// Solidity: function ethHash(address to, address from, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) EthHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthHash(&_MessageTester.CallOpts, to, from, value, blockNumber, messageNum)
}

// EthMessageHash is a free data retrieval call binding the contract method 0x74faf594.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) EthMessageHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "ethMessageHash", to, from, value, blockNumber, messageNum)
	return *ret0, err
}

// EthMessageHash is a free data retrieval call binding the contract method 0x74faf594.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) EthMessageHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthMessageHash(&_MessageTester.CallOpts, to, from, value, blockNumber, messageNum)
}

// EthMessageHash is a free data retrieval call binding the contract method 0x74faf594.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value, uint256 blockNumber, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) EthMessageHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthMessageHash(&_MessageTester.CallOpts, to, from, value, blockNumber, messageNum)
}

// TransactionHash is a free data retrieval call binding the contract method 0xe8151c85.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionHash", chain, to, from, seqNumber, value, data, blockNumber)
	return *ret0, err
}

// TransactionHash is a free data retrieval call binding the contract method 0xe8151c85.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber)
}

// TransactionHash is a free data retrieval call binding the contract method 0xe8151c85.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber)
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x31be08f9.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionMessageHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionMessageHash", chain, to, from, seqNumber, value, data, blockNumber)
	return *ret0, err
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x31be08f9.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionMessageHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber)
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x31be08f9.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber)
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209a88a80f7dbf6eb3c7eddc52926976ff2a84359ecda25eaacdd17bb51b0da0b264736f6c634300050f0032"

// DeployMessages deploys a new Ethereum contract, binding an instance of Messages to it.
func DeployMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messages, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// Messages is an auto generated Go binding around an Ethereum contract.
type Messages struct {
	MessagesCaller     // Read-only binding to the contract
	MessagesTransactor // Write-only binding to the contract
	MessagesFilterer   // Log filterer for contract events
}

// MessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagesSession struct {
	Contract     *Messages         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagesCallerSession struct {
	Contract *MessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagesTransactorSession struct {
	Contract     *MessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagesRaw struct {
	Contract *Messages // Generic contract binding to access the raw methods on
}

// MessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagesCallerRaw struct {
	Contract *MessagesCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagesTransactorRaw struct {
	Contract *MessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessages creates a new instance of Messages, bound to a specific deployed contract.
func NewMessages(address common.Address, backend bind.ContractBackend) (*Messages, error) {
	contract, err := bindMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// NewMessagesCaller creates a new read-only instance of Messages, bound to a specific deployed contract.
func NewMessagesCaller(address common.Address, caller bind.ContractCaller) (*MessagesCaller, error) {
	contract, err := bindMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesCaller{contract: contract}, nil
}

// NewMessagesTransactor creates a new write-only instance of Messages, bound to a specific deployed contract.
func NewMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesTransactor, error) {
	contract, err := bindMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesTransactor{contract: contract}, nil
}

// NewMessagesFilterer creates a new log filterer instance of Messages, bound to a specific deployed contract.
func NewMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesFilterer, error) {
	contract, err := bindMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesFilterer{contract: contract}, nil
}

// bindMessages binds a generic wrapper to an already deployed contract.
func bindMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.MessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820071748e7c73e5f65e1b1ab03af6e0b53935e549fc73ab48434e59fa02406bba364736f6c634300050f0032"

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
