package mockbridge

import (
	"context"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type ArbFactory struct {
	//contract *arbfactory.ArbFactory
	client arbbridge.ArbClient
}

func NewArbFactory(address common.Address, client arbbridge.ArbClient) (*ArbFactory, error) {
	return &ArbFactory{client}, nil
}

func (con *ArbFactory) CreateRollup(
	ctx context.Context,
	vmState common.Hash,
	params valprotocol.ChainParams,
	owner common.Address,
) (common.Address, error) {
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
