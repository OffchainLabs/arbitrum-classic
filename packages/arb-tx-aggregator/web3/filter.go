package web3

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
	"math/big"
)

type PublicFilterAPI interface {
	NewPendingTransactionFilter() rpc.ID
	NewPendingTransactions(ctx context.Context) (*rpc.Subscription, error)
	NewBlockFilter() rpc.ID
	NewHeads(ctx context.Context) (*rpc.Subscription, error)
	Logs(ctx context.Context, crit filters.FilterCriteria) (*rpc.Subscription, error)
	NewFilter(crit filters.FilterCriteria) (rpc.ID, error)
	GetLogs(ctx context.Context, crit filters.FilterCriteria) ([]*types.Log, error)
	UninstallFilter(id rpc.ID) bool
	GetFilterLogs(ctx context.Context, id rpc.ID) ([]*types.Log, error)
	GetFilterChanges(id rpc.ID) (interface{}, error)
}

type filterAPI struct {
	initialHeight *big.Int
	*filters.PublicFilterAPI
}

func (fa *filterAPI) NewFilter(crit filters.FilterCriteria) (rpc.ID, error) {
	fa.updateCrit(&crit)
	return fa.PublicFilterAPI.NewFilter(crit)
}

func (fa *filterAPI) Logs(ctx context.Context, crit filters.FilterCriteria) (*rpc.Subscription, error) {
	fa.updateCrit(&crit)
	return fa.Logs(ctx, crit)
}

func (fa *filterAPI) GetLogs(ctx context.Context, crit filters.FilterCriteria) ([]*types.Log, error) {
	fa.updateCrit(&crit)
	return fa.PublicFilterAPI.GetLogs(ctx, crit)
}

func (fa *filterAPI) updateCrit(crit *filters.FilterCriteria) {
	if crit.FromBlock.Cmp(fa.initialHeight) < 0 {
		crit.FromBlock = fa.initialHeight
	}
}

func NewFilterAPI(server *aggregator.Server, lightMode bool) PublicFilterAPI {
	return &filterAPI{
		initialHeight:   server.InitialBlockHeight(),
		PublicFilterAPI: filters.NewPublicFilterAPI(server, lightMode),
	}
}
