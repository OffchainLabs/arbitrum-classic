package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
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
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var rollupABI abi.ABI
var rollupCreatedID ethcommon.Hash
var nodeCreatedID ethcommon.Hash
var challengeCreatedID ethcommon.Hash
var messageDeliveredID ethcommon.Hash
var messageDeliveredFromOriginID ethcommon.Hash
var l2MessageFromOriginCallABI abi.Method

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(ethbridgecontracts.RollupABI))
	if err != nil {
		panic(err)
	}
	rollupCreatedID = parsedRollup.Events["RollupCreated"].ID
	nodeCreatedID = parsedRollup.Events["NodeCreated"].ID
	challengeCreatedID = parsedRollup.Events["RollupChallengeStarted"].ID
	messageDeliveredID = parsedRollup.Events["MessageDelivered"].ID
	messageDeliveredFromOriginID = parsedRollup.Events["MessageDeliveredFromOrigin"].ID
	l2MessageFromOriginCallABI = parsedRollup.Methods["sendL2MessageFromOrigin"]
	rollupABI = parsedRollup
}

type StakerInfo struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge *common.Address
}

type DeliveredInboxMessage struct {
	BlockHash      common.Hash
	BeforeInboxAcc common.Hash
	Message        inbox.InboxMessage
}

func (d *DeliveredInboxMessage) AfterInboxAcc() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(d.BeforeInboxAcc),
		hashing.Bytes32(d.Message.CommitmentHash()),
	)
}

func (d *DeliveredInboxMessage) Block() *common.BlockId {
	return &common.BlockId{
		Height:     d.Message.ChainTime.BlockNum,
		HeaderHash: d.BlockHash,
	}
}

type RollupWatcher struct {
	con     *ethbridgecontracts.Rollup
	address ethcommon.Address
	client  ethutils.EthClient
}

func NewRollupWatcher(address ethcommon.Address, client ethutils.EthClient) (*RollupWatcher, error) {
	con, err := ethbridgecontracts.NewRollup(address, client)
	if err != nil {
		return nil, err
	}

	return &RollupWatcher{
		con:     con,
		address: address,
		client:  client,
	}, nil
}

func (r *RollupWatcher) LookupCreation(ctx context.Context) (*ethbridgecontracts.RollupRollupCreated, error) {
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{rollupCreatedID}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, errors.New("rollup not created")
	}
	if len(logs) > 1 {
		return nil, errors.New("rollup created multiple times")
	}
	return r.con.ParseRollupCreated(logs[0])
}

func (r *RollupWatcher) LookupNodes(ctx context.Context, nodes []*big.Int) ([]*core.NodeInfo, error) {
	var nodeQuery []ethcommon.Hash
	for _, node := range nodes {
		var nd ethcommon.Hash
		copy(nd[:], math.U256Bytes(node))
		nodeQuery = append(nodeQuery, nd)
	}
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{nodeCreatedID}, nodeQuery},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	infos := make([]*core.NodeInfo, 0, len(logs))
	for _, ethLog := range logs {
		parsedLog, err := r.con.ParseNodeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		proposed := &common.BlockId{
			Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
		}
		infos = append(infos, &core.NodeInfo{
			NodeNum:       parsedLog.NodeNum,
			BlockProposed: proposed,
			Assertion:     core.NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
			InboxMaxCount: parsedLog.InboxMaxCount,
			InboxMaxHash:  parsedLog.InboxMaxHash,
		})
	}
	return infos, nil
}

