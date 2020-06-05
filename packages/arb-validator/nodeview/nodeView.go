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

package nodeview

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"google.golang.org/protobuf/proto"
)

type NodeView struct {
	ns machine.NodeStore
	sr *ckptcontext.SimpleRestore
}

func New(ns machine.NodeStore, db machine.CheckpointStorage) *NodeView {
	return &NodeView{
		ns: ns,
		sr: ckptcontext.NewSimpleRestore(db),
	}
}

func (nv *NodeView) GetNode(height uint64, hash common.Hash) (*structures.Node, error) {
	nodeData, err := nv.ns.GetNode(height, hash)
	if err != nil {
		return nil, err
	}
	nodeBuf := &structures.NodeBuf{}
	if err := proto.Unmarshal(nodeData, nodeBuf); err != nil {
		return nil, err
	}
	nd := nodeBuf.UnmarshalFromCheckpoint(nv.sr)
	return nd, nil
}

func (nv *NodeView) GetNodeHeight(hash common.Hash) (uint64, error) {
	return nv.ns.GetNodeHeight(hash)
}

func (nv *NodeView) GetNodeHash(height uint64) (common.Hash, error) {
	return nv.ns.GetNodeHash(height)
}
