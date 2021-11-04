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

package web3

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"math/big"
)

type Debug struct {
	srv *Server
}

const TraceBlockLookback = 4

func (d *Debug) TraceTransaction(txHash hexutil.Bytes) (interface{}, error) {
	res, _, _, logNumber, err := d.srv.getTransactionInfoByHash(txHash)
	if err != nil || res == nil {
		return nil, err
	}

	blockNumber := res.IncomingRequest.L2BlockNumber.Uint64()
	cursor, err := d.srv.srv.GetExecutionCursorAtBlock(blockNumber-TraceBlockLookback, true)
	if err != nil {
		return nil, err
	}

	debugPrints, err := d.srv.srv.AdvanceExecutionCursorWithTracing(cursor, big.NewInt(100000000000), true, true, logNumber)
	if err != nil {
		return nil, err
	}

	trace, err := evm.GetTrace(debugPrints)
	_ = trace

	return nil, nil
}
