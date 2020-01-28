// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package snapshot

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

// SnapshotABI is the input ABI used to generate the binding from.
const SnapshotABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"stakerLocations\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"snapshot\",\"type\":\"bytes32\"}],\"name\":\"SavedDeadlineStakersSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"latestConfirmed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"snapshot\",\"type\":\"bytes32\"}],\"name\":\"SavedLatestConfirmedSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"snapshot\",\"type\":\"bytes32\"}],\"name\":\"SavedNodeExistsSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"location1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr2\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"location2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"snapshot\",\"type\":\"bytes32\"}],\"name\":\"SavedTwoStakersSnapshot\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMySnapshot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SnapshotFuncSigs maps the 4-byte function signature to its string representation.
var SnapshotFuncSigs = map[string]string{
	"c3f8ae34": "getMySnapshot(uint256)",
}

// SnapshotBin is the compiled bytecode used for deploying new contracts.
var SnapshotBin = "0x6080604052348015600f57600080fd5b5060a98061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c3f8ae3414602d575b600080fd5b604760048036036020811015604157600080fd5b50356059565b60408051918252519081900360200190f35b3360009081526020818152604080832093835292905220549056fea265627a7a72315820300cf0819e98687758ab3df6adb275d761382de6f1d4dd018a518d63261de21a64736f6c63430005100032"

// DeploySnapshot deploys a new Ethereum contract, binding an instance of Snapshot to it.
func DeploySnapshot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Snapshot, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SnapshotBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Snapshot{SnapshotCaller: SnapshotCaller{contract: contract}, SnapshotTransactor: SnapshotTransactor{contract: contract}, SnapshotFilterer: SnapshotFilterer{contract: contract}}, nil
}

// Snapshot is an auto generated Go binding around an Ethereum contract.
type Snapshot struct {
	SnapshotCaller     // Read-only binding to the contract
	SnapshotTransactor // Write-only binding to the contract
	SnapshotFilterer   // Log filterer for contract events
}

// SnapshotCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotSession struct {
	Contract     *Snapshot         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotCallerSession struct {
	Contract *SnapshotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SnapshotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotTransactorSession struct {
	Contract     *SnapshotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SnapshotRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotRaw struct {
	Contract *Snapshot // Generic contract binding to access the raw methods on
}

// SnapshotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotCallerRaw struct {
	Contract *SnapshotCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotTransactorRaw struct {
	Contract *SnapshotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshot creates a new instance of Snapshot, bound to a specific deployed contract.
func NewSnapshot(address common.Address, backend bind.ContractBackend) (*Snapshot, error) {
	contract, err := bindSnapshot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Snapshot{SnapshotCaller: SnapshotCaller{contract: contract}, SnapshotTransactor: SnapshotTransactor{contract: contract}, SnapshotFilterer: SnapshotFilterer{contract: contract}}, nil
}

// NewSnapshotCaller creates a new read-only instance of Snapshot, bound to a specific deployed contract.
func NewSnapshotCaller(address common.Address, caller bind.ContractCaller) (*SnapshotCaller, error) {
	contract, err := bindSnapshot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotCaller{contract: contract}, nil
}

// NewSnapshotTransactor creates a new write-only instance of Snapshot, bound to a specific deployed contract.
func NewSnapshotTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotTransactor, error) {
	contract, err := bindSnapshot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotTransactor{contract: contract}, nil
}

// NewSnapshotFilterer creates a new log filterer instance of Snapshot, bound to a specific deployed contract.
func NewSnapshotFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotFilterer, error) {
	contract, err := bindSnapshot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotFilterer{contract: contract}, nil
}

// bindSnapshot binds a generic wrapper to an already deployed contract.
func bindSnapshot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snapshot *SnapshotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Snapshot.Contract.SnapshotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snapshot *SnapshotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshot.Contract.SnapshotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snapshot *SnapshotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snapshot.Contract.SnapshotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snapshot *SnapshotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Snapshot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snapshot *SnapshotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snapshot *SnapshotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snapshot.Contract.contract.Transact(opts, method, params...)
}

// GetMySnapshot is a free data retrieval call binding the contract method 0xc3f8ae34.
//
// Solidity: function getMySnapshot(uint256 idx) constant returns(bytes32)
func (_Snapshot *SnapshotCaller) GetMySnapshot(opts *bind.CallOpts, idx *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Snapshot.contract.Call(opts, out, "getMySnapshot", idx)
	return *ret0, err
}

