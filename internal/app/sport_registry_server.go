package app

import (
	"context"
	"errors"
	"net/http"
	"sport-results-pocessor/internal/common/adapter/config"
	"sport-results-pocessor/internal/common/adapter/http/middleware"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/service/server/registry"
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
		ctx:           ctx,
		log:           log,
		validationKey: cfg.Server.Registry.ValidationKey,
		service:       sportRegistryService,
	})

	handler = middleware.NewLoggingMiddleware(ctx, handler, log.WithComponent(ctx, "sport-registry-server"))
	handler = middleware.NewRecoverMiddleware(ctx, handler, log.WithComponent(ctx, "sport-registry-server"))

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
	ctx           context.Context
	log           logger.Logger
	validationKey string
	service       *service.SportRegistryService
}

// GetCompetitionLevels implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitionLevels(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetCompetitionTeams implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitionTeams(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetCompetitions implements registry.ServerInterface.
func (s *sportRegistryServer) GetCompetitions(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetCountries implements registry.ServerInterface.
func (s *sportRegistryServer) GetCountries(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetLanguages implements registry.ServerInterface.
func (s *sportRegistryServer) GetLanguages(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetLocations implements registry.ServerInterface.
func (s *sportRegistryServer) GetLocations(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetMatchParticipants implements registry.ServerInterface.
func (s *sportRegistryServer) GetMatchParticipants(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetMatches implements registry.ServerInterface.
func (s *sportRegistryServer) GetMatches(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetPersons implements registry.ServerInterface.
func (s *sportRegistryServer) GetPersons(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetPrizes implements registry.ServerInterface.
func (s *sportRegistryServer) GetPrizes(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetRoles implements registry.ServerInterface.
func (s *sportRegistryServer) GetRoles(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetSports implements registry.ServerInterface.
func (s *sportRegistryServer) GetSports(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetSportsId implements registry.ServerInterface.
func (s *sportRegistryServer) GetSportsId(w http.ResponseWriter, r *http.Request, id int64) {
	panic("unimplemented")
}

// GetStages implements registry.ServerInterface.
func (s *sportRegistryServer) GetStages(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetTeamAchievements implements registry.ServerInterface.
func (s *sportRegistryServer) GetTeamAchievements(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetTeams implements registry.ServerInterface.
func (s *sportRegistryServer) GetTeams(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostCompetitionLevels implements registry.ServerInterface.
func (s *sportRegistryServer) PostCompetitionLevels(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostCompetitionTeams implements registry.ServerInterface.
func (s *sportRegistryServer) PostCompetitionTeams(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostCompetitions implements registry.ServerInterface.
func (s *sportRegistryServer) PostCompetitions(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostCountries implements registry.ServerInterface.
func (s *sportRegistryServer) PostCountries(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostLanguages implements registry.ServerInterface.
func (s *sportRegistryServer) PostLanguages(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostLocations implements registry.ServerInterface.
func (s *sportRegistryServer) PostLocations(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostMatchParticipants implements registry.ServerInterface.
func (s *sportRegistryServer) PostMatchParticipants(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostMatches implements registry.ServerInterface.
func (s *sportRegistryServer) PostMatches(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostPersons implements registry.ServerInterface.
func (s *sportRegistryServer) PostPersons(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostPrizes implements registry.ServerInterface.
func (s *sportRegistryServer) PostPrizes(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostRoles implements registry.ServerInterface.
func (s *sportRegistryServer) PostRoles(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostSports implements registry.ServerInterface.
func (s *sportRegistryServer) PostSports(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostStages implements registry.ServerInterface.
func (s *sportRegistryServer) PostStages(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostTeamAchievements implements registry.ServerInterface.
func (s *sportRegistryServer) PostTeamAchievements(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PostTeams implements registry.ServerInterface.
func (s *sportRegistryServer) PostTeams(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (s *sportRegistryServer) validateKey(key string) error {
	if key != s.validationKey {
		s.log.WithMethod(s.ctx, "validateKey").Error("authenticate failed")
		return errors.New("authenticate failed")
	}

	return nil
}