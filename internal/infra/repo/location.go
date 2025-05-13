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

type LocationRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewLocationRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *LocationRepoFactory {
	return &LocationRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *LocationRepoFactory) Create(ctx context.Context, db pgclient.DB) *LocationRepo {
	return &LocationRepo{
		log: f.log.WithComponent(ctx, "location-repo"),
		db:  db,
	}
}

type LocationRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *LocationRepo) List(ctx context.Context) ([]*model.Location, error) {
	ents, err := entity.Locations().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Location")
		return nil, err
	}

	res := make([]*model.Location, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *LocationRepo) One(ctx context.Context, id int64) (*model.Location, error) {
	ent, err := entity.Locations(
		entity.LocationWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Location level", zap.Int64("Location.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *LocationRepo) Create(ctx context.Context, m *model.Location) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Location", zap.Error(err), zap.Any("Location.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*LocationRepo) toEntity(m *model.Location) *entity.Location {
	return &entity.Location{
		ID:          m.ID,
		CountryID:   m.CountryID,
		City:        null.StringFromPtr(m.City),
		State:       null.StringFromPtr(m.State),
		Address:     null.StringFromPtr(m.Address),
		FullAddress: m.FullAddress,
		CreatedAt:   m.CreatedAt,
	}
}

func (*LocationRepo) fromEntity(e *entity.Location) *model.Location {
	return &model.Location{
		ID:          e.ID,
		CountryID:   e.CountryID,
		City:        e.City.Ptr(),
		State:       e.State.Ptr(),
		Address:     e.Address.Ptr(),
		FullAddress: e.FullAddress,
		CreatedAt:   e.CreatedAt,
	}
}
