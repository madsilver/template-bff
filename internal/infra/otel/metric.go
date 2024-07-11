package otel

import (
	"context"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.uber.org/zap"
)

func InitMeter(ctx context.Context, logger *zap.Logger) *metric.MeterProvider {
	metricExporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint("0.0.0.0:4317"),
		otlpmetricgrpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		logger.Error(err.Error())
	}

	metricResource := metric.WithResource(resource.NewWithAttributes(
		"0.0.0.0:8080",
		semconv.ServiceNameKey.String("template-bff"),
		semconv.ServiceVersionKey.String("0.0.1"),
	))
	metricReader := metric.WithReader(metric.NewPeriodicReader(metricExporter, metric.WithInterval(time.Second)))
	meterProvider := metric.NewMeterProvider(metricResource, metricReader)

	if err := runtime.Start(
		runtime.WithMinimumReadMemStatsInterval(time.Second),
		runtime.WithMeterProvider(meterProvider),
	); err != nil {
		logger.Error(err.Error())
	}

	return meterProvider
}
