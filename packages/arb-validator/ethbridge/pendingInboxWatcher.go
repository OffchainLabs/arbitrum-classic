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

package ethbridge

import (
	"context"
	"math/big"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var transactionMessageDeliveredID ethcommon.Hash
var ethDepositMessageDeliveredID ethcommon.Hash
var depositERC20MessageDeliveredID ethcommon.Hash
var depositERC721MessageDeliveredID ethcommon.Hash

func init() {
	parsedInbox, err := abi.JSON(strings.NewReader(rollup.IGlobalPendingInboxABI))
	if err != nil {
		panic(err)
	}

	transactionMessageDeliveredID = parsedInbox.Events["TransactionMessageDelivered"].ID()
	ethDepositMessageDeliveredID = parsedInbox.Events["EthDepositMessageDelivered"].ID()
	depositERC20MessageDeliveredID = parsedInbox.Events["ERC20DepositMessageDelivered"].ID()
	depositERC721MessageDeliveredID = parsedInbox.Events["ERC721DepositMessageDelivered"].ID()
}

type pendingInboxWatcher struct {
	GlobalPendingInbox *rollup.IGlobalPendingInbox
	contractAddress    ethcommon.Address
	rollupAddress      ethcommon.Address
	client             *ethclient.Client
}

func newPendingInboxWatcher(address ethcommon.Address, rollupAddress ethcommon.Address, client *ethclient.Client) (*pendingInboxWatcher, error) {
	globalPendingContract, err := rollup.NewIGlobalPendingInbox(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}

	return &pendingInboxWatcher{
		GlobalPendingInbox: globalPendingContract,
		contractAddress:    address,
		rollupAddress:      rollupAddress,
		client:             client,
	}, nil
}

func (rw *pendingInboxWatcher) GetEvents(ctx context.Context, blockID *structures.BlockID) ([]arbbridge.Event, error) {
	bh := blockID.HeaderHash.ToEthHash()
	addressIndex := ethcommon.Hash{}
	copy(addressIndex[:], ethcommon.LeftPadBytes(rw.rollupAddress.Bytes(), 32))
	inboxLogs, err := rw.client.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &bh,
		Addresses: []ethcommon.Address{rw.rollupAddress},
		Topics: [][]ethcommon.Hash{{
			transactionMessageDeliveredID,
			ethDepositMessageDeliveredID,
			depositERC20MessageDeliveredID,
			depositERC721MessageDeliveredID,
			addressIndex,
		}},
	})
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(inboxLogs))
	for _, evmLog := range inboxLogs {
		event, err := rw.processEvents(getLogChainInfo(evmLog), evmLog)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (rw *pendingInboxWatcher) processEvents(chainInfo arbbridge.ChainInfo, ethLog types.Log) (arbbridge.Event, error) {
	switch ethLog.Topics[0] {
	case transactionMessageDeliveredID:
		val, err := rw.GlobalPendingInbox.ParseTransactionMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredTransaction{
			Transaction: message.Transaction{
				Chain:       common.NewAddressFromEth(rw.rollupAddress),
				To:          common.NewAddressFromEth(val.To),
				From:        common.NewAddressFromEth(val.From),
				SequenceNum: val.SeqNumber,
				Value:       val.Value,
				Data:        val.Data,
			},
			BlockNum: common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil
	case ethDepositMessageDeliveredID:
		val, err := rw.GlobalPendingInbox.ParseEthDepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredEth{
			Eth: message.Eth{
				To:    common.NewAddressFromEth(val.To),
				From:  common.NewAddressFromEth(val.From),
				Value: val.Value,
			},
			BlockNum:   common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			MessageNum: val.MessageNum,
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil
	case depositERC20MessageDeliveredID:
		val, err := rw.GlobalPendingInbox.ParseERC20DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}
		msg := message.DeliveredERC20{
			ERC20: message.ERC20{
				To:           common.NewAddressFromEth(val.To),
				From:         common.NewAddressFromEth(val.From),
				TokenAddress: common.NewAddressFromEth(val.Erc20),
				Value:        val.Value,
			},
			BlockNum:   common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			MessageNum: val.MessageNum,
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil
	case depositERC721MessageDeliveredID:
		val, err := rw.GlobalPendingInbox.ParseERC721DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredERC721{
			ERC721: message.ERC721{
				To:           common.NewAddressFromEth(val.To),
				From:         common.NewAddressFromEth(val.From),
				TokenAddress: common.NewAddressFromEth(val.Erc721),
				ID:           val.Id,
			},
			BlockNum:   common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			MessageNum: val.MessageNum,
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil
	default:
		return nil, errors2.New("unknown arbitrum event type")
	}
}
