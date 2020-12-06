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

package rollup

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

var logger = log.With().Caller().Str("component", "rollup").Logger()

// WARNING: The code in this file is badly behaved, on purpose. It is for testing only.
//     If you call this in production, you will be sorry.

type WrongAssertionType int

const (
	WrongInboxTopAssertion      = 0
	WrongMessagesSliceAssertion = 1
	WrongExecutionAssertion     = 2
)

type evil_WrongAssertionListener struct {
	*chainlistener.ValidatorChainListener
	kind WrongAssertionType
}

func NewEvil_WrongAssertionListener(
	rollupAddress common.Address,
	actor arbbridge.ArbRollup,
	kind WrongAssertionType,
) *evil_WrongAssertionListener {
	return &evil_WrongAssertionListener{chainlistener.NewValidatorChainListener(context.Background(), rollupAddress, actor), kind}
}

func (lis *evil_WrongAssertionListener) AssertionPrepared(
	ctx context.Context,
	params valprotocol.ChainParams,
	nodeGraph *nodegraph.StakedNodeGraph,
	nodeLocation *structures.Node,
	prepared *chainlistener.PreparedAssertion) {
	badHash := common.Hash{}
	badHash[5] = 37
	switch lis.kind {
	case WrongInboxTopAssertion:
		prepared.AssertionStub.AfterInboxHash = badHash
		logger.Info().Msg("Prepared EVIL inbox top assertion")
	case WrongExecutionAssertion:
		prepared.AssertionStub.AfterMachineHash = badHash
		logger.Info().Msg("Prepared EVIL execution assertion")
	default:
		logger.Fatal().Msg("unrecognized evil listener type")
	}
	lis.ValidatorChainListener.AssertionPrepared(ctx, params, nodeGraph, nodeLocation, prepared)
}
