package tracer

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type OtelSpanWrapper struct {
	oteltrace.Span
}

func (s OtelSpanWrapper) End() {
	s.Span.End()
}

func (s OtelSpanWrapper) AddEvent(name string, opts ...interface{}) {
	s.Span.AddEvent(name)
}

type (
	ITracer interface {
		NewRootSpan(request *http.Request, spanName string) (ISpan, context.Context)
		NewSpanContext(ctx context.Context, spanName string) (ISpan, context.Context)
		AddAttributes(span ISpan, name string, attributes Attributes)
		Shutdown(ctx context.Context, logger logger.ILogger)
	}

	Tracer struct {
		provider *trace.TracerProvider
	}

	Attributes map[string]interface {
	}
)

func NewTracer(ctx context.Context, logger logger.ILogger, appName, env string) ITracer {

	var endpointOption otlptracegrpc.Option
	if env == "local" {
		endpointOption = otlptracegrpc.WithEndpoint("localhost:4317")
	} else {
		endpointOption = otlptracegrpc.WithEndpoint("otel-collector:4317")
	}

	exporter, err := otlptracegrpc.
		New(
			ctx,
			endpointOption,
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
			)),
		)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return &Tracer{
		provider: tp,
	}
}

func convertAttributes(attributes Attributes) []attribute.KeyValue {
	values := make([]attribute.KeyValue, 0, len(attributes))

	for k, v := range attributes {

		switch t := v.(type) {
		case string:
			values = append(values, attribute.String(k, t))
			continue
		case int:
			values = append(values, attribute.Int(k, t))
			continue
		case int64:
			values = append(values, attribute.Int64(k, t))
			continue
		case bool:
			values = append(values, attribute.Bool(k, t))
			continue
		default:
			val, err := json.Marshal(v)
			if err != nil {
				continue
			}

			values = append(values, attribute.String(k, string(val)))
		}

	}

	return values
}

func (t *Tracer) NewRootSpan(request *http.Request, spanName string) (ISpan, context.Context) {
	ctx := otel.
		GetTextMapPropagator().
		Extract(request.Context(), propagation.HeaderCarrier(request.Header))

	return t.NewSpanContext(ctx, spanName)
}

func (t *Tracer) NewSpanContext(ctx context.Context, spanName string) (ISpan, context.Context) {
	appName := env.ProvideAppEnv()
	tracer := otel.Tracer(appName)

	commonLabels := []attribute.KeyValue{
		attribute.String("service.name", appName),
	}

	ctx, span := tracer.Start(
		ctx,
		spanName,
		oteltrace.WithAttributes(commonLabels...),
	)

	span.SetStatus(codes.Ok, spanName)

	return OtelSpanWrapper{span}, ctx
}

func (t *Tracer) AddAttributes(span ISpan, name string, attributes Attributes) {
	values := convertAttributes(attributes)

	span.AddEvent(name, oteltrace.WithAttributes(values...))
}

func (t *Tracer) Shutdown(ctx context.Context, logger logger.ILogger) {
	if t.provider != nil {
		err := t.provider.Shutdown(ctx)
		if err != nil {
			logger.Error(ctx, "failed to shutdown tracer", err)
			return
		}
	}

	logger.Info(ctx, "Tracer shutdown completed successfully")
}
