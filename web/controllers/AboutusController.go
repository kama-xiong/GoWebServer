package controllers

import (
	"GoWebServer/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AboutUsController struct {
}

func (p *AboutUsController) Get() mvc.Result {
	cservice := new(services.ProductCategoryService)
	root := cservice.GetRoot()
	categories := cservice.GetChildrenByName(root.Name)
	return mvc.View{
		Name: "homePage.html",
		Data: iris.Map{"curnav": "aboutus",
			"title":    "Jewelry packaging professional manufacturer",
			"category": categories,
		},
	}
}
