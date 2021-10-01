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

package binding

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/pkg/errors"
)

type Contract struct {
	File     string
	Contract string
}

func GenerateContractsList(folder string, contractNames []string) []Contract {
	var contracts []Contract
	for _, con := range contractNames {
		contracts = append(contracts, Contract{
			File:     filepath.Join(folder, con+".sol", con+".json"),
			Contract: con,
		})
	}
	return contracts
}

type HardHatArtifact struct {
	Format       string        `json:"_format"`
	ContractName string        `json:"contractName"`
	SourceName   string        `json:"sourceName"`
	Abi          []interface{} `json:"abi"`
	Bytecode     string        `json:"bytecode"`
}

func GenerateBinding(artifactPath, contract, pkg string) error {
	code, err := GenerateBindingFromFile(artifactPath, pkg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(contract+".go", []byte(code), 0777)
}

func EthbridgeArtifactsFolder() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("bad path")
	}
	return filepath.Join(filepath.Dir(filename), "../../arb-bridge-eth/build/contracts"), nil
}

func ArbOSArtifactsFolder() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("bad path")
	}
	return filepath.Join(filepath.Dir(filename), "../../arb-os/contracts/artifacts/arbos/builtin"), nil
}

func PeripheralsArtifactsFolder() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("bad path")
	}
	return filepath.Join(filepath.Dir(filename), "../../arb-bridge-peripherals/build/contracts"), nil
}

func GenerateBindingFromFile(inputFile, pkg string) (string, error) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", err
	}
	artifact := HardHatArtifact{}
	if err := json.Unmarshal(data, &artifact); err != nil {
		return "", err
	}
	abiData, err := json.Marshal(artifact.Abi)
	if err != nil {
		return "", err
	}
	return bind.Bind(
		[]string{artifact.ContractName},
		[]string{string(abiData)},
		[]string{artifact.Bytecode},
		nil,
		pkg,
		bind.LangGo,
		nil,
		nil,
	)
}

func GenerateBindingFromSolidity(file, pkg string) (string, error) {
	contracts, err := compiler.CompileSolidity("solc", file)
	if err != nil {
		return "", err
	}
	var contractNames []string
	var abis []string
	var byteCodes []string
	for key := range contracts {
		conName := strings.Split(key, ":")[1]
		contractNames = append(contractNames, conName)
		encodedAbi, err := json.Marshal(contracts[key].Info.AbiDefinition)
		if err != nil {
			return "", err
		}
		abis = append(abis, string(encodedAbi))
		byteCodes = append(byteCodes, contracts[key].Code)
	}
	return bind.Bind(
		contractNames,
		abis,
		byteCodes,
		nil,
		pkg,
		bind.LangGo,
		nil,
		nil,
	)
}
