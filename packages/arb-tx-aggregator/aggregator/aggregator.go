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
	"github.com/ethereum/go-ethereum/trie"
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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
)

type Server struct {
	Client      ethutils.EthClient
	chain       common.Address
	batch       batcher.TransactionBatcher
	db          *txdb.TxDB
	maxCallTime time.Duration
	maxCallGas  *big.Int
}

// NewServer returns a new instance of the Server class
func NewServer(
	client ethutils.EthClient,
	batch batcher.TransactionBatcher,
	rollupAddress common.Address,
	db *txdb.TxDB,
) *Server {
	return &Server{
		Client:      client,
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

func (m *Server) BlockInfo(height uint64) (*machine.BlockInfo, error) {
	return m.db.GetBlock(height)
}

func (m *Server) GetBlockHeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	ethHeader, err := m.Client.HeaderByHash(ctx, hash.ToEthHash())
	if err != nil {
		return nil, err
	}

	currentBlock, err := m.db.GetBlock(ethHeader.Number.Uint64())
	if err != nil {
		return nil, err
	}

	return getBlockHeader(currentBlock, ethHeader)
}

func (m *Server) GetBlockHeaderByNumber(ctx context.Context, height uint64) (*types.Header, error) {
	currentBlock, err := m.db.GetBlock(height)
	if err != nil {
		return nil, err
	}

	var ethHeader *types.Header
	if currentBlock != nil {
		ethHeader, err = m.Client.HeaderByHash(ctx, currentBlock.Hash.ToEthHash())
	} else {
		ethHeader, err = m.Client.HeaderByNumber(ctx, new(big.Int).SetUint64(height))
	}
	if err != nil {
		return nil, err
	}
	return getBlockHeader(currentBlock, ethHeader)
}

func GetBlockFields(currentBlock *machine.BlockInfo, res *evm.BlockInfo) (types.Bloom, uint64, uint64) {
	gasUsed := uint64(0)
	gasLimit := uint64(100000000)
	bloom := types.Bloom{}
	if currentBlock != nil {
		gasUsed = res.BlockStats.GasUsed.Uint64()
		gasLimit = res.GasLimit.Uint64()
		bloom = currentBlock.Bloom
	}
	return bloom, gasLimit, gasUsed
}

func getBlockHeader(currentBlock *machine.BlockInfo, ethHeader *types.Header) (*types.Header, error) {
	var res *evm.BlockInfo
	if currentBlock != nil {
		var err error
		res, err = evm.NewBlockResultFromValue(currentBlock.BlockLog)
		if err != nil {
			return nil, err
		}
	}

	ethHeader.Bloom, ethHeader.GasLimit, ethHeader.GasUsed = GetBlockFields(currentBlock, res)
	return ethHeader, nil
}

func (m *Server) GetBlockInfo(block *machine.BlockInfo) (*evm.BlockInfo, error) {
	if block == nil {
		// No arbitrum block at this height
		return nil, nil
	}

	return evm.NewBlockResultFromValue(block.BlockLog)
}

func (m *Server) GetBlockResults(res *evm.BlockInfo) ([]*evm.TxResult, error) {
	if res == nil {
		return nil, nil
	}
	txCount := res.BlockStats.TxCount.Uint64()
	startLog := res.FirstAVMLog().Uint64()
	results := make([]*evm.TxResult, 0, txCount)
	for i := uint64(0); i < txCount; i++ {
		avmLog, err := m.db.GetLog(startLog + i)
		if err != nil {
			return nil, err
		}
		res, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
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

func (m *Server) GetBlock(ctx context.Context, height uint64) (*types.Block, error) {
	l1BlockInfo, err := m.Client.BlockInfoByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}
	header, err := m.GetBlockHeaderByNumber(ctx, height)
	if err != nil {
		return nil, err
	}

	block, err := m.BlockInfo(height)
	if err != nil {
		return nil, err
	}

	blockInfo, err := m.GetBlockInfo(block)
	if err != nil {
		return nil, err
	}

	results, err := m.GetBlockResults(blockInfo)
	if err != nil {
		return nil, err
	}

	transactions := make([]*types.Transaction, 0, len(results))
	receipts := make([]*types.Receipt, 0, len(results))
	for _, res := range results {
		receipt := res.ToEthReceipt(common.NewHashFromEth(l1BlockInfo.Hash))
		receipts = append(receipts, receipt)
		tx, err := GetTransaction(res.IncomingRequest)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}

	return types.NewBlock(header, transactions, nil, receipts, new(trie.Trie)), nil
}

func GetTransaction(msg evm.IncomingRequest) (*types.Transaction, error) {
	if msg.Kind != message.L2Type {
		return nil, errors.New("result is not a transaction")
	}
	l2msg, err := message.L2Message{Data: msg.Data}.AbstractMessage()
	if err != nil {
		return nil, err
	}
	ethMsg, ok := l2msg.(message.EthConvertable)
	if !ok {
		return nil, errors.New("message not convertible to receipt")
	}
	return ethMsg.AsEthTx(), nil
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

func (m *Server) GetSnapshot(ctx context.Context, blockHeight uint64) (*snapshot.Snapshot, error) {
	height := new(big.Int).SetUint64(blockHeight)
	header, err := m.Client.HeaderByNumber(ctx, height)
	if err != nil {
		return nil, err
	}
	return m.db.GetSnapshot(inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(height),
		Timestamp: new(big.Int).SetUint64(header.Time),
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
