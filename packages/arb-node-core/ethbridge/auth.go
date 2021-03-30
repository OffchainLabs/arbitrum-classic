/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ethbridge

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Stack().Str("component", "ethbridge").Logger()

const (
	smallNonceRepeatCount = 100
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
	err = errors.WithStack(err)

	if auth.Nonce == nil {
		// Not incrementing nonce, so nothing else to do
		if err != nil {
			logger.Error().Err(err).Str("nonce", "nil").Msg("error when nonce not set")
			return addr, nil, err
		}

		logger.Info().Str("nonce", "nil").Hex("sender", t.auth.From.Bytes()).Send()
		return addr, tx, err
	}

	for i := 0; i < smallNonceRepeatCount && err != nil && strings.Contains(err.Error(), smallNonceError); i++ {
		// Increment nonce and try again
		logger.Error().Err(err).Str("nonce", auth.Nonce.String()).Msg("incrementing nonce and submitting tx again")

		t.auth.Nonce = t.auth.Nonce.Add(t.auth.Nonce, big.NewInt(1))
		auth.Nonce = t.auth.Nonce
		addr, tx, _, err = contractFunc(auth)
		err = errors.WithStack(err)

		time.Sleep(100 * time.Millisecond)
	}

	if err != nil {
		logger.Error().Err(err).Str("nonce", auth.Nonce.String()).Send()
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
