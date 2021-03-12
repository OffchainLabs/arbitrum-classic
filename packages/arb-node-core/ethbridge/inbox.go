package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var l2MessageFromOriginCallABI abi.Method

func init() {
	parsedABI, err := abi.JSON(strings.NewReader(ethbridgecontracts.InboxABI))
	if err != nil {
		panic(err)
	}
	l2MessageFromOriginCallABI = parsedABI.Methods["sendL2MessageFromOrigin"]
}

type StandardInboxWatcher struct {
	con     *ethbridgecontracts.Inbox
	address ethcommon.Address
	client  ethutils.EthClient
}

func NewStandardInboxWatcher(address ethcommon.Address, client ethutils.EthClient) (*StandardInboxWatcher, error) {
	con, err := ethbridgecontracts.NewInbox(address, client)
	if err != nil {
		return nil, err
	}
	return &StandardInboxWatcher{
		con:     con,
		address: address,
		client:  client,
	}, nil
}

func (r *StandardInboxWatcher) fillMessageDetails(ctx context.Context, messageNums []*big.Int, messages map[string][]byte) error {
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
		msgNum, msg, err := r.parseMessage(ctx, ethLog)
		if err != nil {
			return err
		}
		messages[string(msgNum.Bytes())] = msg
	}
	return nil
}

func (r *StandardInboxWatcher) parseMessage(ctx context.Context, ethLog types.Log) (*big.Int, []byte, error) {
	if ethLog.Topics[0] == inboxMessageDeliveredID {
		parsedLog, err := r.con.ParseInboxMessageDelivered(ethLog)
		if err != nil {
			return nil, nil, err
		}
		return parsedLog.MessageNum, parsedLog.Data, nil
	} else if ethLog.Topics[0] == inboxMessageFromOriginID {
		tx, _, err := r.client.TransactionByHash(ctx, ethLog.TxHash)
		if err != nil {
			return nil, nil, err
		}
		args := make(map[string]interface{})
		err = l2MessageFromOriginCallABI.Inputs.UnpackIntoMap(args, tx.Data()[4:])
		if err != nil {
			return nil, nil, err
		}
		parsedLog, err := r.con.ParseInboxMessageDeliveredFromOrigin(ethLog)
		if err != nil {
			return nil, nil, err
		}
		return parsedLog.MessageNum, args["messageData"].([]byte), nil
	} else {
		return nil, nil, errors.New("unexpected log type")
	}
}

type StandardInbox struct {
	*StandardInboxWatcher
	auth *TransactAuth
}

func NewStandardInbox(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*StandardInbox, error) {
	watcher, err := NewStandardInboxWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &StandardInbox{
		StandardInboxWatcher: watcher,
		auth:                 auth,
	}, nil
}

func (s *StandardInbox) SendL2MessageFromOrigin(ctx context.Context, data []byte) (common.Hash, error) {
	tx, err := s.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return s.con.SendL2MessageFromOrigin(auth, data)
	})
	if err != nil {
		return common.Hash{}, err
	}
	return common.NewHashFromEth(tx.Hash()), nil
}
