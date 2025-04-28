package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type CurrencyRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCurrencyRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CurrencyRepoFactory {
	return &CurrencyRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CurrencyRepoFactory) Create(ctx context.Context, db pgclient.DB) *CurrencyRepo {
	return &CurrencyRepo{
		log: f.log.WithComponent(ctx, "currency-repo"),
		db:  db,
	}
}

type CurrencyRepo struct {
	log logger.Logger
	db  pgclient.DB
}
