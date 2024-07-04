package domain

import (
	"context"

	"github.com/madsilver/template-bff/internal/domain/dto"
	"github.com/madsilver/template-bff/internal/entity"
)

type HealthUseCase interface {
	Check(ctx context.Context) (*entity.Health, error)
}

type HttpClient interface {
	Send(ctx context.Context, req *dto.Request) ([]byte, int, error)
}
