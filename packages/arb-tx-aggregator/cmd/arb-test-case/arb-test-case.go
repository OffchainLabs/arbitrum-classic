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

package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"io/ioutil"
	golog "log"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
)

var logger zerolog.Logger

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "arb-test-case").Logger()

	if err := generateTestCase(
		"http://localhost:7545",
		common.HexToAddress("0xc68DCee7b8cA57F41D1A417103CB65836E99e013"),
		arbos.Path(),
	); err != nil {
		logger.Fatal().Stack().Err(err).Msg("error generating test")
	}
}

func generateTestCase(ethURL string, rollupAddress common.Address, contract string) error {
	ctx := context.Background()

	ethclint, err := ethutils.NewRPCEthClient(ethURL)
	if err != nil {
		return err
	}

	client := ethbridge.NewEthClient(ethclint)
	rollupWatcher, err := client.NewRollupWatcher(rollupAddress)
	if err != nil {
		return err
	}

	inboxAddress, err := rollupWatcher.InboxAddress(ctx)
	if err != nil {
		return err
	}

	inboxWatcher, err := client.NewGlobalInboxWatcher(inboxAddress, rollupAddress)
	if err != nil {
		return err
	}

	_, eventId, _, _, err := rollupWatcher.GetCreationInfo(ctx)
	if err != nil {
		return err
	}

	events, err := inboxWatcher.GetDeliveredEvents(ctx, eventId.BlockId.Height.AsInt(), nil)
	if err != nil {
		return err
	}

	messages := make([]inbox.InboxMessage, 0, len(events))
	for _, ev := range events {
		messages = append(messages, ev.Message)
	}

	mach, err := cmachine.New(contract)
	if err != nil {
		return err
	}

	// Last value returned is not an error type
	assertion, _, _ := mach.ExecuteAssertion(
		1000000000000,
		true,
		messages,
		true,
	)

	data, err := inbox.TestVectorJSON(messages, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("log.json", data, 0644); err != nil {
		return err
	}
	return nil
}
