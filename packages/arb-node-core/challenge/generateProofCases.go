/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package challenge

import (
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type ExecutionCutJSON struct {
	GasUsed      uint64
	InboxDelta   ethcommon.Hash
	MachineState ethcommon.Hash
	SendAcc      ethcommon.Hash
	SendCount    *hexutil.Big
	LogAcc       ethcommon.Hash
	LogCount     *hexutil.Big
}

type ProofData struct {
	BeforeCut   ExecutionCutJSON
	AfterCut    ExecutionCutJSON
	Proof       hexutil.Bytes
	BufferProof hexutil.Bytes
}

func GenerateProofCases(contract string, maxSteps uint64) ([]*ProofData, []string, error) {
	mach, err := cmachine.New(contract)
	if err != nil {
		return nil, nil, err
	}

	db := core.NewValidatorLookupMock(mach)
	for i := 0; i < 100; i++ {
		db.AddMessage(inbox.NewRandomInboxMessage())
	}

	beforeCut := ExecutionCutJSON{
		GasUsed:      0,
		InboxDelta:   ethcommon.Hash{},
		MachineState: mach.Hash().ToEthHash(),
		SendAcc:      ethcommon.Hash{},
		SendCount:    (*hexutil.Big)(big.NewInt(0)),
		LogAcc:       ethcommon.Hash{},
		LogCount:     (*hexutil.Big)(big.NewInt(0)),
	}
	nextMessageIndex := big.NewInt(0)
	proofs := make([]*ProofData, 0)
	machineStates := make([]string, 0)
	machineStates = append(machineStates, mach.String())
	for i := uint64(0); i < maxSteps; i++ {
		proof, bproof, err := mach.MarshalForProof()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Marshalled proof for opcode 0x%x\n", proof[0])

		messages, err := db.GetMessages(big.NewInt(0), big.NewInt(1))
		if err != nil {
			panic(err)
		}

		a, _, ranSteps := mach.ExecuteAssertionAdvanced(
			1,
			true,
			messages,
			false,
			common.NewHashFromEth(beforeCut.SendAcc),
			common.NewHashFromEth(beforeCut.LogAcc),
		)
		fmt.Println("Ran", ranSteps)
		if ranSteps == 0 {
			break
		}
		if ranSteps != 1 {
			panic("executed incorrect step count")
		}
		machineStates = append(machineStates, mach.String())
		if mach.CurrentStatus() == machine.ErrorStop {
			fmt.Println("Machine stopped in error state")
			return proofs, nil, nil
		}
		if a.InboxMessagesConsumed > 0 {
			fmt.Println("TODO: Inbox is currently unimplemented; stopping")
			return proofs, nil, nil

			inboxDeltaHash, err := db.GetInboxDelta(big.NewInt(0), big.NewInt(1))
			if err != nil {
				return nil, nil, err
			}
			beforeCut.InboxDelta = inboxDeltaHash.ToEthHash()
		}

		afterCut := ExecutionCutJSON{
			GasUsed:      beforeCut.GasUsed + a.NumGas,
			InboxDelta:   ethcommon.Hash{},
			MachineState: mach.Hash().ToEthHash(),
			SendAcc:      ethcommon.Hash{},
			SendCount:    (*hexutil.Big)(new(big.Int).Add(beforeCut.SendCount.ToInt(), big.NewInt(int64(len(a.Sends))))),
			LogAcc:       ethcommon.Hash{},
			LogCount:     (*hexutil.Big)(new(big.Int).Add(beforeCut.LogCount.ToInt(), big.NewInt(int64(len(a.Logs))))),
		}

		proofs = append(proofs, &ProofData{
			BeforeCut:   beforeCut,
			AfterCut:    afterCut,
			Proof:       proof,
			BufferProof: bproof,
		})
		beforeCut = afterCut
		nextMessageIndex = nextMessageIndex.Add(nextMessageIndex, new(big.Int).SetUint64(a.InboxMessagesConsumed))
	}
	return proofs, machineStates, nil
}
