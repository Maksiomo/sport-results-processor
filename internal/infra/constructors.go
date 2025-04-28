package infra

import (
	"sport-results-pocessor/internal/infra/repo"

	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	repo.NewCompetitionLevelRepoFactory,
	repo.NewCompetitionTeamsRepoFactory,
	repo.NewCompetitionRepoFactory,
	repo.NewCompetitionRepoFactory,
	repo.NewCountryRepoFactory,
	repo.NewCurrencyRepoFactory,
	repo.NewLocationRepoFactory,
	repo.NewMatchParticipantRepoFactory,
	repo.NewMatchRepoFactory,
	repo.NewPersonSportRepoFactory,
	repo.NewPersonRepoFactory,
	repo.NewPrizeRepoRepoFactory,
	repo.NewRoleRepoFactory,
	repo.NewSportRepoFactory,
	repo.NewSportRepoFactory,
	repo.NewTeamAchievementsRepoFactory,
	repo.NewTeamPersonRepoFactory,
	repo.NewTeamRepoFactory,
)