package service

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"go.uber.org/zap"
	"math/big"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
	"sport-results-pocessor/internal/common/service/client/blockchain"
	"sport-results-pocessor/internal/common/service/server/registry"
	"sport-results-pocessor/internal/domain/model"
	"sport-results-pocessor/internal/infra/repo"
)

type SportRegistryService struct {
	ctx context.Context
	log logger.Logger
	db  pgclient.DB

	blockchainClient *blockchain.BlockchainRegistryClient

	competitionLevelRepo *repo.CompetitionLevelRepo
	competitionTeamsRepo *repo.CompetitionTeamRepo
	competitionRepo      *repo.CompetitionRepo
	countryRepo          *repo.CountryRepo
	currencyRepo         *repo.CurrencyRepo
	locationRepo         *repo.LocationRepo
	matchParticipantRepo *repo.MatchParticipantRepo
	matchRepo            *repo.MatchRepo
	personSportRepo      *repo.PersonSportRepo
	personRepo           *repo.PersonRepo
	prizeRepo            *repo.PrizeRepo
	roleRepo             *repo.RoleRepo
	sportRepo            *repo.SportRepo
	stageRepo            *repo.StageRepo
	teamAchievementsRepo *repo.TeamAchievementsRepo
	teamPersonRepo       *repo.TeamPersonRepo
	teamRepo             *repo.TeamRepo
}

