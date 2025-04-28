package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type MatchRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewMatchRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *MatchRepoFactory {
	return &MatchRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *MatchRepoFactory) Create(ctx context.Context, db pgclient.DB) *MatchRepo {
	return &MatchRepo{
		log: f.log.WithComponent(ctx, "match-repo"),
		db:  db,
	}
}

type MatchRepo struct {
	log logger.Logger
	db  pgclient.DB
}
