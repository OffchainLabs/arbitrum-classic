// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executionchallenge

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
var BisectionChallengeBin = "0x608060405234801561001057600080fd5b50610882806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806302ad1e4e1461004657806379a9ad851461008a578063ced5c1bf1461013c575b600080fd5b610088600480360360a081101561005c57600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060800135610144565b005b610088600480360360808110156100a057600080fd5b813591908101906040810160208201356401000000008111156100c257600080fd5b8201836020820111156100d457600080fd5b803590602001918460018302840111640100000000831117156100f657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610159565b610088610457565b61015085858585610537565b60065550505050565b60055460ff16600281111561016a57fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b815250906102185760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156101dd5781810151838201526020016101c5565b50505050905090810190601f16801561020a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060035461022543610652565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b815250906102985760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b031633146103145760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b815250906103875760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b5061039783838387600101610659565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b815250906104055760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b50600681905561041361075a565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60035461046343610652565b116104b5576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156104c857fe5b1415610504576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a16104ff61076f565b610535565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a16105356107e3565b565b600060055460ff16600281111561054a57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906105bf5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156101dd5781810151838201526020016101c5565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff1916909117905561061761083b565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e80290565b600080838160205b8851811161074c578089015193506020818a51036020018161067f57fe5b0491505b6000821180156106965750600286066001145b80156106a457508160020a86115b156106b757600286046001019550610683565b600286066107025783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816106fa57fe5b049550610744565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161073d57fe5b0460010195505b602001610661565b505090941495945050505050565b6005805460ff1916600117905561053561083b565b600080546002546001546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b1580156107c757600080fd5b505af11580156107db573d6000803e3d6000fd5b503392505050ff5b600080546001546002546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b1580156107c757600080fd5b60045461084743610652565b0160035556fea265627a7a723158209850b11081f5714a137b7872ff64bb22d8dbc51d320d07d81dcc0b80df958e6c64736f6c63430005110032"

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

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209cc6ea8cfb5d0f6e66ccce67c7494628093cb8581f492f996c0110163c56d51b64736f6c63430005110032"

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

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFuncSigs = map[string]string{
	"ced5c1bf": "timeoutChallenge()",
}

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x608060405234801561001057600080fd5b50610222806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ced5c1bf14610030575b600080fd5b61003861003a565b005b6003546100464361011a565b11610098576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156100ab57fe5b14156100e7576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a16100e2610121565b610118565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1610118610195565b565b6103e80290565b600080546002546001546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561017957600080fd5b505af115801561018d573d6000803e3d6000fd5b503392505050ff5b600080546001546002546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561017957600080fdfea265627a7a7231582026f12cdce58489d144e89925591dcaaf16947f1ea50bd7c07a731feaf290510664736f6c63430005110032"

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

// ChallengeUtilsABI is the input ABI used to generate the binding from.
const ChallengeUtilsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_EXECUTION_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_INBOX_TOP_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_MESSAGES_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALID_CHILD_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeUtilsFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeUtilsFuncSigs = map[string]string{
	"95312727": "INVALID_EXECUTION_TYPE()",
	"a697bcac": "INVALID_INBOX_TOP_TYPE()",
	"d7519b46": "INVALID_MESSAGES_TYPE()",
	"2e179be5": "VALID_CHILD_TYPE()",
}

// ChallengeUtilsBin is the compiled bytecode used for deploying new contracts.
var ChallengeUtilsBin = "0x60c9610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060515760003560e01c80632e179be51460565780639531272714606e578063a697bcac146074578063d7519b4614607a575b600080fd5b605c6080565b60408051918252519081900360200190f35b605c6085565b605c608a565b605c608f565b600381565b600281565b600081565b60018156fea265627a7a72315820d6b82a8f1da234af8e8a5ebc59988c3739058feafa08621f40b955889c0b5e3364736f6c63430005110032"

// DeployChallengeUtils deploys a new Ethereum contract, binding an instance of ChallengeUtils to it.
func DeployChallengeUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeUtils{ChallengeUtilsCaller: ChallengeUtilsCaller{contract: contract}, ChallengeUtilsTransactor: ChallengeUtilsTransactor{contract: contract}, ChallengeUtilsFilterer: ChallengeUtilsFilterer{contract: contract}}, nil
}

// ChallengeUtils is an auto generated Go binding around an Ethereum contract.
type ChallengeUtils struct {
	ChallengeUtilsCaller     // Read-only binding to the contract
	ChallengeUtilsTransactor // Write-only binding to the contract
	ChallengeUtilsFilterer   // Log filterer for contract events
}

// ChallengeUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeUtilsSession struct {
	Contract     *ChallengeUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeUtilsCallerSession struct {
	Contract *ChallengeUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChallengeUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeUtilsTransactorSession struct {
	Contract     *ChallengeUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChallengeUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeUtilsRaw struct {
	Contract *ChallengeUtils // Generic contract binding to access the raw methods on
}

// ChallengeUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeUtilsCallerRaw struct {
	Contract *ChallengeUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeUtilsTransactorRaw struct {
	Contract *ChallengeUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeUtils creates a new instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtils(address common.Address, backend bind.ContractBackend) (*ChallengeUtils, error) {
	contract, err := bindChallengeUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtils{ChallengeUtilsCaller: ChallengeUtilsCaller{contract: contract}, ChallengeUtilsTransactor: ChallengeUtilsTransactor{contract: contract}, ChallengeUtilsFilterer: ChallengeUtilsFilterer{contract: contract}}, nil
}

// NewChallengeUtilsCaller creates a new read-only instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsCaller(address common.Address, caller bind.ContractCaller) (*ChallengeUtilsCaller, error) {
	contract, err := bindChallengeUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsCaller{contract: contract}, nil
}

// NewChallengeUtilsTransactor creates a new write-only instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeUtilsTransactor, error) {
	contract, err := bindChallengeUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsTransactor{contract: contract}, nil
}

// NewChallengeUtilsFilterer creates a new log filterer instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeUtilsFilterer, error) {
	contract, err := bindChallengeUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsFilterer{contract: contract}, nil
}

// bindChallengeUtils binds a generic wrapper to an already deployed contract.
func bindChallengeUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeUtils *ChallengeUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeUtils.Contract.ChallengeUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeUtils *ChallengeUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.ChallengeUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeUtils *ChallengeUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.ChallengeUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeUtils *ChallengeUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeUtils *ChallengeUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeUtils *ChallengeUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.contract.Transact(opts, method, params...)
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDEXECUTIONTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_EXECUTION_TYPE")
	return *ret0, err
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDEXECUTIONTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDEXECUTIONTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDEXECUTIONTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDEXECUTIONTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDINBOXTOPTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_INBOX_TOP_TYPE")
	return *ret0, err
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDINBOXTOPTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDINBOXTOPTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDINBOXTOPTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDINBOXTOPTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDMESSAGESTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_MESSAGES_TYPE")
	return *ret0, err
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDMESSAGESTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDMESSAGESTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDMESSAGESTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDMESSAGESTYPE(&_ChallengeUtils.CallOpts)
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) VALIDCHILDTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "VALID_CHILD_TYPE")
	return *ret0, err
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) VALIDCHILDTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.VALIDCHILDTYPE(&_ChallengeUtils.CallOpts)
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) VALIDCHILDTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.VALIDCHILDTYPE(&_ChallengeUtils.CallOpts)
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[]"

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820f1f511f3dfbb775ccca8f03416243a8f73f5f620d61b74b943ba2554e9fdc65a64736f6c63430005110032"

// DeployDebugPrint deploys a new Ethereum contract, binding an instance of DebugPrint to it.
func DeployDebugPrint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebugPrint, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DebugPrintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// DebugPrint is an auto generated Go binding around an Ethereum contract.
type DebugPrint struct {
	DebugPrintCaller     // Read-only binding to the contract
	DebugPrintTransactor // Write-only binding to the contract
	DebugPrintFilterer   // Log filterer for contract events
}

// DebugPrintCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebugPrintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebugPrintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebugPrintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebugPrintSession struct {
	Contract     *DebugPrint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebugPrintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebugPrintCallerSession struct {
	Contract *DebugPrintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebugPrintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebugPrintTransactorSession struct {
	Contract     *DebugPrintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebugPrintRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebugPrintRaw struct {
	Contract *DebugPrint // Generic contract binding to access the raw methods on
}

// DebugPrintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebugPrintCallerRaw struct {
	Contract *DebugPrintCaller // Generic read-only contract binding to access the raw methods on
}

// DebugPrintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebugPrintTransactorRaw struct {
	Contract *DebugPrintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebugPrint creates a new instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrint(address common.Address, backend bind.ContractBackend) (*DebugPrint, error) {
	contract, err := bindDebugPrint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// NewDebugPrintCaller creates a new read-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintCaller(address common.Address, caller bind.ContractCaller) (*DebugPrintCaller, error) {
	contract, err := bindDebugPrint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintCaller{contract: contract}, nil
}

// NewDebugPrintTransactor creates a new write-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugPrintTransactor, error) {
	contract, err := bindDebugPrint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintTransactor{contract: contract}, nil
}

