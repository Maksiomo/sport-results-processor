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

type CurrencyRepoFactory struct {
	ctx context.Context
	log logger.Logger
}

func NewCurrencyRepoFactory(
	ctx context.Context,
	log logger.Logger,
) *CurrencyRepoFactory {
	return &CurrencyRepoFactory{
		ctx: ctx,
		log: log,
	}
}

func (f *CurrencyRepoFactory) Create(ctx context.Context, db pgclient.DB) *CurrencyRepo {
	return &CurrencyRepo{
		log: f.log.WithComponent(ctx, "currency-repo"),
		db:  db,
	}
}

type CurrencyRepo struct {
	log logger.Logger
	db  pgclient.DB
}

func (r *CurrencyRepo) List(ctx context.Context) ([]*model.Currency, error) {
	ents, err := entity.Currencies().All(ctx, r.db)

	if err != nil {
		r.log.WithMethod(ctx, "List").Error("cannot list Currencies")
		return nil, err
	}

	res := make([]*model.Currency, 0, len(ents))
	for i := range ents {
		res = append(res, r.fromEntity(ents[i]))
	}

	return res, nil
}

func (r *CurrencyRepo) One(ctx context.Context, code string) (*model.Currency, error) {
	ent, err := entity.Currencies(
		entity.CurrencyWhere.Code.EQ(code),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.WithMethod(ctx, "Find").Error("cannot find Currency", zap.String("Currency.code", code))
		return nil, err
	}

	return r.fromEntity(ent), nil
}

func (r *CurrencyRepo) Create(ctx context.Context, m *model.Currency) error {
	ent := r.toEntity(m)
	if err := ent.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.log.WithMethod(ctx, "Create").Error("cannot insert Currency", zap.Error(err), zap.Any("Currency.object", m))
		return err
	}

	return nil
}

func (*CurrencyRepo) toEntity(m *model.Currency) *entity.Currency {
	return &entity.Currency{
		Code:   m.Code,
		Name:   m.Name,
		Symbol: null.StringFromPtr(m.Symbol),
	}
}

func (*CurrencyRepo) fromEntity(e *entity.Currency) *model.Currency {
	return &model.Currency{
		Code:   e.Code,
		Name:   e.Name,
		Symbol: e.Symbol.Ptr(),
	}
}
