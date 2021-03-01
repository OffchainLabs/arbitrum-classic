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

package ethbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

var rollupCreatedID ethcommon.Hash
var stakeCreatedID ethcommon.Hash
var challengeStartedID ethcommon.Hash
var challengeCompletedID ethcommon.Hash
var rollupRefundedID ethcommon.Hash
var rollupPrunedID ethcommon.Hash
var rollupStakeMovedID ethcommon.Hash
var rollupAssertedID ethcommon.Hash
var rollupConfirmedID ethcommon.Hash
var confirmedAssertionID ethcommon.Hash

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(ethbridgecontracts.ArbRollupABI))
	if err != nil {
		panic(err)
	}
	rollupCreatedID = parsedRollup.Events["RollupCreated"].ID
	stakeCreatedID = parsedRollup.Events["RollupStakeCreated"].ID
	challengeStartedID = parsedRollup.Events["RollupChallengeStarted"].ID
	challengeCompletedID = parsedRollup.Events["RollupChallengeCompleted"].ID
	rollupRefundedID = parsedRollup.Events["RollupStakeRefunded"].ID
	rollupPrunedID = parsedRollup.Events["RollupPruned"].ID
	rollupStakeMovedID = parsedRollup.Events["RollupStakeMoved"].ID
	rollupAssertedID = parsedRollup.Events["RollupAsserted"].ID
	rollupConfirmedID = parsedRollup.Events["RollupConfirmed"].ID
	confirmedAssertionID = parsedRollup.Events["ConfirmedAssertion"].ID
}

type ethRollupWatcher struct {
	ArbRollup *ethbridgecontracts.ArbRollup

	rollupAddress ethcommon.Address
	client        ethutils.EthClient
}

func newRollupWatcher(
	rollupAddress ethcommon.Address,
	client ethutils.EthClient,
) (*ethRollupWatcher, error) {
	arbitrumRollupContract, err := ethbridgecontracts.NewArbRollup(rollupAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to arbRollup")
	}

	return &ethRollupWatcher{
		ArbRollup:     arbitrumRollupContract,
		rollupAddress: rollupAddress,
		client:        client,
	}, nil
}

func (vm *ethRollupWatcher) generateTopics() [][]ethcommon.Hash {
	addressIndex := ethcommon.Hash{}
	copy(
		addressIndex[:],
		ethcommon.LeftPadBytes(vm.rollupAddress.Bytes(), 32),
	)
	return [][]ethcommon.Hash{
		{
			stakeCreatedID,
			challengeStartedID,
			challengeCompletedID,
			rollupRefundedID,
			rollupPrunedID,
			rollupStakeMovedID,
			rollupAssertedID,
			rollupConfirmedID,
			confirmedAssertionID,
		},
	}
}

