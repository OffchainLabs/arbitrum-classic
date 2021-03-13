package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"math/big"
	"strings"
	"sync"
)

var logger = log.With().Caller().Str("component", "ethbridge").Logger()

const (
	smallNonceRepeatCount = 5
	smallNonceError       = "Try increasing the gas price or incrementing the nonce."
)

type TransactAuth struct {
	sync.Mutex
	auth *bind.TransactOpts
}

func NewTransactAuth(auth *bind.TransactOpts) *TransactAuth {
	return &TransactAuth{
		auth: auth,
	}
}

func (t *TransactAuth) makeContract(ctx context.Context, contractFunc func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error)) (ethcommon.Address, *types.Transaction, error) {
	auth := t.getAuth(ctx)

	addr, tx, _, err := contractFunc(auth)

	if auth.Nonce == nil {
		// Not incrementing nonce, so nothing else to do
		if err != nil {
			logger.Error().Stack().Err(err).Str("nonce", "nil").Msg("error when nonce not set")
			return addr, nil, err
		}

		logger.Info().Str("nonce", "nil").Hex("sender", t.auth.From.Bytes()).Send()
		return addr, tx, err
	}

	for i := 0; i < smallNonceRepeatCount && err != nil && strings.Contains(err.Error(), smallNonceError); i++ {
		// Increment nonce and try again
		logger.Error().Stack().Err(err).Str("nonce", auth.Nonce.String()).Msg("incrementing nonce and submitting tx again")

		t.auth.Nonce = t.auth.Nonce.Add(t.auth.Nonce, big.NewInt(1))
		auth.Nonce = t.auth.Nonce
		addr, tx, _, err = contractFunc(auth)
	}

	if err != nil {
		logger.Error().Stack().Err(err).Str("nonce", auth.Nonce.String()).Send()
		return addr, nil, err
	}

	// Transaction successful, increment nonce for next time
	logger.Info().Str("nonce", auth.Nonce.String()).Hex("sender", t.auth.From.Bytes()).Send()

	t.auth.Nonce = t.auth.Nonce.Add(t.auth.Nonce, big.NewInt(1))
	return addr, tx, err
}

func (t *TransactAuth) makeTx(ctx context.Context, txFunc func(auth *bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {
	_, tx, err := t.makeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		tx, err := txFunc(auth)
		return ethcommon.BigToAddress(big.NewInt(0)), tx, nil, err
	})

	return tx, err
}

func (t *TransactAuth) getAuth(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     t.auth.From,
		Nonce:    t.auth.Nonce,
		Signer:   t.auth.Signer,
		Value:    t.auth.Value,
		GasPrice: t.auth.GasPrice,
		GasLimit: t.auth.GasLimit,
		Context:  ctx,
	}
}
