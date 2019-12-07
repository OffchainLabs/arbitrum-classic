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
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/observer"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/defender"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/challenger"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type challengerInitiator struct {
	precondition        *protocol.Precondition
	machine             machine.Machine
	challengeEverything bool
}

func (bot *challengerInitiator) initiateBot(ev ethbridge.InitiateChallengeEvent, brdg bridge.Challenge) (challenge.State, error) {
	return challenger.New(bot.precondition, bot.machine, ev.Deadline, bot.challengeEverything), nil
}

type defenderInitiator struct {
	assDef machine.AssertionDefender
}

func (bot *defenderInitiator) initiateBot(ev ethbridge.InitiateChallengeEvent, brdg bridge.Challenge) (challenge.State, error) {
	return defender.New(bot.assDef, ev.Deadline, brdg)
}

type observerInitiator struct {
}

func (bot *observerInitiator) initiateBot(ev ethbridge.InitiateChallengeEvent, brdg bridge.Challenge) (challenge.State, error) {
	return observer.New(ev.Deadline), nil
}

type initiator interface {
	initiateBot(ethbridge.InitiateChallengeEvent, bridge.Challenge) (challenge.State, error)
}

type Challenge struct {
	// Run loop only
	brdge        bridge.Challenge
	init         initiator
	bot          challenge.State
	latestHeader *types.Header
}

func NewChallengerValidator(
	brdge bridge.Challenge,
	latestHeader *types.Header,
	precondition *protocol.Precondition,
	machine machine.Machine,
	challengeEverything bool,

) *Challenge {
	return &Challenge{
		brdge,
		&challengerInitiator{
			precondition:        precondition,
			machine:             machine,
			challengeEverything: challengeEverything,
		},
		nil,
		latestHeader,
	}
}

func NewDefenderValidator(
	brdge bridge.Challenge,
	latestHeader *types.Header,
	assDef machine.AssertionDefender,

) *Challenge {
	return &Challenge{
		brdge,
		&defenderInitiator{
			assDef: assDef,
		},
		nil,
		latestHeader,
	}
}

func NewObserverValidator(
	brdge bridge.Challenge,
	latestHeader *types.Header,

) *Challenge {
	return &Challenge{
		brdge,
		&observerInitiator{},
		nil,
		latestHeader,
	}
}

func (c *Challenge) validatorClosing() {
	fmt.Printf("Challenge Validator: Exiting\n")
	c.brdge.SendMonitorErr(bridge.Error{errors.New("WARNING: c closing"), "WARNING: c closing", false})
}

func (c *Challenge) Run(ctx context.Context, recvChan <-chan ethbridge.Notification) {
	defer c.validatorClosing()
	for {
		select {
		case <-ctx.Done():
			break
		case notification, ok := <-recvChan:
			//log.Printf("Challenge Validator got notification %T: %v\n", notification.Event, notification.Event)
			if !ok {
				fmt.Printf("Challenge Validator: Error in recvChan\n")
				return
			}

			newHeader := notification.Header
			if c.bot != nil && (c.latestHeader == nil || newHeader.Number.Uint64() >= c.latestHeader.Number.Uint64() && newHeader.Hash() != c.latestHeader.Hash()) {
				c.latestHeader = newHeader
				err := c.timeUpdate()
				if err != nil {
					//log.Printf("Validator %v: Error processing time update - %v\n", c.Name, err)
					if errstat, ok := err.(*bridge.Error); ok {
						if !errstat.Recoverable {
							//log.Printf("Validator %v: non recoverable error\n", c.Name)
							c.brdge.SendMonitorErr(*errstat)
							return
						} else {
							//log.Printf("Validator %v: recoverable error - contiuing\n", c.Name)
							c.brdge.SendMonitorErr(*errstat)
						}
					} else {
						c.brdge.SendMonitorErr(bridge.Error{err, "non recoverable error - exiting", false})
						return
					}
				}
			}

			switch ev := notification.Event.(type) {
			case ethbridge.NewTimeEvent:
				break
			case ethbridge.InitiateChallengeEvent:
				bot, err := c.init.initiateBot(ev, c.brdge)
				if err != nil {
					c.brdge.SendMonitorErr(bridge.Error{
						Err:         err,
						Message:     "Failed to initiate challenge bot",
						Recoverable: false,
					})
					return
				}
				c.bot = bot
			case ethbridge.VMEvent:
				err := c.eventUpdate(ev, notification.Header)
				if err != nil {
					//log.Printf("*****Validator %v: error - %v\n", c.Name, err)
					if errstat, ok := err.(*bridge.Error); ok {
						if !errstat.Recoverable {
							//log.Printf("Validator %v: non recoverable error - %v\n", c.Name, err)
							c.brdge.SendMonitorErr(*errstat)
							return
						} else {
							//log.Printf("Validator %v: recoverable error - %v - contiuing\n", c.Name, err)
							c.brdge.SendMonitorErr(*errstat)
						}
					} else {
						log.Println("Error processing event update", err)
						return
					}
				}
			default:
				panic("Should never receive other kinds of events")
			}
		}
	}
}

func (c *Challenge) timeUpdate() error {
	newBot, err := c.bot.UpdateTime(c.latestHeader.Number.Uint64(), c.brdge)
	if err != nil {
		return err
	}
	c.bot = newBot
	return nil
}

func (c *Challenge) eventUpdate(ev ethbridge.VMEvent, header *types.Header) error {
	newBot, err := c.bot.UpdateState(ev, header.Number.Uint64(), c.brdge)
	if err != nil {
		return err
	}
	c.bot = newBot
	return nil
}
