package messageschallenge

import (
	"bytes"
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
)

func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofEthMessageCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) error {
	return CallCheck(ctx, client, from, contractAddress, "oneStepProofEthMessage", _lowerHashA, _lowerHashB, _to, _from, _value, _blockNumber, _timestamp, _messageNum)
}

func (_MessagesChallenge *MessagesChallengeTransactor) BisectCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _chainHashes [][32]byte, _segmentHashes [][32]byte, _chainLength *big.Int) error {
	return CallCheck(ctx, client, from, contractAddress, "bisect", _chainHashes, _segmentHashes, _chainLength)
}

func CallCheck(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(MessagesChallengeABI)))
	if err != nil {
		return err
	}

	return ethutils.CallCheck(ctx, client, from, contractAddress, contractABI, method, params...)
}
