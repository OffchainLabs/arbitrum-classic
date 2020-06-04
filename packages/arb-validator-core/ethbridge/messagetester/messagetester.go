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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820557675814df39add0fa761bb6fc1ca2754618e3660ac97035c79af694b750a9264736f6c634300050d0032"

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
const MessageTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc20Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc20MessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc721Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"erc721MessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"ethHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"ethMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"transactionBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"transactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"prev\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prevSize\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"transactionMessageBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"transactionMessageBatchHashSingle\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"transactionMessageBatchSingleSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"transactionMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MessageTesterFuncSigs maps the 4-byte function signature to its string representation.
var MessageTesterFuncSigs = map[string]string{
	"f3e5743e": "erc20Hash(address,address,address,uint256,uint256,uint256,uint256)",
	"fadffc5f": "erc20MessageHash(address,address,address,uint256,uint256,uint256,uint256)",
	"87c406bf": "erc721Hash(address,address,address,uint256,uint256,uint256,uint256)",
	"dd182d69": "erc721MessageHash(address,address,address,uint256,uint256,uint256,uint256)",
	"bbfd47ce": "ethHash(address,address,uint256,uint256,uint256,uint256)",
	"3bcceb7d": "ethMessageHash(address,address,uint256,uint256,uint256,uint256)",
	"ca06a372": "transactionBatchHash(bytes,uint256,uint256,uint256)",
	"d9214a35": "transactionHash(address,address,address,uint256,uint256,bytes,uint256,uint256,uint256)",
	"a109e062": "transactionMessageBatchHash(bytes32,uint256,address,bytes,uint256,uint256)",
	"94caf578": "transactionMessageBatchHashSingle(uint256,address,bytes,uint256,uint256)",
	"fedc217e": "transactionMessageBatchSingleSender(uint256,address,bytes32,bytes)",
	"63bc3d74": "transactionMessageHash(address,address,address,uint256,uint256,bytes,uint256,uint256)",
}

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
var MessageTesterBin = "0x608060405234801561001057600080fd5b50612026806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063ca06a37211610071578063ca06a37214610403578063d9214a35146104b2578063dd182d691461058f578063f3e5743e146105dd578063fadffc5f1461062b578063fedc217e14610679576100b4565b80633bcceb7d146100b957806363bc3d741461011357806387c406bf146101ea57806394caf57814610238578063a109e062146102f6578063bbfd47ce146103bb575b600080fd5b610101600480360360c08110156100cf57600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060808101359060a00135610755565b60408051918252519081900360200190f35b610101600480360361010081101561012a57600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160808101359181019060c0810160a0820135600160201b81111561017157600080fd5b82018360208201111561018357600080fd5b803590602001918460018302840111600160201b831117156101a457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610772565b610101600480360360e081101561020057600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a08101359060c00135610791565b610101600480360360a081101561024e57600080fd5b8135916001600160a01b0360208201351691810190606081016040820135600160201b81111561027d57600080fd5b82018360208201111561028f57600080fd5b803590602001918460018302840111600160201b831117156102b057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356107ae565b610101600480360360c081101561030c57600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b81111561034257600080fd5b82018360208201111561035457600080fd5b803590602001918460018302840111600160201b8311171561037557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356107c5565b610101600480360360c08110156103d157600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060808101359060a001356107d5565b6101016004803603608081101561041957600080fd5b810190602081018135600160201b81111561043357600080fd5b82018360208201111561044557600080fd5b803590602001918460018302840111600160201b8311171561046657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604001356107e5565b61010160048036036101208110156104c957600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160808101359181019060c0810160a0820135600160201b81111561051057600080fd5b82018360208201111561052257600080fd5b803590602001918460018302840111600160201b8311171561054357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604001356107fe565b610101600480360360e08110156105a557600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a08101359060c00135610826565b610101600480360360e08110156105f357600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a08101359060c0013561083a565b610101600480360360e081101561064157600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a08101359060c0013561084b565b6107396004803603608081101561068f57600080fd5b8135916001600160a01b036020820135169160408201359190810190608081016060820135600160201b8111156106c557600080fd5b8201836020820111156106d757600080fd5b803590602001918460018302840111600160201b831117156106f857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061085f945050505050565b604080516001600160a01b039092168252519081900360200190f35b600061076587878787878761086d565b90505b9695505050505050565b60006107848989898989898989610890565b9998505050505050505050565b60006107a2888888888888886108c9565b98975050505050505050565b60006107686107c087878787876108dc565b6109d7565b6000610765878787878787610aea565b6000610765878787878787610b84565b60006107f385858585610be2565b90505b949350505050565b60006108188a8a8a8a8a8a805190602001208a8a8a610c6f565b9a9950505050505050505050565b60006107a26107c089898989898989610ce0565b60006107a288888888888888610cf9565b60006107a26107c089898989898989610d0c565b60006107f385858585610d25565b6000610877611f96565b610885888888888888610ddb565b90506107a281610fc5565b600061089a611f96565b6108be8a8a8a8a8a8a805190602001206108b78c60008e51610fea565b8b8b61117e565b905061081881610fc5565b60006107a2600389898989898989611418565b6108e4611f96565b600061090785609789016108fe828b63ffffffff61148a16565b61ffff166114a6565b9050600061091788888489610d25565b90506001600160a01b038116610962576040805162461bcd60e51b815260206004820152600b60248201526a696e76616c69642073696760a81b604482015290519081900360640190fd5b61096a611f96565b61098b8760978b01610982828d63ffffffff61148a16565b61ffff16610fea565b9050610784886109a48960028d0163ffffffff6114af16565b846109b88b60168f0163ffffffff6114d216565b6109ce60368f018d6114d290919063ffffffff16565b88878d8d61117e565b6000600360090160ff16826060015160ff1610610a2f576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610a4d578151610a46906114ee565b9050610ae5565b606082015160ff1660011415610a80576020808301518051604082015160608301519290930151610a4693919290611512565b606082015160ff1660021415610ab157602080830151015160011415610aa857508051610ae5565b610a46826115ba565b600360ff16826060015160ff1610158015610ad557506060820151600c60ff909116105b15610ae357610a4682610fc5565bfe5b919050565b6000610af4611f96565b610afe8888611626565b855190915060005b81609782011015610b7b576000610b23888363ffffffff61148a16565b9050828161ffff1660978401011115610b4a57610b3f846109d7565b945050505050610768565b610b52611f96565b610b5f838b8b8b8b6108dc565b9050610b6b85826116aa565b94505061ffff1601609701610b06565b610818836109d7565b60408051600160f81b6020808301919091526001600160601b031960608a811b8216602185015289901b1660358301526049808301889052835180840390910181526069909201909252805191012060009061076590858585611720565b60006107f3600686604051602001808360ff1660ff1660f81b815260010182805190602001908083835b60208310610c2b5780518252601f199092019160209182019101610c0c565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405280519060200120858585611720565b60408051600060208083018290526001600160601b031960608e811b821660218601528d811b821660358601528c901b166049840152605d83018a9052607d8301899052609d8084018990528451808503909101815260bd9093019093528151919092012061081890858585611720565b610ce8611f96565b6107a260038989898989898961175e565b60006107a2600289898989898989611418565b610d14611f96565b6107a260028989898989898961175e565b60006107f384610d3e846002890163ffffffff6114af16565b610d518560168a0163ffffffff6114d216565b610d648660368b0163ffffffff6114d216565b8760405160200180866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b815260140184815260200183815260200182815260200195505050505050604051602081830303815290604052805190602001208360568801611986565b610de3611f96565b6040805160028082526060828101909352816020015b610e01611f96565b815260200190600190039081610df9579050509050610e28886001600160a01b0316611ab9565b81600081518110610e3557fe5b6020026020010181905250610e4986611ab9565b81600181518110610e5657fe5b602090810291909101015260408051600380825260808201909252606091816020015b610e81611f96565b815260200190600190039081610e79579050509050610ea06001611ab9565b81600081518110610ead57fe5b6020026020010181905250610eca886001600160a01b0316611ab9565b81600181518110610ed757fe5b6020026020010181905250610eeb82611b3e565b81600281518110610ef857fe5b602090810291909101015260408051600480825260a08201909252606091816020015b610f23611f96565b815260200190600190039081610f1b579050509050610f4187611ab9565b81600081518110610f4e57fe5b6020026020010181905250610f6286611ab9565b81600181518110610f6f57fe5b6020026020010181905250610f8385611ab9565b81600281518110610f9057fe5b6020026020010181905250610fa482611b3e565b81600381518110610fb157fe5b602002602001018190525061081881611b3e565b6000610fcf611f96565b610fd883611c2d565b9050610fe3816115ba565b9392505050565b610ff2611f96565b602080830490601f8401046000611007611ca3565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b858110156110a857838260008151811061104657fe5b6020026020010181815250506110756107c0611070836020028c018d6114d290919063ffffffff16565b611ab9565b8260018151811061108257fe5b60200260200101818152505060028301925061109e8284611cc4565b9350600101611030565b508385101561112f5760006110c98a601f198b8b010163ffffffff6114d216565b905085602002886020030360080281901b905083826000815181106110ea57fe5b6020026020010181815250506111026107c082611ab9565b8260018151811061110f57fe5b60200260200101818152505060028301925061112b8284611cc4565b9350505b61113b6107c088611ab9565b8160008151811061114857fe5b602002602001018181525050828160018151811061116257fe5b6020026020010181815250506002820191506107848183611ce3565b611186611f96565b60408051600060208083019190915260608d811b6001600160601b031990811660218501528d821b811660358501528c821b166049840152605d83018b9052607d83018a9052609d8084018a90528451808503909101815260bd840180865281519190930120600480845261015d850190955293909260dd015b611208611f96565b81526020019060019003908161120057905050905061122f8b6001600160a01b0316611ab9565b8160008151811061123c57fe5b602002602001018190525061125089611ab9565b8160018151811061125d57fe5b602002602001018190525061127188611ab9565b8160028151811061127e57fe5b6020026020010181905250858160038151811061129757fe5b602090810291909101015260408051600380825260808201909252606091816020015b6112c2611f96565b8152602001906001900390816112ba5790505090506112e16000611ab9565b816000815181106112ee57fe5b602002602001018190525061130b8b6001600160a01b0316611ab9565b8160018151811061131857fe5b602002602001018190525061132c82611b3e565b8160028151811061133957fe5b602090810291909101015260408051600480825260a08201909252606091816020015b611364611f96565b81526020019060019003908161135c57905050905061138287611ab9565b8160008151811061138f57fe5b60200260200101819052506113a386611ab9565b816001815181106113b057fe5b60209081029190910101526113c484611ab9565b816002815181106113d157fe5b60200260200101819052506113e582611b3e565b816003815181106113f257fe5b602002602001018190525061140681611b3e565b9e9d5050505050505050505050505050565b604080516001600160f81b031960f88b901b166020808301919091526001600160601b031960608b811b821660218501528a811b8216603585015289901b166049830152605d80830188905283518084039091018152607d909201909252805191012060009061078490858585611720565b6000816002018351101561149d57600080fd5b50016002015190565b91016020012090565b600081601401835110156114c257600080fd5b500160200151600160601b900490565b600081602001835110156114e557600080fd5b50016020015190565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561156c575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206107f6565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff1660021461160f576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b815160808301516116209190611d02565b92915050565b61162e611f96565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611693565b611680611f96565b8152602001906001900390816116785790505b508152600260208201526040019290925250919050565b6116b2611f96565b6040805160028082526060828101909352816020015b6116d0611f96565b8152602001906001900390816116c857905050905083816000815181106116f357fe5b6020026020010181905250828160018151811061170c57fe5b60200260200101819052506107f681611b3e565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b611766611f96565b60408051600380825260808201909252606091816020015b611786611f96565b81526020019060019003908161177e5790505090506117ad876001600160a01b0316611ab9565b816000815181106117ba57fe5b60200260200101819052506117d7896001600160a01b0316611ab9565b816001815181106117e457fe5b60200260200101819052506117f886611ab9565b8160028151811061180557fe5b602090810291909101015260408051600380825260808201909252606091816020015b611830611f96565b8152602001906001900390816118285790505090506118518b60ff16611ab9565b8160008151811061185e57fe5b602002602001018190525061187b896001600160a01b0316611ab9565b8160018151811061188857fe5b602002602001018190525061189c82611b3e565b816002815181106118a957fe5b602090810291909101015260408051600480825260a08201909252606091816020015b6118d4611f96565b8152602001906001900390816118cc5790505090506118f287611ab9565b816000815181106118ff57fe5b602002602001018190525061191386611ab9565b8160018151811061192057fe5b602002602001018190525061193485611ab9565b8160028151811061194157fe5b602002602001018190525061195582611b3e565b8160038151811061196257fe5b602002602001018190525061197681611b3e565b9c9b505050505050505050505050565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081896040516020018083805190602001908083835b602083106119fc5780518252601f1990920191602091820191016119dd565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201905282519201919091209250611a42915089905088611d3c565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa158015611aa1573d6000803e3d6000fd5b5050604051601f1901519a9950505050505050505050565b611ac1611f96565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611b26565b611b13611f96565b815260200190600190039081611b0b5790505b50815260006020820152600160409091015292915050565b611b46611f96565b611b508251611dca565b611ba1576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015611bd857838181518110611bbb57fe5b602002602001015160800151820191508080600101915050611ba6565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b611c35611f96565b611c3e82611dd1565b611c84576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060611c938360400151611de0565b9050610fe3818460800151611ce3565b60408051600080825260208201909252611cbe816001611cc4565b91505090565b6000611cce611f96565b611cd88484611ce3565b90506107f6816115ba565b611ceb611f96565b6000611cf684611eb8565b90506107f68184611626565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b818101602081810151604083015160609093015160001a9290918401601b841015611d6857601b840193505b8360ff16601b1480611d7d57508360ff16601c145b611dc2576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b6008101590565b60006116208260600151611f78565b6060600882511115611e30576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611e5d578160200160208202803883390190505b50805190915060005b81811015611eaf576000611e8c868381518110611e7f57fe5b60200260200101516109d7565b905080848381518110611e9b57fe5b602090810291909101015250600101611e66565b50909392505050565b6000600882511115611f08576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611f4c578181015183820152602001611f34565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6000600c60ff8316108015611620575050600360ff91909116101590565b6040518060a0016040528060008152602001611fb0611fca565b815260606020820181905260006040830181905291015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a7231582099ad6d835c2728505c3424f00afd5f8d6b095f5f795401cad0ef4e73e9454ca664736f6c634300050d0032"

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

