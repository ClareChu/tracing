package web

import (
	"fmt"
	"github.com/ClareChu/tracing/server/controller"
	"github.com/kataras/iris/v12"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type Dns interface {
	Start(ctx iris.Context)
	Done(ctx iris.Context)
}

func Start(ctx iris.Context) {
	dns := controller.Dns{}
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(ctx.Request().Header))
	if err != nil {
		// Optionally record something about err here
	}

	// Create the span referring to the RPC client if available.
	// If wireContext == nil, a root span will be created.
	serverSpan := opentracing.StartSpan(
		"serverDnsStart",
		ext.RPCServerOption(wireContext))

	ctx2 := opentracing.ContextWithSpan(ctx.Request().Context(), serverSpan)
	rep := dns.Start(ctx2)
	_, err = ctx.JSON(rep)
	fmt.Println(err)
	defer serverSpan.Finish()
}

func Done(ctx iris.Context) {
	dns := controller.Dns{}
	parent := opentracing.GlobalTracer().StartSpan("serverDnsDone")
	ctx2 := opentracing.ContextWithSpan(ctx.Request().Context(), parent)
	rep := dns.Done(ctx2)
	_, err := ctx.JSON(rep)
	fmt.Println(err)
	defer parent.Finish()
}
