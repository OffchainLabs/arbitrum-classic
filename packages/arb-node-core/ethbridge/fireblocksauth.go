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
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks/accounttype"
	"github.com/pkg/errors"
)

type FireblocksTransactAuth struct {
	sync.Mutex
	auth   *bind.TransactOpts
	client ethutils.EthClient
	fb     *fireblocks.Fireblocks
}

func NewFireblocksTransactAuthAdvanced(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	walletConfig *configuration.Wallet,
	usePendingNonce bool,
) (TransactAuth, *fireblocks.Fireblocks, error) {
	err := getNonce(ctx, client, auth, usePendingNonce)
	if err != nil {
		return nil, nil, err
	}

	fb, err := fireblocks.New(walletConfig.Fireblocks)
	if err != nil {
		return nil, nil, err
	}

	transactAuth := &FireblocksTransactAuth{
		auth:   auth,
		client: client,
		fb:     fb,
	}

	if !walletConfig.Fireblocks.DisableHandlePending {
		// Handle any pending transactions left from last time process was running
		err = waitForPendingTransactions(ctx, client, transactAuth, fb)
		if err != nil {
			return nil, nil, err
		}
	}

	return transactAuth, fb, nil
}

func NewFireblocksTransactAuth(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	walletConfig *configuration.Wallet,
) (TransactAuth, *fireblocks.Fireblocks, error) {
	return NewFireblocksTransactAuthAdvanced(ctx, client, auth, walletConfig, true)
}

func waitForPendingTransactions(
	ctx context.Context,
	client ethutils.EthClient,
	transactAuth *FireblocksTransactAuth,
	fb *fireblocks.Fireblocks,
) error {
	for {
		pendingTx, err := fb.ListPendingTransactions()
		if err != nil {
			logger.Error().Err(err).Msg("error listing pending transactions")
			return err
		}

		if len(*pendingTx) == 0 {
			logger.Info().Msg("no pending fireblocks transactions to take care of")
			break
		}

		logger.Info().Int("count", len(*pendingTx)).Msg("pending fireblocks transactions need to be handled")
		for _, details := range *pendingTx {
			if details.Status == fireblocks.Broadcasting {
				logger.
					Info().
					Str("id", details.Id).
					Str("status", details.Status).
					Str("destination", details.DestinationAddress).
					Msg("retrying pending fireblocks transaction")
				// Existing transaction is stuck
				destinationAddress := ethcommon.HexToAddress(details.DestinationAddress)
				baseTx := &types.DynamicFeeTx{
					To:    &destinationAddress,
					Value: big.NewInt(details.Amount),
					Data:  []byte(details.ExtraParameters.ContractCallData),
				}
				rawTx := types.NewTx(baseTx)
				arbTx, err := NewFireblocksArbTransaction(rawTx, &details)
				if err != nil {
					logger.
						Error().
						Err(err).
						Str("id", details.Id).
						Msg("error creating new version of pending transaction")
					return err
				}
				_, err = WaitForReceiptWithResultsAndReplaceByFee(
					ctx,
					client,
					transactAuth.auth.From,
					arbTx,
					"waitForPendingTransactions",
					transactAuth,
					transactAuth,
				)
				if err != nil {
					return err
				}
			} else {
				logger.
					Info().
					Str("id", details.Id).
					Str("status", details.Status).
					Str("destination", details.DestinationAddress).
					Msg("waiting on pending fireblocks transaction")
			}
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(10 * time.Second):
		}
	}

	return nil
}
func (ta *FireblocksTransactAuth) TransactionReceipt(ctx context.Context, tx *ArbTransaction) (*types.Receipt, error) {
	details, err := ta.fb.GetTransaction(tx.Id())
	if err != nil {
		logger.
			Warn().
			Err(err).
			Str("hash", tx.Hash().String()).
			Msg("error getting fireblocks transaction for receipt")
		return nil, errors.Wrapf(err, "error getting fireblocks transaction for receipt: %s", details.Status)
	}
	if ta.fb.IsTransactionStatusFailed(details.Status) {
		logger.
			Error().
			Err(err).
			Str("To", details.DestinationAddress).
			Str("From", details.SourceAddress).
			Str("status", details.Status).
			Str("txhash", details.TxHash).
			Msg("fireblocks transaction failed when getting receipt")
		return nil, errors.Wrapf(err, "fireblocks transaction failed when getting receipt: %s", details.Status)
	}

	return ta.client.TransactionReceipt(ctx, tx.Hash())
}

func (ta *FireblocksTransactAuth) NonceAt(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (uint64, error) {
	return ta.client.NonceAt(ctx, account, blockNumber)
}

func (ta *FireblocksTransactAuth) SendTransaction(ctx context.Context, tx *types.Transaction, replaceTxByHash string) (*ArbTransaction, error) {
	input := fireblocks.CreateTransactionInput{
		DestinationType:     accounttype.OneTimeAddress,
		DestinationId:       tx.To().Hex(),
		DestinationTag:      "",
		AmountWei:           tx.Value(),
		GasLimitWei:         big.NewInt(int64(tx.Gas())),
		GasPriceWei:         tx.GasPrice(),
		MaxPriorityFeeWei:   tx.GasTipCap(),
		MaxTotalGasPriceWei: tx.GasFeeCap(),
		ReplaceTxByHash:     replaceTxByHash,
		CallData:            ethcommon.Bytes2Hex(tx.Data()),
	}
	txResponse, err := ta.fb.CreateContractCall(&input)
	if err != nil {
		return nil, err
	}
	if ta.fb.IsTransactionStatusFailed(txResponse.Status) {
		logger.
			Error().
			Hex("data", tx.Data()).
			Str("id", txResponse.Id).
			Str("status", txResponse.Status).
			Msg("fireblocks transaction failed")
		return nil, errors.New("fireblocks transaction failed")
	}
	logger.Debug().Hex("data", tx.Data()).Msg("sent transaction")

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("ctx done")
		case <-time.After(2 * time.Second):
		}

		details, err := ta.fb.GetTransaction(txResponse.Id)
		if err != nil {
			logger.
				Warn().
				Err(err).
				Hex("to", tx.To().Bytes()).
				Str("id", txResponse.Id).
				Str("status", details.Status).
				Msg("error getting fireblocks transaction")
			return nil, errors.Wrapf(err, "error getting fireblocks transaction: %s", details.Status)
		}

		if ta.fb.IsTransactionStatusFailed(details.Status) {
			logger.
				Error().
				Str("To", details.DestinationAddress).
				Str("From", details.SourceAddress).
				Str("status", details.Status).
				Str("txhash", details.TxHash).
				Str("replaceTxByHash", replaceTxByHash).
				Msg("fireblocks transaction failed")
			return nil, errors.New("fireblocks transaction failed")
		}

		if len(details.TxHash) > 0 {
			return NewFireblocksArbTransaction(tx, details)
		}

		// Hash not returned, keep trying
	}
}

func (ta *FireblocksTransactAuth) Sign(addr ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
	// Fireblocks handles signing, so nothing to do here
	return tx, nil
}

func (ta *FireblocksTransactAuth) getAuth() *bind.TransactOpts {
	return ta.auth
}

func (ta *FireblocksTransactAuth) From() ethcommon.Address {
	return ta.auth.From
}
