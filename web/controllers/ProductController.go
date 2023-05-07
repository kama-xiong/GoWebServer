package controllers

import (
	"GoWebServer/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
}

func (p *ProductController) Get() mvc.Result {
	service := new(services.ProductCategoryService)
	data := service.GetAll()
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data},
	}
}