// NewDebugPrintFilterer creates a new log filterer instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugPrintFilterer, error) {
	contract, err := bindDebugPrint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugPrintFilterer{contract: contract}, nil
}

// bindDebugPrint binds a generic wrapper to an already deployed contract.
func bindDebugPrint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.DebugPrintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transact(opts, method, params...)
}

// ExecutionChallengeABI is the input ABI used to generate the binding from.
const ExecutionChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"machineHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"didInboxInsns\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"logAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"gases\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_machineHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"_didInboxInsns\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_logAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_gases\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_totalSteps\",\"type\":\"uint64\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxValueSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsns\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExecutionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionChallengeFuncSigs = map[string]string{
	"eb83afbf": "bisectAssertion(bytes32,bytes32[],bool[],bytes32[],bytes32[],uint64[],uint64)",
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"82498b81": "oneStepProof(bytes32,bytes32,uint256,bytes32,bool,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
	"ced5c1bf": "timeoutChallenge()",
}

// ExecutionChallengeBin is the compiled bytecode used for deploying new contracts.
var ExecutionChallengeBin = "0x608060405234801561001057600080fd5b50615543806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806302ad1e4e1461005c57806379a9ad85146100a057806382498b8114610150578063ced5c1bf14610241578063eb83afbf14610249575b600080fd5b61009e600480360360a081101561007257600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060800135610408565b005b61009e600480360360808110156100b657600080fd5b81359190810190604081016020820135600160201b8111156100d757600080fd5b8201836020820111156100e957600080fd5b803590602001918460018302840111600160201b8311171561010a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561041d565b61009e600480360361016081101561016757600080fd5b813591602081013591604082013591606081013591608082013515159160a08101359160c08201359160e0810135916101008201359167ffffffffffffffff61012082013516918101906101608101610140820135600160201b8111156101cd57600080fd5b8201836020820111156101df57600080fd5b803590602001918460018302840111600160201b8311171561020057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061071b945050505050565b61009e610963565b61009e600480360360e081101561025f57600080fd5b81359190810190604081016020820135600160201b81111561028057600080fd5b82018360208201111561029257600080fd5b803590602001918460208302840111600160201b831117156102b357600080fd5b919390929091602081019035600160201b8111156102d057600080fd5b8201836020820111156102e257600080fd5b803590602001918460208302840111600160201b8311171561030357600080fd5b919390929091602081019035600160201b81111561032057600080fd5b82018360208201111561033257600080fd5b803590602001918460208302840111600160201b8311171561035357600080fd5b919390929091602081019035600160201b81111561037057600080fd5b82018360208201111561038257600080fd5b803590602001918460208302840111600160201b831117156103a357600080fd5b919390929091602081019035600160201b8111156103c057600080fd5b8201836020820111156103d257600080fd5b803590602001918460208302840111600160201b831117156103f357600080fd5b91935091503567ffffffffffffffff16610a43565b61041485858585610d02565b60065550505050565b60055460ff16600281111561042e57fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b815250906104dc5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156104a1578181015183820152602001610489565b50505050905090810190601f1680156104ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506003546104e943610e1d565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b8152509061055c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b031633146105d85760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b8152509061064b5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5061065b83838387600101610e28565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b815250906106c95760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5060068190556106d7610f2b565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60055460ff16600281111561072c57fe5b600114604051806040016040528060098152602001684249535f535441544560b81b8152509061079d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b506003546107aa43610e1d565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b8152509061081d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146108995760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b506108ac8b8b8b8b8b878c8c8c8c610f47565b6108b4615347565b6108c68c8c8c8b8b8b8b8b8b8b610f85565b9050886108d282610ff9565b14610924576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a16109556110bd565b505050505050505050505050565b60035461096f43610e1d565b116109c1576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156109d457fe5b1415610a10576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a1610a0b611131565b610a41565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1610a416110bd565b565b60055460ff166002811115610a5457fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610ac55760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b50600354610ad243610e1d565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610b455760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610bc15760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b506109556040518060e001604052808e81526020018d8d808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505050908252506040805160208d810282810182019093528d82529283019290918e918e9182918501908490808284376000920191909152505050908252506040805160208b810282810182019093528b82529283019290918c918c91829185019084908082843760009201919091525050509082525060408051602089810282810182019093528982529283019290918a918a918291850190849080828437600092019190915250505090825250604080516020878102828101820190935287825292830192909188918891829185019084908082843760009201919091525050509082525067ffffffffffffffff8416602090910152611189565b600060055460ff166002811115610d1557fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b81525090610d8a5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff19169091179055610de26118ca565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e881025b919050565b600080838160205b88518111610f1b578089015193506020818a510360200181610e4e57fe5b0491505b600082118015610e655750600286066001145b8015610e7357508160020a86115b15610e8657600286046001019550610e52565b60028606610ed1578383604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028681610ec957fe5b049550610f13565b8284604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028681610f0c57fe5b0460010195505b602001610e30565b505085149150505b949350505050565b600580546001919060ff191682805b0217905550610a416118ca565b6000610f5b610f568b8b6118dc565b61198f565b9050610f78610f7360018d848c8c8c8c8c8c8c611aaf565b611b32565b5050505050505050505050565b610f8d615347565b610fea6040518061012001604052808d8152602001610fac8d8d6118dc565b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250611ba8565b9b9a5050505050505050505050565b600060028260e00151141561101057506000610e23565b60018260e00151141561102557506001610e23565b815160208301516110359061198f565b611042846040015161198f565b61104f856060015161198f565b61105c866080015161198f565b8660a001518760c0015160405160200180888152602001878152602001868152602001858152602001848152602001838152602001828152602001975050505050505050604051602081830303815290604052805190602001209050610e23565b600080546001546002546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561111557600080fd5b505af1158015611129573d6000803e3d6000fd5b503392505050ff5b600080546002546001546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561111557600080fd5b6000600182602001515103905081604001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b8152509061120d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5081606001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906112885760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5081608001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906113035760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b508160a001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b8152509061137b5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b50600080805b838110156113cf578460a00151818151811061139957fe5b60200260200101518301925081806113c55750846040015181815181106113bc57fe5b60200260200101515b9150600101611381565b5061147a610f738560c0015186602001516000815181106113ec57fe5b602002602001015187600001518860200151888151811061140957fe5b602002602001015186888b6060015160008151811061142457fe5b60200260200101518c606001518c8151811061143c57fe5b60200260200101518d6080015160008151811061145557fe5b60200260200101518e608001518e8151811061146d57fe5b6020026020010151611aaf565b6060836040519080825280602002602001820160405280156114a6578160200160208202803883390190505b50905061158f6114c48660c0015167ffffffffffffffff1686612cf8565b63ffffffff1686602001516000815181106114db57fe5b6020026020010151876000015188602001516001815181106114f957fe5b6020026020010151896040015160008151811061151257fe5b60200260200101518a60a0015160008151811061152b57fe5b60200260200101518b6060015160008151811061154457fe5b60200260200101518c6060015160018151811061155d57fe5b60200260200101518d6080015160008151811061157657fe5b60200260200101518e6080015160018151811061146d57fe5b8160008151811061159c57fe5b602090810291909101015260015b848110156116e657856040015160018203815181106115c557fe5b6020026020010151156115e0576115dd610f56612d16565b86525b6116c76115fb8760c0015167ffffffffffffffff1687612d5d565b63ffffffff168760200151838151811061161157fe5b602002602001015188600001518960200151856001018151811061163157fe5b60200260200101518a60400151868151811061164957fe5b60200260200101518b60a00151878151811061166157fe5b60200260200101518c60600151888151811061167957fe5b60200260200101518d60600151896001018151811061169457fe5b60200260200101518e608001518a815181106116ac57fe5b60200260200101518f608001518b6001018151811061146d57fe5b8282815181106116d357fe5b60209081029190910101526001016115aa565b506116f081612d70565b6116f8612d7f565b7f99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee4785602001518660400151876060015188608001518960a001518a60c001516003546040518080602001806020018060200180602001806020018867ffffffffffffffff1667ffffffffffffffff16815260200187815260200186810386528d818151815260200191508051906020019060200280838360005b838110156117aa578181015183820152602001611792565b5050505090500186810385528c818151815260200191508051906020019060200280838360005b838110156117e95781810151838201526020016117d1565b5050505090500186810384528b818151815260200191508051906020019060200280838360005b83811015611828578181015183820152602001611810565b5050505090500186810383528a818151815260200191508051906020019060200280838360005b8381101561186757818101518382015260200161184f565b50505050905001868103825289818151815260200191508051906020019060200280838360005b838110156118a657818101518382015260200161188e565b505050509050019c5050505050505050505050505060405180910390a15050505050565b6004546118d643610e1d565b01600355565b6118e46153a5565b6040805160a081018252848152815160608101835260008082526020828101829052845182815280820186529394908501939083019161193a565b6119276153a5565b81526020019060019003908161191f5790505b50905281526040805160008082526020828101909352919092019190611976565b6119636153a5565b81526020019060019003908161195b5790505b5081526002602082015260400183905290505b92915050565b6000611999612d93565b60ff16826060015160ff1614156119bc5781516119b590612d98565b9050610e23565b6119c4612dbc565b60ff16826060015160ff1614156119e2576119b58260200151612dc1565b6119ea612ebe565b60ff16826060015160ff161415611a0c57815160808301516119b59190612ec3565b611a14612f14565b60ff16826060015160ff161415611a4d57611a2d6153a5565b611a3a8360400151612f19565b9050611a458161198f565b915050610e23565b611a5561307b565b60ff16826060015160ff161415611a6e57508051610e23565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604080516001600160c01b031960c09c8d1b8116602080840191909152602883019c909c52604882019a909a52606881019890985295151560f81b60888801529390981b9095166089850152609184015260b183019390935260d182019490945260f1808201929092528351808203909201825261011101909252815191012090565b6006548114604051806040016040528060088152602001672124a9afa82922ab60c11b81525090611ba45760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104a1578181015183820152602001610489565b5050565b611bb0615347565b60008060006060611bbf615347565b611bc7615347565b611bd088613080565b60e08e0151959b509399509297509095509350915060019060009067ffffffffffffffff168714611c3f576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b89604001518015611c53575060ff88166072145b80611c6f57508960400151158015611c6f575060ff8816607214155b611cc0576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a080840180518990039052840151871115611ce75760001960a084015260009150612bb1565b60ff881660011415611d2d57611d268386600081518110611d0457fe5b602002602001015187600181518110611d1957fe5b60200260200101516132d6565b9150612bb1565b60ff881660021415611d6c57611d268386600081518110611d4a57fe5b602002602001015187600181518110611d5f57fe5b6020026020010151613326565b60ff881660031415611dab57611d268386600081518110611d8957fe5b602002602001015187600181518110611d9e57fe5b6020026020010151613367565b60ff881660041415611dea57611d268386600081518110611dc857fe5b602002602001015187600181518110611ddd57fe5b60200260200101516133a8565b60ff881660051415611e2957611d268386600081518110611e0757fe5b602002602001015187600181518110611e1c57fe5b60200260200101516133f9565b60ff881660061415611e6857611d268386600081518110611e4657fe5b602002602001015187600181518110611e5b57fe5b602002602001015161344a565b60ff881660071415611ea757611d268386600081518110611e8557fe5b602002602001015187600181518110611e9a57fe5b602002602001015161349b565b60ff881660081415611efb57611d268386600081518110611ec457fe5b602002602001015187600181518110611ed957fe5b602002602001015188600281518110611eee57fe5b60200260200101516134ec565b60ff881660091415611f4f57611d268386600081518110611f1857fe5b602002602001015187600181518110611f2d57fe5b602002602001015188600281518110611f4257fe5b6020026020010151613554565b60ff8816600a1415611f8e57611d268386600081518110611f6c57fe5b602002602001015187600181518110611f8157fe5b60200260200101516135ad565b60ff881660101415611fcd57611d268386600081518110611fab57fe5b602002602001015187600181518110611fc057fe5b60200260200101516135ee565b60ff88166011141561200c57611d268386600081518110611fea57fe5b602002602001015187600181518110611fff57fe5b602002602001015161362f565b60ff88166012141561204b57611d26838660008151811061202957fe5b60200260200101518760018151811061203e57fe5b6020026020010151613670565b60ff88166013141561208a57611d26838660008151811061206857fe5b60200260200101518760018151811061207d57fe5b60200260200101516136b1565b60ff8816601414156120c957611d2683866000815181106120a757fe5b6020026020010151876001815181106120bc57fe5b60200260200101516136f2565b60ff8816601514156120f357611d2683866000815181106120e657fe5b6020026020010151613729565b60ff88166016141561213257611d26838660008151811061211057fe5b60200260200101518760018151811061212557fe5b602002602001015161376e565b60ff88166017141561217157611d26838660008151811061214f57fe5b60200260200101518760018151811061216457fe5b60200260200101516137af565b60ff8816601814156121b057611d26838660008151811061218e57fe5b6020026020010151876001815181106121a357fe5b60200260200101516137f0565b60ff8816601914156121da57611d2683866000815181106121cd57fe5b6020026020010151613831565b60ff8816601a141561221957611d2683866000815181106121f757fe5b60200260200101518760018151811061220c57fe5b6020026020010151613867565b60ff8816601b141561225857611d26838660008151811061223657fe5b60200260200101518760018151811061224b57fe5b60200260200101516138a8565b60ff88166020141561228257611d26838660008151811061227557fe5b60200260200101516138e9565b60ff8816602114156122ac57611d26838660008151811061229f57fe5b6020026020010151613904565b60ff8816602214156122eb57611d2683866000815181106122c957fe5b6020026020010151876001815181106122de57fe5b602002602001015161391f565b60ff88166030141561231557611d26838660008151811061230857fe5b6020026020010151613985565b60ff88166031141561232a57611d268361398d565b60ff88166032141561233f57611d26836139ae565b60ff88166033141561236957611d26838660008151811061235c57fe5b60200260200101516139c7565b60ff88166034141561239357611d26838660008151811061238657fe5b60200260200101516139d3565b60ff8816603514156123d257611d2683866000815181106123b057fe5b6020026020010151876001815181106123c557fe5b60200260200101516139fe565b60ff8816603614156123e757611d2683613a46565b60ff88166037141561240157611d26838560000151613a73565b60ff88166038141561242b57611d26838660008151811061241e57fe5b6020026020010151613a83565b60ff8816603914156124795761243f6153a5565b61244e8b610100015188613a95565b9097509050612463858263ffffffff613c2716565b612473848263ffffffff613c4116565b50612bb1565b60ff8816603a141561248e57611d2683613c5b565b60ff8816603b14156124a35760019150612bb1565b60ff8816603c14156124b857611d2683613c7b565b60ff8816603d14156124e257611d2683866000815181106124d557fe5b6020026020010151613c8f565b60ff88166040141561250c57611d2683866000815181106124ff57fe5b6020026020010151613cbd565b60ff88166041141561254b57611d26838660008151811061252957fe5b60200260200101518760018151811061253e57fe5b6020026020010151613cdf565b60ff88166042141561259f57611d26838660008151811061256857fe5b60200260200101518760018151811061257d57fe5b60200260200101518860028151811061259257fe5b6020026020010151613d11565b60ff8816604314156125de57611d2683866000815181106125bc57fe5b6020026020010151876001815181106125d157fe5b6020026020010151613d53565b60ff88166044141561263257611d2683866000815181106125fb57fe5b60200260200101518760018151811061261057fe5b60200260200101518860028151811061262557fe5b6020026020010151613d65565b60ff88166050141561267157611d26838660008151811061264f57fe5b60200260200101518760018151811061266457fe5b6020026020010151613d87565b60ff8816605114156126c557611d26838660008151811061268e57fe5b6020026020010151876001815181106126a357fe5b6020026020010151886002815181106126b857fe5b6020026020010151613dfd565b60ff8816605214156126ef57611d2683866000815181106126e257fe5b6020026020010151613e8a565b60ff88166053141561274e576127036153a5565b6127128b610100015188613a95565b9097509050612727858263ffffffff613c2716565b612746848760008151811061273857fe5b602002602001015183613ebd565b925050612bb1565b60ff8816605414156127ba576127626153a5565b6127718b610100015188613a95565b9097509050612786858263ffffffff613c2716565b612746848760008151811061279757fe5b6020026020010151886001815181106127ac57fe5b602002602001015184613f15565b60ff8816606014156127cf57611d2683613f96565b60ff8816606114156128cc576127f983866000815181106127ec57fe5b6020026020010151613f9c565b909250905081156128c3578960c001518a60a0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146128785760405162461bcd60e51b81526004018080602001828103825260258152602001806154c36025913960400191505060405180910390fd5b89608001518a60600151146128be5760405162461bcd60e51b81526004018080602001828103825260278152602001806154e86027913960400191505060405180910390fd5b6128c7565b5060005b612bb1565b60ff881660701415612a03576128f683866000815181106128e957fe5b6020026020010151613fb5565b909250905081156128c357806129515789608001518a606001511461294c5760405162461bcd60e51b815260040180806020018281038252603881526020018061548b6038913960400191505060405180910390fd5b6128be565b60808a01516060808c015160408051602080820193909352808201869052815180820383018152930190528151910120146129bd5760405162461bcd60e51b815260040180806020018281038252602981526020018061541b6029913960400191505060405180910390fd5b8960c001518a60a00151146128be5760405162461bcd60e51b81526004018080602001828103825260268152602001806154446026913960400191505060405180910390fd5b60ff881660721415612a1d57611d26838b60200151613ff4565b60ff881660731415612a325760009150612bb1565b60ff881660741415612a47576128c78361405d565b60ff881660751415612a7157611d268386600081518110612a6457fe5b6020026020010151614067565b60ff881660761415612a8657611d268361408c565b60ff881660771415612a9b57611d26836140a5565b60ff881660781415612ada57611d268386600081518110612ab857fe5b602002602001015187600181518110612acd57fe5b60200260200101516140ee565b60ff881660791415612b2e57611d268386600081518110612af757fe5b602002602001015187600181518110612b0c57fe5b602002602001015188600281518110612b2157fe5b6020026020010151614133565b60ff8816607b1415612b4357611d2683614186565b60ff881660801415612bac57611d268386600081518110612b6057fe5b602002602001015187600181518110612b7557fe5b602002602001015188600281518110612b8a57fe5b602002602001015189600381518110612b9f57fe5b60200260200101516141c9565b600091505b80612c425789608001518a6060015114612bfc5760405162461bcd60e51b81526004018080602001828103825260278152602001806154e86027913960400191505060405180910390fd5b8960c001518a60a0015114612c425760405162461bcd60e51b81526004018080602001828103825260268152602001806154446026913960400191505060405180910390fd5b81612ca35760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c08401511415612c9b57612c96836142e8565b612ca3565b60c083015183525b612cac84610ff9565b8a5114612cea5760405162461bcd60e51b81526004018080602001828103825260228152602001806153f96022913960400191505060405180910390fd5b509098975050505050505050565b6000818381612d0357fe5b06828481612d0d57fe5b04019392505050565b612d1e6153a5565b60408051600080825260208201909252612d5891612d52565b612d3f6153a5565b815260200190600190039081612d375790505b506142f2565b905090565b6000818381612d6857fe5b049392505050565b612d7981614404565b60065550565b600580546002919060ff1916600183610f3a565b600090565b60408051602080820193909352815180820384018152908201909152805191012090565b600190565b6000600282604001515110612dd257fe5b604082015151612e3757612de4612dbc565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b909316602185015260228085019190915282518085039091018152604290930190915281519101209050610e23565b612e3f612dbc565b8260000151612e658460400151600081518110612e5857fe5b602002602001015161198f565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600290565b6000612ecd612f14565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b600390565b612f216153a5565b600882511115612f6f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015612f9c578160200160208202803883390190505b508051909150600160005b82811015612fff57612fbe868281518110612e5857fe5b848281518110612fca57fe5b602002602001018181525050858181518110612fe257fe5b602002602001015160800151820191508080600101915050612fa7565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561304457818101518382015260200161302c565b505050509050019250505060405160208183030381529060405280519060200120905061307181836118dc565b9695505050505050565b606490565b600080606061308d615347565b613095615347565b60006130a083614542565b6130af8761010001518261454c565b935090506130bc836145ec565b9150600087610100015182815181106130d157fe5b602001015160f81c60f81b60f81c905087610100015182600101815181106130f557fe5b016020015160f81c9650600061310a88614655565b60408051838152602080850282010190915290985090915081801561314957816020015b6131366153a5565b81526020019060019003908161312e5790505b5095506002830192508160ff166000148061316757508160ff166001145b6131b8576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166131d8576131d1610f56898760000151614bb2565b855261323e565b6131e06153a5565b6131ef8a610100015185613a95565b9094509050811561321857808760008151811061320857fe5b6020026020010181905250613228565b613228858263ffffffff613c4116565b61323a610f568a886000015184614c0d565b8652505b60ff82165b8181101561327d5761325a8a610100015185613a95565b885189908490811061326857fe5b60209081029190910101529350600101613243565b8651156132ca575060005b8260ff168751038110156132ca576132c2878260018a510303815181106132ab57fe5b602002602001015187613c4190919063ffffffff16565b600101613288565b50505091939550919395565b60006132e183614c88565b15806132f357506132f182614c88565b155b156133005750600061331f565b82518251808201613317878263ffffffff614c9316565b600193505050505b9392505050565b600061333183614c88565b1580613343575061334182614c88565b155b156133505750600061331f565b82518251808202613317878263ffffffff614c9316565b600061337283614c88565b1580613384575061338282614c88565b155b156133915750600061331f565b82518251808203613317878263ffffffff614c9316565b60006133b383614c88565b15806133c557506133c382614c88565b155b156133d25750600061331f565b82518251806133e65760009250505061331f565b808204613317878263ffffffff614c9316565b600061340483614c88565b1580613416575061341482614c88565b155b156134235750600061331f565b82518251806134375760009250505061331f565b808205613317878263ffffffff614c9316565b600061345583614c88565b1580613467575061346582614c88565b155b156134745750600061331f565b82518251806134885760009250505061331f565b808206613317878263ffffffff614c9316565b60006134a683614c88565b15806134b857506134b682614c88565b155b156134c55750600061331f565b82518251806134d95760009250505061331f565b808207613317878263ffffffff614c9316565b60006134f784614c88565b1580613509575061350783614c88565b155b1561351657506000610f23565b8351835183518061352d5760009350505050610f23565b6000818385089050613545898263ffffffff614c9316565b50600198975050505050505050565b600061355f84614c88565b1580613571575061356f83614c88565b155b1561357e57506000610f23565b835183518351806135955760009350505050610f23565b6000818385099050613545898263ffffffff614c9316565b60006135b883614c88565b15806135ca57506135c882614c88565b155b156135d75750600061331f565b8251825180820a613317878263ffffffff614c9316565b60006135f983614c88565b158061360b575061360982614c88565b155b156136185750600061331f565b82518251808210613317878263ffffffff614c9316565b600061363a83614c88565b158061364c575061364a82614c88565b155b156136595750600061331f565b82518251808211613317878263ffffffff614c9316565b600061367b83614c88565b158061368d575061368b82614c88565b155b1561369a5750600061331f565b82518251808212613317878263ffffffff614c9316565b60006136bc83614c88565b15806136ce57506136cc82614c88565b155b156136db5750600061331f565b82518251808213613317878263ffffffff614c9316565b600061371f6137126137038461198f565b61370c8661198f565b14614ca9565b859063ffffffff613c4116565b5060019392505050565b600061373482614c88565b61374e5761374983600063ffffffff614c9316565b613765565b81518015613762858263ffffffff614c9316565b50505b50600192915050565b600061377983614c88565b158061378b575061378982614c88565b155b156137985750600061331f565b82518251808216613317878263ffffffff614c9316565b60006137ba83614c88565b15806137cc57506137ca82614c88565b155b156137d95750600061331f565b82518251808217613317878263ffffffff614c9316565b60006137fb83614c88565b158061380d575061380b82614c88565b155b1561381a5750600061331f565b82518251808218613317878263ffffffff614c9316565b600061383c82614c88565b61384857506000611989565b8151801961385c858263ffffffff614c9316565b506001949350505050565b600061387283614c88565b1580613884575061388282614c88565b155b156138915750600061331f565b8251825181811a613317878263ffffffff614c9316565b60006138b383614c88565b15806138c557506138c382614c88565b155b156138d25750600061331f565b8251825181810b613317878263ffffffff614c9316565b60006137656138f78361198f565b849063ffffffff614c9316565b600061376561391283614ccb565b849063ffffffff613c4116565b600061392a83614c88565b158061393c575061393a82614c88565b155b156139495750600061331f565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120613317878263ffffffff614c9316565b600192915050565b60006139a6826080015183613c4190919063ffffffff16565b506001919050565b60006139a6826060015183613c4190919063ffffffff16565b60609190910152600190565b60006139de82614d2a565b6139ea57506000611989565b6139f38261198f565b835250600192915050565b6000613a0983614d2a565b613a155750600061331f565b613a1e82614c88565b613a2a5750600061331f565b81511561371f57613a3a8361198f565b84525060019392505050565b60006139a6613a66613a59610f56612d16565b61370c856020015161198f565b839063ffffffff613c4116565b6000613765613912836001614d37565b6000613765838363ffffffff613c2716565b6000613a9f6153a5565b83518310613ae5576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b600080613af28686614de8565b91509150613afe612d93565b60ff168160ff161415613b32576000613b178784614e0f565b909350905082613b2682614e78565b94509450505050613c20565b613b3a612dbc565b60ff168160ff161415613b5c57613b518683614f2a565b935093505050613c20565b613b64612ebe565b60ff168160ff161415613b7b57613b518683614fcc565b613b83612f14565b60ff168160ff1610158015613ba45750613b9b615060565b60ff168160ff16105b15613be0576000613bb3612f14565b820390506060613bc4828986615065565b909450905083613bd3826142f2565b9550955050505050613c20565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b613c358260400151826150fe565b82604001819052505050565b613c4f8260200151826150fe565b82602001819052505050565b60006139a6613a66613c6e610f56612d16565b61370c856040015161198f565b60006139a6613a668360c001516001614d37565b6000613c9a82614d2a565b613ca657506000611989565b613caf8261198f565b60c084015250600192915050565b6000613ccf838363ffffffff613c4116565b613765838363ffffffff613c4116565b6000613cf1848363ffffffff613c4116565b613d01848463ffffffff613c4116565b61371f848363ffffffff613c4116565b6000613d23858363ffffffff613c4116565b613d33858463ffffffff613c4116565b613d43858563ffffffff613c4116565b61385c858363ffffffff613c4116565b6000613d01848463ffffffff613c4116565b6000613d77858563ffffffff613c4116565b613d43858463ffffffff613c4116565b6000613d9283614c88565b1580613da45750613da282615174565b155b15613db15750600061331f565b613dba82615181565b60ff16836000015110613dcf5750600061331f565b61371f8260400151846000015181518110613de657fe5b602002602001015185613c4190919063ffffffff16565b6000613e0883615174565b1580613e1a5750613e1884614c88565b155b15613e2757506000610f23565b613e3083615181565b60ff16846000015110613e4557506000610f23565b604083015184518151849183918110613e5a57fe5b6020026020010181905250613e7e613e71826142f2565b879063ffffffff613c4116565b50600195945050505050565b6000613e9582615174565b613ea157506000611989565b613765613ead83615181565b849060ff1663ffffffff614c9316565b6000613ec883614c88565b1580613eda5750613ed882615174565b155b15613ee75750600061331f565b613ef082615181565b60ff16836000015110613f055750600061331f565b613dcf848363ffffffff613c2716565b6000613f2082615174565b1580613f325750613f3084614c88565b155b15613f3f57506000610f23565b613f4882615181565b60ff16846000015110613f5d57506000610f23565b604082015184518151859183918110613f7257fe5b6020026020010181905250613e7e613f89826142f2565b879063ffffffff613c2716565b50600190565b6000806001613faa8461198f565b915091509250929050565b60008061271083608001511115613fd157506000905080613c20565b613fda836151a8565b613fe957506000905080613c20565b6001613faa8461198f565b6000614001610f56612d16565b61400a8361198f565b1415613ccf576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b600061407282614c88565b61407e57506000611989565b505160a09190910152600190565b60006139a68260a0015183614c9390919063ffffffff16565b60408051600160f81b602080830191909152600060218301819052602280840182905284518085039091018152604290930190935281519101206139a690613a66906001614d37565b60006140f983614c88565b6141055750600061331f565b61410e82614d2a565b61411a5750600061331f565b61371f613712846000015161412e8561198f565b614bb2565b600061413e84614c88565b61414a57506000610f23565b61415382614d2a565b61415f57506000610f23565b61385c61417985600001516141738561198f565b86614c0d565b869063ffffffff613c4116565b60408051600080825260208201909252606090826141ba565b6141a76153a5565b81526020019060019003908161419f5790505b509050613765613912826142f2565b60006141d485614c88565b15806141e657506141e484614c88565b155b806141f757506141f583614c88565b155b80614208575061420682614c88565b155b15614215575060006142df565b8451845184511580159061422b57508451600114155b1561424c5761424188600063ffffffff614c9316565b6001925050506142df565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa1580156142ae573d6000803e3d6000fd5b5050604051601f19015191506142d590508b6001600160a01b03831663ffffffff614c9316565b6001955050505050505b95945050505050565b600160e090910152565b6142fa6153a5565b61430482516152a6565b614355576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561438c5783818151811061436f57fe5b60200260200101516080015182019150808060010191505061435a565b506040805160a08101825260008082528251606081018452818152602081810183905284518381528082018652939490850193919290830191906143e6565b6143d36153a5565b8152602001906001900390816143cb5790505b50905281526020810194909452600360408501526060909301525090565b6000815b600181511115614525576060600282516001018161442257fe5b0460405190808252806020026020018201604052801561444c578160200160208202803883390190505b50905060005b815181101561451d5782518160020260010110156144e55782816002028151811061447957fe5b602002602001015183826002026001018151811061449357fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106144d457fe5b602002602001018181525050614515565b8281600202815181106144f457fe5b602002602001015182828151811061450857fe5b6020026020010181815250505b600101614452565b509050614408565b8060008151811061453257fe5b6020026020010151915050919050565b600060e090910152565b6000614556615347565b61455e615347565b600060e08201819052806145728787614e0f565b90965091506145818787614fcc565b602085015295506145928787614fcc565b604085015295506145a38787613a95565b606085015295506145b48787613a95565b608085015295506145c58787614e0f565b60a085015295506145d68787614e0f565b92845260c0840192909252509590945092505050565b6145f4615347565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b600080600183141561466d5750600290506003614bad565b60028314156146825750600290506003614bad565b60038314156146975750600290506003614bad565b60048314156146ac5750600290506004614bad565b60058314156146c15750600290506007614bad565b60068314156146d65750600290506004614bad565b60078314156146eb5750600290506007614bad565b60088314156147005750600390506004614bad565b60098314156147155750600390506004614bad565b600a83141561472a5750600290506019614bad565b601083141561473e57506002905080614bad565b601183141561475257506002905080614bad565b601283141561476657506002905080614bad565b601383141561477a57506002905080614bad565b601483141561478e57506002905080614bad565b60158314156147a257506001905080614bad565b60168314156147b657506002905080614bad565b60178314156147ca57506002905080614bad565b60188314156147de57506002905080614bad565b60198314156147f257506001905080614bad565b601a8314156148075750600290506004614bad565b601b83141561481c5750600290506007614bad565b60208314156148315750600190506007614bad565b60218314156148465750600190506003614bad565b602283141561485b5750600290506008614bad565b603083141561486f57506001905080614bad565b60318314156148845750600090506001614bad565b60328314156148995750600090506001614bad565b60338314156148ae5750600190506002614bad565b60348314156148c35750600190506004614bad565b60358314156148d85750600290506004614bad565b60368314156148ed5750600090506002614bad565b60378314156149025750600090506001614bad565b603883141561491657506001905080614bad565b603983141561492b5750600090506001614bad565b603a8314156149405750600090506002614bad565b603b8314156149555750600090506001614bad565b603c83141561496a5750600090506001614bad565b603d83141561497e57506001905080614bad565b604083141561499257506001905080614bad565b60418314156149a75750600290506001614bad565b60428314156149bc5750600390506001614bad565b60438314156149d15750600290506001614bad565b60448314156149e65750600390506001614bad565b60508314156149fa57506002905080614bad565b6051831415614a0f5750600390506028614bad565b6052831415614a245750600190506002614bad565b6053831415614a395750600190506003614bad565b6054831415614a4e5750600290506029614bad565b6060831415614a635750600090506064614bad565b6061831415614a785750600190506064614bad565b6070831415614a8d5750600190506064614bad565b6072831415614aa25750600090506028614bad565b6073831415614ab75750600090506005614bad565b6074831415614acc575060009050600a614bad565b6075831415614ae15750600190506000614bad565b6076831415614af65750600090506001614bad565b6077831415614b0b5750600090506019614bad565b6078831415614b205750600290506019614bad565b6079831415614b355750600390506019614bad565b607b831415614b4a575060009050600a614bad565b6080831415614b60575060049050614e20614bad565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b614bba6153a5565b6040805160608101825260ff85168152602080820185905282516000808252918101845261331f93830191614c05565b614bf26153a5565b815260200190600190039081614bea5790505b5090526152ad565b614c156153a5565b604080516001808252818301909252606091816020015b614c346153a5565b815260200190600190039081614c2c5790505090508281600081518110614c5757fe5b60200260200101819052506142df60405180606001604052808760ff168152602001868152602001838152506152ad565b6060015160ff161590565b613c4f8260200151614ca483614e78565b6150fe565b614cb16153a5565b8115614cc1576119b56001614e78565b6119b56000614e78565b614cd36153a5565b816060015160ff1660021415614d1a5760405162461bcd60e51b815260040180806020018281038252602181526020018061546a6021913960400191505060405180910390fd5b611989826060015160ff16614e78565b6060015160ff1660011490565b614d3f6153a5565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191614d95565b614d826153a5565b815260200190600190039081614d7a5790505b50905281526040805160008082526020828101909352919092019190614dd1565b614dbe6153a5565b815260200190600190039081614db65790505b508152606460208201526040019290925250919050565b60008082600101848481518110614dfb57fe5b016020015190925060f81c90509250929050565b60008082845110158015614e27575060208385510310155b614e64576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301613faa858563ffffffff61531416565b614e806153a5565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191614ed6565b614ec36153a5565b815260200190600190039081614ebb5790505b50905281526040805160008082526020828101909352919092019190614f12565b614eff6153a5565b815260200190600190039081614ef75790505b50815260006020820152600160409091015292915050565b6000614f346153a5565b82600080614f406153a5565b6000614f4c8986614de8565b9095509350614f5b8986614de8565b9095509250600160ff85161415614f7c57614f768986613a95565b90955091505b614f868986615330565b9095509050600160ff85161415614fb15784614fa3848385614c0d565b965096505050505050613c20565b84614fbc8483614bb2565b9650965050505050509250929050565b6000614fd66153a5565b82845110158015614feb575060408385510310155b615027576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000806150348686615330565b90945091506150438685614e0f565b90945090508361505383836118dc565b9350935050509250929050565b600c90565b60006060600083905060608660ff166040519080825280602002602001820160405280156150ad57816020015b61509a6153a5565b8152602001906001900390816150925790505b50905060005b8760ff168160ff1610156150f1576150cb8784613a95565b8351849060ff85169081106150dc57fe5b602090810291909101015292506001016150b3565b5090969095509350505050565b6151066153a5565b6040805160028082526060828101909352816020015b6151246153a5565b81526020019060019003908161511c579050509050828160008151811061514757fe5b6020026020010181905250838160018151811061516057fe5b6020026020010181905250610f2381612f19565b6060015160ff1660031490565b606081015160009060ff16600314156151a05750604081015151610e23565b506001610e23565b606081015160009060ff166151bf57506001610e23565b606082015160ff16600114156151d757506000610e23565b606082015160ff166002141561522b576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b606082015160ff166003141561528e5760408201515160005b818110156152835761526c8460400151828151811061525f57fe5b60200260200101516151a8565b61527b57600092505050610e23565b600101615244565b506001915050610e23565b606082015160ff1660641415611a6e57506000610e23565b6008101590565b6152b56153a5565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906152fc565b6152e96153a5565b8152602001906001900390816152e15790505b50815260016020820181905260409091015292915050565b6000816020018351101561532757600080fd5b50016020015190565b60008060208301613faa858563ffffffff61531416565b60408051610100810190915260008152602081016153636153a5565b81526020016153706153a5565b815260200161537d6153a5565b815260200161538a6153a5565b81526000602082018190526040820181905260609091015290565b6040518060a00160405280600081526020016153bf6153d9565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726fa265627a7a72315820f7b50dfe606c94c24eda0f5405b7cea23a8370ae10a74b5c7bb8330058c2c0b364736f6c63430005110032"

