/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"
)

type ArbAddresses struct {
	ArbFactory string `json:"ArbFactory"`
}

var parityErr = "Block information is incomplete while ancient block sync is still in progress, before it's finished we can't determine the existence of requested item."
var parityErr2 = "missing required field 'transactionHash' for Log"

func (a ArbAddresses) ArbFactoryAddress() common.Address {
	return common.NewAddressFromEth(ethcommon.HexToAddress(a.ArbFactory))
}

// RBF = Replace By Fee
const rbfInterval time.Duration = time.Minute * 5

type attemptRbfInfo struct {
	attempt func() (ethcommon.Hash, error)
	account ethcommon.Address
	nonce   uint64
}

func waitForReceiptWithResultsSimpleInternal(ctx context.Context, client ethutils.ReceiptFetcher, txHash ethcommon.Hash, rbfInfo *attemptRbfInfo, fb *fireblocks.Fireblocks) (*types.Receipt, error) {
	lastRbf := time.Now()
	for {
		select {
		case <-time.After(time.Second):
			if rbfInfo != nil && time.Since(lastRbf) >= rbfInterval {
				newTxHash, err := rbfInfo.attempt()
				lastRbf = time.Now()
				if err == nil {
					txHash = newTxHash
				} else {
					logger.Warn().Err(err).Msg("failed to replace by fee")
				}
			}
			if fb != nil {
			}
			receipt, err := client.TransactionReceipt(ctx, txHash)
			if receipt == nil {
				if rbfInfo != nil {
					// an alternative tx might've gotten confirmed
					nonce, err := client.NonceAt(ctx, rbfInfo.account, nil)
					if err == nil {
						if nonce >= rbfInfo.nonce {
							return nil, nil
						}
					} else {
						logger.Warn().Err(err).Str("account", rbfInfo.account.String()).Msg("Issue getting pending nonce")
					}
				}
				continue
			}
			if err != nil {
				if err.Error() == ethereum.NotFound.Error() || err.Error() == parityErr {
					continue
				}

				if err.Error() == parityErr2 {
					logger.Warn().Err(err).Hex("tx", txHash.Bytes()).Msg("Issue getting receipt")
					continue
				}
				logger.Error().Err(err).Hex("tx", txHash.Bytes()).Msg("Issue getting receipt")
				return nil, errors.WithStack(err)
			}
			return receipt, nil
		case <-ctx.Done():
			return nil, errors.Errorf("receipt not found")
		}
	}
}

func WaitForReceiptWithResultsSimple(ctx context.Context, client ethutils.ReceiptFetcher, txHash ethcommon.Hash, fb *fireblocks.Fireblocks) (*types.Receipt, error) {
	return waitForReceiptWithResultsSimpleInternal(ctx, client, txHash, nil, fb)
}

func increaseByPercent(original *big.Int, percentage int64) *big.Int {
	threshold := new(big.Int).Mul(original, big.NewInt(100+percentage))
	threshold.Div(threshold, big.NewInt(100))
	return threshold
}

func WaitForReceiptWithResultsAndReplaceByFee(ctx context.Context, client ethutils.EthClient, from ethcommon.Address, tx *types.Transaction, methodName string, auth *TransactAuth, fb *fireblocks.Fireblocks) (*types.Receipt, error) {
	var rbfInfo *attemptRbfInfo
	if auth != nil {
		attemptRbf := func() (ethcommon.Hash, error) {
			auth := auth.getAuth(ctx)
			if auth.GasPrice.Cmp(tx.GasPrice()) <= 0 {
				return tx.Hash(), nil
			}
			var rawTx *types.Transaction
			if tx.Type() == types.DynamicFeeTxType {
				block, err := client.HeaderByNumber(ctx, nil)
				if err != nil {
					return ethcommon.Hash{}, err
				}
				if block.BaseFee == nil {
					return ethcommon.Hash{}, errors.New("attempted to use dynamic fee tx in pre-EIP-1559 block")
				}
				tipCap, err := client.SuggestGasTipCap(ctx)
				if err != nil {
					return ethcommon.Hash{}, err
				}
				if tipCap.Cmp(increaseByPercent(tx.GasTipCap(), 10)) < 0 {
					// We only replace by fee when we'd increase the tip by 10%
					return tx.Hash(), nil
				}
				feeCap := new(big.Int).Mul(block.BaseFee, big.NewInt(2))
				feeCap.Add(feeCap, tipCap)
				minFeeCap := increaseByPercent(tx.GasFeeCap(), 10)
				if feeCap.Cmp(minFeeCap) < 0 {
					feeCap = minFeeCap
				}
				baseTx := &types.DynamicFeeTx{
					ChainID:    tx.ChainId(),
					Nonce:      tx.Nonce(),
					GasTipCap:  tipCap,
					GasFeeCap:  feeCap,
					Gas:        tx.Gas(),
					To:         tx.To(),
					Value:      tx.Value(),
					Data:       tx.Data(),
					AccessList: tx.AccessList(),
				}
				rawTx = types.NewTx(baseTx)
			} else {
				gasPrice, err := client.SuggestGasPrice(ctx)
				if err != nil {
					return ethcommon.Hash{}, err
				}
				if gasPrice.Cmp(increaseByPercent(tx.GasPrice(), 10)) < 0 {
					// We only replace by fee when we'd increase the fee by at least 10%
					return tx.Hash(), nil
				}
				baseTx := &types.LegacyTx{
					Nonce:    tx.Nonce(),
					GasPrice: gasPrice,
					Gas:      tx.Gas(),
					To:       tx.To(),
					Value:    tx.Value(),
					Data:     tx.Data(),
				}
				rawTx = types.NewTx(baseTx)
			}
			newTx, err := auth.Signer(auth.From, rawTx)
			if err != nil {
				return ethcommon.Hash{}, err
			}
			err = client.SendTransaction(ctx, newTx)
			if err != nil {
				return ethcommon.Hash{}, err
			}
			*tx = *newTx
			return newTx.Hash(), nil
		}
		rbfInfo = &attemptRbfInfo{
			attempt: attemptRbf,
			account: auth.auth.From,
			nonce:   tx.Nonce(),
		}
	}
	receipt, err := waitForReceiptWithResultsSimpleInternal(ctx, client, tx.Hash(), rbfInfo, fb)
	if err != nil {
		logger.Warn().Err(err).Hex("tx", tx.Hash().Bytes()).Msg("error while waiting for transaction receipt")
		return nil, errors.WithStack(err)
	}
	if receipt != nil && receipt.Status != 1 {
		logger.Warn().Hex("tx", tx.Hash().Bytes()).Msg("failed transaction")
		callMsg := ethereum.CallMsg{
			From:      from,
			To:        tx.To(),
			Gas:       tx.Gas(),
			GasTipCap: tx.GasTipCap(),
			GasFeeCap: tx.GasFeeCap(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		}
		data, err := client.CallContract(ctx, callMsg, receipt.BlockNumber)
		if err != nil {
			return nil, errors.Wrapf(err, "transaction %v failed", methodName)
		}
		return nil, errors.Errorf("transaction %v failed with tx %v", methodName, string(data))
	}
	return receipt, nil
}

func WaitForReceiptWithResults(ctx context.Context, client ethutils.EthClient, from ethcommon.Address, tx *types.Transaction, methodName string, fb *fireblocks.Fireblocks) (*types.Receipt, error) {
	return WaitForReceiptWithResultsAndReplaceByFee(ctx, client, from, tx, methodName, nil, fb)
}
