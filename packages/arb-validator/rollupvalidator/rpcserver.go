/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package rollupvalidator

import (
	"context"
	"net/http"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/validatorserver"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
)

// Server provides an interface for interacting with a a running coordinator
type RPCServer struct {
	*Server
}

// NewServer returns a new instance of the Server class
func NewRPCServer(
	man *rollupmanager.Manager,
	maxCallTime time.Duration,
) (*RPCServer, error) {
	server, err := NewServer(man, maxCallTime)
	return &RPCServer{Server: server}, err
}

// FindLogs takes a set of parameters and return the list of all logs that match
// the query
func (m *RPCServer) FindLogs(
	r *http.Request,
	args *validatorserver.FindLogsArgs,
	reply *validatorserver.FindLogsReply,
) error {
	ret, err := m.Server.FindLogs(context.Background(), args)
	if ret != nil {
		*reply = *ret
	}
	return err
}

func (m *RPCServer) GetOutputMessage(r *http.Request, args *validatorserver.GetOutputMessageArgs, reply *validatorserver.GetOutputMessageReply) error {
	ret, err := m.Server.GetOutputMessage(context.Background(), args)
	if ret != nil {
		*reply = *ret
	}
	return err
}

// GetMessageResult returns the value output by the VM in response to the
//message with the given hash
func (m *RPCServer) GetMessageResult(
	r *http.Request,
	args *validatorserver.GetMessageResultArgs,
	reply *validatorserver.GetMessageResultReply,
) error {
	ret, err := m.Server.GetMessageResult(context.Background(), args)
	if ret != nil {
		*reply = *ret
	}
	return err
}

// GetAssertionCount returns the total number of finalized assertions
func (m *RPCServer) GetAssertionCount(
	r *http.Request,
	args *validatorserver.GetAssertionCountArgs,
	reply *validatorserver.GetAssertionCountReply,
) error {
	ret, err := m.Server.GetAssertionCount(context.Background(), args)
	if ret != nil {
		*reply = *ret
	}
	return err
}

// GetVMInfo returns current metadata about this VM
func (m *RPCServer) GetVMInfo(
	r *http.Request,
	args *validatorserver.GetVMInfoArgs,
	reply *validatorserver.GetVMInfoReply,
) error {
	ret, err := m.Server.GetVMInfo(context.Background(), args)
	if ret != nil {
		*reply = *ret
	}
	return err
}

// CallMessage takes a request from a client to process in a temporary context
//and return the result
func (m *RPCServer) CallMessage(
	r *http.Request,
	args *validatorserver.CallMessageArgs,
	reply *validatorserver.CallMessageReply,
) error {
	ret, err := m.Server.CallMessage(context.Background(), args)
	if ret != nil {
		*reply = *ret
	}
	return err
}
