package service

import (
	"context"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
	"sport-results-pocessor/internal/common/service/client/blockchain"
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
		teamAchievementsRepo: teamAchievementsRepoFactory.Create(ctx, db),
		teamPersonRepo:       teamPersonRepoFactory.Create(ctx, db),
		teamRepo:             teamRepoFactory.Create(ctx, db),
		blockchainClient:     blockchainClient,
	}
}
