package ethconnection

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection/channelcreator"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChannelCreator struct {
	contract *channelcreator.ChannelCreator
}

func NewVMCreator(address common.Address, client *ethclient.Client) (*ChannelCreator, error) {
	vmCreatorContract, err := channelcreator.NewChannelCreator(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChannelCreator")
	}
	return &ChannelCreator{vmCreatorContract}, nil
}

func (con *ChannelCreator) ParseVMCreated(log *types.Log) (common.Address, [32]byte, *valmessage.VMConfiguration, error) {
	event, err := con.contract.ParseChannelCreated(*log)
	if err != nil {
		return common.Address{}, [32]byte{}, nil, err
	}
	return event.VmAddress, event.VmState, valmessage.NewVMConfiguration(
		uint64(event.GracePeriod),
		event.EscrowRequired,
		common.Address{},
		event.Validators,
		event.MaxExecutionSteps,
		event.Owner,
	), nil
}

func (con *ChannelCreator) LaunchVM(
	auth *bind.TransactOpts,
	config *valmessage.VMConfiguration,
	vmState [32]byte,
) (*types.Transaction, error) {
	var owner common.Address
	copy(owner[:], config.Owner.Value)
	var escrowCurrency common.Address
	copy(escrowCurrency[:], config.EscrowCurrency.Value)
	validatorKeys := make([]common.Address, 0, len(config.AssertKeys))
	for _, key := range config.AssertKeys {
		validatorKeys = append(validatorKeys, protocol.NewAddressFromBuf(key))
	}
	return con.contract.LaunchVM(
		auth,
		vmState,
		uint32(config.GracePeriod),
		config.MaxExecutionStepCount,
		value.NewBigIntFromBuf(config.EscrowRequired),
		owner,
		validatorKeys,
	)
}
