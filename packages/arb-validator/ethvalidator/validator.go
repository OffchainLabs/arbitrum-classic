/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ethvalidator

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"time"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/validator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type EthValidator struct {
	key *ecdsa.PrivateKey

	// Safe public interface
	VMID                [32]byte
	Validators          map[common.Address]validatorInfo
	Bot                 *validator.Validator
	actionChan          chan func(*EthValidator) error
	CompletedCallChan   chan valmessage.FinalizedAssertion
	VMCreatedTxHashChan chan [32]byte

	// Not in thread, but internal only
	serverAddress string
	arbAddresses  ethbridge.ArbAddresses

	// private thread only
	con  *ethbridge.Bridge
	auth *bind.TransactOpts
}

func (val *EthValidator) Address() common.Address {
	return crypto.PubkeyToAddress(val.key.PublicKey)
}

func (val *EthValidator) ValidatorCount() int {
	return len(val.Validators)
}

type validatorInfo struct {
	indexNum uint16
}

type VMResponse struct {
	Message protocol.Message
	Result  value.Value
	Proof   [][32]byte
}

func NewEthValidator(
	name string,
	vmID [32]byte,
	machine machine.Machine,
	key *ecdsa.PrivateKey,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	maxCallSteps int32,
	connectionInfo ethbridge.ArbAddresses,
	ethURL string,
) (*EthValidator, error) {
	auth := bind.NewKeyedTransactor(key)

	con, err := ethbridge.New(ethURL, connectionInfo)
	if err != nil {
		return nil, err
	}
	// auth.Value = big.NewInt(10000000)     // in wei
	auth.GasLimit = uint64(0) // in units
	auth.GasPrice = big.NewInt(10)

	nonce, err := con.PendingNonceAt(auth.From)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	manMap := make(map[common.Address]validatorInfo)
	keys := make([]common.Address, 0, len(config.AssertKeys))
	for _, key := range config.AssertKeys {
		var address common.Address
		copy(address[:], key.Value)
		keys = append(keys, address)
	}
	for i, add := range keys {
		manMap[add] = validatorInfo{uint16(i)}
	}

	_, found := manMap[crypto.PubkeyToAddress(key.PublicKey)]
	if !found {
		return nil, errors.New("key is not a validator of chosen VM")
	}

	bot := validator.NewValidator(name, auth.From, protocol.NewBalanceTracker(), config, machine, challengeEverything, maxCallSteps)

	actionChan := make(chan func(*EthValidator) error, 1024)
	completedCallChan := make(chan valmessage.FinalizedAssertion, 1024)
	unanVMCreatedEventTxHashChan := make(chan [32]byte, 1)

	val := &EthValidator{
		key,
		vmID,
		manMap,
		bot,
		actionChan,
		completedCallChan,
		unanVMCreatedEventTxHashChan,
		ethURL,
		connectionInfo,
		con,
		auth,
	}
	return val, nil
}

func (val *EthValidator) Sign(msgHash [32]byte) ([]byte, error) {
	data := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(msgHash))
	return crypto.Sign(data, val.key)
}

func (val *EthValidator) StartListening() error {
	outChan, errChan, err := val.con.CreateListeners(val.VMID)
	if err != nil {
		return err
	}
	parsedChan := make(chan ethbridge.Notification, 1024)

	val.Bot.Run(parsedChan, val)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			select {
			case parse := <-outChan:
				switch parse.Event.(type) {
				case ethbridge.VMCreatedEvent:
					val.VMCreatedTxHashChan <- parse.TxHash
				default:
					parsedChan <- parse
				}
			case event := <-val.actionChan:
				err := event(val)
				if err != nil {
					log.Fatalf("Error handling send: %v", err)
				}
				val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
			case <-errChan:
				// Ignore error and try to reset connection
				// log.Printf("Validator recieved error: %v", err)
				// fmt.Println("Resetting channels")
				con, err := ethbridge.New(val.serverAddress, val.arbAddresses)
				if err != nil {
					panic(err)
				}
				nonce, err := con.PendingNonceAt(val.auth.From)
				if err != nil {
					panic(err)
				}
				val.auth.Nonce = big.NewInt(int64(nonce))
				val.con = con
				outChan, errChan, err = val.con.CreateListeners(val.VMID)
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	return nil
}

func buildBisectionTree(preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub) *MerkleTree {
	bisectionHashes := make([][32]byte, 0, len(assertions))
	for i := range assertions {
		bisectionBytes := solsha3.SoliditySHA3(
			solsha3.Bytes32(preconditions[i].Hash()),
			solsha3.Bytes32(assertions[i].Hash()),
		)
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], bisectionBytes)
		bisectionHashes = append(bisectionHashes, bisectionHash)
	}
	return NewMerkleTree(bisectionHashes)
}

func (val *EthValidator) FinalizedAssertion(
	assertion *protocol.Assertion,
	newLogCount int,
	signatures [][]byte,
	proposalResults *valmessage.UnanimousUpdateResults,
	onChainTxHash []byte,
) {
	val.CompletedCallChan <- valmessage.FinalizedAssertion{
		Assertion:       assertion,
		NewLogCount:     newLogCount,
		Signatures:      signatures,
		ProposalResults: proposalResults,
		OnChainTxHash:   onChainTxHash,
	}
}

func (val *EthValidator) FinalizedUnanimousAssert(newInboxHash [32]byte, timeBounds protocol.TimeBounds, assertion *protocol.Assertion, signatures [][]byte) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.FinalizedUnanimousAssert(
			val.auth,
			val.VMID,
			newInboxHash,
			timeBounds,
			assertion,
			signatures,
		)
		if err != nil {
			return errors2.Wrap(err, "failed sending finalized unanimous assertion")
		}
		return nil
	}
}

