package ethbridge

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
)

type Bridge struct {
	*BridgeWatcher
	auth *TransactAuth
}

func NewBridge(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*Bridge, error) {
	watcher, err := NewBridgeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Bridge{
		BridgeWatcher: watcher,
		auth:          auth,
	}, nil
}

func (b *Bridge) SendL2MessageFromOrigin() {

}
