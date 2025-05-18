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

func (r *CompetitionLevelRepo) List(ctx context.Context) ([]*model.CompetitionLevel, error) {
	ents, err := entity.CompetitionLevels().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list competition levels")
		return nil, err
	}

	res := make([]*model.CompetitionLevel, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *CompetitionLevelRepo) One(ctx context.Context, id int64) (*model.CompetitionLevel, error) {
	ent, err := entity.CompetitionLevels(
		entity.CompetitionLevelWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "One").Error("cannot find competition level", zap.Int64("competition_level.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *CompetitionLevelRepo) Create(ctx context.Context, m *model.CompetitionLevel) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert competition level", zap.Error(err), zap.Any("competition_level.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*CompetitionLevelRepo) toEntity(m *model.CompetitionLevel) *entity.CompetitionLevel {
	return &entity.CompetitionLevel{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}

func (*CompetitionLevelRepo) fromEntity(e *entity.CompetitionLevel) *model.CompetitionLevel {
	return &model.CompetitionLevel{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
	}
}
