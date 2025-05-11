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

type PersonRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewPersonRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *PersonRepoFactory {
	return &PersonRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *PersonRepoFactory) Create(ctx context.Context, db pgclient.DB) *PersonRepo {
	return &PersonRepo{
		log: f.log.WithComponent(ctx, "person-repo"),
		db:  db,
	}
}

type PersonRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *PersonRepo) List(ctx context.Context) ([]*model.Person, error) {
	ents, err := entity.People().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list competition levels")
		return nil, err
	}

	res := make([]*model.Person, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *PersonRepo) One(ctx context.Context, id int64) (*model.Person, error) {
	ent, err := entity.People(
		entity.PersonWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Person", zap.Int64("Person.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *PersonRepo) Create(ctx context.Context, m *model.Person) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Person", zap.Error(err), zap.Any("Person.object", m))
		return err
	}

	m.ID = ent.ID

	return nil
}

func (*PersonRepo) toEntity(m *model.Person) *entity.Person {
	return &entity.Person{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}

func (*PersonRepo) fromEntity(e *entity.Person) *model.Person {
	return &model.Person{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
	}
}
