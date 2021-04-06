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

package arbos

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"strings"
)

var (
	giveOwnershipABI         abi.Method
	startArbOSUpgradeABI     abi.Method
	continueArbOSUpgradeABI  abi.Method
	finishArbOSUpgradeABI    abi.Method
	setFeesEnabledABI        abi.Method
	setFairGasPriceSenderABI abi.Method
)

func init() {
	arbowner, err := abi.JSON(strings.NewReader(arboscontracts.ArbOwnerABI))
	if err != nil {
		panic(err)
	}

	giveOwnershipABI = arbowner.Methods["giveOwnership"]
	startArbOSUpgradeABI = arbowner.Methods["startCodeUpload"]
	continueArbOSUpgradeABI = arbowner.Methods["continueCodeUpload"]
	finishArbOSUpgradeABI = arbowner.Methods["finishCodeUploadAsArbosUpgrade"]
	setFeesEnabledABI = arbowner.Methods["setFeesEnabled"]
	setFairGasPriceSenderABI = arbowner.Methods["setFairGasPriceSender"]
}

func GiveOwnershipData(newOwnerAddr common.Address) []byte {
	return makeFuncData(giveOwnershipABI, newOwnerAddr.ToEthAddress())
}

func StartArbOSUpgradeData() []byte {
	return makeFuncData(startArbOSUpgradeABI)
}

func ContinueArbOSUpgradeData(data []byte) []byte {
	return makeFuncData(continueArbOSUpgradeABI, data)
}

func FinishArbOSUpgradeData() []byte {
	return makeFuncData(finishArbOSUpgradeABI)
}

func SetFairGasPriceSender(sender common.Address) []byte {
	return makeFuncData(setFairGasPriceSenderABI, sender)
}

func SetFeesEnabled(enabled bool) []byte {
	return makeFuncData(setFeesEnabledABI, enabled)
}
