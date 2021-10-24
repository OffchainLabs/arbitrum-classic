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

package configuration

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

type rpcMock struct {
	chainId uint64
}

func (s *rpcMock) ChainId() hexutil.Uint64 {
	fmt.Println("got chain id")
	return hexutil.Uint64(s.chainId)
}

func TestNodeConfig(t *testing.T) {
	s := rpc.NewServer()
	mock := &rpcMock{chainId: 1}
	err := s.RegisterName("eth", mock)
	test.FailIfError(t, err)
	server := &http.Server{
		Addr:    "127.0.0.1:5344",
		Handler: s,
	}
	defer func() {
		err := server.Shutdown(context.Background())
		test.FailIfError(t, err)
	}()
	go func() {
		_ = server.ListenAndServe()
	}()

	ctx := context.Background()

	mock.chainId = 1
	conf, _, _, _, err := ParseNode(ctx, []string{"arb-node", "--l1.url", "http://127.0.0.1:5344"})
	test.FailIfError(t, err)
	if conf.Rollup.Address != arbitrumOneRollupAddress {
		t.Error("didn't get correct rollup address")
	}

	mock.chainId = 4
	conf, _, _, _, err = ParseNode(ctx, []string{"arb-node", "--l1.url", "http://127.0.0.1:5344"})
	test.FailIfError(t, err)
	if conf.Rollup.Address != rinkebyTestnetRollupAddress {
		t.Error("didn't get correct rollup address")
	}

	mock.chainId = 10
	conf, _, _, _, err = ParseNode(ctx, []string{"--l1.url", "http://127.0.0.1:5344"})
	if err == nil {
		t.Error("config parsing should've failed for unknown chain")
	}
}
