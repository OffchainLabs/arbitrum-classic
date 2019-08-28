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
	"context"
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
	actionChan          chan func(*EthValidator)
	CompletedCallChan   chan valmessage.FinalizedAssertion
	VMCreatedTxHashChan chan [32]byte

	// Not in thread, but internal only
	serverAddress string
	arbAddresses  ethbridge.ArbAddresses

	// private thread only
	con                     *ethbridge.Bridge
	auth                    *bind.TransactOpts
	unprocessedMessageCount uint64
}

func (val *EthValidator) Address() common.Address {
	return crypto.PubkeyToAddress(val.key.PublicKey)
}

func (val *EthValidator) ValidatorCount() int {
	return len(val.Validators)
}

func (val *EthValidator) makeAuth(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     val.auth.From,
		Nonce:    val.auth.Nonce,
		Signer:   val.auth.Signer,
		Value:    val.auth.Value,
		GasPrice: val.auth.GasPrice,
		GasLimit: val.auth.GasLimit,
		Context:  ctx,
	}
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

	actionChan := make(chan func(*EthValidator), 1024)
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
		0,
	}
	return val, nil
}

func (val *EthValidator) Sign(msgHash [32]byte) ([]byte, error) {
	data := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(msgHash))
	return crypto.Sign(data, val.key)
}