// Erc20Hash is a free data retrieval call binding the contract method 0xf3e5743e.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc20Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc20Hash", to, from, erc20, value, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// Erc20Hash is a free data retrieval call binding the contract method 0xf3e5743e.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc20Hash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20Hash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, timestamp, messageNum)
}

// Erc20Hash is a free data retrieval call binding the contract method 0xf3e5743e.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc20Hash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20Hash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, timestamp, messageNum)
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0xfadffc5f.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc20MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc20MessageHash", to, from, erc20, value, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0xfadffc5f.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc20MessageHash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20MessageHash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, timestamp, messageNum)
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0xfadffc5f.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc20MessageHash(to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20MessageHash(&_MessageTester.CallOpts, to, from, erc20, value, blockNumber, timestamp, messageNum)
}

// Erc721Hash is a free data retrieval call binding the contract method 0x87c406bf.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc721Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc721Hash", to, from, erc721, id, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// Erc721Hash is a free data retrieval call binding the contract method 0x87c406bf.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc721Hash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721Hash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, timestamp, messageNum)
}

// Erc721Hash is a free data retrieval call binding the contract method 0x87c406bf.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc721Hash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721Hash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, timestamp, messageNum)
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xdd182d69.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc721MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc721MessageHash", to, from, erc721, id, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xdd182d69.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc721MessageHash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721MessageHash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, timestamp, messageNum)
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xdd182d69.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc721MessageHash(to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721MessageHash(&_MessageTester.CallOpts, to, from, erc721, id, blockNumber, timestamp, messageNum)
}

