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

package dump

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
)

func getFile(taskName string, blockNumber uint64, perFolder, perFile uint64) (*os.File, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get current work dir failed: %w", err)
	}

	logPath := path.Join(cwd, "dump", taskName, strconv.FormatUint(blockNumber/perFolder, 10), strconv.FormatUint(blockNumber/perFile, 10)+".log")
	fmt.Printf("log path: %v, block: %v\n", logPath, blockNumber)
	if err := os.MkdirAll(path.Dir(logPath), 0755); err != nil {
		return nil, fmt.Errorf("mkdir for all parents [%v] failed: %w", path.Dir(logPath), err)
	}

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return nil, fmt.Errorf("create file %s failed: %w", logPath, err)
	}
	return file, nil
}

func BlockDumpLogger(block *types.Block, perFolder, perFile uint64) error {
	file, err := getFile("blocks", block.NumberU64(), perFolder, perFile)
	if err != nil {
		return err
	}
	defer file.Close()

	entry := map[string]interface{}{
		"timestamp":      block.Time(),
		"blockNumber":    block.NumberU64(),
		"blockHash":      block.Hash(),
		"parentHash":     block.ParentHash(),
		"gasLimit":       block.GasLimit(),
		"gasUsed":        block.GasUsed(),
		"miner":          block.Coinbase(),
		"difficulty":     block.Difficulty(),
		"nonce":          block.Nonce(),
		"size":           block.Size(),
		"extdataGasUsed": block.ExtDataGasUsed(),
	}
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(entry); err != nil {
		return fmt.Errorf("failed to encode transaction entry %w", err)
	}
	return nil
}
