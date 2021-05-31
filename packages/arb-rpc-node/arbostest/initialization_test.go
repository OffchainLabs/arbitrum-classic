/*
* Copyright 2021, Offchain Labs, Inc.
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
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestMemoryInitialization(t *testing.T) {
	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	failIfError(t, err)

	senderKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	sender := common.NewAddressFromEth(crypto.PubkeyToAddress(senderKey.PublicKey))

	tx1 := types.NewTx(&types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(0),
		Gas:      1000000000,
		Value:    big.NewInt(0),
		Data:     hexutil.MustDecode(arbostestcontracts.SimpleBin),
	})
	tx1, err = types.SignTx(tx1, types.HomesteadSigner{}, senderKey)
	test.FailIfError(t, err)
	msg1, err := message.NewL2Message(message.NewCompressedECDSAFromEth(tx1))
	test.FailIfError(t, err)

	to := crypto.CreateAddress(crypto.PubkeyToAddress(senderKey.PublicKey), 0)
	tx2 := types.NewTx(&types.LegacyTx{
		Nonce:    1,
		GasPrice: big.NewInt(0),
		Gas:      10000000,
		To:       &to,
		Value:    big.NewInt(0),
		Data:     makeFuncData(t, simpleABI.Methods["debug"]),
	})
	tx2, err = types.SignTx(tx2, types.NewEIP155Signer(chainId), senderKey)
	test.FailIfError(t, err)
	msg2, err := message.NewL2Message(message.NewCompressedECDSAFromEth(tx2))
	test.FailIfError(t, err)

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	ib := &InboxBuilder{}
	options := []message.ChainConfigOption{message.ChainIDConfig{ChainId: chainId}}
	ib.AddMessage(initMsg(t, options), chain, big.NewInt(0), chainTime)
	ib.AddMessage(msg1, sender, big.NewInt(0), chainTime)
	ib.AddMessage(msg2, sender, big.NewInt(0), chainTime)

	results, _ := runTxAssertion(t, ib.Messages)
	checkConstructorResult(t, results[0], common.NewAddressFromEth(to))
	succeededTxCheck(t, results[1])

	if len(results[1].EVMLogs) != 1 {
		t.Fatal("unexpected log count")
	}
	evmLog := results[1].EVMLogs[0]
	t.Log("log data", hexutil.Encode(evmLog.Data))
	for _, byt := range evmLog.Data {
		if byt != 0 {
			t.Fatal("found nonzero byte")
		}
	}
}