func (vm *ethRollupWatcher) GetEvents(
	ctx context.Context,
	blockId *common.BlockId,
	_ *big.Int,
) ([]arbbridge.Event, error) {
	bh := blockId.HeaderHash.ToEthHash()
	rollupLogs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &bh,
		Addresses: []ethcommon.Address{vm.rollupAddress},
		Topics:    vm.generateTopics(),
	})
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(rollupLogs))
	for _, evmLog := range rollupLogs {
		event, err := vm.processEvents(
			getLogChainInfo(evmLog),
			evmLog,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (vm *ethRollupWatcher) GetAllEvents(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]arbbridge.Event, error) {
	inboxLogs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []ethcommon.Address{vm.rollupAddress},
		Topics:    vm.generateTopics(),
	})
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(inboxLogs))
	for _, evmLog := range inboxLogs {
		event, err := vm.processEvents(getLogChainInfo(evmLog), evmLog)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (vm *ethRollupWatcher) processEvents(
	chainInfo arbbridge.ChainInfo,
	ethLog types.Log,
) (arbbridge.Event, error) {
	switch ethLog.Topics[0] {
	case stakeCreatedID:
		eventVal, err := vm.ArbRollup.ParseRollupStakeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeCreatedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			NodeHash:  eventVal.NodeHash,
		}, nil
	case challengeStartedID:
		eventVal, err := vm.ArbRollup.ParseRollupChallengeStarted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeStartedEvent{
			ChainInfo:  chainInfo,
			Asserter:   common.NewAddressFromEth(eventVal.Asserter),
			Challenger: common.NewAddressFromEth(eventVal.Challenger),
			ChallengeType: valprotocol.ChildType(
				eventVal.ChallengeType.Uint64(),
			),
			ChallengeContract: common.NewAddressFromEth(
				eventVal.ChallengeContract,
			),
		}, nil
	case challengeCompletedID:
		eventVal, err := vm.ArbRollup.ParseRollupChallengeCompleted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeCompletedEvent{
			ChainInfo: chainInfo,
			Winner:    common.NewAddressFromEth(eventVal.Winner),
			Loser:     common.NewAddressFromEth(eventVal.Loser),
			ChallengeContract: common.NewAddressFromEth(
				eventVal.ChallengeContract,
			),
		}, nil
	case rollupRefundedID:
		eventVal, err := vm.ArbRollup.ParseRollupStakeRefunded(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeRefundedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
		}, nil
	case rollupPrunedID:
		eventVal, err := vm.ArbRollup.ParseRollupPruned(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.PrunedEvent{
			ChainInfo: chainInfo,
			Leaf:      eventVal.Leaf,
		}, nil
	case rollupStakeMovedID:
		eventVal, err := vm.ArbRollup.ParseRollupStakeMoved(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeMovedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			Location:  eventVal.ToNodeHash,
		}, nil
	case rollupAssertedID:
		eventVal, err := vm.ArbRollup.ParseRollupAsserted(ethLog)
		if err != nil {
			return nil, err
		}
		params := &valprotocol.AssertionParams{
			NumSteps:             eventVal.NumSteps,
			ImportedMessageCount: eventVal.ImportedMessageCount,
		}

		return arbbridge.AssertedEvent{
			ChainInfo:        chainInfo,
			PrevLeafHash:     eventVal.Fields[0],
			MaxInboxTop:      eventVal.Fields[1],
			MaxInboxCount:    eventVal.InboxCount,
			AssertionParams:  params,
			AfterMachineHash: eventVal.Fields[2],
			AfterInboxAcc:    eventVal.Fields[3],
			NumGas:           eventVal.NumArbGas,
			LastMessageHash:  eventVal.Fields[4],
			MessageCount:     eventVal.MessageCount,
			LastLogHash:      eventVal.Fields[5],
			LogCount:         eventVal.LogCount,
		}, nil
	case rollupConfirmedID:
		eventVal, err := vm.ArbRollup.ParseRollupConfirmed(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedEvent{
			ChainInfo: chainInfo,
			NodeHash:  eventVal.NodeHash,
		}, nil
	case confirmedAssertionID:
		eventVal, err := vm.ArbRollup.ParseConfirmedAssertion(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedAssertionEvent{
			ChainInfo:   chainInfo,
			LogsAccHash: hashSliceToHashes(eventVal.LogsAccHash),
		}, nil
	default:
		return nil, errors.New("unknown arbitrum event type")
	}
}

func (vm *ethRollupWatcher) GetParams(
	ctx context.Context,
) (valprotocol.ChainParams, error) {
	invalid := valprotocol.ChainParams{}
	callOpts := &bind.CallOpts{Context: ctx}
	rawParams, err := vm.ArbRollup.VmParams(callOpts)
	if err != nil {
		return invalid, err
	}
	stakeRequired, err := vm.ArbRollup.GetStakeRequired(callOpts)
	if err != nil {
		return invalid, err
	}

	stakeToken, err := vm.ArbRollup.GetStakeToken(callOpts)
	if err != nil {
		return invalid, err
	}
	return valprotocol.ChainParams{
		StakeRequirement: stakeRequired,
		StakeToken:       common.NewAddressFromEth(stakeToken),
		GracePeriod: common.TimeTicks{
			Val: rawParams.GracePeriodTicks,
		},
		MaxExecutionSteps:       rawParams.MaxExecutionSteps,
		ArbGasSpeedLimitPerTick: rawParams.ArbGasSpeedLimitPerTick.Uint64(),
	}, nil
}

func (vm *ethRollupWatcher) InboxAddress(
	ctx context.Context,
) (common.Address, error) {
	addr, err := vm.ArbRollup.GlobalInbox(&bind.CallOpts{Context: ctx})
	return common.NewAddressFromEth(addr), err
}

func (vm *ethRollupWatcher) GetCreationInfo(
	ctx context.Context,
) (common.Hash, arbbridge.ChainInfo, common.Hash, *big.Int, error) {
	addressIndex := ethcommon.Hash{}
	copy(
		addressIndex[:],
		ethcommon.LeftPadBytes(vm.rollupAddress.Bytes(), 32),
	)
	logs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []ethcommon.Address{vm.rollupAddress},
		Topics:    [][]ethcommon.Hash{{rollupCreatedID}},
	})
	if err != nil {
		return common.Hash{}, arbbridge.ChainInfo{}, common.Hash{}, nil, err
	}
	if len(logs) == 0 {
		return common.Hash{},
			arbbridge.ChainInfo{},
			common.Hash{},
			nil,
			errors.New("chain does not exist")
	}
	if len(logs) > 1 {
		return common.Hash{},
			arbbridge.ChainInfo{},
			common.Hash{},
			nil,
			errors.New("more than one chain created with same address")
	}
	ev, err := vm.ArbRollup.ParseRollupCreated(logs[0])
	if err != nil {
		return common.Hash{}, arbbridge.ChainInfo{}, common.Hash{}, nil, err
	}

	header, err := vm.client.HeaderByNumber(ctx, new(big.Int).SetUint64(logs[0].BlockNumber))
	if err != nil {
		return common.Hash{},
			arbbridge.ChainInfo{},
			common.Hash{},
			nil,
			err
	}
	return common.NewHashFromEth(logs[0].TxHash), getLogChainInfo(logs[0]), ev.InitVMHash, new(big.Int).SetUint64(header.Time), nil
}

func (vm *ethRollupWatcher) GetVersion(ctx context.Context) (string, error) {
	return vm.ArbRollup.VERSION(&bind.CallOpts{Context: ctx})
}

func (vm *ethRollupWatcher) IsStaked(address common.Address) (bool, error) {
	return vm.ArbRollup.IsStaked(nil, address.ToEthAddress())
}

func (vm *ethRollupWatcher) VerifyArbChain(ctx context.Context, machHash common.Hash) error {
	simulatedBackend, pks := test.SimulatedBackend()
	simulatedClient := &ethutils.SimulatedEthClient{SimulatedBackend: simulatedBackend}
	simulatedAuth := bind.NewKeyedTransactor(pks[0])
	authClient, err := NewEthAuthClient(ctx, simulatedClient, simulatedAuth)
	if err != nil {
		return err
	}

	simulatedRollupAddr, _, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgecontracts.DeployArbRollup(auth, simulatedBackend)
	})
	if err != nil {
		return err
	}
	simulatedRollup, err := ethbridgecontracts.NewArbRollup(simulatedRollupAddr, simulatedBackend)
	if err != nil {
		return err
	}
	simulatedBackend.Commit()
	validEthBridgeVersion, err := simulatedRollup.VERSION(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}
	ethbridgeVersion, err := vm.GetVersion(ctx)
	if err != nil {
		return err
	}

	if ethbridgeVersion != validEthBridgeVersion {
		return errors.Errorf("VM has EthBridge version %v, but validator implements version %v."+
			" To find a validator version which supports your EthBridge, visit "+
			"https://offchainlabs.com/ethbridge-version-support",
			ethbridgeVersion, validEthBridgeVersion)
	}

	_, _, initialVMHash, _, err := vm.GetCreationInfo(ctx)
	if err != nil {
		return err
	}

	if machHash != initialVMHash {
		return errors.Errorf("ArbChain was initialized with VM with hash %v, but local validator has VM with hash %v", initialVMHash, machHash)
	}
	return nil
}
