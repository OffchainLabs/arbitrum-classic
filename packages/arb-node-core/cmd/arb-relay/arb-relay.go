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
	SequencerFeedAddress string
	broadcastClient      *broadcastclient.BroadcastClient
	broadcaster          *broadcaster.Broadcaster
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

	config, err := configuration.ParseFeed(ctx)
	if err != nil || len(config.Feed.Input.URL) == 0 {
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

	if config.Feed.Input.URL == "" {
		return errors.New("Missing --feed.input.url")
	}

	if config.PProf.Enable {
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
	broadcastClient := broadcastclient.NewBroadcastClient(settings.Input.URL, nil, settings.Input.Timeout)
	broadcastClient.ConfirmedAccumulatorListener = make(chan common.Hash, 1)
	return &ArbRelay{
		SequencerFeedAddress: settings.Input.URL,
		broadcaster:          broadcaster.NewBroadcaster(settings.Output),
		broadcastClient:      broadcastClient,
	}
}

func (ar *ArbRelay) Start(ctx context.Context) (chan bool, error) {
	done := make(chan bool)

	err := ar.broadcaster.Start(ctx)
	if err != nil {
		return nil, errors.New("broadcast unable to start")
	}

	// connect returns
	var messages chan broadcaster.BroadcastFeedMessage
	for {
		messages, err = ar.broadcastClient.Connect(ctx)
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

	go func() {
		defer func() {
			done <- true
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-messages:
				err = ar.broadcaster.BroadcastSingle(msg.FeedItem.PrevAcc, msg.FeedItem.BatchItem, msg.Signature)
				if err != nil {
					logger.
						Error().
						Err(err).
						Hex("PrevAcc", msg.FeedItem.PrevAcc.Bytes()).
						Hex("BatchItem", msg.FeedItem.BatchItem.ToBytesWithSeqNum()).
						Msg("unable to broadcast batch item")
				}
			case ca := <-ar.broadcastClient.ConfirmedAccumulatorListener:
				ar.broadcaster.ConfirmedAccumulator(ca)
			}
		}
	}()

	return done, nil
}

func (ar *ArbRelay) Stop() {
	ar.broadcastClient.Close()
	ar.broadcaster.Stop()
}