func (val *EthValidator) RestartConnection() (chan ethbridge.Notification, chan error, error) {
	con, err := ethbridge.New(val.serverAddress, val.arbAddresses)
	if err != nil {
		return nil, nil, err
	}
	nonce, err := con.PendingNonceAt(val.auth.From)
	if err != nil {
		return nil, nil, err
	}
	val.auth.Nonce = big.NewInt(int64(nonce))
	val.con = con
	return val.con.CreateListeners(val.VMID)
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
			case parse, ok := <-outChan:
				if !ok {
					outChan, errChan, err = val.RestartConnection()
					if err != nil {
						panic(err)
					}
					break
				}
				if _, ok := parse.Event.(ethbridge.VMCreatedEvent); ok {
					val.VMCreatedTxHashChan <- parse.TxHash
				}
				parsedChan <- parse
			case event := <-val.actionChan:
				event(val)
			case <-errChan:
				// Ignore error and try to reset connection
				// log.Printf("Validator recieved error: %v", err)
				// fmt.Println("Resetting channels")
				for {
					outChan, errChan, err = val.RestartConnection()
					if err == nil {
						break
					}
					log.Println("Error: Validator can't connect to blockchain")
					time.Sleep(5 * time.Second)
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

func (val *EthValidator) AddedNewMessages(count uint64) {
	val.actionChan <- func(val *EthValidator) {
		val.unprocessedMessageCount += count
	}
}

func (val *EthValidator) FinalizedAssertion(
	assertion *protocol.Assertion,
	onChainTxHash []byte,
	signatures [][]byte,
	proposalResults *valmessage.UnanimousUpdateResults,
) {
	val.actionChan <- func(val *EthValidator) {
		finalizedAssertion := valmessage.FinalizedAssertion{
			Assertion:       assertion,
			OnChainTxHash:   onChainTxHash,
			Signatures:      signatures,
			ProposalResults: proposalResults,
		}
		val.unprocessedMessageCount -= uint64(len(finalizedAssertion.NewLogs()))
		val.CompletedCallChan <- finalizedAssertion
	}
}

func (val *EthValidator) FinalizedUnanimousAssert(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	signatures [][]byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.FinalizedUnanimousAssert(
			val.makeAuth(ctx),
			val.VMID,
			newInboxHash,
			assertion,
			signatures,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed sending finalized unanimous assertion")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed sending finalized unanimous assertion")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) PendingUnanimousAssert(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	sequenceNum uint64,
	signatures [][]byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.PendingUnanimousAssert(
			val.makeAuth(ctx),
			val.VMID,
			newInboxHash,
			assertion,
			sequenceNum,
			signatures,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed proposing unanimous assertion")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed proposing unanimous assertion")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) ConfirmUnanimousAsserted(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.ConfirmUnanimousAsserted(
			val.makeAuth(ctx),
			val.VMID,
			newInboxHash,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed confirming unanimous assertion")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed confirming unanimous assertion")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) PendingDisputableAssert(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.PendingDisputableAssert(
			val.makeAuth(ctx),
			val.VMID,
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating disputable assertion")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating disputable assertion")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) ConfirmDisputableAsserted(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.ConfirmDisputableAsserted(
			val.makeAuth(ctx),
			val.VMID,
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed confirming disputable assertion")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed confirming disputable assertion")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) InitiateChallenge(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.InitiateChallenge(
			val.makeAuth(ctx),
			val.VMID,
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating challenge")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating challenge")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) BisectAssertion(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertions []*protocol.AssertionStub,
	deadline uint64,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.BisectAssertion(
			val.makeAuth(ctx),
			val.VMID,
			deadline,
			precondition,
			assertions,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating bisection")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating bisection")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) ContinueChallenge(
	ctx context.Context,
	assertionToChallenge uint16,
	preconditions []*protocol.Precondition,
	assertions []*protocol.AssertionStub,
	deadline uint64,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tree := buildBisectionTree(preconditions, assertions)
		root := tree.GetRoot()
		tx, err := val.con.ContinueChallenge(
			val.makeAuth(ctx),
			val.VMID,
			big.NewInt(int64(assertionToChallenge)),
			tree.GetProofFlat(int(assertionToChallenge)),
			root,
			tree.GetNode(int(assertionToChallenge)),
			deadline,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed continuing challenge")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed continuing challenge")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) OneStepProof(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	proof []byte,
	deadline uint64,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.OneStepProof(
			val.makeAuth(ctx),
			val.VMID,
			precondition,
			assertion,
			proof,
			deadline,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed one step proof")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed one step proof")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) AsserterTimedOut(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	deadline uint64,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		preAssBytes := solsha3.SoliditySHA3(
			solsha3.Bytes32(precondition.Hash()),
			solsha3.Bytes32(assertion.Hash()),
		)
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], preAssBytes)
		tx, err := val.con.Challenge.AsserterTimedOut(
			val.makeAuth(ctx),
			val.VMID,
			bisectionHash,
			deadline,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) ChallengerTimedOut(
	ctx context.Context,
	preconditions []*protocol.Precondition,
	assertions []*protocol.AssertionStub,
	deadline uint64,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tree := buildBisectionTree(preconditions, assertions)
		tx, err := val.con.Challenge.ChallengerTimedOut(
			val.makeAuth(ctx),
			val.VMID,
			tree.GetRoot(),
			deadline,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) AdvanceBlockchain(
	ctx context.Context,
	blockCount int,
) error {
	return val.con.AdvanceBlockchain(val.makeAuth(ctx), blockCount)
}

func (val *EthValidator) DepositFunds(
	ctx context.Context,
	amount *big.Int,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		senderArr := [32]byte{}
		copy(senderArr[:], val.Address().Bytes())
		tx, err := val.con.DepositFunds(val.makeAuth(ctx), amount, senderArr)
		if err != nil {
			errChan <- err
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- err
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) GetTokenBalance(
	user [32]byte,
	tokenContract common.Address,
) (*big.Int, error) {
	auth := &bind.CallOpts{
		Pending: false,
		From:    val.auth.From,
		Context: val.auth.Context,
	}
	amt, err := val.con.GetTokenBalance(auth, user, tokenContract)
	return amt, err
}

func (val *EthValidator) WaitForTokenBalance(
	ctx context.Context,
	user [32]byte,
	tokenContract common.Address,
	amount *big.Int,
) error {
	auth := &bind.CallOpts{
		Pending: false,
		From:    val.auth.From,
		Context: ctx,
	}

	for {
		select {
		case _ = <-time.After(time.Second):
			amt, err := val.con.GetTokenBalance(auth, user, tokenContract)
			if err != nil {
				return err
			}
			if amount.Cmp(amt) >= 0 {
				return nil
			}
		case _ = <-ctx.Done():
			return errors.New("balance not reached")
		}
	}
}

func (val *EthValidator) CreateVM(
	ctx context.Context,
	createData *valmessage.CreateVMValidatorRequest,
	signatures [][]byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.CreateVM(
			val.makeAuth(ctx),
			createData,
			hashing.CreateVMHash(createData),
			signatures,
		)
		if err != nil {
			errChan <- err
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- err
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) SendMessage(
	ctx context.Context,
	data value.Value,
	tokenType [21]byte,
	currency *big.Int,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.SendMessage(val.makeAuth(ctx), protocol.NewMessage(data, tokenType, currency, val.VMID))
		if err != nil {
			errChan <- err
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- err
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) ForwardMessage(
	ctx context.Context,
	data value.Value,
	tokenType [21]byte,
	currency *big.Int,
	sig []byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		tx, err := val.con.ForwardMessage(val.makeAuth(ctx), protocol.NewMessage(data, tokenType, currency, val.VMID), sig)
		if err != nil {
			errChan <- err
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- err
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) SendEthMessage(
	ctx context.Context,
	data value.Value,
	amount *big.Int,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	val.actionChan <- func(val *EthValidator) {
		var dataBuf bytes.Buffer
		if err := value.MarshalValue(data, &dataBuf); err != nil {
			errChan <- err
			return
		}
		tx, err := val.con.SendEthMessage(val.makeAuth(ctx), data, val.VMID, amount)
		if err != nil {
			errChan <- err
			return
		}
		val.auth.Nonce.Add(val.auth.Nonce, big.NewInt(1))
		receipt, err := val.con.WaitForReceipt(ctx, tx.Hash())
		if err != nil {
			errChan <- err
			return
		}
		receiptChan <- receipt
	}
	return receiptChan, errChan
}

func (val *EthValidator) UnanimousAssertHash(
	sequenceNum uint64,
	beforeHash [32]byte,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.Assertion,
) ([32]byte, error) {
	return hashing.UnanimousAssertHash(
		val.VMID,
		sequenceNum,
		beforeHash,
		newInboxHash,
		originalInboxHash,
		assertion,
	)
}
