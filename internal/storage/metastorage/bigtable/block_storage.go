package bigtable

import (
	"context"

	"cloud.google.com/go/bigtable"

	"github.com/coinbase/chainstorage/internal/storage/metastorage/internal"
	"github.com/coinbase/chainstorage/internal/utils/instrument"
	"github.com/coinbase/chainstorage/protos/coinbase/chainstorage"
)

type (
	blockStorageImpl struct {
		adminClient                      *bigtable.AdminClient
		client                           *bigtable.Client
		blockStartHeight                 uint64
		instrumentPersistBlockMetas      instrument.Call
		instrumentGetLatestBlock         instrument.Call
		instrumentGetBlockByHash         instrument.Call
		instrumentGetBlockByHeight       instrument.Call
		instrumentGetBlocksByHeightRange instrument.Call
	}
)

func newBlockStorage(params Params, adminClient *bigtable.AdminClient, client *bigtable.Client) (internal.BlockStorage, error) {

	metrics := params.Metrics.SubScope("block_storage_bigtable")
	accessor := blockStorageImpl{
		adminClient:                      adminClient,
		client:                           client,
		blockStartHeight:                 params.Config.Chain.BlockStartHeight,
		instrumentPersistBlockMetas:      instrument.NewCall(metrics, "persist_block_metas"),
		instrumentGetLatestBlock:         instrument.NewCall(metrics, "get_latest_block"),
		instrumentGetBlockByHash:         instrument.NewCall(metrics, "get_block_by_hash"),
		instrumentGetBlockByHeight:       instrument.NewCall(metrics, "get_block_by_height"),
		instrumentGetBlocksByHeightRange: instrument.NewCall(metrics, "get_blocks_by_height_range"),
	}
	return &accessor, nil
}

// GetBlockByHash implements internal.BlockStorage.
func (*blockStorageImpl) GetBlockByHash(ctx context.Context, tag uint32, height uint64, blockHash string) (*chainstorage.BlockMetadata, error) {
	panic("unimplemented")
}

// GetBlockByHeight implements internal.BlockStorage.
func (*blockStorageImpl) GetBlockByHeight(ctx context.Context, tag uint32, height uint64) (*chainstorage.BlockMetadata, error) {
	panic("unimplemented")
}

// GetBlocksByHeightRange implements internal.BlockStorage.
func (*blockStorageImpl) GetBlocksByHeightRange(ctx context.Context, tag uint32, startHeight uint64, endHeight uint64) ([]*chainstorage.BlockMetadata, error) {
	panic("unimplemented")
}

// GetLatestBlock implements internal.BlockStorage.
func (*blockStorageImpl) GetLatestBlock(ctx context.Context, tag uint32) (*chainstorage.BlockMetadata, error) {
	panic("unimplemented")
}

// PersistBlockMetas implements internal.BlockStorage.
func (*blockStorageImpl) PersistBlockMetas(ctx context.Context, updateWatermark bool, blocks []*chainstorage.BlockMetadata, lastBlock *chainstorage.BlockMetadata) error {
	panic("unimplemented")
}
