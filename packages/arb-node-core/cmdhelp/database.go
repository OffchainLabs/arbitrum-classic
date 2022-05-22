package cmdhelp

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
	"math/big"
)

func UpdatePrunePoint(ctx context.Context, rollup *ethbridge.RollupWatcher, lookup core.ArbCoreLookup) error {
	// Prune any stale database entries while we wait
	latestNode, err := rollup.LatestConfirmedNode(ctx)
	if err != nil {
		return errors.Wrap(err, "couldn't get latest confirmed node")
	}

	if latestNode.Cmp(big.NewInt(0)) == 0 {
		logger.Info().Msg("no confirmed nodes so nothing to prune")
		return nil
	}

	// Prune checkpoints up to confirmed node before last confirmed node
	previousConfirmedNode := new(big.Int).Sub(latestNode, big.NewInt(1))
	previousNodeInfo, err := rollup.LookupNode(ctx, previousConfirmedNode)
	if err != nil {
		return errors.Wrap(err, "couldn't lookup previous confirmed node "+previousConfirmedNode.String())
	}

	confirmedGas := previousNodeInfo.AfterState().TotalGasConsumed
	lookup.UpdateCheckpointPruningGas(confirmedGas)
	return nil
}
