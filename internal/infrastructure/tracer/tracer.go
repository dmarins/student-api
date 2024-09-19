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

type (
	ITracer interface {
		NewRootSpan(request *http.Request, spanName string) (oteltrace.Span, context.Context)
		NewSpanWithCtx(parentCtx context.Context, spanName string) (oteltrace.Span, context.Context)
		AddEvent(span oteltrace.Span, name string, attributes Attributes)
	}

	Tracer struct {
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

	return &Tracer{}
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

func (t *Tracer) NewRootSpan(request *http.Request, spanName string) (oteltrace.Span, context.Context) {
	parentContext := otel.
		GetTextMapPropagator().
		Extract(request.Context(), propagation.HeaderCarrier(request.Header))

	return t.NewSpanWithCtx(parentContext, spanName)
}

func (t *Tracer) NewSpanWithCtx(parentCtx context.Context, spanName string) (oteltrace.Span, context.Context) {
	appName := env.ProvideAppEnv()
	tracer := otel.Tracer(appName)

	commonLabels := []attribute.KeyValue{
		attribute.String("service.name", appName),
	}

	ctx, span := tracer.Start(
		parentCtx,
		spanName,
		oteltrace.WithAttributes(commonLabels...),
	)

	span.SetStatus(codes.Ok, spanName)

	return span, ctx
}

func (t *Tracer) AddEvent(span oteltrace.Span, name string, attributes Attributes) {
	values := convertAttributes(attributes)

	span.AddEvent(name, oteltrace.WithAttributes(values...))
}
