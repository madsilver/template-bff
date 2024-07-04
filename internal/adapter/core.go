package adapter

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/labstack/echo/v4"
)

type GraphqlServer interface {
	AddTransport(transport graphql.Transport)
	SetErrorPresenter(f graphql.ErrorPresenterFunc)
	SetRecoverFunc(f graphql.RecoverFunc)
	SetQueryCache(cache graphql.Cache)
	Use(extension graphql.HandlerExtension)
	AroundFields(f graphql.FieldMiddleware)
	AroundRootFields(f graphql.RootFieldMiddleware)
	AroundOperations(f graphql.OperationMiddleware)
	AroundResponses(f graphql.ResponseMiddleware)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type GraphqlHandler interface {
	Handle(c echo.Context) error
}
