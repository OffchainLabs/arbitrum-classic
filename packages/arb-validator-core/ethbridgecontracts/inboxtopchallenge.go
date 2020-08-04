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

// BisectionChallengeABI is the input ABI used to generate the binding from.
const BisectionChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BisectionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var BisectionChallengeFuncSigs = map[string]string{
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"ced5c1bf": "timeoutChallenge()",
}

// BisectionChallengeBin is the compiled bytecode used for deploying new contracts.
var BisectionChallengeBin = "0x608060405234801561001057600080fd5b50610882806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806302ad1e4e1461004657806379a9ad851461008a578063ced5c1bf1461013c575b600080fd5b610088600480360360a081101561005c57600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060800135610144565b005b610088600480360360808110156100a057600080fd5b813591908101906040810160208201356401000000008111156100c257600080fd5b8201836020820111156100d457600080fd5b803590602001918460018302840111640100000000831117156100f657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610159565b610088610457565b61015085858585610537565b60065550505050565b60055460ff16600281111561016a57fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b815250906102185760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156101dd5781810151838201526020016101c5565b50505050905090810190601f16801561020a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060035461022543610652565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b815250906102985760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b031633146103145760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b815250906103875760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b5061039783838387600101610659565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b815250906104055760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b50600681905561041361075a565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60035461046343610652565b116104b5576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156104c857fe5b1415610504576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a16104ff61076f565b610535565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a16105356107e3565b565b600060055460ff16600281111561054a57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906105bf5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff1916909117905561061761083b565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e80290565b600080838160205b8851811161074c578089015193506020818a51036020018161067f57fe5b0491505b6000821180156106965750600286066001145b80156106a457508160020a86115b156106b757600286046001019550610683565b600286066107025783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816106fa57fe5b049550610744565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161073d57fe5b0460010195505b602001610661565b505090941495945050505050565b6005805460ff1916600117905561053561083b565b600080546002546001546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b1580156107c757600080fd5b505af11580156107db573d6000803e3d6000fd5b503392505050ff5b600080546001546002546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b1580156107c757600080fd5b60045461084743610652565b0160035556fea265627a7a72315820244ba7aca16ee7858d2646f6d8d78dcf89787d6f0a5b7776114db87056cb9d6264736f6c63430005110032"

// DeployBisectionChallenge deploys a new Ethereum contract, binding an instance of BisectionChallenge to it.
func DeployBisectionChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BisectionChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BisectionChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BisectionChallenge{BisectionChallengeCaller: BisectionChallengeCaller{contract: contract}, BisectionChallengeTransactor: BisectionChallengeTransactor{contract: contract}, BisectionChallengeFilterer: BisectionChallengeFilterer{contract: contract}}, nil
}

// BisectionChallenge is an auto generated Go binding around an Ethereum contract.
type BisectionChallenge struct {
	BisectionChallengeCaller     // Read-only binding to the contract
	BisectionChallengeTransactor // Write-only binding to the contract
	BisectionChallengeFilterer   // Log filterer for contract events
}

// BisectionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BisectionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BisectionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BisectionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BisectionChallengeSession struct {
	Contract     *BisectionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BisectionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BisectionChallengeCallerSession struct {
	Contract *BisectionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BisectionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BisectionChallengeTransactorSession struct {
	Contract     *BisectionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BisectionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BisectionChallengeRaw struct {
	Contract *BisectionChallenge // Generic contract binding to access the raw methods on
}

// BisectionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BisectionChallengeCallerRaw struct {
	Contract *BisectionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// BisectionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BisectionChallengeTransactorRaw struct {
	Contract *BisectionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBisectionChallenge creates a new instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallenge(address common.Address, backend bind.ContractBackend) (*BisectionChallenge, error) {
	contract, err := bindBisectionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BisectionChallenge{BisectionChallengeCaller: BisectionChallengeCaller{contract: contract}, BisectionChallengeTransactor: BisectionChallengeTransactor{contract: contract}, BisectionChallengeFilterer: BisectionChallengeFilterer{contract: contract}}, nil
}

// NewBisectionChallengeCaller creates a new read-only instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallengeCaller(address common.Address, caller bind.ContractCaller) (*BisectionChallengeCaller, error) {
	contract, err := bindBisectionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeCaller{contract: contract}, nil
}

// NewBisectionChallengeTransactor creates a new write-only instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*BisectionChallengeTransactor, error) {
	contract, err := bindBisectionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeTransactor{contract: contract}, nil
}

