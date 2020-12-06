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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"strings"
	"testing"
)

func TestFailedNestedSend(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)
	chain := common.RandAddress()
	sender := common.RandAddress()
	dest := common.RandAddress()

	runMessage(t, mach, initMsg(), chain)
	depositEth(t, mach, sender, big.NewInt(1000))

	constructorData := hexutil.MustDecode(arbostestcontracts.FailedSendBin)

	failedSendAddress, err := deployContract(t, mach, sender, constructorData, big.NewInt(0), nil)
	failIfError(t, err)

	failedSend, err := abi.JSON(strings.NewReader(arbostestcontracts.FailedSendABI))
	failIfError(t, err)

	failedSendABI := failedSend.Methods["send"]
	failedSendData, err := failedSendABI.Inputs.Pack(dest)
	failIfError(t, err)
	sendTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: failedSendAddress,
		Payment:     big.NewInt(300),
		Data:        append(failedSendABI.ID, failedSendData...),
	}
	_, err = runTransaction(t, mach, sendTx, sender)
	failIfError(t, err)
}
