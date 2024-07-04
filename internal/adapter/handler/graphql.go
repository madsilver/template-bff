package handler

import (


	"github.com/labstack/echo/v4"
	"github.com/madsilver/template-bff/internal/adapter"
)

type graphqlHandler struct {
	graphqlServer adapter.GraphqlServer
}

func NewGraphqlHandler(graphqlServer adapter.GraphqlServer) adapter.GraphqlHandler {
	return &graphqlHandler{
		graphqlServer,
	}
}

func (h *graphqlHandler) Handle(c echo.Context) error {
	h.graphqlServer.ServeHTTP(c.Response(), c.Request())
	return nil
}
