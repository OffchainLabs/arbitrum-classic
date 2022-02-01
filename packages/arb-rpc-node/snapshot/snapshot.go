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

package snapshot

import (
	"context"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"

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

var logger = log.With().Caller().Stack().Str("component", "snapshot").Logger()

type Snapshot struct {
	mach                  machine.Machine
	time                  inbox.ChainTime
	nextInboxSeqNum       *big.Int
	chainId               *big.Int
	arbosVersion          uint64
	arbosRemappingEnabled bool
}

func NewSnapshot(ctx context.Context, mach machine.Machine, time inbox.ChainTime, lastInboxSeq *big.Int) (*Snapshot, error) {
	snap := &Snapshot{
		mach:            mach,
		time:            time,
		nextInboxSeqNum: new(big.Int).Add(lastInboxSeq, big.NewInt(1)),
	}

	ver, err := snap.ArbOSVersion(ctx)
	if err != nil {
		return nil, err
	}
	snap.arbosVersion = ver.Uint64()
	if snap.arbosVersion >= 27 {
		chainId, err := snap.ChainId(ctx)
		if err != nil {
			return nil, err
		}
		snap.chainId = chainId
	}

	if snap.arbosVersion >= 40 {
		arbOwnerMsg := message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1 << 30),
				GasPriceBid: snap.MaxGasPriceBid(),
				DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
				Payment:     big.NewInt(0),
				Data:        arbos.GetChainParameterData(arbos.EnableL1ContractAddressAliasingParamId),
			},
		}
		// Note: this .Call actually uses arbosRemappingEnabled which isn't set yet,
		// but that's fine because the zero address is never rewritten regardless.
		arbOwnerRes, _, err := snap.Call(ctx, arbOwnerMsg, common.Address{}, math.MaxUint64)
		if err != nil {
			return nil, err
		}
		if arbOwnerRes.ResultCode != evm.ReturnCode {
			return nil, errors.New("failed to query ArbOS address remapping state")
		}
		snap.arbosRemappingEnabled = new(big.Int).SetBytes(arbOwnerRes.ReturnData).Sign() > 0
	}

	return snap, nil
}

func (s *Snapshot) ArbosVersion() uint64 {
	return s.arbosVersion
}

func (s *Snapshot) MaxGasPriceBid() *big.Int {
	if s.arbosVersion >= 42 {
		return big.NewInt(1 << 60)
	} else {
		return big.NewInt(0)
	}
}

// AddMessage can only be called if the snapshot is uniquely owned
// If an error is returned, s is unmodified
func (s *Snapshot) AddMessage(ctx context.Context, msg message.Message, sender common.Address, targetHash common.Hash) (*evm.TxResult, error) {
	mach := s.mach.Clone()
	res, _, err := s.addMessage(ctx, msg, sender, targetHash, addMessageMaxAVMGas)
	if err != nil {
		// Revert the machine
		s.mach = mach
	}
	return res, err
}

const addMessageMaxAVMGas = 100000000000

// addMessage can only be called if the snapshot is uniquely owned
// leaves the machine in an undefined state on error
func (s *Snapshot) addMessage(ctx context.Context, msg message.Message, sender common.Address, targetHash common.Hash, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	inboxMsg := message.NewInboxMessage(msg, sender, s.nextInboxSeqNum, big.NewInt(0), chainTime)
	res, debugPrints, err := runTxUnchecked(ctx, s.mach, inboxMsg, maxAVMGas)
	if err != nil {
		return nil, nil, err
	}
	if res.IncomingRequest.MessageID != targetHash {
		return nil, nil, errors.Errorf("call got unexpected result %v instead of %v", res.IncomingRequest.MessageID, targetHash)
	}
	s.nextInboxSeqNum = new(big.Int).Add(s.nextInboxSeqNum, big.NewInt(1))
	return res, debugPrints, nil
}

// AdvanceTime can only be called if the snapshot is uniquely owned
func (s *Snapshot) AdvanceTime(time inbox.ChainTime) {
	s.time = time
}

