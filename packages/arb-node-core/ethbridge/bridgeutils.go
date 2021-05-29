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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/pkg/errors"
)

type BridgeUtils struct {
	con             *ethbridgecontracts.BridgeUtils
	bridgeAddresses [2]ethcommon.Address
}

func NewBridgeUtils(address ethcommon.Address, client ethutils.EthClient, delayedBridge *DelayedBridgeWatcher, sequencerInbox *SequencerInboxWatcher) (*BridgeUtils, error) {
	con, err := ethbridgecontracts.NewBridgeUtils(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &BridgeUtils{
		con:             con,
		bridgeAddresses: [2]ethcommon.Address{delayedBridge.address, sequencerInbox.address},
	}, nil
}

type CountAndAccumulator struct {
	Count       *big.Int
	Accumulator common.Hash
}

func (b *BridgeUtils) GetCountsAndAccumulators(ctx context.Context) (delayedRet, seqRet CountAndAccumulator, err error) {
	contractRet, contractErr := b.con.GetCountsAndAccumulators(&bind.CallOpts{Context: ctx}, b.bridgeAddresses)
	err = contractErr
	if err != nil {
		return
	}
	delayedRet.Count = contractRet.Counts[0]
	delayedRet.Accumulator = contractRet.Accs[0]
	seqRet.Count = contractRet.Counts[1]
	seqRet.Accumulator = contractRet.Accs[1]
	return
}
