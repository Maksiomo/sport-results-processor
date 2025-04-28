package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type PrizeRepoRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewPrizeRepoRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *PrizeRepoRepoFactory {
	return &PrizeRepoRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *PrizeRepoRepoFactory) Create(ctx context.Context, db pgclient.DB) *PrizeRepo {
	return &PrizeRepo{
		log: f.log.WithComponent(ctx, "prize-repo"),
		db:  db,
	}
}

type PrizeRepo struct {
	log logger.Logger
	db  pgclient.DB
}
