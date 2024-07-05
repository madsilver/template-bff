package usecases

import (
	"context"
	"net/http"

	"github.com/madsilver/template-bff/internal/domain"
	"github.com/madsilver/template-bff/internal/domain/dto"
	"github.com/madsilver/template-bff/internal/entity"
)

type healthUseCase struct {
	http   domain.HttpClient
	logger domain.Logger
}

func NewHealthUseCase(http domain.HttpClient, logger domain.Logger) domain.HealthUseCase {
	return &healthUseCase{
		http,
		logger,
	}
}

func (u *healthUseCase) Check(ctx context.Context) (*entity.Health, error) {
	u.logger.Info("health check")

	ent := new(entity.Health)
	_, code, err := u.http.Send(ctx, &dto.Request{
		Method: http.MethodGet,
		URL:    "http://localhost:8080/health",
		Entity: ent,
	})
	ent.StatusCode = code
	return ent, err
}
