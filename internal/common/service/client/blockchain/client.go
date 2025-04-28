package blockchain

import (
	"context"

	"go.uber.org/zap"
)

func NewBlockchainRegistryClient(ctx context.Context, cfg *Config, log *zap.Logger)