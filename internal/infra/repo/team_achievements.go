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

type TeamAchievementsRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewTeamAchievementsRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *TeamAchievementsRepoFactory {
	return &TeamAchievementsRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *TeamAchievementsRepoFactory) Create(ctx context.Context, db pgclient.DB) *TeamAchievementsRepo {
	return &TeamAchievementsRepo{
		log: f.log.WithComponent(ctx, "team-achievements-repo"),
		db:  db,
	}
}

type TeamAchievementsRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *TeamAchievementsRepo) List(ctx context.Context) ([]*model.TeamAchievement, error) {
	ents, err := entity.TeamAchievements().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Team Achievements")
		return nil, err
	}

	res := make([]*model.TeamAchievement, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *TeamAchievementsRepo) One(ctx context.Context, id int64) (*model.TeamAchievement, error) {
	ent, err := entity.TeamAchievements(
		entity.TeamAchievementWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Team Achievement", zap.Int64("TeamAchievement.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *TeamAchievementsRepo) Create(ctx context.Context, m *model.TeamAchievement) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Team Achievement", zap.Error(err), zap.Any("TeamAchievement.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (r *TeamAchievementsRepo) SetTxHash(ctx context.Context, id int64, txHash string) error {
	ct, err := entity.TeamAchievements(
        entity.TeamAchievementWhere.ID.EQ(id),
    ).One(ctx, r.db)
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not get ent", zap.Error(err))
		return err
    }

    // 2. Меняем поле
    ct.TXHash = null.StringFrom(txHash)

    // 3. Обновляем только record_hash
    _, err = ct.Update(ctx, r.db, boil.Whitelist(entity.TeamAchievementColumns.TXHash))
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not update ent", zap.Error(err))
		return err
    }
    return nil
}

func (*TeamAchievementsRepo) toEntity(m *model.TeamAchievement) *entity.TeamAchievement {
	return &entity.TeamAchievement{
		ID:         m.ID,
		TeamID:     m.TeamID,
		PrizeID:    m.PrizeID,
		CreatedAt:  m.CreatedAt,
		TXHash:     null.StringFromPtr(m.TXHash),
	}
}

func (*TeamAchievementsRepo) fromEntity(e *entity.TeamAchievement) *model.TeamAchievement {
	return &model.TeamAchievement{
		ID:         e.ID,
		TeamID:     e.TeamID,
		PrizeID:    e.PrizeID,
		CreatedAt:  e.CreatedAt,
		TXHash:     e.TXHash.Ptr(),
	}
}
