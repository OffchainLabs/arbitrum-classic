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

package arbostest

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestBlockHash(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)

	runMessage(t, mach, initMsg(), chain)

	connAddress, err := deployContract(t, mach, sender, hexutil.MustDecode(arbostestcontracts.OpCodesBin), big.NewInt(0), nil)
	failIfError(t, err)

	snap := snapshot.NewSnapshot(mach.Clone(), chainTime, message.ChainAddressToID(chain), big.NewInt(2))
	ret, err := snap.BasicCall(hexutil.MustDecode("0x9663f88f"), connAddress)
	failIfError(t, err)
	succeededTxCheck(t, ret)
	if !bytes.Equal(ret.ReturnData, common.Hash{}.Bytes()) {
		t.Error("Unexpected block hash result")
	}
}
