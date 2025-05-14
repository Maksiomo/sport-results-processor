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

type MatchRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewMatchRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *MatchRepoFactory {
	return &MatchRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *MatchRepoFactory) Create(ctx context.Context, db pgclient.DB) *MatchRepo {
	return &MatchRepo{
		log: f.log.WithComponent(ctx, "match-repo"),
		db:  db,
	}
}

type MatchRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *MatchRepo) List(ctx context.Context) ([]*model.Match, error) {
	ents, err := entity.Matches().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Matches")
		return nil, err
	}

	res := make([]*model.Match, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *MatchRepo) One(ctx context.Context, id int64) (*model.Match, error) {
	ent, err := entity.Matches(
		entity.MatchWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Match", zap.Int64("Match.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *MatchRepo) Create(ctx context.Context, m *model.Match) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Match", zap.Error(err), zap.Any("Match.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (r *MatchRepo) SetTxHash(ctx context.Context, id int64, txHash string) error {
	ct, err := entity.Matches(
        entity.MatchWhere.ID.EQ(id),
    ).One(ctx, r.db)
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not get ent", zap.Error(err))
		return err
    }

    // 2. Меняем поле
    ct.TXHash = null.StringFrom(txHash)

    // 3. Обновляем только record_hash
    _, err = ct.Update(ctx, r.db, boil.Whitelist(entity.MatchColumns.TXHash))
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not update ent", zap.Error(err))
		return err
    }
    return nil
}

func (*MatchRepo) toEntity(m *model.Match) *entity.Match {
	return &entity.Match{
		ID:         m.ID,
		StageID:    m.StageID,
		MatchTime:  m.MatchTime,
		LocationID: null.Int64FromPtr(m.LocationID),
		Metadata:   null.JSONFrom(m.Metadata),
		CreatedAt:  m.CreatedAt,
		TXHash:     null.StringFromPtr(m.TXHash),
	}
}

func (*MatchRepo) fromEntity(e *entity.Match) *model.Match {
	return &model.Match{
		ID:         e.ID,
		StageID:    e.StageID,
		MatchTime:  e.MatchTime,
		LocationID: e.LocationID.Ptr(),
		Metadata:   e.Metadata.JSON,
		CreatedAt:  e.CreatedAt,
		TXHash:     e.TXHash.Ptr(),
	}
}
