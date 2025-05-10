package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type PrizeRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewPrizeRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *PrizeRepoFactory {
	return &PrizeRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *PrizeRepoFactory) Create(ctx context.Context, db pgclient.DB) *PrizeRepo {
	return &PrizeRepo{
		log: f.log.WithComponent(ctx, "prize-repo"),
		db:  db,
	}
}

type PrizeRepo struct {
	log logger.Logger
	db  pgclient.DB
}
