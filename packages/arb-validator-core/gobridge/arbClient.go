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

package gobridge

import (
	"context"
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var errReorgError = errors.New("reorg occured")
var headerRetryDelay = time.Second * 2
var maxFetchAttempts = 5

func NewEthClient(ethURL string) *goEthdata {
	return getGoEth(ethURL)
}

func (c *goEthdata) SubscribeBlockHeaders(ctx context.Context, startBlockID *common.BlockId) (<-chan arbbridge.MaybeBlockId, error) {
	blockIDChan := make(chan arbbridge.MaybeBlockId, 100)

	blockIDChan <- arbbridge.MaybeBlockId{BlockId: startBlockID}
	prevBlockID := startBlockID
	go func() {
		defer close(blockIDChan)

		for {
			var nextBlock *common.BlockId
			fetchErrorCount := 0
			for {
				if prevBlockID == nil {
					fmt.Println("prevBlockID nil")
				}
				nextHeight := common.NewTimeBlocks(new(big.Int).Add(prevBlockID.Height.AsInt(), big.NewInt(1)))
				n, notFound := c.getBlockFromHeight(nextHeight)
				if notFound == nil {
					// Got next header
					nextBlock = n
					break
				}

				select {
				case <-ctx.Done():
					// Getting header must have failed due to context cancellation
					return
				default:
				}

				if notFound != nil {
					log.Printf("Failed to fetch next header on attempt %v", fetchErrorCount)
					fetchErrorCount++
				}

				if fetchErrorCount >= maxFetchAttempts {
					err := fmt.Sprint("Next header not found after ", fetchErrorCount, " attempts")
					blockIDChan <- arbbridge.MaybeBlockId{Err: errors.New(err)}
					return
				}

				// Header was not found so wait before checking again
				time.Sleep(headerRetryDelay)
			}

			if c.parentHashes[*nextBlock] != prevBlockID.HeaderHash {
				blockIDChan <- arbbridge.MaybeBlockId{Err: errReorgError}
				return
			}

			prevBlockID = nextBlock
			blockIDChan <- arbbridge.MaybeBlockId{BlockId: prevBlockID}
		}
	}()

	return blockIDChan, nil
}

func (c *goEthdata) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	return newArbFactoryWatcher(address, c)
}

func (c *goEthdata) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return newRollupWatcher(address, c)
}

func (c *goEthdata) NewExecutionChallengeWatcher(address common.Address) (arbbridge.ExecutionChallengeWatcher, error) {
	return newExecutionChallengeWatcher(address, c)
}

func (c *goEthdata) NewMessagesChallengeWatcher(address common.Address) (arbbridge.MessagesChallengeWatcher, error) {
	return newMessagesChallengeWatcher(address, c)
}

func (c *goEthdata) NewInboxTopChallengeWatcher(address common.Address) (arbbridge.InboxTopChallengeWatcher, error) {
	return newInboxTopChallengeWatcher(address, c)
}

func (c *goEthdata) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	return newOneStepProof(address, c)
}

func (c *goEthdata) GetBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	return c.balances[account], nil
}

func (c *goEthdata) CurrentBlockId(ctx context.Context) (*common.BlockId, error) {
	return c.LastMinedBlock, nil
}

func (c *goEthdata) BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
	block, err := c.getBlockFromHeight(height)
	if err != nil {
		errstr := fmt.Sprintln("block height", height, " not found")
		return nil, errors.New(errstr)
	}
	return block, nil
}

type TransOpts struct {
	sync.Mutex
	From common.Address
}

type GoArbAuthClient struct {
	*goEthdata
	fromAddr common.Address
}

func NewEthAuthClient(ethURL string, fromAddr common.Address) (*GoArbAuthClient, error) {
	client := NewEthClient(ethURL)

	client.balances[fromAddr] = big.NewInt(1000) // give client a default balance of 1000
	return &GoArbAuthClient{
		goEthdata: client,
		fromAddr:  fromAddr,
	}, nil
}

func (c *GoArbAuthClient) Address() common.Address {
	return c.fromAddr
}

func (c *GoArbAuthClient) NewArbFactory(address common.Address) (arbbridge.ArbFactory, error) {
	return newArbFactory(address, c)
}

func (c *GoArbAuthClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	return newRollupContract(address, c)
}

func (c *GoArbAuthClient) NewGlobalInbox(address common.Address) (arbbridge.GlobalInbox, error) {
	return newGlobalInbox(address, c)
}

func (c *GoArbAuthClient) NewChallengeFactory(address common.Address) (arbbridge.ChallengeFactory, error) {
	return newChallengeFactory(address, c)
}

func (c *GoArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	return NewExecutionChallenge(address, c)
}

func (c *GoArbAuthClient) NewMessagesChallenge(address common.Address) (arbbridge.MessagesChallenge, error) {
	return newMessagesChallenge(address, c)
}

func (c *GoArbAuthClient) NewInboxTopChallenge(address common.Address) (arbbridge.InboxTopChallenge, error) {
	return newInboxTopChallenge(address, c)
}

func (c *GoArbAuthClient) DeployChallengeTest(ctx context.Context, challengeFactory common.Address) (arbbridge.ChallengeTester, error) {
	//c.auth.Lock()
	//defer c.auth.Unlock()
	tester, err := NewChallengeTester(c)
	if err != nil {
		return nil, err
	}
	return tester, nil
}

func (c *GoArbAuthClient) DeployOneStepProof(ctx context.Context) (arbbridge.OneStepProof, error) {
	//c.auth.Lock()
	//defer c.auth.Unlock()
	osp, err := newOneStepProof(c.Address(), c.goEthdata)
	return osp, err
}
