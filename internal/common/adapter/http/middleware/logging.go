package middleware

import (
	"context"
	"net/http"
	"net/http/httputil"
	"sport-results-pocessor/internal/common/adapter/logger"
	"strings"
	"time"

	"go.uber.org/zap"
)

type loggingMiddleware struct {
	ctx     context.Context
	handler http.Handler
	log     logger.Logger
}

func NewLoggingMiddleware(ctx context.Context, handler http.Handler, log logger.Logger) http.Handler {
	return &loggingMiddleware{
		ctx:     ctx,
		handler: handler,
		log:     log,
	}
}

func (m *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := m.log.WithMethod(m.ctx, "ServeHTTP")
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Error("failed to dump incoming request", zap.Error(err))
	}

	lw := &loggingResponseWriter{rw: w}

	start := time.Now()
	m.handler.ServeHTTP(lw, r)

	if err != nil {
		log.Error("failed to dump outgoing response", zap.Error(err))
	}

	var h strings.Builder
	for name, values := range lw.Header() {
		for _, value := range values {
			h.WriteString(name)
			h.WriteRune(':')
			h.WriteString(value)
			h.WriteRune('\n')
		}
	}

	h.WriteString("\r\n")
	if strings.Contains(r.RequestURI, "games/get-games") {
		h.Write([]byte(`{"body":"skipped"}`))
	} else {
		h.Write(lw.body)
	}

	log.Info("request served",
		zap.ByteString("debug.request.data", reqDump),
		zap.String("debug.response.data", h.String()),
		zap.Int("debug.response.status_code", lw.statusCode),
		zap.Duration("debug.request.duration", time.Since(start)),
	)
}

type loggingResponseWriter struct {
	rw         http.ResponseWriter
	statusCode int
	body       []byte
}

func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.rw.WriteHeader(code)
}

func (lw *loggingResponseWriter) Write(body []byte) (int, error) {
	lw.body = body
	return lw.rw.Write(body)
}

func (lw *loggingResponseWriter) Header() http.Header {
	return lw.rw.Header()
}
