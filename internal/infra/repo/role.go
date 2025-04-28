package repo

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type RoleRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewRoleRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *RoleRepoFactory {
	return &RoleRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *RoleRepoFactory) Create(ctx context.Context, db pgclient.DB) *RoleRepo {
	return &RoleRepo{
		log: f.log.WithComponent(ctx, "role-repo"),
		db:  db,
	}
}

type RoleRepo struct {
	log logger.Logger
	db  pgclient.DB
}
