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
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/batcher"
	errors2 "github.com/pkg/errors"
	"math/big"
	"net/http"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type Server struct {
	batch *batcher.Batcher
}

// NewServer returns a new instance of the Server class
func NewServer(
	ctx context.Context,
	globalInbox arbbridge.GlobalInbox,
	rollupAddress common.Address,
) *Server {
	return &Server{batch: batcher.NewBatcher(ctx, globalInbox, rollupAddress)}
}

// SendTransaction takes a request signed transaction message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *Server) SendTransaction(_ *http.Request, args *SendTransactionArgs, _ *SendTransactionReply) error {
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
