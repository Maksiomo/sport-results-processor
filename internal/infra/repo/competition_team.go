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

type CompetitionTeamsRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCompetitionTeamsRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CompetitionTeamsRepoFactory {
	return &CompetitionTeamsRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CompetitionTeamsRepoFactory) Create(ctx context.Context, db pgclient.DB) *CompetitionTeamRepo {
	return &CompetitionTeamRepo{
		log: f.log.WithComponent(ctx, "competition-teams-repo"),
		db:  db,
	}
}

type CompetitionTeamRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *CompetitionTeamRepo) List(ctx context.Context) ([]*model.CompetitionTeam, error) {
	ents, err := entity.CompetitionTeams().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list competition teams")
		return nil, err
	}

	res := make([]*model.CompetitionTeam, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *CompetitionTeamRepo) One(ctx context.Context, id int64) (*model.CompetitionTeam, error) {
	ent, err := entity.CompetitionTeams(
		entity.CompetitionTeamWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "One").Error("cannot find competition team", zap.Int64("competition_team.id", id))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *CompetitionTeamRepo) Create(ctx context.Context, m *model.CompetitionTeam) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert competition level", zap.Error(err), zap.Any("competition_team.object", m))
		return err
	}

	m.ID = ent.ID
	m.CreatedAt = ent.CreatedAt

	return nil
}

func (r *CompetitionTeamRepo) SetTxHash(ctx context.Context, id int64, txHash string) error {
	ct, err := entity.CompetitionTeams(
        entity.CompetitionTeamWhere.ID.EQ(id),
    ).One(ctx, r.db)
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not get ent", zap.Error(err))
		return err
    }

    // 2. Меняем поле
    ct.RecordHash = null.StringFrom(txHash)

    // 3. Обновляем только record_hash
    _, err = ct.Update(ctx, r.db, boil.Whitelist(entity.CompetitionTeamColumns.RecordHash))
    if err != nil {
        r.log.WithMethod(ctx, "SetTxHash").Error("can not update ent", zap.Error(err))
		return err
    }
    return nil
}

func (*CompetitionTeamRepo) toEntity(m *model.CompetitionTeam) *entity.CompetitionTeam {
	return &entity.CompetitionTeam{
		ID:            m.ID,
		TeamID:        m.TeamID,
		CompetitionID: m.CompetitionID,
		CreatedAt:     m.CreatedAt,
		RecordHash:    null.StringFrom(m.RecordHash),
		TXHash:        null.StringFromPtr(m.TXHash),
	}
}

func (*CompetitionTeamRepo) fromEntity(e *entity.CompetitionTeam) *model.CompetitionTeam {
	return &model.CompetitionTeam{
		ID:            e.ID,
		TeamID:        e.TeamID,
		CompetitionID: e.CompetitionID,
		CreatedAt:     e.CreatedAt,
		RecordHash:    e.RecordHash.String,
		TXHash:        &e.TXHash.String,
	}
}
