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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/globalinbox"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

var transactionID ethcommon.Hash
var transactionBatchID ethcommon.Hash
var ethDepositID ethcommon.Hash
var depositERC20ID ethcommon.Hash
var depositERC721ID ethcommon.Hash
var contractTxID ethcommon.Hash
var transactionBatchTxCallABI abi.Method

func init() {
	inbox, err := abi.JSON(strings.NewReader(globalinbox.GlobalInboxABI))
	if err != nil {
		panic(err)
	}
	transactionID = inbox.Events["TransactionMessageDelivered"].ID
	transactionBatchID = inbox.Events["TransactionMessageBatchDelivered"].ID
	ethDepositID = inbox.Events["EthDepositMessageDelivered"].ID
	depositERC20ID = inbox.Events["ERC20DepositMessageDelivered"].ID
	depositERC721ID = inbox.Events["ERC721DepositMessageDelivered"].ID
	contractTxID = inbox.Events["ContractTransactionMessageDelivered"].ID
	transactionBatchTxCallABI = inbox.Methods["deliverTransactionBatch"]
}

type globalInboxWatcher struct {
	GlobalInbox *globalinbox.GlobalInbox

	rollupAddress ethcommon.Address
	inboxAddress  ethcommon.Address
	client        ethutils.EthClient
}

func newGlobalInboxWatcher(
	globalInboxAddress ethcommon.Address,
	rollupAddress ethcommon.Address,
	client ethutils.EthClient,
) (*globalInboxWatcher, error) {
	globalInboxContract, err := globalinbox.NewGlobalInbox(
		globalInboxAddress,
		client,
	)
	if err != nil {
		return nil, errors2.Wrap(err, "failed to connect to inbox")
	}

	return &globalInboxWatcher{
		GlobalInbox:   globalInboxContract,
		rollupAddress: rollupAddress,
		inboxAddress:  globalInboxAddress,
		client:        client,
	}, nil
}

func (vm *globalInboxWatcher) getLogs(
	ctx context.Context,
	fromBlock, toBlock *big.Int,
	blockHash *ethcommon.Hash,
) ([]types.Log, error) {
	addressIndex := ethcommon.Hash{}
	copy(
		addressIndex[:],
		ethcommon.LeftPadBytes(vm.rollupAddress.Bytes(), 32),
	)
	return vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		BlockHash: blockHash,
		Addresses: []ethcommon.Address{vm.inboxAddress},
		Topics: [][]ethcommon.Hash{
			{
				transactionID,
				transactionBatchID,
				ethDepositID,
				depositERC20ID,
				depositERC721ID,
			}, {
				addressIndex,
			},
		},
	})
}

func (vm *globalInboxWatcher) GetDeliveredEvents(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]arbbridge.MessageDeliveredEvent, error) {
	inboxLogs, err := vm.getLogs(ctx, fromBlock, toBlock, nil)
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.MessageDeliveredEvent, 0, len(inboxLogs))
	for _, evmLog := range inboxLogs {
		blockHeader, err := vm.client.HeaderByHash(ctx, evmLog.BlockHash)
		if err != nil {
			return nil, err
		}
		timestamp := new(big.Int).SetUint64(blockHeader.Time)
		msg, err := vm.processLog(ctx, evmLog, timestamp)
		if err != nil {
			return nil, err
		}
		events = append(events, arbbridge.MessageDeliveredEvent{
			ChainInfo: getLogChainInfo(evmLog),
			Message:   msg,
		})
	}
	return events, nil
}

func (vm *globalInboxWatcher) GetEvents(
	ctx context.Context,
	blockId *common.BlockId,
	timestamp *big.Int,
) ([]arbbridge.Event, error) {
	bh := blockId.HeaderHash.ToEthHash()
	inboxLogs, err := vm.getLogs(ctx, nil, nil, &bh)
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(inboxLogs))
	for _, evmLog := range inboxLogs {
		msg, err := vm.processLog(ctx, evmLog, timestamp)
		if err != nil {
			return nil, err
		}
		events = append(events, arbbridge.MessageDeliveredEvent{
			ChainInfo: getLogChainInfo(evmLog),
			Message:   msg,
		})
	}
	return events, nil
}

