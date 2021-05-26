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
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Snapshot struct {
	mach            machine.Machine
	time            inbox.ChainTime
	nextInboxSeqNum *big.Int
	chainId         *big.Int
	arbosVersion    uint64
}

func NewSnapshot(mach machine.Machine, time inbox.ChainTime, chainId *big.Int, lastInboxSeq *big.Int) (*Snapshot, error) {
	snap := &Snapshot{
		mach:            mach,
		time:            time,
		nextInboxSeqNum: new(big.Int).Add(lastInboxSeq, big.NewInt(1)),
		chainId:         chainId,
	}
	ver, err := snap.ArbOSVersion()
	if err != nil {
		return nil, err
	}
	snap.arbosVersion = ver.Uint64()
	return snap, nil
}

// AddMessage can only be called if the snapshot is uniquely owned
// If an error is returned, s is unmodified
func (s *Snapshot) AddMessage(msg message.Message, sender common.Address, targetHash common.Hash) (*evm.TxResult, error) {
	mach := s.mach.Clone()
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	inboxMsg := message.NewInboxMessage(msg, sender, s.nextInboxSeqNum, big.NewInt(0), chainTime)
	res, _, err := runTx(mach, inboxMsg, targetHash, 100000000000)
	if err != nil {
		return nil, err
	}
	s.mach = mach
	s.nextInboxSeqNum = new(big.Int).Add(s.nextInboxSeqNum, big.NewInt(1))
	return res, nil
}

// AdvanceTime can only be called if the snapshot is uniquely owned
func (s *Snapshot) AdvanceTime(time inbox.ChainTime) {
	s.time = time
}

func (s *Snapshot) Clone() *Snapshot {
	return &Snapshot{
		mach: s.mach,
		time: inbox.ChainTime{
			BlockNum:  s.time.BlockNum.Clone(),
			Timestamp: new(big.Int).Set(s.time.Timestamp),
		},
		nextInboxSeqNum: new(big.Int).Set(s.nextInboxSeqNum),
		chainId:         new(big.Int).Set(s.chainId),
	}
}

func (s *Snapshot) Height() *common.TimeBlocks {
	return s.time.BlockNum
}

func (s *Snapshot) EstimateGas(tx *types.Transaction, aggregator, sender common.Address, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	if s.arbosVersion < 3 {

		var dest common.Address
		if tx.To() != nil {
			copy(dest[:], tx.To().Bytes())
		}
		msg := message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      new(big.Int).SetUint64(tx.Gas()),
				GasPriceBid: tx.GasPrice(),
				DestAddress: dest,
				Payment:     tx.Value(),
				Data:        tx.Data(),
			},
		}
		return s.Call(msg, sender)
	} else {
		gasEstimationMessage, err := message.NewGasEstimationMessage(aggregator, big.NewInt(0), message.NewCompressedECDSAFromEth(tx))
		if err != nil {
			return nil, nil, err
		}
		targetHash := hashing.SoliditySHA3(hashing.Uint256(s.chainId), hashing.Uint256(s.nextInboxSeqNum))
		targetHash = hashing.SoliditySHA3(hashing.Bytes32(targetHash), hashing.Uint256(big.NewInt(0)))
		return s.tryTx(gasEstimationMessage, sender, targetHash, maxAVMGas)
	}
}

func (s *Snapshot) Call(msg message.ContractTransaction, sender common.Address) (*evm.TxResult, []value.Value, error) {
	targetHash := hashing.SoliditySHA3(hashing.Uint256(s.chainId), hashing.Uint256(s.nextInboxSeqNum))
	return s.tryTx(message.NewSafeL2Message(msg), sender, targetHash, 100000000000)
}

func (s *Snapshot) tryTx(msg message.Message, sender common.Address, targetHash common.Hash, maxGas uint64) (*evm.TxResult, []value.Value, error) {
	inboxMsg := message.NewInboxMessage(msg, sender, s.nextInboxSeqNum, big.NewInt(0), s.time)
	return runTx(s.mach.Clone(), inboxMsg, targetHash, maxGas)
}

func (s *Snapshot) basicCall(data []byte, dest common.Address) (*evm.TxResult, error) {
	msg := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: dest,
			Payment:     big.NewInt(0),
			Data:        data,
		},
	}
	res, _, err := s.Call(msg, common.Address{})
	return res, err
}

func checkValidResult(res *evm.TxResult) error {
	if res.ResultCode == evm.ReturnCode {
		return nil
	}
	return errors.Errorf("error processing call %v", res.ResultCode)
}

func (s *Snapshot) GetBalance(account common.Address) (*big.Int, error) {
	res, err := s.basicCall(arbos.GetBalanceData(account), common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseBalanceResult(res.ReturnData)
}

func (s *Snapshot) GetTransactionCount(account common.Address) (*big.Int, error) {
	res, err := s.basicCall(arbos.TransactionCountData(account), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseTransactionCountResult(res.ReturnData)
}

func (s *Snapshot) GetCode(account common.Address) ([]byte, error) {
	res, err := s.basicCall(arbos.GetCodeData(account), common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseCodeResult(res.ReturnData)
}

func (s *Snapshot) GetStorageAt(account common.Address, index *big.Int) (*big.Int, error) {
	res, err := s.basicCall(arbos.StorageAtData(account, index), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseGetStorageAtResult(res.ReturnData)
}

func (s *Snapshot) ArbOSVersion() (*big.Int, error) {
	res, err := s.basicCall(arbos.ArbOSVersionData(), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseArbOSVersionResult(res.ReturnData)
}

func (s *Snapshot) GetPricesInWei() ([6]*big.Int, error) {
	res, err := s.basicCall(arbos.GetPricesInWeiData(), common.NewAddressFromEth(arbos.ARB_GAS_INFO_ADDRESS))
	if err != nil {
		return [6]*big.Int{}, err
	}
	if err := checkValidResult(res); err != nil {
		return [6]*big.Int{}, err
	}
	return arbos.ParseGetPricesInWeiResult(res.ReturnData)
}

func runTx(mach machine.Machine, msg inbox.InboxMessage, targetHash common.Hash, maxGas uint64) (*evm.TxResult, []value.Value, error) {
	assertion, debugPrints, steps, err := mach.ExecuteAssertionAdvanced(maxGas, false, nil, []inbox.InboxMessage{msg}, true)
	if err != nil {
		return nil, nil, err
	}

	// If the machine wasn't able to run and it reports that it is currently
	// blocked, return the block reason to give the client more information
	// as opposed to just returning a general "call produced no output"
	if br := mach.IsBlocked(true); steps == 0 && br != nil {
		return nil, debugPrints, errors.Errorf("can't produce solution since machine is blocked %v", br)
	}

	avmLogs := assertion.Logs
	if len(avmLogs) == 0 {
		fmt.Println("Running message didn't produce log")
		fmt.Println("Gas used", assertion.NumGas)
		fmt.Println("mach state after", mach)
		return nil, debugPrints, errors.New("no logs produced by tx")
	}

	res, err := evm.NewTxResultFromValue(avmLogs[len(avmLogs)-1])
	if err != nil {
		return nil, debugPrints, err
	}

	if res.IncomingRequest.MessageID != targetHash {
		return nil, debugPrints, errors.Errorf("call got unexpected result %v instead of %v", res.IncomingRequest.MessageID, targetHash)
	}

	return res, debugPrints, nil
}
