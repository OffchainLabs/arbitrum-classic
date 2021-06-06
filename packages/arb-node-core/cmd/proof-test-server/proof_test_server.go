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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/pkg/errors"
)

type stateData struct {
	gasUsed           *big.Int
	totalMessagesRead *big.Int
	machineHash       common.Hash
	sendAcc           common.Hash
	logAcc            common.Hash
}

type proofData struct {
	before      stateData
	proof       []byte
	bufferProof []byte
	after       stateData
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

func checkProof(proof proofData, osps []*ethbridgetestcontracts.IOneStepProof, sequencerBridge, delayedBridge ethcommon.Address, debug bool) error {
	op := proof.proof[0]
	if debug {
		fmt.Fprintf(os.Stderr, "Proving execution of opcode 0x%x\n", op)
	}
	prover := getProverNum(op)
	machineData, err := osps[prover].ExecuteStep(
		&bind.CallOpts{},
		[2]ethcommon.Address{sequencerBridge, delayedBridge},
		proof.before.totalMessagesRead,
		[2][32]byte{
			proof.before.sendAcc,
			proof.before.logAcc,
		},
		proof.proof,
		proof.bufferProof,
	)
	if err != nil {
		return errors.Wrap(err, "Solidity OSP execution failed")
	}
	correctGasUsed := new(big.Int).Sub(proof.after.gasUsed, proof.before.gasUsed)
	if new(big.Int).SetUint64(machineData.Gas).Cmp(correctGasUsed) != 0 {
		return errors.Errorf("wrong gas %v instead of %v", machineData.Gas, correctGasUsed)
	}
	if machineData.AfterMessagesRead.Cmp(proof.after.totalMessagesRead) != 0 {
		return errors.Errorf("wrong total messages read")
	}
	if machineData.Fields[0] != proof.before.machineHash {
		return errors.Errorf("wrong before machine")
	}
	if machineData.Fields[2] != proof.after.sendAcc {
		return errors.Errorf("wrong send accumulator")
	}
	if machineData.Fields[3] != proof.after.logAcc {
		return errors.Errorf("wrong log accumulator")
	}
	if machineData.Fields[1] != proof.after.machineHash {
		return errors.Errorf("wrong after machine (got %v but expected %v)", common.Hash(machineData.Fields[1]).String(), proof.after.machineHash.String())
	}
	return nil
}

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

func readStateData(reader io.Reader) stateData {
	buf := make([]byte, 32*5)
	_, err := io.ReadFull(reader, buf)
	handleFatalError(err)
	data := stateData{
		gasUsed:           new(big.Int).SetBytes(buf[:32]),
		totalMessagesRead: new(big.Int).SetBytes(buf[32:64]),
	}
	copy(data.machineHash[:], buf[64:96])
	copy(data.sendAcc[:], buf[96:128])
	copy(data.logAcc[:], buf[128:160])
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

func readProofData(reader io.Reader) proofData {
	before := readStateData(reader)
	proof := readBytes(reader)
	bufferProof := readBytes(reader)
	after := readStateData(reader)
	return proofData{
		before:      before,
		proof:       proof,
		bufferProof: bufferProof,
		after:       after,
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

	backend, pks := test.SimulatedBackend(&testing.T{})
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth, err := bind.NewKeyedTransactorWithChainID(pks[0], big.NewInt(1337))
	handleFatalError(err)
	sequencer := common.RandAddress().ToEthAddress()
	maxDelayBlocks := big.NewInt(60)
	maxDelaySeconds := big.NewInt(900)

	osp1Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	handleFatalError(err)
	osp2Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof2(auth, client)
	handleFatalError(err)
	osp3Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProofHash(auth, client)
	handleFatalError(err)
	delayedBridgeAddr, _, _, err := ethbridgecontracts.DeployBridge(auth, client)
	handleFatalError(err)
	sequencerBridgeAddr, _, sequencerCon, err := ethbridgecontracts.DeploySequencerInbox(auth, client)
	handleFatalError(err)
	rollupAddr, _, rollup, err := ethbridgetestcontracts.DeployRollupMock(auth, client)
	handleFatalError(err)
	client.Commit()

	_, err = rollup.SetMock(auth, maxDelayBlocks, maxDelaySeconds)
	handleFatalError(err)
	_, err = sequencerCon.Initialize(auth, delayedBridgeAddr, sequencer, rollupAddr)
	handleFatalError(err)
	client.Commit()

	osp1, err := ethbridgetestcontracts.NewIOneStepProof(osp1Addr, client)
	handleFatalError(err)
	osp2, err := ethbridgetestcontracts.NewIOneStepProof(osp2Addr, client)
	handleFatalError(err)
	osp3, err := ethbridgetestcontracts.NewIOneStepProof(osp3Addr, client)
	handleFatalError(err)
	provers := []*ethbridgetestcontracts.IOneStepProof{osp1, osp2, osp3}

	for {
		proof := readProofData(queryPipe)
		err = checkProof(proof, provers, sequencerBridgeAddr, delayedBridgeAddr, debug)
		retByte := uint8(0)
		if err != nil {
			retByte = 1
			fmt.Fprintf(os.Stderr, "Error verifying proof: %v\n", err)
		}
		_, err = resultPipe.Write([]byte{retByte})
		handleFatalError(err)
	}
}
