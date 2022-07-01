/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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

package arbostest

import (
	"context"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbosmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestInit(t *testing.T) {
	ctx := context.Background()
	cmach, err := cmachine.New(*arbosfile)
	failIfError(t, err)
	mach := arbosmachine.New(cmach)
	assertion, _, _, err := mach.ExecuteAssertion(ctx, 10000000000, false, nil)
	test.FailIfError(t, err)
	t.Log("Startup used", assertion.NumGas, "gas")
}
