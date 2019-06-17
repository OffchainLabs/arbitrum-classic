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

//var rComp [32]byte
//var sComp [32]byte
//copy(rComp[:], signature[:32])
//copy(sComp[:], signature[32:64])
//return valmessage.Signature{
//R: rComp,
//S: sComp,
//V: uint8(int(signature[64])) + 27, // Yes add 27, weird Ethereum quirk
//}, nil

func (val *EthValidator) Sign(msgHash [32]byte) ([]byte, error) {
	data := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(msgHash))
	return crypto.Sign(data, val.key)
}

func (val *EthValidator) StartListening() error {
	outChan, errChan, err := val.con.CreateListeners(val.VmId)
	if err != nil {
		return err
	}

	incomingChan := make(chan valmessage.OutgoingMessage, 1024)

	val.Bot.Run(outChan, incomingChan)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			select {
			case event := <-incomingChan:
				if event != nil {
					err := val.handleSendRequest(event)
					if err != nil {
						log.Fatalf("Error handling send: %v", err)
					}
				}
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
				outChan, errChan, err = val.con.CreateListeners(val.VmId)
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

func (val *EthValidator) handleSendRequest(msg valmessage.OutgoingMessage) error {
	switch msg := msg.(type) {
	case valmessage.FinalizedAssertion:
		val.CompletedCallChan <- msg
	case valmessage.SendProposeUnanimousAssertMessage:
		_, err := val.con.ProposeUnanimousAssert(
			val.auth,
			val.VmId,
			msg.NewInboxHash,
			msg.TimeBounds,
			msg.Assertion,
			msg.SequenceNum,
			msg.Signatures,
		)
		if err != nil {
			return errors2.Wrap(err, "failed proposing unanimous assertion")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendConfirmUnanimousAssertedMessage:
		_, err := val.con.ConfirmUnanimousAsserted(
			val.auth,
			val.VmId,
			msg.NewInboxHash,
			msg.Assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed confirming unanimous assertion")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendUnanimousAssertMessage:
		_, err := val.con.UnanimousAssert(
			val.auth,
			val.VmId,
			msg.NewInboxHash,
			msg.TimeBounds,
			msg.Assertion,
			msg.Signatures,
		)
		if err != nil {
			return errors2.Wrap(err, "failed sending finalized unanimous assertion")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendAssertMessage:
		_, err := val.con.DisputableAssert(
			val.auth,
			val.VmId,
			msg.Precondition,
			msg.Assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed initiating disputable assertion")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendInitiateChallengeMessage:
		_, err := val.con.InitiateChallenge(
			val.auth,
			val.VmId,
			msg.Precondition,
			msg.Assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed initiating challenge")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendBisectionMessage:
		_, err := val.con.BisectChallenge(
			val.auth,
			val.VmId,
			msg.Deadline,
			msg.Precondition,
			msg.Assertions,
		)
		if err != nil {
			return errors2.Wrap(err, "failed initiating bisection")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendContinueChallengeMessage:
		tree := buildBisectionTree(msg.Preconditions, msg.Assertions)
		root := tree.GetRoot()
		_, err := val.con.ContinueChallenge(
			val.auth,
			val.VmId,
			big.NewInt(int64(msg.AssertionToChallenge)),
			tree.GetProofFlat(int(msg.AssertionToChallenge)),
			root,
			tree.GetNode(int(msg.AssertionToChallenge)),
			msg.Deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed continuing challenge")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendOneStepProofMessage:
		_, err := val.con.OneStepProof(
			val.auth,
			val.VmId,
			msg.Precondition,
			msg.Assertion.Stub(),
			msg.Proof,
			msg.Deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed one step proof")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendConfirmedAssertMessage:
		_, err := val.con.ConfirmAsserted(
			val.auth,
			val.VmId,
			msg.Precondition,
			msg.Assertion,
		)
		if err != nil {
			return errors2.Wrap(err, "failed confirming assertion")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendAsserterTimedOutChallengeMessage:
		preAssBytes := solsha3.SoliditySHA3(
			solsha3.Bytes32(msg.Precondition.Hash()),
			solsha3.Bytes32(msg.Assertion.Hash()),
		)
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], preAssBytes)
		_, err := val.con.Challenge.AsserterTimedOut(
			val.auth,
			val.VmId,
			bisectionHash,
			msg.Deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed timing out challenge")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	case valmessage.SendChallengerTimedOutChallengeMessage:
		tree := buildBisectionTree(msg.Preconditions, msg.Assertions)
		_, err := val.con.Challenge.ChallengerTimedOut(
			val.auth,
			val.VmId,
			tree.GetRoot(),
			msg.Deadline,
		)
		if err != nil {
			return errors2.Wrap(err, "failed timing out challenge")
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	default:
		return fmt.Errorf("unhandled valmessage %T: %+v", msg, msg)
	}
	return nil
}

func (val *EthValidator) AdvanceBlockchain(blockCount int) error {
	return val.con.AdvanceBlockchain(val.auth, blockCount)
}

func (val *EthValidator) DepositEth(amount *big.Int) (*types.Transaction, error) {
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
		CreateVMHash(createData),
		signatures,
	)
	val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
	return tx, err
}

func (val *EthValidator) SendMessage(data value.Value, tokenType [21]byte, currency *big.Int) (*types.Transaction, error) {
	tx, err := val.con.SendMessage(val.auth, protocol.NewMessage(data, tokenType, currency, val.VmId))
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
	tx, err := val.con.SendEthMessage(val.auth, data, val.VmId, amount)
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
	return UnanimousAssertHash(
		val.VmId,
		sequenceNum,
		beforeHash,
		timeBounds,
		newInboxHash,
		originalInboxHash,
		assertion,
	)
}
