package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type TeamPersonRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewTeamPersonRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *TeamPersonRepoFactory {
	return &TeamPersonRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *TeamPersonRepoFactory) Create(ctx context.Context, db pgclient.DB) *TeamPersonRepo {
	return &TeamPersonRepo{
		log: f.log.WithComponent(ctx, "team-person-repo"),
		db:  db,
	}
}

type TeamPersonRepo struct {
	log logger.Logger
	db  pgclient.DB
}
