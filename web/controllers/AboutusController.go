package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AboutUsController struct {
}

func (p *AboutUsController) Get() mvc.Result {
	return mvc.View{
		Name: "homePage.html",
		Data: iris.Map{"curnav": "aboutus"},
	}
}
