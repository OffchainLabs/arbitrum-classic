/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
)

func MarshalAssertionForCheckpoint(ctx *checkpointing.CheckpointContext, a *protocol.ExecutionAssertion) *ExecutionAssertionBuf {
	messages := make([]*common.HashBuf, 0, len(a.OutMsgs))
	for _, msg := range a.OutMsgs {
		ctx.AddValue(msg)
		messages = append(messages, msg.Hash().MarshalToBuf())
	}
	logs := make([]*common.HashBuf, 0, len(a.Logs))
	for _, msg := range a.OutMsgs {
		ctx.AddValue(msg)
		logs = append(logs, msg.Hash().MarshalToBuf())
	}
	return &ExecutionAssertionBuf{
		AfterHash:    a.AfterHash.MarshalToBuf(),
		DidInboxInsn: a.DidInboxInsn,
		NumGas:       a.NumGas,
		Messages:     messages,
		Logs:         logs,
	}
}

func (a *ExecutionAssertionBuf) UnmarshalFromCheckpoint(ctx checkpointing.RestoreContext) *protocol.ExecutionAssertion {
	messages := make([]value.Value, 0, len(a.Logs))
	for _, valLog := range a.Messages {
		val := ctx.GetValue(valLog.Unmarshal())
		messages = append(messages, val)
	}

	logs := make([]value.Value, 0, len(a.Logs))
	for _, valLog := range a.Logs {
		val := ctx.GetValue(valLog.Unmarshal())
		logs = append(logs, val)
	}
	return &protocol.ExecutionAssertion{
		AfterHash:    a.AfterHash.Unmarshal(),
		DidInboxInsn: a.DidInboxInsn,
		NumGas:       a.NumGas,
		OutMsgs:      messages,
		Logs:         logs,
	}
}
