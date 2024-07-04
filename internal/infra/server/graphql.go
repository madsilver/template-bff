package server

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	gqlgenHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/madsilver/template-bff/internal/adapter"
	"github.com/madsilver/template-bff/internal/adapter/graph"
	"github.com/madsilver/template-bff/internal/domain/dto"
)

type GraphServer adapter.GraphqlServer

func NewGraphqlServer(resolver *graph.Resolver) GraphServer {
	return gqlgenHandler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: resolver,
				Directives: graph.DirectiveRoot{
					HasAnyRole: HasAnyRole,
				},
			},
		),
	)
}

func HasAnyRole(ctx context.Context, o any, next graphql.Resolver, roles []string) (res any, err error) {
	rctx := ctx.Value(dto.RequestContextKey).(dto.RequestContext)
	if rctx.HasAnyRole(roles...) {
		return next(ctx)
	}
	return nil, errors.New("access denied")
}
