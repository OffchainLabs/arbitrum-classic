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
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var rollupCreatedID ethcommon.Hash
var nodeCreatedID ethcommon.Hash
var challengeCreatedID ethcommon.Hash

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(ethbridgecontracts.RollupUserFacetABI))
	if err != nil {
		panic(err)
	}
	rollupCreatedID = parsedRollup.Events["RollupCreated"].ID
	nodeCreatedID = parsedRollup.Events["NodeCreated"].ID
	challengeCreatedID = parsedRollup.Events["RollupChallengeStarted"].ID
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
	con       *ethbridgecontracts.RollupUserFacet
	address   ethcommon.Address
	fromBlock int64
	client    ethutils.EthClient
}

func NewRollupWatcher(address ethcommon.Address, fromBlock int64, client ethutils.EthClient) (*RollupWatcher, error) {
	con, err := ethbridgecontracts.NewRollupUserFacet(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &RollupWatcher{
		con:       con,
		address:   address,
		fromBlock: fromBlock,
		client:    client,
	}, nil
}

func (r *RollupWatcher) LookupCreation(ctx context.Context) (*ethbridgecontracts.RollupUserFacetRollupCreated, error) {
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(r.fromBlock),
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{rollupCreatedID}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(logs) == 0 {
		return nil, errors.New("rollup not created")
	}
	if len(logs) > 1 {
		return nil, errors.New("rollup created multiple times")
	}
	ev, err := r.con.ParseRollupCreated(logs[0])
	return ev, errors.WithStack(err)
}

func (r *RollupWatcher) LookupNode(ctx context.Context, number *big.Int) (*core.NodeInfo, error) {
	var numberAsHash ethcommon.Hash
	copy(numberAsHash[:], math.U256Bytes(number))
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(r.fromBlock),
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{nodeCreatedID}, {numberAsHash}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
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
		return nil, errors.WithStack(err)
	}
	proposed := &common.BlockId{
		Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
		HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
	}
	return &core.NodeInfo{
		NodeNum:                 parsedLog.NodeNum,
		BlockProposed:           proposed,
		Assertion:               core.NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
		InboxMaxCount:           parsedLog.InboxMaxCount,
		AfterInboxBatchEndCount: parsedLog.AfterInboxBatchEndCount,
		AfterInboxBatchAcc:      parsedLog.AfterInboxBatchAcc,
		NodeHash:                parsedLog.NodeHash,
	}, nil
}

func (r *RollupWatcher) LookupNodeChildren(ctx context.Context, parentHash [32]byte, fromBlock *big.Int) ([]*core.NodeInfo, error) {
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: fromBlock,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{nodeCreatedID}, nil, {parentHash}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	infos := make([]*core.NodeInfo, 0, len(logs))
	lastHash := parentHash
	for i, ethLog := range logs {
		parsedLog, err := r.con.ParseNodeCreated(ethLog)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		proposed := &common.BlockId{
			Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
		}
		lastHashIsSibling := [1]byte{0}
		if i > 0 {
			lastHashIsSibling[0] = 1
		}
		lastHash = hashing.SoliditySHA3(lastHashIsSibling[:], lastHash[:], parsedLog.ExecutionHash[:], parsedLog.AfterInboxBatchAcc[:])
		infos = append(infos, &core.NodeInfo{
			NodeNum:                 parsedLog.NodeNum,
			BlockProposed:           proposed,
			Assertion:               core.NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
			InboxMaxCount:           parsedLog.InboxMaxCount,
			AfterInboxBatchEndCount: parsedLog.AfterInboxBatchEndCount,
			AfterInboxBatchAcc:      parsedLog.AfterInboxBatchAcc,
			NodeHash:                lastHash,
		})
	}
	return infos, nil
}

func (r *RollupWatcher) LookupChallengedNode(ctx context.Context, address common.Address) (*big.Int, error) {
	addressQuery := ethcommon.Hash{}
	copy(addressQuery[12:], address.Bytes())

	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(r.fromBlock),
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{challengeCreatedID}, {addressQuery}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(logs) == 0 {
		return nil, errors.New("no matching challenge")
	}

	if len(logs) > 1 {
		return nil, errors.New("too many matching challenges")
	}

	challenge, err := r.con.ParseRollupChallengeStarted(logs[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return challenge.ChallengedNode, nil
}

func (r *RollupWatcher) StakerInfo(ctx context.Context, staker common.Address) (*StakerInfo, error) {
	info, err := r.con.StakerMap(&bind.CallOpts{Context: ctx}, staker.ToEthAddress())
	if err != nil {
		return nil, errors.WithStack(err)
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
	blocks, err := r.con.MinimumAssertionPeriod(&bind.CallOpts{Context: ctx})
	return blocks, errors.WithStack(err)
}

func (r *RollupWatcher) SequencerBridge(ctx context.Context) (common.Address, error) {
	addr, err := r.con.SequencerBridge(&bind.CallOpts{Context: ctx})
	return common.NewAddressFromEth(addr), errors.WithStack(err)
}

func (r *RollupWatcher) DelayedBridge(ctx context.Context) (common.Address, error) {
	addr, err := r.con.DelayedBridge(&bind.CallOpts{Context: ctx})
	return common.NewAddressFromEth(addr), errors.WithStack(err)
}

func (r *RollupWatcher) StakerCount(ctx context.Context) (*big.Int, error) {
	count, err := r.con.StakerCount(&bind.CallOpts{Context: ctx})
	return count, errors.WithStack(err)
}

func (r *RollupWatcher) ArbGasSpeedLimitPerBlock(ctx context.Context) (*big.Int, error) {
	speed, err := r.con.ArbGasSpeedLimitPerBlock(&bind.CallOpts{Context: ctx})
	return speed, errors.WithStack(err)
}

func (r *RollupWatcher) CurrentRequiredStake(ctx context.Context) (*big.Int, error) {
	stake, err := r.con.CurrentRequiredStake(&bind.CallOpts{Context: ctx})
	return stake, errors.WithStack(err)
}

func (r *RollupWatcher) BaseStake(ctx context.Context) (*big.Int, error) {
	stake, err := r.con.BaseStake(&bind.CallOpts{Context: ctx})
	return stake, errors.WithStack(err)
}

func (r *RollupWatcher) LatestConfirmedNode(ctx context.Context) (*big.Int, error) {
	node, err := r.con.LatestConfirmed(&bind.CallOpts{Context: ctx})
	return node, errors.WithStack(err)
}

func (r *RollupWatcher) FirstUnresolvedNode(ctx context.Context) (*big.Int, error) {
	node, err := r.con.FirstUnresolvedNode(&bind.CallOpts{Context: ctx})
	return node, errors.WithStack(err)
}

func (r *RollupWatcher) LatestNodeCreated(ctx context.Context) (*big.Int, error) {
	node, err := r.con.LatestNodeCreated(&bind.CallOpts{Context: ctx})
	return node, errors.WithStack(err)
}

func (r *RollupWatcher) ConfirmPeriodBlocks(ctx context.Context) (*big.Int, error) {
	blocks, err := r.con.ConfirmPeriodBlocks(&bind.CallOpts{Context: ctx})
	return blocks, errors.WithStack(err)
}

func (r *RollupWatcher) GetNode(ctx context.Context, node core.NodeID) (*NodeWatcher, error) {
	nodeAddress, err := r.con.GetNode(&bind.CallOpts{Context: ctx}, node)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return NewNodeWatcher(nodeAddress, r.client)
}
