/*
* Copyright 2021, Offchain Labs, Inc.
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

package arbos

import (
	"bytes"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	RetryCanceledEvent abi.Event
	RetryRedeemedEvent abi.Event

	createRetryableTicketABI abi.Method
	redeemABI                abi.Method
)

func init() {
	parsedABI, err := abi.JSON(strings.NewReader(arboscontracts.ArbRetryableTxABI))
	if err != nil {
		panic(err)
	}

	creatorABI, err := abi.JSON(strings.NewReader(arboscontracts.RetryableTicketCreatorABI))
	if err != nil {
		panic(err)
	}

	RetryCanceledEvent = parsedABI.Events["Canceled"]
	RetryRedeemedEvent = parsedABI.Events["Redeemed"]
	redeemABI = parsedABI.Methods["redeem"]
	createRetryableTicketABI = creatorABI.Methods["createRetryableTicket"]
}

func CreateRetryableTicketData(msg message.RetryableTx) []byte {
	txData, err := createRetryableTicketABI.Inputs.Pack(
		msg.Destination.ToEthAddress(),
		msg.Value,
		msg.MaxSubmissionCost,
		msg.CreditBack.ToEthAddress(),
		msg.Beneficiary.ToEthAddress(),
		msg.MaxGas,
		msg.GasPriceBid,
		msg.Data,
	)
	if err != nil {
		panic(err)
	}
	return append(createRetryableTicketABI.ID, txData...)
}

func RedeemData(txId common.Hash) []byte {
	return append(redeemABI.ID, txId[:]...)
}

func ParseCreateRetryableTicketTx(tx *types.Transaction) (*message.RetryableTx, error) {
	if !bytes.Equal(tx.Data()[:4], createRetryableTicketABI.ID) {
		return nil, errors.New("bad func id")
	}
	args, err := createRetryableTicketABI.Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		return nil, err
	}
	if len(args) != 8 {
		return nil, errors.New("unexpected arg count")
	}

	dest, ok := args[0].(ethcommon.Address)
	if !ok {
		return nil, errors.New("bad dest")
	}
	value, ok := args[1].(*big.Int)
	if !ok {
		return nil, errors.New("bad value")
	}
	cost, ok := args[2].(*big.Int)
	if !ok {
		return nil, errors.New("bad cost")
	}
	creditBack, ok := args[3].(ethcommon.Address)
	if !ok {
		return nil, errors.New("bad dest")
	}
	beneficiary, ok := args[4].(ethcommon.Address)
	if !ok {
		return nil, errors.New("bad beneficiary")
	}
	maxGas, ok := args[5].(*big.Int)
	if !ok {
		return nil, errors.New("bad max gas")
	}
	gasPrice, ok := args[6].(*big.Int)
	if !ok {
		return nil, errors.New("bad gas price")
	}
	calldata, ok := args[7].([]byte)
	if !ok {
		return nil, errors.New("bad data")
	}

	return &message.RetryableTx{
		Destination:       common.NewAddressFromEth(dest),
		Value:             value,
		Deposit:           tx.Value(),
		MaxSubmissionCost: cost,
		CreditBack:        common.NewAddressFromEth(creditBack),
		Beneficiary:       common.NewAddressFromEth(beneficiary),
		MaxGas:            maxGas,
		GasPriceBid:       gasPrice,
		Data:              calldata,
	}, nil
}
