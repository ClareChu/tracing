package main

import (
	"github.com/ClareChu/tracing/server/grpc"
	"github.com/ClareChu/tracing/server/iris"
	"github.com/ClareChu/tracing/server/tracing"
)

func main() {
	tracer, closer := tracing.NewTracing()
	defer closer.Close()

	s, err := grpc.NewConfig(tracer)
	if err != nil {
		return
	}
	defer s.Stop()
	iris.NewConfig()
}
