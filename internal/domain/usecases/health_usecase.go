package usecases

import (
	"context"
	"net/http"

	"github.com/madsilver/template-bff/internal/domain"
	"github.com/madsilver/template-bff/internal/domain/dto"
	"github.com/madsilver/template-bff/internal/entity"
)

type healthUseCase struct {
	http domain.HttpClient
}

func NewHealthUseCase(http domain.HttpClient) domain.HealthUseCase {
	return &healthUseCase{
		http,
	}
}

func (u *healthUseCase) Check(ctx context.Context) (*entity.Health, error) {
	ent := new(entity.Health)

	_, code, err := u.http.Send(ctx, &dto.Request{
		Method: http.MethodGet,
		URL:    "http://localhost:8080/health",
		Entity: ent,
	})

	ent.StatusCode = code

	return ent, err
}
