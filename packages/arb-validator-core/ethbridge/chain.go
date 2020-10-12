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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"

	"github.com/pkg/errors"

	ethereum "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ArbAddresses struct {
	ArbFactory string `json:"ArbFactory"`
}

func (a ArbAddresses) ArbFactoryAddress() common.Address {
	return common.NewAddressFromEth(ethcommon.HexToAddress(a.ArbFactory))
}

func getLogBlockID(ethLog types.Log) *common.BlockId {
	return &common.BlockId{
		Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
		HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
	}
}

func GetReceiptBlockID(receipt *types.Receipt) *common.BlockId {
	return &common.BlockId{
		Height:     common.NewTimeBlocks(receipt.BlockNumber),
		HeaderHash: common.NewHashFromEth(receipt.BlockHash),
	}
}

func getLogChainInfo(log types.Log) arbbridge.ChainInfo {
	return arbbridge.ChainInfo{
		BlockId:  getLogBlockID(log),
		LogIndex: log.Index,
	}
}

func waitForReceipt(ctx context.Context, client ethutils.EthClient, from ethcommon.Address, tx *types.Transaction, methodName string) error {
	_, err := WaitForReceiptWithResults(ctx, client, from, tx, methodName)
	return err
}

func WaitForReceiptWithResultsSimple(ctx context.Context, client ethutils.EthClient, txHash ethcommon.Hash) (*types.Receipt, error) {
	for {
		select {
		case _ = <-time.After(time.Second):
			receipt, err := client.TransactionReceipt(ctx, txHash)
			if receipt == nil && err == nil {
				continue
			}
			if err != nil {
				if err.Error() == ethereum.NotFound.Error() {
					continue
				}
				log.Println("ERROR getting receipt", err)
				return nil, err
			}
			return receipt, nil
		case _ = <-ctx.Done():
			return nil, errors.New("Receipt not found")
		}
	}
}

func WaitForReceiptWithResults(ctx context.Context, client ethutils.EthClient, from ethcommon.Address, tx *types.Transaction, methodName string) (*types.Receipt, error) {
	receipt, err := WaitForReceiptWithResultsSimple(ctx, client, tx.Hash())
	if err != nil {
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
}
