package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type PersonSportRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewPersonSportRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *PersonSportRepoFactory {
	return &PersonSportRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *PersonSportRepoFactory) Create(ctx context.Context, db pgclient.DB) *PersonSportRepo {
	return &PersonSportRepo{
		log: f.log.WithComponent(ctx, "person-sport-repo"),
		db:  db,
	}
}

type PersonSportRepo struct {
	log logger.Logger
	db  pgclient.DB
}
