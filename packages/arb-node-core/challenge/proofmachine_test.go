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
	"encoding/json"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"runtime"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

type ExecutionCutJSON struct {
	GasUsed           uint64
	TotalMessagesRead *hexutil.Big
	MachineState      ethcommon.Hash
	SendAcc           ethcommon.Hash
	SendCount         *hexutil.Big
	LogAcc            ethcommon.Hash
	LogCount          *hexutil.Big
}

func NewExecutionCutJSONFromCursor(cursor core.ExecutionCursor) ExecutionCutJSON {
	return ExecutionCutJSON{
		GasUsed:           cursor.TotalGasConsumed().Uint64(),
		TotalMessagesRead: (*hexutil.Big)(cursor.TotalMessagesRead()),
		MachineState:      cursor.MachineHash().ToEthHash(),
		SendAcc:           cursor.SendAcc().ToEthHash(),
		SendCount:         (*hexutil.Big)(cursor.TotalSendCount()),
		LogAcc:            cursor.LogAcc().ToEthHash(),
		LogCount:          (*hexutil.Big)(cursor.TotalLogCount()),
	}
}

type proofData struct {
	BeforeCut   ExecutionCutJSON
	AfterCut    ExecutionCutJSON
	Proof       hexutil.Bytes
	BufferProof hexutil.Bytes
}

func generateProofCases(t *testing.T, arbCore *monitor.Monitor) ([]*proofData, []string) {
	var cursors []core.ExecutionCursor
	cursor, err := arbCore.Core.GetExecutionCursor(big.NewInt(0))
	test.FailIfError(t, err)
	cursors = append(cursors, cursor.Clone())
	for {
		err = arbCore.Core.AdvanceExecutionCursor(cursor, big.NewInt(1), true)
		test.FailIfError(t, err)
		if cursor.TotalGasConsumed().Cmp(cursors[len(cursors)-1].TotalGasConsumed()) == 0 {
			break
		}
		cursors = append(cursors, cursor.Clone())
	}

	proofs := make([]*proofData, 0)
	machineStates := make([]string, 0)

	for i := 0; i < len(cursors)-1; i++ {
		mach, err := arbCore.Core.TakeMachine(cursors[i].Clone())
		test.FailIfError(t, err)
		machineStates = append(machineStates, mach.String())

		proof, bproof, err := mach.MarshalForProof()
		test.FailIfError(t, err)

		proofs = append(proofs, &proofData{
			BeforeCut:   NewExecutionCutJSONFromCursor(cursors[i]),
			AfterCut:    NewExecutionCutJSONFromCursor(cursors[i+1]),
			Proof:       proof,
			BufferProof: bproof,
		})
	}
	mach, err := arbCore.Core.TakeMachine(cursors[len(cursors)-1].Clone())
	test.FailIfError(t, err)
	machineStates = append(machineStates, mach.String())
	return proofs, machineStates
}

func getProverNum(op uint8) uint8 {
	if (op >= 0xa1 && op <= 0xa6) || op == 0x70 {
		return 1
	} else if op >= 0x20 && op <= 0x24 {
		return 2
	} else {
		return 0
	}
}

func checkProofs(t *testing.T, proofs []*proofData, osps []*ethbridgetestcontracts.IOneStepProof, sequencerBridge, delayedBridge ethcommon.Address) {
	for _, proof := range proofs {
		op := proof.Proof[0]
		prover := getProverNum(op)
		machineData, err := osps[prover].ExecuteStep(
			&bind.CallOpts{},
			[2]ethcommon.Address{sequencerBridge, delayedBridge},
			proof.BeforeCut.TotalMessagesRead.ToInt(),
			[2][32]byte{
				proof.BeforeCut.SendAcc,
				proof.BeforeCut.LogAcc,
			},
			proof.Proof,
			proof.BufferProof,
		)
		if err != nil {
			t.Logf("Opcode 0x%x with prover %v", op, prover)
			t.Error(err)
			continue
		}
		correctGasUsed := proof.AfterCut.GasUsed - proof.BeforeCut.GasUsed
		if machineData.Gas != correctGasUsed {
			t.Errorf("wrong gas %v instead of %v", machineData.Gas, correctGasUsed)
		}
		if machineData.AfterMessagesRead.Cmp(proof.AfterCut.TotalMessagesRead.ToInt()) != 0 {
			t.Error("wrong total messages read")
		}
		if machineData.Fields[0] != proof.BeforeCut.MachineState {
			t.Error("wrong before machine")
		}
		if machineData.Fields[2] != proof.AfterCut.SendAcc {
			t.Error("wrong log")
		}
		if machineData.Fields[3] != proof.AfterCut.LogAcc {
			t.Error("wrong message")
		}
		if machineData.Fields[1] != proof.AfterCut.MachineState {
			t.Errorf("wrong after machine 0x%x 0x%x", machineData.Fields[1][:], proof.AfterCut.MachineState[:])
		}
	}
}

func TestValidateProof(t *testing.T) {
	testMachines, err := gotest.OpCodeTestFiles()
	test.FailIfError(t, err)
	backend, auths := test.SimulatedBackend(t)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := auths[0]
	sequencer := common.RandAddress().ToEthAddress()
	maxDelayBlocks := big.NewInt(60)
	maxDelaySeconds := big.NewInt(900)

	osp1Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	test.FailIfError(t, err)
	osp2Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof2(auth, client)
	test.FailIfError(t, err)
	osp3Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProofHash(auth, client)
	test.FailIfError(t, err)
	delayedBridgeAddr, _, _, err := ethbridgecontracts.DeployBridge(auth, client)
	test.FailIfError(t, err)
	sequencerAddr, _, sequencerCon, err := ethbridgecontracts.DeploySequencerInbox(auth, client)
	test.FailIfError(t, err)
	rollupAddr, _, rollup, err := ethbridgetestcontracts.DeployRollupMock(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	_, err = rollup.SetMock(auth, maxDelayBlocks, maxDelaySeconds)
	test.FailIfError(t, err)
	_, err = sequencerCon.Initialize(auth, delayedBridgeAddr, sequencer, rollupAddr)
	test.FailIfError(t, err)
	client.Commit()

	osp1, err := ethbridgetestcontracts.NewIOneStepProof(osp1Addr, client)
	test.FailIfError(t, err)
	osp2, err := ethbridgetestcontracts.NewIOneStepProof(osp2Addr, client)
	test.FailIfError(t, err)
	osp3, err := ethbridgetestcontracts.NewIOneStepProof(osp3Addr, client)
	test.FailIfError(t, err)
	provers := []*ethbridgetestcontracts.IOneStepProof{osp1, osp2, osp3}
	t.Log(testMachines)

	for _, machName := range testMachines {
		machName := machName // capture range variable
		t.Run(machName, func(t *testing.T) {
			//t.Parallel()
			arbCore, cancel := monitor.PrepareArbCoreWithMexe(t, machName)
			proofs, _ := generateProofCases(t, arbCore)
			cancel()

			data, err := json.Marshal(proofs)
			test.FailIfError(t, err)

			_, filename, _, ok := runtime.Caller(0)
			if !ok {
				t.Fatal("failed to get filename")
			}

			file := filepath.Join(filepath.Dir(filename), "../../arb-bridge-eth/test/proofs", filepath.Base(machName)+"-proofs.json")
			err = ioutil.WriteFile(file, data, 0644)
			test.FailIfError(t, err)

			checkProofs(t, proofs, provers, sequencerAddr, delayedBridgeAddr)
		})
	}
}
