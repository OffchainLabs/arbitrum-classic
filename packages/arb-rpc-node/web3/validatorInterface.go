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
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

var lookupMessageBatchProof abi.Method

func init() {
	parsedABI, err := abi.JSON(strings.NewReader(arboscontracts.NodeInterfaceABI))
	if err != nil {
		panic(err)
	}
	lookupMessageBatchProof = parsedABI.Methods["lookupMessageBatchProof"]
}

func HandleNodeInterfaceCall(srv *aggregator.Server, calldata []byte) ([]byte, error) {
	if len(calldata) < 4 {
		return nil, errors.New("calldata too short")
	}
	funcID := calldata[:4]
	if bytes.Equal(funcID, lookupMessageBatchProof.ID) {
		return handleLookupMessageBatch(srv, calldata)
	}
	return nil, errors.New("invalid function")
}

func handleLookupMessageBatch(srv *aggregator.Server, calldata []byte) ([]byte, error) {
	inputs, err := lookupMessageBatchProof.Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, err
	}
	batchNum := inputs[0].(*big.Int)
	index := inputs[0].(uint64)
	proof, err := srv.GetL2ToL1Proof(batchNum, index)
	if err != nil {
		return nil, err
	}
	res, err := evm.NewSendResultFromData(proof.Data)
	if err != nil {
		return nil, err
	}
	txRes, ok := res.(*evm.L2ToL1TxResult)
	if !ok {
		return nil, errors.New("unexepected result type")
	}

	return lookupMessageBatchProof.Outputs.PackValues([]interface{}{
		common.NewEthHashesFromHashes(proof.Nodes),
		protocol.PathSliceToInt(proof.Path),
		txRes.L2Sender.ToEthAddress(),
		txRes.L1Dest.ToEthAddress(),
		txRes.L2Block,
		txRes.L1Block,
		txRes.Timestamp,
		txRes.Value,
		txRes.Calldata,
	})
}
