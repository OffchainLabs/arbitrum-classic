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
	"fmt"

	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-validator/valmessage"
	"github.com/offchainlabs/arb-validator/ethbridge"
	"github.com/offchainlabs/arb-validator/validator"
)

func createAddressMerkleTree(addresses []common.Address) *MerkleTree {
	converted := make([][32]byte, 0, len(addresses))
	for _, a := range addresses {
		data := [32]byte{}
		copy(data[:], a.Bytes()[:])
		converted = append(converted, data)
	}
	return NewMerkleTree(converted)
}

type EthValidator struct {
	key *ecdsa.PrivateKey

	// Safe public interface
	VmId              [32]byte
	Validators        map[common.Address]validatorInfo
	Bot               *validator.Validator
	CompletedCallChan chan valmessage.FinalizedAssertion

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
	proof    []byte
	indexNum uint16
}

type VMResponse struct {
	Message protocol.Message
	Result  value.Value
	Proof   [][32]byte
}

func NewEthValidator(
	name string,
	vmId [32]byte,
	machine *vm.Machine,
	key *ecdsa.PrivateKey,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
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
	manTree := createAddressMerkleTree(keys)
	for i, add := range keys {
		manMap[add] = validatorInfo{manTree.GetProofFlat(i), uint16(i)}
	}

	_, found := manMap[crypto.PubkeyToAddress(key.PublicKey)]
	if !found {
		return nil, errors.New("key is not a validator of chosen VM")
	}

	bot := validator.NewValidator(name, auth.From, protocol.NewEmptyInbox(), protocol.NewBalanceTracker(), config, machine, challengeEverything)

	completedCallChan := make(chan valmessage.FinalizedAssertion, 1024)

	return &EthValidator{key, vmId, manMap, bot, completedCallChan, ethURL, connectionInfo, con, auth}, nil
}

func (man *EthValidator) Sign(msgHash [32]byte) (valmessage.Signature, error) {
	data := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(msgHash))
	signature, err := crypto.Sign(data, man.key)
	if err != nil {
		panic(err)
	}
	var rComp [32]byte
	var sComp [32]byte
	copy(rComp[:], signature[:32])
	copy(sComp[:], signature[32:64])
	return valmessage.Signature{
		R: rComp,
		S: sComp,
		V: uint8(int(signature[64])) + 27, // Yes add 27, weird Ethereum quirk
	}, nil
}

