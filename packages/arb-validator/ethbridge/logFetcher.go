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
	"errors"
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
	//debug.PrintStack()
	go func() {
		defer close(logChan)
		defer logSub.Unsubscribe()
		// Get initial old logs
		filter.FromBlock = startHeight
		filter.ToBlock = header.Number
		log.Println("Filter1 from", filter.FromBlock, "to", filter.ToBlock)
		logs, err := client.FilterLogs(ctx, filter)
		if err != nil {
			errChan <- err
			return
		}
		for _, ethLog := range logs {
			logChan <- ethLog
		}

		// Retreive for log from stream
		var ethStreamLog types.Log
		select {
		case <-ctx.Done():
			return
		case ethStreamLog = <-streamingLogChan:
			log.Println("First stream log", ethStreamLog.BlockNumber)
		case err := <-logSub.Err():
			errChan <- err
			return
		}

		// If there was a gap between initial retrieval and the stream, fill it in
		if ethStreamLog.BlockNumber > header.Number.Uint64()+1 {
			filter.FromBlock = new(big.Int).Add(header.Number, big.NewInt(1))
			filter.ToBlock = new(big.Int).Sub(new(big.Int).SetUint64(ethStreamLog.BlockNumber), big.NewInt(1))
			log.Println("Filter2 from", filter.FromBlock, "to", filter.ToBlock)
			logs, err := client.FilterLogs(ctx, filter)
			if err != nil {
				errChan <- err
				return
			}
			for _, ethLog := range logs {
				logChan <- ethLog
			}
		}
		logChan <- ethStreamLog

		for {
			select {
			case <-ctx.Done():
				return
			case ethLog, ok := <-streamingLogChan:
				if !ok {
					errChan <- errors.New("streamingLogChan terminated early2")
					return
				}
				logChan <- ethLog
			case err := <-logSub.Err():
				errChan <- err
				return
			}
		}
	}()

	return nil
}