func (r *RollupWatcher) LookupMessagesByNum(ctx context.Context, messageNums []*big.Int) ([]*DeliveredInboxMessage, error) {
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
		Topics:    [][]ethcommon.Hash{{messageDeliveredID, messageDeliveredFromOriginID}, msgQuery},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	messages := make([]*DeliveredInboxMessage, 0, len(logs))
	for _, ethLog := range logs {
		header, err := r.client.HeaderByHash(ctx, ethLog.BlockHash)
		if err != nil {
			return nil, err
		}
		msg, err := r.parseMessage(ctx, ethLog, new(big.Int).SetUint64(header.Time))
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *RollupWatcher) parseMessage(ctx context.Context, ethLog types.Log, timestamp *big.Int) (*DeliveredInboxMessage, error) {
	chainTime := inbox.ChainTime{
		BlockNum: common.NewTimeBlocks(
			new(big.Int).SetUint64(ethLog.BlockNumber),
		),
		Timestamp: timestamp,
	}
	if ethLog.Topics[0] == messageDeliveredID {
		parsedLog, err := r.con.ParseMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}
		msg := inbox.InboxMessage{
			Kind:        inbox.Type(parsedLog.Kind),
			Sender:      common.NewAddressFromEth(parsedLog.Sender),
			InboxSeqNum: parsedLog.MessageNum,
			Data:        parsedLog.Data,
			ChainTime:   chainTime,
		}
		return &DeliveredInboxMessage{
			BeforeInboxAcc: parsedLog.BeforeInboxAcc,
			Message:        msg,
		}, nil
	} else if ethLog.Topics[0] == messageDeliveredFromOriginID {
		tx, _, err := r.client.TransactionByHash(ctx, ethLog.TxHash)
		if err != nil {
			return nil, err
		}
		args := make(map[string]interface{})
		err = l2MessageFromOriginCallABI.Inputs.UnpackIntoMap(args, tx.Data()[4:])
		if err != nil {
			return nil, err
		}
		parsedLog, err := r.con.ParseMessageDeliveredFromOrigin(ethLog)
		if err != nil {
			return nil, err
		}
		msg := inbox.InboxMessage{
			Kind:        inbox.Type(parsedLog.Kind),
			Sender:      common.NewAddressFromEth(parsedLog.Sender),
			InboxSeqNum: parsedLog.MessageNum,
			Data:        args["messageData"].([]byte),
			ChainTime:   chainTime,
		}
		return &DeliveredInboxMessage{
			BlockHash:      common.NewHashFromEth(ethLog.BlockHash),
			BeforeInboxAcc: parsedLog.BeforeInboxAcc,
			Message:        msg,
		}, nil
	} else {
		return nil, errors.New("unexpected log type")
	}
}

func (r *RollupWatcher) LookupChallengedNode(ctx context.Context, address common.Address) (*big.Int, error) {
	addressQuery := ethcommon.Hash{}
	copy(addressQuery[12:], address.Bytes())

	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{challengeCreatedID}, {addressQuery}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.New("no matching challenge")
	}

	if len(logs) > 1 {
		return nil, errors.New("too many matching challenges")
	}

	challenge, err := r.con.ParseRollupChallengeStarted(logs[0])
	if err != nil {
		return nil, err
	}

	return challenge.ChallengedNode, nil
}

func (r *RollupWatcher) StakerInfo(ctx context.Context, staker common.Address) (*StakerInfo, error) {
	info, err := r.con.StakerMap(&bind.CallOpts{Context: ctx}, staker.ToEthAddress())
	if err != nil {
		return nil, err
	}
	if !info.IsStaked {
		return nil, nil
	}
	stakerInfo := &StakerInfo{
		Index:            info.Index,
		LatestStakedNode: info.LatestStakedNode,
		AmountStaked:     info.AmountStaked,
	}
	emptyAddress := ethcommon.Address{}
	if info.CurrentChallenge != emptyAddress {
		chal := common.NewAddressFromEth(info.CurrentChallenge)
		stakerInfo.CurrentChallenge = &chal
	}
	return stakerInfo, nil
}

func (r *RollupWatcher) MinimumAssertionPeriod(ctx context.Context) (*big.Int, error) {
	return r.con.MinimumAssertionPeriod(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) GetStakers(ctx context.Context) ([]common.Address, error) {
	addresses, err := r.con.GetStakers(&bind.CallOpts{Context: ctx}, big.NewInt(0), math.MaxBig256)
	if err != nil {
		return nil, err
	}
	return common.AddressArrayFromEth(addresses), nil
}

func (r *RollupWatcher) StakerCount(ctx context.Context) (*big.Int, error) {
	return r.con.StakerCount(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) ArbGasSpeedLimitPerBlock(ctx context.Context) (*big.Int, error) {
	return r.con.ArbGasSpeedLimitPerBlock(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) CurrentRequiredStake(ctx context.Context) (*big.Int, error) {
	return r.con.CurrentRequiredStake(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) LatestConfirmedNode(ctx context.Context) (*big.Int, error) {
	return r.con.LatestConfirmed(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) FirstUnresolvedNode(ctx context.Context) (*big.Int, error) {
	return r.con.FirstUnresolvedNode(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) LatestNodeCreated(ctx context.Context) (*big.Int, error) {
	return r.con.LatestNodeCreated(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) ChallengePeriodBlocks(ctx context.Context) (*big.Int, error) {
	return r.con.ChallengePeriodBlocks(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) GetNode(ctx context.Context, node core.NodeID) (*NodeWatcher, error) {
	nodeAddress, err := r.con.Nodes(&bind.CallOpts{Context: ctx}, node)
	if err != nil {
		return nil, err
	}
	return NewNodeWatcher(nodeAddress, r.client)
}
