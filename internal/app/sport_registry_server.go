package app

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/config"
	"sport-results-pocessor/internal/common/adapter/logger"

	"go.uber.org/fx"
)

func SportRegistryServer(
	lf fx.Lifecycle,
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	sportR 
)