func (s *Snapshot) Clone() *Snapshot {
	var chainId *big.Int
	if s.chainId != nil {
		chainId = new(big.Int).Set(s.chainId)
	}
	return &Snapshot{
		mach: s.mach.Clone(),
		time: inbox.ChainTime{
			BlockNum:  s.time.BlockNum.Clone(),
			Timestamp: new(big.Int).Set(s.time.Timestamp),
		},
		nextInboxSeqNum:       new(big.Int).Set(s.nextInboxSeqNum),
		chainId:               chainId,
		arbosRemappingEnabled: s.arbosRemappingEnabled,
	}
}

func (s *Snapshot) Height() *common.TimeBlocks {
	return s.time.BlockNum
}

func (s *Snapshot) EstimateGas(ctx context.Context, tx *types.Transaction, aggregator, sender common.Address, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
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
		return s.Call(ctx, msg, sender, maxAVMGas)
	} else {
		gasEstimationMessage, err := message.NewGasEstimationMessage(aggregator, big.NewInt(0), message.NewCompressedECDSAFromEth(tx))
		if err != nil {
			return nil, nil, err
		}
		var targetHash common.Hash
		if s.chainId != nil {
			targetHash = hashing.SoliditySHA3(hashing.Uint256(s.chainId), hashing.Uint256(s.nextInboxSeqNum))
			targetHash = hashing.SoliditySHA3(hashing.Bytes32(targetHash), hashing.Uint256(big.NewInt(0)))
		}
		inboxMsg := s.makeInboxMessage(gasEstimationMessage, sender)
		return runTx(ctx, s.mach.Clone(), inboxMsg, targetHash, maxAVMGas)
	}
}

func (s *Snapshot) EstimateRetryableGas(ctx context.Context, msg message.RetryableTx, sender common.Address, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	redeemGas := new(big.Int).Set(msg.MaxGas)
	redeemGasPriceBid := new(big.Int).Set(msg.GasPriceBid)
	msg.MaxGas = msg.MaxGas.SetUint64(0)
	msg.GasPriceBid = msg.GasPriceBid.SetUint64(0)
	inboxMsg1 := message.NewInboxMessage(msg, sender, s.nextInboxSeqNum, big.NewInt(0), s.time)
	var targetHash common.Hash
	var ticketHash common.Hash
	if s.chainId != nil {
		targetHash = hashing.SoliditySHA3(hashing.Uint256(s.chainId), hashing.Uint256(s.nextInboxSeqNum))
		ticketHash = hashing.SoliditySHA3(hashing.Bytes32(targetHash), hashing.Uint256(big.NewInt(0)))
	}

	redeemTx := types.NewTx(&types.LegacyTx{
		Nonce:    0,
		GasPrice: redeemGasPriceBid,
		Gas:      redeemGas.Uint64(),
		To:       &arbos.ARB_RETRYABLE_ADDRESS,
		Value:    big.NewInt(0),
		Data:     arbos.RedeemData(ticketHash),
	})
	gasEstimationMessage, err := message.NewGasEstimationMessage(common.Address{}, big.NewInt(0), message.NewCompressedECDSAFromEth(redeemTx))
	if err != nil {
		return nil, nil, err
	}
	estimateSeqNum := new(big.Int).Add(s.nextInboxSeqNum, big.NewInt(1))
	var targetHash2 common.Hash
	if s.chainId != nil {
		targetHash2 = hashing.SoliditySHA3(hashing.Uint256(s.chainId), hashing.Uint256(estimateSeqNum))
		targetHash2 = hashing.SoliditySHA3(hashing.Bytes32(targetHash2), hashing.Uint256(big.NewInt(0)))
	}
	redeemer := sender
	if s.arbosRemappingEnabled {
		redeemer = message.L2RemapAccount(redeemer)
	}
	inboxMsg2 := message.NewInboxMessage(gasEstimationMessage, redeemer, estimateSeqNum, big.NewInt(0), s.time)

	mach := s.mach.Clone()
	assertion, debugPrints, _, err := mach.ExecuteAssertionAdvanced(
		ctx,
		maxAVMGas,
		false,
		nil,
		[]inbox.InboxMessage{inboxMsg2, inboxMsg1},
		true,
		false,
	)
	if err != nil {
		return nil, nil, err
	}

	avmLogs := assertion.Logs
	if len(avmLogs) == 0 {
		return nil, nil, errors.New("no logs emitted processing retryable")
	}
	res, err := evm.NewTxResultFromValue(avmLogs[0])
	if err != nil {
		return nil, nil, err
	}
	if res.ResultCode != evm.ReturnCode {
		return nil, nil, errors.New("ticket creation failed")
	}
	if res.IncomingRequest.MessageID != targetHash {
		return nil, debugPrints, errors.Errorf("ticket creation got unexpected result %v instead of %v", res.IncomingRequest.MessageID, targetHash)
	}

	if len(avmLogs) == 2 {
		// Redeem must have failed
		res2, err := evm.NewTxResultFromValue(avmLogs[1])
		if err != nil {
			return nil, nil, err
		}
		if res2.ResultCode != evm.ReturnCode {
			return nil, nil, evm.HandleCallError(res2, false)
		} else {
			return nil, nil, errors.New("Redeem succeeded, but failed to trigger redeemed tx")
		}
	}

	if len(avmLogs) != 3 {
		return nil, nil, errors.Errorf("unexpected result count %v", len(avmLogs))
	}

	res2, err := evm.NewTxResultFromValue(avmLogs[2])
	if err != nil {
		return nil, nil, err
	}
	if res2.IncomingRequest.MessageID != targetHash2 {
		return nil, debugPrints, errors.Errorf("estimation got unexpected result %v instead of %v", res2.IncomingRequest.MessageID, targetHash2)
	}
	return res2, debugPrints, err
}

