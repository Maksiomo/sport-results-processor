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

type CountryRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCountryRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CountryRepoFactory {
	return &CountryRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CountryRepoFactory) Create(ctx context.Context, db pgclient.DB) *CountryRepo {
	return &CountryRepo{
		log: f.log.WithComponent(ctx, "country-repo"),
		db:  db,
	}
}

type CountryRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *CountryRepo) List(ctx context.Context) ([]*model.Country, error) {
	ents, err := entity.Countries().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list countries")
		return nil, err
	}

	res := make([]*model.Country, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *CountryRepo) One(ctx context.Context, id int64) (*model.Country, error) {
	ent, err := entity.Countries(
		entity.CountryWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "One").Error("cannot find countries", zap.Int64("country.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *CountryRepo) Create(ctx context.Context, m *model.Country) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert country", zap.Error(err), zap.Any("country.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*CountryRepo) toEntity(m *model.Country) *entity.Country {
	return &entity.Country{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}

func (*CountryRepo) fromEntity(e *entity.Country) *model.Country {
	return &model.Country{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
	}
}
