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

package arbostest

import (
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

func initMsg() message.Init {
	return message.Init{
		ChainParams: valprotocol.ChainParams{
			StakeRequirement:        big.NewInt(0),
			StakeToken:              common.Address{},
			GracePeriod:             common.TimeTicks{Val: big.NewInt(0)},
			MaxExecutionSteps:       0,
			ArbGasSpeedLimitPerTick: 0,
		},
		Owner:       common.Address{},
		ExtraConfig: []byte{},
	}
}

func runMessage(t *testing.T, mach machine.Machine, msg message.Message, sender common.Address) ([]*evm.TxResult, []message.OutMessage) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	assertion, steps := mach.ExecuteAssertion(
		1000000000,
		[]inbox.InboxMessage{message.NewInboxMessage(msg, sender, big.NewInt(0), chainTime)},
		0,
	)
	//data, err := value.TestVectorJSON(inbox, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(data))
	t.Log("Ran assertion for", steps, "steps and had", assertion.LogsCount, "logs and", assertion.OutMsgsCount, "messages")
	if mach.CurrentStatus() != machine.Extensive {
		t.Fatal("machine should still be working")
	}
	blockReason := mach.IsBlocked(false)
	if blockReason == nil {
		t.Fatal("machine not blocked")
	}

	if _, ok := blockReason.(machine.InboxBlocked); !ok {
		t.Fatal("Machine blocked for weird reason", blockReason)
	}
	results := make([]*evm.TxResult, 0)
	for _, avmLog := range assertion.ParseLogs() {
		result, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, result)
	}
	sends := make([]message.OutMessage, 0)
	for _, send := range assertion.ParseOutMessages() {
		msg, err := message.NewOutMessageFromValue(send)
		if err != nil {
			t.Fatal(err)
		}
		sends = append(sends, msg)
	}
	return results, sends
}

func runValidTransaction(t *testing.T, mach machine.Machine, msg message.AbstractL2Message, sender common.Address) (*evm.TxResult, error) {
	result, err := runTransaction(t, mach, msg, sender)
	if err != nil {
		return nil, err
	}
	if result.ResultCode != evm.ReturnCode {
		return nil, fmt.Errorf("transaction failed unexpectedly %v", result)
	}
	return result, nil
}

func runTransaction(t *testing.T, mach machine.Machine, msg message.AbstractL2Message, sender common.Address) (*evm.TxResult, error) {
	l2, err := message.NewL2Message(msg)
	if err != nil {
		return nil, err
	}
	results, sends := runMessage(t, mach, l2, sender)
	if len(results) != 1 {
		return nil, fmt.Errorf("unexpected log count %v", len(results))
	}
	if len(sends) != 0 {
		return nil, fmt.Errorf("unexpected send count %v", len(sends))
	}
	return results[0], nil
}

func withdrawEthTx(sequenceNum *big.Int, amount *big.Int, dest common.Address) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     amount,
		Data:        snapshot.WithdrawEthData(dest),
	}
}

func withdrawERC20Tx(sequenceNum *big.Int, amount *big.Int, dest common.Address) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.WithdrawERC20Data(dest, amount),
	}
}

func withdrawERC721Tx(sequenceNum *big.Int, id *big.Int, dest common.Address) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.WithdrawERC721Data(dest, id),
	}
}

func makeConstructorTx(code []byte, sequenceNum *big.Int) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        code,
	}
}

func deployContract(t *testing.T, mach machine.Machine, sender common.Address, code []byte, sequenceNum *big.Int) (common.Address, error) {
	constructorTx := makeConstructorTx(code, sequenceNum)
	constructorResult, err := runValidTransaction(t, mach, constructorTx, sender)
	if err != nil {
		return common.Address{}, err
	}
	return getConstructorResult(constructorResult)
}

func getConstructorResult(constructorResult *evm.TxResult) (common.Address, error) {
	if len(constructorResult.ReturnData) != 32 {
		return common.Address{}, errors.New("unexpected constructor result length")
	}
	var contractAddress common.Address
	copy(contractAddress[:], constructorResult.ReturnData[12:])
	return contractAddress, nil
}

func depositEth(t *testing.T, mach machine.Machine, dest common.Address, amount *big.Int) {
	msg := message.Eth{
		Dest:  dest,
		Value: amount,
	}

	depositResults, sendResults := runMessage(t, mach, msg, dest)
	if len(depositResults) != 0 {
		t.Fatal("deposit should not have had a result")
	}
	if len(sendResults) != 0 {
		t.Fatal("deposit should not trigger sends")
	}
}
