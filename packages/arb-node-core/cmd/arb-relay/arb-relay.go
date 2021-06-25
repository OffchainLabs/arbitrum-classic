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

package main

import (
	"context"
	"fmt"
	golog "log"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

type ArbRelay struct {
	broadcastClients         []*broadcastclient.BroadcastClient
	broadcaster              *broadcaster.Broadcaster
	confirmedAccumulatorChan chan common.Hash
}

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
}

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Print line number that log was created on
	logger = log.With().Caller().Stack().Str("component", "arb-validator").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running relay")
	}
}

func startup() error {
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	config, err := configuration.ParseRelay()
	if err != nil || len(config.Feed.Input.URLs) == 0 {
		fmt.Printf("\n")
		fmt.Printf("Sample usage: arb-relay --conf=<filename> \n")
		fmt.Printf("          or: arb-relay --feed.input.url=<feed websocket>\n\n")
		if err != nil && !strings.Contains(err.Error(), "help requested") {
			fmt.Printf("%s\n", err.Error())
		}

		return nil
	}

	if err := cmdhelp.ParseLogFlags(&config.Log.RPC, &config.Log.Core); err != nil {
		return err
	}

	defer logger.Info().Msg("Cleanly shutting down relay")

	if config.PProfEnable {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	// Start up an arbitrum sequencer relay
	arbRelay := NewArbRelay(config.Feed)
	relayDone, err := arbRelay.Start(ctx)
	if err != nil {
		return err
	}
	defer arbRelay.Stop()

	select {
	case <-cancelChan:
		return nil
	case <-relayDone:
		return nil
	}
}

func NewArbRelay(settings configuration.Feed) *ArbRelay {
	var broadcastClients []*broadcastclient.BroadcastClient
	confirmedAccumulatorChan := make(chan common.Hash, 1)
	for _, address := range settings.Input.URLs {
		client := broadcastclient.NewBroadcastClient(address, nil, settings.Input.Timeout)
		client.ConfirmedAccumulatorListener = confirmedAccumulatorChan
		broadcastClients = append(broadcastClients, client)
	}
	return &ArbRelay{
		broadcaster:              broadcaster.NewBroadcaster(settings.Output),
		broadcastClients:         broadcastClients,
		confirmedAccumulatorChan: confirmedAccumulatorChan,
	}
}

const RECENT_FEED_ITEM_TTL time.Duration = time.Second * 10

func (ar *ArbRelay) Start(ctx context.Context) (chan bool, error) {
	done := make(chan bool)

	err := ar.broadcaster.Start(ctx)
	if err != nil {
		return nil, errors.New("broadcast unable to start")
	}

	// connect returns
	messages := make(chan broadcaster.BroadcastFeedMessage)
	for _, client := range ar.broadcastClients {
		for {
			err = client.ConnectWithChannel(ctx, messages)
			if err == nil {
				break
			}
			logger.Warn().Err(err).
				Msg("failed connect to sequencer broadcast, waiting and retrying")
			select {
			case <-ctx.Done():
				return nil, errors.New("ctx cancelled broadcast client connect")
			case <-time.After(5 * time.Second):
			}
		}
	}

	recentFeedItems := make(map[common.Hash]time.Time)
	go func() {
		defer func() {
			done <- true
		}()
		recentFeedItemsCleanup := time.NewTicker(RECENT_FEED_ITEM_TTL)
		defer recentFeedItemsCleanup.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-messages:
				newAcc := msg.FeedItem.BatchItem.Accumulator
				if recentFeedItems[newAcc] != (time.Time{}) {
					continue
				}
				recentFeedItems[newAcc] = time.Now()
				err = ar.broadcaster.BroadcastSingle(msg.FeedItem.PrevAcc, msg.FeedItem.BatchItem, msg.Signature)
				if err != nil {
					logger.
						Error().
						Err(err).
						Hex("PrevAcc", msg.FeedItem.PrevAcc.Bytes()).
						Hex("BatchItem", msg.FeedItem.BatchItem.ToBytesWithSeqNum()).
						Msg("unable to broadcast batch item")
				}
			case ca := <-ar.confirmedAccumulatorChan:
				ar.broadcaster.ConfirmedAccumulator(ca)
			case <-recentFeedItemsCleanup.C:
				// Clear expired items from recentFeedItems
				recentFeedItemExpiry := time.Now().Add(-RECENT_FEED_ITEM_TTL)
				for acc, created := range recentFeedItems {
					if created.Before(recentFeedItemExpiry) {
						delete(recentFeedItems, acc)
					}
				}
			}
		}
	}()

	return done, nil
}

func (ar *ArbRelay) Stop() {
	for _, client := range ar.broadcastClients {
		client.Close()
	}
	ar.broadcaster.Stop()
}
