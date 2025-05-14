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

type MatchParticipantRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewMatchParticipantRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *MatchParticipantRepoFactory {
	return &MatchParticipantRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *MatchParticipantRepoFactory) Create(ctx context.Context, db pgclient.DB) *MatchParticipantRepo {
	return &MatchParticipantRepo{
		log: f.log.WithComponent(ctx, "match-participant-repo"),
		db:  db,
	}
}

type MatchParticipantRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *MatchParticipantRepo) List(ctx context.Context) ([]*model.MatchParticipant, error) {
	ents, err := entity.MatchParticipants().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Match Participants")
		return nil, err
	}

	res := make([]*model.MatchParticipant, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *MatchParticipantRepo) One(ctx context.Context, id int64) (*model.MatchParticipant, error) {
	ent, err := entity.MatchParticipants(
		entity.MatchParticipantWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Match Participant", zap.Int64("MatchParticipant.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *MatchParticipantRepo) Create(ctx context.Context, m *model.MatchParticipant) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Match Participant", zap.Error(err), zap.Any("MatchParticipant.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (r *MatchParticipantRepo) SetTxHash(ctx context.Context, id int64, txHash string) error {
	ct, err := entity.MatchParticipants(
        entity.MatchParticipantWhere.ID.EQ(id),
    ).One(ctx, r.db)
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not get ent", zap.Error(err))
		return err
    }

    // 2. Меняем поле
    ct.TXHash = null.StringFrom(txHash)

    // 3. Обновляем только record_hash
    _, err = ct.Update(ctx, r.db, boil.Whitelist(entity.MatchParticipantColumns.TXHash))
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not update ent", zap.Error(err))
		return err
    }
    return nil
}

func (*MatchParticipantRepo) toEntity(m *model.MatchParticipant) *entity.MatchParticipant {
	return &entity.MatchParticipant{
		ID:         m.ID,
		MatchID:    m.MatchID,
		TeamID:     m.TeamID,
		Score:      m.Score,
		IsWinner:   m.IsWinner,
		CreatedAt:  m.CreatedAt,
		TXHash:     null.StringFromPtr(m.TXHash),
	}
}

func (*MatchParticipantRepo) fromEntity(e *entity.MatchParticipant) *model.MatchParticipant {
	return &model.MatchParticipant{
		ID:         e.ID,
		MatchID:    e.MatchID,
		TeamID:     e.TeamID,
		Score:      e.Score,
		IsWinner:   e.IsWinner,
		CreatedAt:  e.CreatedAt,
		TXHash:     e.TXHash.Ptr(),
	}
}
