package adapter

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/config"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"

	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	config.NewConfig,

	NewFxLogger,
	fx.Annotate(NewFxPgClient, fx.OnStop(func(db pgclient.DB) error { return db.Close() })),
)

func NewFxLogger(cfg *config.Config) (logger.Logger, error) {
	return logger.NewLogger(cfg.Logger)
}

func NewFxPgClient(ctx context.Context, cfg *config.Config, logg logger.Logger) (pgclient.DB, error) {
	fxLog := logg.WithComponent(ctx, "pgclient")
	return pgclient.NewPool(ctx, cfg.DB, fxLog, pgclient.WithTransactionLogging(fxLog))
}