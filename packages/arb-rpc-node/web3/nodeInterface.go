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
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

var lookupMessageBatchProof abi.Method
var estimateRetryableTicket abi.Method

func init() {
	parsedABI, err := abi.JSON(strings.NewReader(arboscontracts.NodeInterfaceABI))
	if err != nil {
		panic(err)
	}
	lookupMessageBatchProof = parsedABI.Methods["lookupMessageBatchProof"]
	estimateRetryableTicket = parsedABI.Methods["estimateRetryableTicket"]
}

func HandleNodeInterfaceCall(ctx context.Context, srv *Server, calldata []byte, blockNum rpc.BlockNumberOrHash) ([]byte, error) {
	if len(calldata) < 4 {
		return nil, errors.New("calldata too short")
	}
	funcID := calldata[:4]
	if bytes.Equal(funcID, lookupMessageBatchProof.ID) {
		return handleLookupMessageBatch(srv.srv, calldata)
	} else if bytes.Equal(funcID, estimateRetryableTicket.ID) {
		return handleEstimateRetryableTicket(ctx, srv, calldata, blockNum)
	}
	return nil, errors.New("invalid function")
}

func handleLookupMessageBatch(srv *aggregator.Server, calldata []byte) ([]byte, error) {
	inputs, err := lookupMessageBatchProof.Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, err
	}
	batchNum := inputs[0].(*big.Int)
	index := inputs[1].(uint64)
	proof, err := srv.GetL2ToL1Proof(batchNum, index)
	if err != nil {
		return nil, err
	}
	res, err := evm.NewVirtualSendResultFromData(proof.Data)
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

func handleEstimateRetryableTicket(ctx context.Context, srv *Server, calldata []byte, blockNum rpc.BlockNumberOrHash) ([]byte, error) {
	inputs, err := estimateRetryableTicket.Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, err
	}
	sender := inputs[0].(ethcommon.Address)
	deposit := inputs[1].(*big.Int)
	destAddr := inputs[2].(ethcommon.Address)
	l2CallValue := inputs[3].(*big.Int)
	maxSubmissionCost := inputs[4].(*big.Int)
	excessFeeRefundAddress := inputs[5].(ethcommon.Address)
	callValueRefundAddress := inputs[6].(ethcommon.Address)
	maxGas := inputs[7].(*big.Int)
	gasPriceBid := inputs[8].(*big.Int)
	data := inputs[9].([]byte)

	snap, err := srv.getSnapshotForNumberOrHash(ctx, blockNum)
	if err != nil {
		return nil, err
	}
	if gasPriceBid == nil || gasPriceBid.Sign() <= 0 {
		gasPriceBid = snap.MaxGasPriceBid()
	}

	createTicket := message.RetryableTx{
		Destination:       common.NewAddressFromEth(destAddr),
		Value:             l2CallValue,
		Deposit:           deposit,
		MaxSubmissionCost: maxSubmissionCost,
		CreditBack:        common.NewAddressFromEth(excessFeeRefundAddress),
		Beneficiary:       common.NewAddressFromEth(callValueRefundAddress),
		MaxGas:            maxGas,
		GasPriceBid:       gasPriceBid,
		Data:              data,
	}

	res, _, err := snap.EstimateRetryableGas(ctx, createTicket, common.NewAddressFromEth(sender), srv.maxAVMGas)
	if err != nil {
		return nil, err
	}
	if res.ResultCode != evm.ReturnCode {
		return nil, evm.HandleCallError(res, srv.ganacheMode)
	}

	res.FeeStats.UnitsUsed.L1Calldata = big.NewInt(0)
	used := res.FeeStats.TargetGasUsed()
	used = used.Mul(used, big.NewInt(11))
	used = used.Div(used, big.NewInt(10))
	return estimateRetryableTicket.Outputs.PackValues([]interface{}{
		new(big.Int).Add(used, big.NewInt(100)),
		ApplyGasPriceBidFactor(res.GasPrice),
	})
}
