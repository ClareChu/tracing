package iris

import (
	"github.com/ClareChu/tracing/server/iris/web"
	"github.com/kataras/iris/v12"
)

func NewConfig() {
	app := iris.New()
	app.Get("/dns/start", web.Start)
	app.Get("/dns/done", web.Done)
	app.Run(iris.Addr(":8081"))
}
