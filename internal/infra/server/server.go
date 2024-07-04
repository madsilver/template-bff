package server

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/madsilver/template-bff/internal/adapter/handler"
	"github.com/madsilver/template-bff/internal/adapter/presenter"
	"github.com/madsilver/template-bff/internal/infra/server/middleware"
)

func Start(graphqlServer GraphServer) chan error {
	e := echo.New()

	e.Use(middleware.HeaderMapper())

	e.Server.Addr = "0.0.0.0:8080"

	e.POST("query", handler.NewGraphqlHandler(graphqlServer).Handle)
	e.GET("health", health)

	err := make(chan error, 1)

	go func() {
		err <- gracehttp.Serve(e.Server)
	}()

	return err
}

func health(c echo.Context) error {
	return c.JSON(200, presenter.Health{Message: "it works"})
}
