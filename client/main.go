package main

import (
	"github.com/ClareChu/tracing/client/iris"
	"github.com/ClareChu/tracing/client/tracing"
)

func main() {
	closer := tracing.NewTracing()
	defer closer.Close()
	iris.NewConfig()

}
