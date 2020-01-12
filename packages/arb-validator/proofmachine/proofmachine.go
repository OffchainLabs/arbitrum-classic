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

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type Machine struct {
	machine machine.Machine
	ethConn *Connection
}

type Connection struct {
	fromAddress common.Address
	osp         arbbridge.OneStepProof
	client      arbbridge.ArbClient
	proofbounds [2]uint32
}

func NewEthConnection(contractAddress common.Address, key *ecdsa.PrivateKey, ethURL string, proofbounds [2]uint32) (*Connection, error) {
	client, err := ethbridge.NewEthClient(ethURL)
	if err != nil {
		log.Fatal("Connection failure ", err)
	}
	osp, err := client.NewOneStepProof(contractAddress)
	if err != nil {
		log.Fatal(err)
	}

	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	return &Connection{
		fromAddress: common.NewAddressFromEth(keyAddr),
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

func (m *Machine) Hash() common.Hash {
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

func (m *Machine) ExecuteAssertion(maxSteps uint32, timeBounds *protocol.TimeBoundsBlocks, inbox value.TupleValue) (*protocol.ExecutionAssertion, uint32) {
	a := &protocol.ExecutionAssertion{}
	totalSteps := uint32(0)
	stepIncrease := uint32(1)
	stepsRan := 0
	for i := uint32(0); i < maxSteps; i += stepIncrease {
		var proof []byte
		var err error
		// only marshall if we are going to validate (see below)
		if i >= m.ethConn.proofbounds[0] && i <= m.ethConn.proofbounds[1] {
			proof, err = m.MarshalForProof()
			if err != nil {
				log.Println("error marshaling")
			}
		}
		beforeHash := m.Hash()
		a1, ranSteps := m.machine.ExecuteAssertion(stepIncrease, timeBounds, inbox)
		a.AfterHash = a1.AfterHash
		totalSteps += ranSteps
		a.NumGas += a1.NumGas
		a.Logs = append(a.Logs, a1.Logs...)
		a.OutMsgs = append(a.OutMsgs, a1.OutMsgs...)

		if ranSteps == 0 {
			fmt.Println(" machine halted ")
			break
		}
		if ranSteps != 1 {
			log.Println("Num steps = ", ranSteps)
		}
		stepsRan++

		// only marshall and validate if step is within proofbounds
		if i >= m.ethConn.proofbounds[0] && i <= m.ethConn.proofbounds[1] {
			// uncomment to force proof fail
			//beforeHash[0] = 5
			precond := valprotocol.NewPrecondition(beforeHash, timeBounds, inbox)

			res, err := m.ethConn.osp.ValidateProof(context.Background(), precond, valprotocol.NewExecutionAssertionStubFromAssertion(a1), proof)
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

		if a1.DidInboxInsn {
			inbox = value.NewEmptyTuple()
		}
	}
	fmt.Println("Proof mode ran ", stepsRan, " steps")
	return a, totalSteps
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	return m.machine.MarshalForProof()
}

func (m *Machine) Checkpoint(storage machine.CheckpointStorage) bool {
	return m.machine.Checkpoint(storage)
}

func (m *Machine) RestoreCheckpoint(storage machine.CheckpointStorage, machineHash common.Hash) bool {
	return m.machine.RestoreCheckpoint(storage, machineHash)
}
