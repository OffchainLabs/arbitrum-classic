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
	"context"
	"errors"
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChainValidator struct {
	*Validator
	chainBot *ChainBot
}

func NewChainValidator(
	name string,
	address common.Address,
	latestHeader *types.Header,
	balance *protocol.BalanceTracker,
	config *valmessage.VMConfiguration,
	machine machine.Machine,
	challengeEverything bool,
	maxCallSteps int32,
) *ChainValidator {
	c := core.NewCore(
		machine,
		balance,
	)

	valConfig := core.NewValidatorConfig(address, config, challengeEverything, maxCallSteps)
	chainBot := &ChainBot{state.NewWaiting(valConfig, c)}
	val := NewValidator(
		name,
		chainBot,
		latestHeader,
	)
	return &ChainValidator{val, chainBot}
}

func (validator *ChainValidator) Run(ctx context.Context, recvChan <-chan ethconnection.Notification, bridge bridge.ArbVMBridge) {
	defer fmt.Printf("%v: Exiting\n", validator.Name)
	for {
		select {
		case <-ctx.Done():
			break
		case notification, ok := <-recvChan:
			// fmt.Printf("Got valmessage %T: %v\n", event, event)
			if !ok {
				fmt.Printf("%v: Error in recvChan\n", validator.Name)
				return
			}

			newHeader := notification.Header
			if validator.latestHeader == nil || newHeader.Number.Uint64() >= validator.latestHeader.Number.Uint64() && newHeader.Hash() != validator.latestHeader.Hash() {
				validator.latestHeader = newHeader
				validator.timeUpdate(bridge)

				if validator.pendingDisputableRequest != nil {
					pre := validator.pendingDisputableRequest.Precondition
					if !validator.chainBot.GetCore().ValidateAssertion(pre, newHeader.Number.Uint64()) {
						validator.pendingDisputableRequest.ErrorChan <- errors.New("Precondition was invalidated")
						close(validator.pendingDisputableRequest.ErrorChan)
						close(validator.pendingDisputableRequest.ResultChan)
						validator.pendingDisputableRequest = nil
					}
				}
			}

			switch ev := notification.Event.(type) {
			case ethconnection.NewTimeEvent:
				break
			case ethconnection.VMEvent:
				validator.eventUpdate(ev, notification.Header, bridge)
			case ethconnection.MessageDeliveredEvent:
				validator.chainBot.SendMessageToVM(ev.Msg)
			default:
				panic("Should never recieve other kinds of events")
			}
		case action := <-validator.actions:
			action(bridge)
		case <-validator.maybeAssert:
		}

		if bot, ok := validator.chainBot.ChainState.(state.Waiting); ok && validator.pendingDisputableRequest != nil {
			validator.chainBot.updateBot(bot.AttemptAssertion(context.Background(), *validator.pendingDisputableRequest, bridge))
			validator.pendingDisputableRequest = nil
		}
	}
}

func (validator *ChainValidator) timeUpdate(bridge bridge.ArbVMBridge) {
	if validator.challengeBot != nil {
		newBot, err := validator.challengeBot.UpdateTime(validator.latestHeader.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
			return
		}
		validator.challengeBot = newBot
	}
	newBot, err := validator.chainBot.ChainUpdateTime(validator.latestHeader.Number.Uint64(), bridge)
	if err != nil {
		fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
		return
	}
	validator.chainBot.updateBot(newBot)
}

func (validator *ChainValidator) eventUpdate(ev ethconnection.VMEvent, header *types.Header, bridge bridge.ArbVMBridge) {
	if ev.GetIncomingMessageType() == ethconnection.ChallengeMessage {
		if validator.challengeBot == nil {
			panic("challengeBot can't be nil if challenge message is recieved")
		}

		newBot, err := validator.challengeBot.UpdateState(ev, header.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
			return
		}
		validator.challengeBot = newBot
	} else {
		newBot, challengeBot, err := validator.chainBot.ChainUpdateState(ev, header.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, validator.chainBot)
			return
		}
		validator.chainBot.updateBot(newBot)
		if challengeBot != nil {
			validator.challengeBot = challengeBot
		}
	}
}