func NewSportRegistryService(
	ctx context.Context,
	log logger.Logger,
	db pgclient.DB,
	competitionLevelRepoFactory *repo.CompetitionLevelRepoFactory,
	competitionTeamsRepoFactory *repo.CompetitionTeamsRepoFactory,
	competitionRepoFactory *repo.CompetitionRepoFactory,
	countryRepoFactory *repo.CountryRepoFactory,
	currencyRepoFactory *repo.CurrencyRepoFactory,
	locationRepoFactory *repo.LocationRepoFactory,
	matchParticipantRepoFactory *repo.MatchParticipantRepoFactory,
	matchRepoFactory *repo.MatchRepoFactory,
	personSportRepoFactory *repo.PersonSportRepoFactory,
	personRepoFactory *repo.PersonRepoFactory,
	prizeRepoFactory *repo.PrizeRepoFactory,
	roleRepoFactory *repo.RoleRepoFactory,
	sportRepoFactory *repo.SportRepoFactory,
	stageRepoFactory *repo.StageRepoFactory,
	teamAchievementsRepoFactory *repo.TeamAchievementsRepoFactory,
	teamPersonRepoFactory *repo.TeamPersonRepoFactory,
	teamRepoFactory *repo.TeamRepoFactory,
	blockchainClient *blockchain.BlockchainRegistryClient,
) *SportRegistryService {
	return &SportRegistryService{
		ctx:                  ctx,
		db:                   db,
		log:                  log.WithComponent(ctx, "sport-registry-server"),
		competitionLevelRepo: competitionLevelRepoFactory.Create(ctx, db),
		competitionTeamsRepo: competitionTeamsRepoFactory.Create(ctx, db),
		competitionRepo:      competitionRepoFactory.Create(ctx, db),
		countryRepo:          countryRepoFactory.Create(ctx, db),
		currencyRepo:         currencyRepoFactory.Create(ctx, db),
		locationRepo:         locationRepoFactory.Create(ctx, db),
		matchParticipantRepo: matchParticipantRepoFactory.Create(ctx, db),
		matchRepo:            matchRepoFactory.Create(ctx, db),
		personSportRepo:      personSportRepoFactory.Create(ctx, db),
		personRepo:           personRepoFactory.Create(ctx, db),
		prizeRepo:            prizeRepoFactory.Create(ctx, db),
		roleRepo:             roleRepoFactory.Create(ctx, db),
		sportRepo:            sportRepoFactory.Create(ctx, db),
		stageRepo:            stageRepoFactory.Create(ctx, db),
		teamAchievementsRepo: teamAchievementsRepoFactory.Create(ctx, db),
		teamPersonRepo:       teamPersonRepoFactory.Create(ctx, db),
		teamRepo:             teamRepoFactory.Create(ctx, db),
		blockchainClient:     blockchainClient,
	}
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {
	res, err := s.competitionLevelRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return &registry.CompetitionLevel{
		Id:        res.ID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) (*registry.ListCompetitionLevelsResponse, error) {
	res, err := s.competitionLevelRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.CompetitionLevel, 0, len(res))

	for i := range res {
		out = append(out, registry.CompetitionLevel{
			Id:        res[i].ID,
			Name:      res[i].Name,
			CreatedAt: res[i].CreatedAt,
		})
	}

	return &registry.ListCompetitionLevelsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) error {
	buf := &model.CompetitionLevel{
		Name: req.Name,
	}

	err := s.competitionLevelRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindCompetitionTeam(ctx context.Context, id int64) (*registry.CompetitionTeams, error) {
	res, err := s.competitionTeamsRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.CompetitionTeams.ValidateCompetitionTeams(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindCompetitionTeam").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindCompetitionTeam").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.CompetitionTeams{
		Id:            res.ID,
		CompetitionId: res.CompetitionID,
		TeamId:        res.TeamID,
		CreatedAt:     res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListCompetitionTeams(ctx context.Context) (*registry.ListCompetitionTeamsResponse, error) {
	res, err := s.competitionTeamsRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.CompetitionTeams, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.CompetitionTeams.ValidateCompetitionTeams(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListCompetitionTeams").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListCompetitionTeams").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.CompetitionTeams{
			Id:            res[i].ID,
			CompetitionId: res[i].CompetitionID,
			TeamId:        res[i].TeamID,
			CreatedAt:     res[i].CreatedAt,
		})
	}

	return &registry.ListCompetitionTeamsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddCompetitionTeam(ctx context.Context, req registry.NewCompetitionTeams) error {
	buf := &model.CompetitionTeam{
		TeamID:        req.TeamId,
		CompetitionID: req.CompetitionId,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.competitionTeamsRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.CompetitionTeams.RecordCompetitionTeams(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddCompetitionTeam").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.competitionRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	err = s.competitionTeamsRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindCompetition(ctx context.Context, id int64) (*registry.Competition, error) {
	res, err := s.competitionRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.Competition.ValidateCompetition(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindCompetition").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindCompetition").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.Competition{
		Id:         res.ID,
		LevelId:    res.LevelID,
		LocationId: res.LocationID,
		Name:       res.Name,
		SportId:    res.SportID,
		CreatedAt:  res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListCompetitions(ctx context.Context) (*registry.ListCompetitionsResponse, error) {
	res, err := s.competitionRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Competition, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.Competition.ValidateCompetition(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListCompetition").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListCompetition").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.Competition{
			Id:         res[i].ID,
			LevelId:    res[i].LevelID,
			LocationId: res[i].LocationID,
			Name:       res[i].Name,
			SportId:    res[i].SportID,
			CreatedAt:  res[i].CreatedAt,
		})
	}

	return &registry.ListCompetitionsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddCompetition(ctx context.Context, req registry.NewCompetition) error {
	buf := &model.Competition{
		LocationID: req.LocationId,
		LevelID:    req.LevelId,
		SportID:    req.SportId,
		Name:       req.Name,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.competitionRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.Competition.RecordCompetition(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddCompetition").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.competitionRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindCountry(ctx context.Context, id int64) (*registry.Country, error) {
	res, err := s.countryRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return &registry.Country{
		Id:        res.ID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListCountries(ctx context.Context) (*registry.ListCountriesResponse, error) {
	res, err := s.countryRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Country, 0, len(res))

	for i := range res {
		out = append(out, registry.Country{
			Id:        res[i].ID,
			Name:      res[i].Name,
			CreatedAt: res[i].CreatedAt,
		})
	}

	return &registry.ListCountriesResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddCountry(ctx context.Context, req registry.NewCountry) error {
	buf := &model.Country{
		Name: req.Name,
	}

	err := s.countryRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindCurrency(ctx context.Context, code string) (*registry.Currency, error) {
	res, err := s.currencyRepo.One(ctx, code)
	if err != nil {
		return nil, err
	}

	return &registry.Currency{
		Code: res.Code,
		Name: res.Name,
	}, nil
}

func (s *SportRegistryService) ListCurrencies(ctx context.Context) (*registry.ListCurrenciesResponse, error) {
	res, err := s.currencyRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Currency, 0, len(res))

	for i := range res {
		out = append(out, registry.Currency{
			Code: res[i].Code,
			Name: res[i].Name,
		})
	}

	return &registry.ListCurrenciesResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddCurrency(ctx context.Context, req registry.NewCurrency) error {
	buf := &model.Currency{
		Code: req.Code,
		Name: req.Name,
	}

	err := s.currencyRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindLocation(ctx context.Context, id int64) (*registry.Location, error) {
	res, err := s.locationRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return &registry.Location{
		Id:          res.ID,
		Address:     res.Address,
		City:        res.City,
		CountryId:   res.CountryID,
		CreatedAt:   &res.CreatedAt,
		FullAddress: *res.Address,
		State:       res.State,
	}, nil
}

func (s *SportRegistryService) ListLocations(ctx context.Context) (*registry.ListLocationsResponse, error) {
	res, err := s.locationRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Location, 0, len(res))

	for i := range res {
		out = append(out, registry.Location{
			Id:          res[i].ID,
			Address:     res[i].Address,
			City:        res[i].City,
			CountryId:   res[i].CountryID,
			CreatedAt:   &res[i].CreatedAt,
			FullAddress: *res[i].Address,
			State:       res[i].State,
		})
	}

	return &registry.ListLocationsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddLocation(ctx context.Context, req registry.NewLocation) error {
	buf := &model.Location{
		Address:     req.Address,
		City:        req.City,
		CountryID:   req.CountryId,
		FullAddress: *req.Address,
		State:       req.State,
	}

	err := s.locationRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindMatchParticipant(ctx context.Context, id int64) (*registry.MatchParticipant, error) {
	res, err := s.matchParticipantRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.MatchParticipant.ValidateMatchParticipant(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindMatchParticipant").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindMatchParticipant").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.MatchParticipant{
		Id:       res.ID,
		IsWinner: &res.IsWinner,
		MatchId:  res.MatchID,
		Score:    res.Score,
		TeamId:   res.TeamID,
	}, nil
}

func (s *SportRegistryService) ListMatchParticipants(ctx context.Context) (*registry.ListMatchParticipantsResponse, error) {
	res, err := s.matchParticipantRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.MatchParticipant, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.MatchParticipant.ValidateMatchParticipant(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListMatchParticipants").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListMatchParticipants").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.MatchParticipant{
			Id:       res[i].ID,
			IsWinner: &res[i].IsWinner,
			MatchId:  res[i].MatchID,
			Score:    res[i].Score,
			TeamId:   res[i].TeamID,
		})
	}

	return &registry.ListMatchParticipantsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddMatchParticipant(ctx context.Context, req registry.NewMatchParticipant) error {
	buf := &model.MatchParticipant{
		IsWinner: *req.IsWinner,
		MatchID:  req.MatchId,
		Score:    req.Score,
		TeamID:   req.TeamId,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.matchParticipantRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.Competition.RecordCompetition(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddMatchParticipant").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.matchParticipantRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindMatch(ctx context.Context, id int64) (*registry.Match, error) {
	res, err := s.matchRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.Match.ValidateMatch(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindMatch").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindMatch").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	m := make(map[string]interface{})
	if err := json.Unmarshal(res.Metadata, &m); err != nil {
		s.log.WithMethod(ctx, "FindMatch").Error("invalid object metadata", zap.ByteString("object", res.Metadata))
	}

	return &registry.Match{
		Id:         res.ID,
		LocationId: res.LocationID,
		MatchTime:  res.MatchTime,
		Metadata:   &m,
		StageId:    res.StageID,
		CreatedAt:  res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListMatches(ctx context.Context) (*registry.ListMatchesResponse, error) {
	res, err := s.matchRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Match, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.Match.ValidateMatch(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListMatches").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListMatches").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		m := make(map[string]interface{})
		if err := json.Unmarshal(res[i].Metadata, &m); err != nil {
			s.log.WithMethod(ctx, "ListMatches").Error("invalid object metadata", zap.ByteString("object", res[i].Metadata))
		}

		out = append(out, registry.Match{
			Id:         res[i].ID,
			LocationId: res[i].LocationID,
			MatchTime:  res[i].MatchTime,
			Metadata:   &m,
			StageId:    res[i].StageID,
			CreatedAt:  res[i].CreatedAt,
		})
	}

	return &registry.ListMatchesResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddMatch(ctx context.Context, req registry.Match) error {
	m, err := json.Marshal(req.Metadata)
	buf := &model.Match{
		LocationID: req.LocationId,
		MatchTime:  req.MatchTime,
		Metadata:   m,
		StageID:    req.StageId,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.matchRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.Match.RecordMatch(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddMatch").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.matchRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindPersonSports(ctx context.Context, id int64) (*registry.PersonSport, error) {
	res, err := s.personSportRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return &registry.PersonSport{
		Id:        res.ID,
		PersonId:  res.PersonID,
		SportId:   res.SportID,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListPersonSports(ctx context.Context) (*registry.ListPersonSportsResponse, error) {
	res, err := s.personSportRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.PersonSport, 0, len(res))

	for i := range res {
		out = append(out, registry.PersonSport{
			Id:        res[i].ID,
			PersonId:  res[i].PersonID,
			SportId:   res[i].SportID,
			CreatedAt: res[i].CreatedAt,
		})
	}

	return &registry.ListPersonSportsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddPersonSport(ctx context.Context, req registry.NewPersonSport) error {
	buf := &model.PersonSport{
		PersonID: req.PersonId,
		SportID:  req.SportId,
	}

	err := s.personSportRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindPerson(ctx context.Context, id int64) (*registry.Person, error) {
	res, err := s.personRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.Person.ValidatePerson(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindPerson").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindPerson").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.Person{
		Id:        res.ID,
		BirthDate: res.BirthDate,
		CountryId: res.CountryID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListPerson(ctx context.Context) (*registry.ListPersonsResponse, error) {
	res, err := s.personRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Person, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.Person.ValidatePerson(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListPerson").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListPerson").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.Person{
			Id:        res[i].ID,
			BirthDate: res[i].BirthDate,
			Name:      res[i].Name,
			CountryId: res[i].CountryID,
			CreatedAt: res[i].CreatedAt,
		})
	}

	return &registry.ListPersonsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddPerson(ctx context.Context, req registry.NewPerson) error {
	buf := &model.Person{
		Name:      req.Name,
		CountryID: req.CountryId,
		BirthDate: req.BirthDate,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.personRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.Person.RecordPerson(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddPerson").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.personRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindPrize(ctx context.Context, id int64) (*registry.Prize, error) {
	res, err := s.prizeRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.Prize.ValidatePrize(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindPrize").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindPrize").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.Prize{
		Id:            res.ID,
		CompetitionId: res.CompetitionID,
		CurrencyCode:  res.CurrencyCode,
		PlaceBracket:  res.PlaceBracket,
		PrizeAmount:   int(res.PrizeAmount),
		CreatedAt:     res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListPrises(ctx context.Context) (*registry.ListPrizesResponse, error) {
	res, err := s.prizeRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Prize, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.Prize.ValidatePrize(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListPrises").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListPrises").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.Prize{
			Id:            res[i].ID,
			CompetitionId: res[i].CompetitionID,
			CurrencyCode:  res[i].CurrencyCode,
			PlaceBracket:  res[i].PlaceBracket,
			PrizeAmount:   int(res[i].PrizeAmount),
			CreatedAt:     res[i].CreatedAt,
		})
	}

	return &registry.ListPrizesResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddPrize(ctx context.Context, req registry.NewPrize) error {
	buf := &model.Prize{
		CompetitionID: req.CompetitionId,
		CurrencyCode:  req.CurrencyCode,
		PlaceBracket:  req.PlaceBracket,
		PrizeAmount:   int64(req.PrizeAmount),
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.prizeRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.Prize.RecordPrize(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddPrize").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.prizeRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindRole(ctx context.Context, id int64) (*registry.Role, error) {
	res, err := s.roleRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return &registry.Role{
		Id:        res.ID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListRoles(ctx context.Context) (*registry.ListRolesResponse, error) {
	res, err := s.roleRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Role, 0, len(res))

	for i := range res {
		out = append(out, registry.Role{
			Id:        res[i].ID,
			Name:      res[i].Name,
			CreatedAt: res[i].CreatedAt,
		})
	}

	return &registry.ListRolesResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddRole(ctx context.Context, req registry.NewRole) error {
	buf := &model.Role{
		Name: req.Name,
	}

	err := s.roleRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindSport(ctx context.Context, id int64) (*registry.Sport, error) {
	res, err := s.sportRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.Sport.ValidateSport(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindSport").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindSport").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.Sport{
		Id:          res.ID,
		Description: res.Description,
		Name:        res.Name,
		MaxTeamSize: res.MaxTeamSize,
		MinTeamSize: res.MinTeamSize,
		CreatedAt:   res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListSports(ctx context.Context) (*registry.ListSportsResponse, error) {
	res, err := s.sportRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Sport, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.Sport.ValidateSport(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListSports").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListSports").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.Sport{
			Id:          res[i].ID,
			Description: res[i].Description,
			Name:        res[i].Name,
			MaxTeamSize: res[i].MaxTeamSize,
			MinTeamSize: res[i].MinTeamSize,
			CreatedAt:   res[i].CreatedAt,
		})
	}

	return &registry.ListSportsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddSport(ctx context.Context, req registry.NewSport) error {
	buf := &model.Sport{
		Description: req.Description,
		Name:        req.Name,
		MinTeamSize: req.MinTeamSize,
		MaxTeamSize: req.MaxTeamSize,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.sportRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.Sport.RecordSport(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddSport").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.sportRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindStage(ctx context.Context, id int64) (*registry.Stage, error) {
	res, err := s.stageRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return &registry.Stage{
		Id:            res.ID,
		Name:          res.Name,
		CompetitionId: res.CompetitionID,
		StartTime:     res.StartTime,
		EndTime:       res.EndTime,
		CreatedAt:     res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListStages(ctx context.Context) (*registry.ListStagesResponse, error) {
	res, err := s.stageRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Stage, 0, len(res))

	for i := range res {
		out = append(out, registry.Stage{
			Id:            res[i].ID,
			Name:          res[i].Name,
			CompetitionId: res[i].CompetitionID,
			StartTime:     res[i].StartTime,
			EndTime:       res[i].EndTime,
			CreatedAt:     res[i].CreatedAt,
		})
	}

	return &registry.ListStagesResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddStage(ctx context.Context, req registry.NewStage) error {
	buf := &model.Stage{
		Name:          req.Name,
		CompetitionID: req.CompetitionId,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
	}

	err := s.stageRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindTeamAchievement(ctx context.Context, id int64) (*registry.TeamAchievements, error) {
	res, err := s.teamAchievementsRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.TeamAchievements.ValidateTeamAchievements(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindTeamAchievement").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindTeamAchievement").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.TeamAchievements{
		Id:        res.ID,
		PrizeId:   res.PrizeID,
		TeamId:    res.TeamID,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListTeamAchievements(ctx context.Context) (*registry.ListTeamAchievementsResponse, error) {
	res, err := s.teamAchievementsRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.TeamAchievements, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.TeamAchievements.ValidateTeamAchievements(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListTeamAchievements").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListTeamAchievements").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.TeamAchievements{
			Id:        res[i].ID,
			PrizeId:   res[i].PrizeID,
			TeamId:    res[i].TeamID,
			CreatedAt: res[i].CreatedAt,
		})
	}

	return &registry.ListTeamAchievementsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddTeamAchievement(ctx context.Context, req registry.NewTeamAchievements) error {
	buf := &model.TeamAchievement{
		PrizeID: req.PrizeId,
		TeamID:  req.TeamId,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.teamAchievementsRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.TeamAchievements.RecordTeamAchievements(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddTeamAchievement").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.teamAchievementsRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindTeamPerson(ctx context.Context, id int64) (*registry.TeamPerson, error) {
	res, err := s.teamPersonRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.TeamPerson.ValidateTeamPerson(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindTeamPerson").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindTeamPerson").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.TeamPerson{
		Id:        res.ID,
		JoinedAt:  res.JoinedAt,
		LeftAt:    res.LeftAt,
		RoleId:    res.RoleID,
		PersonId:  res.PersonID,
		TeamId:    res.TeamID,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListTeamPersons(ctx context.Context) (*registry.ListTeamPersonsResponse, error) {
	res, err := s.teamPersonRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.TeamPerson, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.TeamPerson.ValidateTeamPerson(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListTeamPersons").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListTeamPersons").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.TeamPerson{
			Id:        res[i].ID,
			JoinedAt:  res[i].JoinedAt,
			LeftAt:    res[i].LeftAt,
			RoleId:    res[i].RoleID,
			PersonId:  res[i].PersonID,
			TeamId:    res[i].TeamID,
			CreatedAt: res[i].CreatedAt,
		})
	}

	return &registry.ListTeamPersonsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddTeamPerson(ctx context.Context, req registry.NewTeamPerson) error {
	buf := &model.TeamPerson{
		RoleID:   req.RoleId,
		TeamID:   req.TeamId,
		PersonID: req.PersonId,
		JoinedAt: req.JoinedAt,
		LeftAt:   req.LeftAt,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.teamPersonRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.TeamPerson.RecordTeamPerson(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddTeamPerson").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.teamPersonRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) FindTeam(ctx context.Context, id int64) (*registry.Team, error) {
	res, err := s.teamRepo.One(ctx, id)
	if err != nil {
		return nil, err
	}

	hash, err := s.calcHash(res)
	if err != nil {
		return nil, err
	}

	hashValid, err := s.blockchainClient.Team.ValidateTeam(s.blockchainClient.CallOpts, big.NewInt(res.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "FindTeam").Error("can no validate object", zap.Error(err))
		return nil, err
	}
	if !hashValid {
		s.log.WithMethod(ctx, "FindTeam").Error("invalid object detected", zap.Any("object", res))
		return nil, err
	}

	return &registry.Team{
		Id:             res.ID,
		CountryId:      res.CountryID,
		FoundationDate: res.FoundationDate,
		Name:           res.Name,
		CreatedAt:      res.CreatedAt,
	}, nil
}

func (s *SportRegistryService) ListTeams(ctx context.Context) (*registry.ListTeamsResponse, error) {
	res, err := s.teamRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]registry.Team, 0, len(res))

	for i := range res {

		hash, err := s.calcHash(res)
		if err != nil {
			return nil, err
		}

		hashValid, err := s.blockchainClient.Team.ValidateTeam(s.blockchainClient.CallOpts, big.NewInt(res[i].ID), hash)
		if err != nil {
			s.log.WithMethod(ctx, "ListTeams").Error("can no validate object", zap.Error(err))
			return nil, err
		}
		if !hashValid {
			s.log.WithMethod(ctx, "ListTeams").Error("invalid object detected", zap.Any("object", res[i]))
			return nil, err
		}

		out = append(out, registry.Team{
			Id:             res[i].ID,
			CountryId:      res[i].CountryID,
			FoundationDate: res[i].FoundationDate,
			Name:           res[i].Name,
			CreatedAt:      res[i].CreatedAt,
		})
	}

	return &registry.ListTeamsResponse{
		Data: out,
	}, nil
}

func (s *SportRegistryService) AddTeam(ctx context.Context, req registry.NewTeam) error {
	buf := &model.Team{
		Name:           req.Name,
		CountryID:      req.CountryId,
		FoundationDate: req.FoundationDate,
	}

	hash, err := s.calcHash(buf)
	if err != nil {
		return err
	}

	err = s.teamRepo.Create(ctx, buf)
	if err != nil {
		return err
	}

	tx, err := s.blockchainClient.Team.RecordTeam(s.blockchainClient.Auth, big.NewInt(buf.ID), hash)
	if err != nil {
		s.log.WithMethod(ctx, "AddTeam").Error("can not record object hash", zap.Error(err))
		return err
	}

	err = s.teamRepo.SetTxHash(ctx, buf.ID, tx.Hash().String())
	if err != nil {
		return err
	}

	return nil
}

func (s *SportRegistryService) calcHash(t any) (string, error) {
	out := ""
	json, err := json.Marshal(t)
	if err != nil {
		s.log.WithMethod(s.ctx, "calcHash").Error("can not marshal t", zap.Any("object", t))
		return out, err
	}

	buf := sha1.Sum(json)
	out = base64.RawURLEncoding.EncodeToString(buf[:])

	return out, nil
}
