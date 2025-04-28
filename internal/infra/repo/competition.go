package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type CompetitionRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCompetitionRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CompetitionRepoFactory {
	return &CompetitionRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CompetitionRepoFactory) Create(ctx context.Context, db pgclient.DB) *CompetitionRepo {
	return &CompetitionRepo{
		log: f.log.WithComponent(ctx, "competition-repo"),
		db:  db,
	}
}

type CompetitionRepo struct {
	log logger.Logger
	db  pgclient.DB
}
