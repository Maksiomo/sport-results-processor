package repo

import (
	"context"
	"database/sql"
	"errors"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
	"sport-results-pocessor/internal/domain/model"
	"sport-results-pocessor/internal/infra/repo/entity"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
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

func (r *StageRepo) List(ctx context.Context) ([]*model.Stage, error) {
	ents, err := entity.Stages().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Stages")
		return nil, err
	}

	res := make([]*model.Stage, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *StageRepo) One(ctx context.Context, id int64) (*model.Stage, error) {
	ent, err := entity.Stages(
		entity.StageWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Stage", zap.Int64("Stage.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *StageRepo) Create(ctx context.Context, m *model.Stage) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Stage", zap.Error(err), zap.Any("Stage.object", m))
		return err
	}

	m.ID = ent.ID

	return nil
}

func (*StageRepo) toEntity(m *model.Stage) *entity.Stage {
	return &entity.Stage{
		ID:            m.ID,
		CompetitionID: m.CompetitionID,
		Name:          m.Name,
		StartTime:     m.StartTime,
		EndTime:       null.TimeFromPtr(m.EndTime),
		CreatedAt:     m.CreatedAt,
	}
}

func (*StageRepo) fromEntity(e *entity.Stage) *model.Stage {
	return &model.Stage{
		ID:            e.ID,
		CompetitionID: e.CompetitionID,
		Name:          e.Name,
		StartTime:     e.StartTime,
		EndTime:       e.EndTime.Ptr(),
		CreatedAt:     e.CreatedAt,
	}
}
