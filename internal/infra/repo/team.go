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

func (r *TeamRepo) List(ctx context.Context) ([]*model.Team, error) {
	ents, err := entity.Teams().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Teams")
		return nil, err
	}

	res := make([]*model.Team, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *TeamRepo) One(ctx context.Context, id int64) (*model.Team, error) {
	ent, err := entity.Teams(
		entity.TeamWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Team", zap.Int64("Team.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *TeamRepo) Create(ctx context.Context, m *model.Team) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Team", zap.Error(err), zap.Any("Team.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*TeamRepo) toEntity(m *model.Team) *entity.Team {
	return &entity.Team{
		ID:             m.ID,
		Name:           m.Name,
		CountryID:      m.CountryID,
		FoundationDate: m.FoundationDate,
		CreatedAt:      m.CreatedAt,
		RecordHash:     null.BytesFrom(m.RecordHash),
		TXHash:         null.StringFromPtr(m.TXHash),
	}
}

func (*TeamRepo) fromEntity(e *entity.Team) *model.Team {
	return &model.Team{
		ID:             e.ID,
		Name:           e.Name,
		CountryID:      e.CountryID,
		FoundationDate: e.FoundationDate,
		CreatedAt:      e.CreatedAt,
		RecordHash:     e.RecordHash.Bytes,
		TXHash:         e.TXHash.Ptr(),
	}
}
