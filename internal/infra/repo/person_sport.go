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

type PersonSportRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewPersonSportRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *PersonSportRepoFactory {
	return &PersonSportRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *PersonSportRepoFactory) Create(ctx context.Context, db pgclient.DB) *PersonSportRepo {
	return &PersonSportRepo{
		log: f.log.WithComponent(ctx, "person-sport-repo"),
		db:  db,
	}
}

type PersonSportRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *PersonSportRepo) List(ctx context.Context) ([]*model.PersonSport, error) {
	ents, err := entity.PersonSports().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Person Sports")
		return nil, err
	}

	res := make([]*model.PersonSport, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *PersonSportRepo) One(ctx context.Context, id int64) (*model.PersonSport, error) {
	ent, err := entity.PersonSports(
		entity.PersonSportWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Person Sport", zap.Int64("PersonSport.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *PersonSportRepo) Create(ctx context.Context, m *model.PersonSport) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Person Sport", zap.Error(err), zap.Any("PersonSport.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*PersonSportRepo) toEntity(m *model.PersonSport) *entity.PersonSport {
	return &entity.PersonSport{
		ID:        m.ID,
		PersonID:  m.PersonID,
		SportID:   m.SportID,
		CreatedAt: m.CreatedAt,
	}
}

func (*PersonSportRepo) fromEntity(e *entity.PersonSport) *model.PersonSport {
	return &model.PersonSport{
		ID:        e.ID,
		PersonID:  e.PersonID,
		SportID:   e.SportID,
		CreatedAt: e.CreatedAt,
	}
}
