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

package aggregator

import (
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

func GenerateRPCServer(server *Server) (*rpc.Server, error) {
	arbServer := NewRPCServer(server)
	s := rpc.NewServer()
	s.RegisterCodec(
		json.NewCodec(),
		"application/json",
	)
	s.RegisterCodec(
		json.NewCodec(),
		"application/json;charset=UTF-8",
	)

	if err := s.RegisterService(arbServer, "Aggregator"); err != nil {
		return nil, err
	}
	return s, nil
}