func (val *EthValidator) PendingUnanimousAssert(newInboxHash [32]byte, timeBounds protocol.TimeBounds, assertion *protocol.Assertion, sequenceNum uint64, signatures [][]byte) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.PendingUnanimousAssert(
			val.auth,
			val.VMID,
			newInboxHash,
			timeBounds,
			assertion,
			sequenceNum,
			signatures,
		)
		if err != nil {
			return errors2.Wrap(err, "failed proposing unanimous assertion")
		}
		return nil
	}
}

func (val *EthValidator) ConfirmUnanimousAsserted(newInboxHash [32]byte, assertion *protocol.Assertion) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.ConfirmUnanimousAsserted(
			val.auth,
			val.VMID,
			newInboxHash,
			assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed confirming unanimous assertion")
		}
		return nil
	}
}

func (val *EthValidator) PendingDisputableAssert(precondition *protocol.Precondition, assertion *protocol.Assertion) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.PendingDisputableAssert(
			val.auth,
			val.VMID,
			precondition,
			assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed initiating disputable assertion")
		}
		return nil
	}
}

func (val *EthValidator) ConfirmDisputableAsserted(
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.ConfirmDisputableAsserted(
			val.auth,
			val.VMID,
			precondition,
			assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed confirming assertion")
		}
		return nil
	}
}

func (val *EthValidator) InitiateChallenge(precondition *protocol.Precondition, assertion *protocol.AssertionStub) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.InitiateChallenge(
			val.auth,
			val.VMID,
			precondition,
			assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed initiating challenge")
		}
		return nil
	}
}

func (val *EthValidator) BisectAssertion(precondition *protocol.Precondition, assertions []*protocol.AssertionStub, deadline uint64) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.BisectAssertion(
			val.auth,
			val.VMID,
			deadline,
			precondition,
			assertions,
		)
		if err != nil {
			return errors2.Wrap(err, "failed initiating bisection")
		}
		return nil
	}
}

func (val *EthValidator) ContinueChallenge(assertionToChallenge uint16, preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub, deadline uint64) {
	val.actionChan <- func(val *EthValidator) error {
		tree := buildBisectionTree(preconditions, assertions)
		root := tree.GetRoot()
		_, err := val.con.ContinueChallenge(
			val.auth,
			val.VMID,
			big.NewInt(int64(assertionToChallenge)),
			tree.GetProofFlat(int(assertionToChallenge)),
			root,
			tree.GetNode(int(assertionToChallenge)),
			deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed continuing challenge")
		}
		return nil
	}
}

func (val *EthValidator) OneStepProof(precondition *protocol.Precondition, assertion *protocol.AssertionStub, proof []byte, deadline uint64) {
	val.actionChan <- func(val *EthValidator) error {
		_, err := val.con.OneStepProof(
			val.auth,
			val.VMID,
			precondition,
			assertion,
			proof,
			deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed one step proof")
		}
		return nil
	}
}

func (val *EthValidator) AsserterTimedOut(precondition *protocol.Precondition, assertion *protocol.AssertionStub, deadline uint64) {
	val.actionChan <- func(val *EthValidator) error {
		preAssBytes := solsha3.SoliditySHA3(
			solsha3.Bytes32(precondition.Hash()),
			solsha3.Bytes32(assertion.Hash()),
		)
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], preAssBytes)
		_, err := val.con.Challenge.AsserterTimedOut(
			val.auth,
			val.VMID,
			bisectionHash,
			deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed timing out challenge")
		}
		return nil
	}
}

func (val *EthValidator) ChallengerTimedOut(preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub, deadline uint64) {
	val.actionChan <- func(val *EthValidator) error {
		tree := buildBisectionTree(preconditions, assertions)
		_, err := val.con.Challenge.ChallengerTimedOut(
			val.auth,
			val.VMID,
			tree.GetRoot(),
			deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed timing out challenge")
		}
		return nil
	}
}

func (val *EthValidator) AdvanceBlockchain(blockCount int) error {
	return val.con.AdvanceBlockchain(val.auth, blockCount)
}

func (val *EthValidator) DepositFunds(amount *big.Int) (*types.Transaction, error) {
	senderArr := [32]byte{}
	copy(senderArr[:], val.Address().Bytes())
	tx, err := val.con.DepositFunds(val.auth, amount, senderArr)
	val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (val *EthValidator) CreateVM(createData *valmessage.CreateVMValidatorRequest, signatures [][]byte) (*types.Transaction, error) {
	tx, err := val.con.CreateVM(
		val.auth,
		createData,
		hashing.CreateVMHash(createData),
		signatures,
	)
	val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (val *EthValidator) SendMessage(data value.Value, tokenType [21]byte, currency *big.Int) (*types.Transaction, error) {
	tx, err := val.con.SendMessage(val.auth, protocol.NewMessage(data, tokenType, currency, val.VMID))
	val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (val *EthValidator) SendEthMessage(
	data value.Value,
	amount *big.Int,
) (*types.Transaction, error) {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(data, &dataBuf); err != nil {
		return nil, err
	}
	tx, err := val.con.SendEthMessage(val.auth, data, val.VMID, amount)
	val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (val *EthValidator) UnanimousAssertHash(
	sequenceNum uint64,
	beforeHash [32]byte,
	timeBounds protocol.TimeBounds,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.Assertion,
) ([32]byte, error) {
	return hashing.UnanimousAssertHash(
		val.VMID,
		sequenceNum,
		beforeHash,
		timeBounds,
		newInboxHash,
		originalInboxHash,
		assertion,
	)
}
