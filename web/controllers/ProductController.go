package controllers

import (
	"GoWebServer/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
}

func (p *ProductController) GetPaperbox() mvc.Result {
	service := new(services.ProductCategoryService)
	data := service.GetAll()
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": "paperbox",
		},
	}
}
func (p *ProductController) GetCoveredbox() mvc.Result {
	service := new(services.ProductCategoryService)
	data := service.GetChildrenByName("PaperBox")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": "coveredbox",
		},
	}
}
