package bigtable

import (
	"context"

	"cloud.google.com/go/bigtable"
	"golang.org/x/xerrors"

	"go.uber.org/fx"

	"github.com/coinbase/chainstorage/internal/storage/metastorage/internal"
	"github.com/coinbase/chainstorage/internal/utils/fxparams"
)

type (
	metaStorageImpl struct {
		internal.BlockStorage
		internal.EventStorage
	}

	Params struct {
		fx.In
		fxparams.Params
	}

	metaStorageFactory struct {
		params Params
	}
)

func NewMetaStorage(params Params) (internal.Result, error) {
	ctx := context.Background()
	config := params.Config.GCP

	adminClient, err := bigtable.NewAdminClient(ctx, config.Project, config.BigTableInstance)
	if err != nil {
		return internal.Result{}, xerrors.Errorf("failed to create bigtable admin client: %w", err)
	}
	client, err := bigtable.NewClient(ctx, config.Project, config.BigTableInstance)
	if err != nil {
		return internal.Result{}, xerrors.Errorf("failed to create bigtable client: %w", err)
	}
	blockStorage, err := newBlockStorage(params, adminClient, client)
	if err != nil {
		return internal.Result{}, xerrors.Errorf("failed create new BlockStorage: %w", err)
	}

	eventStorage, err := newEventStorage(params, adminClient, client)
	if err != nil {
		return internal.Result{}, xerrors.Errorf("failed create new EventStorage: %w", err)
	}

	metaStorage := &metaStorageImpl{
		BlockStorage: blockStorage,
		EventStorage: eventStorage,
	}

	return internal.Result{
		BlockStorage: blockStorage,
		EventStorage: eventStorage,
		MetaStorage:  metaStorage,
	}, nil
}

// Create implements internal.MetaStorageFactory.
func (f *metaStorageFactory) Create() (internal.Result, error) {
	return NewMetaStorage(f.params)
}

func NewFactory(params Params) internal.MetaStorageFactory {
	return &metaStorageFactory{params}
}
