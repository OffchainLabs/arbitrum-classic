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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"
)

var validatorABI abi.ABI
var walletCreatedID ethcommon.Hash

func init() {
	parsedValidator, err := abi.JSON(strings.NewReader(ethbridgecontracts.ValidatorABI))
	if err != nil {
		panic(err)
	}
	validatorABI = parsedValidator

	parsedValidatorWalletCreator, err := abi.JSON(strings.NewReader(ethbridgecontracts.ValidatorWalletCreatorABI))
	if err != nil {
		panic(err)
	}
	walletCreatedID = parsedValidatorWalletCreator.Events["WalletCreated"].ID
}

type ValidatorWallet struct {
	con               *ethbridgecontracts.Validator
	address           *ethcommon.Address
	onWalletCreated   func(ethcommon.Address)
	client            ethutils.EthClient
	auth              transactauth.TransactAuth
	rollupAddress     ethcommon.Address
	walletFactoryAddr ethcommon.Address
	rollupFromBlock   int64
	blockSearchSize   int64
}

func NewValidator(
	address *ethcommon.Address,
	walletFactoryAddr,
	rollupAddress ethcommon.Address,
	client ethutils.EthClient,
	auth transactauth.TransactAuth,
	rollupFromBlock int64,
	blockSearchSize int64,
	onWalletCreated func(ethcommon.Address),
) (*ValidatorWallet, error) {
	var con *ethbridgecontracts.Validator
	if address != nil {
		var err error
		con, err = ethbridgecontracts.NewValidator(*address, client)
		if err != nil {
			return nil, err
		}
	}
	return &ValidatorWallet{
		con:               con,
		address:           address,
		onWalletCreated:   onWalletCreated,
		client:            client,
		auth:              auth,
		rollupAddress:     rollupAddress,
		walletFactoryAddr: walletFactoryAddr,
		rollupFromBlock:   rollupFromBlock,
		blockSearchSize:   blockSearchSize,
	}, nil
}

// May be the nil if the wallet hasn't been deployed yet
func (v *ValidatorWallet) Address() *ethcommon.Address {
	return v.address
}

func (v *ValidatorWallet) From() common.Address {
	return common.NewAddressFromEth(v.auth.From())
}

func (v *ValidatorWallet) RollupAddress() common.Address {
	return common.NewAddressFromEth(v.rollupAddress)
}

func (v *ValidatorWallet) executeTransaction(ctx context.Context, tx *types.Transaction) (*arbtransaction.ArbTransaction, error) {
	return transactauth.MakeTx(ctx, v.auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = tx.Value()
		return v.con.ExecuteTransaction(auth, tx.Data(), *tx.To(), tx.Value())
	})
}

func (v *ValidatorWallet) CreateWalletIfNeeded(ctx context.Context) error {
	if v.con != nil {
		return nil
	}
	if v.address == nil {
		addr, err := CreateValidatorWallet(ctx, v.walletFactoryAddr, v.rollupFromBlock, v.blockSearchSize, v.auth, v.client)
		if err != nil {
			return err
		}
		v.address = &addr
		if v.onWalletCreated != nil {
			v.onWalletCreated(addr)
		}
	}
	con, err := ethbridgecontracts.NewValidator(*v.address, v.client)
	if err != nil {
		return err
	}
	v.con = con
	return nil
}

func combineTxes(txes []*types.Transaction) ([][]byte, []ethcommon.Address, []*big.Int, *big.Int) {
	totalAmount := big.NewInt(0)
	data := make([][]byte, 0, len(txes))
	dest := make([]ethcommon.Address, 0, len(txes))
	amount := make([]*big.Int, 0, len(txes))

	for _, tx := range txes {
		data = append(data, tx.Data())
		dest = append(dest, *tx.To())
		amount = append(amount, tx.Value())
		totalAmount = totalAmount.Add(totalAmount, tx.Value())
	}
	return data, dest, amount, totalAmount
}

