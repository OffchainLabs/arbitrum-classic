package ethbridge

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var rollupCreatedID ethcommon.Hash
var nodeCreatedID ethcommon.Hash
var challengeCreatedID ethcommon.Hash
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
	l2MessageFromOriginCallABI = parsedRollup.Methods["sendL2MessageFromOrigin"]
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

func (r *RollupWatcher) LookupNode(ctx context.Context, number *big.Int) (*core.NodeInfo, error) {
	var numberAsHash ethcommon.Hash
	copy(numberAsHash[:], math.U256Bytes(number))
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{nodeCreatedID}, {numberAsHash}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, errors.New("Couldn't find requested node")
	}
	if len(logs) > 1 {
		return nil, errors.New("Found multiple instances of requested node")
	}
	ethLog := logs[0]
	parsedLog, err := r.con.ParseNodeCreated(ethLog)
	if err != nil {
		return nil, err
	}
	proposed := &common.BlockId{
		Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
		HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
	}
	return &core.NodeInfo{
		NodeNum:       parsedLog.NodeNum,
		BlockProposed: proposed,
		Assertion:     core.NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
		InboxMaxCount: parsedLog.InboxMaxCount,
		NodeHash:      parsedLog.NodeHash,
	}, nil
}

func (r *RollupWatcher) LookupNodeChildren(ctx context.Context, parentHash [32]byte) ([]*core.NodeInfo, error) {
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{nodeCreatedID}, {}, {parentHash}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	infos := make([]*core.NodeInfo, 0, len(logs))
	lastHash := parentHash
	for i, ethLog := range logs {
		parsedLog, err := r.con.ParseNodeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		proposed := &common.BlockId{
			Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
		}
		lastHashIsSibling := [1]byte{0}
		if i > 0 {
			lastHashIsSibling[0] = 1
		}
		lastHash = hashing.SoliditySHA3(lastHashIsSibling, lastHash, parsedLog.ExecutionHash, parsedLog.AfterInboxHash)
		infos = append(infos, &core.NodeInfo{
			NodeNum:       parsedLog.NodeNum,
			BlockProposed: proposed,
			Assertion:     core.NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
			InboxMaxCount: parsedLog.InboxMaxCount,
			NodeHash:      lastHash,
		})
	}
	return infos, nil
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

func (r *RollupWatcher) Bridge(ctx context.Context) (common.Address, error) {
	addr, err := r.con.Bridge(&bind.CallOpts{Context: ctx})
	return common.NewAddressFromEth(addr), err
}

func (r *RollupWatcher) StakeToken(ctx context.Context) (common.Address, error) {
	addr, err := r.con.StakeToken(&bind.CallOpts{Context: ctx})
	return common.NewAddressFromEth(addr), err
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

func (r *RollupWatcher) ConfirmPeriodBlocks(ctx context.Context) (*big.Int, error) {
	return r.con.ConfirmPeriodBlocks(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) GetNode(ctx context.Context, node core.NodeID) (*NodeWatcher, error) {
	nodeAddress, err := r.con.GetNode(&bind.CallOpts{Context: ctx}, node)
	if err != nil {
		return nil, err
	}
	return NewNodeWatcher(nodeAddress, r.client)
}
