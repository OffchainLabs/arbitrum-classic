package ethconnection

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection/arblauncher"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ArbLauncher struct {
	contract *arblauncher.ArbLauncher
	client   *ethclient.Client
}

func NewArbLauncher(address common.Address, client *ethclient.Client) (*ArbLauncher, error) {
	vmCreatorContract, err := arblauncher.NewArbLauncher(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ArbLauncher")
	}
	return &ArbLauncher{vmCreatorContract, client}, nil
}

func (con *ArbLauncher) ParseChannelCreated(log *types.Log) (common.Address, error) {
	event, err := con.contract.ParseChannelCreated(*log)
	if err != nil {
		return common.Address{}, err
	}
	return event.VmAddress, nil
}

func (con *ArbLauncher) LaunchChannel(
	auth *bind.TransactOpts,
	config *valmessage.VMConfiguration,
	vmState [32]byte,
) (common.Address, error) {
	var owner common.Address
	copy(owner[:], config.Owner.Value)
	var escrowCurrency common.Address
	copy(escrowCurrency[:], config.EscrowCurrency.Value)
	validatorKeys := make([]common.Address, 0, len(config.AssertKeys))
	for _, key := range config.AssertKeys {
		validatorKeys = append(validatorKeys, protocol.NewAddressFromBuf(key))
	}
	tx, err := con.contract.LaunchChannel(
		auth,
		vmState,
		uint32(config.GracePeriod),
		config.MaxExecutionStepCount,
		value.NewBigIntFromBuf(config.EscrowRequired),
		owner,
		validatorKeys,
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
	event, err := con.contract.ParseChannelCreated(*receipt.Logs[0])
	if err != nil {
		return common.Address{}, err
	}
	return event.VmAddress, nil
}

func (con *ArbLauncher) LaunchChain(
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
