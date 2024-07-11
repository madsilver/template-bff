package main

import (
	"context"

	"github.com/madsilver/template-bff/internal/adapter/graph"
	"github.com/madsilver/template-bff/internal/infra/http"
	"github.com/madsilver/template-bff/internal/infra/otel"
	"github.com/madsilver/template-bff/internal/infra/server"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	logger.Info("running template-bff")

	client := http.NewHttpClient()
	resolver := graph.NewResolver(client, logger)
	graphServer := server.NewGraphqlServer(resolver)

	initOtel(logger)

	err := <-server.Start(graphServer)
	if err == nil {
		logger.Warn("shutdown")
	} else {
		logger.Panic("app down")
	}
}

func initOtel(logger *zap.Logger) {
	ctx := context.Background()

	// tracer := otel.InitTracer(ctx, logger)
	// defer func() {
	// 	if err := tracer.Shutdown(ctx); err != nil {
	// 		logger.Error(err.Error())
	// 	}
	// }()

	meter := otel.InitMeter(ctx, logger)
	defer func() {
		if err := meter.Shutdown(ctx); err != nil {
			logger.Info(err.Error())
		}
	}()
}
