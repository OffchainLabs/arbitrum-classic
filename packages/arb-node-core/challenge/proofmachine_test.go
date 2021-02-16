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
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
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
	db := core.NewValidatorLookupMock(mach)
	for i := 0; i < 100; i++ {
		db.AddMessage(inbox.NewRandomInboxMessage())
	}

	beforeCut := ExecutionCutJSON{
		GasUsed:           0,
		TotalMessagesRead: (*hexutil.Big)(big.NewInt(0)),
		MachineState:      mach.Hash().ToEthHash(),
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

		messages, err := db.GetMessages(big.NewInt(0), big.NewInt(1))
		if err != nil {
			return nil, nil, err
		}

		a, _, ranSteps := mach.ExecuteAssertionAdvanced(
			1,
			true,
			messages,
			false,
			common.NewHashFromEth(beforeCut.SendAcc),
			common.NewHashFromEth(beforeCut.LogAcc),
		)
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

		afterCut := ExecutionCutJSON{
			GasUsed:           beforeCut.GasUsed + a.NumGas,
			TotalMessagesRead: (*hexutil.Big)(new(big.Int).Add(beforeCut.TotalMessagesRead.ToInt(), new(big.Int).SetUint64(a.InboxMessagesConsumed))),
			MachineState:      mach.Hash().ToEthHash(),
			SendAcc:           ethcommon.Hash{},
			SendCount:         (*hexutil.Big)(new(big.Int).Add(beforeCut.SendCount.ToInt(), big.NewInt(int64(len(a.Sends))))),
			LogAcc:            ethcommon.Hash{},
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

func runTestValidateProof(t *testing.T, contract string, osp *ethbridgetestcontracts.OneStepProof, osp2 *ethbridgetestcontracts.OneStepProof2, bridge ethcommon.Address) {
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
		opcode := proof.Proof[len(proof.Proof)-1]
		t.Run(strconv.FormatUint(uint64(opcode), 10), func(t *testing.T) {
			var err error
			var machineData struct {
				Gas               uint64
				TotalMessagesRead *big.Int
				Fields            [4][32]byte
			}

			op := proof.Proof[0]
			t.Log("Opcode", opcode)
			if (op >= 0xa1 && op <= 0xa6) || op == 0x70 {
				machineData, err = osp2.ExecuteStep(
					&bind.CallOpts{Context: ctx},
					proof.BeforeCut.TotalMessagesRead.ToInt(),
					proof.BeforeCut.SendAcc,
					proof.BeforeCut.LogAcc,
					proof.Proof,
					proof.BufferProof,
				)
			} else {
				var messagesRead [32]byte
				copy(messagesRead[:], math.U256Bytes(proof.BeforeCut.TotalMessagesRead.ToInt()))
				ret, err := osp.ExecuteStepDebug(
					&bind.CallOpts{Context: ctx},
					bridge,
					[3][32]byte{
						messagesRead,
						proof.BeforeCut.SendAcc,
						proof.BeforeCut.LogAcc,
					},
					proof.Proof,
				)
				test.FailIfError(t, err)
				machineData.Fields = ret.Fields
				machineData.Gas = ret.Gas
			}
			test.FailIfError(t, err)
			correctGasUsed := proof.AfterCut.GasUsed - proof.BeforeCut.GasUsed
			if machineData.Gas != correctGasUsed {
				t.Fatalf("wrong gas %v instead of %v", machineData.Gas, correctGasUsed)
			}
			if machineData.TotalMessagesRead.Cmp(proof.AfterCut.TotalMessagesRead.ToInt()) != 0 {
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
	testMachines := gotest.OpCodeTestFiles()
	backend, pks := test.SimulatedBackend()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := bind.NewKeyedTransactor(pks[0])

	_, _, osp, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	test.FailIfError(t, err)
	_, _, osp2, err := ethbridgetestcontracts.DeployOneStepProof2(auth, client)
	test.FailIfError(t, err)
	bridgeAddr, _, _, err := ethbridgecontracts.DeployBridge(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	for _, machName := range testMachines {
		machName := machName // capture range variable
		t.Run(machName, func(t *testing.T) {
			//t.Parallel()
			runTestValidateProof(t, machName, osp, osp2, bridgeAddr)
		})
	}
}
