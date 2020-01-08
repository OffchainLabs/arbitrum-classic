package arbbridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type OneStepProof interface {
	ValidateProof(
		auth *bind.CallOpts,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertionStub,
		proof []byte,
	) (*big.Int, error)
}
