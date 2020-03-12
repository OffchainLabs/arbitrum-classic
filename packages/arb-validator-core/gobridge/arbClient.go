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

var headerRetryDelay = time.Second * 2

func newEthClient(ethURL string) *goEthdata {
	return getGoEth(ethURL)
}

func (c *goEthdata) SubscribeBlockHeaders(ctx context.Context, startBlockID *common.BlockId) (<-chan arbbridge.MaybeBlockId, error) {
	blockIdChan := make(chan arbbridge.MaybeBlockId, 100)

	if startBlockID == nil {
		startBlockID = c.rootBlock
	}
	blockIdChan <- arbbridge.MaybeBlockId{BlockId: startBlockID}
	prevBlockID := startBlockID
	go func() {
		defer close(blockIdChan)

		for {
			var nextBlock *common.BlockId
			for {
				c.goEthMutex.Lock()
				nextBlock = nil
				if prevBlockID.Height.Cmp(c.LastMinedBlock.Height) < 0 {
					nextBlock = c.blockNumbers[prevBlockID.Height.AsInt().Uint64()+1]
				}
				c.goEthMutex.Unlock()

				if nextBlock != nil {
					// new block has been mined
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
			blockIdChan <- arbbridge.MaybeBlockId{BlockId: prevBlockID}
		}
	}()

	return blockIdChan, nil
}

func (c *goEthdata) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newArbFactoryWatcher(address, c)
}

func (c *goEthdata) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	return newRollupWatcher(address, c), nil
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

	client.ethWallet[fromAddr] = big.NewInt(10 * 1000 * 1000 * 1000 * 1000 * 1000) // give client a default balance
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
	return newChallengeFactory(address, c)
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
	if !challengeFactory.Equals(c.challengeFactoryContract) {
		return nil, errors.New("invalid challengeFactory")
	}
	c.challengeTester = &challengeTester{
		contract:         c.getNextAddress(),
		challengeFactory: c.challengeFactory,
		client:           c.client,
	}
	return NewChallengeTester(c.challengeTester.contract, c)
}

func (c *GoArbAuthClient) DeployOneStepProof(ctx context.Context) (arbbridge.OneStepProof, error) {
	c.goEthMutex.Lock()
	defer c.goEthMutex.Unlock()
	c.oneStepProof = &oneStepProof{
		oneStepProofContract: c.getNextAddress(),
		client:               c.client,
	}
	osp, err := newOneStepProof(c.oneStepProof.oneStepProofContract, c.goEthdata)
	return osp, err
}

func deployRollupFactory(m *goEthdata) {
	m.rollups = make(map[common.Address]*arbRollup)
	m.arbFactoryContract = &arbFactory{
		rollupContractAddress: m.getNextAddress(),
		client:                nil,
	}
}

func deployGlobalInbox(m *goEthdata) {
	m.globalInbox = &globalInbox{
		ethData:         m,
		inbox:           make(map[common.Address]*inbox),
		contractAddress: m.getNextAddress(),
	}
}

func deployChallengeFactory(m *goEthdata) {
	m.challengeFactory = &challengeFactory{
		challengeFactoryContract: m.getNextAddress(),
		client:                   m,
	}
	m.challengeFactory.challenges = make(map[common.Address]*challenge)
}
