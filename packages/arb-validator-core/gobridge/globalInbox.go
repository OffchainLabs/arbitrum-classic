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

package gobridge

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	//"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type globalInbox struct {
	inbox   *inbox
	client  *GoArbAuthClient
	address common.Address
}

func newGlobalInbox(address common.Address, client *GoArbAuthClient) (*globalInbox, error) {
	client.GoEthClient.inbox[address] = &inbox{common.Hash{}, big.NewInt(0)}
	return &globalInbox{
		inbox:   client.GoEthClient.inbox[address],
		client:  client,
		address: address,
	}, nil
}

func (con *globalInbox) SendTransactionMessage(ctx context.Context, data []byte, vmAddress common.Address, contactAddress common.Address, amount *big.Int, seqNumber *big.Int) error {
	msgHash := hashing.SoliditySHA3(
		hashing.Uint8(0), // TRANSACTION_MSG
		hashing.Address(vmAddress),
		hashing.Address(contactAddress),
		hashing.Uint256(seqNumber),
		hashing.Uint256(amount),
		data,
		hashing.TimeBlocks(con.client.GoEthClient.LastMinedBlock.Height),
	)

	con.client.GoEthClient.deliverMessage(vmAddress, msgHash)
	msg := message.DeliveredTransaction{
		Transaction: message.Transaction{
			Chain:       con.address,
			To:          contactAddress,
			From:        vmAddress,
			SequenceNum: seqNumber,
			Value:       amount,
			Data:        data,
		},
		BlockNum: con.client.GoEthClient.LastMinedBlock.Height,
	}

	con.client.GoEthClient.pubMsg(nil, arbbridge.MaybeEvent{
		Event: arbbridge.MessageDeliveredEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: con.client.GoEthClient.getCurrentBlock(),
			},
			Message: msg,
		},
	})

	return nil
}

func (con *globalInbox) DeliverTransactionBatch(
	ctx context.Context,
	chain common.Address,
	transactions []message.Transaction,
	signatures [][65]byte,
) error {
	tos := make([]ethcommon.Address, 0, len(transactions))
	seqNums := make([]*big.Int, 0, len(transactions))
	amounts := make([]*big.Int, 0, len(transactions))
	messageLengths := make([]*big.Int, 0, len(transactions))
	data := make([]byte, 0)
	signaturesFlat := make([]byte, 0, len(transactions)*65)
	for i, tx := range transactions {
		tos = append(tos, tx.To.ToEthAddress())
		seqNums = append(seqNums, tx.SequenceNum)
		amounts = append(amounts, tx.Value)
		messageLengths = append(messageLengths, big.NewInt(int64(len(tx.Data))))
		data = append(data, tx.Data...)
		signaturesFlat = append(signaturesFlat, signatures[i][:]...)
		msgHash := transactionHash(
			chain,
			tx.To,
			tx.From,
			tx.SequenceNum,
			tx.Value,
			tx.Data,
			con.client.GoEthClient.getCurrentBlock().Height,
		)

		con.inbox.addMessageToInbox(msgHash)

		con.client.GoEthClient.pubMsg(nil, arbbridge.MaybeEvent{
			Event: arbbridge.MessageDeliveredEvent{
				ChainInfo: arbbridge.ChainInfo{
					BlockId: con.client.GoEthClient.getCurrentBlock(),
				},
				Message: message.DeliveredTransaction{
					Transaction: message.Transaction{
						Chain:       chain,
						To:          tx.To,
						From:        tx.From,
						SequenceNum: tx.SequenceNum,
						Value:       tx.Value,
						Data:        tx.Data,
					},
					BlockNum: con.client.GoEthClient.getCurrentBlock().Height,
				},
			},
		})
	}
	return nil
}

func (con *globalInbox) DepositEthMessage(
	ctx context.Context,
	vmAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	// depositEth
	if _, ok := con.client.GoEthClient.ethWallet[vmAddress]; ok {
		con.client.GoEthClient.ethWallet[vmAddress] = new(big.Int).Add(con.client.GoEthClient.ethWallet[vmAddress], value)
	}
	//deliverEthMessage
	msgNum := new(big.Int).Add(con.inbox.count, big.NewInt(1))

	msg := message.DeliveredEth{
		Eth: message.Eth{
			To:    destination,
			From:  vmAddress,
			Value: value,
		},
		BlockNum:   con.client.GoEthClient.LastMinedBlock.Height,
		MessageNum: msgNum,
	}
	msgHash := msg.CommitmentHash()

	con.inbox.addMessageToInbox(msgHash)

	con.client.GoEthClient.pubMsg(nil, arbbridge.MaybeEvent{
		Event: arbbridge.MessageDeliveredEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: con.client.GoEthClient.getCurrentBlock(),
			},
			Message: msg,
		},
	})
	return nil
}

func (con *globalInbox) DepositERC20Message(
	ctx context.Context,
	vmAddress common.Address,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	msgNum := new(big.Int).Add(con.inbox.count, big.NewInt(1))
	msg := message.DeliveredERC20{
		ERC20: message.ERC20{
			To:           destination,
			From:         vmAddress,
			TokenAddress: tokenAddress,
			Value:        value,
		},
		BlockNum:   con.client.GoEthClient.getCurrentBlock().Height,
		MessageNum: msgNum,
	}
	msgHash := msg.CommitmentHash()
	con.inbox.addMessageToInbox(msgHash)

	con.client.GoEthClient.pubMsg(nil, arbbridge.MaybeEvent{
		Event: arbbridge.MessageDeliveredEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: con.client.GoEthClient.getCurrentBlock(),
			},
			Message: msg,
		},
	})
	return nil
}

func (con *globalInbox) DepositERC721Message(
	ctx context.Context,
	vmAddress common.Address,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	msgNum := new(big.Int).Add(con.inbox.count, big.NewInt(1))
	msg := message.DeliveredERC721{
		ERC721: message.ERC721{
			To:           tokenAddress,
			From:         vmAddress,
			TokenAddress: destination,
			Id:           value,
		},
		BlockNum:   con.client.GoEthClient.getCurrentBlock().Height,
		MessageNum: msgNum,
	}
	msgHash := msg.CommitmentHash()
	con.inbox.addMessageToInbox(msgHash)

	con.client.GoEthClient.pubMsg(nil, arbbridge.MaybeEvent{
		Event: arbbridge.MessageDeliveredEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: con.client.GoEthClient.getCurrentBlock(),
			},
			Message: msg,
		},
	})
	return nil
}

func (con *globalInbox) GetTokenBalance(
	ctx context.Context,
	user common.Address,
	tokenContract common.Address,
) (*big.Int, error) {

	// TODO: fill in
	//return con.GlobalPendingInbox.GetTokenBalance(
	//	auth,
	//	tokenContract,
	//	user,
	//)
	return big.NewInt(0), nil
}
