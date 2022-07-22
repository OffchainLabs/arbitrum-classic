package nitroexport

import (
	"context"
	"errors"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/rs/zerolog/log"
)

type ExportRPCServer struct {
	db         *CrossDB
	arbcore    core.ArbCore
	pathPrefix string
}

func NewExportRPCServer(ctx context.Context, txDB *txdb.TxDB, arbcore core.ArbCore, pathPrefix string) (*ExportRPCServer, error) {
	ethDbPath := path.Join(pathPrefix, "nitro")
	err := os.MkdirAll(ethDbPath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	db, err := NewCrossDB(txDB, ethDbPath)
	if err != nil {
		return nil, err
	}
	db.Start(ctx)
	return &ExportRPCServer{
		db:         db,
		arbcore:    arbcore,
		pathPrefix: pathPrefix,
	}, nil
}

func (r *ExportRPCServer) blockNumToU6(blockNumber rpc.BlockNumber) (uint64, error) {
	if blockNumber == rpc.LatestBlockNumber {
		latest, err := r.db.txDB.LatestBlock()
		if err != nil {
			return 0, err
		}
		return latest.Header.Number.Uint64(), nil
	}
	if blockNumber < 0 {
		return 0, errors.New("unsupported block number")
	}
	return uint64(blockNumber), nil
}

func (r *ExportRPCServer) ExportHistory(blockNumber rpc.BlockNumber) error {
	blockU64, err := r.blockNumToU6(blockNumber)
	if err != nil {
		return err
	}
	r.db.UpdateTargetBlock(blockU64)
	return r.db.CurrentError()
}

func (r *ExportRPCServer) ExportOutbox(batchNumber hexutil.Uint64) error {
	r.db.UpdateTargetBatch(uint64(batchNumber))
	return r.db.CurrentError()
}

func (r *ExportRPCServer) ExportOutboxStatus() (hexutil.Uint64, error) {
	batches := r.db.BatchesExported()
	return hexutil.Uint64(batches), r.db.CurrentError()
}

func (r *ExportRPCServer) ExportHistoryStatus() (*hexutil.Uint64, error) {
	blocks, err := r.db.BlocksExported()
	if err != nil {
		return nil, err
	}
	err = r.db.CurrentError()
	if err != nil {
		return nil, err
	}
	if blocks == 0 {
		return nil, nil
	}
	hexLastBlock := hexutil.Uint64(blocks - 1)
	return &hexLastBlock, nil
}

func (r *ExportRPCServer) ExportState(blockNumber rpc.BlockNumber) error {
	blockU64, err := r.blockNumToU6(blockNumber)
	if err != nil {
		return err
	}
	err = ExportState(r.arbcore, blockU64, path.Join(r.pathPrefix, "state", hexutil.EncodeUint64(blockU64)))
	if err != nil {
		log.Error().Err(err).Uint64("blockNumber", blockU64).Msg("export state failed")
		return err
	}
	log.Info().Uint64("blockNumber", blockU64).Msg("State export done")
	return nil
}
