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

package rollup

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	arbrollup "github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/rollup"
	"log"
	"math/big"
	"strings"
)

type Observer struct {
}

type RollupAssertedEvent struct {
	Fields               [6][32]byte
	ImportedMessageCount *big.Int
	TimeBoundsBlocks     [2]*big.Int
	DidInboxInsn         bool
	NumSteps             uint32
	NumArbGas            uint64
}

type RollupConfirmedEvent struct {
	NodeHash [32]byte
}

type RollupPrunedEvent struct {
	NodeHash [32]byte
}

type RollupStakeCreatedEvent struct {
	Staker      common.Address
	NodeHash    [32]byte
	BlockNumber *big.Int
}

type RollupStakeMovedEvent struct {
	Address    common.Address
	ToNodeHash [32]byte
}

type RollupStakeRefundedEvent struct {
	Staker common.Address
}

type RollupChallengeStartedEvent struct {
	Asserter          common.Address
	Challenger        common.Address
	ChallengeType     *big.Int
	ChallengeContract common.Address
}

type RollupChallengeCompletedEvent struct {
	ChallengeContract common.Address
	Winner            common.Address
	Loser             common.Address
}

func NewObserver(chain *Chain, clnt *ethclient.Client, rawUrl string) (*Observer, error) {
	pClient, err := ethclient.Dial(rawUrl)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{chain.rollupAddr},
	}
	logs := make(chan types.Log)
	sub, err := pClient.SubscribeFilterLogs(
		context.Background(),
		query,
		logs,
	)
	if err != nil {
		return nil, err
	}

	rollupAbi, err := abi.JSON(strings.NewReader(string(arbrollup.ArbRollupABI)))
	if err != nil {
		return nil, err
	}

	rollupAssertedSigHash := calcSigHash("RollupAsserted(bytes32[6],uint,uint128[2],bool,uint32,uint64)")
	rollupConfirmedSigHash := calcSigHash("RollupConfirmed(bytes32)")
	rollupPrunedSigHash := calcSigHash("RollupPruned(bytes32)")
	rollupStakeCreatedSigHash := calcSigHash("RollupStakeCreated(address,bytes32,uint)")
	rollupStakeMovedSigHash := calcSigHash("RollupStakeMoved(address,bytes32)")
	rollupStakeRefundedSigHash := calcSigHash("RollupStakeRefunded(address)")
	rollupChallengeStartedSigHash := calcSigHash("RollupChallengeStarted(address,address,uint,address)")
	rollupChallengeCompletedSigHash := calcSigHash("RollupChallengeCompleted(address,address,address)")

	go func() {
		defer sub.Unsubscribe()
		lastBlockNumberSeen := uint64(0)
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logs:
				if vLog.BlockNumber > lastBlockNumberSeen {
					lastBlockNumberSeen = vLog.BlockNumber
					chain.notifyNewBlockNumber(lastBlockNumberSeen)

				}
				switch vLog.Topics[0] {
				case rollupAssertedSigHash:
					var event RollupAssertedEvent
					err := rollupAbi.Unpack(&event, "RollupAsserted", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				case rollupConfirmedSigHash:
					var event RollupConfirmedEvent
					err := rollupAbi.Unpack(&event, "RollupConfirmed", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				case rollupPrunedSigHash:
					var event RollupPrunedEvent
					err := rollupAbi.Unpack(&event, "RollupPruned", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				case rollupStakeCreatedSigHash:
					var event RollupStakeCreatedEvent
					err := rollupAbi.Unpack(&event, "RollupStakeCreated", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				case rollupStakeMovedSigHash:
					var event RollupStakeMovedEvent
					err := rollupAbi.Unpack(&event, "RollupStakeMoved", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				case rollupStakeRefundedSigHash:
					var event RollupStakeRefundedEvent
					err := rollupAbi.Unpack(&event, "RollupStakeRefunded", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				case rollupChallengeStartedSigHash:
					var event RollupChallengeStartedEvent
					err := rollupAbi.Unpack(&event, "RollupChallengeStarted", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				case rollupChallengeCompletedSigHash:
					var event RollupChallengeCompletedEvent
					err := rollupAbi.Unpack(&event, "RollupChallengeCompleted", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					// do operation for event
				default:
					log.Fatal("unknown log event type")
				}
			}
		}
	}()
	return &Observer{}, nil
}

func calcSigHash(sig string) common.Hash {
	return crypto.Keccak256Hash([]byte(sig))
}
