package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type CompetitionTeamsRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCompetitionTeamsRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CompetitionTeamsRepoFactory {
	return &CompetitionTeamsRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CompetitionTeamsRepoFactory) Create(ctx context.Context, db pgclient.DB) *CompetitionTeamsRepo {
	return &CompetitionTeamsRepo{
		log: f.log.WithComponent(ctx, "competition-teams-repo"),
		db:  db,
	}
}

type CompetitionTeamsRepo struct {
	log logger.Logger
	db  pgclient.DB
}
