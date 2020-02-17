package gobridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"math/big"
)

type ArbFactory struct {
	rollupAddress common.Address
	client        *GoArbClient
}

func newArbFactory(address common.Address, client *GoArbClient) (*ArbFactory, error) {
	return &ArbFactory{client.GoEthClient.arbFactory, client}, nil
}

func (con *ArbFactory) CreateRollup(
	ctx context.Context,
	vmState common.Hash,
	params valprotocol.ChainParams,
	owner common.Address,
) (common.Address, error) {
	events := make(map[*common.BlockId][]arbbridge.Event)
	addr := con.client.GoEthClient.getNextAddress()
	vmProto := hashing.SoliditySHA3(
		hashing.Bytes32(vmState),
		hashing.Bytes32(value.NewEmptyTuple().Hash()),
		hashing.Uint256(big.NewInt(0)),
	)
	innerHash := hashing.SoliditySHA3(
		hashing.Bytes32(vmProto),
		hashing.Uint256(big.NewInt(0)),
		hashing.Uint256(big.NewInt(0)),
		hashing.Uint256(big.NewInt(0)),
	)
	initialNode := hashing.SoliditySHA3(
		hashing.Uint256(big.NewInt(0)),
		hashing.Bytes32(innerHash),
	)

	con.client.GoEthClient.rollups[addr] = &rollupData{
		initVMHash:              vmState,
		VMstate:                 machine.Extensive,
		state:                   Uninitialized,
		gracePeriod:             params.GracePeriod,
		maxSteps:                params.MaxExecutionSteps,
		maxTimeBoundsWidth:      params.MaxTimeBoundsWidth,
		arbGasSpeedLimitPerTick: params.ArbGasSpeedLimitPerTick,
		escrowRequired:          params.StakeRequirement,
		owner:                   owner,
		events:                  events,
		creation:                con.client.GoEthClient.getCurrentBlock(),
		stakers:                 make(map[common.Address]*staker),
		leaves:                  make(map[common.Hash]bool),
		lastConfirmed:           initialNode,
		contractAddress:         addr,
	}
	con.client.GoEthClient.rollups[addr].leaves[initialNode] = true

	//event, err := con.contract.ParseRollupCreated(*receipt.Logs[0])
	//if err != nil {
	//	return common.Address{}, err
	//}
	return addr, nil
}

type arbFactoryWatcher struct {
	rollupAddress common.Address
	client        *GoArbClient
}

func newArbFactoryWatcher(address common.Address, client *GoArbClient) (*arbFactoryWatcher, error) {
	//vmCreatorContract, err := arbfactory.newArbFactory(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to arbFactory")
	//}
	return &arbFactoryWatcher{address, client}, nil
}

func (con *arbFactoryWatcher) GlobalInboxAddress() (common.Address, error) {
	return con.rollupAddress, nil
}

func (con *arbFactoryWatcher) ChallengeFactoryAddress() (common.Address, error) {
	return con.rollupAddress, nil
}
