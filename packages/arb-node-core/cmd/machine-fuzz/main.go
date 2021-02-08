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

package main

/*
#include "stdint.h"
extern void HF_ITER(uint8_t** buf, size_t* len);
void crash() { __builtin_trap(); }
*/
import "C"

import (
	"context"
	"fmt"
	"os"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"

	"io/ioutil"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
)

func runFuzzValidateProof(contract string, osp *ethbridgetestcontracts.OneStepProof, osp2 *ethbridgetestcontracts.OneStepProof2) {
	ctx := context.Background()
	proofs, _, err := challenge.GenerateProofCases(contract, 20)
	if err != nil {
		println("Error generating proofs: " + err.Error())
		return
	}

	/*
		data, err := json.Marshal(proofs)
		if err != nil {
			panic(err)
		}

		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			err := errors.New("failed to get filename")
			panic(err)
		}

		file := filepath.Join(filepath.Dir(filename), "../../arb-bridge-eth/test/proofs", filepath.Base(contract)+"-proofs.json")

		if err := ioutil.WriteFile(file, data, 0644); err != nil {
			panic(err)
		}
	*/

	for _, proof := range proofs {
		op := proof.Proof[0]
		fmt.Printf("Checking proof of opcode 0x%x\n", op)
		var err error
		var machineData struct {
			Gas    uint64
			Fields [5][32]byte
		}
		machineFields := [3][32]byte{
			proof.BeforeCut.InboxDelta,
			proof.BeforeCut.SendAcc,
			proof.BeforeCut.LogAcc,
		}
		if (op >= 0xa1 && op <= 0xa6) || op == 0x70 {
			machineData, err = osp2.ExecuteStep(
				&bind.CallOpts{Context: ctx},
				machineFields,
				proof.Proof,
				proof.BufferProof,
			)
		} else {
			machineData, err = osp.ExecuteStep(
				&bind.CallOpts{Context: ctx},
				machineFields,
				proof.Proof,
			)
		}
		if err != nil {
			panic(err)
		}

		//t.Log("Opcode", opcode)
		if err != nil {
			panic(err)
		}
		correctGasUsed := proof.AfterCut.GasUsed - proof.BeforeCut.GasUsed
		if machineData.Gas != correctGasUsed {
			panic(fmt.Sprintf("wrong gas %v instead of %v", machineData.Gas, correctGasUsed))
		}
		if machineData.Fields[0] != proof.BeforeCut.MachineState {
			panic("wrong before machine")
		}
		if machineData.Fields[2] != proof.AfterCut.InboxDelta {
			panic("wrong DidInboxInsn")
		}
		/*
			if machineData.Fields[3] != proof.AfterCut.SendAcc {
				println("Machine state:")
				println(hex.EncodeToString(machineData.Fields[3][:]))
				println(hex.EncodeToString(machineData.Fields[4][:]))
				println("Proof state:")
				println(hex.EncodeToString(proof.AfterCut.SendAcc[:]))
				println(hex.EncodeToString(proof.AfterCut.LogAcc[:]))
				panic("wrong log")
			}
			if machineData.Fields[4] != proof.AfterCut.LogAcc {
				println("Machine state:")
				println(hex.EncodeToString(machineData.Fields[3][:]))
				println(hex.EncodeToString(machineData.Fields[4][:]))
				println("Proof state:")
				println(hex.EncodeToString(proof.AfterCut.SendAcc[:]))
				println(hex.EncodeToString(proof.AfterCut.LogAcc[:]))
				panic("wrong message")
			}
		*/
		if machineData.Fields[1] != proof.AfterCut.MachineState {
			panic("wrong after machine")
		}
	}
}

func main() {
	clnt, pks := test.SimulatedBackend()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}
	auth := bind.NewKeyedTransactor(pks[0])
	_, _, osp, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	if err != nil {
		panic(err)
	}
	_, _, osp2, err := ethbridgetestcontracts.DeployOneStepProof2(auth, client)
	if err != nil {
		panic(err)
	}
	client.Commit()

	if len(os.Args) > 1 {
		runFuzzValidateProof(os.Args[1], osp, osp2)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			print(err)
			C.crash()
		}
	}()

Iter:
	for {
		var ptr *C.uint8_t
		var size C.size_t
		C.HF_ITER(&ptr, &size)
		contract := C.GoBytes(unsafe.Pointer(ptr), C.int(size))

		for _, b := range contract {
			if b < ' ' || b > '~' {
				continue Iter
			}
		}

		contractFile, err := ioutil.TempFile("/tmp", "arbitrum-proofmachine-fuzz")
		if err != nil {
			panic(err)
		}
		contractFile.Close()
		if err := ioutil.WriteFile(contractFile.Name(), contract, 0644); err != nil {
			panic(err)
		}

		runFuzzValidateProof(contractFile.Name(), osp, osp2)

		os.Remove(contractFile.Name())
	}
}
