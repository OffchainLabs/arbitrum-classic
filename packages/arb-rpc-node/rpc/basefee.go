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

package rpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"
)

const adjustInterval = time.Minute * 60
const adjustMinPercentChange = 10
const adjustMinSampleSize = 4

func tryAdjustBaseFee(ctx context.Context, batcher *batcher.SequencerBatcher, arbAggregator *arboscontracts.ArbAggregator) error {
	newBaseFee, sampleSize := batcher.RecommendedBaseFee()
	if sampleSize < adjustMinSampleSize {
		logger.Debug().Int("sampleSize", sampleSize).Msg("not adjusting base fee as sample size is too low")
		return nil
	}
	auth := batcher.GetTransactAuth()
	aggAddr := auth.From()
	currentBaseFee, err := arbAggregator.GetTxBaseFee(&bind.CallOpts{Context: ctx}, aggAddr)
	if err != nil {
		return err
	}
	newBaseFeeBig := big.NewInt(int64(newBaseFee))
	diffPercent := new(big.Int).Sub(newBaseFeeBig, currentBaseFee)
	diffPercent.Mul(diffPercent, big.NewInt(100))
	diffPercent.Div(diffPercent, currentBaseFee)
	diffPercent.Abs(diffPercent)
	if diffPercent.Cmp(big.NewInt(int64(adjustMinPercentChange))) < 0 {
		logger.
			Debug().
			Str("currentBaseFee", currentBaseFee.String()).
			Str("newBaseFee", newBaseFeeBig.String()).
			Msg("not adjusting base fee as difference is too small")
		return nil
	}
	tx, err := transactauth.MakeTx(ctx, auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return arbAggregator.SetTxBaseFee(auth, aggAddr, newBaseFeeBig)
	})
	if err != nil {
		return err
	}
	logger.
		Info().
		Str("oldBaseFee", currentBaseFee.String()).
		Str("newBaseFee", newBaseFeeBig.String()).
		Str("txHash", tx.Hash().String()).
		Msg("adjusting sequencer aggregator base fee")
	return nil
}

func AutoAdjustBaseFee(ctx context.Context, batcher *batcher.SequencerBatcher, srv *aggregator.Server) {
	go func() {
		client := web3.NewEthClient(srv, false)
		var arbAggregator *arboscontracts.ArbAggregator
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(adjustInterval):
			}
			if arbAggregator == nil {
				var err error
				arbAggregator, err = arboscontracts.NewArbAggregator(arbos.ARB_AGGREGATOR_ADDRESS, client)
				if err != nil {
					logger.Error().Err(err).Msg("failed to bind to ArbAggregator")
					continue
				}
			}
			err := tryAdjustBaseFee(ctx, batcher, arbAggregator)
			if err != nil {
				logger.Error().Err(err).Msg("failed to adjust aggregator base fee")
			}
		}
	}()
}
