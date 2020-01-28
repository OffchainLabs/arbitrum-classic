package executionchallenge

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/common"
)

func (_BisectionChallenge *BisectionChallengeTransactor) ChooseSegmentCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) error {
	return callCheck(ctx, client, from, contractAddress, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) BisectAssertionCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _beforeInbox [32]byte, _timeBoundsBlocks [2]*big.Int, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) error {
	return callCheck(ctx, client, from, contractAddress, "bisectAssertion", _beforeInbox, _timeBoundsBlocks, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProofCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBoundsBlocks [2]*big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) error {
	return callCheck(ctx, client, from, contractAddress, "oneStepProof", _beforeHash, _beforeInbox, _timeBoundsBlocks, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

func (_Challenge *ChallengeTransactor) TimeoutChallengeCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address) error {
	return callCheck(ctx, client, from, contractAddress, "timeoutChallenge")
}

func callCheck(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(ExecutionChallengeABI)))
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
