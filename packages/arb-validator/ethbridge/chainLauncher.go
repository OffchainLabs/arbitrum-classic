package ethbridge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/chainlauncher"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChainLauncher struct {
	contract *chainlauncher.ChainLauncher
	client   *ethclient.Client
}

func NewChainLauncher(address common.Address, client *ethclient.Client) (*ChainLauncher, error) {
	vmCreatorContract, err := chainlauncher.NewChainLauncher(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ArbLauncher")
	}
	return &ChainLauncher{vmCreatorContract, client}, nil
}

func (con *ChainLauncher) ParseChainCreated(log *types.Log) (common.Address, error) {
	event, err := con.contract.ParseChainCreated(*log)
	if err != nil {
		return common.Address{}, err
	}
	return event.VmAddress, nil
}

func (con *ChainLauncher) LaunchChain(
	auth *bind.TransactOpts,
	config *valmessage.VMConfiguration,
	vmState [32]byte,
) (common.Address, error) {
	var owner common.Address
	copy(owner[:], config.Owner.Value)
	var escrowCurrency common.Address
	copy(escrowCurrency[:], config.EscrowCurrency.Value)
	tx, err := con.contract.LaunchChain(
		auth,
		vmState,
		uint32(config.GracePeriod),
		config.MaxExecutionStepCount,
		value.NewBigIntFromBuf(config.EscrowRequired),
		owner,
	)
	if err != nil {
		return common.Address{}, err
	}
	receipt, err := waitForReceipt(auth.Context, con.client, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	if len(receipt.Logs) != 1 {
		return common.Address{}, errors2.New("Wrong receipt count")
	}
	event, err := con.contract.ParseChainCreated(*receipt.Logs[0])
	if err != nil {
		return common.Address{}, err
	}
	return event.VmAddress, nil
}
