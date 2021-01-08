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

type ValidatorUtilsWatcher struct {
	con           *ethbridgecontracts.ValidatorUtils
	rollupAddress ethcommon.Address
}

func NewValidatorUtilsWatcher(address, rollupAddress ethcommon.Address, client ethutils.EthClient) (*ValidatorUtilsWatcher, error) {
	con, err := ethbridgecontracts.NewValidatorUtils(address, client)
	if err != nil {
		return nil, err
	}

	return &ValidatorUtilsWatcher{
		con:           con,
		rollupAddress: rollupAddress,
	}, nil
}

func (v *ValidatorUtilsWatcher) RefundableStakers(ctx context.Context) ([]common.Address, error) {
	addresses, err := v.con.RefundableStakers(&bind.CallOpts{Context: ctx}, v.rollupAddress)
	if err != nil {
		return nil, err
	}
	return common.AddressArrayFromEth(addresses), nil
}

func (v *ValidatorUtilsWatcher) SuccessorNodes(ctx context.Context, node NodeID) ([]*big.Int, error) {
	return v.con.SuccessorNodes(&bind.CallOpts{Context: ctx}, v.rollupAddress, node)
}

func (v *ValidatorUtilsWatcher) StakedNodes(ctx context.Context, staker common.Address) ([]*big.Int, error) {
	return v.con.StakedNodes(&bind.CallOpts{Context: ctx}, v.rollupAddress, staker.ToEthAddress())
}

type ValidatorUtils struct {
	*ValidatorUtilsWatcher
	auth *TransactAuth
}

func NewValidatorUtils(address, rollupAddress ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*ValidatorUtils, error) {
	watcher, err := NewValidatorUtilsWatcher(address, rollupAddress, client)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtils{
		ValidatorUtilsWatcher: watcher,
		auth:                  auth,
	}, nil
}

func (v *ValidatorUtils) RejectNextNode(ctx context.Context, stakers []common.Address) (*types.Transaction, error) {
	return v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return v.con.RefundStakers(auth, v.rollupAddress, common.AddressArrayToEth(stakers))
	})
}
