// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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

// InboxABI is the input ABI used to generate the binding from.
const InboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"BuddyContractPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InboxFuncSigs maps the 4-byte function signature to its string representation.
var InboxFuncSigs = map[string]string{
	"6f5dfdca": "deployL2ContractPair(uint256,uint256,uint256,bytes)",
	"afcc220b": "depositEthMessage(address)",
	"917cae02": "inboxMaxCount()",
	"efefa7e5": "inboxMaxValue()",
	"b75436bb": "sendL2Message(bytes)",
	"1fe927cf": "sendL2MessageFromOrigin(bytes)",
}

// InboxBin is the compiled bytecode used for deploying new contracts.
var InboxBin = "0x608060405234801561001057600080fd5b50610682806100206000396000f3fe6080604052600436106100555760003560e01c80631fe927cf1461005a5780636f5dfdca146100d9578063917cae0214610169578063afcc220b14610190578063b75436bb146101b6578063efefa7e514610233575b600080fd5b34801561006657600080fd5b506100d76004803603602081101561007d57600080fd5b81019060208101813564010000000081111561009857600080fd5b8201836020820111156100aa57600080fd5b803590602001918460018302840111640100000000831117156100cc57600080fd5b509092509050610248565b005b3480156100e557600080fd5b506100d7600480360360808110156100fc57600080fd5b8135916020810135916040820135919081019060808101606082013564010000000081111561012a57600080fd5b82018360208201111561013c57600080fd5b8035906020019184600183028401116401000000008311171561015e57600080fd5b5090925090506102fe565b34801561017557600080fd5b5061017e6103d6565b60408051918252519081900360200190f35b6100d7600480360360208110156101a657600080fd5b50356001600160a01b03166103dc565b3480156101c257600080fd5b506100d7600480360360208110156101d957600080fd5b8101906020810181356401000000008111156101f457600080fd5b82018360208201111561020657600080fd5b8035906020019184600183028401116401000000008311171561022857600080fd5b509092509050610416565b34801561023f57600080fd5b5061017e61045c565b33321461028a576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b6000806102b56003338686604051808383808284376040519201829003909120935061046292505050565b60408051600381523360208201528151939550919350839285927f852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab92908290030190a350505050565b6103073361049f565b610358576040805162461bcd60e51b815260206004820152601a60248201527f6d7573742062652063616c6c656420627920636f6e7472616374000000000000604482015290519081900360640190fd5b6103a460053387878787876040516020018086815260200185815260200184815260200183838082843780830192505050955050505050506040516020818303038152906040526104db565b60405133907f49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f90600090a25050505050565b60015481565b604080516001600160a01b0383166020820152348183015281518082038301815260609091019091526104139060009033906104db565b50565b61045860033384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104db92505050565b5050565b60005481565b60015460008054909182918261047c88884342878b6105b2565b90506104888282610620565b600055506001828101905590969095509350505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906104d357508115155b949350505050565b6000806104f085858580519060200120610462565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561056f578181015183820152602001610557565b50505050905090810190601f16801561059c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b60408051602080820194909452808201929092528051808303820181526060909201905280519101209056fea26469706673582212206f68a10189e51677f97fc3a7ead2b528a3bf914f1b5cc4ba9cd967a547d5a35c64736f6c634300060c0033"

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Inbox *InboxCaller) InboxMaxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "inboxMaxCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Inbox *InboxSession) InboxMaxCount() (*big.Int, error) {
	return _Inbox.Contract.InboxMaxCount(&_Inbox.CallOpts)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Inbox *InboxCallerSession) InboxMaxCount() (*big.Int, error) {
	return _Inbox.Contract.InboxMaxCount(&_Inbox.CallOpts)
}

// InboxMaxValue is a free data retrieval call binding the contract method 0xefefa7e5.
//
// Solidity: function inboxMaxValue() view returns(bytes32)
func (_Inbox *InboxCaller) InboxMaxValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "inboxMaxValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxMaxValue is a free data retrieval call binding the contract method 0xefefa7e5.
//
// Solidity: function inboxMaxValue() view returns(bytes32)
func (_Inbox *InboxSession) InboxMaxValue() ([32]byte, error) {
	return _Inbox.Contract.InboxMaxValue(&_Inbox.CallOpts)
}

// InboxMaxValue is a free data retrieval call binding the contract method 0xefefa7e5.
//
// Solidity: function inboxMaxValue() view returns(bytes32)
func (_Inbox *InboxCallerSession) InboxMaxValue() ([32]byte, error) {
	return _Inbox.Contract.InboxMaxValue(&_Inbox.CallOpts)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxTransactor) DepositEthMessage(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEthMessage", to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthMessage(&_Inbox.TransactOpts, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxTransactorSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthMessage(&_Inbox.TransactOpts, to)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// InboxBuddyContractPairIterator is returned from FilterBuddyContractPair and is used to iterate over the raw logs and unpacked data for BuddyContractPair events raised by the Inbox contract.
type InboxBuddyContractPairIterator struct {
	Event *InboxBuddyContractPair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxBuddyContractPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxBuddyContractPair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxBuddyContractPair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxBuddyContractPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxBuddyContractPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxBuddyContractPair represents a BuddyContractPair event raised by the Inbox contract.
type InboxBuddyContractPair struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractPair is a free log retrieval operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) FilterBuddyContractPair(opts *bind.FilterOpts, sender []common.Address) (*InboxBuddyContractPairIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxBuddyContractPairIterator{contract: _Inbox.contract, event: "BuddyContractPair", logs: logs, sub: sub}, nil
}

// WatchBuddyContractPair is a free log subscription operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) WatchBuddyContractPair(opts *bind.WatchOpts, sink chan<- *InboxBuddyContractPair, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxBuddyContractPair)
				if err := _Inbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuddyContractPair is a log parse operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) ParseBuddyContractPair(log types.Log) (*InboxBuddyContractPair, error) {
	event := new(InboxBuddyContractPair)
	if err := _Inbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Inbox contract.
type InboxMessageDeliveredIterator struct {
	Event *InboxMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxMessageDelivered represents a MessageDelivered event raised by the Inbox contract.
type InboxMessageDelivered struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*InboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &InboxMessageDeliveredIterator{contract: _Inbox.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxMessageDelivered, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) ParseMessageDelivered(log types.Log) (*InboxMessageDelivered, error) {
	event := new(InboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxMessageDeliveredFromOriginIterator is returned from FilterMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for MessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxMessageDeliveredFromOriginIterator struct {
	Event *InboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxMessageDeliveredFromOrigin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxMessageDeliveredFromOrigin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxMessageDeliveredFromOrigin represents a MessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxMessageDeliveredFromOrigin struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Inbox *InboxFilterer) FilterMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*InboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &InboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "MessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Inbox *InboxFilterer) WatchMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxMessageDeliveredFromOrigin, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDeliveredFromOrigin is a log parse operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Inbox *InboxFilterer) ParseMessageDeliveredFromOrigin(log types.Log) (*InboxMessageDeliveredFromOrigin, error) {
	event := new(InboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220360a944955baccdb8a59cb1cc677eb7088a2170c3c5166c60239fa30fcfd92a164736f6c634300060c0033"

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
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// NodeABI is the input ABI used to generate the binding from.
const NodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkConfirmInvalid\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"}],\"name\":\"checkConfirmOutOfOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStakerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"}],\"name\":\"checkConfirmValid\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"removeStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeFuncSigs maps the 4-byte function signature to its string representation.
var NodeFuncSigs = map[string]string{
	"2466696e": "addStaker(address)",
	"5b8b2280": "challengeHash()",
	"1a8a092b": "checkConfirmInvalid(uint256)",
	"284426b2": "checkConfirmOutOfOrder(uint256)",
	"6cf00e7e": "checkConfirmValid(uint256,uint256)",
	"97bdc510": "confirmData()",
	"2edfb42a": "deadlineBlock()",
	"83197ef0": "destroy()",
	"a406b374": "initialize(address,bytes32,bytes32,bytes32,uint256,uint256)",
	"6f791d29": "isMaster()",
	"479c9254": "prev()",
	"96a9fdc0": "removeStaker(address)",
	"dff69787": "stakerCount()",
	"9168ae72": "stakers(address)",
	"701da98e": "stateHash()",
}

// NodeBin is the compiled bytecode used for deploying new contracts.
var NodeBin = "0x608060405234801561001057600080fd5b506000805460ff191660011790556106dd8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636f791d291161009757806396a9fdc01161006657806396a9fdc0146101fb57806397bdc51014610221578063a406b37414610229578063dff697871461026d576100f5565b80636f791d29146101a9578063701da98e146101c557806383197ef0146101cd5780639168ae72146101d5576100f5565b80632edfb42a116100d35780632edfb42a1461015c578063479c9254146101765780635b8b22801461017e5780636cf00e7e14610186576100f5565b80631a8a092b146100fa5780632466696e14610119578063284426b21461013f575b600080fd5b6101176004803603602081101561011057600080fd5b5035610275565b005b6101176004803603602081101561012f57600080fd5b50356001600160a01b0316610305565b6101176004803603602081101561015557600080fd5b50356103e1565b6101646103f0565b60408051918252519081900360200190f35b6101646103f6565b6101646103fc565b6101176004803603604081101561019c57600080fd5b5080359060200135610402565b6101b161051d565b604080519115158252519081900360200190f35b610164610526565b61011761052c565b6101b1600480360360208110156101eb57600080fd5b50356001600160a01b031661057c565b6101176004803603602081101561021157600080fd5b50356001600160a01b0316610591565b610164610663565b610117600480360360c081101561023f57600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a00135610669565b6101646106a1565b6005544310156102be576040805162461bcd60e51b815260206004820152600f60248201526e4245464f52455f444541444c494e4560881b604482015290519081900360640190fd5b8060065414610302576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b50565b6008546001600160a01b03163314610352576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b6001600160a01b03811660009081526007602052604090205460ff16156103b1576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b6001600160a01b03166000908152600760205260409020805460ff19166001908117909155600680549091019055565b80600454141561030257600080fd5b60055481565b60045481565b60025481565b60055443101561044b576040805162461bcd60e51b815260206004820152600f60248201526e4245464f52455f444541444c494e4560881b604482015290519081900360640190fd5b8060045414610490576040805162461bcd60e51b815260206004820152600c60248201526b24a72b20a624a22fa82922ab60a11b604482015290519081900360640190fd5b81600654146104d7576040805162461bcd60e51b815260206004820152600e60248201526d1393d517d0531317d4d51052d15160921b604482015290519081900360640190fd5b60008211610519576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f5354414b45525360b01b604482015290519081900360640190fd5b5050565b60005460ff1690565b60015481565b6008546001600160a01b03163314610579576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b33ff5b60076020526000908152604090205460ff1681565b6008546001600160a01b031633146105de576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b6001600160a01b03811660009081526007602052604090205460ff16610638576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6001600160a01b03166000908152600760205260409020805460ff1916905560068054600019019055565b60035481565b600880546001600160a01b0319166001600160a01b039790971696909617909555600193909355600291909155600355600455600555565b6006548156fea2646970667358221220c6584dd6e86e12a693a74657354a1893fd6e16cb85d48d6717fc607f6c2793c564736f6c634300060c0033"

// DeployNode deploys a new Ethereum contract, binding an instance of Node to it.
func DeployNode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Node, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_Node *NodeCaller) ChallengeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "challengeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_Node *NodeSession) ChallengeHash() ([32]byte, error) {
	return _Node.Contract.ChallengeHash(&_Node.CallOpts)
}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_Node *NodeCallerSession) ChallengeHash() ([32]byte, error) {
	return _Node.Contract.ChallengeHash(&_Node.CallOpts)
}

// CheckConfirmInvalid is a free data retrieval call binding the contract method 0x1a8a092b.
//
// Solidity: function checkConfirmInvalid(uint256 zombieStakerCount) view returns()
func (_Node *NodeCaller) CheckConfirmInvalid(opts *bind.CallOpts, zombieStakerCount *big.Int) error {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "checkConfirmInvalid", zombieStakerCount)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmInvalid is a free data retrieval call binding the contract method 0x1a8a092b.
//
// Solidity: function checkConfirmInvalid(uint256 zombieStakerCount) view returns()
func (_Node *NodeSession) CheckConfirmInvalid(zombieStakerCount *big.Int) error {
	return _Node.Contract.CheckConfirmInvalid(&_Node.CallOpts, zombieStakerCount)
}

// CheckConfirmInvalid is a free data retrieval call binding the contract method 0x1a8a092b.
//
// Solidity: function checkConfirmInvalid(uint256 zombieStakerCount) view returns()
func (_Node *NodeCallerSession) CheckConfirmInvalid(zombieStakerCount *big.Int) error {
	return _Node.Contract.CheckConfirmInvalid(&_Node.CallOpts, zombieStakerCount)
}

// CheckConfirmOutOfOrder is a free data retrieval call binding the contract method 0x284426b2.
//
// Solidity: function checkConfirmOutOfOrder(uint256 latestConfirmed) view returns()
func (_Node *NodeCaller) CheckConfirmOutOfOrder(opts *bind.CallOpts, latestConfirmed *big.Int) error {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "checkConfirmOutOfOrder", latestConfirmed)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmOutOfOrder is a free data retrieval call binding the contract method 0x284426b2.
//
// Solidity: function checkConfirmOutOfOrder(uint256 latestConfirmed) view returns()
func (_Node *NodeSession) CheckConfirmOutOfOrder(latestConfirmed *big.Int) error {
	return _Node.Contract.CheckConfirmOutOfOrder(&_Node.CallOpts, latestConfirmed)
}

// CheckConfirmOutOfOrder is a free data retrieval call binding the contract method 0x284426b2.
//
// Solidity: function checkConfirmOutOfOrder(uint256 latestConfirmed) view returns()
func (_Node *NodeCallerSession) CheckConfirmOutOfOrder(latestConfirmed *big.Int) error {
	return _Node.Contract.CheckConfirmOutOfOrder(&_Node.CallOpts, latestConfirmed)
}

