package main

import (
	"context"
	"sport-results-pocessor/internal/common/adapter"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/service"
	"sport-results-pocessor/internal/infra"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap/zapcore"
)

func main() {
	fxApp := fx.New(
		fx.Provide(context.Background),
		fx.WithLogger(func(ctx context.Context, log logger.Logger) fxevent.Logger {
			l := &fxevent.ZapLogger{Logger: log.WithComponent(ctx, "fx").UnZap()}
			l.UseLogLevel(zapcore.DebugLevel)
			return l
		}),
		fx.StartTimeout(time.Second*10),
		fx.StopTimeout(time.Second*10),

		adapter.Constructors,
		infra.Constructors,
		service.Constructors,
		fx.Invoke(
			
		),
	)

	fxApp.Run()
}