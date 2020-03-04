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

package gobridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type messageTester struct {
	tester common.Address
	client *GoArbAuthClient
}

func NewMessageTester(client *GoArbAuthClient) (*messageTester, error) {
	tester := client.getNextAddress()

	return &messageTester{client: client, tester: tester}, nil
}

func (m *messageTester) TransactionHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	msgHash := transactionHash(
		chain,
		to,
		from,
		seqNumber,
		value,
		data,
		common.NewTimeBlocks(blockNumber),
	)
	return msgHash, nil
}

func (m *messageTester) TransactionMessageHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte, blockNumber *big.Int) ([32]byte, error) {
	msg := message.DeliveredTransaction{
		Transaction: message.Transaction{
			Chain:       chain,
			To:          to,
			From:        from,
			SequenceNum: seqNumber,
			Value:       value,
			Data:        data,
		},
		BlockNum: common.NewTimeBlocks(blockNumber),
	}
	return message.DeliveredValue(msg).Hash().ToEthHash(), nil
}

func (m *messageTester) EthHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	msg := message.DeliveredEth{
		Eth: message.Eth{
			To:    to,
			From:  from,
			Value: value,
		},
		BlockNum:   common.NewTimeBlocks(blockNumber),
		MessageNum: messageNum,
	}
	return msg.CommitmentHash().ToEthHash(), nil
}

func (m *messageTester) EthMessageHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	msg := message.DeliveredEth{
		Eth: message.Eth{
			To:    to,
			From:  from,
			Value: value,
		},
		BlockNum:   common.NewTimeBlocks(blockNumber),
		MessageNum: messageNum,
	}
	return message.DeliveredValue(msg).Hash().ToEthHash(), nil
}

func (m *messageTester) Erc20Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	msg := message.DeliveredERC20{
		ERC20: message.ERC20{
			To:           to,
			From:         from,
			TokenAddress: erc20,
			Value:        value,
		},
		BlockNum:   common.NewTimeBlocks(blockNumber),
		MessageNum: messageNum,
	}
	return msg.CommitmentHash().ToEthHash(), nil
}

func (m *messageTester) Erc20MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	msg := message.DeliveredERC20{
		ERC20: message.ERC20{
			To:           to,
			From:         from,
			TokenAddress: erc20,
			Value:        value,
		},
		BlockNum:   common.NewTimeBlocks(blockNumber),
		MessageNum: messageNum,
	}
	return message.DeliveredValue(msg).Hash().ToEthHash(), nil
}

func (m *messageTester) Erc721Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	msg := message.DeliveredERC721{
		ERC721: message.ERC721{
			To:           to,
			From:         from,
			TokenAddress: erc721,
			Id:           id,
		},
		BlockNum:   common.NewTimeBlocks(blockNumber),
		MessageNum: messageNum,
	}
	return msg.CommitmentHash().ToEthHash(), nil
}

func (m *messageTester) Erc721MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int, blockNumber *big.Int, messageNum *big.Int) ([32]byte, error) {
	msg := message.DeliveredERC721{
		ERC721: message.ERC721{
			To:           to,
			From:         from,
			TokenAddress: erc721,
			Id:           id,
		},
		BlockNum:   common.NewTimeBlocks(blockNumber),
		MessageNum: messageNum,
	}
	return message.DeliveredValue(msg).Hash().ToEthHash(), nil
}
