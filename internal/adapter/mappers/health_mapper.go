package mappers

import (
	"github.com/madsilver/template-bff/internal/adapter/presenter"
	"github.com/madsilver/template-bff/internal/entity"
)

func HealthPresenter(ent *entity.Health) *presenter.Health {
	if ent == nil {
		return nil
	}
	return &presenter.Health{
		Message:    ent.Message,
		StatusCode: ent.StatusCode,
	}
}
