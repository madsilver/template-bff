package main

import (
	"github.com/madsilver/template-bff/internal/adapter/graph"
	"github.com/madsilver/template-bff/internal/infra/http"
	"github.com/madsilver/template-bff/internal/infra/server"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	logger.Info("running template-bff")

	client := http.NewHttpClient()
	resolver := graph.NewResolver(client)
	graphServer := server.NewGraphqlServer(resolver)

	err := <-server.Start(graphServer)
	if err == nil {
		logger.Warn("shutdown")
	} else {
		logger.Panic("app down")
	}
}
