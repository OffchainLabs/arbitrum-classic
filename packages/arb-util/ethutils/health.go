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

package ethutils

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

type OpenEthereumNetPeers struct {
	//Result field for parityNetPeers
	Active    int                `json:"active"`
	Connected int                `json:"connected"`
	Max       int                `json:"max"`
	Peers     []OpenEthereumPeer `json:"peers"`
}

type OpenEthereumPeer struct {
	//Peer client ID
	Id string `json:"id"`
}

type ParityPeers struct {
	url         string
	peerMinimum int
}

func (p ParityPeers) Execute(ctx context.Context) (interface{}, error) {
	client, err := rpc.DialContext(ctx, p.url)
	if err != nil {
		return nil, err
	}
	ret := OpenEthereumNetPeers{}
	if err := client.CallContext(ctx, &ret, "parity_netPeers"); err != nil {
		return nil, err
	}
	details := fmt.Sprintf("L1 peer count = %v", ret.Connected)
	if ret.Connected < p.peerMinimum {
		return details, errors.Errorf("not enough L1 peers, needed %v", ret.Connected)
	}
	return details, nil
}

func (p ParityPeers) Name() string {
	return "parity-peers"
}

type L1Ready struct {
	Url          string
	MaxBlockDiff uint64
}

func NewL1ReadyCheck(url string, config configuration.Healthcheck) L1Ready {
	return L1Ready{
		Url:          url,
		MaxBlockDiff: uint64(config.MaxL1BlockDiff),
	}
}

func (c L1Ready) Execute(ctx context.Context) (interface{}, error) {
	client, err := ethclient.DialContext(ctx, c.Url)
	if err != nil {
		return "failed dialing node", err
	}
	syncing, err := client.SyncProgress(ctx)
	if err != nil {
		return "couldn't get syncing info", err
	}
	if syncing == nil {
		return "not syncing", nil
	}
	if syncing.HighestBlock-syncing.CurrentBlock > c.MaxBlockDiff {
		return nil, errors.Errorf("L1 node still syncing %v/%v", syncing.CurrentBlock, syncing.HighestBlock)
	}
	return "synced within tolerance", nil
}

func (c L1Ready) Name() string {
	return "l1-syncing"
}
