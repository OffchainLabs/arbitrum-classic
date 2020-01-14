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
	client   arbbridge.ArbClient
}

func newArbFactory(address common.Address, client arbbridge.ArbClient) (*ArbFactory, error) {
	return &ArbFactory{address, client}, nil
}

func (con *ArbFactory) CreateRollup(
	ctx context.Context,
	vmState common.Hash,
	params structures.ChainParams,
	owner common.Address,
) (common.Address, error) {
	//con.contract =
	//tx, err := con.contract.CreateRollup(
	//	auth,
	//	vmState,
	//	params.GracePeriod.Val,
	//	new(big.Int).SetUint64(params.ArbGasSpeedLimitPerTick),
	//	params.MaxExecutionSteps,
	//	params.StakeRequirement,
	//	owner,
	//)
	//if err != nil {
	//	return common.Address{}, errors2.Wrap(err, "Failed to call to ChainFactory.CreateChain")
	//}
	//receipt, err := waitForReceiptWithResults(auth.Context, con.client, auth.From, tx, "CreateChain")
	//if err != nil {
	//	return common.Address{}, err
	//}
	//if len(receipt.Logs) != 1 {
	//	return common.Address{}, errors2.New("Wrong receipt count")
	//}
	//event, err := con.contract.ParseRollupCreated(*receipt.Logs[0])
	//if err != nil {
	//	return common.Address{}, err
	//}
	return common.Address{}, nil
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
