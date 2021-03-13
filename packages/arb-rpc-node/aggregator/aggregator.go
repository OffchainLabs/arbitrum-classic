/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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

package aggregator

import (
	"context"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

var logger = log.With().Caller().Str("component", "aggregator").Logger()

type Server struct {
	chain      common.Address
	batch      batcher.TransactionBatcher
	db         *txdb.TxDB
	maxCallGas *big.Int
	scope      event.SubscriptionScope
}

// NewServer returns a new instance of the Server class
func NewServer(
	batch batcher.TransactionBatcher,
	rollupAddress common.Address,
	db *txdb.TxDB,
) *Server {
	return &Server{
		chain:      rollupAddress,
		batch:      batch,
		db:         db,
		maxCallGas: big.NewInt(1000000000),
	}
}

// SendTransaction takes a request signed transaction l2message from a Client
// and puts it in a queue to be included in the next transaction batch
func (m *Server) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return m.batch.SendTransaction(ctx, tx)
}

func (m *Server) GetBlockCount() (uint64, error) {
	latest, err := m.db.BlockCount()
	if err != nil {
		return 0, err
	}
	return latest, nil
}

func (m *Server) BlockNum(block *rpc.BlockNumber) (uint64, error) {
	if *block == rpc.LatestBlockNumber || *block == rpc.PendingBlockNumber {
		latest, err := m.db.LatestBlock()
		if err != nil {
			return 0, err
		}
		return latest.Header.Number.Uint64(), nil
	} else if *block >= 0 {
		return uint64(*block), nil
	} else {
		return 0, errors.Errorf("unsupported BlockNumber: %v", block.Int64())
	}
}

// GetMessageResult returns the value output by the VM in response to the
// l2message with the given hash
func (m *Server) GetRequestResult(requestId common.Hash) (*evm.TxResult, error) {
	return m.db.GetRequest(requestId)
}

