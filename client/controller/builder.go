package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type BuilderController interface {
	Get(ctx iris.Context)
}

type Builder struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Get(ctx iris.Context) {
	parent := opentracing.GlobalTracer().StartSpan("hello")
	//ctx1 := context.Background()
	//c := opentracing.ContextWithSpan(context.Background(), parent)
	//ctx.Set("ctx", c)

	parent.LogFields(
		log.String("event", "zdsdsdsdsd"),
	)

	parent.LogFields(
		log.String("event", "xxxxx"),
	)
	defer parent.Finish()
	name := ctx.URLParam("name")
	build := &Builder{
		Id:       1,
		Name:     name,
		Password: "12345678",
	}
	ctx.JSON(build)
	child := opentracing.GlobalTracer().StartSpan(
		"world", opentracing.ChildOf(parent.Context()))
	child.LogKV("xxx", "dsdsdsdsdsd")
	child.SetTag("hh", "xx")
	child.SetBaggageItem("BaggageKey", "BaggageValue")
	value := child.BaggageItem("BaggageKey")
	fmt.Println(value)
	defer child.Finish()
	child1 := opentracing.GlobalTracer().StartSpan(
		"world1", opentracing.ChildOf(child.Context()))
	defer child1.Finish()
}
