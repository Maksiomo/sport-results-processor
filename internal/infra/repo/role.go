package repo

import (
	"context"
	"database/sql"
	"errors"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
	"sport-results-pocessor/internal/domain/model"
	"sport-results-pocessor/internal/infra/repo/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
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

func (r *RoleRepo) List(ctx context.Context) ([]*model.Role, error) {
	ents, err := entity.Roles().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Roles")
		return nil, err
	}

	res := make([]*model.Role, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *RoleRepo) One(ctx context.Context, id int64) (*model.Role, error) {
	ent, err := entity.Roles(
		entity.RoleWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Role", zap.Int64("Role.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *RoleRepo) Create(ctx context.Context, m *model.Role) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Role", zap.Error(err), zap.Any("Role.object", m))
		return err
	}

	m.ID = ent.ID

	return nil
}

func (*RoleRepo) toEntity(m *model.Role) *entity.Role {
	return &entity.Role{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}

func (*RoleRepo) fromEntity(e *entity.Role) *model.Role {
	return &model.Role{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
	}
}
