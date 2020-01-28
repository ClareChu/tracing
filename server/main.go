package main

import (
	"github.com/ClareChu/tracing/server/iris"
	"github.com/ClareChu/tracing/server/tracing"
)

func main() {
	closer := tracing.NewTracing()
	defer closer.Close()
	iris.NewConfig()

}
