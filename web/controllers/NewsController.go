package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type NewsController struct {
}

func (p *NewsController) Get() mvc.Result {
	return mvc.View{
		Name: "News/News.html",
		Data: iris.Map{"curnav": "news"},
	}
}
