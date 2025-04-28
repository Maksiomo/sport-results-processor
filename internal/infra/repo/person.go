package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type PersonRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewPersonRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *PersonRepoFactory {
	return &PersonRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *PersonRepoFactory) Create(ctx context.Context, db pgclient.DB) *PersonRepo {
	return &PersonRepo{
		log: f.log.WithComponent(ctx, "person-repo"),
		db:  db,
	}
}

type PersonRepo struct {
	log logger.Logger
	db  pgclient.DB
}