func (s *Snapshot) Call(ctx context.Context, msg message.ContractTransaction, sender common.Address, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	var targetHash common.Hash
	if s.chainId != nil {
		targetHash = hashing.SoliditySHA3(hashing.Uint256(s.chainId), hashing.Uint256(s.nextInboxSeqNum))
	}
	if s.arbosRemappingEnabled {
		sender = message.L1RemapAccount(sender)
	}
	inboxMsg := s.makeInboxMessage(message.NewSafeL2Message(msg), sender)
	return runTx(ctx, s.mach.Clone(), inboxMsg, targetHash, maxAVMGas)
}

type EthCallOverride struct {
	Nonce     *hexutil.Uint64                    `json:"nonce"`
	Code      *hexutil.Bytes                     `json:"code"`
	Balance   *hexutil.Big                       `json:"balance"`
	State     *map[ethcommon.Hash]ethcommon.Hash `json:"state"`
	StateDiff *map[ethcommon.Hash]ethcommon.Hash `json:"stateDiff"`
}

func (s *Snapshot) CallWithOverrides(ctx context.Context, msg message.ContractTransaction, sender common.Address, overrides *map[ethcommon.Address]EthCallOverride, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	snap := s.Clone()
	snap.mach = snap.mach.Clone()
	if overrides != nil {
		// We'll be mutating this scapshot so we need to make a copy
		for address, override := range *overrides {
			account := common.NewAddressFromEth(address)
			if override.Nonce != nil {
				err := snap.setNonce(ctx, account, uint64(*override.Nonce))
				if err != nil {
					return nil, nil, err
				}
			}
			if override.Balance != nil {
				err := snap.setBalance(ctx, account, override.Balance.ToInt())
				if err != nil {
					return nil, nil, err
				}
			}
			if override.Code != nil {
				err := snap.setCode(ctx, account, *override.Code)
				if err != nil {
					return nil, nil, err
				}
			}
			if override.State != nil {
				storage := make(map[common.Hash]common.Hash)
				for key, val := range *override.State {
					storage[common.NewHashFromEth(key)] = common.NewHashFromEth(val)
				}
				err := snap.setState(ctx, account, storage)
				if err != nil {
					return nil, nil, err
				}
			}
			if override.StateDiff != nil {
				for key, val := range *override.StateDiff {
					err := snap.store(ctx, account, common.NewHashFromEth(key), common.NewHashFromEth(val))
					if err != nil {
						return nil, nil, err
					}
				}
			}
		}
	}

	var targetHash common.Hash
	if snap.chainId != nil {
		targetHash = hashing.SoliditySHA3(hashing.Uint256(snap.chainId), hashing.Uint256(snap.nextInboxSeqNum))
	}
	if snap.arbosRemappingEnabled {
		sender = message.L1RemapAccount(sender)
	}
	inboxMsg := snap.makeInboxMessage(message.NewSafeL2Message(msg), sender)
	return runTx(ctx, snap.mach, inboxMsg, targetHash, maxAVMGas)
}