// EthHash is a free data retrieval call binding the contract method 0xbbfd47ce.
//
// Solidity: function ethHash(address to, address from, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) EthHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "ethHash", to, from, value, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// EthHash is a free data retrieval call binding the contract method 0xbbfd47ce.
//
// Solidity: function ethHash(address to, address from, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) EthHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthHash(&_MessageTester.CallOpts, to, from, value, blockNumber, timestamp, messageNum)
}

// EthHash is a free data retrieval call binding the contract method 0xbbfd47ce.
//
// Solidity: function ethHash(address to, address from, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) EthHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthHash(&_MessageTester.CallOpts, to, from, value, blockNumber, timestamp, messageNum)
}

// EthMessageHash is a free data retrieval call binding the contract method 0x3bcceb7d.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) EthMessageHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "ethMessageHash", to, from, value, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// EthMessageHash is a free data retrieval call binding the contract method 0x3bcceb7d.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) EthMessageHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthMessageHash(&_MessageTester.CallOpts, to, from, value, blockNumber, timestamp, messageNum)
}

// EthMessageHash is a free data retrieval call binding the contract method 0x3bcceb7d.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) EthMessageHash(to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthMessageHash(&_MessageTester.CallOpts, to, from, value, blockNumber, timestamp, messageNum)
}

