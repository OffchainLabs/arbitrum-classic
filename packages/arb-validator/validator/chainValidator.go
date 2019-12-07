/*
 * Copyright 2019, Offchain Labs, Inc.
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

package validator

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChainValidator struct {
	*Validator
	chainBot *ChainBot
}

func NewChainValidator(
	name string,
	b bridge.Bridge,
	address common.Address,
	latestHeader *types.Header,
	config *valmessage.VMConfiguration,
	machine machine.Machine,
	challengeEverything bool,
	maxCallSteps int32,
) *ChainValidator {
	c := core.NewCore(
		machine,
	)

	valConfig := core.NewValidatorConfig(address, config, challengeEverything, maxCallSteps)
	chainBot := &ChainBot{state.NewWaiting(valConfig, c), b}
	val := NewValidator(
		name,
		chainBot,
		latestHeader,
	)
	return &ChainValidator{val, chainBot}
}