func (s *Snapshot) makeInboxMessage(msg message.Message, sender common.Address) inbox.InboxMessage {
	return message.NewInboxMessage(msg, sender, s.nextInboxSeqNum, big.NewInt(0), s.time)
}

func runTx(ctx context.Context, mach machine.Machine, inboxMsg inbox.InboxMessage, targetHash common.Hash, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	res, debugPrints, err := runTxUnchecked(ctx, mach, inboxMsg, maxAVMGas)
	if err != nil {
		return nil, nil, err
	}
	var emptyHash common.Hash
	if targetHash != emptyHash && res.IncomingRequest.MessageID != targetHash {
		return nil, debugPrints, errors.Errorf("call got unexpected result %v instead of %v", res.IncomingRequest.MessageID, targetHash)
	}
	return res, debugPrints, nil
}

func (s *Snapshot) basicCallUnsafe(ctx context.Context, data []byte, dest common.Address) (*evm.TxResult, []value.Value, error) {
	msg := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: s.MaxGasPriceBid(),
			DestAddress: dest,
			Payment:     big.NewInt(0),
			Data:        data,
		},
	}
	inboxMsg := message.NewInboxMessage(message.NewSafeL2Message(msg), common.Address{}, s.nextInboxSeqNum, big.NewInt(0), s.time)
	return runTxUnchecked(ctx, s.mach.Clone(), inboxMsg, 1000000000)
}

func (s *Snapshot) basicCall(ctx context.Context, data []byte, dest common.Address) (*evm.TxResult, error) {
	msg := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: s.MaxGasPriceBid(),
			DestAddress: dest,
			Payment:     big.NewInt(0),
			Data:        data,
		},
	}
	res, _, err := s.Call(ctx, msg, common.Address{}, math.MaxUint64)
	return res, err
}

func (s *Snapshot) basicAddMessage(ctx context.Context, data []byte, dest common.Address) (*evm.TxResult, error) {
	msg := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: s.MaxGasPriceBid(),
			DestAddress: dest,
			Payment:     big.NewInt(0),
			Data:        data,
		},
	}
	res, _, err := s.AddContractMessage(ctx, msg, common.Address{}, addMessageMaxAVMGas)
	return res, err
}

func (s *Snapshot) AddContractMessage(ctx context.Context, msg message.ContractTransaction, sender common.Address, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	if s.arbosRemappingEnabled && sender != (common.Address{}) {
		sender = message.L1RemapAccount(sender)
	}
	var targetHash common.Hash
	if s.chainId != nil {
		targetHash = hashing.SoliditySHA3(hashing.Uint256(s.chainId), hashing.Uint256(s.nextInboxSeqNum))
	}
	return s.addMessage(ctx, message.NewSafeL2Message(msg), sender, targetHash, maxAVMGas)
}

func (s *Snapshot) addArbosTestMessage(ctx context.Context, data []byte) error {
	res, err := s.basicAddMessage(ctx, (data), common.NewAddressFromEth(arbos.ARB_TEST_ADDRESS))
	if err != nil {
		return err
	}
	if err := checkValidResult(res); err != nil {
		return err
	}
	return nil
}

