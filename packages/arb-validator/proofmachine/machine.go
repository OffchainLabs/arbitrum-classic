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

package proofmachine

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/testmachine"
	"log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	testmachine *testmachine.Machine
	//cppmachine *cmachine.Machine
	//gomachine  *gomachine.Machine
	//proofMode  bool
	data *ProofMachData
}

type ProofMachData struct {
	fromAddress common.Address
	osp         *ethbridge.OneStepProof
	balance     *protocol.BalanceTracker
	client      *ethclient.Client
}

func New(codeFile string, warnMode bool) (*Machine, error) {
	tm, err := testmachine.New(codeFile, warnMode)
	if err != nil {
		err = fmt.Errorf("Test machine error: %v ", err)
	}
	return &Machine{
		tm,
		//true,
		nil,
	}, err
}

func (m *Machine) Hash() [32]byte {
	return m.testmachine.Hash()
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{m.testmachine.Clone().(*testmachine.Machine), m.data}
}

func (m *Machine) CurrentStatus() machine.Status {
	return m.testmachine.CurrentStatus()
}

func (m *Machine) LastBlockReason() machine.BlockReason {
	return m.testmachine.LastBlockReason()
}

func (m *Machine) CanSpend(tokenType protocol.TokenType, currency *big.Int) bool {
	return m.testmachine.CanSpend(tokenType, currency)
}

func (m *Machine) InboxHash() value.HashOnlyValue {
	return m.testmachine.InboxHash()
}

func (m *Machine) PendingMessageCount() uint64 {
	return m.testmachine.PendingMessageCount()
}

func (m *Machine) SendOnchainMessage(msg protocol.Message) {
	m.testmachine.SendOnchainMessage(msg)
}

func (m *Machine) DeliverOnchainMessage() {
	m.testmachine.DeliverOnchainMessage()
}

func (m *Machine) SendOffchainMessages(msgs []protocol.Message) {
	m.testmachine.SendOffchainMessages(msgs)
}

func (m *Machine) ProofMachineData(contractAddress common.Address, key *ecdsa.PrivateKey, ethURL string, balance *protocol.BalanceTracker) {
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		log.Fatal("Connection failure ", err)
	}
	osp, err := ethbridge.NewOneStepProof(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	pd := ProofMachData{
		fromAddress: keyAddr,
		osp:         osp,
		client:      client,
		balance:     balance,
	}
	m.data = &pd
}

func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	if m.data == nil {
		log.Println("Proof data not set")
		return m.testmachine.ExecuteAssertion(maxSteps, timeBounds)
	}

	a := &protocol.Assertion{}
	stepIncrease := int32(1)
	stepsRan := 0
	for i := int32(0); i < maxSteps; i += stepIncrease {
		proof, err := m.MarshalForProof()
		steps := int32(stepIncrease)
		beforeHash := m.Hash()
		inboxHash := m.InboxHash()

		a1 := m.testmachine.ExecuteAssertion(steps, timeBounds)
		a.AfterHash = a1.AfterHash
		a.NumSteps += a1.NumSteps
		a.Logs = append(a.Logs, a1.Logs...)
		a.OutMsgs = append(a.OutMsgs, a1.OutMsgs...)
		if m.testmachine.CurrentStatus() != machine.Extensive {
			fmt.Println(" machine status = ", m.testmachine.CurrentStatus())
			break
		}
		if a1.NumSteps == 0 {
			fmt.Println(" machine halted ")
			break
		}
		if a1.NumSteps != 1 {
			log.Println("Num steps = ", a1.NumSteps)
		}
		stepsRan++

		spentBalance := protocol.NewTokenTrackerFromMessages(a1.OutMsgs)
		balance := m.data.balance.Clone()
		_ = balance.SpendAllTokens(spentBalance)
		callOpts := &bind.CallOpts{
			Pending: true,
			From:    m.data.fromAddress,
			Context: context.Background(),
		}
		// uncomment to force proof fail
		//beforeHash[0] = 5
		precond := &protocol.Precondition{
			BeforeHash:    beforeHash,
			TimeBounds:    timeBounds,
			BeforeBalance: spentBalance,
			BeforeInbox:   inboxHash,
		}

		res, err := m.data.osp.ValidateProof(callOpts, precond, a1.Stub(), proof)
		if err != nil {
			log.Fatal("Proof invalid", err)
		}
		if res.Cmp(big.NewInt(0)) == 0 {
			log.Println("Proof valid")
		} else {
			log.Fatalln("Proof invalid")
		}
		fmt.Println("Proof mode ran ", stepsRan, " steps")
	}
	return a
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	return m.testmachine.MarshalForProof()
}
