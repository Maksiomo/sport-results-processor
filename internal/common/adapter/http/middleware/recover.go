package middleware

import (
	"context"
	"net/http"
	"sport-results-pocessor/internal/common/adapter/logger"

	"go.uber.org/zap"
)

type recoverMiddleware struct {
	ctx     context.Context
	handler http.Handler
	log     logger.Logger
}

func NewRecoverMiddleware(ctx context.Context, handler http.Handler, log logger.Logger) http.Handler {
	return &recoverMiddleware{
		ctx:     ctx,
		handler: handler,
		log:     log,
	}
}

func (m *recoverMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				m.log.WithMethod(m.ctx, "ServeHTTP").Error("panic recovered", zap.Error(r.(error)))
			default:
				m.log.WithMethod(m.ctx, "ServeHTTP").Error("panic recovered", zap.Any("panic", r))
			}

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()

	m.handler.ServeHTTP(w, r)
}
