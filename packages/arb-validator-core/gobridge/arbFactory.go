package gobridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type arbFactory struct {
	rollupContractAddress common.Address
	client                *GoArbAuthClient
}

func deployRollupFactory(m *goEthdata) {
	m.rollups = make(map[common.Address]*arbRollup)
	m.arbFactoryContract = &arbFactory{
		rollupContractAddress: m.getNextAddress(),
		client:                nil,
	}
}

func newArbFactory(address common.Address, client *GoArbAuthClient) (*arbFactory, error) {
	client.arbFactoryContract.client = client
	return client.arbFactoryContract, nil
}

func (con *arbFactory) CreateRollup(
	ctx context.Context,
	vmState common.Hash,
	params valprotocol.ChainParams,
	owner common.Address,
) (common.Address, error) {
	con.client.goEthMutex.Lock()
	defer con.client.goEthMutex.Unlock()
	addr := con.client.getNextAddress()

	newGlobalInbox(addr, con.client)
	newRollup(con, addr, vmState, params, owner)

	return addr, nil
}

type arbFactoryWatcher struct {
	rollupAddress common.Address
	client        *goEthdata
}

func newArbFactoryWatcher(address common.Address, client *goEthdata) (*arbFactoryWatcher, error) {
	return &arbFactoryWatcher{address, client}, nil
}

func (con *arbFactoryWatcher) GlobalInboxAddress() (common.Address, error) {
	con.client.goEthMutex.Lock()
	defer con.client.goEthMutex.Unlock()
	return con.rollupAddress, nil
}

func (con *arbFactoryWatcher) ChallengeFactoryAddress() (common.Address, error) {
	con.client.goEthMutex.Lock()
	defer con.client.goEthMutex.Unlock()
	return con.rollupAddress, nil
}
