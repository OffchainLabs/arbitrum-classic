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

import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"os"
	"testing"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/proofmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

var proofCheckers = make(map[int]*proofmachine.ProofChecker)
var debug bool

func getProofChecker(index int) (*proofmachine.ProofChecker, error) {
	checker, ok := proofCheckers[index]
	if ok {
		return checker, nil
	}
	backend, auths := test.SimulatedBackend(&testing.T{})
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := auths[0]
	checker, err := proofmachine.NewProofChecker(auth, client)
	if err != nil {
		return nil, err
	}
	proofCheckers[index] = checker
	return checker, nil
}

//export CheckProof
func CheckProof(checkerIndex C.int, dataPtr unsafe.Pointer, length C.int) C.int {
	data := C.GoBytes(dataPtr, length)
	reader := bytes.NewReader(data)
	proof, err := readProofData(reader)
	if err != nil {
		fmt.Println("error reading proof data", err)
		return -1
	}
	if debug {
		fmt.Fprintf(os.Stderr, "Proving execution of opcode 0x%x\n", proof.Proof[0])
	}
	checker, err := getProofChecker(int(checkerIndex))
	if err != nil {
		fmt.Println("error getting checker", err)
		return -1
	}
	proofErrors := checker.CheckProof(proof)
	retByte := C.int(0)
	if len(proofErrors) > 0 {
		retByte = 1
		fmt.Fprintln(os.Stderr, "Error verifying proof:")
	}
	for _, err := range proofErrors {
		fmt.Fprintln(os.Stderr, err)
	}
	return retByte
}

func readStateData(reader io.Reader) (*core.ExecutionState, error) {
	buf := make([]byte, 32*5)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}
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
	return data, nil
}

func readBytes(reader io.Reader) ([]byte, error) {
	buf := make([]byte, 8)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}
	buf = make([]byte, int(binary.BigEndian.Uint64(buf)))
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func readProofData(reader io.Reader) (*proofmachine.ProofData, error) {
	before, err := readStateData(reader)
	if err != nil {
		return nil, err
	}
	proof, err := readBytes(reader)
	if err != nil {
		return nil, err
	}
	bufferProof, err := readBytes(reader)
	if err != nil {
		return nil, err
	}
	after, err := readStateData(reader)
	if err != nil {
		return nil, err
	}
	return &proofmachine.ProofData{
		Assertion: &core.Assertion{
			Before: before,
			After:  after,
		},
		Proof:       proof,
		BufferProof: bufferProof,
	}, nil
}

func main() {
}
