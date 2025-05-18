package app

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sport-results-pocessor/internal/common/adapter/config"
	"sport-results-pocessor/internal/common/adapter/helper/util"
	"sport-results-pocessor/internal/common/adapter/http/middleware"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/service/server/registry"
	server "sport-results-pocessor/internal/common/service/server/registry"
	"sport-results-pocessor/internal/infra/service"
	"strconv"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func SportRegistryServer(
	lf fx.Lifecycle,
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	sportRegistryService *service.SportRegistryService,
) error {
	handler := registry.Handler(&sportRegistryServer{
		ctx:     ctx,
		log:     log,
		service: sportRegistryService,
	})

	handler = middleware.NewLoggingMiddleware(ctx, handler, log.WithComponent(ctx, "sport-registry-server"))
	handler = middleware.NewRecoverMiddleware(ctx, handler, log.WithComponent(ctx, "sport-registry-server"))
	auth := middleware.NewAuthMiddleware(cfg.Server.Registry.ValidationKey, log.WithComponent(ctx, "sport-registry-server"))
	handler = auth.Auth(handler)

	httpServer := &http.Server{
		Addr:              ":" + strconv.Itoa(cfg.Server.Registry.Port),
		Handler:           handler,
		IdleTimeout:       8 * time.Second,
		ReadHeaderTimeout: time.Second,
	}

	lf.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info("starting customer server", zap.String("address", httpServer.Addr))

			go func() {
				if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					if log != nil {
						log.Error("cannot to start customer server", zap.Error(err))
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := httpServer.Shutdown(ctx); err != nil {
				log.Error("cannot to shutdown customer server", zap.Error(err))
			} else {
				log.Info("customer server stopped")
			}
			return nil
		},
	})

	return nil
}

type sportRegistryServer struct {
	ctx     context.Context
	log     logger.Logger
	service *service.SportRegistryService
}

