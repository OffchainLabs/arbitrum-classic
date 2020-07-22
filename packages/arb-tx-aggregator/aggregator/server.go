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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	errors2 "github.com/pkg/errors"
	"math/big"
	"net/http"
	"strconv"
)

type RPCServer struct {
	srv *Server
}

// NewServer returns a new instance of the Server class
func NewRPCServer(srv *Server) *RPCServer {
	return &RPCServer{srv: srv}
}

// SendTransaction takes a request signed transaction l2message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *RPCServer) SendTransaction(request *http.Request, args *evm.SendTransactionArgs, reply *evm.SendTransactionReply) error {
	encodedTx, err := hexutil.Decode(args.SignedTransaction)
	if err != nil {
		return errors2.Wrap(err, "error decoding signed transaction")
	}

	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(encodedTx, tx); err != nil {
		return err
	}
	hash, err := m.srv.SendTransaction(request.Context(), tx)
	if err != nil {
		return err
	}
	reply.TransactionHash = hash.String()
	return nil
}

//FindLogs takes a set of parameters and return the list of all logs that match
//the query
func (m *RPCServer) FindLogs(
	request *http.Request,
	args *evm.FindLogsArgs,
	reply *evm.FindLogsReply,
) error {
	addresses := make([]ethcommon.Address, 0, len(args.Addresses))
	for _, addr := range args.Addresses {
		addresses = append(addresses, ethcommon.HexToAddress(addr))
	}

	topicGroups := make([][]ethcommon.Hash, 0, len(args.TopicGroups))
	for _, topicGroup := range args.TopicGroups {
		topics := make([]ethcommon.Hash, 0, len(topicGroup.Topics))
		for _, topic := range topicGroup.Topics {
			topics = append(topics, ethcommon.HexToHash(topic))
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
	logs, err := m.srv.FindLogs(request.Context(), fromHeight, toHeight, addresses, topicGroups)
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

func (m *RPCServer) GetBlockCount(
	request *http.Request,
	_ *evm.BlockCountArgs,
	reply *evm.BlockCountReply,
) error {
	var err error
	reply.Height, err = m.srv.GetBlockCount(request.Context())
	return err
}

func (m *RPCServer) GetOutputMessage(
	request *http.Request,
	args *evm.GetOutputMessageArgs,
	reply *evm.GetOutputMessageReply,
) error {
	return m.srv.GetOutputMessage(request.Context(), args, reply)
}

// GetMessageResult returns the value output by the VM in response to the
//l2message with the given hash
func (m *RPCServer) GetRequestResult(
	request *http.Request,
	args *evm.GetRequestResultArgs,
	reply *evm.GetRequestResultReply,
) error {
	decoded, err := hexutil.Decode(args.TxHash)
	if err != nil {
		return err
	}
	var requestId common.Hash
	copy(requestId[:], decoded)
	val, err := m.srv.GetRequestResult(request.Context(), requestId)
	if err != nil {
		// Request was not found so return nil rawVal
		reply.RawVal = ""
		return nil
	}
	var buf bytes.Buffer
	if err := value.MarshalValue(val, &buf); err != nil {
		return err
	}
	reply.RawVal = hexutil.Encode(buf.Bytes())
	return nil
}

// GetVMInfo returns current metadata about this VM
func (m *RPCServer) GetChainAddress(
	request *http.Request,
	args *evm.GetChainAddressArgs,
	reply *evm.GetChainAddressReply,
) error {
	chain, err := m.srv.GetChainAddress(request.Context())
	if err != nil {
		return err
	}
	reply.ChainAddress = chain.Hex()
	return nil
}

func (m *RPCServer) BlockInfo(
	request *http.Request,
	args *evm.BlockInfoArgs,
	reply *evm.BlockInfoReply,
) error {
	info, err := m.srv.BlockInfo(request.Context(), args.Height)
	if err != nil {
		return err
	}
	reply.Hash = info.Hash.String()
	reply.StartLog = info.StartLog
	reply.LogCount = info.LogCount
	reply.StartMessage = info.StartMessage
	reply.MessageCount = info.MessageCount
	reply.Bloom = hexutil.Encode(info.Bloom.Bytes())
	return nil
}

// Call takes a request from a client to process in a temporary context
// and return the result
func (m *RPCServer) Call(
	request *http.Request,
	args *evm.CallMessageArgs,
	reply *evm.CallMessageReply,
) error {
	return m.callImpl(request, args, reply, m.srv.Call)
}

// PendingCall takes a request from a client to process in a temporary context
// and return the result
func (m *RPCServer) PendingCall(
	request *http.Request,
	args *evm.CallMessageArgs,
	reply *evm.CallMessageReply,
) error {
	return m.callImpl(request, args, reply, m.srv.PendingCall)
}

func (m *RPCServer) callImpl(
	request *http.Request,
	args *evm.CallMessageArgs,
	reply *evm.CallMessageReply,
	call func(ctx context.Context, msg l2message.Call, sender ethcommon.Address) (value.Value, error),
) error {
	var sender ethcommon.Address
	if len(args.Sender) > 0 {
		sender = ethcommon.HexToAddress(args.Sender)
	}
	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return err
	}

	callMsg := l2message.NewCallFromData(dataBytes)
	if callMsg.MaxGas.Cmp(big.NewInt(0)) == 0 || callMsg.MaxGas.Cmp(m.srv.maxCallGas) > 0 {
		callMsg.MaxGas = m.srv.maxCallGas
	}
	val, err := call(request.Context(), callMsg, sender)
	var buf bytes.Buffer
	_ = value.MarshalValue(val, &buf) // error can only occur from writes and bytes.Buffer is safe
	reply.RawVal = hexutil.Encode(buf.Bytes())
	return nil
}