// DeployExecutionChallenge deploys a new Ethereum contract, binding an instance of ExecutionChallenge to it.
func DeployExecutionChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExecutionChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExecutionChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutionChallenge{ExecutionChallengeCaller: ExecutionChallengeCaller{contract: contract}, ExecutionChallengeTransactor: ExecutionChallengeTransactor{contract: contract}, ExecutionChallengeFilterer: ExecutionChallengeFilterer{contract: contract}}, nil
}

// ExecutionChallenge is an auto generated Go binding around an Ethereum contract.
type ExecutionChallenge struct {
	ExecutionChallengeCaller     // Read-only binding to the contract
	ExecutionChallengeTransactor // Write-only binding to the contract
	ExecutionChallengeFilterer   // Log filterer for contract events
}

// ExecutionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionChallengeSession struct {
	Contract     *ExecutionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionChallengeCallerSession struct {
	Contract *ExecutionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExecutionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionChallengeTransactorSession struct {
	Contract     *ExecutionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExecutionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionChallengeRaw struct {
	Contract *ExecutionChallenge // Generic contract binding to access the raw methods on
}

// ExecutionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionChallengeCallerRaw struct {
	Contract *ExecutionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionChallengeTransactorRaw struct {
	Contract *ExecutionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionChallenge creates a new instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallenge(address common.Address, backend bind.ContractBackend) (*ExecutionChallenge, error) {
	contract, err := bindExecutionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallenge{ExecutionChallengeCaller: ExecutionChallengeCaller{contract: contract}, ExecutionChallengeTransactor: ExecutionChallengeTransactor{contract: contract}, ExecutionChallengeFilterer: ExecutionChallengeFilterer{contract: contract}}, nil
}

// NewExecutionChallengeCaller creates a new read-only instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeCaller(address common.Address, caller bind.ContractCaller) (*ExecutionChallengeCaller, error) {
	contract, err := bindExecutionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeCaller{contract: contract}, nil
}

// NewExecutionChallengeTransactor creates a new write-only instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionChallengeTransactor, error) {
	contract, err := bindExecutionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeTransactor{contract: contract}, nil
}

// NewExecutionChallengeFilterer creates a new log filterer instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionChallengeFilterer, error) {
	contract, err := bindExecutionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeFilterer{contract: contract}, nil
}

