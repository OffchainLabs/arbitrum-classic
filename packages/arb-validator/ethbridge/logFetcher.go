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

package ethbridge

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getLogs(
	ctx context.Context,
	client *ethclient.Client,
	filter ethereum.FilterQuery,
	startHeight *big.Int,
	logChan chan types.Log,
	errChan chan error,
) error {
	streamingLogChan := make(chan types.Log)
	logSub, err := client.SubscribeFilterLogs(ctx, filter, streamingLogChan)
	if err != nil {
		return err
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	go func() {
		defer close(logChan)
		defer logSub.Unsubscribe()
		// Get initial old logs
		filter.FromBlock = startHeight
		filter.ToBlock = header.Number
		logs, err := client.FilterLogs(ctx, filter)
		if err != nil {
			errChan <- err
			return
		}
		for _, ethLog := range logs {
			log.Println("getLogs1", ethLog.BlockNumber, ethLog.TxIndex, ethLog.Index)
			logChan <- ethLog
		}

		// Retreive for log from stream
		ethLog := <-streamingLogChan
		log.Println("getLogs2", ethLog.BlockNumber, ethLog.TxIndex, ethLog.Index)

		// If there was a gap between initial retrieval and the stream, fill it in
		if ethLog.BlockNumber > header.Number.Uint64() {
			filter.FromBlock = header.Number
			filter.ToBlock = new(big.Int).Sub(new(big.Int).SetUint64(ethLog.BlockNumber), big.NewInt(1))
			logs, err := client.FilterLogs(ctx, filter)
			if err != nil {
				errChan <- err
				return
			}
			for _, ethLog := range logs {
				log.Println("getLogs3", ethLog.BlockNumber, ethLog.TxIndex, ethLog.Index)
				logChan <- ethLog
			}
		}
		logChan <- ethLog

		for {
			select {
			case <-ctx.Done():
				return
			case ethLog := <-streamingLogChan:
				log.Println("getLogs4", ethLog.BlockNumber, ethLog.TxIndex, ethLog.Index)
				logChan <- ethLog
			case err := <-logSub.Err():
				errChan <- err
			}
		}
	}()

	return nil
}
