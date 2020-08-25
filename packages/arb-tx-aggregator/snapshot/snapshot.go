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

package snapshot

import (
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"math/big"
)

type Snapshot struct {
	mach            machine.Machine
	time            inbox.ChainTime
	nextInboxSeqNum *big.Int
	chainId         *big.Int
}

func NewSnapshot(mach machine.Machine, time inbox.ChainTime, chainId *big.Int, lastInboxSeq *big.Int) *Snapshot {
	return &Snapshot{
		mach:            mach,
		time:            time,
		nextInboxSeqNum: new(big.Int).Add(lastInboxSeq, big.NewInt(1)),
		chainId:         chainId,
	}
}

func NewSnapshotWithMessage(snap *Snapshot, msg message.Message, sender common.Address) (*Snapshot, *evm.TxResult, error) {
	mach := snap.mach.Clone()
	inboxMsg := message.NewInboxMessage(msg, sender, snap.nextInboxSeqNum, snap.time)
	res, err := runTx(mach, inboxMsg, snap.chainId)
	if err != nil {
		return nil, nil, err
	}
	return &Snapshot{
		mach:            mach,
		time:            snap.time,
		nextInboxSeqNum: new(big.Int).Add(snap.nextInboxSeqNum, big.NewInt(1)),
		chainId:         snap.chainId,
	}, res, nil
}

func (s *Snapshot) Height() *common.TimeBlocks {
	return s.time.BlockNum
}

func (s *Snapshot) Call(msg message.ContractTransaction, sender common.Address) (*evm.TxResult, error) {
	return s.TryTx(message.NewSafeL2Message(msg), sender)
}

func (s *Snapshot) TryTx(msg message.Message, sender common.Address) (*evm.TxResult, error) {
	inboxMsg := message.NewInboxMessage(msg, sender, s.nextInboxSeqNum, s.time)
	return runTx(s.mach.Clone(), inboxMsg, s.chainId)
}

func (s *Snapshot) makeBasicCall(data []byte, dest common.Address) (*evm.TxResult, error) {
	msg := message.ContractTransaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		DestAddress: dest,
		Payment:     big.NewInt(0),
		Data:        data,
	}
	return s.Call(msg, common.Address{})
}

func (s *Snapshot) GetBalance(account common.Address) (*big.Int, error) {
	res, err := s.makeBasicCall(getBalanceData(account), common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS))
	if err != nil {
		return nil, err
	}
	return parseBalanceResult(res)
}

func (s *Snapshot) GetTransactionCount(account common.Address) (*big.Int, error) {
	res, err := s.makeBasicCall(getTransactionCountData(account), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	return parseTransactionCountResult(res)
}

func (s *Snapshot) GetCode(account common.Address) ([]byte, error) {
	res, err := s.makeBasicCall(getCodeData(account), common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS))
	if err != nil {
		return nil, err
	}
	return parseCodeResult(res)
}

func (s *Snapshot) GetStorageAt(account common.Address, index *big.Int) (*big.Int, error) {
	res, err := s.makeBasicCall(getStorageAtData(account, index), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	return parseGetStorageAtResult(res)
}

func runTx(mach machine.Machine, msg inbox.InboxMessage, chainId *big.Int) (*evm.TxResult, error) {
	assertion, steps := mach.ExecuteAssertion(100000000, []inbox.InboxMessage{msg}, 0)

	// If the machine wasn't able to run and it reports that it is currently
	// blocked, return the block reason to give the client more information
	// as opposed to just returning a general "call produced no output"
	if br := mach.IsBlocked(true); steps == 0 && br != nil {
		return nil, fmt.Errorf("can't produce solution since machine is blocked %v", br)
	}

	avmLogs := assertion.ParseLogs()
	if len(avmLogs) != 1 {
		return nil, fmt.Errorf("unexpected log count %v", len(avmLogs))
	}

	res, err := evm.NewTxResultFromValue(avmLogs[0])
	if err != nil {
		return nil, err
	}

	targetHash := hashing.SoliditySHA3(hashing.Uint256(chainId), hashing.Uint256(msg.InboxSeqNum))
	if res.IncomingRequest.MessageID != targetHash {
		return nil, fmt.Errorf("call got unexpected result %v instead of %v", res.IncomingRequest.MessageID, targetHash)
	}

	return res, nil
}
