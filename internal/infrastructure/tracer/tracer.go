package tracer

import (
	"context"

	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type (
	ITracer interface {
	}

	Tracer struct {
	}
)

func NewTracer(ctx context.Context, logger logger.ILogger, appName string, appVersion string) ITracer {
	exporter, err := otlptracegrpc.
		New(
			ctx,
			otlptracegrpc.WithEndpoint("otel-collector:4317"),
			otlptracegrpc.WithInsecure(),
		)
	if err != nil {
		logger.Fatal(ctx, "failed to initialize tracer", err)
	}

	tp := trace.
		NewTracerProvider(
			trace.WithBatcher(exporter),
			trace.WithResource(resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(appName),
				semconv.ServiceVersionKey.String(appVersion),
			)),
		)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return &Tracer{}
}