// TransactionBatchHash is a free data retrieval call binding the contract method 0xca06a372.
//
// Solidity: function transactionBatchHash(bytes transactions, uint256 blockNum, uint256 blockTimestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionBatchHash(opts *bind.CallOpts, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionBatchHash", transactions, blockNum, blockTimestamp, messageNum)
	return *ret0, err
}

// TransactionBatchHash is a free data retrieval call binding the contract method 0xca06a372.
//
// Solidity: function transactionBatchHash(bytes transactions, uint256 blockNum, uint256 blockTimestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionBatchHash(transactions []byte, blockNum *big.Int, blockTimestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionBatchHash(&_MessageTester.CallOpts, transactions, blockNum, blockTimestamp, messageNum)
}

// TransactionBatchHash is a free data retrieval call binding the contract method 0xca06a372.
//
// Solidity: function transactionBatchHash(bytes transactions, uint256 blockNum, uint256 blockTimestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionBatchHash(transactions []byte, blockNum *big.Int, blockTimestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionBatchHash(&_MessageTester.CallOpts, transactions, blockNum, blockTimestamp, messageNum)
}

// TransactionHash is a free data retrieval call binding the contract method 0xd9214a35.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionHash", chain, to, from, seqNumber, value, data, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// TransactionHash is a free data retrieval call binding the contract method 0xd9214a35.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber, timestamp, messageNum)
}

