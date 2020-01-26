package rollup

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

// Solidity: function confirmValid(uint256 deadlineTicks, bytes _messages, bytes32 logsAcc, bytes32 vmProtoStateHash, address[] stakerAddresses, bytes32[] stakerProofs, uint256[] stakerProofOffsets) returns()
func (_ArbRollup *ArbRollupTransactor) ConfirmValidCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, deadlineTicks *big.Int, _messages []byte, logsAcc [32]byte, vmProtoStateHash [32]byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) error {
	return CallCheck(ctx, client, from, contractAddress, "confirmValid", deadlineTicks, _messages, logsAcc, vmProtoStateHash, stakerAddresses, stakerProofs, stakerProofOffsets)
}

// Solidity: function confirmInvalid(uint256 deadlineTicks, bytes32 challengeNodeData, uint256 branch, bytes32 vmProtoStateHash, address[] stakerAddresses, bytes32[] stakerProofs, uint256[] stakerProofOffsets) returns()
func (_ArbRollup *ArbRollupTransactor) ConfirmInvalidCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, deadlineTicks *big.Int, challengeNodeData [32]byte, branch *big.Int, vmProtoStateHash [32]byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) error {
	return CallCheck(ctx, client, from, contractAddress, "confirmInvalid", deadlineTicks, challengeNodeData, branch, vmProtoStateHash, stakerAddresses, stakerProofs, stakerProofOffsets)
}

func CallCheck(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(ArbRollupABI)))
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
