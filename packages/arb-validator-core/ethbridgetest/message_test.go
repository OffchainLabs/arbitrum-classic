/*
 * Copyright 2019, Offchain Labs, Inc.
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

package ethbridgetest

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/messagetester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
)

func init() {

}

func TestMessageHashing(t *testing.T) {
	auth, err := test.SetupAuth("27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa")
	if err != nil {
		t.Fatal(err)
	}
	client, err := ethclient.Dial(test.GetEthUrl())
	if err != nil {
		t.Fatal(err)
	}
	_, tx, tester, err := messagetester.DeployMessageTester(auth, client)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(context.Background(), client, auth.From, tx, "DeployMessageTester")
	if err != nil {
		t.Fatal(err)
	}

	addr1 := common.Address{}
	addr1[0] = 76
	addr1[19] = 93

	addr2 := common.Address{}
	addr2[0] = 43
	addr2[19] = 12

	addr3 := common.Address{}
	addr3[0] = 73
	addr3[19] = 85

	t.Run("Transaction Message", func(t *testing.T) {
		msg := message.DeliveredTransaction{
			Transaction: message.Transaction{
				Chain:       addr3,
				To:          addr1,
				From:        addr2,
				SequenceNum: big.NewInt(74563),
				Value:       big.NewInt(89735406),
				Data:        []byte{65, 23, 68, 87, 12},
			},
			BlockNum: common.NewTimeBlocks(big.NewInt(87962345)),
		}
		bridgeHash, err := tester.TransactionHash(
			nil,
			msg.Chain.ToEthAddress(),
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.SequenceNum,
			msg.Value,
			msg.Data,
			msg.BlockNum.AsInt(),
		)
		if err != nil {
			t.Fatal(err)
		}
		if bridgeHash != msg.CommitmentHash().ToEthHash() {
			t.Error("Ethbridge calculated wrong hash")
		}

		messageBridgeHash, err := tester.TransactionMessageHash(
			nil,
			msg.Chain.ToEthAddress(),
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.SequenceNum,
			msg.Value,
			msg.Data,
			msg.BlockNum.AsInt(),
		)
		if err != nil {
			t.Fatal(err)
		}

		if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
			t.Error("Ethbridge calculated wrong message hash")
		}
	})

	t.Run("Eth Message", func(t *testing.T) {
		msg := message.DeliveredEth{
			Eth: message.Eth{
				To:    addr1,
				From:  addr2,
				Value: big.NewInt(89735406),
			},
			BlockNum:   common.NewTimeBlocks(big.NewInt(87962345)),
			MessageNum: big.NewInt(98742),
		}
		bridgeHash, err := tester.EthHash(
			nil,
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.Value,
			msg.BlockNum.AsInt(),
			msg.MessageNum,
		)
		if err != nil {
			t.Fatal(err)
		}
		if bridgeHash != msg.CommitmentHash().ToEthHash() {
			t.Error("Ethbridge calculated wrong hash")
		}

		messageBridgeHash, err := tester.EthMessageHash(
			nil,
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.Value,
			msg.BlockNum.AsInt(),
			msg.MessageNum,
		)
		if err != nil {
			t.Fatal(err)
		}

		// if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
		// 	t.Error("Ethbridge calculated wrong message hash")
		// }
	})

	t.Run("ERC20 Message", func(t *testing.T) {
		msg := message.DeliveredERC20{
			ERC20: message.ERC20{
				To:           addr1,
				From:         addr2,
				TokenAddress: addr3,
				Value:        big.NewInt(89735406),
			},
			BlockNum:   common.NewTimeBlocks(big.NewInt(87962345)),
			MessageNum: big.NewInt(98742),
		}
		bridgeHash, err := tester.Erc20Hash(
			nil,
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.TokenAddress.ToEthAddress(),
			msg.Value,
			msg.BlockNum.AsInt(),
			msg.MessageNum,
		)
		if err != nil {
			t.Fatal(err)
		}
		if bridgeHash != msg.CommitmentHash().ToEthHash() {
			t.Error("Ethbridge calculated wrong hash")
		}

		messageBridgeHash, err := tester.Erc20MessageHash(
			nil,
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.TokenAddress.ToEthAddress(),
			msg.Value,
			msg.BlockNum.AsInt(),
			msg.MessageNum,
		)
		if err != nil {
			t.Fatal(err)
		}

		if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
			t.Error("Ethbridge calculated wrong message hash")
		}
	})

	t.Run("ERC721 Message", func(t *testing.T) {
		msg := message.DeliveredERC721{
			ERC721: message.ERC721{
				To:           addr1,
				From:         addr2,
				TokenAddress: addr3,
				Id:           big.NewInt(89735406),
			},
			BlockNum:   common.NewTimeBlocks(big.NewInt(87962345)),
			MessageNum: big.NewInt(98742),
		}
		bridgeHash, err := tester.Erc721Hash(
			nil,
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.TokenAddress.ToEthAddress(),
			msg.Id,
			msg.BlockNum.AsInt(),
			msg.MessageNum,
		)
		if err != nil {
			t.Fatal(err)
		}
		if bridgeHash != msg.CommitmentHash().ToEthHash() {
			t.Error("Ethbridge calculated wrong hash")
		}

		messageBridgeHash, err := tester.Erc721MessageHash(
			nil,
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.TokenAddress.ToEthAddress(),
			msg.Id,
			msg.BlockNum.AsInt(),
			msg.MessageNum,
		)
		if err != nil {
			t.Fatal(err)
		}

		if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
			t.Error("Ethbridge calculated wrong message hash")
		}
	})

}
