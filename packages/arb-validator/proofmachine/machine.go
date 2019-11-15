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
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	machine machine.Machine
	ethConn *Connection
}

type Connection struct {
	fromAddress common.Address
	osp         *ethbridge.OneStepProof
	client      *ethclient.Client
	proofbounds [2]uint64
}

func NewEthConnection(contractAddress common.Address, key *ecdsa.PrivateKey, ethURL string, proofbounds [2]uint64) (*Connection, error) {
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		log.Fatal("Connection failure ", err)
	}
	osp, err := ethbridge.NewOneStepProof(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	return &Connection{
		fromAddress: keyAddr,
		osp:         osp,
		client:      client,
		proofbounds: proofbounds,
	}, err
}

func New(mach machine.Machine, ethConn *Connection) (*Machine, error) {
	return &Machine{
		machine: mach,
		ethConn: ethConn,
	}, nil
}

func (m *Machine) Hash() [32]byte {
	return m.machine.Hash()
}

func (m *Machine) PrintState() {
	m.machine.PrintState()
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{m.machine.Clone(), m.ethConn}
}

func (m *Machine) CurrentStatus() machine.Status {
	return m.machine.CurrentStatus()
}

func (m *Machine) LastBlockReason() machine.BlockReason {
	return m.machine.LastBlockReason()
}

func (m *Machine) CanSpend(tokenType protocol.TokenType, currency *big.Int) bool {
	return m.machine.CanSpend(tokenType, currency)
}

func (m *Machine) InboxHash() value.HashOnlyValue {
	return m.machine.InboxHash()
}

func (m *Machine) PendingMessageCount() uint64 {
	return m.machine.PendingMessageCount()
}

func (m *Machine) SendOnchainMessage(msg protocol.Message) {
	m.machine.SendOnchainMessage(msg)
}

func (m *Machine) DeliverOnchainMessage() {
	m.machine.DeliverOnchainMessage()
}

func (m *Machine) SendOffchainMessages(msgs []protocol.Message) {
	m.machine.SendOffchainMessages(msgs)
}

func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	a := &protocol.Assertion{}
	stepIncrease := int32(1)
	stepsRan := 0
	for i := int32(0); i < maxSteps; i += stepIncrease {
		var proof []byte
		var err error
		// only marshall if we are going to validate (see below)
		if i >= int32(m.ethConn.proofbounds[0]) && i <= int32(m.ethConn.proofbounds[1]) {
			proof, err = m.MarshalForProof()
			if err != nil {
				log.Println("error marshaling")
			}
		}
		steps := int32(stepIncrease)
		beforeHash := m.Hash()
		inboxHash := m.InboxHash()
		a1 := m.machine.ExecuteAssertion(steps, timeBounds)
		a.AfterHash = a1.AfterHash
		a.NumSteps += a1.NumSteps
		a.Logs = append(a.Logs, a1.Logs...)
		a.OutMsgs = append(a.OutMsgs, a1.OutMsgs...)

		if a1.NumSteps == 0 {
			fmt.Println(" machine halted ")
			break
		}
		if a1.NumSteps != 1 {
			log.Println("Num steps = ", a1.NumSteps)
		}
		stepsRan++

		// only marshall and validate if step is within proofbounds
		if i >= int32(m.ethConn.proofbounds[0]) && i <= int32(m.ethConn.proofbounds[1]) {
			spentBalance := protocol.NewTokenTrackerFromMessages(a1.OutMsgs)
			callOpts := &bind.CallOpts{
				Pending: true,
				From:    m.ethConn.fromAddress,
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

			res, err := m.ethConn.osp.ValidateProof(callOpts, precond, a1.Stub(), proof)
			if err != nil {
				log.Println("Machine ended with error:")
				m.PrintState()
				log.Fatal("Proof invalid ", err)
			}
			if res.Cmp(big.NewInt(0)) == 0 {
				log.Println("Proof valid")
			} else {
				log.Println("Machine ended with invalid proof:")
				m.PrintState()
				log.Fatalln("Proof invalid")
			}
		}
	}
	fmt.Println("Proof mode ran ", stepsRan, " steps")
	return a
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	return m.machine.MarshalForProof()
}
