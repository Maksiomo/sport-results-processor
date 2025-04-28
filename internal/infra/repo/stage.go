package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type StageRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewStageRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *StageRepoFactory {
	return &StageRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *StageRepoFactory) Create(ctx context.Context, db pgclient.DB) *StageRepo {
	return &StageRepo{
		log: f.log.WithComponent(ctx, "stage-repo"),
		db:  db,
	}
}

type StageRepo struct {
	log logger.Logger
	db  pgclient.DB
}
