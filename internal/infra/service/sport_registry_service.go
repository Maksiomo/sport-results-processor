package service

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
	"sport-results-pocessor/internal/common/service/client/blockchain"
	"sport-results-pocessor/internal/common/service/server/registry"
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

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {

}

func (s *SportRegistryService) FindCompetitionTeam(ctx context.Context, id int64) (*registry.CompetitionTeams, error) {

}
 
func (s *SportRegistryService) ListCompetitionTeams(ctx context.Context) ([]*registry.CompetitionTeams, error) {

}

func (s *SportRegistryService) AddCompetitionTeam(ctx context.Context, req registry.NewCompetitionTeams) (error) {
	
}

func (s *SportRegistryService) FindCompetition(ctx context.Context, id int64) (*registry.Competition, error) {

}
 
func (s *SportRegistryService) ListCompetitions(ctx context.Context) ([]*registry.Competition, error) {

}

func (s *SportRegistryService) AddCompetition(ctx context.Context, req registry.NewCompetition) (error) {
	
}

func (s *SportRegistryService) FindCountry(ctx context.Context, id int64) (*registry.Country, error) {

}
 
func (s *SportRegistryService) ListCountries(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCountry(ctx context.Context, req registry.NewCountry) (error) {
	
}

func (s *SportRegistryService) FindLocation(ctx context.Context, id int64) (*registry.Location, error) {

}
 
func (s *SportRegistryService) ListLocations(ctx context.Context) ([]*registry.Location, error) {

}

func (s *SportRegistryService) AddLocation(ctx context.Context, req registry.NewLocation) (error) {
	
}

func (s *SportRegistryService) FindMatchParticipant(ctx context.Context, id int64) (*registry.MatchParticipant, error) {

}
 
func (s *SportRegistryService) ListMatchParticipants(ctx context.Context) ([]*registry.MatchParticipant, error) {

}

func (s *SportRegistryService) AddMatchParticipant(ctx context.Context, req registry.NewMatchParticipant) (error) {
	
}

func (s *SportRegistryService) FindMatch(ctx context.Context, id int64) (*registry.Match, error) {

}
 
func (s *SportRegistryService) ListMatches(ctx context.Context) ([]*registry.Match, error) {

}

func (s *SportRegistryService) AddMatch(ctx context.Context, req registry.Match) (error) {
	
}

func (s *SportRegistryService) FindPeopleSports(ctx context.Context, id int64) (*registry.PeopleSports, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}

func (s *SportRegistryService) FindCompetitionLevel(ctx context.Context, id int64) (*registry.CompetitionLevel, error) {

}
 
func (s *SportRegistryService) ListCompetitionLevels(ctx context.Context) ([]*registry.CompetitionLevel, error) {

}

func (s *SportRegistryService) AddCompetitionLevel(ctx context.Context, req registry.NewCompetitionLevel) (error) {
	
}
