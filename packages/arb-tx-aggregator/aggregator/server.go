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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type Server struct {
	chain       common.Address
	batch       *batcher.Batcher
	db          *txdb.TxDB
	maxCallTime time.Duration
}

// NewServer returns a new instance of the Server class
func NewServer(
	ctx context.Context,
	globalInbox arbbridge.GlobalInbox,
	rollupAddress common.Address,
	db *txdb.TxDB,
) *Server {
	return &Server{
		chain: rollupAddress,
		batch: batcher.NewBatcher(ctx, globalInbox, rollupAddress),
		db:    db,
	}
}

// SendTransaction takes a request signed transaction message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *Server) SendTransaction(_ *http.Request, args *evm.SendTransactionArgs, _ *evm.SendTransactionReply) error {
	destBytes, err := hexutil.Decode(args.DestAddress)
	if err != nil {
		return errors2.Wrap(err, "error decoding Dest")
	}
	var dest common.Address
	copy(dest[:], destBytes)

	maxGas, valid := new(big.Int).SetString(args.MaxGas, 10)
	if !valid {
		return errors.New("invalid MaxGas")
	}

	gasPriceBid, valid := new(big.Int).SetString(args.GasPriceBid, 10)
	if !valid {
		return errors.New("invalid GasPriceBid")
	}

	sequenceNum, valid := new(big.Int).SetString(args.SequenceNum, 10)
	if !valid {
		return errors.New("invalid sequence num")
	}

	paymentInt, valid := new(big.Int).SetString(args.Payment, 10)
	if !valid {
		return errors.New("invalid Payment")
	}

	data, err := hexutil.Decode(args.Data)
	if err != nil {
		return errors2.Wrap(err, "error decoding data")
	}

	tx := message.Transaction{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		SequenceNum: sequenceNum,
		DestAddress: dest,
		Payment:     paymentInt,
		Data:        data,
	}

	pubkeyBytes, err := hexutil.Decode(args.Pubkey)
	if err != nil {
		return errors2.Wrap(err, "error decoding pubkey")
	}

	signature, err := hexutil.Decode(args.Signature)
	if err != nil {
		return errors2.Wrap(err, "error decoding signature")
	}

	return m.batch.SendTransaction(tx, pubkeyBytes, signature)
}

//FindLogs takes a set of parameters and return the list of all logs that match
//the query
func (m *Server) FindLogs(
	request *http.Request,
	args *evm.FindLogsArgs,
	reply *evm.FindLogsReply,
) error {
	addresses := make([]common.Address, 0, 1)
	for _, addr := range args.Addresses {
		addresses = append(addresses, common.HexToAddress(addr))
	}

	topicGroups := make([][]common.Hash, 0, len(args.TopicGroups))
	for _, topicGroup := range args.TopicGroups {
		topics := make([]common.Hash, 0, len(topicGroup.Topics))
		for _, topic := range topicGroup.Topics {
			topics = append(topics, common.NewHashFromEth(ethcommon.HexToHash(topic)))
		}
		topicGroups = append(topicGroups, topics)
	}

	var fromHeight *uint64
	if len(args.FromHeight) != 0 {
		intVal, err := strconv.ParseUint(args.FromHeight[2:], 16, 64)
		if err != nil {
			return errors2.Wrap(err, "bad from height")
		}
		fromHeight = &intVal
	}

	var toHeight *uint64
	if len(args.ToHeight) != 0 && args.ToHeight != "latest" {
		intVal, err := strconv.ParseUint(args.ToHeight[2:], 16, 64)
		if err != nil {
			return errors2.Wrap(err, "bad to height")
		}
		toHeight = &intVal
	}

	logs, err := m.db.FindLogs(request.Context(), fromHeight, toHeight, addresses, topicGroups)
	if err != nil {
		return err
	}

	logInfos := make([]*evm.FullLogBuf, 0, len(logs))
	for _, l := range logs {
		logInfos = append(logInfos, l.Marshal())
	}

	reply.Logs = logInfos
	return nil
}

func (m *Server) GetBlockCount(
	_ *http.Request,
	_ *evm.BlockCountArgs,
	reply *evm.BlockCountReply,
) error {
	id, err := m.db.LatestBlock()
	if err != nil {
		return err
	}
	reply.Height = id.Height.AsInt().Uint64()
	return nil
}

