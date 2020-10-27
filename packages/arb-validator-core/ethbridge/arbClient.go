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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type EthArbClient struct {
	client ethutils.EthClient
}

func NewEthClient(client ethutils.EthClient) *EthArbClient {
	return &EthArbClient{client}
}

var reorgError = errors.New("reorg occured")
var headerRetryDelay = time.Second * 2
var maxFetchAttempts = 5

func (c *EthArbClient) SubscribeBlockHeadersAfter(ctx context.Context, prevBlockId *common.BlockId) (<-chan arbbridge.MaybeBlockId, error) {
	blockIdChan := make(chan arbbridge.MaybeBlockId, 100)
	if err := c.subscribeBlockHeadersAfter(ctx, prevBlockId, blockIdChan); err != nil {
		return nil, err
	}
	return blockIdChan, nil
}

func (c *EthArbClient) SubscribeBlockHeaders(ctx context.Context, startBlockId *common.BlockId) (<-chan arbbridge.MaybeBlockId, error) {
	blockIdChan := make(chan arbbridge.MaybeBlockId, 100)

	startHeader, err := c.client.HeaderByHash(ctx, startBlockId.HeaderHash.ToEthHash())
	if err != nil {
		return nil, errors2.Wrapf(err, "can't find initial header %v", startBlockId)
	}
	blockIdChan <- arbbridge.MaybeBlockId{BlockId: startBlockId, Timestamp: new(big.Int).SetUint64(startHeader.Time)}
	prevBlockId := startBlockId
	if err := c.subscribeBlockHeadersAfter(ctx, prevBlockId, blockIdChan); err != nil {
		return nil, err
	}
	return blockIdChan, nil
}

func (c *EthArbClient) subscribeBlockHeadersAfter(ctx context.Context, prevBlockId *common.BlockId, blockIdChan chan<- arbbridge.MaybeBlockId) error {
	prevBlockIdCheck, err := c.BlockIdForHeight(ctx, prevBlockId.Height)
	if err != nil {
		return err
	}

	if !prevBlockId.Equals(prevBlockIdCheck) {
		return fmt.Errorf("can't subscribe to headers, block hash %v doesn't match expected value %v", prevBlockIdCheck, prevBlockId)
	}

	go func() {
		defer close(blockIdChan)

		for {
			var blockInfo *ethutils.BlockInfo
			fetchErrorCount := 0
			for {
				var err error
				targetHeight := new(big.Int).Add(prevBlockId.Height.AsInt(), big.NewInt(1))
				blockInfo, err = c.client.BlockInfoByNumber(ctx, targetHeight)
				if err == nil && (*big.Int)(blockInfo.Number).Cmp(targetHeight) == 0 {
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

			if blockInfo.ParentHash != prevBlockId.HeaderHash.ToEthHash() {
				blockIdChan <- arbbridge.MaybeBlockId{Err: reorgError}
				return
			}

			prevBlockId = &common.BlockId{
				Height:     common.NewTimeBlocks((*big.Int)(blockInfo.Number)),
				HeaderHash: common.NewHashFromEth(blockInfo.Hash),
			}
			blockIdChan <- arbbridge.MaybeBlockId{BlockId: prevBlockId, Timestamp: new(big.Int).SetUint64(uint64(blockInfo.Time))}
		}
	}()
	return nil
}

func (c *EthArbClient) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	return newArbFactoryWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return newRollupWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewGlobalInboxWatcher(address common.Address, rollupAddress common.Address) (arbbridge.GlobalInboxWatcher, error) {
	return newGlobalInboxWatcher(address.ToEthAddress(), rollupAddress.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewExecutionChallengeWatcher(address common.Address) (arbbridge.ExecutionChallengeWatcher, error) {
	return newExecutionChallengeWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewInboxTopChallengeWatcher(address common.Address) (arbbridge.InboxTopChallengeWatcher, error) {
	return newInboxTopChallengeWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewIERC20Watcher(address common.Address) (arbbridge.IERC20Watcher, error) {
	return newIERC20Watcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) GetBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	return c.client.BalanceAt(ctx, account.ToEthAddress(), nil)
}

func (c *EthArbClient) BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
	var num *big.Int
	if height != nil {
		num = height.AsInt()
	}
	blockInfo, err := c.client.BlockInfoByNumber(ctx, num)
	if err != nil {
		return nil, err
	}
	return &common.BlockId{
		Height:     common.NewTimeBlocks((*big.Int)(blockInfo.Number)),
		HeaderHash: common.NewHashFromEth(blockInfo.Hash),
	}, nil
}

func (c *EthArbClient) TimestampForBlockHash(ctx context.Context, hash common.Hash) (*big.Int, error) {
	header, err := c.client.HeaderByHash(ctx, hash.ToEthHash())
	if err != nil {
		return nil, err
	}
	if header == nil {
		return nil, errors.New("couldn't get header at height")
	}
	return new(big.Int).SetUint64(header.Time), nil
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

func NewEthAuthClient(client ethutils.EthClient, auth *bind.TransactOpts) *EthArbAuthClient {
	return &EthArbAuthClient{
		EthArbClient: NewEthClient(client),
		auth:         &TransactAuth{auth: auth},
	}
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

func (c *EthArbAuthClient) NewGlobalInbox(address common.Address, rollupAddress common.Address) (arbbridge.GlobalInbox, error) {
	return newGlobalInbox(address.ToEthAddress(), rollupAddress.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewChallengeFactory(address common.Address) (arbbridge.ChallengeFactory, error) {
	return newChallengeFactory(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	return newExecutionChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewInboxTopChallenge(address common.Address) (arbbridge.InboxTopChallenge, error) {
	return newInboxTopChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewIERC20(address common.Address) (arbbridge.IERC20, error) {
	return newIERC20(address.ToEthAddress(), c.client, c.auth)
}