func (man *EthValidator) StartListening() error {
	outChan, errChan, err := man.con.CreateListeners(man.VmId)
	if err != nil {
		return err
	}

	incomingChan := make(chan valmessage.OutgoingMessage, 1024)
	outgoingChan := make(chan valmessage.IncomingValidatorMessage, 1024)

	man.Bot.Run(outgoingChan, incomingChan)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			select {
			case val := <-outChan:
				err := man.handleEvent(val, outgoingChan)
				if err != nil {
					log.Fatalf("Error handling event: %v", err)
				}
			case event := <-incomingChan:
				if event != nil {
					err := man.handleSendRequest(event)
					if err != nil {
						log.Fatalf("Error handling send: %v", err)
					}
				}
			case <-errChan:
				// Ignore error and try to reset connection
				// log.Printf("Validator recieved error: %v", err)
				// fmt.Println("Resetting channels")
				con, err := ethbridge.New(man.serverAddress, man.arbAddresses)
				if err != nil {
					panic(err)
				}
				nonce, err := con.PendingNonceAt(man.auth.From)
				if err != nil {
					panic(err)
				}
				man.auth.Nonce = big.NewInt(int64(nonce))
				man.con = con
				outChan, errChan, err = man.con.CreateListeners(man.VmId)
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	return nil
}

func (man *EthValidator) handleEvent(note ethbridge.Notification, outgoingChan chan valmessage.IncomingValidatorMessage) error {
	switch ev := note.Event.(type) {
	case ethbridge.VMCreatedEvent:
		// fmt.Printf("Created vm with state %x\n", Val.VmState)
	case ethbridge.MessageDeliveredEvent:
		fmt.Println("VM recieved on-chain message")
		outgoingChan <- valmessage.IncomingMessageMessage{Msg: ev.Msg, Header: note.Header}
	case ethbridge.FinalUnanimousAssertEvent:
		msg := valmessage.FinalUnanimousAssertMessage{
			UnanHash: ev.UnanHash,
		}
		outgoingChan <- valmessage.BridgeMessage{Message: msg, Header: note.Header}
	case ethbridge.ProposedUnanimousAssertEvent:
		msg := valmessage.ProposedUnanimousAssertMessage{
			UnanHash:    ev.UnanHash,
			SequenceNum: ev.SequenceNum,
		}
		outgoingChan <- valmessage.BridgeMessage{Message: msg, Header: note.Header}
	case ethbridge.ConfirmedUnanimousAssertEvent:
		msg := valmessage.ConfirmedUnanimousAssertMessage{
			SequenceNum: ev.SequenceNum,
		}
		outgoingChan <- valmessage.BridgeMessage{Message: msg, Header: note.Header}
	case ethbridge.DisputableAssertionEvent:
		assertMessage := valmessage.AssertMessage{Precondition: ev.Precondition, Assertion: ev.Assertion, Asserter: ev.Asserter}
		outgoingChan <- valmessage.BridgeMessage{Message: assertMessage, Header: note.Header}
	case ethbridge.ConfirmedAssertEvent:
		outgoingChan <- valmessage.BridgeMessage{Message: valmessage.ConfirmedAssertMessage{}, Header: note.Header}
	case ethbridge.InitiateChallengeEvent:
		outgoingChan <- valmessage.BridgeMessage{Message: valmessage.InitiateChallengeMessage{Challenger: ev.Challenger}, Header: note.Header}
	case ethbridge.BisectionEvent:
		outgoingChan <- valmessage.BridgeMessage{Message: valmessage.BisectMessage{Assertions: ev.Assertions}, Header: note.Header}
	case ethbridge.ContinueChallengeEvent:
		challengeMessage := valmessage.ContinueChallengeMessage{ChallengedAssertion: uint16(ev.ChallengedAssertion)}
		outgoingChan <- valmessage.BridgeMessage{Message: challengeMessage, Header: note.Header}
	case ethbridge.AsserterTimeoutEvent:
		msg := valmessage.AsserterTimeoutMessage{}
		outgoingChan <- valmessage.BridgeMessage{Message: msg, Header: note.Header}
	case ethbridge.ChallengerTimeoutEvent:
		msg := valmessage.ChallengerTimeoutMessage{}
		outgoingChan <- valmessage.BridgeMessage{Message: msg, Header: note.Header}
	case ethbridge.OneStepProofEvent:
		msg := valmessage.OneStepProofMessage{}
		outgoingChan <- valmessage.BridgeMessage{Message: msg, Header: note.Header}
	case *types.Header:
		outgoingChan <- valmessage.TimeUpdateMessage{Header: ev}
	default:
		fmt.Println("Unknown event: ", ev)
	}
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

func LogProof(a *protocol.Assertion, index int) ([][32]byte, error) {
	if index < len(a.Logs) {
		return nil, errors.New("log index out of range")
	}
	proof := make([][32]byte, 0, len(a.Logs)-1-index)
	for i := len(a.Logs) - 1; i > index; i-- {
		proof = append(proof, a.Logs[i].Hash())
	}
	return proof, nil
}

func (man *EthValidator) handleSendRequest(msg valmessage.OutgoingMessage) error {
	switch msg := msg.(type) {
	case valmessage.FinalizedAssertion:
		man.CompletedCallChan <- msg
	case valmessage.SendProposeUnanimousAssertMessage:
		_, err := man.ProposeUnanimousAssert(msg.NewInboxHash, msg.TimeBounds, msg.Assertion, msg.SequenceNum, msg.Signatures)
		if err != nil {
			return errors2.Wrap(err, "failed proposing unanimous assertion")
		}
	case valmessage.SendConfirmUnanimousAssertedMessage:
		_, err := man.ConfirmUnanimousAsserted(msg.NewInboxHash, msg.Assertion)
		if err != nil {
			return errors2.Wrap(err, "failed confirming unanimous assertion")
		}
	case valmessage.SendUnanimousAssertMessage:
		_, err := man.UnanimousAssert(msg.NewInboxHash, msg.TimeBounds, msg.Assertion, msg.Signatures)
		if err != nil {
			return errors2.Wrap(err, "failed sending finalized unanimous assertion")
		}
	case valmessage.SendAssertMessage:
		_, err := man.DisputableAssert(msg.Precondition, msg.Assertion)
		if err != nil {
			return errors2.Wrap(err, "failed initiating disputable assertion")
		}
	case valmessage.SendInitiateChallengeMessage:
		_, err := man.InitiateChallenge(msg.Precondition, msg.Assertion)
		if err != nil {
			return errors2.Wrap(err, "failed initiating challenge")
		}
	case valmessage.SendBisectionMessage:
		_, err := man.BisectChallenge(msg.Deadline, msg.Precondition, msg.Assertions)
		if err != nil {
			return errors2.Wrap(err, "failed initiating bisection")
		}
	case valmessage.SendContinueChallengeMessage:
		tree := buildBisectionTree(msg.Preconditions, msg.Assertions)
		root := tree.GetRoot()
		_, err := man.ContinueChallenge(
			big.NewInt(int64(msg.AssertionToChallenge)),
			tree.GetProofFlat(int(msg.AssertionToChallenge)),
			root,
			tree.GetNode(int(msg.AssertionToChallenge)),
			msg.Deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed continuing challenge")
		}
	case valmessage.SendOneStepProofMessage:
		_, err := man.OneStepProof(
			msg.Precondition,
			msg.Assertion.Stub(),
			msg.Proof,
			msg.Deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed one step proof")
		}
	case valmessage.SendConfirmedAssertMessage:
		_, err := man.ConfirmAsserted(msg.Precondition, msg.Assertion)
		if err != nil {
			return errors2.Wrap(err, "failed confirming assertion")
		}
	case valmessage.SendAsserterTimedOutChallengeMessage:
		preAssBytes := solsha3.SoliditySHA3(
			solsha3.Bytes32(msg.Precondition.Hash()),
			solsha3.Bytes32(msg.Assertion.Hash()),
		)
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], preAssBytes)
		_, err := man.AsserterTimedOutChallenge(
			bisectionHash,
			msg.Deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed timing out challenge")
		}
	case valmessage.SendChallengerTimedOutChallengeMessage:
		tree := buildBisectionTree(msg.Preconditions, msg.Assertions)
		_, err := man.ChallengerTimedOut(tree.GetRoot(), msg.Deadline)
		if err != nil {
			return errors2.Wrap(err, "failed timing out challenge")
		}
	default:
		return fmt.Errorf("unhandled valmessage %T: %+v", msg, msg)
	}
	return nil
}

