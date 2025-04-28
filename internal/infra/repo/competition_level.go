package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type CompetitionLevelRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCompetitionLevelRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CompetitionLevelRepoFactory {
	return &CompetitionLevelRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CompetitionLevelRepoFactory) Create(ctx context.Context, db pgclient.DB) *CompetitionLevelRepo {
	return &CompetitionLevelRepo{
		log: f.log.WithComponent(ctx, "competition-level-repo"),
		db:  db,
	}
}

type CompetitionLevelRepo struct {
	log logger.Logger
	db  pgclient.DB
}
