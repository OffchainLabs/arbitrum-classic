// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengetester

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

// ChallengeTesterABI is the input ABI used to generate the binding from.
const ChallengeTesterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengeFactory_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeTemplate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cloneAddress\",\"type\":\"address\"}],\"name\":\"ChallengeInfo\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"asserterAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"challengerPeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"challengerDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"startChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeTesterFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeTesterFuncSigs = map[string]string{
	"6bc3cd22": "resolveChallenge(address,address,uint256)",
	"8f43ee32": "startChallenge(address,address,uint128,bytes32,uint256)",
}

// ChallengeTesterBin is the compiled bytecode used for deploying new contracts.
var ChallengeTesterBin = "0x608060405234801561001057600080fd5b506040516102073803806102078339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556101a2806100656000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636bc3cd221461003b5780638f43ee3214610073575b600080fd5b6100716004803603606081101561005157600080fd5b506001600160a01b038135811691602081013590911690604001356100be565b005b610071600480360360a081101561008957600080fd5b506001600160a01b0381358116916020810135909116906001600160801b0360408201351690606081013590608001356100c3565b505050565b600080546040805163432ed0e160e11b81526001600160a01b03898116600483015288811660248301526001600160801b038816604483015260648201879052608482018690529151919092169263865da1c29260a480820193602093909283900390910190829087803b15801561013a57600080fd5b505af115801561014e573d6000803e3d6000fd5b505050506040513d602081101561016457600080fd5b5050505050505056fea265627a7a72315820ce59b2f20905041d69be3a83804dbcb1fc83fb2331f31cb7d719060ec4c6c09b64736f6c634300050c0032"

// DeployChallengeTester deploys a new Ethereum contract, binding an instance of ChallengeTester to it.
func DeployChallengeTester(auth *bind.TransactOpts, backend bind.ContractBackend, challengeFactory_ common.Address) (common.Address, *types.Transaction, *ChallengeTester, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeTesterBin), backend, challengeFactory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeTester{ChallengeTesterCaller: ChallengeTesterCaller{contract: contract}, ChallengeTesterTransactor: ChallengeTesterTransactor{contract: contract}, ChallengeTesterFilterer: ChallengeTesterFilterer{contract: contract}}, nil
}

// ChallengeTester is an auto generated Go binding around an Ethereum contract.
type ChallengeTester struct {
	ChallengeTesterCaller     // Read-only binding to the contract
	ChallengeTesterTransactor // Write-only binding to the contract
	ChallengeTesterFilterer   // Log filterer for contract events
}

// ChallengeTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeTesterSession struct {
	Contract     *ChallengeTester  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeTesterCallerSession struct {
	Contract *ChallengeTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ChallengeTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTesterTransactorSession struct {
	Contract     *ChallengeTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChallengeTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeTesterRaw struct {
	Contract *ChallengeTester // Generic contract binding to access the raw methods on
}

// ChallengeTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeTesterCallerRaw struct {
	Contract *ChallengeTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTesterTransactorRaw struct {
	Contract *ChallengeTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeTester creates a new instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTester(address common.Address, backend bind.ContractBackend) (*ChallengeTester, error) {
	contract, err := bindChallengeTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeTester{ChallengeTesterCaller: ChallengeTesterCaller{contract: contract}, ChallengeTesterTransactor: ChallengeTesterTransactor{contract: contract}, ChallengeTesterFilterer: ChallengeTesterFilterer{contract: contract}}, nil
}

// NewChallengeTesterCaller creates a new read-only instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTesterCaller(address common.Address, caller bind.ContractCaller) (*ChallengeTesterCaller, error) {
	contract, err := bindChallengeTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTesterCaller{contract: contract}, nil
}

// NewChallengeTesterTransactor creates a new write-only instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTesterTransactor, error) {
	contract, err := bindChallengeTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTesterTransactor{contract: contract}, nil
}

// NewChallengeTesterFilterer creates a new log filterer instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeTesterFilterer, error) {
	contract, err := bindChallengeTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeTesterFilterer{contract: contract}, nil
}

