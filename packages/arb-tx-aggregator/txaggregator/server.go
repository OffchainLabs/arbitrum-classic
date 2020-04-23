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

package txaggregator

import (
	"context"
	"errors"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

const maxTransactions = 50

type Server struct {
	rollupAddress common.Address
	globalInbox   arbbridge.GlobalInbox

	sync.Mutex
	valid        bool
	transactions []message.Transaction
	signatures   [][65]byte
}

// NewServer returns a new instance of the Server class
func NewServer(ctx context.Context, globalInbox arbbridge.GlobalInbox, rollupAddress common.Address) *Server {
	server := &Server{rollupAddress: rollupAddress, globalInbox: globalInbox, valid: true}

	go func() {
		ticker := time.NewTicker(common.NewTimeBlocksInt(30).Duration())
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				server.sendBatch(ctx)
			}
		}
	}()
	return server
}

func (m *Server) sendBatch(ctx context.Context) {
	var txes []message.Transaction
	var sigs [][65]byte

	m.Lock()
	if !m.valid {
		return
	}

	if len(m.transactions) > maxTransactions {
		txes = m.transactions[:maxTransactions]
		sigs = m.signatures[:maxTransactions]
		m.transactions = m.transactions[maxTransactions:]
		m.signatures = m.signatures[maxTransactions:]
	} else {
		txes = m.transactions
		m.transactions = nil
		sigs = m.signatures
		m.signatures = nil
	}
	m.Unlock()

	err := m.globalInbox.DeliverTransactionBatch(ctx, m.rollupAddress, txes, sigs)
	if err != nil {
		log.Println("Transaction aggregator failed: ", err)
		m.Lock()
		m.valid = false
		m.Unlock()
	}
}

// CallMessage takes a request from a client to process in a temporary context and return the result
func (m *Server) SendTransaction(ctx context.Context, args *SendTransactionArgs) (*SendTransactionReply, error) {
	toBytes, err := hexutil.Decode(args.To)
	if err != nil {
		return nil, err
	}
	var to common.Address
	copy(toBytes[:], toBytes)

	sequenceNumBytes, err := hexutil.Decode(args.SequenceNum)
	if err != nil {
		return nil, err
	}
	sequenceNum := new(big.Int).SetBytes(sequenceNumBytes[:])

	valueBytes, err := hexutil.Decode(args.Value)
	if err != nil {
		return nil, err
	}
	valueInt := new(big.Int).SetBytes(valueBytes[:])

	data, err := hexutil.Decode(args.Data)
	if err != nil {
		return nil, err
	}

	pubkey, err := hexutil.Decode(args.Pubkey)
	if err != nil {
		return nil, err
	}

	signature, err := hexutil.Decode(args.Data)
	if err != nil {
		return nil, err
	}

	if len(signature) != 65 {
		return nil, errors.New("SendTransaction: Signature of wrong length")
	}

	// Convert sig with normalized v
	if signature[64] == 27 {
		signature[64] = 0
	} else if signature[64] == 28 {
		signature[64] = 1
	}

	txDataHash := hashing.SoliditySHA3(
		hashing.Address(m.rollupAddress),
		hashing.Address(to),
		hashing.Uint256(sequenceNum),
		hashing.Uint256(valueInt),
		data,
	)

	messageHash := hashing.SoliditySHA3WithPrefix(txDataHash[:])
	if !crypto.VerifySignature(pubkey, messageHash[:], signature) {
		return nil, errors.New("Invalid signature")
	}

	m.Lock()
	defer m.Unlock()

	if !m.valid {
		return nil, errors.New("Tx aggregator is not running")
	}

	m.transactions = append(m.transactions, message.Transaction{
		To:          to,
		SequenceNum: sequenceNum,
		Value:       valueInt,
		Data:        data,
	})

	var sigData [65]byte
	copy(sigData[:], signature)
	m.signatures = append(m.signatures, sigData)
	return &SendTransactionReply{}, nil
}
