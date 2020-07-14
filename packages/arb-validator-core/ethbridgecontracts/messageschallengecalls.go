package ethbridgecontracts

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
)

func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofEthMessageCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _lowerHashA [32]byte, _preImageBHash [32]byte, _preImageBSize *big.Int, _to common.Address, _from common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) error {
	return callCheckMessages(ctx, client, from, contractAddress, "oneStepProofEthMessage", _lowerHashA, _preImageBHash, _preImageBSize, _to, _from, _value, _blockNumber, _timestamp, _messageNum)
}

func (_MessagesChallenge *MessagesChallengeTransactor) BisectCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _chainHashes [][32]byte, _segmentHashes [][32]byte, _chainLength *big.Int) error {
	return callCheckMessages(ctx, client, from, contractAddress, "bisect", _chainHashes, _segmentHashes, _chainLength)
}

func callCheckMessages(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(MessagesChallengeABI)))
	if err != nil {
		return err
	}

	return ethutils.CallCheck(ctx, client, from, contractAddress, contractABI, method, params...)
}
