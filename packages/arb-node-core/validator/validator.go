package validator

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"math/big"
)

type Validator struct {
	rollup         *ethbridge.Rollup
	validatorUtils *ethbridge.ValidatorUtils
	client         ethutils.EthClient
}

func (v *Validator) removeOldStakers(ctx context.Context) (*types.Transaction, error) {
	stakersToEliminate, err := v.validatorUtils.RefundableStakers(ctx)
	if err != nil {
		return nil, err
	}
	return v.validatorUtils.RefundStakers(ctx, stakersToEliminate)
}

func (v *Validator) resolveNextNode(ctx context.Context) error {
	config, err := v.validatorUtils.GetConfig(ctx)
	if err != nil {
		return err
	}
	latestConfirmedIndex, err := v.rollup.LatestConfirmedNode(ctx)
	if err != nil {
		return err
	}

	firstUnresolvedIndex, err := v.rollup.FirstUnresolvedNode(ctx)
	if err != nil {
		return err
	}

	firstUnresolved, err := v.rollup.GetNode(ctx, firstUnresolvedIndex)
	if err != nil {
		return err
	}

	firstUnresolvedPrev, err := firstUnresolved.Prev(ctx)
	if err != nil {
		return err
	}

	if firstUnresolvedPrev.Cmp(latestConfirmedIndex) != 0 {
		v.rollup.RejectNextNodeOutOfOrder(ctx)
		return nil
	}

	currentBlock, err := v.client.BlockInfoByNumber(ctx, nil)
	if err != nil {
		return err
	}
	height := (*big.Int)(currentBlock.Number)

	deadline, err := firstUnresolved.DeadlineBlock(ctx)
	if err != nil {
		return err
	}

	if height.Cmp(deadline) < 0 {
		// We're not past the deadline so there's nothing to do
		return nil
	}

	return nil
}
