package infra

import (
	"sport-results-pocessor/internal/infra/repo"
	"sport-results-pocessor/internal/infra/service"

	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	repo.NewCompetitionLevelRepoFactory,
	repo.NewCompetitionTeamsRepoFactory,
	repo.NewCompetitionRepoFactory,
	repo.NewCountryRepoFactory,
	repo.NewCurrencyRepoFactory,
	repo.NewLocationRepoFactory,
	repo.NewMatchParticipantRepoFactory,
	repo.NewMatchRepoFactory,
	repo.NewPersonSportRepoFactory,
	repo.NewPersonRepoFactory,
	repo.NewPrizeRepoFactory,
	repo.NewRoleRepoFactory,
	repo.NewSportRepoFactory,
	repo.NewTeamAchievementsRepoFactory,
	repo.NewTeamPersonRepoFactory,
	repo.NewTeamRepoFactory,

	service.NewSportRegistryService,
)
