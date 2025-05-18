package middleware

import (
	"encoding/json"
	"net/http"
	"sport-results-pocessor/internal/common/adapter/logger"

	"github.com/friendsofgo/errors"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	secretKey string
	log       logger.Logger
}

func NewAuthMiddleware(key string, log logger.Logger) AuthMiddleware {
	return AuthMiddleware{
		secretKey: key,
		log:       log,
	}
}

func (a *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		onTimeRW := oneTimeResponseWriter{ResponseWriter: w}
		a.auth(&onTimeRW, r)
		if onTimeRW.written {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *AuthMiddleware) auth(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("Sport-Registry-ApiKey")
	if key == "" {
		msg := "Sport-Registry-ApiKey is not provided"
		a.log.Error(msg)
		a.writeError(w, msg)
		return
	}

	if a.secretKey != key {
		msg := "Sport-Registry-ApiKey is not valid"
		a.log.Error(msg)
		a.writeError(w, msg)
		return
	}
}

func (a *AuthMiddleware) writeError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	err := errors.New(message)
	if err := json.NewEncoder(w).Encode(err); err != nil {
		a.log.Error("failed to encode err body", zap.Error(err))
	}
}

type oneTimeResponseWriter struct {
	http.ResponseWriter
	written bool
}

func (w *oneTimeResponseWriter) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *oneTimeResponseWriter) Write(body []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(body)
}
