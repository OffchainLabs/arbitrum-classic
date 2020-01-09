package arbbridge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ArbFactory interface {
	CreateRollup(
		auth *bind.TransactOpts,
		vmState [32]byte,
		params structures.ChainParams,
		owner common.Address,
	) (common.Address, error)
}
