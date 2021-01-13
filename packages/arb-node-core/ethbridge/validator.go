package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type Validator struct {
	con     *ethbridgecontracts.Validator
	address ethcommon.Address
	client  ethutils.EthClient
	auth    *TransactAuth
}

func NewValidator(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*Validator, error) {
	con, err := ethbridgecontracts.NewValidator(address, client)
	if err != nil {
		return nil, err
	}
	return &Validator{
		con:     con,
		address: address,
		client:  client,
		auth:    auth,
	}, nil
}

func (v *Validator) Address() common.Address {
	return common.NewAddressFromEth(v.address)
}

func (v *Validator) ExecuteTransaction(ctx context.Context, tx *RawTransaction) (*types.Transaction, error) {
	return v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = tx.Amount
		return v.con.ExecuteTransaction(auth, tx.Data, tx.Dest, tx.Amount)
	})
}

func (v *Validator) ExecuteTransactions(ctx context.Context, txes []*RawTransaction) (*types.Transaction, error) {
	if len(txes) == 0 {
		return nil, nil
	}

	if len(txes) == 1 {
		return v.ExecuteTransaction(ctx, txes[0])
	}

	totalAmount := big.NewInt(0)
	data := make([][]byte, 0, len(txes))
	dest := make([]ethcommon.Address, 0, len(txes))
	amount := make([]*big.Int, 0, len(txes))

	for _, tx := range txes {
		data = append(data, tx.Data)
		dest = append(dest, tx.Dest)
		amount = append(amount, tx.Amount)
		totalAmount = totalAmount.Add(totalAmount, tx.Amount)
	}

	return v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = totalAmount
		return v.con.ExecuteTransactions(auth, data, dest, amount)
	})
}
