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

func NewExportRpcServer(txdb *txdb.TxDB, arbcore core.ArbCore, pathPrefix string) (*ExportRpcServer, error) {
	ethDbPath := path.Join(pathPrefix, "nitro", "l2chaindata")
	err := os.MkdirAll(ethDbPath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	db, err := NewCrossDB(txdb, ethDbPath)
	if err != nil {
		return nil, err
	}
	return &ExportRpcServer{
		db:         db,
		arbcore:    arbcore,
		pathPrefix: pathPrefix,
	}, nil
}

func (r *ExportRpcServer) ExportBlocks(ctx context.Context, blockNumber rpc.BlockNumber) error {
	return r.db.FillerUp(ctx, uint64(blockNumber))
}

func (r *ExportRpcServer) ExportState(blockNumber rpc.BlockNumber) error {
	return ExportState(r.arbcore, uint64(blockNumber), path.Join(r.pathPrefix, "state", hexutil.EncodeUint64(uint64(blockNumber))))
}
