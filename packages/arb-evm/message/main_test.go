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

package message

//
//import (
//	"context"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	ethcommon "github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
//	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
//	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
//	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
//	"os"
//	"testing"
//)

//var tester *ethbridgetestcontracts.MessageTester

//func TestMain(m *testing.M) {
//	ctx := context.Background()
//	backend, pks := test.SimulatedBackend()
//	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
//	auth := bind.NewKeyedTransactor(pks[0])
//	authClient, err := ethbridge.NewEthAuthClient(ctx, client, auth)
//	if err != nil {
//		logger.Fatal().Stack().Err(err).Send()
//	}
//
//	messageTesterAddr, _, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
//		return ethbridgetestcontracts.DeployMessageTester(auth, client)
//	})
//	if err != nil {
//		logger.Fatal().Stack().Err(err).Send()
//	}
//
//	client.Commit()
//
//	tester, err = ethbridgetestcontracts.NewMessageTester(messageTesterAddr, client)
//	if err != nil {
//		logger.Fatal().Stack().Err(err).Send()
//	}
//
//	code := m.Run()
//	os.Exit(code)
//}
