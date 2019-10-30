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

package testmachine

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/channel"
	"log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/goloader"
	gomachine "github.com/offchainlabs/arbitrum/packages/arb-avm-go/vm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	cppmachine *cmachine.Machine
	gomachine  *gomachine.Machine
	proofMode  bool
	data       *ProofMachData
}

type ProofMachData struct {
	fromAddress common.Address
	coord       *channel.ValidatorCoordinator
	balance     *protocol.BalanceTracker
}

func New(codeFile string, warnMode bool, proofMode bool) (*Machine, error) {
	gm, gmerr := goloader.LoadMachineFromFile(codeFile, warnMode)
	cm, cmerr := cmachine.New(codeFile)
	var err error
	if gmerr != nil {
		if cmerr != nil {
			err = fmt.Errorf("Go machine error: %v, cpp machine error: %v ", gmerr, cmerr)
		} else {
			err = fmt.Errorf("Go machine error: %v", gmerr)
		}
	} else if cmerr != nil {
		err = fmt.Errorf("cpp machine error: %v ", cmerr)
	}
	return &Machine{
		cm,
		gm,
		proofMode,
		nil,
	}, err
}

func (m *Machine) Hash() [32]byte {
	h1 := m.cppmachine.Hash()
	h2 := m.gomachine.Hash()
	if h1 != h2 {
		log.Fatalln("Hash error at pc", m.gomachine.GetPC())
	}
	return h1
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{m.cppmachine.Clone().(*cmachine.Machine), m.gomachine.Clone().(*gomachine.Machine), m.proofMode, m.data}
}

func (m *Machine) CurrentStatus() machine.Status {
	b1 := m.cppmachine.CurrentStatus()
	b2 := m.gomachine.CurrentStatus()
	if b1 != b2 {
		log.Fatalln("CurrentStatus error at pc", m.gomachine.GetPC())
	}
	return b1
}

func (m *Machine) LastBlockReason() machine.BlockReason {
	b1 := m.cppmachine.LastBlockReason()
	b2 := m.gomachine.LastBlockReason()
	if b1 == nil || b2 == nil {
		if b1 != nil || b2 != nil {
			log.Fatalln("LastBlockReason error at pc", m.gomachine.GetPC())
		}
		return nil
	}
	if !b1.Equals(b2) {
		log.Fatalln("LastBlockReason error at pc", m.gomachine.GetPC())
	}
	return b1
}

func (m *Machine) CanSpend(tokenType protocol.TokenType, currency *big.Int) bool {
	b1 := m.cppmachine.CanSpend(tokenType, currency)
	b2 := m.gomachine.CanSpend(tokenType, currency)
	if b1 != b2 {
		log.Fatalln("CanSpend error at pc", m.gomachine.GetPC())
	}
	return b1
}

func (m *Machine) InboxHash() value.HashOnlyValue {
	h1 := m.cppmachine.InboxHash()
	h2 := m.gomachine.InboxHash()
	if !h1.Equal(h2) {
		log.Fatalln("InboxHash error at pc", m.gomachine.GetPC())
	}
	return h1
}

func (m *Machine) PendingMessageCount() uint64 {
	h1 := m.cppmachine.PendingMessageCount()
	h2 := m.gomachine.PendingMessageCount()
	if h1 != h2 {
		log.Fatalln("PendingMessageCount error", h1, h2, "at pc", m.gomachine.GetPC())
	}
	return h1
}

func (m *Machine) SendOnchainMessage(msg protocol.Message) {
	m.cppmachine.SendOnchainMessage(msg)
	m.gomachine.SendOnchainMessage(msg)
}

func (m *Machine) DeliverOnchainMessage() {
	m.cppmachine.DeliverOnchainMessage()
	m.gomachine.DeliverOnchainMessage()
}

func (m *Machine) SendOffchainMessages(msgs []protocol.Message) {
	m.cppmachine.SendOffchainMessages(msgs)
	m.gomachine.SendOffchainMessages(msgs)
}

