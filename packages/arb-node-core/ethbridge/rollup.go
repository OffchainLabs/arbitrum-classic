package ethbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var nodeCreatedID ethcommon.Hash
var messageDeliveredID ethcommon.Hash
var messageDeliveredFromOriginID ethcommon.Hash
var l2MessageFromOriginCallABI abi.Method

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(ethbridgecontracts.RollupABI))
	if err != nil {
		panic(err)
	}
	nodeCreatedID = parsedRollup.Events["NodeCreated"].ID
	messageDeliveredID = parsedRollup.Events["MessageDelivered"].ID
	messageDeliveredFromOriginID = parsedRollup.Events["MessageDeliveredFromOrigin"].ID
	l2MessageFromOriginCallABI = parsedRollup.Methods["sendL2MessageFromOrigin"]
}

type StakerInfo struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge *common.Address
}

type NodeInfo struct {
	NodeNum   NodeID
	Assertion *Assertion
}

type DeliveredInboxMessage struct {
	BeforeInboxAcc common.Hash
	Message        inbox.InboxMessage
}

func (d *DeliveredInboxMessage) AfterInboxAcc() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(d.BeforeInboxAcc),
		hashing.Bytes32(d.Message.CommitmentHash()),
	)
}

type NodeID *big.Int

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

func (r *RollupWatcher) LookupNodes(ctx context.Context, nodes []*big.Int) ([]*NodeInfo, error) {
	var nodeQuery []ethcommon.Hash
	for _, node := range nodes {
		var nd ethcommon.Hash
		copy(nd[:], math.U256Bytes(node))
		nodeQuery = append(nodeQuery, nd)
	}
	query := ethereum.FilterQuery{
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
	infos := make([]*NodeInfo, 0, len(logs))
	for _, ethLog := range logs {
		parsedLog, err := r.con.ParseNodeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		infos = append(infos, &NodeInfo{
			NodeNum:   parsedLog.NodeNum,
			Assertion: NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
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
			Kind:        parsedLog.Kind,
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
			Kind:        parsedLog.Kind,
			Sender:      common.NewAddressFromEth(parsedLog.Sender),
			InboxSeqNum: parsedLog.MessageNum,
			Data:        args["messageData"].([]byte),
			ChainTime:   chainTime,
		}
		return &DeliveredInboxMessage{
			BeforeInboxAcc: parsedLog.BeforeInboxAcc,
			Message:        msg,
		}, nil
	} else {
		return nil, errors.New("unexpected log type")
	}
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

func (r *RollupWatcher) StakerCount(ctx context.Context) (*big.Int, error) {
	return r.con.StakerCount(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) CurrentRequiredStake(ctx context.Context) (*big.Int, error) {
	return r.con.CurrentRequiredStake(&bind.CallOpts{Context: ctx})
}

func (r *Rollup) LatestConfirmedNode(ctx context.Context) (*big.Int, error) {
	return r.con.LatestConfirmed(&bind.CallOpts{Context: ctx})
}

func (r *Rollup) FirstUnresolvedNode(ctx context.Context) (*big.Int, error) {
	return r.con.FirstUnresolvedNode(&bind.CallOpts{Context: ctx})
}

func (r *Rollup) ChallengePeriodBlocks(ctx context.Context) (*big.Int, error) {
	return r.con.ChallengePeriodBlocks(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) GetNode(ctx context.Context, node NodeID) (*NodeWatcher, error) {
	nodeAddress, err := r.con.Nodes(&bind.CallOpts{Context: ctx}, node)
	if err != nil {
		return nil, err
	}
	return NewNodeWatcher(nodeAddress, r.client)
}

type Rollup struct {
	*RollupWatcher
	auth *TransactAuth
}

func NewRollup(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*Rollup, error) {
	watcher, err := NewRollupWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Rollup{
		RollupWatcher: watcher,
		auth:          auth,
	}, nil
}

func (r *Rollup) RejectNextNode(ctx context.Context, node NodeID, staker common.Address) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RejectNextNode(auth, node, staker.ToEthAddress())
	})
}

func (r *Rollup) RejectNextNodeOutOfOrder(ctx context.Context) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RejectNextNodeOutOfOrder(auth)
	})
}

func (r *Rollup) ConfirmNextNode(
	ctx context.Context,
	logAcc common.Hash,
	sends [][]byte,
) (*types.Transaction, error) {
	var sendsData []byte
	sendLengths := make([]*big.Int, 0, len(sends))
	for _, msg := range sends {
		sendsData = append(sendsData, msg...)
		sendLengths = append(sendLengths, new(big.Int).SetInt64(int64(len(msg))))
	}
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ConfirmNextNode(auth, logAcc, sendsData, sendLengths)
	})
}

func (r *Rollup) NewStakeOnExistingNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.NewStakeOnExistingNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
		)
	})
}

func (r *Rollup) AddStakeOnExistingNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.AddStakeOnExistingNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
		)
	})
}

func (r *Rollup) NewStakeOnNewNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
	prev NodeID,
	assertion *Assertion,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.NewStakeOnNewNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
			prev,
			assertion.BytesFields(),
			assertion.IntFields(),
		)
	})
}

func (r *Rollup) AddStakeOnNewNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
	assertion *Assertion,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.AddStakeOnNewNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
			assertion.BytesFields(),
			assertion.IntFields(),
		)
	})
}

func (r *Rollup) ReturnOldDeposit(ctx context.Context, staker common.Address) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ReturnOldDeposit(auth, staker.ToEthAddress())
	})
}

func (r *Rollup) AddToDeposit(ctx context.Context, address common.Address, amount *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = amount
		return r.con.AddToDeposit(auth, address.ToEthAddress())
	})
}

func (r *Rollup) ReduceDeposit(ctx context.Context, amount *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ReduceDeposit(auth, amount)
	})
}

func (r *Rollup) CreateChallenge(
	ctx context.Context,
	staker1 common.Address,
	node1 NodeID,
	staker2 common.Address,
	node2 NodeID,
	assertion *Assertion,
	inboxMaxHash common.Hash,
	inboxMaxCount *big.Int,
	arbGasSpeedLimitPerBlock *big.Int,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.CreateChallenge(
			auth,
			staker1.ToEthAddress(),
			node1,
			staker2.ToEthAddress(),
			node2,
			assertion.InboxConsistencyHash(inboxMaxHash, inboxMaxCount),
			assertion.InboxDeltaHash(),
			assertion.ExecutionHash(),
			assertion.CheckTime(arbGasSpeedLimitPerBlock),
		)
	})
}

func (r *Rollup) RemoveZombie(ctx context.Context, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RemoveZombie(auth, zombieNum, maxNodes)
	})
}

func (r *Rollup) RemoveOldZombies(ctx context.Context, startIndex *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RemoveOldZombies(auth, startIndex)
	})
}
