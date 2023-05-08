package controllers

import (
	"GoWebServer/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strings"
)

type ProductController struct {
}

func (p *ProductController) GetBy(kind string) mvc.Result {
	//menuservice := new(services.MenuService)
	service := new(services.ProductCategoryService)
	//menus := menuservice.GetAll()
	data := service.GetChildrenByName(kind)
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": strings.ToLower(kind),
		},
	}

}