// bindChallengeTester binds a generic wrapper to an already deployed contract.
func bindChallengeTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeTester *ChallengeTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeTester.Contract.ChallengeTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeTester *ChallengeTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ChallengeTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeTester *ChallengeTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ChallengeTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeTester *ChallengeTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeTester *ChallengeTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeTester *ChallengeTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeTester.Contract.contract.Transact(opts, method, params...)
}

// ResolveChallenge is a free data retrieval call binding the contract method 0x6bc3cd22.
//
// Solidity: function resolveChallenge(address winner, address loser, uint256 challengeType) constant returns()
func (_ChallengeTester *ChallengeTesterCaller) ResolveChallenge(opts *bind.CallOpts, winner common.Address, loser common.Address, challengeType *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _ChallengeTester.contract.Call(opts, out, "resolveChallenge", winner, loser, challengeType)
	return err
}

// ResolveChallenge is a free data retrieval call binding the contract method 0x6bc3cd22.
//
// Solidity: function resolveChallenge(address winner, address loser, uint256 challengeType) constant returns()
func (_ChallengeTester *ChallengeTesterSession) ResolveChallenge(winner common.Address, loser common.Address, challengeType *big.Int) error {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.CallOpts, winner, loser, challengeType)
}

// ResolveChallenge is a free data retrieval call binding the contract method 0x6bc3cd22.
//
// Solidity: function resolveChallenge(address winner, address loser, uint256 challengeType) constant returns()
func (_ChallengeTester *ChallengeTesterCallerSession) ResolveChallenge(winner common.Address, loser common.Address, challengeType *big.Int) error {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.CallOpts, winner, loser, challengeType)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x8f43ee32.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterTransactor) StartChallenge(opts *bind.TransactOpts, asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.contract.Transact(opts, "startChallenge", asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x8f43ee32.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.Contract.StartChallenge(&_ChallengeTester.TransactOpts, asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x8f43ee32.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterTransactorSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.Contract.StartChallenge(&_ChallengeTester.TransactOpts, asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// ChallengeTesterChallengeInfoIterator is returned from FilterChallengeInfo and is used to iterate over the raw logs and unpacked data for ChallengeInfo events raised by the ChallengeTester contract.
type ChallengeTesterChallengeInfoIterator struct {
	Event *ChallengeTesterChallengeInfo // Event containing the contract specifics and raw log

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
func (it *ChallengeTesterChallengeInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeTesterChallengeInfo)
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
		it.Event = new(ChallengeTesterChallengeInfo)
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
func (it *ChallengeTesterChallengeInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeTesterChallengeInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeTesterChallengeInfo represents a ChallengeInfo event raised by the ChallengeTester contract.
type ChallengeTesterChallengeInfo struct {
	ChallengeTemplate common.Address
	Nonce             *big.Int
	CodeHash          [32]byte
	CloneAddress      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChallengeInfo is a free log retrieval operation binding the contract event 0xf6894746d43f1ffb08f3b07337aed4e8f9d76fb78b92516ba96d89eb9976812a.
//
// Solidity: event ChallengeInfo(address challengeTemplate, uint256 nonce, bytes32 codeHash, address cloneAddress)
func (_ChallengeTester *ChallengeTesterFilterer) FilterChallengeInfo(opts *bind.FilterOpts) (*ChallengeTesterChallengeInfoIterator, error) {

	logs, sub, err := _ChallengeTester.contract.FilterLogs(opts, "ChallengeInfo")
	if err != nil {
		return nil, err
	}
	return &ChallengeTesterChallengeInfoIterator{contract: _ChallengeTester.contract, event: "ChallengeInfo", logs: logs, sub: sub}, nil
}

// WatchChallengeInfo is a free log subscription operation binding the contract event 0xf6894746d43f1ffb08f3b07337aed4e8f9d76fb78b92516ba96d89eb9976812a.
//
// Solidity: event ChallengeInfo(address challengeTemplate, uint256 nonce, bytes32 codeHash, address cloneAddress)
func (_ChallengeTester *ChallengeTesterFilterer) WatchChallengeInfo(opts *bind.WatchOpts, sink chan<- *ChallengeTesterChallengeInfo) (event.Subscription, error) {

	logs, sub, err := _ChallengeTester.contract.WatchLogs(opts, "ChallengeInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeTesterChallengeInfo)
				if err := _ChallengeTester.contract.UnpackLog(event, "ChallengeInfo", log); err != nil {
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

// ParseChallengeInfo is a log parse operation binding the contract event 0xf6894746d43f1ffb08f3b07337aed4e8f9d76fb78b92516ba96d89eb9976812a.
//
// Solidity: event ChallengeInfo(address challengeTemplate, uint256 nonce, bytes32 codeHash, address cloneAddress)
func (_ChallengeTester *ChallengeTesterFilterer) ParseChallengeInfo(log types.Log) (*ChallengeTesterChallengeInfo, error) {
	event := new(ChallengeTesterChallengeInfo)
	if err := _ChallengeTester.contract.UnpackLog(event, "ChallengeInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IChallengeFactoryABI is the input ABI used to generate the binding from.
const IChallengeFactoryABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"generateCloneAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IChallengeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeFactoryFuncSigs = map[string]string{
	"865da1c2": "createChallenge(address,address,uint256,bytes32,uint256)",
	"729406c8": "generateCloneAddress(address,address,uint256)",
}

// IChallengeFactory is an auto generated Go binding around an Ethereum contract.
type IChallengeFactory struct {
	IChallengeFactoryCaller     // Read-only binding to the contract
	IChallengeFactoryTransactor // Write-only binding to the contract
	IChallengeFactoryFilterer   // Log filterer for contract events
}

// IChallengeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeFactorySession struct {
	Contract     *IChallengeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeFactoryCallerSession struct {
	Contract *IChallengeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeFactoryTransactorSession struct {
	Contract     *IChallengeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeFactoryRaw struct {
	Contract *IChallengeFactory // Generic contract binding to access the raw methods on
}

// IChallengeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeFactoryCallerRaw struct {
	Contract *IChallengeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeFactoryTransactorRaw struct {
	Contract *IChallengeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeFactory creates a new instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactory(address common.Address, backend bind.ContractBackend) (*IChallengeFactory, error) {
	contract, err := bindIChallengeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactory{IChallengeFactoryCaller: IChallengeFactoryCaller{contract: contract}, IChallengeFactoryTransactor: IChallengeFactoryTransactor{contract: contract}, IChallengeFactoryFilterer: IChallengeFactoryFilterer{contract: contract}}, nil
}

// NewIChallengeFactoryCaller creates a new read-only instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryCaller(address common.Address, caller bind.ContractCaller) (*IChallengeFactoryCaller, error) {
	contract, err := bindIChallengeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryCaller{contract: contract}, nil
}

// NewIChallengeFactoryTransactor creates a new write-only instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeFactoryTransactor, error) {
	contract, err := bindIChallengeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryTransactor{contract: contract}, nil
}

// NewIChallengeFactoryFilterer creates a new log filterer instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeFactoryFilterer, error) {
	contract, err := bindIChallengeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryFilterer{contract: contract}, nil
}

// bindIChallengeFactory binds a generic wrapper to an already deployed contract.
func bindIChallengeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeFactory *IChallengeFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeFactory.Contract.IChallengeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeFactory *IChallengeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.IChallengeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeFactory *IChallengeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.IChallengeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeFactory *IChallengeFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeFactory *IChallengeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeFactory *IChallengeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.contract.Transact(opts, method, params...)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_IChallengeFactory *IChallengeFactoryCaller) GenerateCloneAddress(opts *bind.CallOpts, asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IChallengeFactory.contract.Call(opts, out, "generateCloneAddress", asserter, challenger, challengeType)
	return *ret0, err
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_IChallengeFactory *IChallengeFactorySession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _IChallengeFactory.Contract.GenerateCloneAddress(&_IChallengeFactory.CallOpts, asserter, challenger, challengeType)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_IChallengeFactory *IChallengeFactoryCallerSession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _IChallengeFactory.Contract.GenerateCloneAddress(&_IChallengeFactory.CallOpts, asserter, challenger, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_IChallengeFactory *IChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _IChallengeFactory.contract.Transact(opts, "createChallenge", _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_IChallengeFactory *IChallengeFactorySession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.CreateChallenge(&_IChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_IChallengeFactory *IChallengeFactoryTransactorSession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.CreateChallenge(&_IChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}
