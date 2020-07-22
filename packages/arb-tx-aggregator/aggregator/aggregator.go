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
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type Server struct {
	chain       common.Address
	batch       *batcher.Batcher
	db          *txdb.TxDB
	maxCallTime time.Duration
	maxCallGas  *big.Int
}

// NewServer returns a new instance of the Server class
func NewServer(
	ctx context.Context,
	globalInbox arbbridge.GlobalInbox,
	rollupAddress common.Address,
	db *txdb.TxDB,
) *Server {
	return &Server{
		chain:       rollupAddress,
		batch:       batcher.NewBatcher(ctx, globalInbox, rollupAddress),
		db:          db,
		maxCallTime: 0,
		maxCallGas:  big.NewInt(100000000),
	}
}

// SendTransaction takes a request signed transaction l2message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *Server) SendTransaction(_ context.Context, tx *types.Transaction) (common.Hash, error) {
	return m.batch.SendTransaction(tx)
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

func (m *Server) GetBlockCount(_ context.Context) (uint64, error) {
	id, err := m.db.LatestBlock()
	if err != nil {
		return 0, err
	}
	return id.Height.AsInt().Uint64(), nil
}

func (m *Server) GetOutputMessage(
	_ context.Context,
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
func (m *Server) GetRequestResult(_ context.Context, requestId common.Hash) (value.Value, error) {
	return m.db.GetRequest(requestId)
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetChainAddress(_ context.Context) (ethcommon.Address, error) {
	return m.chain.ToEthAddress(), nil
}

func (m *Server) BlockInfo(_ context.Context, height uint64) (machine.BlockInfo, error) {
	return m.db.GetBlock(height)
}

// Call takes a request from a client to process in a temporary context
// and return the result
func (m *Server) Call(ctx context.Context, msg l2message.Call, sender ethcommon.Address) (value.Value, error) {
	mach, blockId := m.db.CallInfo()
	return m.executeCall(mach, blockId, msg, sender)
}

// PendingCall takes a request from a client to process in a temporary context
// and return the result
func (m *Server) PendingCall(ctx context.Context, msg l2message.Call, sender ethcommon.Address) (value.Value, error) {
	mach, blockId := m.db.CallInfo()
	return m.executeCall(mach, blockId, msg, sender)
}

func (m *Server) executeCall(mach machine.Machine, blockId *common.BlockId, msg l2message.Call, sender ethcommon.Address) (value.Value, error) {
	seq, _ := new(big.Int).SetString("999999999999999999999999", 10)
	inboxMsg := message.NewInboxMessage(
		message.L2Message{Data: l2message.L2MessageAsData(msg)},
		common.NewAddressFromEth(sender),
		seq,
		message.ChainTime{
			BlockNum:  blockId.Height,
			Timestamp: big.NewInt(time.Now().Unix()),
		},
	)

	inbox := value.NewEmptyTuple()
	inbox = value.NewTuple2(inbox, inboxMsg.AsValue())
	assertion, steps := mach.ExecuteAssertion(
		// Call execution is only limited by wall time, so use a massive max steps as an approximation to infinity
		10000000000000000,
		inbox,
		m.maxCallTime,
	)

	// If the machine wasn't able to run and it reports that it is currently
	// blocked, return the block reason to give the client more information
	// as opposed to just returning a general "call produced no output"
	if br := mach.IsBlocked(true); steps == 0 && br != nil {
		log.Println("can't produce solution since machine is blocked", br)
		return nil, fmt.Errorf("can't produce solution since machine is blocked %v", br)
	}

	log.Println("Executed call for", steps, "steps")

	results := assertion.ParseLogs()
	if len(results) == 0 {
		return nil, errors.New("call produced no output")
	}
	lastLogVal := results[len(results)-1]
	lastLog, err := evm.NewResultFromValue(lastLogVal)
	if err != nil {
		return nil, err
	}
	if lastLog.L1Message.MessageID() != inboxMsg.MessageID() {
		// Last produced log is not the call we sent
		return nil, errors.New("call took too long to execute")
	}
	return results[len(results)-1], nil
}
