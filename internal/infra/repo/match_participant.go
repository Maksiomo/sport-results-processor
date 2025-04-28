package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type MatchParticipantRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewMatchParticipantRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *MatchParticipantRepoFactory {
	return &MatchParticipantRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *MatchParticipantRepoFactory) Create(ctx context.Context, db pgclient.DB) *MatchParticipantRepo {
	return &MatchParticipantRepo{
		log: f.log.WithComponent(ctx, "match-participant-repo"),
		db:  db,
	}
}

type MatchParticipantRepo struct {
	log logger.Logger
	db  pgclient.DB
}
