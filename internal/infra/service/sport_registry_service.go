package service

import (
	"context"
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
	countryRepo          *repo.CompetitionRepo
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
	countryRepoFactory *repo.CompetitionRepoFactory,
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

}

func (s *SportRegistryService) ListCompetitionTeams(ctx context.Context) (*registry.ListCompetitionTeamsResponse, error) {

}

func (s *SportRegistryService) AddCompetitionTeam(ctx context.Context, req registry.NewCompetitionTeams) error {

}

func (s *SportRegistryService) FindCompetition(ctx context.Context, id int64) (*registry.Competition, error) {

}

func (s *SportRegistryService) ListCompetitions(ctx context.Context) (*registry.ListCompetitionsResponse, error) {

}

func (s *SportRegistryService) AddCompetition(ctx context.Context, req registry.NewCompetition) error {

}

func (s *SportRegistryService) FindCountry(ctx context.Context, id int64) (*registry.Country, error) {

}

func (s *SportRegistryService) ListCountries(ctx context.Context) (*registry.ListCountriesResponse, error) {

}

func (s *SportRegistryService) AddCountry(ctx context.Context, req registry.NewCountry) error {

}

func (s *SportRegistryService) FindCurrency(ctx context.Context, code string) (*registry.Currency, error) {

}

func (s *SportRegistryService) ListCurrencies(ctx context.Context) (*registry.ListCurrenciesResponse, error) {

}

func (s *SportRegistryService) AddCurrency(ctx context.Context, req registry.NewCurrency) error {

}

func (s *SportRegistryService) FindLocation(ctx context.Context, id int64) (*registry.Location, error) {

}

func (s *SportRegistryService) ListLocations(ctx context.Context) (*registry.ListLocationsResponse, error) {

}

func (s *SportRegistryService) AddLocation(ctx context.Context, req registry.NewLocation) error {

}

func (s *SportRegistryService) FindMatchParticipant(ctx context.Context, id int64) (*registry.MatchParticipant, error) {

}

func (s *SportRegistryService) ListMatchParticipants(ctx context.Context) (*registry.ListMatchParticipantsResponse, error) {

}

func (s *SportRegistryService) AddMatchParticipant(ctx context.Context, req registry.NewMatchParticipant) error {

}

func (s *SportRegistryService) FindMatch(ctx context.Context, id int64) (*registry.Match, error) {

}

func (s *SportRegistryService) ListMatches(ctx context.Context) (*registry.ListMatchParticipantsResponse, error) {

}

func (s *SportRegistryService) AddMatch(ctx context.Context, req registry.Match) error {

}

func (s *SportRegistryService) FindPersonSports(ctx context.Context, id int64) (*registry.PersonSport, error) {

}

func (s *SportRegistryService) ListPersonSports(ctx context.Context) (*registry.ListPersonSportsResponse, error) {

}

func (s *SportRegistryService) AddPersonSport(ctx context.Context, req registry.NewPersonSport) error {

}

func (s *SportRegistryService) FindPerson(ctx context.Context, id int64) (*registry.Person, error) {

}

func (s *SportRegistryService) ListPerson(ctx context.Context) (*registry.ListPersonsResponse, error) {

}

func (s *SportRegistryService) AddPerson(ctx context.Context, req registry.NewPerson) error {

}

func (s *SportRegistryService) FindPrize(ctx context.Context, id int64) (*registry.Prize, error) {

}

func (s *SportRegistryService) ListPrises(ctx context.Context) (*registry.ListPrizesResponse, error) {

}

func (s *SportRegistryService) AddPrize(ctx context.Context, req registry.NewPrize) error {

}

func (s *SportRegistryService) FindRole(ctx context.Context, id int64) (*registry.Role, error) {

}

func (s *SportRegistryService) ListRoles(ctx context.Context) ([]*registry.ListRolesResponse, error) {

}

func (s *SportRegistryService) AddRole(ctx context.Context, req registry.NewRole) error {

}

func (s *SportRegistryService) FindSport(ctx context.Context, id int64) (*registry.Sport, error) {

}

func (s *SportRegistryService) ListSports(ctx context.Context) (*registry.ListSportsResponse, error) {

}

func (s *SportRegistryService) AddSport(ctx context.Context, req registry.NewSport) error {

}

func (s *SportRegistryService) FindStage(ctx context.Context, id int64) (*registry.Stage, error) {

}

func (s *SportRegistryService) ListStages(ctx context.Context) ([]*registry.ListStagesResponse, error) {

}

func (s *SportRegistryService) AddStage(ctx context.Context, req registry.NewStage) error {

}

func (s *SportRegistryService) FindTeamAchievement(ctx context.Context, id int64) (*registry.TeamAchievements, error) {

}

func (s *SportRegistryService) ListTeamAchievements(ctx context.Context) (*registry.ListTeamAchievementsResponse, error) {

}

func (s *SportRegistryService) AddTeamAchievement(ctx context.Context, req registry.NewTeamAchievements) error {

}

func (s *SportRegistryService) FindTeamPerson(ctx context.Context, id int64) (*registry.TeamPerson, error) {

}

func (s *SportRegistryService) ListTeamPersons(ctx context.Context) (*registry.ListTeamPersonsResponse, error) {

}

func (s *SportRegistryService) AddTeamPerson(ctx context.Context, req registry.NewTeamPerson) error {

}

func (s *SportRegistryService) FindTeam(ctx context.Context, id int64) (*registry.Team, error) {

}

func (s *SportRegistryService) ListTeams(ctx context.Context) ([]*registry.ListTeamsResponse, error) {

}

func (s *SportRegistryService) AddTeam(ctx context.Context, req registry.NewTeam) error {

}
