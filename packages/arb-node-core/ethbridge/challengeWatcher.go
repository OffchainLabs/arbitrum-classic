/*
 * Copyright 2021, Offchain Labs, Inc.
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
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
)

var bisectedID ethcommon.Hash

func init() {
	parsedChallenge, err := abi.JSON(strings.NewReader(ethbridgecontracts.ChallengeABI))
	if err != nil {
		panic(err)
	}
	bisectedID = parsedChallenge.Events["Bisected"].ID
}

type ChallengeTurn uint8

const (
	NONE ChallengeTurn = iota
	ASSERTER_TURN
	CHALLENGER_TURN
)

type ChallengeWatcher struct {
	con     *ethbridgecontracts.Challenge
	address ethcommon.Address
	client  ethutils.EthClient
}

func NewChallengeWatcher(address ethcommon.Address, client ethutils.EthClient) (*ChallengeWatcher, error) {
	con, err := ethbridgecontracts.NewChallenge(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &ChallengeWatcher{
		con:     con,
		address: address,
		client:  client,
	}, nil
}

func (c *ChallengeWatcher) Address() common.Address {
	return common.NewAddressFromEth(c.address)
}

func (c *ChallengeWatcher) Turn(ctx context.Context) (ChallengeTurn, error) {
	rawTurn, err := c.con.Turn(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return ChallengeTurn(rawTurn), nil
}

func (c *ChallengeWatcher) Asserter(ctx context.Context) (common.Address, error) {
	asserter, err := c.con.Asserter(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return common.NewAddressFromEth(asserter), nil
}

func (c *ChallengeWatcher) Challenger(ctx context.Context) (common.Address, error) {
	challenger, err := c.con.Challenger(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return common.NewAddressFromEth(challenger), nil
}

func (c *ChallengeWatcher) CurrentResponder(ctx context.Context) (common.Address, error) {
	responder, err := c.con.CurrentResponder(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return common.NewAddressFromEth(responder), nil
}

func (c *ChallengeWatcher) ChallengeState(ctx context.Context) (common.Hash, error) {
	challengeState, err := c.con.ChallengeState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, errors.WithStack(err)
	}
	return common.NewHashFromEth(challengeState), nil
}

func (c *ChallengeWatcher) LookupBisection(ctx context.Context, challengeState common.Hash) (*core.Bisection, error) {
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Addresses: []ethcommon.Address{c.address},
		Topics:    [][]ethcommon.Hash{{bisectedID}, {challengeState.ToEthHash()}},
	}
	logs, err := c.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(logs) == 0 {
		return nil, nil
	}
	if len(logs) > 1 {
		return nil, errors.New("too many matching bisections")
	}

	parsedLog, err := c.con.ParseBisected(logs[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cuts := make([]core.Cut, 0, len(parsedLog.ChainHashes))
	for _, ch := range parsedLog.ChainHashes {
		cuts = append(cuts, core.NewSimpleCut(ch))
	}
	challengeSegment := &core.ChallengeSegment{
		Start:  parsedLog.ChallengedSegmentStart,
		Length: parsedLog.ChallengedSegmentLength,
	}
	return &core.Bisection{
		ChallengedSegment: challengeSegment,
		Cuts:              cuts,
	}, nil
}