// NewBisectionChallengeFilterer creates a new log filterer instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*BisectionChallengeFilterer, error) {
	contract, err := bindBisectionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeFilterer{contract: contract}, nil
}

// bindBisectionChallenge binds a generic wrapper to an already deployed contract.
func bindBisectionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BisectionChallenge *BisectionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BisectionChallenge.Contract.BisectionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BisectionChallenge *BisectionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.BisectionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BisectionChallenge *BisectionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.BisectionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BisectionChallenge *BisectionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BisectionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BisectionChallenge *BisectionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BisectionChallenge *BisectionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.contract.Transact(opts, method, params...)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_BisectionChallenge *BisectionChallengeTransactor) ChooseSegment(opts *bind.TransactOpts, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.contract.Transact(opts, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_BisectionChallenge *BisectionChallengeSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.ChooseSegment(&_BisectionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_BisectionChallenge *BisectionChallengeTransactorSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.ChooseSegment(&_BisectionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_BisectionChallenge *BisectionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.contract.Transact(opts, "initializeBisection", _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_BisectionChallenge *BisectionChallengeSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.InitializeBisection(&_BisectionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_BisectionChallenge *BisectionChallengeTransactorSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.InitializeBisection(&_BisectionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_BisectionChallenge *BisectionChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BisectionChallenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_BisectionChallenge *BisectionChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _BisectionChallenge.Contract.TimeoutChallenge(&_BisectionChallenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_BisectionChallenge *BisectionChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _BisectionChallenge.Contract.TimeoutChallenge(&_BisectionChallenge.TransactOpts)
}

// BisectionChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the BisectionChallenge contract.
type BisectionChallengeAsserterTimedOutIterator struct {
	Event *BisectionChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeAsserterTimedOut)
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
		it.Event = new(BisectionChallengeAsserterTimedOut)
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
func (it *BisectionChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the BisectionChallenge contract.
type BisectionChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*BisectionChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeAsserterTimedOutIterator{contract: _BisectionChallenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *BisectionChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeAsserterTimedOut)
				if err := _BisectionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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

// ParseAsserterTimedOut is a log parse operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*BisectionChallengeAsserterTimedOut, error) {
	event := new(BisectionChallengeAsserterTimedOut)
	if err := _BisectionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the BisectionChallenge contract.
type BisectionChallengeChallengerTimedOutIterator struct {
	Event *BisectionChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeChallengerTimedOut)
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
		it.Event = new(BisectionChallengeChallengerTimedOut)
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
func (it *BisectionChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the BisectionChallenge contract.
type BisectionChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*BisectionChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeChallengerTimedOutIterator{contract: _BisectionChallenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *BisectionChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeChallengerTimedOut)
				if err := _BisectionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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

// ParseChallengerTimedOut is a log parse operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*BisectionChallengeChallengerTimedOut, error) {
	event := new(BisectionChallengeChallengerTimedOut)
	if err := _BisectionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionChallengeContinuedIterator is returned from FilterContinued and is used to iterate over the raw logs and unpacked data for Continued events raised by the BisectionChallenge contract.
type BisectionChallengeContinuedIterator struct {
	Event *BisectionChallengeContinued // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeContinued)
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
		it.Event = new(BisectionChallengeContinued)
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
func (it *BisectionChallengeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeContinued represents a Continued event raised by the BisectionChallenge contract.
type BisectionChallengeContinued struct {
	SegmentIndex  *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContinued is a free log retrieval operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) FilterContinued(opts *bind.FilterOpts) (*BisectionChallengeContinuedIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeContinuedIterator{contract: _BisectionChallenge.contract, event: "Continued", logs: logs, sub: sub}, nil
}

// WatchContinued is a free log subscription operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) WatchContinued(opts *bind.WatchOpts, sink chan<- *BisectionChallengeContinued) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeContinued)
				if err := _BisectionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
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

// ParseContinued is a log parse operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) ParseContinued(log types.Log) (*BisectionChallengeContinued, error) {
	event := new(BisectionChallengeContinued)
	if err := _BisectionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the BisectionChallenge contract.
type BisectionChallengeInitiatedChallengeIterator struct {
	Event *BisectionChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeInitiatedChallenge)
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
		it.Event = new(BisectionChallengeInitiatedChallenge)
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
func (it *BisectionChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the BisectionChallenge contract.
type BisectionChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*BisectionChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeInitiatedChallengeIterator{contract: _BisectionChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *BisectionChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeInitiatedChallenge)
				if err := _BisectionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*BisectionChallengeInitiatedChallenge, error) {
	event := new(BisectionChallengeInitiatedChallenge)
	if err := _BisectionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFuncSigs = map[string]string{
	"ced5c1bf": "timeoutChallenge()",
}

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x608060405234801561001057600080fd5b50610222806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ced5c1bf14610030575b600080fd5b61003861003a565b005b6003546100464361011a565b11610098576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156100ab57fe5b14156100e7576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a16100e2610121565b610118565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1610118610195565b565b6103e80290565b600080546002546001546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561017957600080fd5b505af115801561018d573d6000803e3d6000fd5b503392505050ff5b600080546001546002546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561017957600080fdfea265627a7a72315820b191210c168eb9e3ac7bc532ca875f096b3f03cbd50b32070b2cfb1905eddae364736f6c63430005110032"

// DeployChallenge deploys a new Ethereum contract, binding an instance of Challenge to it.
func DeployChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Challenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_Challenge *ChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_Challenge *ChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _Challenge.Contract.TimeoutChallenge(&_Challenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_Challenge *ChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _Challenge.Contract.TimeoutChallenge(&_Challenge.TransactOpts)
}

// ChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the Challenge contract.
type ChallengeAsserterTimedOutIterator struct {
	Event *ChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *ChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeAsserterTimedOut)
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
		it.Event = new(ChallengeAsserterTimedOut)
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
func (it *ChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the Challenge contract.
type ChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*ChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &ChallengeAsserterTimedOutIterator{contract: _Challenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *ChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeAsserterTimedOut)
				if err := _Challenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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

// ParseAsserterTimedOut is a log parse operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*ChallengeAsserterTimedOut, error) {
	event := new(ChallengeAsserterTimedOut)
	if err := _Challenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the Challenge contract.
type ChallengeChallengerTimedOutIterator struct {
	Event *ChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *ChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeChallengerTimedOut)
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
		it.Event = new(ChallengeChallengerTimedOut)
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
func (it *ChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the Challenge contract.
type ChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*ChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &ChallengeChallengerTimedOutIterator{contract: _Challenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *ChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeChallengerTimedOut)
				if err := _Challenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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

// ParseChallengerTimedOut is a log parse operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*ChallengeChallengerTimedOut, error) {
	event := new(ChallengeChallengerTimedOut)
	if err := _Challenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the Challenge contract.
type ChallengeInitiatedChallengeIterator struct {
	Event *ChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeInitiatedChallenge)
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
		it.Event = new(ChallengeInitiatedChallenge)
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
func (it *ChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the Challenge contract.
type ChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_Challenge *ChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeInitiatedChallengeIterator{contract: _Challenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_Challenge *ChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeInitiatedChallenge)
				if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_Challenge *ChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ChallengeInitiatedChallenge, error) {
	event := new(ChallengeInitiatedChallenge)
	if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InboxTopChallengeABI is the input ABI used to generate the binding from.
const InboxTopChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_chainHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_chainLength\",\"type\":\"uint256\"}],\"name\":\"bisect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_lowerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InboxTopChallengeFuncSigs maps the 4-byte function signature to its string representation.
var InboxTopChallengeFuncSigs = map[string]string{
	"37423267": "bisect(bytes32[],uint256)",
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"df9ce01b": "oneStepProof(bytes32,bytes32)",
	"ced5c1bf": "timeoutChallenge()",
}

// InboxTopChallengeBin is the compiled bytecode used for deploying new contracts.
var InboxTopChallengeBin = "0x608060405234801561001057600080fd5b506110ee806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806302ad1e4e1461005c57806337423267146100a057806379a9ad8514610110578063ced5c1bf146101c2578063df9ce01b146101ca575b600080fd5b61009e600480360360a081101561007257600080fd5b506001600160a01b038135811691602081013582169160408201351690606081013590608001356101ed565b005b61009e600480360360408110156100b657600080fd5b8101906020810181356401000000008111156100d157600080fd5b8201836020820111156100e357600080fd5b8035906020019184602083028401116401000000008311171561010557600080fd5b919350915035610202565b61009e6004803603608081101561012657600080fd5b8135919081019060408101602082013564010000000081111561014857600080fd5b82018360208201111561015a57600080fd5b8035906020019184600183028401116401000000008311171561017c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356105a7565b61009e610868565b61009e600480360360408110156101e057600080fd5b5080359060200135610948565b6101f985858585610b14565b60065550505050565b60055460ff16600281111561021357fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906102c15760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561028657818101518382015260200161026e565b50505050905090810190601f1680156102b35780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506003546102ce43610c2f565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906103415760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146103bd5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b5060001982016103f96103f485856000816103d457fe5b905060200201358686858181106103e757fe5b9050602002013585610c36565b610c6d565b600182116104385760405162461bcd60e51b81526004018080602001828103825260218152602001806110996021913960400191505060405180910390fd5b606081604051908082528060200260200182016040528015610464578160200160208202803883390190505b5090506104a18585600081811061047757fe5b905060200201358686600181811061048b57fe5b9050602002013561049c8686610cdf565b610c36565b816000815181106104ae57fe5b602090810291909101015260015b82811015610519576104fa8686838181106104d357fe5b905060200201358787846001018181106104e957fe5b9050602002013561049c8787610cfd565b82828151811061050657fe5b60209081029190910101526001016104bc565b5061052381610d10565b61052b610d1f565b7f6ccb624e36453fb82f1a793715d74763283e458adc397e1f7d73dcdf604afbb185858560035460405180806020018481526020018381526020018281038252868682818152602001925060200280828437600083820152604051601f909101601f191690920182900397509095505050505050a15050505050565b60055460ff1660028111156105b857fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b815250906106295760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b5060035461063643610c2f565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b815250906106a95760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b031633146107255760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b815250906107985760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b506107a883838387600101610d3c565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b815250906108165760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b506006819055610824610e3d565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60035461087443610c2f565b116108c6576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156108d957fe5b1415610915576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a1610910610e50565b610946565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1610946610ec4565b565b60055460ff16600281111561095957fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906109ca5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b506003546109d743610c2f565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610a4a5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610ac65760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b50610adf6103f483610ad88585610f1c565b6001610c36565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1610b10610ec4565b5050565b600060055460ff166002811115610b2757fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b81525090610b9c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff19169091179055610bf4610f48565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e80290565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b6006548114604051806040016040528060088152602001672124a9afa82922ab60c11b81525090610b105760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561028657818101518382015260200161026e565b6000818381610cea57fe5b06828481610cf457fe5b04019392505050565b6000818381610d0857fe5b049392505050565b610d1981610f5a565b60065550565b600580546002919060ff19166001835b0217905550610946610f48565b600080838160205b88518111610e2f578089015193506020818a510360200181610d6257fe5b0491505b600082118015610d795750600286066001145b8015610d8757508160020a86115b15610d9a57600286046001019550610d66565b60028606610de5578383604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028681610ddd57fe5b049550610e27565b8284604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028681610e2057fe5b0460010195505b602001610d44565b505090941495945050505050565b600580546001919060ff19168280610d2f565b600080546002546001546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b158015610ea857600080fd5b505af1158015610ebc573d6000803e3d6000fd5b503392505050ff5b600080546001546002546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b158015610ea857600080fd5b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600454610f5443610c2f565b01600355565b6000815b60018151111561107b5760606002825160010181610f7857fe5b04604051908082528060200260200182016040528015610fa2578160200160208202803883390190505b50905060005b815181101561107357825181600202600101101561103b57828160020281518110610fcf57fe5b6020026020010151838260020260010181518110610fe957fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012082828151811061102a57fe5b60200260200101818152505061106b565b82816002028151811061104a57fe5b602002602001015182828151811061105e57fe5b6020026020010181815250505b600101610fa8565b509050610f5e565b8060008151811061108857fe5b602002602001015191505091905056fe43616e27742062697365637420636861696e206f66206c657373207468616e2032a265627a7a72315820ff3ba1814ed36382ed61a59c183fa43d8738ba814a43ad3f5b74137e9c52e86064736f6c63430005110032"

// DeployInboxTopChallenge deploys a new Ethereum contract, binding an instance of InboxTopChallenge to it.
func DeployInboxTopChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InboxTopChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxTopChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxTopChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InboxTopChallenge{InboxTopChallengeCaller: InboxTopChallengeCaller{contract: contract}, InboxTopChallengeTransactor: InboxTopChallengeTransactor{contract: contract}, InboxTopChallengeFilterer: InboxTopChallengeFilterer{contract: contract}}, nil
}

// InboxTopChallenge is an auto generated Go binding around an Ethereum contract.
type InboxTopChallenge struct {
	InboxTopChallengeCaller     // Read-only binding to the contract
	InboxTopChallengeTransactor // Write-only binding to the contract
	InboxTopChallengeFilterer   // Log filterer for contract events
}

// InboxTopChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxTopChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTopChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTopChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTopChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxTopChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTopChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxTopChallengeSession struct {
	Contract     *InboxTopChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InboxTopChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxTopChallengeCallerSession struct {
	Contract *InboxTopChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InboxTopChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTopChallengeTransactorSession struct {
	Contract     *InboxTopChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InboxTopChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxTopChallengeRaw struct {
	Contract *InboxTopChallenge // Generic contract binding to access the raw methods on
}

// InboxTopChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxTopChallengeCallerRaw struct {
	Contract *InboxTopChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTopChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTopChallengeTransactorRaw struct {
	Contract *InboxTopChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInboxTopChallenge creates a new instance of InboxTopChallenge, bound to a specific deployed contract.
func NewInboxTopChallenge(address common.Address, backend bind.ContractBackend) (*InboxTopChallenge, error) {
	contract, err := bindInboxTopChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InboxTopChallenge{InboxTopChallengeCaller: InboxTopChallengeCaller{contract: contract}, InboxTopChallengeTransactor: InboxTopChallengeTransactor{contract: contract}, InboxTopChallengeFilterer: InboxTopChallengeFilterer{contract: contract}}, nil
}

// NewInboxTopChallengeCaller creates a new read-only instance of InboxTopChallenge, bound to a specific deployed contract.
func NewInboxTopChallengeCaller(address common.Address, caller bind.ContractCaller) (*InboxTopChallengeCaller, error) {
	contract, err := bindInboxTopChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeCaller{contract: contract}, nil
}

// NewInboxTopChallengeTransactor creates a new write-only instance of InboxTopChallenge, bound to a specific deployed contract.
func NewInboxTopChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTopChallengeTransactor, error) {
	contract, err := bindInboxTopChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeTransactor{contract: contract}, nil
}

// NewInboxTopChallengeFilterer creates a new log filterer instance of InboxTopChallenge, bound to a specific deployed contract.
func NewInboxTopChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxTopChallengeFilterer, error) {
	contract, err := bindInboxTopChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeFilterer{contract: contract}, nil
}

// bindInboxTopChallenge binds a generic wrapper to an already deployed contract.
func bindInboxTopChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxTopChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxTopChallenge *InboxTopChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InboxTopChallenge.Contract.InboxTopChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxTopChallenge *InboxTopChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.InboxTopChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxTopChallenge *InboxTopChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.InboxTopChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxTopChallenge *InboxTopChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InboxTopChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxTopChallenge *InboxTopChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxTopChallenge *InboxTopChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.contract.Transact(opts, method, params...)
}

// Bisect is a paid mutator transaction binding the contract method 0x37423267.
//
// Solidity: function bisect(bytes32[] _chainHashes, uint256 _chainLength) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactor) Bisect(opts *bind.TransactOpts, _chainHashes [][32]byte, _chainLength *big.Int) (*types.Transaction, error) {
	return _InboxTopChallenge.contract.Transact(opts, "bisect", _chainHashes, _chainLength)
}

// Bisect is a paid mutator transaction binding the contract method 0x37423267.
//
// Solidity: function bisect(bytes32[] _chainHashes, uint256 _chainLength) returns()
func (_InboxTopChallenge *InboxTopChallengeSession) Bisect(_chainHashes [][32]byte, _chainLength *big.Int) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.Bisect(&_InboxTopChallenge.TransactOpts, _chainHashes, _chainLength)
}

// Bisect is a paid mutator transaction binding the contract method 0x37423267.
//
// Solidity: function bisect(bytes32[] _chainHashes, uint256 _chainLength) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactorSession) Bisect(_chainHashes [][32]byte, _chainLength *big.Int) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.Bisect(&_InboxTopChallenge.TransactOpts, _chainHashes, _chainLength)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactor) ChooseSegment(opts *bind.TransactOpts, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.contract.Transact(opts, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_InboxTopChallenge *InboxTopChallengeSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.ChooseSegment(&_InboxTopChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactorSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.ChooseSegment(&_InboxTopChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.contract.Transact(opts, "initializeBisection", _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_InboxTopChallenge *InboxTopChallengeSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.InitializeBisection(&_InboxTopChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactorSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.InitializeBisection(&_InboxTopChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xdf9ce01b.
//
// Solidity: function oneStepProof(bytes32 _lowerHash, bytes32 _value) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactor) OneStepProof(opts *bind.TransactOpts, _lowerHash [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.contract.Transact(opts, "oneStepProof", _lowerHash, _value)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xdf9ce01b.
//
// Solidity: function oneStepProof(bytes32 _lowerHash, bytes32 _value) returns()
func (_InboxTopChallenge *InboxTopChallengeSession) OneStepProof(_lowerHash [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.OneStepProof(&_InboxTopChallenge.TransactOpts, _lowerHash, _value)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xdf9ce01b.
//
// Solidity: function oneStepProof(bytes32 _lowerHash, bytes32 _value) returns()
func (_InboxTopChallenge *InboxTopChallengeTransactorSession) OneStepProof(_lowerHash [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.OneStepProof(&_InboxTopChallenge.TransactOpts, _lowerHash, _value)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_InboxTopChallenge *InboxTopChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxTopChallenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_InboxTopChallenge *InboxTopChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.TimeoutChallenge(&_InboxTopChallenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_InboxTopChallenge *InboxTopChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _InboxTopChallenge.Contract.TimeoutChallenge(&_InboxTopChallenge.TransactOpts)
}

// InboxTopChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the InboxTopChallenge contract.
type InboxTopChallengeAsserterTimedOutIterator struct {
	Event *InboxTopChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *InboxTopChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxTopChallengeAsserterTimedOut)
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
		it.Event = new(InboxTopChallengeAsserterTimedOut)
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
func (it *InboxTopChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxTopChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxTopChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the InboxTopChallenge contract.
type InboxTopChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_InboxTopChallenge *InboxTopChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*InboxTopChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _InboxTopChallenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeAsserterTimedOutIterator{contract: _InboxTopChallenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_InboxTopChallenge *InboxTopChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *InboxTopChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _InboxTopChallenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxTopChallengeAsserterTimedOut)
				if err := _InboxTopChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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

// ParseAsserterTimedOut is a log parse operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_InboxTopChallenge *InboxTopChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*InboxTopChallengeAsserterTimedOut, error) {
	event := new(InboxTopChallengeAsserterTimedOut)
	if err := _InboxTopChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InboxTopChallengeBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the InboxTopChallenge contract.
type InboxTopChallengeBisectedIterator struct {
	Event *InboxTopChallengeBisected // Event containing the contract specifics and raw log

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
func (it *InboxTopChallengeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxTopChallengeBisected)
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
		it.Event = new(InboxTopChallengeBisected)
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
func (it *InboxTopChallengeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxTopChallengeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxTopChallengeBisected represents a Bisected event raised by the InboxTopChallenge contract.
type InboxTopChallengeBisected struct {
	ChainHashes   [][32]byte
	TotalLength   *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x6ccb624e36453fb82f1a793715d74763283e458adc397e1f7d73dcdf604afbb1.
//
// Solidity: event Bisected(bytes32[] chainHashes, uint256 totalLength, uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) FilterBisected(opts *bind.FilterOpts) (*InboxTopChallengeBisectedIterator, error) {

	logs, sub, err := _InboxTopChallenge.contract.FilterLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeBisectedIterator{contract: _InboxTopChallenge.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x6ccb624e36453fb82f1a793715d74763283e458adc397e1f7d73dcdf604afbb1.
//
// Solidity: event Bisected(bytes32[] chainHashes, uint256 totalLength, uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *InboxTopChallengeBisected) (event.Subscription, error) {

	logs, sub, err := _InboxTopChallenge.contract.WatchLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxTopChallengeBisected)
				if err := _InboxTopChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x6ccb624e36453fb82f1a793715d74763283e458adc397e1f7d73dcdf604afbb1.
//
// Solidity: event Bisected(bytes32[] chainHashes, uint256 totalLength, uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) ParseBisected(log types.Log) (*InboxTopChallengeBisected, error) {
	event := new(InboxTopChallengeBisected)
	if err := _InboxTopChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InboxTopChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the InboxTopChallenge contract.
type InboxTopChallengeChallengerTimedOutIterator struct {
	Event *InboxTopChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *InboxTopChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxTopChallengeChallengerTimedOut)
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
		it.Event = new(InboxTopChallengeChallengerTimedOut)
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
func (it *InboxTopChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxTopChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxTopChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the InboxTopChallenge contract.
type InboxTopChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_InboxTopChallenge *InboxTopChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*InboxTopChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _InboxTopChallenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeChallengerTimedOutIterator{contract: _InboxTopChallenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_InboxTopChallenge *InboxTopChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *InboxTopChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _InboxTopChallenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxTopChallengeChallengerTimedOut)
				if err := _InboxTopChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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

// ParseChallengerTimedOut is a log parse operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_InboxTopChallenge *InboxTopChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*InboxTopChallengeChallengerTimedOut, error) {
	event := new(InboxTopChallengeChallengerTimedOut)
	if err := _InboxTopChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InboxTopChallengeContinuedIterator is returned from FilterContinued and is used to iterate over the raw logs and unpacked data for Continued events raised by the InboxTopChallenge contract.
type InboxTopChallengeContinuedIterator struct {
	Event *InboxTopChallengeContinued // Event containing the contract specifics and raw log

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
func (it *InboxTopChallengeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxTopChallengeContinued)
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
		it.Event = new(InboxTopChallengeContinued)
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
func (it *InboxTopChallengeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxTopChallengeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxTopChallengeContinued represents a Continued event raised by the InboxTopChallenge contract.
type InboxTopChallengeContinued struct {
	SegmentIndex  *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContinued is a free log retrieval operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) FilterContinued(opts *bind.FilterOpts) (*InboxTopChallengeContinuedIterator, error) {

	logs, sub, err := _InboxTopChallenge.contract.FilterLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeContinuedIterator{contract: _InboxTopChallenge.contract, event: "Continued", logs: logs, sub: sub}, nil
}

// WatchContinued is a free log subscription operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) WatchContinued(opts *bind.WatchOpts, sink chan<- *InboxTopChallengeContinued) (event.Subscription, error) {

	logs, sub, err := _InboxTopChallenge.contract.WatchLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxTopChallengeContinued)
				if err := _InboxTopChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
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

// ParseContinued is a log parse operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) ParseContinued(log types.Log) (*InboxTopChallengeContinued, error) {
	event := new(InboxTopChallengeContinued)
	if err := _InboxTopChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InboxTopChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the InboxTopChallenge contract.
type InboxTopChallengeInitiatedChallengeIterator struct {
	Event *InboxTopChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *InboxTopChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxTopChallengeInitiatedChallenge)
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
		it.Event = new(InboxTopChallengeInitiatedChallenge)
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
func (it *InboxTopChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxTopChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxTopChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the InboxTopChallenge contract.
type InboxTopChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*InboxTopChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _InboxTopChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeInitiatedChallengeIterator{contract: _InboxTopChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *InboxTopChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _InboxTopChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxTopChallengeInitiatedChallenge)
				if err := _InboxTopChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_InboxTopChallenge *InboxTopChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*InboxTopChallengeInitiatedChallenge, error) {
	event := new(InboxTopChallengeInitiatedChallenge)
	if err := _InboxTopChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InboxTopChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the InboxTopChallenge contract.
type InboxTopChallengeOneStepProofCompletedIterator struct {
	Event *InboxTopChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *InboxTopChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxTopChallengeOneStepProofCompleted)
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
		it.Event = new(InboxTopChallengeOneStepProofCompleted)
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
func (it *InboxTopChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxTopChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxTopChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the InboxTopChallenge contract.
type InboxTopChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_InboxTopChallenge *InboxTopChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*InboxTopChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _InboxTopChallenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &InboxTopChallengeOneStepProofCompletedIterator{contract: _InboxTopChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_InboxTopChallenge *InboxTopChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *InboxTopChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _InboxTopChallenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxTopChallengeOneStepProofCompleted)
				if err := _InboxTopChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_InboxTopChallenge *InboxTopChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*InboxTopChallengeOneStepProofCompleted, error) {
	event := new(InboxTopChallengeOneStepProofCompleted)
	if err := _InboxTopChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}
