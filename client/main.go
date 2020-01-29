package main

import (
	"github.com/ClareChu/tracing/client/grpc"
	"github.com/ClareChu/tracing/client/iris"
	"github.com/ClareChu/tracing/client/tracing"
)

func main() {
	tracer, closer := tracing.NewTracing()
	defer closer.Close()
	grpc.NewConfig(tracer)
	iris.NewConfig()

}
