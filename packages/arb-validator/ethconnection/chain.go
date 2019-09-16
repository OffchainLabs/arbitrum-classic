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

package ethconnection

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type ArbAddresses struct {
	VMCreatorAddress   string `json:"vmCreator"`
	GlobalPendingInbox string `json:"globalPendingInbox"`
}

func waitForReceipt(ctx context.Context, client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	for {
		select {
		case _ = <-time.After(time.Second):
			receipt, err := client.TransactionReceipt(context.Background(), hash)
			if err != nil {
				return nil, err
			}
			if receipt.Status != 1 {
				return nil, errors.New("Transaction failed")
			}
			return receipt, nil
		case _ = <-ctx.Done():
			return nil, errors.New("Receipt not found")
		}
	}
}
