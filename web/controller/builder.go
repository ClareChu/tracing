package controller

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/opentracing/opentracing-go"
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
	defer child.Finish()
}
