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

type CompetitionRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCompetitionRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CompetitionRepoFactory {
	return &CompetitionRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CompetitionRepoFactory) Create(ctx context.Context, db pgclient.DB) *CompetitionRepo {
	return &CompetitionRepo{
		log: f.log.WithComponent(ctx, "competition-repo"),
		db:  db,
	}
}

type CompetitionRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *CompetitionRepo) List(ctx context.Context) ([]*model.Competition, error) {
	ents, err := entity.Competitions().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list competitions")
		return nil, err
	}

	res := make([]*model.Competition, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *CompetitionRepo) One(ctx context.Context, id int64) (*model.Competition, error) {
	ent, err := entity.Competitions(
		entity.CompetitionWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find competition", zap.Int64("competition.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *CompetitionRepo) Create(ctx context.Context, m *model.Competition) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert competition", zap.Error(err), zap.Any("competition.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*CompetitionRepo) toEntity(m *model.Competition) *entity.Competition {
	return &entity.Competition{
		ID:         m.ID,
		Name:       m.Name,
		SportID:    m.SportID,
		LocationID: m.LocationID,
		LevelID:    m.LevelID,
		CreatedAt:  m.CreatedAt,
		RecordHash: null.StringFrom(m.RecordHash),
		TXHash:     null.StringFromPtr(m.TXHash),
	}
}

func (*CompetitionRepo) fromEntity(e *entity.Competition) *model.Competition {
	return &model.Competition{
		ID:         e.ID,
		Name:       e.Name,
		SportID:    e.SportID,
		LocationID: e.LocationID,
		LevelID:    e.LevelID,
		CreatedAt:  e.CreatedAt,
		RecordHash: e.RecordHash.String,
		TXHash:     &e.TXHash.String,
	}
}
