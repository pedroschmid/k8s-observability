package tracing

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var Tracer = otel.Tracer("tracing-app")

func InitTracer() func() {
	ctx := context.Background()

	endpoint := os.Getenv("OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = "tempo.monitoring.svc.cluster.local:4318"
	}

	log.Printf("🚀 Initializing OpenTelemetry tracer (exporting to %s)...", endpoint)

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("❌ Failed to create OTLP exporter: %v", err)
	}

	res, _ := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("tracing-app"),
		),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	log.Println("✅ Tracer initialized successfully")

	return func() {
		log.Println("🛑 Shutting down tracer...")
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("⚠️ Error shutting down tracer: %v", err)
		} else {
			log.Println("💤 Tracer shut down cleanly")
		}
	}
}
