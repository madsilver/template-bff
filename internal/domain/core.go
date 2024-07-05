package domain

import (
	"context"

	"github.com/madsilver/template-bff/internal/domain/dto"
	"github.com/madsilver/template-bff/internal/entity"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HealthUseCase interface {
	Check(ctx context.Context) (*entity.Health, error)
}

type HttpClient interface {
	Send(ctx context.Context, req *dto.Request) ([]byte, int, error)
}

type Logger interface {
	Sugar() *zap.SugaredLogger
	Named(s string) *zap.Logger 
	WithOptions(opts ...zap.Option) *zap.Logger 
	With(fields ...zap.Field) *zap.Logger 
	WithLazy(fields ...zap.Field) *zap.Logger 
	Level() zapcore.Level
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
	Log(lvl zapcore.Level, msg string, fields ...zap.Field) 
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field) 
	Warn(msg string, fields ...zap.Field) 
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field) 
	Fatal(msg string, fields ...zap.Field)
	Sync() error
	Core() zapcore.Core
	Name() string
}
