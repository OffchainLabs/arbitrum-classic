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
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"net/http"
	"sort"
	"sync"
	"time"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

const maxTransactions = 200

const signatureLength = 65

type DecodedBatchTx struct {
	tx     message.BatchTx
	sender common.Address
}

type Server struct {
	rollupAddress common.Address
	globalInbox   arbbridge.GlobalInbox

	sync.Mutex
	valid        bool
	transactions []DecodedBatchTx
}

// NewServer returns a new instance of the Server class
func NewServer(
	ctx context.Context,
	globalInbox arbbridge.GlobalInbox,
	rollupAddress common.Address,
) *Server {
	server := &Server{
		rollupAddress: rollupAddress,
		globalInbox:   globalInbox,
		valid:         true,
	}

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				server.Lock()
				// Keep sending in spin loop until we can't anymore
				sentFull := false
				for server.valid && len(server.transactions) >= maxTransactions {
					server.sendBatch(ctx)
					sentFull = true
				}
				// If we have've sent any batches, send a partial
				if !sentFull && server.valid && len(server.transactions) > 0 {
					server.sendBatch(ctx)
				}
				server.Unlock()
			}
		}
	}()
	return server
}

// prepareTransactions reorders the transactions such that the position of each
// user is maintained, but the transactions of that user are swapped to be in
// sequence number order
func prepareTransactions(txes []DecodedBatchTx) message.TransactionBatch {
	transactionsBySender := make(map[common.Address][]DecodedBatchTx)
	for _, tx := range txes {
		transactionsBySender[tx.sender] = append(transactionsBySender[tx.sender], tx)
	}

	for _, txes := range transactionsBySender {
		sort.SliceStable(txes, func(i, j int) bool {
			return txes[i].tx.Transaction.SequenceNum.Cmp(txes[j].tx.Transaction.SequenceNum) < 0
		})
	}

	batchTxes := make([]message.BatchTx, 0, len(txes))
	for _, tx := range txes {
		nextTx := transactionsBySender[tx.sender][0]
		transactionsBySender[tx.sender] = transactionsBySender[tx.sender][1:]
		batchTxes = append(batchTxes, nextTx.tx)
	}
	return message.TransactionBatch{Transactions: batchTxes}
}

func (m *Server) sendBatch(ctx context.Context) {
	var txes []DecodedBatchTx

	if len(m.transactions) > maxTransactions {
		txes = m.transactions[:maxTransactions]
		m.transactions = m.transactions[maxTransactions:]
	} else {
		txes = m.transactions
		m.transactions = nil
	}
	m.Unlock()

	log.Println("Submitting batch with", len(txes), "transactions")

	err := m.globalInbox.SendL2MessageNoWait(
		ctx,
		m.rollupAddress,
		message.L2Message{Msg: prepareTransactions(txes)},
	)

	m.Lock()
	if err != nil {
		log.Println("transaction aggregator failed: ", err)
		m.valid = false
	}
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

	signature, err := hexutil.Decode(args.Signature)
	if err != nil {
		return errors2.Wrap(err, "error decoding signature")
	}
	if len(signature) != signatureLength {
		return errors.New("signature of wrong length")
	}

	chainId := new(big.Int).SetBytes(m.rollupAddress[12:])
	signer := types.NewEIP155Signer(chainId)
	signedTx, err := tx.AsEthTx().WithSignature(signer, signature)
	if err != nil {
		return err
	}
	ethSender, err := signer.Sender(signedTx)
	if err != nil {
		return err
	}
	sender := common.NewAddressFromEth(ethSender)

	var sigData [signatureLength]byte
	copy(sigData[:], signature)

	batchTx := message.BatchTx{
		Transaction: tx,
		Signature:   sigData,
	}

	log.Println("Got tx: ", tx, "with hash", tx.MessageID(sender), "from", sender)

	m.Lock()
	defer m.Unlock()

	if !m.valid {
		return errors.New("tx aggregator is not running")
	}

	m.transactions = append(m.transactions, DecodedBatchTx{
		tx:     batchTx,
		sender: sender,
	})
	return nil
}
