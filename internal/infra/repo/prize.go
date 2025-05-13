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

type PrizeRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewPrizeRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *PrizeRepoFactory {
	return &PrizeRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *PrizeRepoFactory) Create(ctx context.Context, db pgclient.DB) *PrizeRepo {
	return &PrizeRepo{
		log: f.log.WithComponent(ctx, "prize-repo"),
		db:  db,
	}
}

type PrizeRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *PrizeRepo) List(ctx context.Context) ([]*model.Prize, error) {
	ents, err := entity.Prizes().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list competition levels")
		return nil, err
	}

	res := make([]*model.Prize, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *PrizeRepo) One(ctx context.Context, id int64) (*model.Prize, error) {
	ent, err := entity.Prizes(
		entity.PrizeWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Prize", zap.Int64("Prize.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *PrizeRepo) Create(ctx context.Context, m *model.Prize) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Prize", zap.Error(err), zap.Any("Prize.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*PrizeRepo) toEntity(m *model.Prize) *entity.Prize {
	return &entity.Prize{
		ID:            m.ID,
		CompetitionID: m.CompetitionID,
		PlaceBracket:  m.PlaceBracket,
		CurrencyCode:  m.CurrencyCode,
		CreatedAt:     m.CreatedAt,
		RecordHash:    null.StringFrom(m.RecordHash),
		TXHash:        null.StringFromPtr(m.TXHash),
	}
}

func (*PrizeRepo) fromEntity(e *entity.Prize) *model.Prize {
	return &model.Prize{
		ID:            e.ID,
		CompetitionID: e.CompetitionID,
		PlaceBracket:  e.PlaceBracket,
		CurrencyCode:  e.CurrencyCode,
		PrizeAmount:   e.PrizeAmount,
		CreatedAt:     e.CreatedAt,
		RecordHash:    e.RecordHash.String,
		TXHash:        e.TXHash.Ptr(),
	}
}
