package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type TeamRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewTeamRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *TeamRepoFactory {
	return &TeamRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *TeamRepoFactory) Create(ctx context.Context, db pgclient.DB) *TeamRepo {
	return &TeamRepo{
		log: f.log.WithComponent(ctx, "team-repo"),
		db:  db,
	}
}

type TeamRepo struct {
	log logger.Logger
	db  pgclient.DB
}
