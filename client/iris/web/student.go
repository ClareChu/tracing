package web

import (
	"fmt"
	"github.com/ClareChu/tracing/client/controller"
	"github.com/kataras/iris/v12"
	"github.com/opentracing/opentracing-go"
)

func Get(ctx iris.Context) {
	id, err := ctx.URLParamInt64("id")
	dto := controller.Student{
		Id: id,
	}
	parent := opentracing.GlobalTracer().StartSpan("getStudent")
	ctx2 := opentracing.ContextWithSpan(ctx.Request().Context(), parent)
	rep := dto.Get(ctx2)
	_, err = ctx.JSON(rep)
	fmt.Println(err)
	defer parent.Finish()
}