// TransactionHash is a free data retrieval call binding the contract method 0xd9214a35.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber, uint256 timestamp, uint256 messageNum) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber, timestamp, messageNum)
}

// TransactionMessageBatchHash is a free data retrieval call binding the contract method 0xa109e062.
//
// Solidity: function transactionMessageBatchHash(bytes32 prev, uint256 prevSize, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionMessageBatchHash(opts *bind.CallOpts, prev [32]byte, prevSize *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionMessageBatchHash", prev, prevSize, chain, transactions, blockNum, blockTimestamp)
	return *ret0, err
}

// TransactionMessageBatchHash is a free data retrieval call binding the contract method 0xa109e062.
//
// Solidity: function transactionMessageBatchHash(bytes32 prev, uint256 prevSize, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionMessageBatchHash(prev [32]byte, prevSize *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHash(&_MessageTester.CallOpts, prev, prevSize, chain, transactions, blockNum, blockTimestamp)
}

// TransactionMessageBatchHash is a free data retrieval call binding the contract method 0xa109e062.
//
// Solidity: function transactionMessageBatchHash(bytes32 prev, uint256 prevSize, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageBatchHash(prev [32]byte, prevSize *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHash(&_MessageTester.CallOpts, prev, prevSize, chain, transactions, blockNum, blockTimestamp)
}

// TransactionMessageBatchHashSingle is a free data retrieval call binding the contract method 0x94caf578.
//
// Solidity: function transactionMessageBatchHashSingle(uint256 start, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionMessageBatchHashSingle(opts *bind.CallOpts, start *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionMessageBatchHashSingle", start, chain, transactions, blockNum, blockTimestamp)
	return *ret0, err
}

// TransactionMessageBatchHashSingle is a free data retrieval call binding the contract method 0x94caf578.
//
// Solidity: function transactionMessageBatchHashSingle(uint256 start, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionMessageBatchHashSingle(start *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHashSingle(&_MessageTester.CallOpts, start, chain, transactions, blockNum, blockTimestamp)
}

// TransactionMessageBatchHashSingle is a free data retrieval call binding the contract method 0x94caf578.
//
// Solidity: function transactionMessageBatchHashSingle(uint256 start, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageBatchHashSingle(start *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHashSingle(&_MessageTester.CallOpts, start, chain, transactions, blockNum, blockTimestamp)
}