func checkValidResult(res *evm.TxResult) error {
	if res.ResultCode == evm.ReturnCode {
		return nil
	}
	return errors.Errorf("error processing call %v", res.ResultCode)
}

func (s *Snapshot) GetBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	res, err := s.basicCall(ctx, arbos.GetBalanceData(account), common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseBalanceResult(res.ReturnData)
}

func (s *Snapshot) GetTransactionCount(ctx context.Context, account common.Address) (*big.Int, error) {
	res, err := s.basicCall(ctx, arbos.TransactionCountData(account), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseTransactionCountResult(res.ReturnData)
}

func (s *Snapshot) GetCode(ctx context.Context, account common.Address) ([]byte, error) {
	res, err := s.basicCall(ctx, arbos.GetCodeData(account), common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseCodeResult(res.ReturnData)
}

func (s *Snapshot) GetStorageAt(ctx context.Context, account common.Address, index *big.Int) (*big.Int, error) {
	res, err := s.basicCall(ctx, arbos.StorageAtData(account, index), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseGetStorageAtResult(res.ReturnData)
}

func (s *Snapshot) setNonce(ctx context.Context, account common.Address, nonce uint64) error {
	return s.addArbosTestMessage(ctx, arbos.SetNonceData(account, nonce))
}

func (s *Snapshot) setBalance(ctx context.Context, account common.Address, balance *big.Int) error {
	return s.addArbosTestMessage(ctx, arbos.SetBalanceData(account, balance))
}

func (s *Snapshot) setState(ctx context.Context, account common.Address, storage map[common.Hash]common.Hash) error {
	return s.addArbosTestMessage(ctx, arbos.SetStateData(account, storage))
}

func (s *Snapshot) setCode(ctx context.Context, account common.Address, code []byte) error {
	return s.addArbosTestMessage(ctx, arbos.SetCodeData(account, code))
}

func (s *Snapshot) store(ctx context.Context, account common.Address, key, val common.Hash) error {
	return s.addArbosTestMessage(ctx, arbos.StoreData(account, key, val))
}

func (s *Snapshot) ArbOSVersion(ctx context.Context) (*big.Int, error) {
	res, _, err := s.basicCallUnsafe(ctx, arbos.ArbOSVersionData(), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseArbOSVersionResult(res.ReturnData)
}

func (s *Snapshot) ChainId(ctx context.Context) (*big.Int, error) {
	res, _, err := s.basicCallUnsafe(ctx, arbos.ChainIdData(), common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS))
	if err != nil {
		return nil, err
	}
	if err := checkValidResult(res); err != nil {
		return nil, err
	}
	return arbos.ParseChainIdResult(res.ReturnData)
}

func (s *Snapshot) GetPricesInWei(ctx context.Context) ([6]*big.Int, error) {
	res, err := s.basicCall(ctx, arbos.GetPricesInWeiData(), common.NewAddressFromEth(arbos.ARB_GAS_INFO_ADDRESS))
	if err != nil {
		return [6]*big.Int{}, err
	}
	if err := checkValidResult(res); err != nil {
		return [6]*big.Int{}, err
	}
	return arbos.ParseGetPricesInWeiResult(res.ReturnData)
}

func runTxUnchecked(ctx context.Context, mach machine.Machine, msg inbox.InboxMessage, maxAVMGas uint64) (*evm.TxResult, []value.Value, error) {
	assertion, debugPrints, steps, err := mach.ExecuteAssertionAdvanced(ctx, maxAVMGas, false, nil, []inbox.InboxMessage{msg}, true, false)
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
		logger.Info().Uint64("gasused", assertion.NumGas).Msg("Running message didn't produce log")
		return nil, debugPrints, errors.New("transaction ran out of gas")
	}

	res, err := evm.NewTxResultFromValue(avmLogs[len(avmLogs)-1])
	return res, debugPrints, err
}