// Not thread safe! Don't call this from multiple threads at the same time.
func (v *ValidatorWallet) ExecuteTransactions(ctx context.Context, builder *BuilderBackend) (*arbtransaction.ArbTransaction, error) {
	txes := builder.transactions
	if len(txes) == 0 {
		return nil, nil
	}

	if len(txes) == 1 {
		arbTx, err := v.executeTransaction(ctx, txes[0])
		if err != nil {
			return nil, err
		}
		builder.transactions = nil
		return arbTx, nil
	}

	totalAmount := big.NewInt(0)
	data := make([][]byte, 0, len(txes))
	dest := make([]ethcommon.Address, 0, len(txes))
	amount := make([]*big.Int, 0, len(txes))

	for _, tx := range txes {
		data = append(data, tx.Data())
		dest = append(dest, *tx.To())
		amount = append(amount, tx.Value())
		totalAmount = totalAmount.Add(totalAmount, tx.Value())
	}

	err := v.CreateWalletIfNeeded(ctx)
	if err != nil {
		return nil, err
	}

	arbTx, err := transactauth.MakeTx(ctx, v.auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = totalAmount
		return v.con.ExecuteTransactions(auth, data, dest, amount)
	})
	if err != nil {
		return nil, err
	}
	builder.transactions = nil
	return arbTx, nil
}

func (v *ValidatorWallet) ReturnOldDeposits(ctx context.Context, stakers []common.Address) (*arbtransaction.ArbTransaction, error) {
	return transactauth.MakeTx(ctx, v.auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return v.con.ReturnOldDeposits(auth, v.rollupAddress, common.AddressArrayToEth(stakers))
	})
}

func (v *ValidatorWallet) TimeoutChallenges(ctx context.Context, challenges []common.Address) (*arbtransaction.ArbTransaction, error) {
	return transactauth.MakeTx(ctx, v.auth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return v.con.TimeoutChallenges(auth, common.AddressArrayToEth(challenges))
	})
}

func CreateValidatorWallet(
	ctx context.Context,
	validatorWalletFactoryAddr ethcommon.Address,
	initialFromBlock int64,
	blockSearchSize int64,
	transactAuth transactauth.TransactAuth,
	client ethutils.EthClient,
) (ethcommon.Address, error) {
	walletCreator, err := ethbridgecontracts.NewValidatorWalletCreator(validatorWalletFactoryAddr, client)
	if err != nil {
		return ethcommon.Address{}, errors.WithStack(err)
	}

	latestHeader, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return ethcommon.Address{}, errors.WithStack(err)
	}
	latestBlockHeight := latestHeader.Number.Int64()
	currentFromBlock := initialFromBlock
	var currentToBlock int64
	if blockSearchSize > 0 {
		currentToBlock = initialFromBlock + blockSearchSize
	} else {
		// Search all blocks at once, must use log caching better than go-ethereum for large block searches
		currentToBlock = latestBlockHeight
	}
	var logs []types.Log
	for len(logs) == 0 && currentFromBlock <= latestBlockHeight {
		logger.Debug().Int64("fromBlock", currentFromBlock).Int64("toBlock", currentToBlock).Msg("searching for validator smart contract")
		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(currentFromBlock),
			ToBlock:   big.NewInt(currentToBlock),
			Addresses: []ethcommon.Address{validatorWalletFactoryAddr},
			Topics:    [][]ethcommon.Hash{{walletCreatedID}, nil, {transactAuth.From().Hash()}},
		}
		logs, err = client.FilterLogs(ctx, query)
		if err != nil {
			return ethcommon.Address{}, errors.WithStack(err)
		}
		currentFromBlock = currentToBlock + 1
		currentToBlock = currentFromBlock + blockSearchSize
	}
	if len(logs) > 1 {
		return ethcommon.Address{}, errors.New("more than one validator wallet created for address")
	} else if len(logs) == 1 {
		log := logs[0]
		parsed, err := walletCreator.ParseWalletCreated(log)
		if err != nil {
			return ethcommon.Address{}, err
		}
		return parsed.WalletAddress, err
	}

	arbTx, err := transactauth.MakeTx(ctx, transactAuth, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return walletCreator.CreateWallet(auth)
	})
	if err != nil {
		return ethcommon.Address{}, err
	}

	simulatedBackend, ok := client.(*ethutils.SimulatedEthClient)
	if ok {
		simulatedBackend.Commit()
	}

	receipt, err := transactauth.WaitForReceiptWithResultsAndReplaceByFee(
		ctx,
		client,
		transactAuth.From(),
		arbTx,
		"CreateWallet",
		transactAuth,
		transactAuth,
	)
	if err != nil {
		return ethcommon.Address{}, err
	}
	ev, err := walletCreator.ParseWalletCreated(*receipt.Logs[len(receipt.Logs)-1])
	if err != nil {
		return ethcommon.Address{}, errors.WithStack(err)
	}
	return ev.WalletAddress, nil
}
