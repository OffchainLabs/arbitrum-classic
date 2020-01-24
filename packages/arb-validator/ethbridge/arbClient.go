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

package ethbridge

import (
	"context"
	"errors"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/challengetester"
)

type EthArbClient struct {
	client *ethclient.Client
}

func NewEthClient(ethURL string) (*EthArbClient, error) {
	client, err := ethclient.Dial(ethURL)
	return &EthArbClient{client}, err
}

var reorgError = errors.New("reorg occured")
var headerRetryDelay = time.Second * 2
var maxFetchAttempts = 5

func (c *EthArbClient) SubscribeBlockHeaders(ctx context.Context, startBlockId *structures.BlockId) (<-chan arbbridge.MaybeBlockId, error) {
	blockIdChan := make(chan arbbridge.MaybeBlockId, 100)

	blockIdChan <- arbbridge.MaybeBlockId{BlockId: startBlockId}
	prevBlockId := startBlockId
	go func() {
		defer close(blockIdChan)

		for {
			var nextHeader *types.Header
			fetchErrorCount := 0
			for {
				var err error
				nextHeader, err = c.client.HeaderByNumber(ctx, new(big.Int).Add(prevBlockId.Height.AsInt(), big.NewInt(1)))
				if err == nil {
					// Got next header
					break
				}

				select {
				case <-ctx.Done():
					// Getting header must have failed due to context cancellation
					return
				default:
				}

				if err != nil && err.Error() != ethereum.NotFound.Error() {
					log.Printf("Failed to fetch next header on attempt %v with error: %v", fetchErrorCount, err)
					fetchErrorCount++
				}

				if fetchErrorCount >= maxFetchAttempts {
					blockIdChan <- arbbridge.MaybeBlockId{Err: err}
					return
				}

				// Header was not found so wait before checking again
				time.Sleep(headerRetryDelay)
			}

			if nextHeader.ParentHash != prevBlockId.HeaderHash.ToEthHash() {
				blockIdChan <- arbbridge.MaybeBlockId{Err: reorgError}
				return
			}

			prevBlockId = getBlockID(nextHeader)
			blockIdChan <- arbbridge.MaybeBlockId{BlockId: prevBlockId}
		}
	}()

	return blockIdChan, nil
}

func (c *EthArbClient) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	return newArbFactoryWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return newRollupWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewExecutionChallengeWatcher(address common.Address) (arbbridge.ExecutionChallengeWatcher, error) {
	return newExecutionChallengeWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewMessagesChallengeWatcher(address common.Address) (arbbridge.MessagesChallengeWatcher, error) {
	return newMessagesChallengeWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewPendingTopChallengeWatcher(address common.Address) (arbbridge.PendingTopChallengeWatcher, error) {
	return newPendingTopChallengeWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	return newOneStepProof(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) CurrentBlockId(ctx context.Context) (*structures.BlockId, error) {
	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return getBlockID(header), nil
}

func (c *EthArbClient) BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*structures.BlockId, error) {
	header, err := c.client.HeaderByNumber(ctx, height.AsInt())
	if err != nil {
		return nil, err
	}
	return getBlockID(header), nil
}

type TransactAuth struct {
	sync.Mutex
	auth *bind.TransactOpts
}

func (t *TransactAuth) getAuth(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     t.auth.From,
		Nonce:    t.auth.Nonce,
		Signer:   t.auth.Signer,
		Value:    t.auth.Value,
		GasPrice: t.auth.GasPrice,
		GasLimit: t.auth.GasLimit,
		Context:  ctx,
	}
}

type EthArbAuthClient struct {
	*EthArbClient
	auth *TransactAuth
}

func NewEthAuthClient(ethURL string, auth *bind.TransactOpts) (*EthArbAuthClient, error) {
	client, err := NewEthClient(ethURL)
	if err != nil {
		return nil, err
	}
	return &EthArbAuthClient{
		EthArbClient: client,
		auth:         &TransactAuth{auth: auth},
	}, nil
}

func (c *EthArbAuthClient) Address() common.Address {
	return common.NewAddressFromEth(c.auth.auth.From)
}

func (c *EthArbAuthClient) NewArbFactory(address common.Address) (arbbridge.ArbFactory, error) {
	return newArbFactory(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	return newRollup(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewPendingInbox(address common.Address) (arbbridge.PendingInbox, error) {
	return newPendingInbox(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewChallengeFactory(address common.Address) (arbbridge.ChallengeFactory, error) {
	return newChallengeFactory(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	return newExecutionChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewMessagesChallenge(address common.Address) (arbbridge.MessagesChallenge, error) {
	return newMessagesChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewPendingTopChallenge(address common.Address) (arbbridge.PendingTopChallenge, error) {
	return newPendingTopChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) DeployChallengeTest(ctx context.Context, challengeFactory common.Address) (*ChallengeTester, error) {
	c.auth.Lock()
	defer c.auth.Unlock()
	testerAddress, tx, _, err := challengetester.DeployChallengeTester(c.auth.auth, c.client, challengeFactory.ToEthAddress())
	if err != nil {
		return nil, err
	}
	if err := waitForReceipt(
		ctx,
		c.client,
		c.auth.auth.From,
		tx,
		"DeployChallengeTester",
	); err != nil {
		return nil, err
	}
	tester, err := NewChallengeTester(testerAddress, c.client, c.auth)
	if err != nil {
		return nil, err
	}
	return tester, nil
}
