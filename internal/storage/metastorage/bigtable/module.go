package bigtable

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(fx.Annotated{
		Name:   "metastorage/bigtable",
		Target: NewFactory,
	}),
)
