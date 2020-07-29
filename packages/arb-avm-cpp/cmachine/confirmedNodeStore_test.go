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

package cmachine

import (
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"os"
	"testing"
)

var dePath = "dbPath"

var nodeHeight = uint64(3)
var nodeHash = common.Hash{54}
var nodeData = []byte{5, 6, 7}

var nodeHeight2 = uint64(5)
var nodeHash2 = common.Hash{56}

func TestConfirmedNodeStore(t *testing.T) {
	checkpointStorage, err := NewCheckpoint(dePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := checkpointStorage.Initialize(codeFile); err != nil {
		t.Fatal(err)
	}
	defer checkpointStorage.CloseCheckpointStorage()

	ns := checkpointStorage.GetConfirmedNodeStore()

	t.Run("EmptyTrue", func(t *testing.T) {
		if !ns.Empty() {
			t.Error("should be empty")
		}
	})

	t.Run("Put", func(t *testing.T) {
		if err := ns.PutNode(nodeHeight, nodeHash, nodeData); err != nil {
			t.Error(err)
		}
	})

	t.Run("EmptyFalse", func(t *testing.T) {
		if ns.Empty() {
			t.Error("shouldn't be empty")
		}
	})

	t.Run("MaxHeight", func(t *testing.T) {
		if ns.MaxHeight() != nodeHeight {
			t.Error("wrong height")
		}
	})

	t.Run("Get", func(t *testing.T) {
		_, err := ns.GetNode(nodeHeight2, nodeHash2)
		if err == nil {
			t.Error("should fail")
		}
		data, err := ns.GetNode(nodeHeight, nodeHash)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(data, nodeData) {
			t.Error("not equal")
		}
	})

	t.Run("GetHash", func(t *testing.T) {
		_, err := ns.GetNodeHash(nodeHeight2)
		if err == nil {
			t.Error("should fail")
		}
		hash, err := ns.GetNodeHash(nodeHeight)
		if err != nil {
			t.Error(err)
		}
		if hash != nodeHash {
			t.Error("not equal")
		}
	})

	t.Run("GetHeight", func(t *testing.T) {
		_, err := ns.GetNodeHeight(nodeHash2)
		if err == nil {
			t.Error("should fail")
		}
		height, err := ns.GetNodeHeight(nodeHash)
		if err != nil {
			t.Error(err)
		}
		if height != nodeHeight {
			t.Error("not equal")
		}
	})

	if err := os.RemoveAll(dePath); err != nil {
		t.Fatal(err)
	}
}
