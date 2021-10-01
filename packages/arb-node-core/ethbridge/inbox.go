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
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"
)

var l2MessageFromOriginCallABI abi.Method

func init() {
	parsedABI, err := abi.JSON(strings.NewReader(ethbridgecontracts.InboxABI))
	if err != nil {
		panic(err)
	}
	l2MessageFromOriginCallABI = parsedABI.Methods["sendL2MessageFromOrigin"]
}

type StandardInboxWatcher struct {
	con     *ethbridgecontracts.Inbox
	address ethcommon.Address
	client  ethutils.EthClient
}

func NewStandardInboxWatcher(address ethcommon.Address, client ethutils.EthClient) (*StandardInboxWatcher, error) {
	con, err := ethbridgecontracts.NewInbox(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &StandardInboxWatcher{
		con:     con,
		address: address,
		client:  client,
	}, nil
}

func (r *StandardInboxWatcher) fillMessageDetails(
	ctx context.Context,
	messageNums []*big.Int,
	txData map[string]*types.Transaction,
	messages map[string][]byte,
	minBlockNum, maxBlockNum uint64,
) error {
	msgQuery := make([]ethcommon.Hash, 0, len(messageNums))
	for _, messageNum := range messageNums {
		var msgNumBytes ethcommon.Hash
		copy(msgNumBytes[:], math.U256Bytes(messageNum))
		msgQuery = append(msgQuery, msgNumBytes)
	}

	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: new(big.Int).SetUint64(minBlockNum),
		// Not sure whether this is inclusive or exclusive so adding 1 just in case
		ToBlock:   new(big.Int).SetUint64(maxBlockNum),
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{inboxMessageDeliveredID, inboxMessageFromOriginID}, msgQuery},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, ethLog := range logs {
		msgNum, msg, err := r.parseMessage(txData, ethLog)
		if err != nil {
			return err
		}
		messages[string(msgNum.Bytes())] = msg
	}
	return nil
}

func (r *StandardInboxWatcher) parseMessage(txData map[string]*types.Transaction, ethLog types.Log) (*big.Int, []byte, error) {
	if ethLog.Topics[0] == inboxMessageDeliveredID {
		parsedLog, err := r.con.ParseInboxMessageDelivered(ethLog)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		return parsedLog.MessageNum, parsedLog.Data, nil
	} else if ethLog.Topics[0] == inboxMessageFromOriginID {
		parsedLog, err := r.con.ParseInboxMessageDeliveredFromOrigin(ethLog)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		tx, ok := txData[string(parsedLog.MessageNum.Bytes())]
		if !ok {
			return nil, nil, errors.New("didn't have tx data")
		}
		args := make(map[string]interface{})
		err = l2MessageFromOriginCallABI.Inputs.UnpackIntoMap(args, tx.Data()[4:])
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		return parsedLog.MessageNum, args["messageData"].([]byte), nil
	} else {
		return nil, nil, errors.New("unexpected log type")
	}
}

type StandardInbox struct {
	*StandardInboxWatcher
	auth transactauth.TransactAuth
}

func NewStandardInbox(address ethcommon.Address, client ethutils.EthClient, auth transactauth.TransactAuth) (*StandardInbox, error) {
	watcher, err := NewStandardInboxWatcher(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &StandardInbox{
		StandardInboxWatcher: watcher,
		auth:                 auth,
	}, nil
}

func (s *StandardInbox) Sender() common.Address {
	return common.NewAddressFromEth(s.auth.From())
}

func (s *StandardInbox) SendL2MessageFromOrigin(ctx context.Context, data []byte) (*arbtransaction.ArbTransaction, error) {
	return transactauth.MakeTx(ctx, s.auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return s.con.SendL2MessageFromOrigin(auth, data)
	})
}

func AddSequencerL2BatchFromOrigin(
	ctx context.Context,
	inbox *ethbridgecontracts.SequencerInbox,
	auth transactauth.TransactAuth,
	transactions []byte,
	lengths []*big.Int,
	sectionsMetadata []*big.Int,
	afterAcc [32]byte,
) (*arbtransaction.ArbTransaction, error) {
	arbTx, err := transactauth.MakeTx(ctx, auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return inbox.AddSequencerL2BatchFromOrigin(auth, transactions, lengths, sectionsMetadata, afterAcc)
	})
	if err != nil {
		return nil, err
	}
	return arbTx, nil
}

// AddSequencerL2BatchFromOriginCustomNonce is like AddSequencerL2BatchFromOrigin but with a custom nonce that will
// be incremented on success.  This is to handle the case when a stuck transaction is present on startup.
func AddSequencerL2BatchFromOriginCustomNonce(
	ctx context.Context,
	inbox *ethbridgecontracts.SequencerInbox,
	auth transactauth.TransactAuth,
	nonce *big.Int,
	transactions []byte,
	lengths []*big.Int,
	sectionsMetadata []*big.Int,
	afterAcc [32]byte,
	gasRefunder ethcommon.Address,
	gasRefunderExtraGas uint64,
) (*arbtransaction.ArbTransaction, error) {
	rawAuth := auth.GetAuth(ctx)
	arbTx, err := transactauth.MakeTxCustomNonce(ctx, auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		if gasRefunder != (ethcommon.Address{}) {
			tx, err := inbox.AddSequencerL2BatchFromOriginWithGasRefunder(auth, transactions, lengths, sectionsMetadata, afterAcc, gasRefunder)
			if err != nil {
				return nil, err
			}
			newGasLimit := tx.Gas() + gasRefunderExtraGas
			if tx.Type() == types.DynamicFeeTxType {
				tx = types.NewTx(&types.DynamicFeeTx{
					ChainID:    tx.ChainId(),
					Nonce:      tx.Nonce(),
					GasTipCap:  tx.GasTipCap(),
					GasFeeCap:  tx.GasFeeCap(),
					Gas:        newGasLimit,
					To:         tx.To(),
					Value:      tx.Value(),
					Data:       tx.Data(),
					AccessList: tx.AccessList(),
				})
			} else {
				tx = types.NewTx(&types.LegacyTx{
					Nonce:    tx.Nonce(),
					GasPrice: tx.GasPrice(),
					Gas:      newGasLimit,
					To:       tx.To(),
					Value:    tx.Value(),
					Data:     tx.Data(),
				})
			}
			return auth.Signer(auth.From, tx)
		} else {
			return inbox.AddSequencerL2BatchFromOrigin(auth, transactions, lengths, sectionsMetadata, afterAcc)
		}
	}, nonce)
	if err != nil {
		return nil, err
	}
	nonce.Add(nonce, big.NewInt(1))
	if rawAuth.Nonce.Cmp(nonce) < 0 {
		rawAuth.Nonce.Set(nonce)
	}

	return arbTx, nil
}
