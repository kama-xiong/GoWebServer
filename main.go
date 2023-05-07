package main

import (
	"GoWebServer/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("info")
	//全局设置
	//app.Use(before)
	//设置模板
	app.RegisterView(iris.Django("./web/views", ".html").Reload(true))
	//app.RegisterView(iris.HTML("./web/views", ".html").Reload(true).Layout("shared/_layout.html"))

	//加载静态文件
	app.HandleDir("/static", "./web/static")

	//错误处理
	//app.OnErrorCode(iris.StatusNotFound, routeFunc.NotFound)
	//app.OnErrorCode(iris.StatusInternalServerError, routeFunc.InternalServerError)

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("curnav", "home")
		ctx.View("homePage.html")

	})
	mvc.Configure(app.Party("/product"), productsMvc)
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))

}

func productsMvc(app *mvc.Application) {
	//app.Router.Use(middleware.BasicAuth)
	app.Handle(new(controllers.ProductController))
}