func (m *Server) GetOutputMessage(
	_ *http.Request,
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
//message with the given hash
func (m *Server) GetRequestResult(
	_ *http.Request,
	args *evm.GetRequestResultArgs,
	reply *evm.GetRequestResultReply,
) error {
	decoded, err := hexutil.Decode(args.TxHash)
	if err != nil {
		return err
	}
	var requestId common.Hash
	copy(requestId[:], decoded)
	index, startLogIndex, res, err := m.db.GetRequest(requestId)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := value.MarshalValue(res, &buf); err != nil {
		return err
	}
	reply.RawVal = hexutil.Encode(buf.Bytes())
	reply.Index = index
	reply.StartLogIndex = startLogIndex
	return nil
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetChainAddress(
	_ *http.Request,
	_ *evm.GetChainAddressArgs,
	reply *evm.GetChainAddressReply,
) error {
	reply.ChainAddress = m.chain.Hex()
	return nil
}

func (m *Server) BlockInfo(
	_ *http.Request,
	args *evm.BlockInfoArgs,
	reply *evm.BlockInfoReply,
) error {
	info, err := m.db.GetBlock(args.Height)
	if err != nil {
		return err
	}
	reply.Hash = info.Hash.String()
	reply.StartLog = info.StartLog
	reply.LogCount = info.LogCount
	reply.StartMessage = info.StartMessage
	reply.MessageCount = info.MessageCount
	reply.Bloom = info.Bloom.String()
	return nil
}

// CallMessage takes a request from a client to process in a temporary context
// and return the result
func (m *Server) Call(
	_ *http.Request,
	args *evm.CallMessageArgs,
	reply *evm.CallMessageReply,
) error {
	mach, blockId := m.db.CallInfo()
	ret, err := m.executeCall(mach, blockId, args)
	if err != nil {
		return err
	}
	reply.RawVal = ret
	return nil
}

// PendingCall takes a request from a client to process in a temporary context
// and return the result
func (m *Server) PendingCall(
	_ *http.Request,
	args *evm.CallMessageArgs,
	reply *evm.CallMessageReply,
) error {
	mach, blockId := m.db.CallInfo()
	ret, err := m.executeCall(mach, blockId, args)
	if err != nil {
		return err
	}
	reply.RawVal = ret
	return nil
}

func (m *Server) executeCall(mach machine.Machine, blockId *common.BlockId, args *evm.CallMessageArgs) (string, error) {
	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return "", err
	}

	var sender common.Address
	if len(args.Sender) > 0 {
		senderBytes, err := hexutil.Decode(args.Sender)
		if err != nil {
			return "", err
		}
		copy(sender[:], senderBytes)
	}

	seq, _ := new(big.Int).SetString("999999999999999999999999", 10)

	callMsg := message.NewCallFromData(dataBytes)
	inboxMsg := message.NewInboxMessage(
		message.L2Message{Msg: callMsg},
		sender,
		seq,
		message.ChainTime{
			BlockNum:  blockId.Height,
			Timestamp: big.NewInt(time.Now().Unix()),
		},
	)

	inbox := value.NewEmptyTuple()
	inbox = value.NewTuple2(inboxMsg.AsValue(), inbox)
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
		return "", fmt.Errorf("can't produce solution since machine is blocked %v", br)
	}

	log.Println("Executed call for", steps, "steps")

	results := assertion.ParseLogs()
	if len(results) == 0 {
		return "", errors.New("call produced no output")
	}
	lastLogVal := results[len(results)-1]
	lastLog, err := evm.NewResultFromValue(lastLogVal)
	if err != nil {
		return "", err
	}
	if lastLog.L1Message.MessageID() != inboxMsg.MessageID() {
		// Last produced log is not the call we sent
		return "", errors.New("call took too long to execute")
	}

	result := results[len(results)-1]
	var buf bytes.Buffer
	_ = value.MarshalValue(result, &buf) // error can only occur from writes and bytes.Buffer is safe
	return hexutil.Encode(buf.Bytes()), nil
}
