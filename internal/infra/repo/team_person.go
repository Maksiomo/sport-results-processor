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

type TeamPersonRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewTeamPersonRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *TeamPersonRepoFactory {
	return &TeamPersonRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *TeamPersonRepoFactory) Create(ctx context.Context, db pgclient.DB) *TeamPersonRepo {
	return &TeamPersonRepo{
		log: f.log.WithComponent(ctx, "team-person-repo"),
		db:  db,
	}
}

type TeamPersonRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *TeamPersonRepo) List(ctx context.Context) ([]*model.TeamPerson, error) {
	ents, err := entity.TeamPeople().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Team People")
		return nil, err
	}

	res := make([]*model.TeamPerson, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *TeamPersonRepo) One(ctx context.Context, id int64) (*model.TeamPerson, error) {
	ent, err := entity.TeamPeople(
		entity.TeamPersonWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Team Person", zap.Int64("TeamPerson.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *TeamPersonRepo) Create(ctx context.Context, m *model.TeamPerson) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Team Person", zap.Error(err), zap.Any("TeamPerson.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (*TeamPersonRepo) toEntity(m *model.TeamPerson) *entity.TeamPerson {
	return &entity.TeamPerson{
		ID:         m.ID,
		TeamID:     m.TeamID,
		PersonID:   m.PersonID,
		RoleID:     m.RoleID,
		JoinedAt:   m.JoinedAt,
		LeftAt:     null.TimeFromPtr(m.LeftAt),
		CreatedAt:  m.CreatedAt,
		RecordHash: null.StringFrom(m.RecordHash),
		TXHash:     null.StringFromPtr(m.TXHash),
	}
}

func (*TeamPersonRepo) fromEntity(e *entity.TeamPerson) *model.TeamPerson {
	return &model.TeamPerson{
		ID:         e.ID,
		TeamID:     e.TeamID,
		PersonID:   e.PersonID,
		RoleID:     e.RoleID,
		JoinedAt:   e.JoinedAt,
		LeftAt:     e.LeftAt.Ptr(),
		CreatedAt:  e.CreatedAt,
		RecordHash: e.RecordHash.String,
		TXHash:     e.TXHash.Ptr(),
	}
}
