package web3

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type DevopsEth struct {
}

func (eth *DevopsEth) Syncing() bool {
	return false
}

type DevopsNet struct {
	chainId uint64
}

func (net *DevopsNet) PeerCount() hexutil.Uint {
	return 1
}

func registerDevopsStubs(s *rpc.Server) error {
	if err := s.RegisterName("net", &DevopsNet{}); err != nil {
		return err
	}

	return s.RegisterName("eth", &DevopsEth{})
}
