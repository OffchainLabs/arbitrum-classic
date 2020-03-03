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
	"errors"
	"math/big"
	"strings"

	errors2 "github.com/pkg/errors"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

var rollupCreatedID ethcommon.Hash
var rollupStakeCreatedID ethcommon.Hash
var rollupChallengeStartedID ethcommon.Hash
var rollupChallengeCompletedID ethcommon.Hash
var rollupRefundedID ethcommon.Hash
var rollupPrunedID ethcommon.Hash
var rollupStakeMovedID ethcommon.Hash
var rollupAssertedID ethcommon.Hash
var rollupConfirmedID ethcommon.Hash
var confirmedAssertionID ethcommon.Hash
var transactionMessageDeliveredID ethcommon.Hash
var ethDepositMessageDeliveredID ethcommon.Hash
var depositERC20MessageDeliveredID ethcommon.Hash
var depositERC721MessageDeliveredID ethcommon.Hash
var contractTransactionMessageDeliveredID ethcommon.Hash

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(rollup.ArbRollupABI))
	if err != nil {
		panic(err)
	}
	parsedInbox, err := abi.JSON(strings.NewReader(rollup.IGlobalInboxABI))
	if err != nil {
		panic(err)
	}
	rollupCreatedID = parsedRollup.Events["RollupCreated"].ID()
	rollupStakeCreatedID = parsedRollup.Events["RollupStakeCreated"].ID()
	rollupChallengeStartedID = parsedRollup.Events["RollupChallengeStarted"].ID()
	rollupChallengeCompletedID = parsedRollup.Events["RollupChallengeCompleted"].ID()
	rollupRefundedID = parsedRollup.Events["RollupStakeRefunded"].ID()
	rollupPrunedID = parsedRollup.Events["RollupPruned"].ID()
	rollupStakeMovedID = parsedRollup.Events["RollupStakeMoved"].ID()
	rollupAssertedID = parsedRollup.Events["RollupAsserted"].ID()
	rollupConfirmedID = parsedRollup.Events["RollupConfirmed"].ID()
	confirmedAssertionID = parsedRollup.Events["ConfirmedAssertion"].ID()

	transactionMessageDeliveredID = parsedInbox.Events["TransactionMessageDelivered"].ID()
	ethDepositMessageDeliveredID = parsedInbox.Events["EthDepositMessageDelivered"].ID()
	depositERC20MessageDeliveredID = parsedInbox.Events["ERC20DepositMessageDelivered"].ID()
	depositERC721MessageDeliveredID = parsedInbox.Events["ERC721DepositMessageDelivered"].ID()
	contractTransactionMessageDeliveredID = parsedInbox.Events["ContractTransactionMessageDelivered"].ID()
}

type ethRollupWatcher struct {
	ArbRollup   *rollup.ArbRollup
	GlobalInbox *rollup.IGlobalInbox

	rollupAddress ethcommon.Address
	inboxAddress  ethcommon.Address
	client        *ethclient.Client
}

func newRollupWatcher(rollupAddress ethcommon.Address, client *ethclient.Client) (*ethRollupWatcher, error) {
	arbitrumRollupContract, err := rollup.NewArbRollup(rollupAddress, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to arbRollup")
	}

	globalInboxAddress, err := arbitrumRollupContract.GlobalInbox(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to get GlobalInbox address")
	}
	globalInboxContract, err := rollup.NewIGlobalInbox(globalInboxAddress, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to GlobalInbox")
	}

	return &ethRollupWatcher{
		ArbRollup:     arbitrumRollupContract,
		GlobalInbox:   globalInboxContract,
		rollupAddress: rollupAddress,
		inboxAddress:  globalInboxAddress,
		client:        client,
	}, nil
}