// TransactionMessageBatchSingleSender is a free data retrieval call binding the contract method 0xfedc217e.
//
// Solidity: function transactionMessageBatchSingleSender(uint256 start, address chain, bytes32 dataHash, bytes transactions) pure returns(address)
func (_MessageTester *MessageTesterCaller) TransactionMessageBatchSingleSender(opts *bind.CallOpts, start *big.Int, chain common.Address, dataHash [32]byte, transactions []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionMessageBatchSingleSender", start, chain, dataHash, transactions)
	return *ret0, err
}

// TransactionMessageBatchSingleSender is a free data retrieval call binding the contract method 0xfedc217e.
//
// Solidity: function transactionMessageBatchSingleSender(uint256 start, address chain, bytes32 dataHash, bytes transactions) pure returns(address)
func (_MessageTester *MessageTesterSession) TransactionMessageBatchSingleSender(start *big.Int, chain common.Address, dataHash [32]byte, transactions []byte) (common.Address, error) {
	return _MessageTester.Contract.TransactionMessageBatchSingleSender(&_MessageTester.CallOpts, start, chain, dataHash, transactions)
}

// TransactionMessageBatchSingleSender is a free data retrieval call binding the contract method 0xfedc217e.
//
// Solidity: function transactionMessageBatchSingleSender(uint256 start, address chain, bytes32 dataHash, bytes transactions) pure returns(address)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageBatchSingleSender(start *big.Int, chain common.Address, dataHash [32]byte, transactions []byte) (common.Address, error) {
	return _MessageTester.Contract.TransactionMessageBatchSingleSender(&_MessageTester.CallOpts, start, chain, dataHash, transactions)
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x63bc3d74.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber, uint256 timestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionMessageHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int, timestamp *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionMessageHash", chain, to, from, seqNumber, value, data, blockNumber, timestamp)
	return *ret0, err
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x63bc3d74.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber, uint256 timestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionMessageHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int, timestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber, timestamp)
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x63bc3d74.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data, uint256 blockNumber, uint256 timestamp) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int, timestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data, blockNumber, timestamp)
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d42220e8b02c83e30c430f597d81c68382cafa1f5dd7f9bf99a4900f39750f3264736f6c634300050d0032"

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

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820da47ac1e8fd8322ca46a10fb9d5476806e9902c358d9f32e427a6053d0c0323464736f6c634300050d0032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// Protocol is an auto generated Go binding around an Ethereum contract.
type Protocol struct {
	ProtocolCaller     // Read-only binding to the contract
	ProtocolTransactor // Write-only binding to the contract
	ProtocolFilterer   // Log filterer for contract events
}

// ProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolSession struct {
	Contract     *Protocol         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolCallerSession struct {
	Contract *ProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTransactorSession struct {
	Contract     *ProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRaw struct {
	Contract *Protocol // Generic contract binding to access the raw methods on
}

// ProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolCallerRaw struct {
	Contract *ProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTransactorRaw struct {
	Contract *ProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocol creates a new instance of Protocol, bound to a specific deployed contract.
func NewProtocol(address common.Address, backend bind.ContractBackend) (*Protocol, error) {
	contract, err := bindProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// NewProtocolCaller creates a new read-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolCaller(address common.Address, caller bind.ContractCaller) (*ProtocolCaller, error) {
	contract, err := bindProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolCaller{contract: contract}, nil
}

// NewProtocolTransactor creates a new write-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTransactor, error) {
	contract, err := bindProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTransactor{contract: contract}, nil
}

// NewProtocolFilterer creates a new log filterer instance of Protocol, bound to a specific deployed contract.
func NewProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolFilterer, error) {
	contract, err := bindProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolFilterer{contract: contract}, nil
}

// bindProtocol binds a generic wrapper to an already deployed contract.
func bindProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.ProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transact(opts, method, params...)
}

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[]"

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208f23efe075c03059d87f4e411631a7d649fa24f434c9b3b30f7e20b12e0b098e64736f6c634300050d0032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d227c8d9e79e0d2ec554c8283bdb98d60c2ea52067463a7da48c451870d2d7dd64736f6c634300050d0032"

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
