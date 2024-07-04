package graph

import (
	"github.com/madsilver/template-bff/internal/domain"
	"github.com/madsilver/template-bff/internal/domain/usecases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	HealthUseCase domain.HealthUseCase
}

func NewResolver(http domain.HttpClient) *Resolver {
	return &Resolver{
		HealthUseCase: usecases.NewHealthUseCase(http),
	}
}
