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

package structures

import (
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"google.golang.org/protobuf/proto"
	"testing"
)

var contractPath = arbos.Path()

func TestMarshalNode(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := NewInitialNode(mach.Clone())
	ms := NewRandomMessageStack(20)
	nextNode := NewRandomNodeFromValidPrev(node, ms, 5)

	checkpointContext := ckptcontext.NewCheckpointContext()
	nodeBuf := nextNode.MarshalForCheckpoint(checkpointContext, true)

	rawNodeData, err := proto.Marshal(nodeBuf)
	if err != nil {
		t.Fatal(err)
	}

	nodeBuf2 := &NodeBuf{}
	if err := proto.Unmarshal(rawNodeData, nodeBuf2); err != nil {
		t.Fatal(err)
	}
	unmarshaledNode, err := nodeBuf2.UnmarshalFromCheckpoint(checkpointContext)
	if err != nil {
		t.Fatal(err)
	}

	if !nextNode.EqualsFull(unmarshaledNode) {
		t.Error("nodes don't match")
	}
}
