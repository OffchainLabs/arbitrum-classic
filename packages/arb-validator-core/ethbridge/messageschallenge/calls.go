package messageschallenge

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
)

func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofEthMessageCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) error {
	return CallCheck(ctx, client, from, contractAddress, "oneStepProofEthMessage", _lowerHashA, _lowerHashB, _to, _from, _value, _blockNumber, _timestamp, _messageNum)
}

func CallCheck(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(MessagesChallengeABI)))
	if err != nil {
		return err
	}

	// Pack the input, call and unpack the results
	input, err := contractABI.Pack(method, params...)
	if err != nil {
		return err
	}
	var (
		msg    = ethereum.CallMsg{From: from, To: &contractAddress, Data: input}
		output []byte
	)

	output, err = client.PendingCallContract(ctx, msg)
	if err != nil {
		return err
	}

	if len(output) < 69 {
		return fmt.Errorf("%v had short output %v, %v", method, len(output), output)
	}
	length := new(big.Int).SetBytes(output[36:68])
	if uint64(len(output)) < 68+length.Uint64()+1 {
		return fmt.Errorf("%v had short output %v, %v", method, len(output), output)
	}
	return fmt.Errorf("%v returned val: %v", method, string(output[68:68+length.Uint64()]))
}
