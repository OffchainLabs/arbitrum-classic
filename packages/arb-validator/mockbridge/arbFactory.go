package mockbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/arbfactory"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ArbFactory struct {
	contract common.Address
	client   *MockArbClient
}

func newArbFactory(address common.Address, client *MockArbClient) (*ArbFactory, error) {
	return &ArbFactory{address, client}, nil
}

func (con *ArbFactory) CreateRollup(
	ctx context.Context,
	vmState common.Hash,
	params structures.ChainParams,
	owner common.Address,
) (common.Address, error) {
	events := make(map[*structures.BlockId][]arbbridge.Event)
	con.client.MockEthClient.rollups[owner] = &rollupData{Uninitialized,
		params.GracePeriod,
		params.MaxExecutionSteps,
		params.StakeRequirement,
		owner,
		events,
		con.client.MockEthClient.LatestBlock,
	}
	//event, err := con.contract.ParseRollupCreated(*receipt.Logs[0])
	//if err != nil {
	//	return common.Address{}, err
	//}
	return owner, nil
}

type arbFactoryWatcher struct {
	contract *arbfactory.ArbFactory
	client   arbbridge.ArbClient
}

func newArbFactoryWatcher(address common.Address, client arbbridge.ArbClient) (*arbFactoryWatcher, error) {
	//vmCreatorContract, err := arbfactory.newArbFactory(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to arbFactory")
	//}
	return &arbFactoryWatcher{nil, client}, nil
}

func (con *arbFactoryWatcher) GlobalPendingInboxAddress() (common.Address, error) {
	addr, err := con.contract.GlobalInboxAddress(nil)
	return common.NewAddressFromEth(addr), err
}

func (con *arbFactoryWatcher) ChallengeFactoryAddress() (common.Address, error) {
	addr, err := con.contract.ChallengeFactoryAddress(nil)
	return common.NewAddressFromEth(addr), err
}
