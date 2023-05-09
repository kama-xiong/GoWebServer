package main

import (
	"GoWebServer/services"
	"GoWebServer/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("info")
	//全局设置
	app.UseGlobal(before)
	//app.DoneGlobal(after)
	//app.Done(after)
	//设置模板
	app.RegisterView(iris.Django("./web/views", ".html").Reload(true))
	//app.RegisterView(iris.HTML("./web/views", ".html").Reload(true).Layout("shared/_layout.html"))

	//加载静态文件
	app.HandleDir("/static", "./web/static")

	//错误处理
	//app.OnErrorCode(iris.StatusNotFound, routeFunc.NotFound)
	//app.OnErrorCode(iris.StatusInternalServerError, routeFunc.InternalServerError)

	//app.Get("/", mainHandler)
	app.Get("/initialdata", func(ctx iris.Context) {
		service := services.MenuService{}
		service.InitialData()
		ctx.Next()
	})
	mvc.Configure(app.Party("/"), webMvc)
	//mvc.Configure(app.Party("/product"), productsMvc)
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))

}

func productsMvc(app *mvc.Application) {
	//app.Router.Use(middleware.BasicAuth)

	app.Handle(new(controllers.ProductController))

}
func webMvc(app *mvc.Application) {

	app.Party("/").Handle(new(controllers.AboutUsController))
	app.Party("/aboutus").Handle(new(controllers.AboutUsController))
	app.Party("/contactus").Handle(new(controllers.ContactUsController))
	app.Party("/news").Handle(new(controllers.NewsController))
	app.Party("/newproducts").Handle(new(controllers.NewProductsController))
	app.Party("/product").Handle(new(controllers.ProductController))

}
func before(ctx iris.Context) {
	menus := new(services.MenuService).GetAll()
	ctx.ViewData("menus", menus)
	ctx.Next()
}
