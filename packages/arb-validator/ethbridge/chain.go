/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
	"fmt"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ArbAddresses struct {
	ArbFactory   string `json:"ArbFactory"`
	OneStepProof string `json:"OneStepProof"`
}

func (a ArbAddresses) ArbFactoryAddress() common.Address {
	return common.NewAddressFromEth(ethcommon.HexToAddress(a.ArbFactory))
}

func (a ArbAddresses) OneStepProofAddress() common.Address {
	return common.NewAddressFromEth(ethcommon.HexToAddress(a.OneStepProof))
}

func getBlockID(header *types.Header) *structures.BlockId {
	return &structures.BlockId{
		Height:     common.NewTimeBlocks(header.Number),
		HeaderHash: common.NewHashFromEth(header.Hash()),
	}
}

func getLogBlockID(ethLog types.Log) *structures.BlockId {
	return &structures.BlockId{
		Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
		HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
	}
}

func getTxBlockID(receipt *types.Receipt) *structures.BlockId {
	return &structures.BlockId{
		Height:     common.NewTimeBlocks(receipt.BlockNumber),
		HeaderHash: common.NewHashFromEth(receipt.BlockHash),
	}
}

func getLogChainInfo(log types.Log) arbbridge.ChainInfo {
	return arbbridge.ChainInfo{
		BlockId:  getLogBlockID(log),
		LogIndex: log.Index,
		TxHash:   log.TxHash,
	}
}

func waitForReceipt(ctx context.Context, client *ethclient.Client, from ethcommon.Address, tx *types.Transaction, methodName string) error {
	_, err := WaitForReceiptWithResults(ctx, client, from, tx, methodName)
	return err
}
func WaitForReceiptWithResults(ctx context.Context, client *ethclient.Client, from ethcommon.Address, tx *types.Transaction, methodName string) (*types.Receipt, error) {
	for {
		select {
		case _ = <-time.After(time.Second):
			receipt, err := client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				if err.Error() == ethereum.NotFound.Error() {
					continue
				}
				return nil, err
			}
			if receipt.Status != 1 {
				data, err := receipt.MarshalJSON()
				if err != nil {
					return nil, errors.New("Failed unmarshalling receipt")
				}
				callMsg := ethereum.CallMsg{
					From:     from,
					To:       tx.To(),
					Gas:      tx.Gas(),
					GasPrice: tx.GasPrice(),
					Value:    tx.Value(),
					Data:     tx.Data(),
				}
				data, err = client.CallContract(ctx, callMsg, receipt.BlockNumber)
				if err != nil {
					return nil, fmt.Errorf("Transaction %v failed with error %v", methodName, err)
				}
				return nil, fmt.Errorf("Transaction %v failed with tx %v", methodName, string(data))
			}
			return receipt, nil
		case _ = <-ctx.Done():
			return nil, errors.New("Receipt not found")
		}
	}
}