func (vm *globalInboxWatcher) GetAllReceived(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]message.Received, error) {
	deliveredEvents, err := vm.GetDeliveredEvents(ctx, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}
	events := make([]message.Received, 0, len(deliveredEvents))
	for _, ev := range deliveredEvents {
		events = append(events, ev.Message)
	}
	return events, nil
}

func (vm *globalInboxWatcher) processLog(
	ctx context.Context,
	evmLog types.Log,
	timestamp *big.Int,
) (message.Received, error) {
	msg, err := vm.processMessageDeliveredEvents(ctx, evmLog)
	if err != nil {
		return message.Received{}, err
	}

	return message.Received{
		Message: msg,
		ChainTime: message.ChainTime{
			BlockNum: common.NewTimeBlocks(
				new(big.Int).SetUint64(evmLog.BlockNumber),
			),
			Timestamp: timestamp,
		},
	}, nil
}

func (vm *globalInboxWatcher) processMessageDeliveredEvents(
	ctx context.Context,
	ethLog types.Log,
) (message.Message, error) {
	switch ethLog.Topics[0] {
	case transactionID:
		val, err := vm.GlobalInbox.ParseTransactionMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		return message.Transaction{
			Chain:       common.NewAddressFromEth(vm.rollupAddress),
			To:          common.NewAddressFromEth(val.To),
			From:        common.NewAddressFromEth(val.From),
			SequenceNum: val.SeqNumber,
			Value:       val.Value,
			Data:        val.Data,
		}, nil

	case transactionBatchID:
		tx, _, err := vm.client.TransactionByHash(ctx, ethLog.TxHash)
		if err != nil {
			return nil, err
		}

		type TransactionBatchTxCallArgs struct {
			Chain        ethcommon.Address
			Transactions []byte
		}

		var args TransactionBatchTxCallArgs
		err = transactionBatchTxCallABI.Inputs.Unpack(&args, tx.Data()[4:])
		if err != nil {
			return nil, err
		}

		return message.TransactionBatch{
			Chain:  common.NewAddressFromEth(vm.rollupAddress),
			TxData: args.Transactions,
		}, nil
	case ethDepositID:
		val, err := vm.GlobalInbox.ParseEthDepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		return message.Eth{
			To:    common.NewAddressFromEth(val.To),
			From:  common.NewAddressFromEth(val.From),
			Value: val.Value,
		}, nil

	case depositERC20ID:
		val, err := vm.GlobalInbox.ParseERC20DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		return message.ERC20{
			To:           common.NewAddressFromEth(val.To),
			From:         common.NewAddressFromEth(val.From),
			TokenAddress: common.NewAddressFromEth(val.Erc20),
			Value:        val.Value,
		}, nil

	case depositERC721ID:
		val, err := vm.GlobalInbox.ParseERC721DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		return message.ERC721{
			To:           common.NewAddressFromEth(val.To),
			From:         common.NewAddressFromEth(val.From),
			TokenAddress: common.NewAddressFromEth(val.Erc721),
			Id:           val.Id,
		}, nil
	case contractTxID:
		val, err := vm.GlobalInbox.ParseContractTransactionMessageDelivered(
			ethLog,
		)
		if err != nil {
			return nil, err
		}

		return message.ContractTransaction{
			To:    common.NewAddressFromEth(val.To),
			From:  common.NewAddressFromEth(val.From),
			Value: val.Value,
			Data:  val.Data,
		}, nil
	default:
		return nil, errors2.New("unknown arbitrum event type")
	}
}

func (con *globalInboxWatcher) GetERC20Balance(
	ctx context.Context,
	user common.Address,
	tokenContract common.Address,
) (*big.Int, error) {
	return con.GlobalInbox.GetERC20Balance(
		&bind.CallOpts{Context: ctx},
		tokenContract.ToEthAddress(),
		user.ToEthAddress(),
	)
}

func (con *globalInboxWatcher) GetEthBalance(
	ctx context.Context,
	user common.Address,
) (*big.Int, error) {
	return con.GlobalInbox.GetEthBalance(
		&bind.CallOpts{Context: ctx},
		user.ToEthAddress(),
	)
}
