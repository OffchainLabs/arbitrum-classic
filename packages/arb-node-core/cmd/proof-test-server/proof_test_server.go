/*
 * Copyright 2020, Offchain Labs, Inc.
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

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"os"
	"testing"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/proofmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func handleFatalError(err error) {
	if err != nil {
		if errors.Is(err, io.EOF) {
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "Fatal error: %v\n", err)
			os.Exit(1)
		}
	}
}

func readStateData(reader io.Reader) *core.ExecutionState {
	buf := make([]byte, 32*5)
	_, err := io.ReadFull(reader, buf)
	handleFatalError(err)
	data := &core.ExecutionState{
		InboxAcc:          common.Hash{},
		TotalMessagesRead: new(big.Int).SetBytes(buf[32:64]),
		TotalGasConsumed:  new(big.Int).SetBytes(buf[:32]),
		TotalSendCount:    nil,
		TotalLogCount:     nil,
	}
	copy(data.MachineHash[:], buf[64:96])
	copy(data.SendAcc[:], buf[96:128])
	copy(data.LogAcc[:], buf[128:160])
	return data
}

func readBytes(reader io.Reader) []byte {
	buf := make([]byte, 8)
	_, err := io.ReadFull(reader, buf)
	handleFatalError(err)
	buf = make([]byte, int(binary.BigEndian.Uint64(buf)))
	_, err = io.ReadFull(reader, buf)
	handleFatalError(err)
	return buf
}

func readProofData(reader io.Reader) *proofmachine.ProofData {
	before := readStateData(reader)
	proof := readBytes(reader)
	bufferProof := readBytes(reader)
	after := readStateData(reader)
	return &proofmachine.ProofData{
		Assertion: &core.Assertion{
			Before: before,
			After:  after,
		},
		Proof:       proof,
		BufferProof: bufferProof,
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: proof_test_server [query pipe] [result path]\n")
		os.Exit(1)
	}

	queryPipe, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	handleFatalError(err)
	resultPipe, err := os.OpenFile(os.Args[2], os.O_WRONLY, 0)
	handleFatalError(err)

	buf := make([]byte, 1)
	_, err = io.ReadFull(queryPipe, buf)
	handleFatalError(err)
	globalFlags := buf[0]
	debug := globalFlags&(1<<0) != 0

	if debug {
		fmt.Fprintf(os.Stderr, "Starting Go proof test server in debug mode\n")
	}

	buf[0] = 0
	_, err = resultPipe.Write(buf)
	handleFatalError(err)

	backend, auths := test.SimulatedBackend(&testing.T{})
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := auths[0]
	proofChecker, err := proofmachine.NewProofChecker(auth, client)
	handleFatalError(err)

	for {
		proof := readProofData(queryPipe)
		if debug {
			fmt.Fprintf(os.Stderr, "Proving execution of opcode 0x%x\n", proof.Proof[0])
		}
		proofErrors := proofChecker.CheckProof(proof)
		retByte := uint8(0)
		if len(proofErrors) > 0 {
			retByte = 1
			fmt.Fprintln(os.Stderr, "Error verifying proof:")
		}
		for _, err := range proofChecker.CheckProof(proof) {
			fmt.Fprintln(os.Stderr, err)
		}
		_, err = resultPipe.Write([]byte{retByte})
		handleFatalError(err)
	}
}
