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

type SportRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewSportRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *SportRepoFactory {
	return &SportRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *SportRepoFactory) Create(ctx context.Context, db pgclient.DB) *SportRepo {
	return &SportRepo{
		log: f.log.WithComponent(ctx, "sport-repo"),
		db:  db,
	}
}

type SportRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *SportRepo) List(ctx context.Context) ([]*model.Sport, error) {
	ents, err := entity.Sports().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Sports")
		return nil, err
	}

	res := make([]*model.Sport, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *SportRepo) One(ctx context.Context, id int64) (*model.Sport, error) {
	ent, err := entity.Sports(
		entity.SportWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Sport", zap.Int64("Sport.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *SportRepo) Create(ctx context.Context, m *model.Sport) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Sport", zap.Error(err), zap.Any("Sport.object", m))
		return err
	}

	m.ID = ent.ID

	return nil
}

func (*SportRepo) toEntity(m *model.Sport) *entity.Sport {
	return &entity.Sport{
		ID:          m.ID,
		Name:        m.Name,
		MinTeamSize: m.MinTeamSize,
		MaxTeamSize: m.MaxTeamSize,
		Description: null.StringFromPtr(m.Description),
		CreatedAt:   m.CreatedAt,
		RecordHash:  null.BytesFrom(m.RecordHash),
		TXHash:      null.StringFromPtr(m.TXHash),
	}
}

func (*SportRepo) fromEntity(e *entity.Sport) *model.Sport {
	return &model.Sport{
		ID:          e.ID,
		Name:        e.Name,
		MinTeamSize: e.MinTeamSize,
		MaxTeamSize: e.MaxTeamSize,
		Description: e.Description.Ptr(),
		CreatedAt:   e.CreatedAt,
		RecordHash:  e.RecordHash.Bytes,
		TXHash:      e.TXHash.Ptr(),
	}
}
