package main

import (
	"fmt"
	"github.com/ClareChu/tracing/client/controller"
	"github.com/kataras/iris/v12"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {
	cfg := config.Configuration{
		ServiceName: "client",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "localhost:6831", // 替换host
		},
	}
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	fmt.Println(err)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	app := iris.New()

	app.Get("/", controller.Get)

	app.Run(iris.Addr(":8080"))
}