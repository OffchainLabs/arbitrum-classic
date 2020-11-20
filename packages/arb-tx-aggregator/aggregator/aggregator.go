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

package aggregator

import (
	"bytes"
	"context"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Server struct {
	chain       common.Address
	batch       batcher.TransactionBatcher
	db          *txdb.TxDB
	maxCallTime time.Duration
	maxCallGas  *big.Int
}

// NewServer returns a new instance of the Server class
func NewServer(
	batch batcher.TransactionBatcher,
	rollupAddress common.Address,
	db *txdb.TxDB,
) *Server {
	return &Server{
		chain:       rollupAddress,
		batch:       batch,
		db:          db,
		maxCallTime: 0,
		maxCallGas:  big.NewInt(100000000),
	}
}

// SendTransaction takes a request signed transaction l2message from a Client
// and puts it in a queue to be included in the next transaction batch
func (m *Server) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return m.batch.SendTransaction(ctx, tx)
}

//FindLogs takes a set of parameters and return the list of all logs that match
//the query
func (m *Server) FindLogs(ctx context.Context, fromHeight, toHeight *uint64, addresses []ethcommon.Address, topics [][]ethcommon.Hash) ([]evm.FullLog, error) {
	topicGroups := make([][]common.Hash, 0)
	for _, group := range topics {
		topicGroups = append(topicGroups, common.HashArrayFromEth(group))
	}

	return m.db.FindLogs(
		ctx,
		fromHeight,
		toHeight,
		common.AddressArrayFromEth(addresses),
		topicGroups,
	)
}

func (m *Server) GetBlockCount() uint64 {
	id := m.db.LatestBlockId()
	return id.Height.AsInt().Uint64()
}

func (m *Server) GetOutputMessage(
	args *evm.GetOutputMessageArgs,
	reply *evm.GetOutputMessageReply,
) error {
	msg, err := m.db.GetMessage(args.Index)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := value.MarshalValue(msg, &buf); err != nil {
		return err
	}
	reply.RawVal = hexutil.Encode(buf.Bytes())
	return nil
}

// GetMessageResult returns the value output by the VM in response to the
//l2message with the given hash
func (m *Server) GetRequestResult(requestId common.Hash) (value.Value, error) {
	return m.db.GetRequest(requestId)
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetChainAddress() ethcommon.Address {
	return m.chain.ToEthAddress()
}

func (m *Server) BlockInfoByNumber(height uint64) (*machine.BlockInfo, error) {
	return m.db.GetBlock(height)
}

func (m *Server) BlockInfoByHash(hash common.Hash) (*machine.BlockInfo, error) {
	return m.db.GetBlockWithHash(hash)
}

func (m *Server) GetBlockResults(block *machine.BlockInfo) ([]*evm.TxResult, error) {
	if block.BlockLog == nil {
		// No arb block at this height
		return nil, nil
	}

	res, err := evm.NewBlockResultFromValue(block.BlockLog)
	if err != nil {
		return nil, err
	}
	return m.db.GetBlockResults(res)
}

func (m *Server) GetTxInBlockAtIndexResults(res *evm.BlockInfo, index uint64) (*evm.TxResult, error) {
	txCount := res.BlockStats.TxCount.Uint64()
	if index >= txCount {
		return nil, errors.New("index out of bounds")
	}
	startLog := res.FirstAVMLog().Uint64()
	avmLog, err := m.db.GetLog(startLog + index)
	if err != nil {
		return nil, err
	}
	return evm.NewTxResultFromValue(avmLog)
}

func (m *Server) AdjustGas(msg message.Call) message.Call {
	if msg.MaxGas.Cmp(big.NewInt(0)) == 0 || msg.MaxGas.Cmp(m.maxCallGas) > 0 {
		msg.MaxGas = m.maxCallGas
	}
	return msg
}

// Call takes a request from a Client to process in a temporary context
// and return the result
func (m *Server) Call(msg message.Call, sender ethcommon.Address) (*evm.TxResult, error) {
	msg = m.AdjustGas(msg)
	return m.db.LatestSnapshot().Call(msg, common.NewAddressFromEth(sender))
}

// PendingCall takes a request from a Client to process in a temporary context
// and return the result
func (m *Server) PendingCall(msg message.Call, sender ethcommon.Address) (*evm.TxResult, error) {
	return m.Call(msg, sender)
}

func (m *Server) GetSnapshot(blockHeight uint64) (*snapshot.Snapshot, error) {
	info, err := m.BlockInfoByNumber(blockHeight)
	if err != nil || info == nil {
		return nil, err
	}
	return m.db.GetSnapshot(inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(new(big.Int).SetUint64(blockHeight)),
		Timestamp: new(big.Int).SetUint64(info.Header.Time),
	}), nil
}

func (m *Server) LatestSnapshot() *snapshot.Snapshot {
	return m.db.LatestSnapshot()
}

func (m *Server) PendingSnapshot() *snapshot.Snapshot {
	pending := m.batch.PendingSnapshot()
	if pending == nil {
		return m.LatestSnapshot()
	}
	return pending
}

func (m *Server) PendingTransactionCount(ctx context.Context, account common.Address) *uint64 {
	return m.batch.PendingTransactionCount(ctx, account)
}
