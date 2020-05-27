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

package txaggregator

import (
	context "context"
	"net/http"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

// RPCServer converts Server over to the interface required for automatic RPC
// interface generation
type RPCServer struct {
	*Server
}

func NewRPCServer(ctx context.Context, globalInbox arbbridge.GlobalInbox, rollupAddress common.Address) *RPCServer {
	return &RPCServer{Server: NewServer(ctx, globalInbox, rollupAddress)}
}

// SendTransaction converts the server implementation of SendTransaction to the
// required rpc server interface
func (m *RPCServer) SendTransaction(r *http.Request, args *SendTransactionArgs, reply *SendTransactionReply) error {
	ret, err := m.Server.SendTransaction(context.Background(), args)
	if ret != nil {
		*reply = *ret
	}
	return err
}
