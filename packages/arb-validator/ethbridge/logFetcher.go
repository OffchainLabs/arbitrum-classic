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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type maybeLog struct {
	log types.Log
	err error
}

func getLogs(
	ctx context.Context,
	client *ethclient.Client,
	filter ethereum.FilterQuery,
	startHeight *common.TimeBlocks,
	startIndex uint,
) (<-chan maybeLog, error) {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	streamingLogChan := make(chan types.Log)
	logSub, err := client.SubscribeFilterLogs(ctx, filter, streamingLogChan)
	if err != nil {
		return nil, err
	}
	logChan := make(chan maybeLog, 1024)
	go func() {
		defer close(logChan)
		defer logSub.Unsubscribe()

		// Get initial old logs
		filter.FromBlock = startHeight.AsInt()
		filter.ToBlock = header.Number
		log.Println("Filter1 from", filter.FromBlock, "to", filter.ToBlock)
		logs, err := client.FilterLogs(ctx, filter)
		if err != nil {
			logChan <- maybeLog{err: err}
			return
		}
		for _, ethLog := range logs {
			if ethLog.BlockNumber > startHeight.AsInt().Uint64() || ethLog.Index >= startIndex {
				logChan <- maybeLog{log: ethLog}
			}
		}

		// Retrieve for log from stream
		var ethStreamLog types.Log
		var ok bool
		select {
		case <-ctx.Done():
			return
		case ethStreamLog, ok = <-streamingLogChan:
			if !ok {
				logChan <- maybeLog{err: errors.New("streamingLogChan terminated early1")}
				return
			}
			log.Println("First stream log", ethStreamLog.BlockNumber)
		case err := <-logSub.Err():
			logChan <- maybeLog{err: err}
			return
		}

		// If there was a gap between initial retrieval and the stream, fill it in
		if ethStreamLog.BlockNumber > header.Number.Uint64()+1 {
			filter.FromBlock = new(big.Int).Add(header.Number, big.NewInt(1))
			filter.ToBlock = new(big.Int).Sub(new(big.Int).SetUint64(ethStreamLog.BlockNumber), big.NewInt(1))
			log.Println("Filter2 from", filter.FromBlock, "to", filter.ToBlock)
			logs, err := client.FilterLogs(ctx, filter)
			if err != nil {
				logChan <- maybeLog{err: err}
				return
			}
			for _, ethLog := range logs {
				logChan <- maybeLog{log: ethLog}
			}
		}

		logChan <- maybeLog{log: ethStreamLog}

		for {
			select {
			case <-ctx.Done():
				return
			case ethLog, ok := <-streamingLogChan:
				if !ok {
					logChan <- maybeLog{err: errors.New("streamingLogChan terminated early2")}
					return
				}
				logChan <- maybeLog{log: ethLog}
			case err := <-logSub.Err():
				logChan <- maybeLog{err: err}
				return
			}
		}
	}()

	return logChan, nil
}

func getEvents(
	ctx context.Context,
	client *ethclient.Client,
	filter ethereum.FilterQuery,
	startHeight *common.TimeBlocks,
	startIndex uint,
	logTransformer func(types.Log) (arbbridge.Event, error),
) (<-chan arbbridge.MaybeEvent, error) {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	streamingLogChan := make(chan types.Log)
	logSub, err := client.SubscribeFilterLogs(ctx, filter, streamingLogChan)
	if err != nil {
		return nil, err
	}
	logChan := make(chan arbbridge.MaybeEvent, 1024)
	go func() {
		defer close(logChan)
		defer logSub.Unsubscribe()

		// Get initial old logs
		filter.FromBlock = startHeight.AsInt()
		filter.ToBlock = header.Number
		log.Println("Filter1 from", filter.FromBlock, "to", filter.ToBlock)
		logs, err := client.FilterLogs(ctx, filter)
		if err != nil {
			logChan <- arbbridge.MaybeEvent{Err: err}
			return
		}
		for _, ethLog := range logs {
			if ethLog.BlockNumber > startHeight.AsInt().Uint64() || ethLog.Index >= startIndex {
				event, err := logTransformer(ethLog)
				if err != nil {
					logChan <- arbbridge.MaybeEvent{Err: err}
					return
				}
				logChan <- arbbridge.MaybeEvent{Event: event}
			}
		}

		// Retrieve for log from stream
		var ethStreamLog types.Log
		var ok bool
		select {
		case <-ctx.Done():
			return
		case ethStreamLog, ok = <-streamingLogChan:
			if !ok {
				logChan <- arbbridge.MaybeEvent{Err: errors.New("streamingLogChan terminated early1")}
				return
			}
			log.Println("First stream log", ethStreamLog.BlockNumber)
		case err := <-logSub.Err():
			logChan <- arbbridge.MaybeEvent{Err: err}
			return
		}

		// If there was a gap between initial retrieval and the stream, fill it in
		if ethStreamLog.BlockNumber > header.Number.Uint64()+1 {
			filter.FromBlock = new(big.Int).Add(header.Number, big.NewInt(1))
			filter.ToBlock = new(big.Int).Sub(new(big.Int).SetUint64(ethStreamLog.BlockNumber), big.NewInt(1))
			log.Println("Filter2 from", filter.FromBlock, "to", filter.ToBlock)
			logs, err := client.FilterLogs(ctx, filter)
			if err != nil {
				logChan <- arbbridge.MaybeEvent{Err: err}
				return
			}
			for _, ethLog := range logs {
				event, err := logTransformer(ethLog)
				if err != nil {
					logChan <- arbbridge.MaybeEvent{Err: err}
					return
				}
				logChan <- arbbridge.MaybeEvent{Event: event}
			}
		}

		event, err := logTransformer(ethStreamLog)
		if err != nil {
			logChan <- arbbridge.MaybeEvent{Err: err}
			return
		}
		logChan <- arbbridge.MaybeEvent{Event: event}

		for {
			select {
			case <-ctx.Done():
				return
			case ethLog, ok := <-streamingLogChan:
				if !ok {
					logChan <- arbbridge.MaybeEvent{Err: errors.New("streamingLogChan terminated early2")}
					return
				}
				event, err := logTransformer(ethLog)
				if err != nil {
					logChan <- arbbridge.MaybeEvent{Err: err}
					return
				}
				logChan <- arbbridge.MaybeEvent{Event: event}
			case err := <-logSub.Err():
				logChan <- arbbridge.MaybeEvent{Err: err}
				return
			}
		}
	}()

	return logChan, nil
}

func nextBlockHash(
	ctx context.Context,
	client *ethclient.Client,
	prevBlock *structures.BlockId,
) (*structures.BlockId, error) {
	if prevBlock == nil {
		header, err := client.HeaderByNumber(ctx, big.NewInt(0))
		if err != nil {
			return nil, err
		}
		return getBlockID(header), nil
	} else {
		header, err := client.HeaderByNumber(ctx, new(big.Int).Add(prevBlock.Height.AsInt(), big.NewInt(1)))
		if err != nil {
			return nil, err
		}
		if header.ParentHash != prevBlock.HeaderHash.ToEthHash() {
			return nil, err
		}
		return getBlockID(header), nil
	}
}

func getNextLogs(
	ctx context.Context,
	client *ethclient.Client,
	filter ethereum.FilterQuery,
	prevBlock *structures.BlockId,
) ([]types.Log, error) {
	nextBlock, err := nextBlockHash(ctx, client, prevBlock)
	if err != nil {
		return nil, err
	}
	blockHash := nextBlock.HeaderHash.ToEthHash()
	return client.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: filter.Addresses,
		Topics:    filter.Topics,
	})
}
