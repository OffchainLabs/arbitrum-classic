package ethbridge

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/arbfactory"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ArbFactory struct {
	contract *arbfactory.ArbFactory
	client   *ethclient.Client
}

func NewArbFactory(address common.Address, client *ethclient.Client) (*ArbFactory, error) {
	vmCreatorContract, err := arbfactory.NewArbFactory(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ArbFactory")
	}
	return &ArbFactory{vmCreatorContract, client}, nil
}

func (con *ArbFactory) CreateRollup(
	auth *bind.TransactOpts,
	vmState [32]byte,
	params structures.ChainParams,
	owner common.Address,
) (common.Address, error) {
	tx, err := con.contract.CreateRollup(
		auth,
		vmState,
		params.GracePeriod.Val,
		new(big.Int).SetUint64(params.ArbGasSpeedLimitPerTick),
		params.MaxExecutionSteps,
		params.StakeRequirement,
		owner,
	)
	if err != nil {
		return common.Address{}, errors2.Wrap(err, "Failed to call to ChainFactory.CreateChain")
	}
	receipt, err := waitForReceipt(auth.Context, con.client, auth.From, tx, "CreateChain")
	if err != nil {
		return common.Address{}, err
	}
	if len(receipt.Logs) != 1 {
		return common.Address{}, errors2.New("Wrong receipt count")
	}
	event, err := con.contract.ParseRollupCreated(*receipt.Logs[0])
	if err != nil {
		return common.Address{}, err
	}
	return event.VmAddress, nil
}