// GetMySnapshot is a free data retrieval call binding the contract method 0xc3f8ae34.
//
// Solidity: function getMySnapshot(uint256 idx) constant returns(bytes32)
func (_Snapshot *SnapshotSession) GetMySnapshot(idx *big.Int) ([32]byte, error) {
	return _Snapshot.Contract.GetMySnapshot(&_Snapshot.CallOpts, idx)
}

// GetMySnapshot is a free data retrieval call binding the contract method 0xc3f8ae34.
//
// Solidity: function getMySnapshot(uint256 idx) constant returns(bytes32)
func (_Snapshot *SnapshotCallerSession) GetMySnapshot(idx *big.Int) ([32]byte, error) {
	return _Snapshot.Contract.GetMySnapshot(&_Snapshot.CallOpts, idx)
}

// SnapshotSavedDeadlineStakersSnapshotIterator is returned from FilterSavedDeadlineStakersSnapshot and is used to iterate over the raw logs and unpacked data for SavedDeadlineStakersSnapshot events raised by the Snapshot contract.
type SnapshotSavedDeadlineStakersSnapshotIterator struct {
	Event *SnapshotSavedDeadlineStakersSnapshot // Event containing the contract specifics and raw log

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
func (it *SnapshotSavedDeadlineStakersSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotSavedDeadlineStakersSnapshot)
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
		it.Event = new(SnapshotSavedDeadlineStakersSnapshot)
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
func (it *SnapshotSavedDeadlineStakersSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotSavedDeadlineStakersSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotSavedDeadlineStakersSnapshot represents a SavedDeadlineStakersSnapshot event raised by the Snapshot contract.
type SnapshotSavedDeadlineStakersSnapshot struct {
	Client          common.Address
	DeadlineTicks   *big.Int
	StakerLocations [][32]byte
	Snapshot        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSavedDeadlineStakersSnapshot is a free log retrieval operation binding the contract event 0x74ea644e404c0a911c3b4cceddbacb969cff8a6efc7efae69d1d9be45e6f1983.
//
// Solidity: event SavedDeadlineStakersSnapshot(address client, uint256 deadlineTicks, bytes32[] stakerLocations, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) FilterSavedDeadlineStakersSnapshot(opts *bind.FilterOpts) (*SnapshotSavedDeadlineStakersSnapshotIterator, error) {

	logs, sub, err := _Snapshot.contract.FilterLogs(opts, "SavedDeadlineStakersSnapshot")
	if err != nil {
		return nil, err
	}
	return &SnapshotSavedDeadlineStakersSnapshotIterator{contract: _Snapshot.contract, event: "SavedDeadlineStakersSnapshot", logs: logs, sub: sub}, nil
}

// WatchSavedDeadlineStakersSnapshot is a free log subscription operation binding the contract event 0x74ea644e404c0a911c3b4cceddbacb969cff8a6efc7efae69d1d9be45e6f1983.
//
// Solidity: event SavedDeadlineStakersSnapshot(address client, uint256 deadlineTicks, bytes32[] stakerLocations, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) WatchSavedDeadlineStakersSnapshot(opts *bind.WatchOpts, sink chan<- *SnapshotSavedDeadlineStakersSnapshot) (event.Subscription, error) {

	logs, sub, err := _Snapshot.contract.WatchLogs(opts, "SavedDeadlineStakersSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotSavedDeadlineStakersSnapshot)
				if err := _Snapshot.contract.UnpackLog(event, "SavedDeadlineStakersSnapshot", log); err != nil {
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

// ParseSavedDeadlineStakersSnapshot is a log parse operation binding the contract event 0x74ea644e404c0a911c3b4cceddbacb969cff8a6efc7efae69d1d9be45e6f1983.
//
// Solidity: event SavedDeadlineStakersSnapshot(address client, uint256 deadlineTicks, bytes32[] stakerLocations, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) ParseSavedDeadlineStakersSnapshot(log types.Log) (*SnapshotSavedDeadlineStakersSnapshot, error) {
	event := new(SnapshotSavedDeadlineStakersSnapshot)
	if err := _Snapshot.contract.UnpackLog(event, "SavedDeadlineStakersSnapshot", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SnapshotSavedLatestConfirmedSnapshotIterator is returned from FilterSavedLatestConfirmedSnapshot and is used to iterate over the raw logs and unpacked data for SavedLatestConfirmedSnapshot events raised by the Snapshot contract.
type SnapshotSavedLatestConfirmedSnapshotIterator struct {
	Event *SnapshotSavedLatestConfirmedSnapshot // Event containing the contract specifics and raw log

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
func (it *SnapshotSavedLatestConfirmedSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotSavedLatestConfirmedSnapshot)
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
		it.Event = new(SnapshotSavedLatestConfirmedSnapshot)
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
func (it *SnapshotSavedLatestConfirmedSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotSavedLatestConfirmedSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotSavedLatestConfirmedSnapshot represents a SavedLatestConfirmedSnapshot event raised by the Snapshot contract.
type SnapshotSavedLatestConfirmedSnapshot struct {
	Client          common.Address
	LatestConfirmed [32]byte
	Snapshot        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSavedLatestConfirmedSnapshot is a free log retrieval operation binding the contract event 0x01e422abe4c07a9d56a7f67171e2f76848c67d632c74dcd3a5cbb50c63c22ca2.
//
// Solidity: event SavedLatestConfirmedSnapshot(address client, bytes32 latestConfirmed, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) FilterSavedLatestConfirmedSnapshot(opts *bind.FilterOpts) (*SnapshotSavedLatestConfirmedSnapshotIterator, error) {

	logs, sub, err := _Snapshot.contract.FilterLogs(opts, "SavedLatestConfirmedSnapshot")
	if err != nil {
		return nil, err
	}
	return &SnapshotSavedLatestConfirmedSnapshotIterator{contract: _Snapshot.contract, event: "SavedLatestConfirmedSnapshot", logs: logs, sub: sub}, nil
}

// WatchSavedLatestConfirmedSnapshot is a free log subscription operation binding the contract event 0x01e422abe4c07a9d56a7f67171e2f76848c67d632c74dcd3a5cbb50c63c22ca2.
//
// Solidity: event SavedLatestConfirmedSnapshot(address client, bytes32 latestConfirmed, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) WatchSavedLatestConfirmedSnapshot(opts *bind.WatchOpts, sink chan<- *SnapshotSavedLatestConfirmedSnapshot) (event.Subscription, error) {

	logs, sub, err := _Snapshot.contract.WatchLogs(opts, "SavedLatestConfirmedSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotSavedLatestConfirmedSnapshot)
				if err := _Snapshot.contract.UnpackLog(event, "SavedLatestConfirmedSnapshot", log); err != nil {
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

// ParseSavedLatestConfirmedSnapshot is a log parse operation binding the contract event 0x01e422abe4c07a9d56a7f67171e2f76848c67d632c74dcd3a5cbb50c63c22ca2.
//
// Solidity: event SavedLatestConfirmedSnapshot(address client, bytes32 latestConfirmed, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) ParseSavedLatestConfirmedSnapshot(log types.Log) (*SnapshotSavedLatestConfirmedSnapshot, error) {
	event := new(SnapshotSavedLatestConfirmedSnapshot)
	if err := _Snapshot.contract.UnpackLog(event, "SavedLatestConfirmedSnapshot", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SnapshotSavedNodeExistsSnapshotIterator is returned from FilterSavedNodeExistsSnapshot and is used to iterate over the raw logs and unpacked data for SavedNodeExistsSnapshot events raised by the Snapshot contract.
type SnapshotSavedNodeExistsSnapshotIterator struct {
	Event *SnapshotSavedNodeExistsSnapshot // Event containing the contract specifics and raw log

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
func (it *SnapshotSavedNodeExistsSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotSavedNodeExistsSnapshot)
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
		it.Event = new(SnapshotSavedNodeExistsSnapshot)
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
func (it *SnapshotSavedNodeExistsSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotSavedNodeExistsSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotSavedNodeExistsSnapshot represents a SavedNodeExistsSnapshot event raised by the Snapshot contract.
type SnapshotSavedNodeExistsSnapshot struct {
	Client   common.Address
	NodeHash [32]byte
	Snapshot [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSavedNodeExistsSnapshot is a free log retrieval operation binding the contract event 0x98c23aaa7c1edd19d6d64f4d7938f7b7286545e1cb67d19a6d4d500f75890556.
//
// Solidity: event SavedNodeExistsSnapshot(address client, bytes32 nodeHash, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) FilterSavedNodeExistsSnapshot(opts *bind.FilterOpts) (*SnapshotSavedNodeExistsSnapshotIterator, error) {

	logs, sub, err := _Snapshot.contract.FilterLogs(opts, "SavedNodeExistsSnapshot")
	if err != nil {
		return nil, err
	}
	return &SnapshotSavedNodeExistsSnapshotIterator{contract: _Snapshot.contract, event: "SavedNodeExistsSnapshot", logs: logs, sub: sub}, nil
}

// WatchSavedNodeExistsSnapshot is a free log subscription operation binding the contract event 0x98c23aaa7c1edd19d6d64f4d7938f7b7286545e1cb67d19a6d4d500f75890556.
//
// Solidity: event SavedNodeExistsSnapshot(address client, bytes32 nodeHash, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) WatchSavedNodeExistsSnapshot(opts *bind.WatchOpts, sink chan<- *SnapshotSavedNodeExistsSnapshot) (event.Subscription, error) {

	logs, sub, err := _Snapshot.contract.WatchLogs(opts, "SavedNodeExistsSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotSavedNodeExistsSnapshot)
				if err := _Snapshot.contract.UnpackLog(event, "SavedNodeExistsSnapshot", log); err != nil {
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

// ParseSavedNodeExistsSnapshot is a log parse operation binding the contract event 0x98c23aaa7c1edd19d6d64f4d7938f7b7286545e1cb67d19a6d4d500f75890556.
//
// Solidity: event SavedNodeExistsSnapshot(address client, bytes32 nodeHash, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) ParseSavedNodeExistsSnapshot(log types.Log) (*SnapshotSavedNodeExistsSnapshot, error) {
	event := new(SnapshotSavedNodeExistsSnapshot)
	if err := _Snapshot.contract.UnpackLog(event, "SavedNodeExistsSnapshot", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SnapshotSavedTwoStakersSnapshotIterator is returned from FilterSavedTwoStakersSnapshot and is used to iterate over the raw logs and unpacked data for SavedTwoStakersSnapshot events raised by the Snapshot contract.
type SnapshotSavedTwoStakersSnapshotIterator struct {
	Event *SnapshotSavedTwoStakersSnapshot // Event containing the contract specifics and raw log

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
func (it *SnapshotSavedTwoStakersSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotSavedTwoStakersSnapshot)
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
		it.Event = new(SnapshotSavedTwoStakersSnapshot)
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
func (it *SnapshotSavedTwoStakersSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotSavedTwoStakersSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotSavedTwoStakersSnapshot represents a SavedTwoStakersSnapshot event raised by the Snapshot contract.
type SnapshotSavedTwoStakersSnapshot struct {
	Client    common.Address
	Addr1     common.Address
	Location1 [32]byte
	Addr2     common.Address
	Location2 [32]byte
	Snapshot  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSavedTwoStakersSnapshot is a free log retrieval operation binding the contract event 0x18410fe1e6e0f23683d25f04192e6b95d64681eee6fc2abe5241d2f58cbfbe52.
//
// Solidity: event SavedTwoStakersSnapshot(address client, address addr1, bytes32 location1, address addr2, bytes32 location2, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) FilterSavedTwoStakersSnapshot(opts *bind.FilterOpts) (*SnapshotSavedTwoStakersSnapshotIterator, error) {

	logs, sub, err := _Snapshot.contract.FilterLogs(opts, "SavedTwoStakersSnapshot")
	if err != nil {
		return nil, err
	}
	return &SnapshotSavedTwoStakersSnapshotIterator{contract: _Snapshot.contract, event: "SavedTwoStakersSnapshot", logs: logs, sub: sub}, nil
}

// WatchSavedTwoStakersSnapshot is a free log subscription operation binding the contract event 0x18410fe1e6e0f23683d25f04192e6b95d64681eee6fc2abe5241d2f58cbfbe52.
//
// Solidity: event SavedTwoStakersSnapshot(address client, address addr1, bytes32 location1, address addr2, bytes32 location2, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) WatchSavedTwoStakersSnapshot(opts *bind.WatchOpts, sink chan<- *SnapshotSavedTwoStakersSnapshot) (event.Subscription, error) {

	logs, sub, err := _Snapshot.contract.WatchLogs(opts, "SavedTwoStakersSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotSavedTwoStakersSnapshot)
				if err := _Snapshot.contract.UnpackLog(event, "SavedTwoStakersSnapshot", log); err != nil {
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

// ParseSavedTwoStakersSnapshot is a log parse operation binding the contract event 0x18410fe1e6e0f23683d25f04192e6b95d64681eee6fc2abe5241d2f58cbfbe52.
//
// Solidity: event SavedTwoStakersSnapshot(address client, address addr1, bytes32 location1, address addr2, bytes32 location2, bytes32 snapshot)
func (_Snapshot *SnapshotFilterer) ParseSavedTwoStakersSnapshot(log types.Log) (*SnapshotSavedTwoStakersSnapshot, error) {
	event := new(SnapshotSavedTwoStakersSnapshot)
	if err := _Snapshot.contract.UnpackLog(event, "SavedTwoStakersSnapshot", log); err != nil {
		return nil, err
	}
	return event, nil
}
