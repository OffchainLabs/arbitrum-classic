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

package ethutils

import (
	"context"
	"github.com/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func CallCheck(ctx context.Context, client EthClient, from common.Address, contractAddress common.Address, contractABI abi.ABI, method string, params ...interface{}) error {
	// Pack the input, call and unpack the results
	input, err := contractABI.Pack(method, params...)
	if err != nil {
		return err
	}
	var (
		msg    = ethereum.CallMsg{From: from, To: &contractAddress, Data: input}
		output []byte
	)

	output, err = client.PendingCallContract(ctx, msg)
	if err != nil {
		return errors.WithStack(err)
	}

	if len(output) < 69 {
		return errors.Errorf("%v had short output %v, %v", method, len(output), output)
	}
	length := new(big.Int).SetBytes(output[36:68])
	if uint64(len(output)) < 68+length.Uint64()+1 {
		return errors.Errorf("%v had short output %v, %v", method, len(output), output)
	}
	return errors.Errorf("%v returned val: %v", method, string(output[68:68+length.Uint64()]))
}