func (m *Machine) ProofMachineData(contractAddress common.Address, coordinator *channel.ValidatorCoordinator, key *ecdsa.PrivateKey) {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	pd := ProofMachData{
		coord:       coordinator,
		fromAddress: keyAddr,
	}
	m.data = &pd
}

func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	if m.proofMode {
		if m.data == nil {
			log.Println("Proof data not set")
			return m.ExecAssertion(maxSteps, timeBounds)
		}
		gomach := m.gomachine.Clone()
		cppmach := m.cppmachine.Clone()
		var timeBounds [2]uint64

		//a := &protocol.Assertion{}
		stepIncrease := int32(1)
		maxSteps := int32(1000)
		stepsRan := 0
		for i := int32(0); i < maxSteps; i += stepIncrease {
			timeBounds[0] = uint64(i)
			timeBounds[1] = uint64(i + stepIncrease)
			proof, err := gomach.MarshalForProof()
			steps := int32(stepIncrease)
			beforeHash := gomach.Hash()
			inboxHash := gomach.InboxHash()

			pcStart := m.gomachine.GetPC()
			a1 := cppmach.ExecuteAssertion(steps, timeBounds)
			a2 := gomach.ExecuteAssertion(steps, timeBounds)
			if !a1.Equals(a2) {
				pcEnd := m.gomachine.GetPC()
				log.Fatalln("ExecuteAssertion error after running step", pcStart, pcEnd, a1, a2)
			}
			if a1.NumSteps == 0 {
				fmt.Println(" machine halted ")
				break
			}
			if a1.NumSteps != 1 {
				log.Println("Num steps = ", a1.NumSteps)
			}
			stepsRan++
			fmt.Println("executed up to step ", i)
			spentBalance := protocol.NewTokenTrackerFromMessages(a1.OutMsgs)
			balance := m.data.coord.ChannelVal.GetBalance().Clone()
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

			res, err := m.data.coord.Val.Validator.OneStepProof.ValidateProof(callOpts, precond, a1.Stub(), proof)
			if err != nil {
				log.Fatal("Proof invalid", err)
			}
			if res.Cmp(big.NewInt(0)) == 0 {
				log.Println("Proof valid")
			} else {
				log.Fatal("Proof invalid")
			}
		}
		fmt.Println("Proof mode ran ", stepsRan, " steps")
	}
	return m.ExecAssertion(maxSteps, timeBounds)
}

func (m *Machine) ExecAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	a := &protocol.Assertion{}
	stepIncrease := int32(50)
	for i := int32(0); i < maxSteps; i += stepIncrease {
		steps := maxSteps - i
		if steps > stepIncrease {
			steps = stepIncrease
		}

		pcStart := m.gomachine.GetPC()
		a1 := m.cppmachine.ExecuteAssertion(steps, timeBounds)
		a2 := m.gomachine.ExecuteAssertion(steps, timeBounds)

		if !a1.Equals(a2) {
			pcEnd := m.gomachine.GetPC()
			log.Fatalln("ExecuteAssertion error after running step", pcStart, pcEnd, a1, a2)
		}
		a.AfterHash = a1.AfterHash
		a.NumSteps += a1.NumSteps
		a.Logs = append(a.Logs, a1.Logs...)
		a.OutMsgs = append(a.OutMsgs, a1.OutMsgs...)

		if a1.NumSteps < uint32(steps) {
			break
		}
	}
	fmt.Println("Assertion ran", a.NumSteps, "steps")
	return a
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	h1, err1 := m.cppmachine.MarshalForProof()
	h2, err2 := m.gomachine.MarshalForProof()

	if err1 != nil {
		return nil, err1
	}

	if err2 != nil {
		return nil, err2
	}

	if !bytes.Equal(h1, h2) {
		log.Fatalln("MarshalForProof error at pc", m.gomachine.GetPC())
	}
	return h1, nil
}
