package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracing() io.Closer {
	cfg := config.Configuration{
		ServiceName: "client",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
			SamplingServerURL: "localhost:5778",
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "localhost:6831", // 替换host
		},
	}
	tracer, closer, _ := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	opentracing.SetGlobalTracer(tracer)
	return closer
}