// GetCompetitionLevels implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitionLevels(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListCompetitionLevels(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCompetitionLevelsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitionLevelsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindCompetitionLevel(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCompetitionTeams implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitionTeams(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListCompetitionTeams(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCompetitionTeamsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitionTeamsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindCompetitionTeam(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCompetitions implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitions(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListCompetitions(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCompetitionsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitionsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindTeam(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCountries implements registry.ServerInterface.
func (s *sportRegistryServer) GetCountries(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListCountries(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCountriesId implements registry.ServerInterface.
func (s *sportRegistryServer) GetCountriesId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindCountry(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCurrencies implements registry.ServerInterface.
func (s *sportRegistryServer) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListCurrencies(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetCurrenciesId implements registry.ServerInterface.
func (s *sportRegistryServer) GetCurrenciesId(w http.ResponseWriter, r *http.Request, id string) {
	res, err := s.service.FindCurrency(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetLocations implements registry.ServerInterface.
func (s *sportRegistryServer) GetLocations(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListLocations(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetLocationsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetLocationsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindLocation(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetMatchParticipants implements registry.ServerInterface.
func (s *sportRegistryServer) GetMatchParticipants(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListMatchParticipants(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetMatchParticipantsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetMatchParticipantsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindMatchParticipant(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetMatches implements registry.ServerInterface.
func (s *sportRegistryServer) GetMatches(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListMatches(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetMatchesId implements registry.ServerInterface.
func (s *sportRegistryServer) GetMatchesId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindMatch(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPersons implements registry.ServerInterface.
func (s *sportRegistryServer) GetPersons(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListPerson(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPersonsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetPersonsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindPerson(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPersonsSport implements registry.ServerInterface.
func (s *sportRegistryServer) GetPersonsSport(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListPersonSports(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPersonsSportId implements registry.ServerInterface.
func (s *sportRegistryServer) GetPersonsSportId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindPersonSports(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPersonsTeam implements registry.ServerInterface.
func (s *sportRegistryServer) GetPersonsTeam(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListTeamPersons(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPersonsTeamsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetPersonsTeamsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindTeamPerson(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPrizes implements registry.ServerInterface.
func (s *sportRegistryServer) GetPrizes(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListPrises(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetPrizesId implements registry.ServerInterface.
func (s *sportRegistryServer) GetPrizesId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindPrize(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetRoles implements registry.ServerInterface.
func (s *sportRegistryServer) GetRoles(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListRoles(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetRolesId implements registry.ServerInterface.
func (s *sportRegistryServer) GetRolesId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindRole(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetSports implements registry.ServerInterface.
func (s *sportRegistryServer) GetSports(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListSports(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetSportsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetSportsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindSport(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetStages implements registry.ServerInterface.
func (s *sportRegistryServer) GetStages(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListStages(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetStagesId implements registry.ServerInterface.
func (s *sportRegistryServer) GetStagesId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindStage(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetTeamAchievements implements registry.ServerInterface.
func (s *sportRegistryServer) GetTeamAchievements(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListTeamAchievements(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetTeamAchievementsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetTeamAchievementsId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindTeamAchievement(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetTeamId implements registry.ServerInterface.
func (s *sportRegistryServer) GetTeamId(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.service.FindTeam(s.ctx, id)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// GetTeams implements registry.ServerInterface.
func (s *sportRegistryServer) GetTeams(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.ListTeams(s.ctx)
	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, res, http.StatusOK)
}

// PostCompetitionLevels implements registry.ServerInterface.
func (s *sportRegistryServer) PostCompetitionLevels(w http.ResponseWriter, r *http.Request) {
	request := &server.NewCompetitionLevel{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddCompetitionLevel(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostCompetitionTeams implements registry.ServerInterface.
func (s *sportRegistryServer) PostCompetitionTeams(w http.ResponseWriter, r *http.Request) {
	request := &server.NewCompetitionTeams{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddCompetitionTeam(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostCompetitions implements registry.ServerInterface.
func (s *sportRegistryServer) PostCompetitions(w http.ResponseWriter, r *http.Request) {
	request := &server.NewCompetition{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddCompetition(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostCountries implements registry.ServerInterface.
func (s *sportRegistryServer) PostCountries(w http.ResponseWriter, r *http.Request) {
	request := &server.NewCountry{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddCountry(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostCurrencies implements registry.ServerInterface.
func (s *sportRegistryServer) PostCurrencies(w http.ResponseWriter, r *http.Request) {
	request := &server.NewCurrency{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddCurrency(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostLocations implements registry.ServerInterface.
func (s *sportRegistryServer) PostLocations(w http.ResponseWriter, r *http.Request) {
	request := &server.NewLocation{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddLocation(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostMatchParticipants implements registry.ServerInterface.
func (s *sportRegistryServer) PostMatchParticipants(w http.ResponseWriter, r *http.Request) {
	request := &server.NewMatchParticipant{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddMatchParticipant(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostMatches implements registry.ServerInterface.
func (s *sportRegistryServer) PostMatches(w http.ResponseWriter, r *http.Request) {
	request := &server.NewMatch{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddMatch(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostPersons implements registry.ServerInterface.
func (s *sportRegistryServer) PostPersons(w http.ResponseWriter, r *http.Request) {
	request := &server.NewPerson{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddPerson(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostPersonsSport implements registry.ServerInterface.
func (s *sportRegistryServer) PostPersonsSport(w http.ResponseWriter, r *http.Request) {
	request := &server.NewPersonSport{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddPersonSport(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostPersonsTeam implements registry.ServerInterface.
func (s *sportRegistryServer) PostPersonsTeam(w http.ResponseWriter, r *http.Request) {
	request := &server.NewTeamPerson{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddTeamPerson(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostPrizes implements registry.ServerInterface.
func (s *sportRegistryServer) PostPrizes(w http.ResponseWriter, r *http.Request) {
	request := &server.NewPrize{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddPrize(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostRoles implements registry.ServerInterface.
func (s *sportRegistryServer) PostRoles(w http.ResponseWriter, r *http.Request) {
	request := &server.NewRole{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddRole(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostSports implements registry.ServerInterface.
func (s *sportRegistryServer) PostSports(w http.ResponseWriter, r *http.Request) {
	request := &server.NewSport{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddSport(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostStages implements registry.ServerInterface.
func (s *sportRegistryServer) PostStages(w http.ResponseWriter, r *http.Request) {
	request := &server.NewStage{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddStage(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostTeamAchievements implements registry.ServerInterface.
func (s *sportRegistryServer) PostTeamAchievements(w http.ResponseWriter, r *http.Request) {
	request := &server.NewTeamAchievements{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddTeamAchievement(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

// PostTeams implements registry.ServerInterface.
func (s *sportRegistryServer) PostTeams(w http.ResponseWriter, r *http.Request) {
	request := &server.NewTeam{}

	if err := util.Unmarshal(s.ctx, s.log, request, r.Body); err != nil {
		s.log.WithMethod(s.ctx, "PostCompetitionLevels").Error("cannot unmarshal request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.AddTeam(s.ctx, request)

	if err != nil {
		s.writeResponse(w, err, http.StatusInternalServerError)
	}

	s.writeResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

func (s *sportRegistryServer) writeResponse(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")

	bytes, err := json.Marshal(response)
	if err != nil {
		s.log.WithMethod(s.ctx, "writeResponse").Error("cannot marshal response", zap.Error(err), zap.Any("debug.response.object", response))
		bytes = []byte(`{"message":"Internal Server Error"}`)
		statusCode = http.StatusInternalServerError
	}

	_, err = w.Write(bytes)
	if err != nil {
		s.log.WithMethod(s.ctx, "writeResponse").Error("cannot write response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
