package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type TeamAchievementsRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewTeamAchievementsRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *TeamAchievementsRepoFactory {
	return &TeamAchievementsRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *TeamAchievementsRepoFactory) Create(ctx context.Context, db pgclient.DB) *TeamAchievementsRepo {
	return &TeamAchievementsRepo{
		log: f.log.WithComponent(ctx, "team-achievements-repo"),
		db:  db,
	}
}

type TeamAchievementsRepo struct {
	log logger.Logger
	db  pgclient.DB
}