// CheckConfirmValid is a free data retrieval call binding the contract method 0x6cf00e7e.
//
// Solidity: function checkConfirmValid(uint256 totalStakerCount, uint256 latestConfirmed) view returns()
func (_Node *NodeCaller) CheckConfirmValid(opts *bind.CallOpts, totalStakerCount *big.Int, latestConfirmed *big.Int) error {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "checkConfirmValid", totalStakerCount, latestConfirmed)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmValid is a free data retrieval call binding the contract method 0x6cf00e7e.
//
// Solidity: function checkConfirmValid(uint256 totalStakerCount, uint256 latestConfirmed) view returns()
func (_Node *NodeSession) CheckConfirmValid(totalStakerCount *big.Int, latestConfirmed *big.Int) error {
	return _Node.Contract.CheckConfirmValid(&_Node.CallOpts, totalStakerCount, latestConfirmed)
}

// CheckConfirmValid is a free data retrieval call binding the contract method 0x6cf00e7e.
//
// Solidity: function checkConfirmValid(uint256 totalStakerCount, uint256 latestConfirmed) view returns()
func (_Node *NodeCallerSession) CheckConfirmValid(totalStakerCount *big.Int, latestConfirmed *big.Int) error {
	return _Node.Contract.CheckConfirmValid(&_Node.CallOpts, totalStakerCount, latestConfirmed)
}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_Node *NodeCaller) ConfirmData(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "confirmData")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_Node *NodeSession) ConfirmData() ([32]byte, error) {
	return _Node.Contract.ConfirmData(&_Node.CallOpts)
}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_Node *NodeCallerSession) ConfirmData() ([32]byte, error) {
	return _Node.Contract.ConfirmData(&_Node.CallOpts)
}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_Node *NodeCaller) DeadlineBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "deadlineBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_Node *NodeSession) DeadlineBlock() (*big.Int, error) {
	return _Node.Contract.DeadlineBlock(&_Node.CallOpts)
}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_Node *NodeCallerSession) DeadlineBlock() (*big.Int, error) {
	return _Node.Contract.DeadlineBlock(&_Node.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Node *NodeCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Node *NodeSession) IsMaster() (bool, error) {
	return _Node.Contract.IsMaster(&_Node.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Node *NodeCallerSession) IsMaster() (bool, error) {
	return _Node.Contract.IsMaster(&_Node.CallOpts)
}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_Node *NodeCaller) Prev(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "prev")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_Node *NodeSession) Prev() (*big.Int, error) {
	return _Node.Contract.Prev(&_Node.CallOpts)
}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_Node *NodeCallerSession) Prev() (*big.Int, error) {
	return _Node.Contract.Prev(&_Node.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Node *NodeCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Node *NodeSession) StakerCount() (*big.Int, error) {
	return _Node.Contract.StakerCount(&_Node.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Node *NodeCallerSession) StakerCount() (*big.Int, error) {
	return _Node.Contract.StakerCount(&_Node.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bool)
func (_Node *NodeCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "stakers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bool)
func (_Node *NodeSession) Stakers(arg0 common.Address) (bool, error) {
	return _Node.Contract.Stakers(&_Node.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bool)
func (_Node *NodeCallerSession) Stakers(arg0 common.Address) (bool, error) {
	return _Node.Contract.Stakers(&_Node.CallOpts, arg0)
}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_Node *NodeCaller) StateHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "stateHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_Node *NodeSession) StateHash() ([32]byte, error) {
	return _Node.Contract.StateHash(&_Node.CallOpts)
}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_Node *NodeCallerSession) StateHash() ([32]byte, error) {
	return _Node.Contract.StateHash(&_Node.CallOpts)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns()
func (_Node *NodeTransactor) AddStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "addStaker", staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns()
func (_Node *NodeSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _Node.Contract.AddStaker(&_Node.TransactOpts, staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns()
func (_Node *NodeTransactorSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _Node.Contract.AddStaker(&_Node.TransactOpts, staker)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Node *NodeTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Node *NodeSession) Destroy() (*types.Transaction, error) {
	return _Node.Contract.Destroy(&_Node.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Node *NodeTransactorSession) Destroy() (*types.Transaction, error) {
	return _Node.Contract.Destroy(&_Node.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_Node *NodeTransactor) Initialize(opts *bind.TransactOpts, _rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "initialize", _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_Node *NodeSession) Initialize(_rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Initialize(&_Node.TransactOpts, _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_Node *NodeTransactorSession) Initialize(_rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Initialize(&_Node.TransactOpts, _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_Node *NodeTransactor) RemoveStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "removeStaker", staker)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_Node *NodeSession) RemoveStaker(staker common.Address) (*types.Transaction, error) {
	return _Node.Contract.RemoveStaker(&_Node.TransactOpts, staker)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_Node *NodeTransactorSession) RemoveStaker(staker common.Address) (*types.Transaction, error) {
	return _Node.Contract.RemoveStaker(&_Node.TransactOpts, staker)
}

// OutboxABI is the input ABI used to generate the binding from.
const OutboxABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxFuncSigs maps the 4-byte function signature to its string representation.
var OutboxFuncSigs = map[string]string{
	"c4fb000c": "executeTransaction(uint256,bytes,uint256,address,uint256,bytes)",
}

// OutboxBin is the compiled bytecode used for deploying new contracts.
var OutboxBin = "0x608060405234801561001057600080fd5b5061040d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c4fb000c14610030575b600080fd5b610114600480360360c081101561004657600080fd5b8135919081019060408101602082013564010000000081111561006857600080fd5b82018360208201111561007a57600080fd5b8035906020019184600183028401116401000000008311171561009c57600080fd5b919390928235926001600160a01b036020820135169260408201359290916080810190606001356401000000008111156100d557600080fd5b8201836020820111156100e757600080fd5b8035906020019184600183028401116401000000008311171561010957600080fd5b509092509050610116565b005b60008460601b60601c6001600160a01b031684848460405160200180858152602001848152602001838380828437808301925050509450505050506040516020818303038152906040528051906020012090506101ad8989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b925086915061022c9050565b6000856001600160a01b0316858585604051808383808284376040519201945060009350909150508083038185875af1925050503d806000811461020d576040519150601f19603f3d011682016040523d82523d6000602084013e610212565b606091505b505090508061022057600080fd5b50505050505050505050565b600160001b8118905060006102458483856001016102ca565b5090506000858154811061025557fe5b600091825260208220015460408051630ad0379b60e01b8152600481018590526024810187905290516001600160a01b0390921692630ad0379b9260448084019382900301818387803b1580156102ab57600080fd5b505af11580156102bf573d6000803e3d6000fd5b505050505050505050565b60008080848160205b885181116103c9578089015193506020818a5103602001816102f157fe5b0491505b6000821180156103085750600287066001145b801561031657508160020a87115b1561032e5760029096046001908101969401936102f5565b6002870661037957838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002878161037157fe5b0496506103bb565b82846040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600287816103b457fe5b0460010196505b6001909401936020016102d3565b50909350505093509391505056fea2646970667358221220edffdc416f2b54a387c70a8c9821b1f767bc7c6998a6176ea12665e0aff569ef64736f6c634300060c0033"

// DeployOutbox deploys a new Ethereum contract, binding an instance of Outbox to it.
func DeployOutbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Outbox, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// Outbox is an auto generated Go binding around an Ethereum contract.
type Outbox struct {
	OutboxCaller     // Read-only binding to the contract
	OutboxTransactor // Write-only binding to the contract
	OutboxFilterer   // Log filterer for contract events
}

// OutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxSession struct {
	Contract     *Outbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxCallerSession struct {
	Contract *OutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxTransactorSession struct {
	Contract     *OutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxRaw struct {
	Contract *Outbox // Generic contract binding to access the raw methods on
}

// OutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxCallerRaw struct {
	Contract *OutboxCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxTransactorRaw struct {
	Contract *OutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutbox creates a new instance of Outbox, bound to a specific deployed contract.
func NewOutbox(address common.Address, backend bind.ContractBackend) (*Outbox, error) {
	contract, err := bindOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// NewOutboxCaller creates a new read-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxCaller(address common.Address, caller bind.ContractCaller) (*OutboxCaller, error) {
	contract, err := bindOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxCaller{contract: contract}, nil
}

// NewOutboxTransactor creates a new write-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxTransactor, error) {
	contract, err := bindOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxTransactor{contract: contract}, nil
}

// NewOutboxFilterer creates a new log filterer instance of Outbox, bound to a specific deployed contract.
func NewOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxFilterer, error) {
	contract, err := bindOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxFilterer{contract: contract}, nil
}

// bindOutbox binds a generic wrapper to an already deployed contract.
func bindOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.OutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transact(opts, method, params...)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransaction", outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// OutboxEntryABI is the input ABI used to generate the binding from.
const OutboxEntryABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"calcRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"spendOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxEntryFuncSigs maps the 4-byte function signature to its string representation.
var OutboxEntryFuncSigs = map[string]string{
	"0ad0379b": "spendOutput(bytes32,uint256)",
}

// OutboxEntryBin is the compiled bytecode used for deploying new contracts.
var OutboxEntryBin = "0x608060405234801561001057600080fd5b506040516101373803806101378339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319163317905560015560df806100586000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630ad0379b14602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135604f565b005b6000546001600160a01b03163314606557600080fd5b60008181526002602052604090205460ff1615608057600080fd5b6001548214608d57600080fd5b6000908152600260205260409020805460ff191660011790555056fea2646970667358221220ac2ea98d21448678a949baebc8abc30d5a0bb7897571455392215b23603f3c9e64736f6c634300060c0033"

// DeployOutboxEntry deploys a new Ethereum contract, binding an instance of OutboxEntry to it.
func DeployOutboxEntry(auth *bind.TransactOpts, backend bind.ContractBackend, root [32]byte) (common.Address, *types.Transaction, *OutboxEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxEntryBin), backend, root)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// OutboxEntry is an auto generated Go binding around an Ethereum contract.
type OutboxEntry struct {
	OutboxEntryCaller     // Read-only binding to the contract
	OutboxEntryTransactor // Write-only binding to the contract
	OutboxEntryFilterer   // Log filterer for contract events
}

// OutboxEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxEntrySession struct {
	Contract     *OutboxEntry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxEntryCallerSession struct {
	Contract *OutboxEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OutboxEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxEntryTransactorSession struct {
	Contract     *OutboxEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OutboxEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxEntryRaw struct {
	Contract *OutboxEntry // Generic contract binding to access the raw methods on
}

// OutboxEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxEntryCallerRaw struct {
	Contract *OutboxEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxEntryTransactorRaw struct {
	Contract *OutboxEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxEntry creates a new instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntry(address common.Address, backend bind.ContractBackend) (*OutboxEntry, error) {
	contract, err := bindOutboxEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// NewOutboxEntryCaller creates a new read-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryCaller(address common.Address, caller bind.ContractCaller) (*OutboxEntryCaller, error) {
	contract, err := bindOutboxEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryCaller{contract: contract}, nil
}

// NewOutboxEntryTransactor creates a new write-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxEntryTransactor, error) {
	contract, err := bindOutboxEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryTransactor{contract: contract}, nil
}

// NewOutboxEntryFilterer creates a new log filterer instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxEntryFilterer, error) {
	contract, err := bindOutboxEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryFilterer{contract: contract}, nil
}

// bindOutboxEntry binds a generic wrapper to an already deployed contract.
func bindOutboxEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.OutboxEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transact(opts, method, params...)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 index) returns()
func (_OutboxEntry *OutboxEntryTransactor) SpendOutput(opts *bind.TransactOpts, calcRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "spendOutput", calcRoot, index)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 index) returns()
func (_OutboxEntry *OutboxEntrySession) SpendOutput(calcRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, calcRoot, index)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 index) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) SpendOutput(calcRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, calcRoot, index)
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"BuddyContractPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"indexed\":false,\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"SentLogs\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"addStakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"addStakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkNoRecentStake\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"logAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"staker1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum1\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"staker2Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum2\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"inboxConsistencyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inboxDeltaHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionCheckTime\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"newStakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"newStakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"contractNode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxReduction\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"successorWithStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectNextNodeOutOfOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerList\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupFuncSigs maps the 4-byte function signature to its string representation.
var RollupFuncSigs = map[string]string{
	"7ba3ca01": "addStakeOnExistingNode(bytes32,uint256,uint256)",
	"af46618b": "addStakeOnNewNode(bytes32,uint256,uint256,bytes32[7],uint256[10])",
	"45c5b2c7": "addToDeposit(address)",
	"5e8ef106": "arbGasSpeedLimitPerBlock()",
	"76e7e23b": "baseStake()",
	"5dbaf68b": "challengeFactory()",
	"46c2781a": "challengePeriodBlocks()",
	"be211c9a": "checkNoRecentStake()",
	"73f33b06": "checkUnresolved()",
	"fa7803e6": "completeChallenge(address,address)",
	"396b8cbc": "confirmNextNode(bytes32,bytes,uint256[])",
	"04a28064": "countStakedZombies(address)",
	"58aab3d3": "createChallenge(address,uint256,address,uint256,bytes32,bytes32,bytes32,uint256)",
	"4d26732d": "currentRequiredStake()",
	"6f5dfdca": "deployL2ContractPair(uint256,uint256,uint256,bytes)",
	"afcc220b": "depositEthMessage(address)",
	"c4fb000c": "executeTransaction(uint256,bytes,uint256,address,uint256,bytes)",
	"d735e21d": "firstUnresolvedNode()",
	"ad71bd36": "getStakers(uint256,uint256)",
	"917cae02": "inboxMaxCount()",
	"efefa7e5": "inboxMaxValue()",
	"65f7f80d": "latestConfirmed()",
	"7ba9534a": "latestNodeCreated()",
	"ad432faf": "newStakeOnExistingNode(bytes32,uint256,uint256)",
	"9a4fcae7": "newStakeOnNewNode(bytes32,uint256,uint256,uint256,bytes32[7],uint256[10])",
	"d93fe9c4": "nodeFactory()",
	"1c53c280": "nodes(uint256)",
	"1e83d30f": "reduceDeposit(uint256)",
	"0e1ef04c": "rejectNextNode(uint256,address)",
	"4802c739": "rejectNextNodeOutOfOrder()",
	"edfd03ed": "removeOldZombies(uint256)",
	"7e2d2155": "removeZombie(uint256,uint256)",
	"7427be51": "returnOldDeposit(address)",
	"b75436bb": "sendL2Message(bytes)",
	"1fe927cf": "sendL2MessageFromOrigin(bytes)",
	"51ed6a30": "stakeToken()",
	"dff69787": "stakerCount()",
	"348e50c6": "stakerList(uint256)",
	"729cfe3b": "stakerMap(address)",
}

// RollupBin is the compiled bytecode used for deploying new contracts.
var RollupBin = "0x60806040523480156200001157600080fd5b50604051620043143803806200431483398181016040526101208110156200003857600080fd5b815160208301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0180519751999b989a969995989497939692959194919392820192846401000000008211156200009157600080fd5b908301906020820185811115620000a757600080fd5b8251640100000000811182820188101715620000c257600080fd5b82525081516020918201929091019080838360005b83811015620000f1578181015183820152602001620000d7565b50505050905090810190601f1680156200011f5780820380516001836020036101000a031916815260200191505b506040525050600f80546001600160a01b03199081166001600160a01b0387811691909117909255601080549091169185169190911790555060006200017843828c81808080806200036e602090811b620021a417901c565b6010546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905260848201819052915193945090926001600160a01b039092169163d45ab2b59160a48082019260209290919082900301818787803b158015620001e957600080fd5b505af1158015620001fe573d6000803e3d6000fd5b505050506040513d60208110156200021557600080fd5b81019080805190602001909291905050509050806006600080815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555089600b8190555088600c8190555087600d8190555086600e60006101000a8154816001600160a01b0302191690836001600160a01b03160217905550620003578a8a8a8a60601b6001600160601b0319168a60601b6001600160601b031916886040516020018087815260200186815260200185815260200184815260200183815260200182805190602001908083835b60208310620003125780518252601f199092019160209182019101620002f1565b6001836020036101000a0380198251168184511680821785525050505050509050019650505050505050604051602081830303815290604052620003c860201b60201c565b5050600160045550620005c5975050505050505050565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b620003d660043083620003d9565b50565b600080620003f685858580519060200120620004bb60201b60201c565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620004775781810151838201526020016200045d565b50505050905090810190601f168015620004a55780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b600080600060015490506000805490506000620004e888884342878b6200051860201b620021fe1760201c565b90506200050182826200059960201b6200226c1760201c565b600055506001828101905590969095509350505050565b6040805160f89790971b7fff000000000000000000000000000000000000000000000000000000000000001660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b613d3f80620005d56000396000f3fe6080604052600436106102305760003560e01c80637427be511161012e578063afcc220b116100ab578063d93fe9c41161006f578063d93fe9c414610a77578063dff6978714610a8c578063edfd03ed14610aa1578063efefa7e514610acb578063fa7803e614610ae057610230565b8063afcc220b146108bf578063b75436bb146108e5578063be211c9a14610960578063c4fb000c14610975578063d735e21d14610a6257610230565b8063917cae02116100f2578063917cae02146107845780639a4fcae714610799578063ad432faf146107d4578063ad71bd36146107fd578063af46618b1461087d57610230565b80637427be51146106c157806376e7e23b146106f45780637ba3ca01146107095780637ba9534a1461073f5780637e2d21551461075457610230565b80634802c739116101bc5780635e8ef106116101805780635e8ef1061461058b57806365f7f80d146105a05780636f5dfdca146105b5578063729cfe3b1461064357806373f33b06146106ac57610230565b80634802c739146104d75780634d26732d146104ec57806351ed6a301461050157806358aab3d3146105165780635dbaf68b1461057657610230565b80631fe927cf116102035780631fe927cf14610325578063348e50c6146103a0578063396b8cbc146103ca57806345c5b2c71461049c57806346c2781a146104c257610230565b806304a28064146102355780630e1ef04c1461027a5780631c53c280146102b55780631e83d30f146102fb575b600080fd5b34801561024157600080fd5b506102686004803603602081101561025857600080fd5b50356001600160a01b0316610b1b565b60408051918252519081900360200190f35b34801561028657600080fd5b506102b36004803603604081101561029d57600080fd5b50803590602001356001600160a01b0316610bde565b005b3480156102c157600080fd5b506102df600480360360208110156102d857600080fd5b5035610efb565b604080516001600160a01b039092168252519081900360200190f35b34801561030757600080fd5b506102b36004803603602081101561031e57600080fd5b5035610f16565b34801561033157600080fd5b506102b36004803603602081101561034857600080fd5b810190602081018135600160201b81111561036257600080fd5b82018360208201111561037457600080fd5b803590602001918460018302840111600160201b8311171561039557600080fd5b509092509050610f91565b3480156103ac57600080fd5b506102df600480360360208110156103c357600080fd5b5035611047565b3480156103d657600080fd5b506102b3600480360360608110156103ed57600080fd5b81359190810190604081016020820135600160201b81111561040e57600080fd5b82018360208201111561042057600080fd5b803590602001918460018302840111600160201b8311171561044157600080fd5b919390929091602081019035600160201b81111561045e57600080fd5b82018360208201111561047057600080fd5b803590602001918460208302840111600160201b8311171561049157600080fd5b50909250905061106e565b6102b3600480360360208110156104b257600080fd5b50356001600160a01b0316611301565b3480156104ce57600080fd5b5061026861132e565b3480156104e357600080fd5b506102b3611334565b3480156104f857600080fd5b506102686113c2565b34801561050d57600080fd5b506102df611533565b34801561052257600080fd5b506102b3600480360361010081101561053a57600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060808101359060a08101359060c08101359060e00135611542565b34801561058257600080fd5b506102df611579565b34801561059757600080fd5b50610268611588565b3480156105ac57600080fd5b5061026861158e565b3480156105c157600080fd5b506102b3600480360360808110156105d857600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561060557600080fd5b82018360208201111561061757600080fd5b803590602001918460018302840111600160201b8311171561063857600080fd5b509092509050611594565b34801561064f57600080fd5b506106766004803603602081101561066657600080fd5b50356001600160a01b031661166c565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b3480156106b857600080fd5b506102b36116a8565b3480156106cd57600080fd5b506102b3600480360360208110156106e457600080fd5b50356001600160a01b0316611702565b34801561070057600080fd5b506102686117b5565b34801561071557600080fd5b506102b36004803603606081101561072c57600080fd5b50803590602081013590604001356117bb565b34801561074b57600080fd5b50610268611823565b34801561076057600080fd5b506102b36004803603604081101561077757600080fd5b5080359060200135611829565b34801561079057600080fd5b50610268611a44565b6102b360048036036102a08110156107b057600080fd5b50803590602081013590604081013590606081013590608081019061016001611a4a565b6102b3600480360360608110156107ea57600080fd5b5080359060208101359060400135611afa565b34801561080957600080fd5b5061082d6004803603604081101561082057600080fd5b5080359060200135611b91565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610869578181015183820152602001610851565b505050509050019250505060405180910390f35b34801561088957600080fd5b506102b360048036036102808110156108a157600080fd5b50803590602081013590604081013590606081019061014001611c59565b6102b3600480360360208110156108d557600080fd5b50356001600160a01b0316611d19565b3480156108f157600080fd5b506102b36004803603602081101561090857600080fd5b810190602081018135600160201b81111561092257600080fd5b82018360208201111561093457600080fd5b803590602001918460018302840111600160201b8311171561095557600080fd5b509092509050611d53565b34801561096c57600080fd5b506102b3611d99565b34801561098157600080fd5b506102b3600480360360c081101561099857600080fd5b81359190810190604081016020820135600160201b8111156109b957600080fd5b8201836020820111156109cb57600080fd5b803590602001918460018302840111600160201b831117156109ec57600080fd5b919390928235926001600160a01b03602082013516926040820135929091608081019060600135600160201b811115610a2457600080fd5b820183602082011115610a3657600080fd5b803590602001918460018302840111600160201b83111715610a5757600080fd5b509092509050611de3565b348015610a6e57600080fd5b50610268611ef9565b348015610a8357600080fd5b506102df611eff565b348015610a9857600080fd5b50610268611f0e565b348015610aad57600080fd5b506102b360048036036020811015610ac457600080fd5b5035611f14565b348015610ad757600080fd5b5061026861202c565b348015610aec57600080fd5b506102b360048036036040811015610b0357600080fd5b506001600160a01b0381358116916020013516612032565b600a5460009081805b82811015610bd6576000600a8281548110610b3b57fe5b60009182526020918290206002909102018054604080516348b4573960e11b81526001600160a01b039283166004820152905192945090891692639168ae7292602480840193829003018186803b158015610b9557600080fd5b505afa158015610ba9573d6000803e3d6000fd5b505050506040513d6020811015610bbf57600080fd5b505115610bcd576001909201915b50600101610b24565b509392505050565b610be66116a8565b610bee611d99565b6004548211610c37576040805162461bcd60e51b815260206004820152601060248201526f535543434553534f525f544f5f4c4f5760801b604482015290519081900360640190fd5b600554821115610c82576040805162461bcd60e51b81526020600482015260116024820152700a6aa86868aa6a69ea4bea89ebe90928e9607b1b604482015290519081900360640190fd5b6001600160a01b038116600090815260096020526040902060030154600160a01b900460ff16610ce6576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000828152600660209081526040918290205460035483516311e7249560e21b815293516001600160a01b03909216939092849263479c9254926004808201939291829003018186803b158015610d3c57600080fd5b505afa158015610d50573d6000803e3d6000fd5b505050506040513d6020811015610d6657600080fd5b505114610daa576040805162461bcd60e51b815260206004820152600d60248201526c2120a22fa9aaa1a1a2a9a9a7a960991b604482015290519081900360640190fd5b806001600160a01b0316639168ae72836040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610df757600080fd5b505afa158015610e0b573d6000803e3d6000fd5b505050506040513d6020811015610e2157600080fd5b5051610e61576040805162461bcd60e51b815260206004820152600a6024820152692120a22fa9aa20a5a2a960b11b604482015290519081900360640190fd5b610e6b6000611f14565b6004546000908152600660205260409020546001600160a01b031680631a8a092b610e9582610b1b565b6040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610ec957600080fd5b505afa158015610edd573d6000803e3d6000fd5b50505050610eec600454612298565b50506004805460010190555050565b6006602052600090815260409020546001600160a01b031681565b336000908152600960205260409020610f2e8161231a565b6000610f386113c2565b905080826002015411610f4a57600080fd5b600282015481900383811115610f5d5750825b604051339082156108fc029083906000818181858888f19350505050158015610f8a573d6000803e3d6000fd5b5050505050565b333214610fd3576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600080610ffe600333868660405180838380828437604051920182900390912093506123b192505050565b60408051600381523360208201528151939550919350839285927f852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab92908290030190a350505050565b6008818154811061105457fe5b6000918252602090912001546001600160a01b0316905081565b6110766116a8565b61107e611d99565b6004546000908152600660205260408120546001600160a01b0316906110a390611f14565b806001600160a01b0316636cf00e7e6110bb83610b1b565b600880549050016003546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561110057600080fd5b505afa158015611114573d6000803e3d6000fd5b50505050600061118a86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506123ee92505050565b9050611196818861226c565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b1580156111cf57600080fd5b505afa1580156111e3573d6000803e3d6000fd5b505050506040513d60208110156111f957600080fd5b50511461123c576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b6112ac86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506124ee92505050565b6112b7600354612298565b60048054600381905560010190556040805188815290517fe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e849181900360200190a150505050505050565b6001600160a01b03811660009081526009602052604090206113228161231a565b60020180543401905550565b600b5481565b61133c6116a8565b60048054600090815260066020526040808220546003548251631422135960e11b81529485015290516001600160a01b0390911692839263284426b292602480840193829003018186803b15801561139357600080fd5b505afa1580156113a7573d6000803e3d6000fd5b505050506113b6600454612298565b50600480546001019055565b600354600090815260066020908152604080832054815163176fda1560e11b81529151600019936001600160a01b0390921692632edfb42a926004808301939192829003018186803b15801561141757600080fd5b505afa15801561142b573d6000803e3d6000fd5b505050506040513d602081101561144157600080fd5b5051431015611454575050600d54611530565b600354600090815260066020908152604080832054815163176fda1560e11b815291516001600160a01b0390911692632edfb42a9260048082019391829003018186803b1580156114a457600080fd5b505afa1580156114b8573d6000803e3d6000fd5b505050506040513d60208110156114ce57600080fd5b5051600b544391909103915060009082816114e557fe5b04905060ff8111156114f5575060ff5b600019600282900a0180611507575060015b600d54848161151257fe5b048111156115265783945050505050611530565b600d540293505050505b90565b600e546001600160a01b031681565b61156f8888888860405180608001604052808a8152602001898152602001888152602001878152506125b8565b5050505050505050565b600f546001600160a01b031681565b600c5481565b60035481565b61159d33612b67565b6115ee576040805162461bcd60e51b815260206004820152601a60248201527f6d7573742062652063616c6c656420627920636f6e7472616374000000000000604482015290519081900360640190fd5b61163a6005338787878787604051602001808681526020018581526020018481526020018383808284378083019250505095505050505050604051602081830303815290604052612ba3565b60405133907f49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f90600090a25050505050565b6009602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6003546004541180156116bf575060055460045411155b611700576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b565b6001600160a01b038116600090815260096020526040902060035460018201541115611762576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b61176b8161231a565b600281015461177982612c7a565b6040516001600160a01b0384169082156108fc029083906000818181858888f193505050501580156117af573d6000803e3d6000fd5b50505050565b600d5481565b3360009081526009602052604090206003810154600160a01b900460ff16611817576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6117af84848484612db4565b60055481565b600a54821115611871576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b6000600a838154811061188057fe5b9060005260206000209060020201905060008160010154905060005b600454821180156118ac57508381105b1561199057600082815260066020526040808220548554825163025aa7f760e61b81526001600160a01b039182166004820152925191169283926396a9fdc0926024808301939282900301818387803b15801561190857600080fd5b505af115801561191c573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561195957600080fd5b505afa15801561196d573d6000803e3d6000fd5b505050506040513d602081101561198357600080fd5b505192505060010161189c565b600454821015611a3857600a805460001981019081106119ac57fe5b9060005260206000209060020201600a86815481106119c757fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a805480611a0a57fe5b60008281526020812060026000199093019283020180546001600160a01b0319168155600101559055610f8a565b50600191909101555050565b60015481565b6000611a54612f34565b90506003548414611a95576040805162461bcd60e51b81526020600480830191909152602482015263282922ab60e11b604482015290519081900360640190fd5b611af187878784876007806020026040519081016040528092919082600760200280828437600092019190915250506040805161014081810190925291508990600a9083908390808284376000920191909152506130b1915050565b50505050505050565b6000611b04612f34565b6003546000848152600660209081526040918290205482516311e7249560e21b8152925194955092936001600160a01b039093169263479c9254926004808201939291829003018186803b158015611b5b57600080fd5b505afa158015611b6f573d6000803e3d6000fd5b505050506040513d6020811015611b8557600080fd5b50511461181757600080fd5b600854606090838301811115611ba657508282015b60608167ffffffffffffffff81118015611bbf57600080fd5b50604051908082528060200260200182016040528015611be9578160200160208202803683370190505b50905060005b82811015611c5057600881870181548110611c0657fe5b9060005260206000200160009054906101000a90046001600160a01b0316828281518110611c3057fe5b6001600160a01b0390921660209283029190910190910152600101611bef565b50949350505050565b3360009081526009602052604090206003810154600160a01b900460ff16611cb5576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b611d1186868684876007806020026040519081016040528092919082600760200280828437600092019190915250506040805161014081810190925291508990600a9083908390808284376000920191909152506130b1915050565b505050505050565b604080516001600160a01b038316602082015234818301528151808203830181526060909101909152611d50906000903390612ba3565b50565b611d9560033384848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612ba392505050565b5050565b600b5460075443031015611700576040805162461bcd60e51b815260206004820152600c60248201526b524543454e545f5354414b4560a01b604482015290519081900360640190fd5b60008460601b60601c6001600160a01b03168484846040516020018085815260200184815260200183838082843780830192505050945050505050604051602081830303815290604052805190602001209050611e7a8989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b925086915061326e9050565b6000856001600160a01b0316858585604051808383808284376040519201945060009350909150508083038185875af1925050503d8060008114611eda576040519150601f19603f3d011682016040523d82523d6000602084013e611edf565b606091505b5050905080611eed57600080fd5b50505050505050505050565b60045481565b6010546001600160a01b031681565b60085490565b600a54815b81811015612027576000600a8281548110611f3057fe5b906000526020600020906002020190505b6004548160010154108015611f565750600083115b1561201e57600a6001840381548110611f6b57fe5b9060005260206000209060020201600a8381548110611f8657fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a805480611fc957fe5b60008281526020812060026000199093019283020180546001600160a01b0319168155600101559055600a80548390811061200057fe5b90600052602060002090600202019050828060019003935050611f41565b50600101611f19565b505050565b60005481565b6001600160a01b038083166000908152600960205260408082208484168352912060038201549192909116331461206857600080fd5b60038101546001600160a01b0316331461208157600080fd5b8160020154816002015411156120e35760028083015490820154604051919003906001600160a01b0385169082156108fc029083906000818181858888f193505050501580156120d5573d6000803e3d6000fd5b506002820180549190910390555b60028181015483820180549183900490910190556003830180546001600160a01b0319908116909155604080518082019091526001600160a01b03868116825260018086015460208401908152600a80549283018155600052925194027fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a88101805495909216949093169390931790925590517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a9909101556117af81612c7a565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008181526006602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b1580156122e457600080fd5b505af11580156122f8573d6000803e3d6000fd5b50505060009182525060066020526040902080546001600160a01b0319169055565b6003810154600160a01b900460ff16612367576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b60038101546001600160a01b031615611d50576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001546000805490918291826123cb88884342878b6121fe565b90506123d7828261226c565b600055506001828101905590969095509350505050565b80518251600091829182805b838110156124a157600087828151811061241057fe5b60200260200101519050838187011115612460576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868a01810182902060408051808401969096528581019190915280518086038201815260609095019052835193019290922091909401936001016123fa565b508184146124e4576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b9695505050505050565b80516000805b82811015610f8a57600060ff1685838151811061250d57fe5b016020015160f81c141561259357600061252a866001850161330c565b905060028160405161253b90613b39565b90815260405190819003602001906000f08015801561255e573d6000803e3d6000fd5b5081546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b03909216919091179055505b83818151811061259f57fe5b60200260200101518201915080806001019150506124f4565b8184106125fa576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b600554821115612640576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b836003541061268a576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b600084815260066020908152604080832054858452928190205481516311e7249560e21b815291516001600160a01b039485169490911692839263479c92549260048083019392829003018186803b1580156126e557600080fd5b505afa1580156126f9573d6000803e3d6000fd5b505050506040513d602081101561270f57600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b15801561275157600080fd5b505afa158015612765573d6000803e3d6000fd5b505050506040513d602081101561277b57600080fd5b5051146127bb576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b6001600160a01b03808816600090815260096020526040808220928816825290206127e58261231a565b6127ee8161231a565b836001600160a01b0316639168ae728a6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561283b57600080fd5b505afa15801561284f573d6000803e3d6000fd5b505050506040513d602081101561286557600080fd5b50516128ad576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b826001600160a01b0316639168ae72886040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156128fa57600080fd5b505afa15801561290e573d6000803e3d6000fd5b505050506040513d602081101561292457600080fd5b505161296c576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b6129888560000151866020015187604001518860600151613365565b846001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b1580156129c157600080fd5b505afa1580156129d5573d6000803e3d6000fd5b505050506040513d60208110156129eb57600080fd5b505114612a2b576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b600f5485516020808801516040808a015160608b0151600b54835163877c8c2b60e01b815260048101979097526024870194909452604486019190915260648501526001600160a01b038e811660848601528c811660a486015260c485019290925251600094919091169263877c8c2b9260e480830193919282900301818787803b158015612ab957600080fd5b505af1158015612acd573d6000803e3d6000fd5b505050506040513d6020811015612ae357600080fd5b5051600380850180546001600160a01b038085166001600160a01b03199283168117909355928601805490911682179055604080518e84168152928c16602084015282810191909152519192507f5356de01101f6e8d5aea55e44b91b527b2c4507b5263f1d5111579896b823735919081900360600190a150505050505050505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590612b9b57508115155b949350505050565b600080612bb8858585805190602001206123b1565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612c37578181015183820152602001612c1f565b50505050905090810190601f168015612c645780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b8054600880546000919083908110612c8e57fe5b600091825260209091200154600880546001600160a01b039092169250906000198101908110612cba57fe5b600091825260209091200154600880546001600160a01b039092169184908110612ce057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816009600060088581548110612d2057fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556008805480612d5057fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0392909216815260099091526040812081815560018101829055600281019190915560030180546001600160a81b03191690555050565b83834014612dff576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b612e08826133a3565b6000828152600660209081526040918290205482516311e7249560e21b815292516001600160a01b0390911692839263479c925492600480840193829003018186803b158015612e5757600080fd5b505afa158015612e6b573d6000803e3d6000fd5b505050506040513d6020811015612e8157600080fd5b5051600183015414612ecc576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b6040805163123334b760e11b815233600482015290516001600160a01b03831691632466696e91602480830192600092919082900301818387803b158015612f1357600080fd5b505af1158015612f27573d6000803e3d6000fd5b5050505050600101555050565b33600090815260096020526040812060030154600160a01b900460ff1615612f94576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b612f9c6113c2565b341015612fe3576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b506008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee381018054336001600160a01b031991821681179092556040805160a0810182529384526003805460208087019182523487850190815260006060890181815260808a018b81529882526009909352949094209651875590519686019690965590516002850155935193830180549251929091166001600160a01b039094169390931760ff60a01b1916600160a01b91151591909102179091554360075590565b858540146130fc576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b6005546001018414613140576040805162461bcd60e51b81526020600482015260086024820152674e4f44455f4e554d60c01b604482015290519081900360640190fd5b613148613b46565b61315283836133c0565b905060006131648286600101546134df565b9050806001600160a01b0316632466696e336040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156131b557600080fd5b505af11580156131c9573d6000803e3d6000fd5b5050505060055485600101819055506005547fb1012d13bcf3816d09c42da8f91e287015ea0f3c3d13bae725983c3c24b42ce185856040518083600760200280838360005b8381101561322657818101518382015260200161320e565b5050505090500182600a60200280838360005b83811015613251578181015183820152602001613239565b505050509050019250505060405180910390a25050505050505050565b600160001b81189050600061328784838560010161389b565b5090506002858154811061329757fe5b600091825260208220015460408051630ad0379b60e01b8152600481018590526024810187905290516001600160a01b0390921692630ad0379b9260448084019382900301818387803b1580156132ed57600080fd5b505af1158015613301573d6000803e3d6000fd5b505050505050505050565b6000816020018351101561335c576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b60045481101580156133b757506005548111155b611d5057600080fd5b6133c8613b46565b60408051610220810182528351815260208085015181830152855182840152850151606080830191909152848301516080808401919091529085015160a0808401919091529085015160c083015284015160e082015290840151610100820152610120810183600660200201518152602001836007600a811061344757fe5b602002015181526020018460036007811061345e57fe5b60200201518152602001836008600a811061347557fe5b602002015181526020018460046007811061348c57fe5b60200201518152602001836009600a81106134a357fe5b60200201518152602001846005600781106134ba57fe5b60200201518152602001846006600781106134d157fe5b602002015190529392505050565b600081815260066020908152604080832054815163380ed4c760e11b815291516001600160a01b0390911692839263701da98e9260048083019392829003018186803b15801561352e57600080fd5b505afa158015613542573d6000803e3d6000fd5b505050506040513d602081101561355857600080fd5b5051613563856139a8565b146135a7576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b83608001516001540384610120015111156135fa576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b6000816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561363557600080fd5b505afa158015613649573d6000803e3d6000fd5b505050506040513d602081101561365f57600080fd5b50518551600b54600c54929350439190910391600a909104908202818310156136bc576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b87608001518860e00151038861012001511015806136df57508088610140015110155b61371c576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b806004028861014001511115613765576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600b544301848110156137755750835b6000600c548a61014001518161378757fe5b04905080820191506000601060009054906101000a90046001600160a01b03166001600160a01b031663d45ab2b56137c18d6001546139e0565b6137d18e60015460005488613a2f565b6137da8f613abb565b8e886040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b15801561382e57600080fd5b505af1158015613842573d6000803e3d6000fd5b505050506040513d602081101561385857600080fd5b50516005805460010190819055600090815260066020526040902080546001600160a01b0319166001600160a01b0383161790559b9a5050505050505050505050565b60008080848160205b8851811161399a578089015193506020818a5103602001816138c257fe5b0491505b6000821180156138d95750600287066001145b80156138e757508160020a87115b156138ff5760029096046001908101969401936138c6565b6002870661394a57838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002878161394257fe5b04965061398c565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002878161398557fe5b0460010196505b6001909401936020016138a4565b509093505050935093915050565b60006139da826000015183602001518460400151856060015186608001518760a001518860c001518960e001516121a4565b92915050565b6000613a2843846101400151856020015101856102000151866101e001518761012001518860800151018861018001518960a0015101896101c001518a60c0015101896121a4565b9392505050565b600080613a5d8661012001518760800151870303876101200151886080015188030386896101e00151613365565b90506000613a9a876101200151886101200151613a828a6101e001516000801b61226c565b613a958b606001518c610100015161226c565b613365565b9050613ab08282613aaa8a613ad1565b87613365565b979650505050505050565b60006139da826101600151836101a0015161226c565b6101408101516101008201516000916139da918190613b019085613af781808080613365565b8860400151613365565b613a956000801b876101400151613b2e8961016001518a61018001518b6101a001518c6101c00151613365565b896102000151613365565b61013780613bd383390190565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091529056fe608060405234801561001057600080fd5b506040516101373803806101378339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319163317905560015560df806100586000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630ad0379b14602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135604f565b005b6000546001600160a01b03163314606557600080fd5b60008181526002602052604090205460ff1615608057600080fd5b6001548214608d57600080fd5b6000908152600260205260409020805460ff191660011790555056fea2646970667358221220ac2ea98d21448678a949baebc8abc30d5a0bb7897571455392215b23603f3c9e64736f6c634300060c0033a2646970667358221220293901b5c321753e9ab462236de8c2aa6e23bc94a89df82aae67456004e80aff64736f6c634300060c0033"

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupBin), backend, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _challengeFactory, _nodeFactory, _extraConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _Rollup.Contract.ArbGasSpeedLimitPerBlock(&_Rollup.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _Rollup.Contract.ArbGasSpeedLimitPerBlock(&_Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupSession) BaseStake() (*big.Int, error) {
	return _Rollup.Contract.BaseStake(&_Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupCallerSession) BaseStake() (*big.Int, error) {
	return _Rollup.Contract.BaseStake(&_Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupSession) ChallengeFactory() (common.Address, error) {
	return _Rollup.Contract.ChallengeFactory(&_Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _Rollup.Contract.ChallengeFactory(&_Rollup.CallOpts)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint256)
func (_Rollup *RollupCaller) ChallengePeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengePeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint256)
func (_Rollup *RollupSession) ChallengePeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ChallengePeriodBlocks(&_Rollup.CallOpts)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint256)
func (_Rollup *RollupCallerSession) ChallengePeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ChallengePeriodBlocks(&_Rollup.CallOpts)
}

// CheckNoRecentStake is a free data retrieval call binding the contract method 0xbe211c9a.
//
// Solidity: function checkNoRecentStake() view returns()
func (_Rollup *RollupCaller) CheckNoRecentStake(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "checkNoRecentStake")

	if err != nil {
		return err
	}

	return err

}

// CheckNoRecentStake is a free data retrieval call binding the contract method 0xbe211c9a.
//
// Solidity: function checkNoRecentStake() view returns()
func (_Rollup *RollupSession) CheckNoRecentStake() error {
	return _Rollup.Contract.CheckNoRecentStake(&_Rollup.CallOpts)
}

// CheckNoRecentStake is a free data retrieval call binding the contract method 0xbe211c9a.
//
// Solidity: function checkNoRecentStake() view returns()
func (_Rollup *RollupCallerSession) CheckNoRecentStake() error {
	return _Rollup.Contract.CheckNoRecentStake(&_Rollup.CallOpts)
}

// CheckUnresolved is a free data retrieval call binding the contract method 0x73f33b06.
//
// Solidity: function checkUnresolved() view returns()
func (_Rollup *RollupCaller) CheckUnresolved(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "checkUnresolved")

	if err != nil {
		return err
	}

	return err

}

// CheckUnresolved is a free data retrieval call binding the contract method 0x73f33b06.
//
// Solidity: function checkUnresolved() view returns()
func (_Rollup *RollupSession) CheckUnresolved() error {
	return _Rollup.Contract.CheckUnresolved(&_Rollup.CallOpts)
}

// CheckUnresolved is a free data retrieval call binding the contract method 0x73f33b06.
//
// Solidity: function checkUnresolved() view returns()
func (_Rollup *RollupCallerSession) CheckUnresolved() error {
	return _Rollup.Contract.CheckUnresolved(&_Rollup.CallOpts)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupCaller) CountStakedZombies(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "countStakedZombies", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _Rollup.Contract.CountStakedZombies(&_Rollup.CallOpts, node)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupCallerSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _Rollup.Contract.CountStakedZombies(&_Rollup.CallOpts, node)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupSession) FirstUnresolvedNode() (*big.Int, error) {
	return _Rollup.Contract.FirstUnresolvedNode(&_Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _Rollup.Contract.FirstUnresolvedNode(&_Rollup.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 startIndex, uint256 max) view returns(address[])
func (_Rollup *RollupCaller) GetStakers(opts *bind.CallOpts, startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getStakers", startIndex, max)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 startIndex, uint256 max) view returns(address[])
func (_Rollup *RollupSession) GetStakers(startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	return _Rollup.Contract.GetStakers(&_Rollup.CallOpts, startIndex, max)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 startIndex, uint256 max) view returns(address[])
func (_Rollup *RollupCallerSession) GetStakers(startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	return _Rollup.Contract.GetStakers(&_Rollup.CallOpts, startIndex, max)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Rollup *RollupCaller) InboxMaxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "inboxMaxCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Rollup *RollupSession) InboxMaxCount() (*big.Int, error) {
	return _Rollup.Contract.InboxMaxCount(&_Rollup.CallOpts)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Rollup *RollupCallerSession) InboxMaxCount() (*big.Int, error) {
	return _Rollup.Contract.InboxMaxCount(&_Rollup.CallOpts)
}

// InboxMaxValue is a free data retrieval call binding the contract method 0xefefa7e5.
//
// Solidity: function inboxMaxValue() view returns(bytes32)
func (_Rollup *RollupCaller) InboxMaxValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "inboxMaxValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxMaxValue is a free data retrieval call binding the contract method 0xefefa7e5.
//
// Solidity: function inboxMaxValue() view returns(bytes32)
func (_Rollup *RollupSession) InboxMaxValue() ([32]byte, error) {
	return _Rollup.Contract.InboxMaxValue(&_Rollup.CallOpts)
}

// InboxMaxValue is a free data retrieval call binding the contract method 0xefefa7e5.
//
// Solidity: function inboxMaxValue() view returns(bytes32)
func (_Rollup *RollupCallerSession) InboxMaxValue() ([32]byte, error) {
	return _Rollup.Contract.InboxMaxValue(&_Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupSession) LatestConfirmed() (*big.Int, error) {
	return _Rollup.Contract.LatestConfirmed(&_Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestConfirmed() (*big.Int, error) {
	return _Rollup.Contract.LatestConfirmed(&_Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupSession) LatestNodeCreated() (*big.Int, error) {
	return _Rollup.Contract.LatestNodeCreated(&_Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _Rollup.Contract.LatestNodeCreated(&_Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupSession) NodeFactory() (common.Address, error) {
	return _Rollup.Contract.NodeFactory(&_Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupCallerSession) NodeFactory() (common.Address, error) {
	return _Rollup.Contract.NodeFactory(&_Rollup.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_Rollup *RollupCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "nodes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_Rollup *RollupSession) Nodes(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.Nodes(&_Rollup.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_Rollup *RollupCallerSession) Nodes(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.Nodes(&_Rollup.CallOpts, arg0)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupSession) StakeToken() (common.Address, error) {
	return _Rollup.Contract.StakeToken(&_Rollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupCallerSession) StakeToken() (common.Address, error) {
	return _Rollup.Contract.StakeToken(&_Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupSession) StakerCount() (*big.Int, error) {
	return _Rollup.Contract.StakerCount(&_Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupCallerSession) StakerCount() (*big.Int, error) {
	return _Rollup.Contract.StakerCount(&_Rollup.CallOpts)
}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_Rollup *RollupCaller) StakerList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_Rollup *RollupSession) StakerList(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.StakerList(&_Rollup.CallOpts, arg0)
}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_Rollup *RollupCallerSession) StakerList(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.StakerList(&_Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0x729cfe3b.
//
// Solidity: function stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})

	outstruct.Index = out[0].(*big.Int)
	outstruct.LatestStakedNode = out[1].(*big.Int)
	outstruct.AmountStaked = out[2].(*big.Int)
	outstruct.CurrentChallenge = out[3].(common.Address)
	outstruct.IsStaked = out[4].(bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0x729cfe3b.
//
// Solidity: function stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0x729cfe3b.
//
// Solidity: function stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// AddStakeOnExistingNode is a paid mutator transaction binding the contract method 0x7ba3ca01.
//
// Solidity: function addStakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) returns()
func (_Rollup *RollupTransactor) AddStakeOnExistingNode(opts *bind.TransactOpts, blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addStakeOnExistingNode", blockHash, blockNumber, nodeNum)
}

// AddStakeOnExistingNode is a paid mutator transaction binding the contract method 0x7ba3ca01.
//
// Solidity: function addStakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) returns()
func (_Rollup *RollupSession) AddStakeOnExistingNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AddStakeOnExistingNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum)
}

// AddStakeOnExistingNode is a paid mutator transaction binding the contract method 0x7ba3ca01.
//
// Solidity: function addStakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) returns()
func (_Rollup *RollupTransactorSession) AddStakeOnExistingNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AddStakeOnExistingNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum)
}

// AddStakeOnNewNode is a paid mutator transaction binding the contract method 0xaf46618b.
//
// Solidity: function addStakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupTransactor) AddStakeOnNewNode(opts *bind.TransactOpts, blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addStakeOnNewNode", blockHash, blockNumber, nodeNum, assertionBytes32Fields, assertionIntFields)
}

// AddStakeOnNewNode is a paid mutator transaction binding the contract method 0xaf46618b.
//
// Solidity: function addStakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupSession) AddStakeOnNewNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AddStakeOnNewNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum, assertionBytes32Fields, assertionIntFields)
}

// AddStakeOnNewNode is a paid mutator transaction binding the contract method 0xaf46618b.
//
// Solidity: function addStakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupTransactorSession) AddStakeOnNewNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AddStakeOnNewNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum, assertionBytes32Fields, assertionIntFields)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupTransactor) AddToDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addToDeposit", stakerAddress)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupSession) AddToDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupTransactorSession) AddToDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "completeChallenge", winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupTransactorSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winningStaker, losingStaker)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0x396b8cbc.
//
// Solidity: function confirmNextNode(bytes32 logAcc, bytes sendsData, uint256[] sendLengths) returns()
func (_Rollup *RollupTransactor) ConfirmNextNode(opts *bind.TransactOpts, logAcc [32]byte, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "confirmNextNode", logAcc, sendsData, sendLengths)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0x396b8cbc.
//
// Solidity: function confirmNextNode(bytes32 logAcc, bytes sendsData, uint256[] sendLengths) returns()
func (_Rollup *RollupSession) ConfirmNextNode(logAcc [32]byte, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmNextNode(&_Rollup.TransactOpts, logAcc, sendsData, sendLengths)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0x396b8cbc.
//
// Solidity: function confirmNextNode(bytes32 logAcc, bytes sendsData, uint256[] sendLengths) returns()
func (_Rollup *RollupTransactorSession) ConfirmNextNode(logAcc [32]byte, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmNextNode(&_Rollup.TransactOpts, logAcc, sendsData, sendLengths)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x58aab3d3.
//
// Solidity: function createChallenge(address staker1Address, uint256 nodeNum1, address staker2Address, uint256 nodeNum2, bytes32 inboxConsistencyHash, bytes32 inboxDeltaHash, bytes32 executionHash, uint256 executionCheckTime) returns()
func (_Rollup *RollupTransactor) CreateChallenge(opts *bind.TransactOpts, staker1Address common.Address, nodeNum1 *big.Int, staker2Address common.Address, nodeNum2 *big.Int, inboxConsistencyHash [32]byte, inboxDeltaHash [32]byte, executionHash [32]byte, executionCheckTime *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "createChallenge", staker1Address, nodeNum1, staker2Address, nodeNum2, inboxConsistencyHash, inboxDeltaHash, executionHash, executionCheckTime)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x58aab3d3.
//
// Solidity: function createChallenge(address staker1Address, uint256 nodeNum1, address staker2Address, uint256 nodeNum2, bytes32 inboxConsistencyHash, bytes32 inboxDeltaHash, bytes32 executionHash, uint256 executionCheckTime) returns()
func (_Rollup *RollupSession) CreateChallenge(staker1Address common.Address, nodeNum1 *big.Int, staker2Address common.Address, nodeNum2 *big.Int, inboxConsistencyHash [32]byte, inboxDeltaHash [32]byte, executionHash [32]byte, executionCheckTime *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, staker1Address, nodeNum1, staker2Address, nodeNum2, inboxConsistencyHash, inboxDeltaHash, executionHash, executionCheckTime)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x58aab3d3.
//
// Solidity: function createChallenge(address staker1Address, uint256 nodeNum1, address staker2Address, uint256 nodeNum2, bytes32 inboxConsistencyHash, bytes32 inboxDeltaHash, bytes32 executionHash, uint256 executionCheckTime) returns()
func (_Rollup *RollupTransactorSession) CreateChallenge(staker1Address common.Address, nodeNum1 *big.Int, staker2Address common.Address, nodeNum2 *big.Int, inboxConsistencyHash [32]byte, inboxDeltaHash [32]byte, executionHash [32]byte, executionCheckTime *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, staker1Address, nodeNum1, staker2Address, nodeNum2, inboxConsistencyHash, inboxDeltaHash, executionHash, executionCheckTime)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Rollup *RollupTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Rollup *RollupSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.DeployL2ContractPair(&_Rollup.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Rollup *RollupTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.DeployL2ContractPair(&_Rollup.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Rollup *RollupTransactor) DepositEthMessage(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "depositEthMessage", to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Rollup *RollupSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.DepositEthMessage(&_Rollup.TransactOpts, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Rollup *RollupTransactorSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.DepositEthMessage(&_Rollup.TransactOpts, to)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Rollup *RollupTransactor) ExecuteTransaction(opts *bind.TransactOpts, outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "executeTransaction", outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Rollup *RollupSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ExecuteTransaction(&_Rollup.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Rollup *RollupTransactorSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ExecuteTransaction(&_Rollup.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// NewStakeOnExistingNode is a paid mutator transaction binding the contract method 0xad432faf.
//
// Solidity: function newStakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) payable returns()
func (_Rollup *RollupTransactor) NewStakeOnExistingNode(opts *bind.TransactOpts, blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "newStakeOnExistingNode", blockHash, blockNumber, nodeNum)
}

// NewStakeOnExistingNode is a paid mutator transaction binding the contract method 0xad432faf.
//
// Solidity: function newStakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) payable returns()
func (_Rollup *RollupSession) NewStakeOnExistingNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.NewStakeOnExistingNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum)
}

// NewStakeOnExistingNode is a paid mutator transaction binding the contract method 0xad432faf.
//
// Solidity: function newStakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) payable returns()
func (_Rollup *RollupTransactorSession) NewStakeOnExistingNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.NewStakeOnExistingNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum)
}

// NewStakeOnNewNode is a paid mutator transaction binding the contract method 0x9a4fcae7.
//
// Solidity: function newStakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, uint256 prev, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) payable returns()
func (_Rollup *RollupTransactor) NewStakeOnNewNode(opts *bind.TransactOpts, blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, prev *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "newStakeOnNewNode", blockHash, blockNumber, nodeNum, prev, assertionBytes32Fields, assertionIntFields)
}

// NewStakeOnNewNode is a paid mutator transaction binding the contract method 0x9a4fcae7.
//
// Solidity: function newStakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, uint256 prev, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) payable returns()
func (_Rollup *RollupSession) NewStakeOnNewNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, prev *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.NewStakeOnNewNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum, prev, assertionBytes32Fields, assertionIntFields)
}

// NewStakeOnNewNode is a paid mutator transaction binding the contract method 0x9a4fcae7.
//
// Solidity: function newStakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, uint256 prev, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) payable returns()
func (_Rollup *RollupTransactorSession) NewStakeOnNewNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, prev *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.NewStakeOnNewNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum, prev, assertionBytes32Fields, assertionIntFields)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 maxReduction) returns()
func (_Rollup *RollupTransactor) ReduceDeposit(opts *bind.TransactOpts, maxReduction *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "reduceDeposit", maxReduction)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 maxReduction) returns()
func (_Rollup *RollupSession) ReduceDeposit(maxReduction *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, maxReduction)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 maxReduction) returns()
func (_Rollup *RollupTransactorSession) ReduceDeposit(maxReduction *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, maxReduction)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x0e1ef04c.
//
// Solidity: function rejectNextNode(uint256 successorWithStake, address stakerAddress) returns()
func (_Rollup *RollupTransactor) RejectNextNode(opts *bind.TransactOpts, successorWithStake *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "rejectNextNode", successorWithStake, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x0e1ef04c.
//
// Solidity: function rejectNextNode(uint256 successorWithStake, address stakerAddress) returns()
func (_Rollup *RollupSession) RejectNextNode(successorWithStake *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNode(&_Rollup.TransactOpts, successorWithStake, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x0e1ef04c.
//
// Solidity: function rejectNextNode(uint256 successorWithStake, address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) RejectNextNode(successorWithStake *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNode(&_Rollup.TransactOpts, successorWithStake, stakerAddress)
}

// RejectNextNodeOutOfOrder is a paid mutator transaction binding the contract method 0x4802c739.
//
// Solidity: function rejectNextNodeOutOfOrder() returns()
func (_Rollup *RollupTransactor) RejectNextNodeOutOfOrder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "rejectNextNodeOutOfOrder")
}

// RejectNextNodeOutOfOrder is a paid mutator transaction binding the contract method 0x4802c739.
//
// Solidity: function rejectNextNodeOutOfOrder() returns()
func (_Rollup *RollupSession) RejectNextNodeOutOfOrder() (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNodeOutOfOrder(&_Rollup.TransactOpts)
}

// RejectNextNodeOutOfOrder is a paid mutator transaction binding the contract method 0x4802c739.
//
// Solidity: function rejectNextNodeOutOfOrder() returns()
func (_Rollup *RollupTransactorSession) RejectNextNodeOutOfOrder() (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNodeOutOfOrder(&_Rollup.TransactOpts)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupTransactor) RemoveOldZombies(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeOldZombies", startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldZombies(&_Rollup.TransactOpts, startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupTransactorSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldZombies(&_Rollup.TransactOpts, startIndex)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupTransactor) RemoveZombie(opts *bind.TransactOpts, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeZombie", zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveZombie(&_Rollup.TransactOpts, zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupTransactorSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveZombie(&_Rollup.TransactOpts, zombieNum, maxNodes)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupTransactor) ReturnOldDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "returnOldDeposit", stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ReturnOldDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ReturnOldDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Rollup *RollupTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Rollup *RollupSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2Message(&_Rollup.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Rollup *RollupTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2Message(&_Rollup.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Rollup *RollupTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Rollup *RollupSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2MessageFromOrigin(&_Rollup.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Rollup *RollupTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2MessageFromOrigin(&_Rollup.TransactOpts, messageData)
}

// RollupBuddyContractPairIterator is returned from FilterBuddyContractPair and is used to iterate over the raw logs and unpacked data for BuddyContractPair events raised by the Rollup contract.
type RollupBuddyContractPairIterator struct {
	Event *RollupBuddyContractPair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupBuddyContractPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBuddyContractPair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupBuddyContractPair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupBuddyContractPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBuddyContractPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBuddyContractPair represents a BuddyContractPair event raised by the Rollup contract.
type RollupBuddyContractPair struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractPair is a free log retrieval operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Rollup *RollupFilterer) FilterBuddyContractPair(opts *bind.FilterOpts, sender []common.Address) (*RollupBuddyContractPairIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return &RollupBuddyContractPairIterator{contract: _Rollup.contract, event: "BuddyContractPair", logs: logs, sub: sub}, nil
}

// WatchBuddyContractPair is a free log subscription operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Rollup *RollupFilterer) WatchBuddyContractPair(opts *bind.WatchOpts, sink chan<- *RollupBuddyContractPair, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBuddyContractPair)
				if err := _Rollup.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuddyContractPair is a log parse operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Rollup *RollupFilterer) ParseBuddyContractPair(log types.Log) (*RollupBuddyContractPair, error) {
	event := new(RollupBuddyContractPair)
	if err := _Rollup.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Rollup contract.
type RollupMessageDeliveredIterator struct {
	Event *RollupMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMessageDelivered represents a MessageDelivered event raised by the Rollup contract.
type RollupMessageDelivered struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Rollup *RollupFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*RollupMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &RollupMessageDeliveredIterator{contract: _Rollup.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Rollup *RollupFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *RollupMessageDelivered, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMessageDelivered)
				if err := _Rollup.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Rollup *RollupFilterer) ParseMessageDelivered(log types.Log) (*RollupMessageDelivered, error) {
	event := new(RollupMessageDelivered)
	if err := _Rollup.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMessageDeliveredFromOriginIterator is returned from FilterMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for MessageDeliveredFromOrigin events raised by the Rollup contract.
type RollupMessageDeliveredFromOriginIterator struct {
	Event *RollupMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMessageDeliveredFromOrigin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupMessageDeliveredFromOrigin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMessageDeliveredFromOrigin represents a MessageDeliveredFromOrigin event raised by the Rollup contract.
type RollupMessageDeliveredFromOrigin struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Rollup *RollupFilterer) FilterMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*RollupMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &RollupMessageDeliveredFromOriginIterator{contract: _Rollup.contract, event: "MessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Rollup *RollupFilterer) WatchMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *RollupMessageDeliveredFromOrigin, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMessageDeliveredFromOrigin)
				if err := _Rollup.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDeliveredFromOrigin is a log parse operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Rollup *RollupFilterer) ParseMessageDeliveredFromOrigin(log types.Log) (*RollupMessageDeliveredFromOrigin, error) {
	event := new(RollupMessageDeliveredFromOrigin)
	if err := _Rollup.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the Rollup contract.
type RollupNodeCreatedIterator struct {
	Event *RollupNodeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNodeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupNodeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNodeCreated represents a NodeCreated event raised by the Rollup contract.
type RollupNodeCreated struct {
	NodeNum                *big.Int
	AssertionBytes32Fields [7][32]byte
	AssertionIntFields     [10]*big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0xb1012d13bcf3816d09c42da8f91e287015ea0f3c3d13bae725983c3c24b42ce1.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields)
func (_Rollup *RollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int) (*RollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupNodeCreatedIterator{contract: _Rollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0xb1012d13bcf3816d09c42da8f91e287015ea0f3c3d13bae725983c3c24b42ce1.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields)
func (_Rollup *RollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *RollupNodeCreated, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNodeCreated)
				if err := _Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeCreated is a log parse operation binding the contract event 0xb1012d13bcf3816d09c42da8f91e287015ea0f3c3d13bae725983c3c24b42ce1.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields)
func (_Rollup *RollupFilterer) ParseNodeCreated(log types.Log) (*RollupNodeCreated, error) {
	event := new(RollupNodeCreated)
	if err := _Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the Rollup contract.
type RollupRollupChallengeStartedIterator struct {
	Event *RollupRollupChallengeStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollupChallengeStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRollupChallengeStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the Rollup contract.
type RollupRollupChallengeStarted struct {
	Asserter          common.Address
	Challenger        common.Address
	ChallengeContract common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0x5356de01101f6e8d5aea55e44b91b527b2c4507b5263f1d5111579896b823735.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, address challengeContract)
func (_Rollup *RollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts) (*RollupRollupChallengeStartedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return &RollupRollupChallengeStartedIterator{contract: _Rollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0x5356de01101f6e8d5aea55e44b91b527b2c4507b5263f1d5111579896b823735.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, address challengeContract)
func (_Rollup *RollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *RollupRollupChallengeStarted) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollupChallengeStarted)
				if err := _Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0x5356de01101f6e8d5aea55e44b91b527b2c4507b5263f1d5111579896b823735.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, address challengeContract)
func (_Rollup *RollupFilterer) ParseRollupChallengeStarted(log types.Log) (*RollupRollupChallengeStarted, error) {
	event := new(RollupRollupChallengeStarted)
	if err := _Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupSentLogsIterator is returned from FilterSentLogs and is used to iterate over the raw logs and unpacked data for SentLogs events raised by the Rollup contract.
type RollupSentLogsIterator struct {
	Event *RollupSentLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupSentLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupSentLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupSentLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupSentLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupSentLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupSentLogs represents a SentLogs event raised by the Rollup contract.
type RollupSentLogs struct {
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSentLogs is a free log retrieval operation binding the contract event 0xe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e84.
//
// Solidity: event SentLogs(bytes32 logsAccHash)
func (_Rollup *RollupFilterer) FilterSentLogs(opts *bind.FilterOpts) (*RollupSentLogsIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "SentLogs")
	if err != nil {
		return nil, err
	}
	return &RollupSentLogsIterator{contract: _Rollup.contract, event: "SentLogs", logs: logs, sub: sub}, nil
}

// WatchSentLogs is a free log subscription operation binding the contract event 0xe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e84.
//
// Solidity: event SentLogs(bytes32 logsAccHash)
func (_Rollup *RollupFilterer) WatchSentLogs(opts *bind.WatchOpts, sink chan<- *RollupSentLogs) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "SentLogs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupSentLogs)
				if err := _Rollup.contract.UnpackLog(event, "SentLogs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSentLogs is a log parse operation binding the contract event 0xe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e84.
//
// Solidity: event SentLogs(bytes32 logsAccHash)
func (_Rollup *RollupFilterer) ParseSentLogs(log types.Log) (*RollupSentLogs, error) {
	event := new(RollupSentLogs)
	if err := _Rollup.contract.UnpackLog(event, "SentLogs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorABI is the input ABI used to generate the binding from.
const RollupCreatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"name\":\"createRollup\",\"outputs\":[{\"internalType\":\"contractRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupCreatorFuncSigs maps the 4-byte function signature to its string representation.
var RollupCreatorFuncSigs = map[string]string{
	"d798f5be": "createRollup(bytes32,uint256,uint256,uint256,address,address,bytes)",
}

// RollupCreatorBin is the compiled bytecode used for deploying new contracts.
var RollupCreatorBin = "0x608060405234801561001057600080fd5b506040516145ea3803806145ea8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556145708061007a6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063d798f5be14610030575b600080fd5b6100d1600480360360e081101561004657600080fd5b8135916020810135916040820135916060810135916001600160a01b03608083013581169260a08101359091169181019060e0810160c082013564010000000081111561009257600080fd5b8201836020820111156100a457600080fd5b803590602001918460018302840111640100000000831117156100c657600080fd5b5090925090506100ed565b604080516001600160a01b039092168252519081900360200190f35b6000805460015460405183928c928c928c928c928c928c926001600160a01b039081169216908c908c9061012090610219565b808b81526020018a8152602001898152602001888152602001876001600160a01b03168152602001866001600160a01b03168152602001856001600160a01b03168152602001846001600160a01b031681526020018060200182810382528484828181526020019250808284376000838201819052604051601f909201601f19169093018190039e509c50909a5050505050505050505050f0801580156101cb573d6000803e3d6000fd5b50604080516001600160a01b038316815290519192507f84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c919081900360200190a19998505050505050505050565b614314806102278339019056fe60806040523480156200001157600080fd5b50604051620043143803806200431483398181016040526101208110156200003857600080fd5b815160208301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0180519751999b989a969995989497939692959194919392820192846401000000008211156200009157600080fd5b908301906020820185811115620000a757600080fd5b8251640100000000811182820188101715620000c257600080fd5b82525081516020918201929091019080838360005b83811015620000f1578181015183820152602001620000d7565b50505050905090810190601f1680156200011f5780820380516001836020036101000a031916815260200191505b506040525050600f80546001600160a01b03199081166001600160a01b0387811691909117909255601080549091169185169190911790555060006200017843828c81808080806200036e602090811b620021a417901c565b6010546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905260848201819052915193945090926001600160a01b039092169163d45ab2b59160a48082019260209290919082900301818787803b158015620001e957600080fd5b505af1158015620001fe573d6000803e3d6000fd5b505050506040513d60208110156200021557600080fd5b81019080805190602001909291905050509050806006600080815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555089600b8190555088600c8190555087600d8190555086600e60006101000a8154816001600160a01b0302191690836001600160a01b03160217905550620003578a8a8a8a60601b6001600160601b0319168a60601b6001600160601b031916886040516020018087815260200186815260200185815260200184815260200183815260200182805190602001908083835b60208310620003125780518252601f199092019160209182019101620002f1565b6001836020036101000a0380198251168184511680821785525050505050509050019650505050505050604051602081830303815290604052620003c860201b60201c565b5050600160045550620005c5975050505050505050565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b620003d660043083620003d9565b50565b600080620003f685858580519060200120620004bb60201b60201c565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620004775781810151838201526020016200045d565b50505050905090810190601f168015620004a55780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b600080600060015490506000805490506000620004e888884342878b6200051860201b620021fe1760201c565b90506200050182826200059960201b6200226c1760201c565b600055506001828101905590969095509350505050565b6040805160f89790971b7fff000000000000000000000000000000000000000000000000000000000000001660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b613d3f80620005d56000396000f3fe6080604052600436106102305760003560e01c80637427be511161012e578063afcc220b116100ab578063d93fe9c41161006f578063d93fe9c414610a77578063dff6978714610a8c578063edfd03ed14610aa1578063efefa7e514610acb578063fa7803e614610ae057610230565b8063afcc220b146108bf578063b75436bb146108e5578063be211c9a14610960578063c4fb000c14610975578063d735e21d14610a6257610230565b8063917cae02116100f2578063917cae02146107845780639a4fcae714610799578063ad432faf146107d4578063ad71bd36146107fd578063af46618b1461087d57610230565b80637427be51146106c157806376e7e23b146106f45780637ba3ca01146107095780637ba9534a1461073f5780637e2d21551461075457610230565b80634802c739116101bc5780635e8ef106116101805780635e8ef1061461058b57806365f7f80d146105a05780636f5dfdca146105b5578063729cfe3b1461064357806373f33b06146106ac57610230565b80634802c739146104d75780634d26732d146104ec57806351ed6a301461050157806358aab3d3146105165780635dbaf68b1461057657610230565b80631fe927cf116102035780631fe927cf14610325578063348e50c6146103a0578063396b8cbc146103ca57806345c5b2c71461049c57806346c2781a146104c257610230565b806304a28064146102355780630e1ef04c1461027a5780631c53c280146102b55780631e83d30f146102fb575b600080fd5b34801561024157600080fd5b506102686004803603602081101561025857600080fd5b50356001600160a01b0316610b1b565b60408051918252519081900360200190f35b34801561028657600080fd5b506102b36004803603604081101561029d57600080fd5b50803590602001356001600160a01b0316610bde565b005b3480156102c157600080fd5b506102df600480360360208110156102d857600080fd5b5035610efb565b604080516001600160a01b039092168252519081900360200190f35b34801561030757600080fd5b506102b36004803603602081101561031e57600080fd5b5035610f16565b34801561033157600080fd5b506102b36004803603602081101561034857600080fd5b810190602081018135600160201b81111561036257600080fd5b82018360208201111561037457600080fd5b803590602001918460018302840111600160201b8311171561039557600080fd5b509092509050610f91565b3480156103ac57600080fd5b506102df600480360360208110156103c357600080fd5b5035611047565b3480156103d657600080fd5b506102b3600480360360608110156103ed57600080fd5b81359190810190604081016020820135600160201b81111561040e57600080fd5b82018360208201111561042057600080fd5b803590602001918460018302840111600160201b8311171561044157600080fd5b919390929091602081019035600160201b81111561045e57600080fd5b82018360208201111561047057600080fd5b803590602001918460208302840111600160201b8311171561049157600080fd5b50909250905061106e565b6102b3600480360360208110156104b257600080fd5b50356001600160a01b0316611301565b3480156104ce57600080fd5b5061026861132e565b3480156104e357600080fd5b506102b3611334565b3480156104f857600080fd5b506102686113c2565b34801561050d57600080fd5b506102df611533565b34801561052257600080fd5b506102b3600480360361010081101561053a57600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060808101359060a08101359060c08101359060e00135611542565b34801561058257600080fd5b506102df611579565b34801561059757600080fd5b50610268611588565b3480156105ac57600080fd5b5061026861158e565b3480156105c157600080fd5b506102b3600480360360808110156105d857600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561060557600080fd5b82018360208201111561061757600080fd5b803590602001918460018302840111600160201b8311171561063857600080fd5b509092509050611594565b34801561064f57600080fd5b506106766004803603602081101561066657600080fd5b50356001600160a01b031661166c565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b3480156106b857600080fd5b506102b36116a8565b3480156106cd57600080fd5b506102b3600480360360208110156106e457600080fd5b50356001600160a01b0316611702565b34801561070057600080fd5b506102686117b5565b34801561071557600080fd5b506102b36004803603606081101561072c57600080fd5b50803590602081013590604001356117bb565b34801561074b57600080fd5b50610268611823565b34801561076057600080fd5b506102b36004803603604081101561077757600080fd5b5080359060200135611829565b34801561079057600080fd5b50610268611a44565b6102b360048036036102a08110156107b057600080fd5b50803590602081013590604081013590606081013590608081019061016001611a4a565b6102b3600480360360608110156107ea57600080fd5b5080359060208101359060400135611afa565b34801561080957600080fd5b5061082d6004803603604081101561082057600080fd5b5080359060200135611b91565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610869578181015183820152602001610851565b505050509050019250505060405180910390f35b34801561088957600080fd5b506102b360048036036102808110156108a157600080fd5b50803590602081013590604081013590606081019061014001611c59565b6102b3600480360360208110156108d557600080fd5b50356001600160a01b0316611d19565b3480156108f157600080fd5b506102b36004803603602081101561090857600080fd5b810190602081018135600160201b81111561092257600080fd5b82018360208201111561093457600080fd5b803590602001918460018302840111600160201b8311171561095557600080fd5b509092509050611d53565b34801561096c57600080fd5b506102b3611d99565b34801561098157600080fd5b506102b3600480360360c081101561099857600080fd5b81359190810190604081016020820135600160201b8111156109b957600080fd5b8201836020820111156109cb57600080fd5b803590602001918460018302840111600160201b831117156109ec57600080fd5b919390928235926001600160a01b03602082013516926040820135929091608081019060600135600160201b811115610a2457600080fd5b820183602082011115610a3657600080fd5b803590602001918460018302840111600160201b83111715610a5757600080fd5b509092509050611de3565b348015610a6e57600080fd5b50610268611ef9565b348015610a8357600080fd5b506102df611eff565b348015610a9857600080fd5b50610268611f0e565b348015610aad57600080fd5b506102b360048036036020811015610ac457600080fd5b5035611f14565b348015610ad757600080fd5b5061026861202c565b348015610aec57600080fd5b506102b360048036036040811015610b0357600080fd5b506001600160a01b0381358116916020013516612032565b600a5460009081805b82811015610bd6576000600a8281548110610b3b57fe5b60009182526020918290206002909102018054604080516348b4573960e11b81526001600160a01b039283166004820152905192945090891692639168ae7292602480840193829003018186803b158015610b9557600080fd5b505afa158015610ba9573d6000803e3d6000fd5b505050506040513d6020811015610bbf57600080fd5b505115610bcd576001909201915b50600101610b24565b509392505050565b610be66116a8565b610bee611d99565b6004548211610c37576040805162461bcd60e51b815260206004820152601060248201526f535543434553534f525f544f5f4c4f5760801b604482015290519081900360640190fd5b600554821115610c82576040805162461bcd60e51b81526020600482015260116024820152700a6aa86868aa6a69ea4bea89ebe90928e9607b1b604482015290519081900360640190fd5b6001600160a01b038116600090815260096020526040902060030154600160a01b900460ff16610ce6576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000828152600660209081526040918290205460035483516311e7249560e21b815293516001600160a01b03909216939092849263479c9254926004808201939291829003018186803b158015610d3c57600080fd5b505afa158015610d50573d6000803e3d6000fd5b505050506040513d6020811015610d6657600080fd5b505114610daa576040805162461bcd60e51b815260206004820152600d60248201526c2120a22fa9aaa1a1a2a9a9a7a960991b604482015290519081900360640190fd5b806001600160a01b0316639168ae72836040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610df757600080fd5b505afa158015610e0b573d6000803e3d6000fd5b505050506040513d6020811015610e2157600080fd5b5051610e61576040805162461bcd60e51b815260206004820152600a6024820152692120a22fa9aa20a5a2a960b11b604482015290519081900360640190fd5b610e6b6000611f14565b6004546000908152600660205260409020546001600160a01b031680631a8a092b610e9582610b1b565b6040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610ec957600080fd5b505afa158015610edd573d6000803e3d6000fd5b50505050610eec600454612298565b50506004805460010190555050565b6006602052600090815260409020546001600160a01b031681565b336000908152600960205260409020610f2e8161231a565b6000610f386113c2565b905080826002015411610f4a57600080fd5b600282015481900383811115610f5d5750825b604051339082156108fc029083906000818181858888f19350505050158015610f8a573d6000803e3d6000fd5b5050505050565b333214610fd3576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600080610ffe600333868660405180838380828437604051920182900390912093506123b192505050565b60408051600381523360208201528151939550919350839285927f852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab92908290030190a350505050565b6008818154811061105457fe5b6000918252602090912001546001600160a01b0316905081565b6110766116a8565b61107e611d99565b6004546000908152600660205260408120546001600160a01b0316906110a390611f14565b806001600160a01b0316636cf00e7e6110bb83610b1b565b600880549050016003546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561110057600080fd5b505afa158015611114573d6000803e3d6000fd5b50505050600061118a86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506123ee92505050565b9050611196818861226c565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b1580156111cf57600080fd5b505afa1580156111e3573d6000803e3d6000fd5b505050506040513d60208110156111f957600080fd5b50511461123c576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b6112ac86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506124ee92505050565b6112b7600354612298565b60048054600381905560010190556040805188815290517fe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e849181900360200190a150505050505050565b6001600160a01b03811660009081526009602052604090206113228161231a565b60020180543401905550565b600b5481565b61133c6116a8565b60048054600090815260066020526040808220546003548251631422135960e11b81529485015290516001600160a01b0390911692839263284426b292602480840193829003018186803b15801561139357600080fd5b505afa1580156113a7573d6000803e3d6000fd5b505050506113b6600454612298565b50600480546001019055565b600354600090815260066020908152604080832054815163176fda1560e11b81529151600019936001600160a01b0390921692632edfb42a926004808301939192829003018186803b15801561141757600080fd5b505afa15801561142b573d6000803e3d6000fd5b505050506040513d602081101561144157600080fd5b5051431015611454575050600d54611530565b600354600090815260066020908152604080832054815163176fda1560e11b815291516001600160a01b0390911692632edfb42a9260048082019391829003018186803b1580156114a457600080fd5b505afa1580156114b8573d6000803e3d6000fd5b505050506040513d60208110156114ce57600080fd5b5051600b544391909103915060009082816114e557fe5b04905060ff8111156114f5575060ff5b600019600282900a0180611507575060015b600d54848161151257fe5b048111156115265783945050505050611530565b600d540293505050505b90565b600e546001600160a01b031681565b61156f8888888860405180608001604052808a8152602001898152602001888152602001878152506125b8565b5050505050505050565b600f546001600160a01b031681565b600c5481565b60035481565b61159d33612b67565b6115ee576040805162461bcd60e51b815260206004820152601a60248201527f6d7573742062652063616c6c656420627920636f6e7472616374000000000000604482015290519081900360640190fd5b61163a6005338787878787604051602001808681526020018581526020018481526020018383808284378083019250505095505050505050604051602081830303815290604052612ba3565b60405133907f49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f90600090a25050505050565b6009602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6003546004541180156116bf575060055460045411155b611700576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b565b6001600160a01b038116600090815260096020526040902060035460018201541115611762576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b61176b8161231a565b600281015461177982612c7a565b6040516001600160a01b0384169082156108fc029083906000818181858888f193505050501580156117af573d6000803e3d6000fd5b50505050565b600d5481565b3360009081526009602052604090206003810154600160a01b900460ff16611817576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6117af84848484612db4565b60055481565b600a54821115611871576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b6000600a838154811061188057fe5b9060005260206000209060020201905060008160010154905060005b600454821180156118ac57508381105b1561199057600082815260066020526040808220548554825163025aa7f760e61b81526001600160a01b039182166004820152925191169283926396a9fdc0926024808301939282900301818387803b15801561190857600080fd5b505af115801561191c573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561195957600080fd5b505afa15801561196d573d6000803e3d6000fd5b505050506040513d602081101561198357600080fd5b505192505060010161189c565b600454821015611a3857600a805460001981019081106119ac57fe5b9060005260206000209060020201600a86815481106119c757fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a805480611a0a57fe5b60008281526020812060026000199093019283020180546001600160a01b0319168155600101559055610f8a565b50600191909101555050565b60015481565b6000611a54612f34565b90506003548414611a95576040805162461bcd60e51b81526020600480830191909152602482015263282922ab60e11b604482015290519081900360640190fd5b611af187878784876007806020026040519081016040528092919082600760200280828437600092019190915250506040805161014081810190925291508990600a9083908390808284376000920191909152506130b1915050565b50505050505050565b6000611b04612f34565b6003546000848152600660209081526040918290205482516311e7249560e21b8152925194955092936001600160a01b039093169263479c9254926004808201939291829003018186803b158015611b5b57600080fd5b505afa158015611b6f573d6000803e3d6000fd5b505050506040513d6020811015611b8557600080fd5b50511461181757600080fd5b600854606090838301811115611ba657508282015b60608167ffffffffffffffff81118015611bbf57600080fd5b50604051908082528060200260200182016040528015611be9578160200160208202803683370190505b50905060005b82811015611c5057600881870181548110611c0657fe5b9060005260206000200160009054906101000a90046001600160a01b0316828281518110611c3057fe5b6001600160a01b0390921660209283029190910190910152600101611bef565b50949350505050565b3360009081526009602052604090206003810154600160a01b900460ff16611cb5576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b611d1186868684876007806020026040519081016040528092919082600760200280828437600092019190915250506040805161014081810190925291508990600a9083908390808284376000920191909152506130b1915050565b505050505050565b604080516001600160a01b038316602082015234818301528151808203830181526060909101909152611d50906000903390612ba3565b50565b611d9560033384848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612ba392505050565b5050565b600b5460075443031015611700576040805162461bcd60e51b815260206004820152600c60248201526b524543454e545f5354414b4560a01b604482015290519081900360640190fd5b60008460601b60601c6001600160a01b03168484846040516020018085815260200184815260200183838082843780830192505050945050505050604051602081830303815290604052805190602001209050611e7a8989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b925086915061326e9050565b6000856001600160a01b0316858585604051808383808284376040519201945060009350909150508083038185875af1925050503d8060008114611eda576040519150601f19603f3d011682016040523d82523d6000602084013e611edf565b606091505b5050905080611eed57600080fd5b50505050505050505050565b60045481565b6010546001600160a01b031681565b60085490565b600a54815b81811015612027576000600a8281548110611f3057fe5b906000526020600020906002020190505b6004548160010154108015611f565750600083115b1561201e57600a6001840381548110611f6b57fe5b9060005260206000209060020201600a8381548110611f8657fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a805480611fc957fe5b60008281526020812060026000199093019283020180546001600160a01b0319168155600101559055600a80548390811061200057fe5b90600052602060002090600202019050828060019003935050611f41565b50600101611f19565b505050565b60005481565b6001600160a01b038083166000908152600960205260408082208484168352912060038201549192909116331461206857600080fd5b60038101546001600160a01b0316331461208157600080fd5b8160020154816002015411156120e35760028083015490820154604051919003906001600160a01b0385169082156108fc029083906000818181858888f193505050501580156120d5573d6000803e3d6000fd5b506002820180549190910390555b60028181015483820180549183900490910190556003830180546001600160a01b0319908116909155604080518082019091526001600160a01b03868116825260018086015460208401908152600a80549283018155600052925194027fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a88101805495909216949093169390931790925590517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a9909101556117af81612c7a565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008181526006602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b1580156122e457600080fd5b505af11580156122f8573d6000803e3d6000fd5b50505060009182525060066020526040902080546001600160a01b0319169055565b6003810154600160a01b900460ff16612367576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b60038101546001600160a01b031615611d50576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001546000805490918291826123cb88884342878b6121fe565b90506123d7828261226c565b600055506001828101905590969095509350505050565b80518251600091829182805b838110156124a157600087828151811061241057fe5b60200260200101519050838187011115612460576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868a01810182902060408051808401969096528581019190915280518086038201815260609095019052835193019290922091909401936001016123fa565b508184146124e4576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b9695505050505050565b80516000805b82811015610f8a57600060ff1685838151811061250d57fe5b016020015160f81c141561259357600061252a866001850161330c565b905060028160405161253b90613b39565b90815260405190819003602001906000f08015801561255e573d6000803e3d6000fd5b5081546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b03909216919091179055505b83818151811061259f57fe5b60200260200101518201915080806001019150506124f4565b8184106125fa576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b600554821115612640576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b836003541061268a576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b600084815260066020908152604080832054858452928190205481516311e7249560e21b815291516001600160a01b039485169490911692839263479c92549260048083019392829003018186803b1580156126e557600080fd5b505afa1580156126f9573d6000803e3d6000fd5b505050506040513d602081101561270f57600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b15801561275157600080fd5b505afa158015612765573d6000803e3d6000fd5b505050506040513d602081101561277b57600080fd5b5051146127bb576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b6001600160a01b03808816600090815260096020526040808220928816825290206127e58261231a565b6127ee8161231a565b836001600160a01b0316639168ae728a6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561283b57600080fd5b505afa15801561284f573d6000803e3d6000fd5b505050506040513d602081101561286557600080fd5b50516128ad576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b826001600160a01b0316639168ae72886040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156128fa57600080fd5b505afa15801561290e573d6000803e3d6000fd5b505050506040513d602081101561292457600080fd5b505161296c576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b6129888560000151866020015187604001518860600151613365565b846001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b1580156129c157600080fd5b505afa1580156129d5573d6000803e3d6000fd5b505050506040513d60208110156129eb57600080fd5b505114612a2b576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b600f5485516020808801516040808a015160608b0151600b54835163877c8c2b60e01b815260048101979097526024870194909452604486019190915260648501526001600160a01b038e811660848601528c811660a486015260c485019290925251600094919091169263877c8c2b9260e480830193919282900301818787803b158015612ab957600080fd5b505af1158015612acd573d6000803e3d6000fd5b505050506040513d6020811015612ae357600080fd5b5051600380850180546001600160a01b038085166001600160a01b03199283168117909355928601805490911682179055604080518e84168152928c16602084015282810191909152519192507f5356de01101f6e8d5aea55e44b91b527b2c4507b5263f1d5111579896b823735919081900360600190a150505050505050505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590612b9b57508115155b949350505050565b600080612bb8858585805190602001206123b1565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612c37578181015183820152602001612c1f565b50505050905090810190601f168015612c645780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b8054600880546000919083908110612c8e57fe5b600091825260209091200154600880546001600160a01b039092169250906000198101908110612cba57fe5b600091825260209091200154600880546001600160a01b039092169184908110612ce057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816009600060088581548110612d2057fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556008805480612d5057fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0392909216815260099091526040812081815560018101829055600281019190915560030180546001600160a81b03191690555050565b83834014612dff576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b612e08826133a3565b6000828152600660209081526040918290205482516311e7249560e21b815292516001600160a01b0390911692839263479c925492600480840193829003018186803b158015612e5757600080fd5b505afa158015612e6b573d6000803e3d6000fd5b505050506040513d6020811015612e8157600080fd5b5051600183015414612ecc576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b6040805163123334b760e11b815233600482015290516001600160a01b03831691632466696e91602480830192600092919082900301818387803b158015612f1357600080fd5b505af1158015612f27573d6000803e3d6000fd5b5050505050600101555050565b33600090815260096020526040812060030154600160a01b900460ff1615612f94576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b612f9c6113c2565b341015612fe3576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b506008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee381018054336001600160a01b031991821681179092556040805160a0810182529384526003805460208087019182523487850190815260006060890181815260808a018b81529882526009909352949094209651875590519686019690965590516002850155935193830180549251929091166001600160a01b039094169390931760ff60a01b1916600160a01b91151591909102179091554360075590565b858540146130fc576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b6005546001018414613140576040805162461bcd60e51b81526020600482015260086024820152674e4f44455f4e554d60c01b604482015290519081900360640190fd5b613148613b46565b61315283836133c0565b905060006131648286600101546134df565b9050806001600160a01b0316632466696e336040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156131b557600080fd5b505af11580156131c9573d6000803e3d6000fd5b5050505060055485600101819055506005547fb1012d13bcf3816d09c42da8f91e287015ea0f3c3d13bae725983c3c24b42ce185856040518083600760200280838360005b8381101561322657818101518382015260200161320e565b5050505090500182600a60200280838360005b83811015613251578181015183820152602001613239565b505050509050019250505060405180910390a25050505050505050565b600160001b81189050600061328784838560010161389b565b5090506002858154811061329757fe5b600091825260208220015460408051630ad0379b60e01b8152600481018590526024810187905290516001600160a01b0390921692630ad0379b9260448084019382900301818387803b1580156132ed57600080fd5b505af1158015613301573d6000803e3d6000fd5b505050505050505050565b6000816020018351101561335c576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b60045481101580156133b757506005548111155b611d5057600080fd5b6133c8613b46565b60408051610220810182528351815260208085015181830152855182840152850151606080830191909152848301516080808401919091529085015160a0808401919091529085015160c083015284015160e082015290840151610100820152610120810183600660200201518152602001836007600a811061344757fe5b602002015181526020018460036007811061345e57fe5b60200201518152602001836008600a811061347557fe5b602002015181526020018460046007811061348c57fe5b60200201518152602001836009600a81106134a357fe5b60200201518152602001846005600781106134ba57fe5b60200201518152602001846006600781106134d157fe5b602002015190529392505050565b600081815260066020908152604080832054815163380ed4c760e11b815291516001600160a01b0390911692839263701da98e9260048083019392829003018186803b15801561352e57600080fd5b505afa158015613542573d6000803e3d6000fd5b505050506040513d602081101561355857600080fd5b5051613563856139a8565b146135a7576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b83608001516001540384610120015111156135fa576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b6000816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561363557600080fd5b505afa158015613649573d6000803e3d6000fd5b505050506040513d602081101561365f57600080fd5b50518551600b54600c54929350439190910391600a909104908202818310156136bc576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b87608001518860e00151038861012001511015806136df57508088610140015110155b61371c576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b806004028861014001511115613765576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600b544301848110156137755750835b6000600c548a61014001518161378757fe5b04905080820191506000601060009054906101000a90046001600160a01b03166001600160a01b031663d45ab2b56137c18d6001546139e0565b6137d18e60015460005488613a2f565b6137da8f613abb565b8e886040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b15801561382e57600080fd5b505af1158015613842573d6000803e3d6000fd5b505050506040513d602081101561385857600080fd5b50516005805460010190819055600090815260066020526040902080546001600160a01b0319166001600160a01b0383161790559b9a5050505050505050505050565b60008080848160205b8851811161399a578089015193506020818a5103602001816138c257fe5b0491505b6000821180156138d95750600287066001145b80156138e757508160020a87115b156138ff5760029096046001908101969401936138c6565b6002870661394a57838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002878161394257fe5b04965061398c565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002878161398557fe5b0460010196505b6001909401936020016138a4565b509093505050935093915050565b60006139da826000015183602001518460400151856060015186608001518760a001518860c001518960e001516121a4565b92915050565b6000613a2843846101400151856020015101856102000151866101e001518761012001518860800151018861018001518960a0015101896101c001518a60c0015101896121a4565b9392505050565b600080613a5d8661012001518760800151870303876101200151886080015188030386896101e00151613365565b90506000613a9a876101200151886101200151613a828a6101e001516000801b61226c565b613a958b606001518c610100015161226c565b613365565b9050613ab08282613aaa8a613ad1565b87613365565b979650505050505050565b60006139da826101600151836101a0015161226c565b6101408101516101008201516000916139da918190613b019085613af781808080613365565b8860400151613365565b613a956000801b876101400151613b2e8961016001518a61018001518b6101a001518c6101c00151613365565b896102000151613365565b61013780613bd383390190565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091529056fe608060405234801561001057600080fd5b506040516101373803806101378339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319163317905560015560df806100586000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630ad0379b14602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135604f565b005b6000546001600160a01b03163314606557600080fd5b60008181526002602052604090205460ff1615608057600080fd5b6001548214608d57600080fd5b6000908152600260205260409020805460ff191660011790555056fea2646970667358221220ac2ea98d21448678a949baebc8abc30d5a0bb7897571455392215b23603f3c9e64736f6c634300060c0033a2646970667358221220293901b5c321753e9ab462236de8c2aa6e23bc94a89df82aae67456004e80aff64736f6c634300060c0033a26469706673582212202763f7457499496bab1f6794ee8e22be3164b8a2ed584a3f0bbaaddecac2747b64736f6c634300060c0033"

// DeployRollupCreator deploys a new Ethereum contract, binding an instance of RollupCreator to it.
func DeployRollupCreator(auth *bind.TransactOpts, backend bind.ContractBackend, _challengeFactory common.Address, _nodeFactory common.Address) (common.Address, *types.Transaction, *RollupCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCreatorBin), backend, _challengeFactory, _nodeFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// RollupCreator is an auto generated Go binding around an Ethereum contract.
type RollupCreator struct {
	RollupCreatorCaller     // Read-only binding to the contract
	RollupCreatorTransactor // Write-only binding to the contract
	RollupCreatorFilterer   // Log filterer for contract events
}

// RollupCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCreatorSession struct {
	Contract     *RollupCreator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCreatorCallerSession struct {
	Contract *RollupCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RollupCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCreatorTransactorSession struct {
	Contract     *RollupCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RollupCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCreatorRaw struct {
	Contract *RollupCreator // Generic contract binding to access the raw methods on
}

// RollupCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCreatorCallerRaw struct {
	Contract *RollupCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCreatorTransactorRaw struct {
	Contract *RollupCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCreator creates a new instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreator(address common.Address, backend bind.ContractBackend) (*RollupCreator, error) {
	contract, err := bindRollupCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// NewRollupCreatorCaller creates a new read-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorCaller(address common.Address, caller bind.ContractCaller) (*RollupCreatorCaller, error) {
	contract, err := bindRollupCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorCaller{contract: contract}, nil
}

// NewRollupCreatorTransactor creates a new write-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCreatorTransactor, error) {
	contract, err := bindRollupCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorTransactor{contract: contract}, nil
}

// NewRollupCreatorFilterer creates a new log filterer instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCreatorFilterer, error) {
	contract, err := bindRollupCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorFilterer{contract: contract}, nil
}

// bindRollupCreator binds a generic wrapper to an already deployed contract.
func bindRollupCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.RollupCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transact(opts, method, params...)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xd798f5be.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactor) CreateRollup(opts *bind.TransactOpts, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "createRollup", _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xd798f5be.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorSession) CreateRollup(_machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xd798f5be.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactorSession) CreateRollup(_machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// RollupCreatorRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupCreator contract.
type RollupCreatorRollupCreatedIterator struct {
	Event *RollupCreatorRollupCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupCreatorRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorRollupCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupCreatorRollupCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupCreatorRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorRollupCreated represents a RollupCreated event raised by the RollupCreator contract.
type RollupCreatorRollupCreated struct {
	RollupAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupCreatorRollupCreatedIterator, error) {

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupCreatorRollupCreatedIterator{contract: _RollupCreator.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupCreatorRollupCreated) (event.Subscription, error) {

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorRollupCreated)
				if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupCreated is a log parse operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) ParseRollupCreated(log types.Log) (*RollupCreatorRollupCreated, error) {
	event := new(RollupCreatorRollupCreated)
	if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupLibABI is the input ABI used to generate the binding from.
const RollupLibABI = "[]"

// RollupLibBin is the compiled bytecode used for deploying new contracts.
var RollupLibBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200941154b89717787b716ccdc7897ac010467da1cfc08b2046d7234ff04f88d3864736f6c634300060c0033"

// DeployRollupLib deploys a new Ethereum contract, binding an instance of RollupLib to it.
func DeployRollupLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupLib, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupLib{RollupLibCaller: RollupLibCaller{contract: contract}, RollupLibTransactor: RollupLibTransactor{contract: contract}, RollupLibFilterer: RollupLibFilterer{contract: contract}}, nil
}

// RollupLib is an auto generated Go binding around an Ethereum contract.
type RollupLib struct {
	RollupLibCaller     // Read-only binding to the contract
	RollupLibTransactor // Write-only binding to the contract
	RollupLibFilterer   // Log filterer for contract events
}

// RollupLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupLibSession struct {
	Contract     *RollupLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupLibCallerSession struct {
	Contract *RollupLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RollupLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupLibTransactorSession struct {
	Contract     *RollupLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RollupLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupLibRaw struct {
	Contract *RollupLib // Generic contract binding to access the raw methods on
}

// RollupLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupLibCallerRaw struct {
	Contract *RollupLibCaller // Generic read-only contract binding to access the raw methods on
}

// RollupLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupLibTransactorRaw struct {
	Contract *RollupLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupLib creates a new instance of RollupLib, bound to a specific deployed contract.
func NewRollupLib(address common.Address, backend bind.ContractBackend) (*RollupLib, error) {
	contract, err := bindRollupLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupLib{RollupLibCaller: RollupLibCaller{contract: contract}, RollupLibTransactor: RollupLibTransactor{contract: contract}, RollupLibFilterer: RollupLibFilterer{contract: contract}}, nil
}

// NewRollupLibCaller creates a new read-only instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibCaller(address common.Address, caller bind.ContractCaller) (*RollupLibCaller, error) {
	contract, err := bindRollupLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupLibCaller{contract: contract}, nil
}

// NewRollupLibTransactor creates a new write-only instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupLibTransactor, error) {
	contract, err := bindRollupLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupLibTransactor{contract: contract}, nil
}

// NewRollupLibFilterer creates a new log filterer instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupLibFilterer, error) {
	contract, err := bindRollupLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupLibFilterer{contract: contract}, nil
}

// bindRollupLib binds a generic wrapper to an already deployed contract.
func bindRollupLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupLib *RollupLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupLib.Contract.RollupLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupLib *RollupLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupLib.Contract.RollupLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupLib *RollupLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupLib.Contract.RollupLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupLib *RollupLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupLib *RollupLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupLib *RollupLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupLib.Contract.contract.Transact(opts, method, params...)
}
