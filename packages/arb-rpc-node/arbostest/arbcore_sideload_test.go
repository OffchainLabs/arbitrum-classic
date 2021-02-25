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

package arbostest

import (
	"bytes"
	"context"
	"crypto"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestArbCoreSideload(t *testing.T) {
	var precompileNum byte = 2
	data := common.RandBytes(100)
	sha256 := crypto.SHA256.New()
	_, err := sha256.Write(data)
	failIfError(t, err)
	correct := sha256.Sum(nil)

	tmpDir, err := ioutil.TempDir("", "arbitrum")
	failIfError(t, err)
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()
	storage, err := cmachine.NewArbStorage(tmpDir)
	test.FailIfError(t, err)
	defer storage.CloseArbStorage()

	err = storage.Initialize(arbos.Path())
	test.FailIfError(t, err)

	arbCore := storage.GetArbCore()
	started := arbCore.StartThread()
	if !started {
		t.Fatal("failed to start thread")
	}

	messages := []inbox.InboxMessage{
		message.NewInboxMessage(
			initMsg(),
			chain,
			big.NewInt(0),
			inbox.ChainTime{
				BlockNum:  common.NewTimeBlocksInt(0),
				Timestamp: big.NewInt(0),
			},
		),
	}
	_, err = core.DeliverMessagesAndWait(arbCore, messages, common.Hash{}, true)
	test.FailIfError(t, err)
	for {
		if arbCore.MachineIdle() {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	// ArbCore should backtrack and look for the latest previous sideload
	sideloadMachine, err := arbCore.GetMachineForSideload(5)
	test.FailIfError(t, err)

	precompileAddress := ethcommon.BytesToAddress([]byte{precompileNum})

	ctx := context.Background()
	backend, _ := test.SimulatedBackend()

	ethCall := ethereum.CallMsg{
		From:     ethcommon.Address{},
		To:       &precompileAddress,
		Gas:      1000000,
		GasPrice: big.NewInt(0),
		Value:    big.NewInt(0),
		Data:     data,
	}
	ethRes, err := backend.CallContract(ctx, ethCall, nil)
	failIfError(t, err)

	tx := message.Transaction{
		MaxGas:      big.NewInt(100000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(precompileAddress),
		Payment:     big.NewInt(0),
		Data:        data,
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	sideloadMessages := []inbox.InboxMessage{
		message.NewInboxMessage(message.NewSafeL2Message(tx), sender, big.NewInt(1), chainTime),
	}

	assertion, _, _ := sideloadMachine.ExecuteAssertionAdvanced(0, false, nil, false, sideloadMessages, true, common.Hash{}, common.Hash{})
	results := processTxResults(t, assertion.Logs)

	res := results[0]
	succeededTxCheck(t, res)
	if res.IncomingRequest.Kind != message.L2Type {
		t.Fatal("wrong request type")
	}
	_, err = message.L2Message{Data: res.IncomingRequest.Data}.AbstractMessage()
	failIfError(t, err)

	if !bytes.Equal(res.ReturnData, correct) {
		t.Logf("Got result 0x%x", res.ReturnData)
		t.Logf("Wanted result 0x%x", correct)
		t.Error("calculated result incorrectly")
	}

	if !bytes.Equal(ethRes, correct) {
		t.Logf("Got result 0x%x", res.ReturnData)
		t.Logf("Wanted result 0x%x", correct)
		t.Error("calculated result incorrectly")
	}
}
