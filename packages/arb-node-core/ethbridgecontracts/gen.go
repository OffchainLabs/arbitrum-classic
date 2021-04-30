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

package ethbridgecontracts

import (
	"path/filepath"

	"github.com/offchainlabs/arbitrum/packages/arb-util/binding"
)

//go:generate go run createBindings.go

func RunBindingGen() error {
	artifactsFolder, err := binding.EthbridgeArtifactsFolder()
	if err != nil {
		return err
	}
	base := filepath.Join(artifactsFolder, "contracts")
	oz := filepath.Join(artifactsFolder, "@openzeppelin")
	contracts := binding.GenerateContractsList(
		filepath.Join(base, "challenge"),
		[]string{"Challenge"},
	)
	contracts = append(contracts, binding.GenerateContractsList(
		filepath.Join(base, "rollup"),
		[]string{"Rollup", "RollupCreator", "INode"},
	)...)
	contracts = append(contracts, binding.GenerateContractsList(
		filepath.Join(base, "bridge"),
		[]string{"Bridge", "Inbox", "Outbox", "OutboxEntry"},
	)...)
	contracts = append(contracts, binding.GenerateContractsList(
		filepath.Join(base, "validator"),
		[]string{"Validator", "ValidatorUtils", "ValidatorWalletCreator"},
	)...)
	contracts = append(contracts, binding.GenerateContractsList(
		filepath.Join(oz, "contracts", "token", "ERC20"),
		[]string{"IERC20"},
	)...)

	for _, con := range contracts {
		err := binding.GenerateBinding(con.File, con.Contract, "ethbridgecontracts")
		if err != nil {
			return err
		}
	}
	return nil
}
