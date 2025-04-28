package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type LocationRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewLocationRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *LocationRepoFactory {
	return &LocationRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *LocationRepoFactory) Create(ctx context.Context, db pgclient.DB) *LocationRepo {
	return &LocationRepo{
		log: f.log.WithComponent(ctx, "location-repo"),
		db:  db,
	}
}

type LocationRepo struct {
	log logger.Logger
	db  pgclient.DB
}