// bindExecutionChallenge binds a generic wrapper to an already deployed contract.
func bindExecutionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionChallenge *ExecutionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExecutionChallenge.Contract.ExecutionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionChallenge *ExecutionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ExecutionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionChallenge *ExecutionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ExecutionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionChallenge *ExecutionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExecutionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionChallenge *ExecutionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionChallenge *ExecutionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.contract.Transact(opts, method, params...)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xeb83afbf.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, bytes32[] _machineHashes, bool[] _didInboxInsns, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) BisectAssertion(opts *bind.TransactOpts, _beforeInbox [32]byte, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "bisectAssertion", _beforeInbox, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xeb83afbf.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, bytes32[] _machineHashes, bool[] _didInboxInsns, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) BisectAssertion(_beforeInbox [32]byte, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.BisectAssertion(&_ExecutionChallenge.TransactOpts, _beforeInbox, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xeb83afbf.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, bytes32[] _machineHashes, bool[] _didInboxInsns, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) BisectAssertion(_beforeInbox [32]byte, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.BisectAssertion(&_ExecutionChallenge.TransactOpts, _beforeInbox, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) ChooseSegment(opts *bind.TransactOpts, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ChooseSegment(&_ExecutionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ChooseSegment(&_ExecutionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "initializeBisection", _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.InitializeBisection(&_ExecutionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.InitializeBisection(&_ExecutionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x82498b81.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint256 _beforeInboxValueSize, bytes32 _afterHash, bool _didInboxInsns, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProof(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "oneStepProof", _beforeHash, _beforeInbox, _beforeInboxValueSize, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x82498b81.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint256 _beforeInboxValueSize, bytes32 _afterHash, bool _didInboxInsns, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProof(&_ExecutionChallenge.TransactOpts, _beforeHash, _beforeInbox, _beforeInboxValueSize, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x82498b81.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint256 _beforeInboxValueSize, bytes32 _afterHash, bool _didInboxInsns, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProof(&_ExecutionChallenge.TransactOpts, _beforeHash, _beforeInbox, _beforeInboxValueSize, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.TimeoutChallenge(&_ExecutionChallenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.TimeoutChallenge(&_ExecutionChallenge.TransactOpts)
}

// ExecutionChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the ExecutionChallenge contract.
type ExecutionChallengeAsserterTimedOutIterator struct {
	Event *ExecutionChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeAsserterTimedOut)
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
		it.Event = new(ExecutionChallengeAsserterTimedOut)
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
func (it *ExecutionChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the ExecutionChallenge contract.
type ExecutionChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*ExecutionChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeAsserterTimedOutIterator{contract: _ExecutionChallenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeAsserterTimedOut)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*ExecutionChallengeAsserterTimedOut, error) {
	event := new(ExecutionChallengeAsserterTimedOut)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ExecutionChallenge contract.
type ExecutionChallengeBisectedAssertionIterator struct {
	Event *ExecutionChallengeBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeBisectedAssertion)
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
		it.Event = new(ExecutionChallengeBisectedAssertion)
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
func (it *ExecutionChallengeBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeBisectedAssertion represents a BisectedAssertion event raised by the ExecutionChallenge contract.
type ExecutionChallengeBisectedAssertion struct {
	MachineHashes [][32]byte
	DidInboxInsns []bool
	MessageAccs   [][32]byte
	LogAccs       [][32]byte
	Gases         []uint64
	TotalSteps    uint64
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0x99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee47.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bool[] didInboxInsns, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint64 totalSteps, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*ExecutionChallengeBisectedAssertionIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeBisectedAssertionIterator{contract: _ExecutionChallenge.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0x99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee47.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bool[] didInboxInsns, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint64 totalSteps, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeBisectedAssertion)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0x99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee47.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bool[] didInboxInsns, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint64 totalSteps, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseBisectedAssertion(log types.Log) (*ExecutionChallengeBisectedAssertion, error) {
	event := new(ExecutionChallengeBisectedAssertion)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the ExecutionChallenge contract.
type ExecutionChallengeChallengerTimedOutIterator struct {
	Event *ExecutionChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeChallengerTimedOut)
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
		it.Event = new(ExecutionChallengeChallengerTimedOut)
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
func (it *ExecutionChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the ExecutionChallenge contract.
type ExecutionChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*ExecutionChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeChallengerTimedOutIterator{contract: _ExecutionChallenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeChallengerTimedOut)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*ExecutionChallengeChallengerTimedOut, error) {
	event := new(ExecutionChallengeChallengerTimedOut)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeContinuedIterator is returned from FilterContinued and is used to iterate over the raw logs and unpacked data for Continued events raised by the ExecutionChallenge contract.
type ExecutionChallengeContinuedIterator struct {
	Event *ExecutionChallengeContinued // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeContinued)
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
		it.Event = new(ExecutionChallengeContinued)
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
func (it *ExecutionChallengeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeContinued represents a Continued event raised by the ExecutionChallenge contract.
type ExecutionChallengeContinued struct {
	SegmentIndex  *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContinued is a free log retrieval operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterContinued(opts *bind.FilterOpts) (*ExecutionChallengeContinuedIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeContinuedIterator{contract: _ExecutionChallenge.contract, event: "Continued", logs: logs, sub: sub}, nil
}

// WatchContinued is a free log subscription operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchContinued(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeContinued) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeContinued)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
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
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseContinued(log types.Log) (*ExecutionChallengeContinued, error) {
	event := new(ExecutionChallengeContinued)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ExecutionChallenge contract.
type ExecutionChallengeInitiatedChallengeIterator struct {
	Event *ExecutionChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeInitiatedChallenge)
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
		it.Event = new(ExecutionChallengeInitiatedChallenge)
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
func (it *ExecutionChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the ExecutionChallenge contract.
type ExecutionChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ExecutionChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeInitiatedChallengeIterator{contract: _ExecutionChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeInitiatedChallenge)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ExecutionChallengeInitiatedChallenge, error) {
	event := new(ExecutionChallengeInitiatedChallenge)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ExecutionChallenge contract.
type ExecutionChallengeOneStepProofCompletedIterator struct {
	Event *ExecutionChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeOneStepProofCompleted)
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
		it.Event = new(ExecutionChallengeOneStepProofCompleted)
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
func (it *ExecutionChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the ExecutionChallenge contract.
type ExecutionChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*ExecutionChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeOneStepProofCompletedIterator{contract: _ExecutionChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeOneStepProofCompleted)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*ExecutionChallengeOneStepProofCompleted, error) {
	event := new(ExecutionChallengeOneStepProofCompleted)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HashingABI is the input ABI used to generate the binding from.
const HashingABI = "[]"

// HashingBin is the compiled bytecode used for deploying new contracts.
var HashingBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200a16caa6b13b51d9a0a96faa7003a2f51f1e55db1ef9dd57c37ab46cee5652c964736f6c63430005110032"

// DeployHashing deploys a new Ethereum contract, binding an instance of Hashing to it.
func DeployHashing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hashing, error) {
	parsed, err := abi.JSON(strings.NewReader(HashingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hashing{HashingCaller: HashingCaller{contract: contract}, HashingTransactor: HashingTransactor{contract: contract}, HashingFilterer: HashingFilterer{contract: contract}}, nil
}

// Hashing is an auto generated Go binding around an Ethereum contract.
type Hashing struct {
	HashingCaller     // Read-only binding to the contract
	HashingTransactor // Write-only binding to the contract
	HashingFilterer   // Log filterer for contract events
}

// HashingCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashingSession struct {
	Contract     *Hashing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashingCallerSession struct {
	Contract *HashingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HashingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashingTransactorSession struct {
	Contract     *HashingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HashingRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashingRaw struct {
	Contract *Hashing // Generic contract binding to access the raw methods on
}

// HashingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashingCallerRaw struct {
	Contract *HashingCaller // Generic read-only contract binding to access the raw methods on
}

// HashingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashingTransactorRaw struct {
	Contract *HashingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashing creates a new instance of Hashing, bound to a specific deployed contract.
func NewHashing(address common.Address, backend bind.ContractBackend) (*Hashing, error) {
	contract, err := bindHashing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hashing{HashingCaller: HashingCaller{contract: contract}, HashingTransactor: HashingTransactor{contract: contract}, HashingFilterer: HashingFilterer{contract: contract}}, nil
}

// NewHashingCaller creates a new read-only instance of Hashing, bound to a specific deployed contract.
func NewHashingCaller(address common.Address, caller bind.ContractCaller) (*HashingCaller, error) {
	contract, err := bindHashing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashingCaller{contract: contract}, nil
}

// NewHashingTransactor creates a new write-only instance of Hashing, bound to a specific deployed contract.
func NewHashingTransactor(address common.Address, transactor bind.ContractTransactor) (*HashingTransactor, error) {
	contract, err := bindHashing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashingTransactor{contract: contract}, nil
}

// NewHashingFilterer creates a new log filterer instance of Hashing, bound to a specific deployed contract.
func NewHashingFilterer(address common.Address, filterer bind.ContractFilterer) (*HashingFilterer, error) {
	contract, err := bindHashing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashingFilterer{contract: contract}, nil
}

// bindHashing binds a generic wrapper to an already deployed contract.
func bindHashing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashing *HashingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hashing.Contract.HashingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashing *HashingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashing.Contract.HashingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashing *HashingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashing.Contract.HashingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashing *HashingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hashing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashing *HashingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashing *HashingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashing.Contract.contract.Transact(opts, method, params...)
}

// IBisectionChallengeABI is the input ABI used to generate the binding from.
const IBisectionChallengeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBisectionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var IBisectionChallengeFuncSigs = map[string]string{
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
}

// IBisectionChallenge is an auto generated Go binding around an Ethereum contract.
type IBisectionChallenge struct {
	IBisectionChallengeCaller     // Read-only binding to the contract
	IBisectionChallengeTransactor // Write-only binding to the contract
	IBisectionChallengeFilterer   // Log filterer for contract events
}

// IBisectionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBisectionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBisectionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBisectionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBisectionChallengeSession struct {
	Contract     *IBisectionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IBisectionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBisectionChallengeCallerSession struct {
	Contract *IBisectionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IBisectionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBisectionChallengeTransactorSession struct {
	Contract     *IBisectionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IBisectionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBisectionChallengeRaw struct {
	Contract *IBisectionChallenge // Generic contract binding to access the raw methods on
}

// IBisectionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBisectionChallengeCallerRaw struct {
	Contract *IBisectionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// IBisectionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBisectionChallengeTransactorRaw struct {
	Contract *IBisectionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBisectionChallenge creates a new instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallenge(address common.Address, backend bind.ContractBackend) (*IBisectionChallenge, error) {
	contract, err := bindIBisectionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallenge{IBisectionChallengeCaller: IBisectionChallengeCaller{contract: contract}, IBisectionChallengeTransactor: IBisectionChallengeTransactor{contract: contract}, IBisectionChallengeFilterer: IBisectionChallengeFilterer{contract: contract}}, nil
}

// NewIBisectionChallengeCaller creates a new read-only instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeCaller(address common.Address, caller bind.ContractCaller) (*IBisectionChallengeCaller, error) {
	contract, err := bindIBisectionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeCaller{contract: contract}, nil
}

// NewIBisectionChallengeTransactor creates a new write-only instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBisectionChallengeTransactor, error) {
	contract, err := bindIBisectionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeTransactor{contract: contract}, nil
}

// NewIBisectionChallengeFilterer creates a new log filterer instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBisectionChallengeFilterer, error) {
	contract, err := bindIBisectionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeFilterer{contract: contract}, nil
}

// bindIBisectionChallenge binds a generic wrapper to an already deployed contract.
func bindIBisectionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBisectionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBisectionChallenge *IBisectionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBisectionChallenge.Contract.IBisectionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBisectionChallenge *IBisectionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.IBisectionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBisectionChallenge *IBisectionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.IBisectionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBisectionChallenge *IBisectionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBisectionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBisectionChallenge *IBisectionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBisectionChallenge *IBisectionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.contract.Transact(opts, method, params...)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.contract.Transact(opts, "initializeBisection", _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.InitializeBisection(&_IBisectionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeTransactorSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.InitializeBisection(&_IBisectionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// IStakingABI is the input ABI used to generate the binding from.
const IStakingABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IStakingFuncSigs maps the 4-byte function signature to its string representation.
var IStakingFuncSigs = map[string]string{
	"396f51cf": "resolveChallenge(address,address)",
}

// IStaking is an auto generated Go binding around an Ethereum contract.
type IStaking struct {
	IStakingCaller     // Read-only binding to the contract
	IStakingTransactor // Write-only binding to the contract
	IStakingFilterer   // Log filterer for contract events
}

// IStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingSession struct {
	Contract     *IStaking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingCallerSession struct {
	Contract *IStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingTransactorSession struct {
	Contract     *IStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingRaw struct {
	Contract *IStaking // Generic contract binding to access the raw methods on
}

// IStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingCallerRaw struct {
	Contract *IStakingCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingTransactorRaw struct {
	Contract *IStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStaking creates a new instance of IStaking, bound to a specific deployed contract.
func NewIStaking(address common.Address, backend bind.ContractBackend) (*IStaking, error) {
	contract, err := bindIStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStaking{IStakingCaller: IStakingCaller{contract: contract}, IStakingTransactor: IStakingTransactor{contract: contract}, IStakingFilterer: IStakingFilterer{contract: contract}}, nil
}

// NewIStakingCaller creates a new read-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingCaller(address common.Address, caller bind.ContractCaller) (*IStakingCaller, error) {
	contract, err := bindIStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingCaller{contract: contract}, nil
}

// NewIStakingTransactor creates a new write-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingTransactor, error) {
	contract, err := bindIStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingTransactor{contract: contract}, nil
}

// NewIStakingFilterer creates a new log filterer instance of IStaking, bound to a specific deployed contract.
func NewIStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingFilterer, error) {
	contract, err := bindIStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingFilterer{contract: contract}, nil
}

// bindIStaking binds a generic wrapper to an already deployed contract.
func bindIStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.IStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transact(opts, method, params...)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_IStaking *IStakingTransactor) ResolveChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "resolveChallenge", winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_IStaking *IStakingSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IStaking.Contract.ResolveChallenge(&_IStaking.TransactOpts, winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_IStaking *IStakingTransactorSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IStaking.Contract.ResolveChallenge(&_IStaking.TransactOpts, winner, loser)
}

// MachineABI is the input ABI used to generate the binding from.
const MachineABI = "[]"

// MachineBin is the compiled bytecode used for deploying new contracts.
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820619a01b5090ea8892065e467fd2f900ee3ca98aec0a583616b484cbfafcf307e64736f6c63430005110032"

// DeployMachine deploys a new Ethereum contract, binding an instance of Machine to it.
func DeployMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Machine, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// Machine is an auto generated Go binding around an Ethereum contract.
type Machine struct {
	MachineCaller     // Read-only binding to the contract
	MachineTransactor // Write-only binding to the contract
	MachineFilterer   // Log filterer for contract events
}

// MachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineSession struct {
	Contract     *Machine          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineCallerSession struct {
	Contract *MachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineTransactorSession struct {
	Contract     *MachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineRaw struct {
	Contract *Machine // Generic contract binding to access the raw methods on
}

// MachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineCallerRaw struct {
	Contract *MachineCaller // Generic read-only contract binding to access the raw methods on
}

// MachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineTransactorRaw struct {
	Contract *MachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachine creates a new instance of Machine, bound to a specific deployed contract.
func NewMachine(address common.Address, backend bind.ContractBackend) (*Machine, error) {
	contract, err := bindMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// NewMachineCaller creates a new read-only instance of Machine, bound to a specific deployed contract.
func NewMachineCaller(address common.Address, caller bind.ContractCaller) (*MachineCaller, error) {
	contract, err := bindMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineCaller{contract: contract}, nil
}

// NewMachineTransactor creates a new write-only instance of Machine, bound to a specific deployed contract.
func NewMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineTransactor, error) {
	contract, err := bindMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTransactor{contract: contract}, nil
}

// NewMachineFilterer creates a new log filterer instance of Machine, bound to a specific deployed contract.
func NewMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineFilterer, error) {
	contract, err := bindMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineFilterer{contract: contract}, nil
}

// bindMachine binds a generic wrapper to an already deployed contract.
func bindMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.MachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transact(opts, method, params...)
}

// MarshalingABI is the input ABI used to generate the binding from.
const MarshalingABI = "[]"

// MarshalingBin is the compiled bytecode used for deploying new contracts.
var MarshalingBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209b0c1d3b5eda3a21d7269632ba939a0fb3ac18f738de2c9350fb50f42ecd3abe64736f6c63430005110032"

// DeployMarshaling deploys a new Ethereum contract, binding an instance of Marshaling to it.
func DeployMarshaling(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Marshaling, error) {
	parsed, err := abi.JSON(strings.NewReader(MarshalingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarshalingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Marshaling{MarshalingCaller: MarshalingCaller{contract: contract}, MarshalingTransactor: MarshalingTransactor{contract: contract}, MarshalingFilterer: MarshalingFilterer{contract: contract}}, nil
}

// Marshaling is an auto generated Go binding around an Ethereum contract.
type Marshaling struct {
	MarshalingCaller     // Read-only binding to the contract
	MarshalingTransactor // Write-only binding to the contract
	MarshalingFilterer   // Log filterer for contract events
}

// MarshalingCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarshalingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarshalingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarshalingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarshalingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarshalingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarshalingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarshalingSession struct {
	Contract     *Marshaling       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarshalingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarshalingCallerSession struct {
	Contract *MarshalingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MarshalingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarshalingTransactorSession struct {
	Contract     *MarshalingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MarshalingRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarshalingRaw struct {
	Contract *Marshaling // Generic contract binding to access the raw methods on
}

// MarshalingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarshalingCallerRaw struct {
	Contract *MarshalingCaller // Generic read-only contract binding to access the raw methods on
}

// MarshalingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarshalingTransactorRaw struct {
	Contract *MarshalingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarshaling creates a new instance of Marshaling, bound to a specific deployed contract.
func NewMarshaling(address common.Address, backend bind.ContractBackend) (*Marshaling, error) {
	contract, err := bindMarshaling(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marshaling{MarshalingCaller: MarshalingCaller{contract: contract}, MarshalingTransactor: MarshalingTransactor{contract: contract}, MarshalingFilterer: MarshalingFilterer{contract: contract}}, nil
}

// NewMarshalingCaller creates a new read-only instance of Marshaling, bound to a specific deployed contract.
func NewMarshalingCaller(address common.Address, caller bind.ContractCaller) (*MarshalingCaller, error) {
	contract, err := bindMarshaling(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarshalingCaller{contract: contract}, nil
}

// NewMarshalingTransactor creates a new write-only instance of Marshaling, bound to a specific deployed contract.
func NewMarshalingTransactor(address common.Address, transactor bind.ContractTransactor) (*MarshalingTransactor, error) {
	contract, err := bindMarshaling(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarshalingTransactor{contract: contract}, nil
}

// NewMarshalingFilterer creates a new log filterer instance of Marshaling, bound to a specific deployed contract.
func NewMarshalingFilterer(address common.Address, filterer bind.ContractFilterer) (*MarshalingFilterer, error) {
	contract, err := bindMarshaling(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarshalingFilterer{contract: contract}, nil
}

// bindMarshaling binds a generic wrapper to an already deployed contract.
func bindMarshaling(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarshalingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marshaling *MarshalingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Marshaling.Contract.MarshalingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marshaling *MarshalingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marshaling.Contract.MarshalingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marshaling *MarshalingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marshaling.Contract.MarshalingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marshaling *MarshalingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Marshaling.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marshaling *MarshalingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marshaling.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marshaling *MarshalingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marshaling.Contract.contract.Transact(opts, method, params...)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[]"

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
var MerkleLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582091f0cf8a8d490f7587e2923ffc15a0c546101f3531c51caa0f561f62b1abd9b464736f6c63430005110032"

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[]"

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205d3d78ee395b8872deae421ff902768b2d07b8c2a9169f4535aff334d36d9e9c64736f6c63430005110032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// OneStepProof is an auto generated Go binding around an Ethereum contract.
type OneStepProof struct {
	OneStepProofCaller     // Read-only binding to the contract
	OneStepProofTransactor // Write-only binding to the contract
	OneStepProofFilterer   // Log filterer for contract events
}

// OneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofSession struct {
	Contract     *OneStepProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofCallerSession struct {
	Contract *OneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTransactorSession struct {
	Contract     *OneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofRaw struct {
	Contract *OneStepProof // Generic contract binding to access the raw methods on
}

// OneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofCallerRaw struct {
	Contract *OneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTransactorRaw struct {
	Contract *OneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof creates a new instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProof(address common.Address, backend bind.ContractBackend) (*OneStepProof, error) {
	contract, err := bindOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// NewOneStepProofCaller creates a new read-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofCaller, error) {
	contract, err := bindOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofCaller{contract: contract}, nil
}

// NewOneStepProofTransactor creates a new write-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTransactor, error) {
	contract, err := bindOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTransactor{contract: contract}, nil
}

// NewOneStepProofFilterer creates a new log filterer instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofFilterer, error) {
	contract, err := bindOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofFilterer{contract: contract}, nil
}

// bindOneStepProof binds a generic wrapper to an already deployed contract.
func bindOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.OneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transact(opts, method, params...)
}

// RollupTimeABI is the input ABI used to generate the binding from.
const RollupTimeABI = "[]"

// RollupTimeBin is the compiled bytecode used for deploying new contracts.
var RollupTimeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820bc5e1c1239ee257a8886402d63fd0e30faaddb124d4cfd39eb5105366f8e112b64736f6c63430005110032"

// DeployRollupTime deploys a new Ethereum contract, binding an instance of RollupTime to it.
func DeployRollupTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTime, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// RollupTime is an auto generated Go binding around an Ethereum contract.
type RollupTime struct {
	RollupTimeCaller     // Read-only binding to the contract
	RollupTimeTransactor // Write-only binding to the contract
	RollupTimeFilterer   // Log filterer for contract events
}

// RollupTimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTimeSession struct {
	Contract     *RollupTime       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupTimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTimeCallerSession struct {
	Contract *RollupTimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupTimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTimeTransactorSession struct {
	Contract     *RollupTimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupTimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTimeRaw struct {
	Contract *RollupTime // Generic contract binding to access the raw methods on
}

// RollupTimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTimeCallerRaw struct {
	Contract *RollupTimeCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTimeTransactorRaw struct {
	Contract *RollupTimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTime creates a new instance of RollupTime, bound to a specific deployed contract.
func NewRollupTime(address common.Address, backend bind.ContractBackend) (*RollupTime, error) {
	contract, err := bindRollupTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// NewRollupTimeCaller creates a new read-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeCaller(address common.Address, caller bind.ContractCaller) (*RollupTimeCaller, error) {
	contract, err := bindRollupTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeCaller{contract: contract}, nil
}

// NewRollupTimeTransactor creates a new write-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTimeTransactor, error) {
	contract, err := bindRollupTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeTransactor{contract: contract}, nil
}

// NewRollupTimeFilterer creates a new log filterer instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTimeFilterer, error) {
	contract, err := bindRollupTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTimeFilterer{contract: contract}, nil
}

// bindRollupTime binds a generic wrapper to an already deployed contract.
func bindRollupTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.RollupTimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582074a5343624f96b1b5461fa24f07e16f0444f98c70cb6e0869342f386e3cc720364736f6c63430005110032"

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
