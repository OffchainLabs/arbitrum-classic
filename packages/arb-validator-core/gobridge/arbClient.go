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
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var errReorgError = errors.New("reorg occured")
var headerRetryDelay = time.Second * 2
var maxFetchAttempts = 5

func newEthClient(ethURL string) *goEthdata {
	return getGoEth(ethURL)
}

func (c *goEthdata) SubscribeBlockHeaders(ctx context.Context, startBlockID *common.BlockId) (<-chan arbbridge.MaybeBlockId, error) {
	blockIDChan := make(chan arbbridge.MaybeBlockId, 100)

	if startBlockID == nil {
		startBlockID = c.rootBlock
	}
	blockIDChan <- arbbridge.MaybeBlockId{BlockId: startBlockID}
	prevBlockID := startBlockID
	go func() {
		defer close(blockIDChan)

		for {
			var nextBlock *common.BlockId
			for {
				c.goEthMutex.Lock()
				var headerHash common.Hash
				nextBlock = nil
				if prevBlockID.Height.Cmp(c.LastMinedBlock.Height) < 0 {
					nextBlock = c.blockNumbers[prevBlockID.Height.AsInt().Uint64()+1]
					headerHash = c.parentHashes[*nextBlock]
				}
				c.goEthMutex.Unlock()

				if nextBlock != nil {
					if headerHash != prevBlockID.HeaderHash {
						blockIDChan <- arbbridge.MaybeBlockId{Err: errReorgError}
						return
					}
					break
				}

				select {
				case <-ctx.Done():
					// Getting header must have failed due to context cancellation
					return
				default:
				}

				// Header was not found so wait before checking again
				time.Sleep(headerRetryDelay)
			}
			prevBlockID = nextBlock
			blockIDChan <- arbbridge.MaybeBlockId{BlockId: prevBlockID}
		}
	}()

	return blockIDChan, nil
}

func (c *goEthdata) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newArbFactoryWatcher(address, c)
}

func (c *goEthdata) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newRollupWatcher(address, c)
}

func (c *goEthdata) NewExecutionChallengeWatcher(address common.Address) (arbbridge.ExecutionChallengeWatcher, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newChallengeWatcher(address, c)
}

func (c *goEthdata) NewMessagesChallengeWatcher(address common.Address) (arbbridge.MessagesChallengeWatcher, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newChallengeWatcher(address, c)
}

func (c *goEthdata) NewInboxTopChallengeWatcher(address common.Address) (arbbridge.InboxTopChallengeWatcher, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newChallengeWatcher(address, c)
}

func (c *goEthdata) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newOneStepProof(address, c)
}

func (c *goEthdata) GetBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return c.ethWallet[account], nil
}

func (c *goEthdata) CurrentBlockId(ctx context.Context) (*common.BlockId, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return c.LastMinedBlock, nil
}

func (c *goEthdata) BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	block, err := c.getBlockFromHeight(height)
	if err != nil {
		errstr := fmt.Sprintln("block height", height, " not found")
		return nil, errors.New(errstr)
	}
	return block, nil
}

type TransOpts struct {
	From common.Address
}

type GoArbAuthClient struct {
	*goEthdata
	fromAddr common.Address
}

func NewEthAuthClient(ethURL string, fromAddr common.Address) *GoArbAuthClient {
	client := newEthClient(ethURL)
	client.goEthMutex.Lock()
	defer client.goEthMutex.Unlock()

	client.ethWallet[fromAddr] = big.NewInt(1000) // give client a default balance of 1000
	return &GoArbAuthClient{
		goEthdata: client,
		fromAddr:  fromAddr,
	}
}

func (c *GoArbAuthClient) Address() common.Address {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return c.fromAddr
}

func (c *GoArbAuthClient) NewArbFactory(address common.Address) (arbbridge.ArbFactory, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newArbFactory(address, c)
}

func (c *GoArbAuthClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return getRollupContract(address, c)
}

func (c *GoArbAuthClient) NewGlobalInbox(address common.Address) (arbbridge.GlobalInbox, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newGlobalInbox(address, c)
}

func (c *GoArbAuthClient) NewChallengeFactory(address common.Address) (arbbridge.ChallengeFactory, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	//return newChallengeFactory(address, c)
	return nil, nil
}

func (c *GoArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newExecutionChallenge(address, c)
}

func (c *GoArbAuthClient) NewMessagesChallenge(address common.Address) (arbbridge.MessagesChallenge, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newMessagesChallenge(address, c)
}

func (c *GoArbAuthClient) NewInboxTopChallenge(address common.Address) (arbbridge.InboxTopChallenge, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newInboxTopChallenge(address, c)
}

func (c *GoArbAuthClient) DeployChallengeTest(ctx context.Context, challengeFactory common.Address) (arbbridge.ChallengeTester, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	tester, err := NewChallengeTester(c)
	if err != nil {
		return nil, err
	}
	return tester, nil
}

func (c *GoArbAuthClient) DeployOneStepProof(ctx context.Context) (arbbridge.OneStepProof, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	osp, err := newOneStepProof(c.fromAddr, c.goEthdata)
	return osp, err
}