func (man *EthValidator) AdvanceBlockchain(blockCount int) error {
	return man.con.AdvanceBlockchain(man.auth, blockCount)
}

func (man *EthValidator) DepositEth(amount *big.Int) (*types.Transaction, error) {
	senderArr := [32]byte{}
	copy(senderArr[:], man.Address().Bytes())
	tx, err := man.con.DepositFunds(man.auth, amount, senderArr)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) CreateVM(createData *valmessage.CreateVMValidatorRequest, signatures []valmessage.Signature) (*types.Transaction, error) {
	tx, err := man.con.CreateVM(
		man.auth,
		createData,
		CreateVMHash(createData),
		signatures,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) SendMessage(data value.Value, tokenType [21]byte, currency *big.Int) (*types.Transaction, error) {
	tx, err := man.con.SendMessage(man.auth, protocol.NewMessage(data, tokenType, currency, man.VmId))
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) SendEthMessage(
	data value.Value,
	amount *big.Int,
) (*types.Transaction, error) {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(data, &dataBuf); err != nil {
		return nil, err
	}
	tx, err := man.con.SendEthMessage(man.auth, data, man.VmId, amount)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) UnanimousAssertHash(
	sequenceNum uint64,
	beforeHash [32]byte,
	timeBounds protocol.TimeBounds,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.Assertion,
) ([32]byte, error) {
	return UnanimousAssertHash(
		man.VmId,
		sequenceNum,
		beforeHash,
		timeBounds,
		newInboxHash,
		originalInboxHash,
		assertion,
	)
}

func (man *EthValidator) ProposeUnanimousAssert(
	newInboxHash [32]byte,
	timeBounds protocol.TimeBounds,
	assertion *protocol.Assertion,
	sequenceNum uint64,
	signatures []valmessage.Signature,
) (*types.Transaction, error) {
	tx, err := man.con.ProposeUnanimousAssert(
		man.auth,
		man.VmId,
		newInboxHash,
		timeBounds,
		assertion,
		sequenceNum,
		signatures,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) ConfirmUnanimousAsserted(
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
) (*types.Transaction, error) {
	tx, err := man.con.ConfirmUnanimousAsserted(
		man.auth,
		man.VmId,
		newInboxHash,
		assertion,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) UnanimousAssert(
	newInboxHash [32]byte,
	timeBounds protocol.TimeBounds,
	assertion *protocol.Assertion,
	signatures []valmessage.Signature,
) (*types.Transaction, error) {
	tx, err := man.con.UnanimousAssert(
		man.auth,
		man.VmId,
		newInboxHash,
		timeBounds,
		assertion,
		signatures,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) DisputableAssert(
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Transaction, error) {
	tx, err := man.con.DisputableAssert(
		man.auth,
		man.VmId,
		precondition,
		assertion,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) ConfirmAsserted(
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Transaction, error) {
	tx, err := man.con.ConfirmAsserted(
		man.auth,
		man.VmId,
		precondition,
		assertion,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) InitiateChallenge(
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
) (*types.Transaction, error) {
	tx, err := man.con.InitiateChallenge(
		man.auth,
		man.VmId,
		precondition,
		assertion,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) BisectChallenge(
	deadline uint64,
	precondition *protocol.Precondition,
	assertions []*protocol.Assertion,
) (*types.Transaction, error) {
	tx, err := man.con.BisectChallenge(
		man.auth,
		man.VmId,
		deadline,
		precondition,
		assertions,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) ContinueChallenge(
	assertionToChallenge *big.Int,
	bisectionProof []byte,
	bisectionRoot [32]byte,
	bisectionHash [32]byte,
	deadline uint64,
) (*types.Transaction, error) {
	tx, err := man.con.ContinueChallenge(
		man.auth,
		man.VmId,
		assertionToChallenge,
		bisectionProof,
		bisectionRoot,
		bisectionHash,
		deadline,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) OneStepProof(
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	proof []byte,
	deadline uint64,
) (*types.Transaction, error) {
	tx, err := man.con.OneStepProof(
		man.auth,
		man.VmId,
		precondition,
		assertion,
		proof,
		deadline,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) AsserterTimedOutChallenge(
	bisectionHash [32]byte,
	deadline uint64,
) (*types.Transaction, error) {
	tx, err := man.con.Challenge.AsserterTimedOut(
		man.auth,
		man.VmId,
		bisectionHash,
		deadline,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (man *EthValidator) ChallengerTimedOut(
	bisectionRoot [32]byte,
	deadline uint64,
) (*types.Transaction, error) {
	tx, err := man.con.Challenge.ChallengerTimedOut(
		man.auth,
		man.VmId,
		bisectionRoot,
		deadline,
	)
	man.auth.Nonce.Add(man.auth.Nonce, big.NewInt(1))
	return tx, err
}
