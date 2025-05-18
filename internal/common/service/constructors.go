package service

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/config"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/service/client/blockchain"

	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	NewFxBlockchainClient,
)

func NewFxBlockchainClient(ctx context.Context, cfg *config.Config, log logger.Logger) (*blockchain.BlockchainRegistryClient, error) {
	return blockchain.NewBlockchainRegistryClient(ctx, cfg.Client.Blockchain, log)
}
