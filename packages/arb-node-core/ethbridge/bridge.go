package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"math/big"
	"strings"
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
	fillMessageDetails(ctx context.Context, messageNums []*big.Int, messages map[string]inbox.InboxMessage) error
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
	}, nil
}

func (r *BridgeWatcher) LookupMessageBlock(ctx context.Context, messageNum *big.Int) (*common.BlockId, error) {
	var msgNumBytes ethcommon.Hash
	copy(msgNumBytes[:], math.U256Bytes(messageNum))

	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
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

func (r *BridgeWatcher) LookupMessagesByNum(ctx context.Context, messageNums []*big.Int) ([]*DeliveredInboxMessage, error) {
	msgQuery := make([]ethcommon.Hash, 0, len(messageNums))
	for _, messageNum := range messageNums {
		var msgNumBytes ethcommon.Hash
		copy(msgNumBytes[:], math.U256Bytes(messageNum))
		msgQuery = append(msgQuery, msgNumBytes)
	}

	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{messageDeliveredID}, msgQuery},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	messagesByInbox := make(map[ethcommon.Address][]*big.Int)
	rawMessages := make(map[string]*DeliveredInboxMessage)
	for _, ethLog := range logs {
		parsedLog, err := r.con.ParseMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}
		messagesByInbox[parsedLog.Inbox] = append(messagesByInbox[parsedLog.Inbox], parsedLog.MessageIndex)
		rawMessages[string(parsedLog.MessageIndex.Bytes())] = &DeliveredInboxMessage{
			BlockHash:      common.NewHashFromEth(ethLog.BlockHash),
			BeforeInboxAcc: parsedLog.BeforeInboxAcc,
		}
	}

	inboxMessages := make(map[string]inbox.InboxMessage)
	for con, indexes := range messagesByInbox {
		inboxGetter, err := r.getInboxGetter(con)
		if err != nil {
			return nil, err
		}
		if err := inboxGetter.fillMessageDetails(ctx, indexes, inboxMessages); err != nil {
			return nil, err
		}
	}

	messages := make([]*DeliveredInboxMessage, 0, len(messageNums))
	for _, msgNum := range messageNums {
		rawMsg, ok := rawMessages[string(msgNum.Bytes())]
		if !ok {
			return nil, errors.New("message not found")
		}

		inboxMsg, ok := inboxMessages[string(msgNum.Bytes())]
		if !ok {
			return nil, errors.New("message not found")
		}

		rawMsg.Message = inboxMsg
		messages = append(messages, rawMsg)
	}
	return messages, nil
}

func (r *BridgeWatcher) getInboxGetter(inboxAddress ethcommon.Address) (InboxMessageGetter, error) {
	curInbox, ok := r.inboxes[inboxAddress]
	if ok {
		return curInbox, nil
	}

	return nil, nil
}

type StandardInboxMessageGetter struct {
	con     *ethbridgecontracts.Inbox
	address ethcommon.Address
	client  ethutils.EthClient
}

func (r *StandardInboxMessageGetter) fillMessageDetails(ctx context.Context, messageNums []*big.Int, messages map[string]inbox.InboxMessage) error {
	msgQuery := make([]ethcommon.Hash, 0, len(messageNums))
	for _, messageNum := range messageNums {
		var msgNumBytes ethcommon.Hash
		copy(msgNumBytes[:], math.U256Bytes(messageNum))
		msgQuery = append(msgQuery, msgNumBytes)
	}

	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{inboxMessageDeliveredID, inboxMessageFromOriginID}, msgQuery},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	for _, ethLog := range logs {
		header, err := r.client.HeaderByHash(ctx, ethLog.BlockHash)
		if err != nil {
			return err
		}
		msg, err := r.parseMessage(ctx, ethLog, new(big.Int).SetUint64(header.Time))
		if err != nil {
			return err
		}
		messages[string(msg.InboxSeqNum.Bytes())] = msg
	}
	return nil
}

func (r *StandardInboxMessageGetter) parseMessage(ctx context.Context, ethLog types.Log, timestamp *big.Int) (inbox.InboxMessage, error) {
	chainTime := inbox.ChainTime{
		BlockNum: common.NewTimeBlocks(
			new(big.Int).SetUint64(ethLog.BlockNumber),
		),
		Timestamp: timestamp,
	}
	if ethLog.Topics[0] == inboxMessageDeliveredID {
		parsedLog, err := r.con.ParseInboxMessageDelivered(ethLog)
		if err != nil {
			return inbox.InboxMessage{}, err
		}
		return inbox.InboxMessage{
			Kind:        inbox.Type(parsedLog.Kind),
			Sender:      common.NewAddressFromEth(parsedLog.Sender),
			InboxSeqNum: parsedLog.MessageNum,
			Data:        parsedLog.Data,
			ChainTime:   chainTime,
		}, nil
	} else if ethLog.Topics[0] == inboxMessageFromOriginID {
		tx, _, err := r.client.TransactionByHash(ctx, ethLog.TxHash)
		if err != nil {
			return inbox.InboxMessage{}, err
		}
		args := make(map[string]interface{})
		err = l2MessageFromOriginCallABI.Inputs.UnpackIntoMap(args, tx.Data()[4:])
		if err != nil {
			return inbox.InboxMessage{}, err
		}
		parsedLog, err := r.con.ParseInboxMessageDeliveredFromOrigin(ethLog)
		if err != nil {
			return inbox.InboxMessage{}, err
		}
		return inbox.InboxMessage{
			Kind:        inbox.Type(parsedLog.Kind),
			Sender:      common.NewAddressFromEth(parsedLog.Sender),
			InboxSeqNum: parsedLog.MessageNum,
			Data:        args["messageData"].([]byte),
			ChainTime:   chainTime,
		}, nil
	} else {
		return inbox.InboxMessage{}, errors.New("unexpected log type")
	}
}
