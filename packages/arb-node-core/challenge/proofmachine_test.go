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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
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

type proofData struct {
	BeforeCut   ExecutionCutJSON
	AfterCut    ExecutionCutJSON
	Proof       hexutil.Bytes
	BufferProof hexutil.Bytes
}

func generateProofCases(contract string) ([]*proofData, []string, error) {
	mach, err := cmachine.New(contract)
	if err != nil {
		return nil, nil, err
	}

	maxSteps := uint64(100000)
	messages := make([]inbox.InboxMessage, 0)
	for i := 0; i < 100; i++ {
		messages = append(messages, inbox.NewRandomInboxMessage())
	}

	hash := mach.Hash()
	beforeCut := ExecutionCutJSON{
		GasUsed:           0,
		TotalMessagesRead: (*hexutil.Big)(big.NewInt(0)),
		MachineState:      hash.ToEthHash(),
		SendAcc:           ethcommon.Hash{},
		SendCount:         (*hexutil.Big)(big.NewInt(0)),
		LogAcc:            ethcommon.Hash{},
		LogCount:          (*hexutil.Big)(big.NewInt(0)),
	}

	nextMessageIndex := big.NewInt(0)

	proofs := make([]*proofData, 0)
	machineStates := make([]string, 0)
	machineStates = append(machineStates, mach.String())
	for i := uint64(0); i < maxSteps; i++ {
		proof, bproof, err := mach.MarshalForProof()
		if err != nil {
			return nil, nil, err
		}

		messages := messages[:1]

		a, _, ranSteps, err := mach.ExecuteAssertionAdvanced(
			1,
			true,
			messages,
			nil,
			false,
			common.NewHashFromEth(beforeCut.SendAcc),
			common.NewHashFromEth(beforeCut.LogAcc),
		)
		if err != nil {
			return nil, nil, err
		}
		if ranSteps == 0 {
			break
		}
		if ranSteps != 1 {
			return nil, nil, errors.New("executed incorrect step count")
		}
		machineStates = append(machineStates, mach.String())
		if mach.CurrentStatus() == machine.ErrorStop {
			fmt.Println("Machine stopped in error state")
			return proofs, nil, nil
		}

		hash := mach.Hash()
		afterCut := ExecutionCutJSON{
			GasUsed:           beforeCut.GasUsed + a.NumGas,
			TotalMessagesRead: (*hexutil.Big)(new(big.Int).Add(beforeCut.TotalMessagesRead.ToInt(), new(big.Int).SetUint64(a.InboxMessagesConsumed))),
			MachineState:      hash.ToEthHash(),
			SendAcc:           a.SendAcc.ToEthHash(),
			SendCount:         (*hexutil.Big)(new(big.Int).Add(beforeCut.SendCount.ToInt(), big.NewInt(int64(len(a.Sends))))),
			LogAcc:            a.LogAcc.ToEthHash(),
			LogCount:          (*hexutil.Big)(new(big.Int).Add(beforeCut.LogCount.ToInt(), big.NewInt(int64(len(a.Logs))))),
		}

		proofs = append(proofs, &proofData{
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

func getProverNum(op uint8) uint8 {
	if (op >= 0xa1 && op <= 0xa6) || op == 0x70 {
		return 1
	} else if op >= 0x20 && op <= 0x24 {
		return 2
	} else {
		return 0
	}
}

func runTestValidateProof(t *testing.T, contract string, osps []*ethbridgetestcontracts.IOneStepProof, sequencerBridge, delayedBridge ethcommon.Address) {
	t.Log("proof test contact: ", contract)
	ctx := context.Background()

	proofs, _, err := generateProofCases(contract)
	if err != nil {
		t.Fatal(err)
	}

	data, err := json.Marshal(proofs)
	if err != nil {
		t.Fatal(err)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		err := errors.New("failed to get filename")
		t.Fatal(err)
	}

	file := filepath.Join(filepath.Dir(filename), "../../arb-bridge-eth/test/proofs", filepath.Base(contract)+"-proofs.json")
	if err := ioutil.WriteFile(file, data, 0644); err != nil {
		t.Fatal(err)
	}

	for _, proof := range proofs {
		op := proof.Proof[0]
		t.Run(strconv.FormatUint(uint64(op), 10), func(t *testing.T) {
			prover := getProverNum(op)
			t.Logf("Opcode 0x%x with prover %v", op, prover)
			machineData, err := osps[prover].ExecuteStep(
				&bind.CallOpts{Context: ctx},
				[2]ethcommon.Address{sequencerBridge, delayedBridge},
				proof.BeforeCut.TotalMessagesRead.ToInt(),
				[2][32]byte{
					proof.BeforeCut.SendAcc,
					proof.BeforeCut.LogAcc,
				},
				proof.Proof,
				proof.BufferProof,
			)
			test.FailIfError(t, err)
			correctGasUsed := proof.AfterCut.GasUsed - proof.BeforeCut.GasUsed
			if machineData.Gas != correctGasUsed {
				t.Fatalf("wrong gas %v instead of %v", machineData.Gas, correctGasUsed)
			}
			if machineData.AfterMessagesRead.Cmp(proof.AfterCut.TotalMessagesRead.ToInt()) != 0 {
				t.Fatal("wrong total messages read")
			}
			if machineData.Fields[0] != proof.BeforeCut.MachineState {
				t.Fatal("wrong before machine")
			}
			if machineData.Fields[2] != proof.AfterCut.SendAcc {
				t.Fatal("wrong log")
			}
			if machineData.Fields[3] != proof.AfterCut.LogAcc {
				t.Fatal("wrong message")
			}
			if machineData.Fields[1] != proof.AfterCut.MachineState {
				t.Fatalf("wrong after machine 0x%x 0x%x", machineData.Fields[1][:], proof.AfterCut.MachineState[:])
			}
		})
	}
}

func TestValidateProof(t *testing.T) {
	testMachines, err := gotest.OpCodeTestFiles()
	test.FailIfError(t, err)
	backend, pks := test.SimulatedBackend(t)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := bind.NewKeyedTransactor(pks[0])
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
	client.Commit()
	_, err = sequencerCon.Initialize(auth, delayedBridgeAddr, sequencer, maxDelayBlocks, maxDelaySeconds)
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
			runTestValidateProof(t, machName, provers, sequencerAddr, delayedBridgeAddr)
		})
	}
}
