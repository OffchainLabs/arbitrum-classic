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

package observer

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"math/big"
)

func CalculateCatchupFetch(ctx context.Context, start *big.Int, clnt arbbridge.ChainTimeGetter, maxReorg *big.Int) (*big.Int, error) {
	currentLocalHeight := start
	currentOnChain, err := clnt.BlockIdForHeight(ctx, nil)
	if err != nil {
		return nil, err
	}
	currentL1Height := currentOnChain.Height.AsInt()

	fastCatchupEndHeight := new(big.Int).Sub(currentL1Height, maxReorg)
	if currentLocalHeight.Cmp(fastCatchupEndHeight) >= 0 {
		return nil, nil
	}

	fetchSize := new(big.Int).Sub(fastCatchupEndHeight, currentLocalHeight)
	if fetchSize.Cmp(big.NewInt(1)) <= 0 {
		return nil, nil
	}
	if fetchSize.Cmp(maxReorg) >= 0 {
		fetchSize = maxReorg
	}
	fetchEnd := new(big.Int).Add(currentLocalHeight, fetchSize)
	fetchEnd = fetchEnd.Sub(fetchEnd, big.NewInt(1))

	if new(big.Int).Sub(fetchEnd, start).Cmp(big.NewInt(10)) < 0 {
		// If the remaining safe fetch amount is too small, exit fast catchup mode
		return nil, nil
	}
	return fetchEnd, nil
}
