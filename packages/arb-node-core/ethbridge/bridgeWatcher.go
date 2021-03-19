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

package ethbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/monitor"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
)

var bridgeABI abi.ABI
var messageDeliveredID ethcommon.Hash
var inboxMessageDeliveredID ethcommon.Hash
var inboxMessageFromOriginID ethcommon.Hash

func init() {
	parsedBridgeABI, err := abi.JSON(strings.NewReader(ethbridgecontracts.BridgeABI))
	if err != nil {
		panic(err)
	}
	messageDeliveredID = parsedBridgeABI.Events["MessageDelivered"].ID
	bridgeABI = parsedBridgeABI

	parsedInboxABI, err := abi.JSON(strings.NewReader(ethbridgecontracts.InboxABI))
	if err != nil {
		panic(err)
	}
	inboxMessageDeliveredID = parsedInboxABI.Events["InboxMessageDelivered"].ID
	inboxMessageFromOriginID = parsedInboxABI.Events["InboxMessageDeliveredFromOrigin"].ID
}

type InboxMessageGetter interface {
	fillMessageDetails(ctx context.Context, messageNums []*big.Int, txData map[string]*types.Transaction, messages map[string][]byte) error
}

type BridgeWatcher struct {
	con     *ethbridgecontracts.Bridge
	address ethcommon.Address
	client  ethutils.EthClient

	inboxes map[ethcommon.Address]InboxMessageGetter
}

func NewBridgeWatcher(address ethcommon.Address, client ethutils.EthClient) (*BridgeWatcher, error) {
	con, err := ethbridgecontracts.NewBridge(address, client)
	if err != nil {
		return nil, err
	}

	return &BridgeWatcher{
		con:     con,
		address: address,
		client:  client,
		inboxes: make(map[ethcommon.Address]InboxMessageGetter),
	}, nil
}

func (r *BridgeWatcher) CurrentBlockHeight(ctx context.Context) (*big.Int, error) {
	latestHeader, err := r.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return latestHeader.Number, nil
}

func (r *BridgeWatcher) LookupMessagesInRange(ctx context.Context, from, to *big.Int) ([]*DeliveredInboxMessage, error) {
	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: from,
		ToBlock:   to,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{messageDeliveredID}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	for _, evmLog := range logs {
		monitor.GlobalMonitor.ReaderGotBatch(common.NewHashFromEth(evmLog.TxHash))
	}
	return r.logsToDeliveredMessages(ctx, logs)
}

func (r *BridgeWatcher) LookupMessageBlock(ctx context.Context, messageNum *big.Int) (*common.BlockId, error) {
	var msgNumBytes ethcommon.Hash
	copy(msgNumBytes[:], math.U256Bytes(messageNum))

	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{messageDeliveredID}, {msgNumBytes}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, errors.New("log not found")
	}
	if len(logs) > 1 {
		return nil, errors.New("too many logs")
	}
	ethLog := logs[0]
	return &common.BlockId{
		Height:     common.NewTimeBlocksInt(int64(ethLog.BlockNumber)),
		HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
	}, nil
}

type DeliveredInboxMessageList []*DeliveredInboxMessage

func (d DeliveredInboxMessageList) Len() int {
	return len(d)
}

func (d DeliveredInboxMessageList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DeliveredInboxMessageList) Less(i, j int) bool {
	return d[i].Message.InboxSeqNum.Cmp(d[j].Message.InboxSeqNum) < 0
}

func (r *BridgeWatcher) logsToDeliveredMessages(ctx context.Context, logs []types.Log) ([]*DeliveredInboxMessage, error) {
	messagesByInbox := make(map[ethcommon.Address][]*big.Int)
	rawMessages := make(map[string]*ethbridgecontracts.BridgeMessageDelivered)
	rawTransactions := make(map[string]*types.Transaction)
	for _, ethLog := range logs {
		parsedLog, err := r.con.ParseMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}
		messagesByInbox[parsedLog.Inbox] = append(messagesByInbox[parsedLog.Inbox], parsedLog.MessageIndex)
		messageKey := string(parsedLog.MessageIndex.Bytes())
		rawMessages[messageKey] = parsedLog

		txData, err := r.client.TransactionInBlock(ctx, ethLog.BlockHash, ethLog.TxIndex)
		if err != nil {
			return nil, err
		}
		rawTransactions[messageKey] = txData
	}

	messageData := make(map[string][]byte)
	for con, indexes := range messagesByInbox {
		inboxGetter, err := r.getInboxGetter(con)
		if err != nil {
			return nil, err
		}
		if err := inboxGetter.fillMessageDetails(ctx, indexes, rawTransactions, messageData); err != nil {
			return nil, err
		}
	}

	blockTimes := make(map[ethcommon.Hash]*big.Int)

	messages := make([]*DeliveredInboxMessage, 0, len(logs))
	for msgNum, rawMsg := range rawMessages {
		data, ok := messageData[msgNum]
		if !ok {
			return nil, errors.New("message not found")
		}
		if hashing.SoliditySHA3(data) != rawMsg.MessageDataHash {
			return nil, errors.New("found message data with mismatched hash")
		}

		blockTime, ok := blockTimes[rawMsg.Raw.BlockHash]
		if !ok {
			header, err := r.client.HeaderByHash(ctx, rawMsg.Raw.BlockHash)
			if err != nil {
				return nil, err
			}
			blockTime = new(big.Int).SetUint64(header.Time)
			blockTimes[rawMsg.Raw.BlockHash] = blockTime
		}

		msg := &DeliveredInboxMessage{
			BlockHash:      common.NewHashFromEth(rawMsg.Raw.BlockHash),
			BeforeInboxAcc: rawMsg.BeforeInboxAcc,
			Message: inbox.InboxMessage{
				Kind:        inbox.Type(rawMsg.Kind),
				Sender:      common.NewAddressFromEth(rawMsg.Sender),
				InboxSeqNum: rawMsg.MessageIndex,
				GasPrice:    rawTransactions[string(rawMsg.MessageIndex.Bytes())].GasPrice(),
				Data:        data,
				ChainTime: inbox.ChainTime{
					BlockNum: common.NewTimeBlocks(
						new(big.Int).SetUint64(rawMsg.Raw.BlockNumber),
					),
					Timestamp: blockTime,
				},
			},
		}
		messages = append(messages, msg)
	}

	sort.Sort(DeliveredInboxMessageList(messages))

	return messages, nil
}

func (r *BridgeWatcher) getInboxGetter(inboxAddress ethcommon.Address) (InboxMessageGetter, error) {
	curInbox, ok := r.inboxes[inboxAddress]
	if ok {
		return curInbox, nil
	}

	curInbox, err := NewStandardInboxWatcher(inboxAddress, r.client)
	if err != nil {
		return nil, err
	}
	r.inboxes[inboxAddress] = curInbox
	return curInbox, nil
}
