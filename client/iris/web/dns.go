package web

import (
	"context"
	"fmt"
	"github.com/ClareChu/tracing/client/controller"
	"github.com/kataras/iris/v12"
	"github.com/opentracing/opentracing-go"
)

type Dns interface {
	Start(ctx iris.Context)
	Done(ctx iris.Context)
}

func Start(ctx iris.Context) {
	dns := controller.Dns{}
	parent := opentracing.GlobalTracer().StartSpan("dnsStart")
	ctx2 := context.Background()
	ctx2 = opentracing.ContextWithSpan(ctx2, parent)
	rep := dns.Start(ctx2)
	_, err := ctx.JSON(rep)
	fmt.Println(err)
	defer parent.Finish()
}

func Done(ctx iris.Context) {
	dns := controller.Dns{}
	parent := opentracing.GlobalTracer().StartSpan("dnsDone")
	ctx2 := context.Background()
	ctx2 = opentracing.ContextWithSpan(ctx2, parent)
	rep := dns.Done(ctx2)
	_, err := ctx.JSON(rep)
	fmt.Println(err)
	defer parent.Finish()
}
