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

package chainlistener

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"log"
)

func InitiateChallenge(
	ctx context.Context,
	rollup arbbridge.ArbRollup,
	opp *nodegraph.ChallengeOpportunity) ([]arbbridge.Event, error) {
	return rollup.StartChallenge(
		ctx,
		opp.Asserter(),
		opp.Challenger(),
		opp.PrevNodeHash(),
		opp.DeadlineTicks().Val,
		opp.AsserterNodeType(),
		opp.ChallengerNodeType(),
		opp.AsserterVMProtoHash(),
		opp.ChallengerVMProtoHash(),
		opp.AsserterProof(),
		opp.ChallengerProof(),
		opp.AsserterNodeHash(),
		opp.ChallengerDataHash(),
		opp.ChallengerPeriodTicks(),
	)
}

func LogChallengeResult(err error) {
	if err != nil {
		log.Println("Failed to initiate challenge", err)
	} else {
		log.Println("Successfully initiated challenge")
	}
}
