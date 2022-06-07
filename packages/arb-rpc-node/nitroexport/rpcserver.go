package nitroexport

import (
	"context"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

type ExportRpcServer struct {
	db         *CrossDB
	arbcore    core.ArbCore
	pathPrefix string
}

func NewExportRpcServer(ctx context.Context, txdb *txdb.TxDB, arbcore core.ArbCore, pathPrefix string) (*ExportRpcServer, error) {
	ethDbPath := path.Join(pathPrefix, "nitro", "l2chaindata")
	err := os.MkdirAll(ethDbPath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	db, err := NewCrossDB(txdb, ethDbPath)
	if err != nil {
		return nil, err
	}
	db.Start(ctx)
	return &ExportRpcServer{
		db:         db,
		arbcore:    arbcore,
		pathPrefix: pathPrefix,
	}, nil
}

func (r *ExportRpcServer) ExportHistory(ctx context.Context, blockNumber rpc.BlockNumber) error {
	r.db.UpdateTarget(uint64(blockNumber))
	return r.db.CurrentError()
}

func (r *ExportRpcServer) ExportHistoryStatus(ctx context.Context) (uint64, error) {
	blocks, err := r.db.BlocksExported()
	if err != nil {
		return blocks, err
	}
	return blocks, r.db.CurrentError()
}

func (r *ExportRpcServer) ExportState(ctx context.Context, blockNumber rpc.BlockNumber) error {
	// TODO: stop state export on context cancel
	return ExportState(r.arbcore, uint64(blockNumber), path.Join(r.pathPrefix, "state", hexutil.EncodeUint64(uint64(blockNumber))))
}
