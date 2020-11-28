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
	errors2 "github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
)

var messageDeliveredID ethcommon.Hash
var messageDeliveredFromOriginID ethcommon.Hash
var l2MessageFromOriginCallABI abi.Method

func init() {
	inboxABI, err := abi.JSON(strings.NewReader(ethbridgecontracts.GlobalInboxABI))
	if err != nil {
		panic(err)
	}
	messageDeliveredID = inboxABI.Events["MessageDelivered"].ID
	messageDeliveredFromOriginID = inboxABI.Events["MessageDeliveredFromOrigin"].ID
	l2MessageFromOriginCallABI = inboxABI.Methods["sendL2MessageFromOrigin"]
}

type globalInboxWatcher struct {
	GlobalInbox *ethbridgecontracts.GlobalInbox

	rollupAddress ethcommon.Address
	inboxAddress  ethcommon.Address
	client        ethutils.EthClient
}

func newGlobalInboxWatcher(
	globalInboxAddress ethcommon.Address,
	rollupAddress ethcommon.Address,
	client ethutils.EthClient,
) (*globalInboxWatcher, error) {
	globalInboxContract, err := ethbridgecontracts.NewGlobalInbox(
		globalInboxAddress,
		client,
	)
	if err != nil {
		return nil, errors2.WithStack(errors2.Wrap(err, "failed to connect to inbox"))
	}

	return &globalInboxWatcher{
		GlobalInbox:   globalInboxContract,
		rollupAddress: rollupAddress,
		inboxAddress:  globalInboxAddress,
		client:        client,
	}, nil
}

func (gi *globalInboxWatcher) getLogs(
	ctx context.Context,
	fromBlock, toBlock *big.Int,
	blockHash *ethcommon.Hash,
) ([]types.Log, error) {
	addressIndex := ethcommon.Hash{}
	copy(
		addressIndex[:],
		ethcommon.LeftPadBytes(gi.rollupAddress.Bytes(), 32),
	)
	return gi.client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		BlockHash: blockHash,
		Addresses: []ethcommon.Address{gi.inboxAddress},
		Topics: [][]ethcommon.Hash{
			{
				messageDeliveredID,
				messageDeliveredFromOriginID,
			}, {
				addressIndex,
			},
		},
	})
}

func (gi *globalInboxWatcher) GetDeliveredEvents(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]arbbridge.MessageDeliveredEvent, error) {
	inboxLogs, err := gi.getLogs(ctx, fromBlock, toBlock, nil)
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.MessageDeliveredEvent, 0, len(inboxLogs))
	for _, evmLog := range inboxLogs {
		blockHeader, err := gi.client.HeaderByHash(ctx, evmLog.BlockHash)
		if err != nil {
			return nil, err
		}
		timestamp := new(big.Int).SetUint64(blockHeader.Time)
		ev, err := gi.processLog(ctx, evmLog, timestamp)
		if err != nil {
			return nil, err
		}
		events = append(events, ev)
	}
	return events, nil
}

func (gi *globalInboxWatcher) GetEvents(
	ctx context.Context,
	blockId *common.BlockId,
	timestamp *big.Int,
) ([]arbbridge.Event, error) {
	evs, err := gi.GetDeliveredEventsInBlock(ctx, blockId, timestamp)
	if err != nil {
		return nil, err
	}
	events := make([]arbbridge.Event, 0, len(evs))
	for _, ev := range evs {
		events = append(events, ev)
	}
	return events, nil
}

func (gi *globalInboxWatcher) GetDeliveredEventsInBlock(
	ctx context.Context,
	blockId *common.BlockId,
	timestamp *big.Int,
) ([]arbbridge.MessageDeliveredEvent, error) {
	bh := blockId.HeaderHash.ToEthHash()
	inboxLogs, err := gi.getLogs(ctx, nil, nil, &bh)
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.MessageDeliveredEvent, 0, len(inboxLogs))
	for _, evmLog := range inboxLogs {
		ev, err := gi.processLog(ctx, evmLog, timestamp)
		if err != nil {
			return nil, err
		}
		events = append(events, ev)
	}
	return events, nil
}

func (gi *globalInboxWatcher) parseMessageFromOrigin(evmLog types.Log, timestamp *big.Int, msgData []byte) (arbbridge.MessageDeliveredEvent, error) {
	chainTime := inbox.ChainTime{
		BlockNum: common.NewTimeBlocks(
			new(big.Int).SetUint64(evmLog.BlockNumber),
		),
		Timestamp: timestamp,
	}
	val, err := gi.GlobalInbox.ParseMessageDeliveredFromOrigin(evmLog)
	if err != nil {
		return arbbridge.MessageDeliveredEvent{}, err
	}
	return arbbridge.MessageDeliveredEvent{
		ChainInfo: getLogChainInfo(evmLog),
		Message: inbox.InboxMessage{
			Kind:        inbox.Type(val.Kind),
			Sender:      common.NewAddressFromEth(val.Sender),
			InboxSeqNum: val.InboxSeqNum,
			Data:        msgData,
			ChainTime:   chainTime,
		},
	}, nil
}

func (gi *globalInboxWatcher) processLog(
	ctx context.Context,
	evmLog types.Log,
	timestamp *big.Int,
) (arbbridge.MessageDeliveredEvent, error) {
	chainInfo := getLogChainInfo(evmLog)
	chainTime := inbox.ChainTime{
		BlockNum: common.NewTimeBlocks(
			new(big.Int).SetUint64(evmLog.BlockNumber),
		),
		Timestamp: timestamp,
	}
	switch evmLog.Topics[0] {
	case messageDeliveredID:
		val, err := gi.GlobalInbox.ParseMessageDelivered(evmLog)
		if err != nil {
			return arbbridge.MessageDeliveredEvent{}, err
		}
		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message: inbox.InboxMessage{
				Kind:        inbox.Type(val.Kind),
				Sender:      common.NewAddressFromEth(val.Sender),
				InboxSeqNum: val.InboxSeqNum,
				Data:        val.Data,
				ChainTime:   chainTime,
			},
		}, nil

	case messageDeliveredFromOriginID:
		tx, _, err := gi.client.TransactionByHash(ctx, evmLog.TxHash)
		if err != nil {
			return arbbridge.MessageDeliveredEvent{}, err
		}
		args := make(map[string]interface{})
		err = l2MessageFromOriginCallABI.Inputs.UnpackIntoMap(args, tx.Data()[4:])
		if err != nil {
			return arbbridge.MessageDeliveredEvent{}, err
		}
		return gi.parseMessageFromOrigin(evmLog, timestamp, args["messageData"].([]byte))
	default:
		return arbbridge.MessageDeliveredEvent{}, errors2.New("unknown arbitrum event type")
	}
}

func (gi *globalInboxWatcher) GetERC20Balance(
	ctx context.Context,
	user common.Address,
	tokenContract common.Address,
) (*big.Int, error) {
	return gi.GlobalInbox.GetERC20Balance(
		&bind.CallOpts{Context: ctx},
		tokenContract.ToEthAddress(),
		user.ToEthAddress(),
	)
}

func (gi *globalInboxWatcher) GetEthBalance(
	ctx context.Context,
	user common.Address,
) (*big.Int, error) {
	return gi.GlobalInbox.GetEthBalance(
		&bind.CallOpts{Context: ctx},
		user.ToEthAddress(),
	)
}
