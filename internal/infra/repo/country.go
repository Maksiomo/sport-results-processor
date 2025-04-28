package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type CountryRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCountryRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CountryRepoFactory {
	return &CountryRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CountryRepoFactory) Create(ctx context.Context, db pgclient.DB) *CountryRepo {
	return &CountryRepo{
		log: f.log.WithComponent(ctx, "country-repo"),
		db:  db,
	}
}

type CountryRepo struct {
	log logger.Logger
	db  pgclient.DB
}