func (m *Server) GetL2ToL1Proof(batchNumber *big.Int, index uint64) (*evm.MerkleRootProof, error) {
	batch, err := m.db.GetMessageBatch(batchNumber)
	if err != nil {
		return nil, err
	}
	if batch == nil {
		return nil, errors.New("batch doesn't exist")
	}
	return batch.GenerateProof(index)
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetChainAddress() ethcommon.Address {
	return m.chain.ToEthAddress()
}

func (m *Server) BlockInfoByNumber(height uint64) (*machine.BlockInfo, error) {
	return m.db.GetBlock(height)
}

func (m *Server) BlockLogFromInfo(block *machine.BlockInfo) (*evm.BlockInfo, error) {
	return m.db.GetL2Block(block)
}

func (m *Server) BlockInfoByHash(hash common.Hash) (*machine.BlockInfo, error) {
	return m.db.GetBlockWithHash(hash)
}

func (m *Server) GetMachineBlockResults(block *machine.BlockInfo) (*evm.BlockInfo, []*evm.TxResult, error) {
	return m.db.GetBlockResults(block)
}

func (m *Server) GetTxInBlockAtIndexResults(res *machine.BlockInfo, index uint64) (*evm.TxResult, error) {
	avmLog, err := core.GetZeroOrOneLog(m.db.Lookup, new(big.Int).SetUint64(res.InitialLogIndex()+index))
	if err != nil || avmLog == nil {
		return nil, err
	}
	evmRes, err := evm.NewTxResultFromValue(avmLog)
	if err != nil {
		return nil, err
	}
	if evmRes.IncomingRequest.L2BlockNumber.Cmp(res.Header.Number) != 0 {
		return nil, nil
	}
	return evmRes, nil
}

func (m *Server) AdjustGas(msg message.ContractTransaction) message.ContractTransaction {
	if msg.MaxGas.Cmp(big.NewInt(0)) == 0 || msg.MaxGas.Cmp(m.maxCallGas) > 0 {
		msg.MaxGas = m.maxCallGas
	}
	return msg
}

// Call takes a request from a Client to process in a temporary context
// and return the result
func (m *Server) Call(msg message.ContractTransaction, sender ethcommon.Address) (*evm.TxResult, error) {
	msg = m.AdjustGas(msg)
	snap, err := m.db.LatestSnapshot()
	if err != nil {
		return nil, err
	}
	return snap.Call(msg, common.NewAddressFromEth(sender))
}

// PendingCall takes a request from a Client to process in a temporary context
// and return the result
func (m *Server) PendingCall(msg message.ContractTransaction, sender ethcommon.Address) (*evm.TxResult, error) {
	return m.Call(msg, sender)
}

func (m *Server) GetSnapshot(blockHeight uint64) (*snapshot.Snapshot, error) {
	return m.db.GetSnapshot(blockHeight)
}

func (m *Server) LatestSnapshot() (*snapshot.Snapshot, error) {
	return m.db.LatestSnapshot()
}

func (m *Server) PendingSnapshot() (*snapshot.Snapshot, error) {
	pending, err := m.batch.PendingSnapshot()
	if err != nil {
		return nil, err
	}
	if pending == nil {
		return m.LatestSnapshot()
	}
	return pending, nil
}

func (m *Server) PendingTransactionCount(ctx context.Context, account common.Address) *uint64 {
	return m.batch.PendingTransactionCount(ctx, account)
}

func (m *Server) ChainDb() ethdb.Database {
	return nil
}

func (m *Server) HeaderByNumber(_ context.Context, blockNumber rpc.BlockNumber) (*types.Header, error) {
	height, err := m.BlockNum(&blockNumber)
	if err != nil {
		return nil, err
	}

	info, err := m.db.GetBlock(height)
	if err != nil || info == nil {
		return nil, err
	}

	return info.Header, nil
}

func (m *Server) HeaderByHash(_ context.Context, blockHash ethcommon.Hash) (*types.Header, error) {
	info, err := m.BlockInfoByHash(common.NewHashFromEth(blockHash))
	if err != nil || info == nil {
		return nil, err
	}

	return info.Header, nil
}

func (m *Server) GetReceipts(_ context.Context, blockHash ethcommon.Hash) (types.Receipts, error) {
	info, err := m.db.GetBlockWithHash(common.NewHashFromEth(blockHash))
	if err != nil || info == nil {
		return nil, err
	}
	_, results, err := m.GetMachineBlockResults(info)
	if err != nil || results == nil {
		return nil, err
	}
	receipts := make(types.Receipts, 0, len(results))
	for _, res := range results {
		receipts = append(receipts, res.ToEthReceipt(common.NewHashFromEth(blockHash)))
	}
	return receipts, nil
}

func (m *Server) GetLogs(_ context.Context, blockHash ethcommon.Hash) ([][]*types.Log, error) {
	info, err := m.db.GetBlockWithHash(common.NewHashFromEth(blockHash))
	if err != nil || info == nil {
		return nil, err
	}
	_, results, err := m.GetMachineBlockResults(info)
	if err != nil || results == nil {
		return nil, err
	}
	logs := make([][]*types.Log, 0, len(results))
	for _, res := range results {
		logs = append(logs, res.EthLogs(common.NewHashFromEth(blockHash)))
	}
	return logs, nil
}

func (m *Server) BloomStatus() (uint64, uint64) {
	return 0, 0
}

func (m *Server) ServiceFilter(_ context.Context, _ *bloombits.MatcherSession) {
	// Currently not implemented
}

func (m *Server) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
	return m.scope.Track(m.batch.SubscribeNewTxsEvent(ch))
}

func (m *Server) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return m.scope.Track(m.db.SubscribePendingLogsEvent(ch))
}

func (m *Server) SubscribeChainEvent(ch chan<- ethcore.ChainEvent) event.Subscription {
	return m.scope.Track(m.db.SubscribeChainEvent(ch))
}

func (m *Server) SubscribeChainHeadEvent(ch chan<- ethcore.ChainEvent) event.Subscription {
	return m.scope.Track(m.db.SubscribeChainHeadEvent(ch))
}

func (m *Server) SubscribeChainSideEvent(ch chan<- ethcore.ChainEvent) event.Subscription {
	return m.scope.Track(m.db.SubscribeChainSideEvent(ch))
}

func (m *Server) SubscribeRemovedLogsEvent(ch chan<- ethcore.RemovedLogsEvent) event.Subscription {
	return m.scope.Track(m.db.SubscribeRemovedLogsEvent(ch))
}

func (m *Server) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return m.scope.Track(m.db.SubscribeLogsEvent(ch))
}

func (m *Server) SubscribeBlockProcessingEvent(ch chan<- []*types.Log) event.Subscription {
	return m.scope.Track(m.db.SubscribeBlockProcessingEvent(ch))
}