func (vm *ethRollupWatcher) GetEvents(ctx context.Context, blockId *common.BlockId) ([]arbbridge.Event, error) {
	bh := blockId.HeaderHash.ToEthHash()
	addressIndex := ethcommon.Hash{}
	copy(addressIndex[:], ethcommon.LeftPadBytes(vm.rollupAddress.Bytes(), 32))
	inboxLogs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &bh,
		Addresses: []ethcommon.Address{vm.inboxAddress},
		Topics: [][]ethcommon.Hash{
			{
				transactionMessageDeliveredID,
				ethDepositMessageDeliveredID,
				depositERC20MessageDeliveredID,
				depositERC721MessageDeliveredID,
			}, {
				addressIndex,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	rollupLogs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &bh,
		Addresses: []ethcommon.Address{vm.rollupAddress},
		Topics: [][]ethcommon.Hash{
			{
				rollupStakeCreatedID,
				rollupChallengeStartedID,
				rollupChallengeCompletedID,
				rollupRefundedID,
				rollupPrunedID,
				rollupStakeMovedID,
				rollupAssertedID,
				rollupConfirmedID,
				confirmedAssertionID,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(inboxLogs)+len(inboxLogs))

	for _, evmLog := range inboxLogs {
		event, err := vm.processEvents(getLogChainInfo(evmLog), evmLog)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	for _, evmLog := range rollupLogs {
		event, err := vm.processEvents(getLogChainInfo(evmLog), evmLog)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (vm *ethRollupWatcher) ProcessMessageDeliveredEvents(chainInfo arbbridge.ChainInfo, ethLog types.Log) (arbbridge.Event, error) {
	if ethLog.Topics[0] == transactionMessageDeliveredID {
		val, err := vm.GlobalInbox.ParseTransactionMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredTransaction{
			Transaction: message.Transaction{
				Chain:       common.NewAddressFromEth(vm.rollupAddress),
				To:          common.NewAddressFromEth(val.To),
				From:        common.NewAddressFromEth(val.From),
				SequenceNum: val.SeqNumber,
				Value:       val.Value,
				Data:        val.Data,
			},
			BlockNum: common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil

	} else if ethLog.Topics[0] == ethDepositMessageDeliveredID {
		val, err := vm.GlobalInbox.ParseEthDepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredEth{
			Eth: message.Eth{
				To:    common.NewAddressFromEth(val.To),
				From:  common.NewAddressFromEth(val.From),
				Value: val.Value,
			},
			BlockNum:   common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			MessageNum: val.MessageNum,
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil

	} else if ethLog.Topics[0] == depositERC20MessageDeliveredID {
		val, err := vm.GlobalInbox.ParseERC20DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredERC20{
			ERC20: message.ERC20{
				To:           common.NewAddressFromEth(val.To),
				From:         common.NewAddressFromEth(val.From),
				TokenAddress: common.NewAddressFromEth(val.Erc20),
				Value:        val.Value,
			},
			BlockNum:   common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			MessageNum: val.MessageNum,
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil

	} else if ethLog.Topics[0] == depositERC721MessageDeliveredID {
		val, err := vm.GlobalInbox.ParseERC721DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredERC721{
			ERC721: message.ERC721{
				To:           common.NewAddressFromEth(val.To),
				From:         common.NewAddressFromEth(val.From),
				TokenAddress: common.NewAddressFromEth(val.Erc721),
				Id:           val.Id,
			},
			BlockNum:   common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			MessageNum: val.MessageNum,
		}

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil
	} else if ethLog.Topics[0] == contractTransactionMessageDeliveredID {
		val, err := vm.GlobalInbox.ParseContractTransactionMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		msg := message.DeliveredContractTransaction{
			ContractTransaction: message.ContractTransaction{
				To:    common.NewAddressFromEth(val.To),
				From:  common.NewAddressFromEth(val.From),
				Value: val.Value,
				Data:  val.Data,
			},
			BlockNum:   common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			MessageNum: val.MessageNum,
		}
		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			Message:   msg,
		}, nil
	} else {
		return nil, errors2.New("unknown arbitrum event type")
	}
}

func (vm *ethRollupWatcher) processEvents(chainInfo arbbridge.ChainInfo, ethLog types.Log) (arbbridge.Event, error) {
	if ethLog.Topics[0] == rollupStakeCreatedID {
		eventVal, err := vm.ArbRollup.ParseRollupStakeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeCreatedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			NodeHash:  eventVal.NodeHash,
		}, nil
	} else if ethLog.Topics[0] == rollupChallengeStartedID {
		eventVal, err := vm.ArbRollup.ParseRollupChallengeStarted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeStartedEvent{
			ChainInfo:         chainInfo,
			Asserter:          common.NewAddressFromEth(eventVal.Asserter),
			Challenger:        common.NewAddressFromEth(eventVal.Challenger),
			ChallengeType:     valprotocol.ChildType(eventVal.ChallengeType.Uint64()),
			ChallengeContract: common.NewAddressFromEth(eventVal.ChallengeContract),
		}, nil
	} else if ethLog.Topics[0] == rollupChallengeCompletedID {
		eventVal, err := vm.ArbRollup.ParseRollupChallengeCompleted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeCompletedEvent{
			ChainInfo:         chainInfo,
			Winner:            common.NewAddressFromEth(eventVal.Winner),
			Loser:             common.NewAddressFromEth(eventVal.Loser),
			ChallengeContract: common.NewAddressFromEth(eventVal.ChallengeContract),
		}, nil
	} else if ethLog.Topics[0] == rollupRefundedID {
		eventVal, err := vm.ArbRollup.ParseRollupStakeRefunded(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeRefundedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
		}, nil
	} else if ethLog.Topics[0] == rollupPrunedID {
		eventVal, err := vm.ArbRollup.ParseRollupPruned(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.PrunedEvent{
			ChainInfo: chainInfo,
			Leaf:      eventVal.Leaf,
		}, nil
	} else if ethLog.Topics[0] == rollupStakeMovedID {
		eventVal, err := vm.ArbRollup.ParseRollupStakeMoved(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeMovedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			Location:  eventVal.ToNodeHash,
		}, nil
	} else if ethLog.Topics[0] == rollupAssertedID {
		eventVal, err := vm.ArbRollup.ParseRollupAsserted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.AssertedEvent{
			ChainInfo:    chainInfo,
			PrevLeafHash: eventVal.Fields[0],
			Params: &valprotocol.AssertionParams{
				NumSteps: eventVal.NumSteps,
				TimeBounds: &protocol.TimeBoundsBlocks{
					common.NewTimeBlocks(eventVal.TimeBoundsBlocks[0]),
					common.NewTimeBlocks(eventVal.TimeBoundsBlocks[1]),
				},
				ImportedMessageCount: eventVal.ImportedMessageCount,
			},
			Claim: &valprotocol.AssertionClaim{
				AfterInboxTop:         eventVal.Fields[2],
				ImportedMessagesSlice: eventVal.Fields[3],
				AssertionStub: &valprotocol.ExecutionAssertionStub{
					AfterHash:        eventVal.Fields[4],
					DidInboxInsn:     eventVal.DidInboxInsn,
					NumGas:           eventVal.NumArbGas,
					FirstMessageHash: [32]byte{},
					LastMessageHash:  eventVal.Fields[5],
					FirstLogHash:     [32]byte{},
					LastLogHash:      eventVal.Fields[6],
				},
			},
			MaxInboxTop:   eventVal.Fields[1],
			MaxInboxCount: eventVal.InboxCount,
		}, nil
	} else if ethLog.Topics[0] == rollupConfirmedID {
		eventVal, err := vm.ArbRollup.ParseRollupConfirmed(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedEvent{
			ChainInfo: chainInfo,
			NodeHash:  eventVal.NodeHash,
		}, nil
	} else if ethLog.Topics[0] == confirmedAssertionID {
		eventVal, err := vm.ArbRollup.ParseConfirmedAssertion(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedAssertionEvent{
			LogsAccHash: hashSliceToHashes(eventVal.LogsAccHash),
		}, nil
	}
	return vm.ProcessMessageDeliveredEvents(chainInfo, ethLog)
}

func (vm *ethRollupWatcher) GetParams(ctx context.Context) (valprotocol.ChainParams, error) {
	rawParams, err := vm.ArbRollup.VmParams(nil)
	if err != nil {
		return valprotocol.ChainParams{}, err
	}
	stakeRequired, err := vm.ArbRollup.GetStakeRequired(nil)
	if err != nil {
		return valprotocol.ChainParams{}, err
	}
	return valprotocol.ChainParams{
		StakeRequirement:        stakeRequired,
		GracePeriod:             common.TimeTicks{rawParams.GracePeriodTicks},
		MaxExecutionSteps:       rawParams.MaxExecutionSteps,
		MaxTimeBoundsWidth:      rawParams.MaxTimeBoundsWidth,
		ArbGasSpeedLimitPerTick: rawParams.ArbGasSpeedLimitPerTick.Uint64(),
	}, nil
}

func (vm *ethRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	addr, err := vm.ArbRollup.GlobalInbox(nil)
	return common.NewAddressFromEth(addr), err
}

func (con *ethRollupWatcher) GetCreationInfo(ctx context.Context) (*common.BlockId, common.Hash, error) {
	addressIndex := ethcommon.Hash{}
	copy(addressIndex[:], ethcommon.LeftPadBytes(con.rollupAddress.Bytes(), 32))
	logs, err := con.client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []ethcommon.Address{con.rollupAddress},
		Topics:    [][]ethcommon.Hash{{rollupCreatedID}},
	})
	if err != nil {
		return nil, common.Hash{}, err
	}
	if len(logs) != 1 {
		return nil, common.Hash{}, errors.New("more than one chain created with same address")
	}
	ev, err := con.ArbRollup.ParseRollupCreated(logs[0])
	if err != nil {
		return nil, common.Hash{}, err
	}

	return getLogBlockID(logs[0]), ev.InitVMHash, nil
}

func (con *ethRollupWatcher) GetVersion(ctx context.Context) (string, error) {
	return con.ArbRollup.VERSION(&bind.CallOpts{Context: ctx})
}
