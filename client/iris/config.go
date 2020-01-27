package iris

import (
	"github.com/ClareChu/tracing/client/controller"
	"github.com/ClareChu/tracing/client/iris/web"
	"github.com/kataras/iris/v12"
)

func NewConfig() {
	app := iris.New()
	app.Get("/", controller.Get)
	app.Get("/dns/start", web.Start)
	app.Get("/dns/done", web.Done)
	app.Run(iris.Addr(":8080"))
}
