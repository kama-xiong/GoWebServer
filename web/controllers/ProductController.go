package controllers

import (
	"GoWebServer/models"
	"GoWebServer/services"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strings"
)

type ProductController struct {
}

//func (p *ProductController) GetBy(kind string) mvc.Result {
//	//menuservice := new(services.MenuService)
//	service := new(services.ProductCategoryService)
//	//menus := menuservice.GetAll()
//	data := service.GetChildrenByName(kind)
//	return mvc.View{
//		Name: "ProductCategory/productCategory.html",
//		Data: iris.Map{"content": data,
//			"curnav": strings.ToLower(kind),
//		},
//	}
//}

func (p *ProductController) BeforeActivation(b mvc.BeforeActivation) {
	println("before activation...")
	b.Handle("GET", "/{kind:string}/{pageNum:int}", "PageHandler")
	b.Handle("GET", "/{kind:string}", "CategoryHandler")
}
func (p *ProductController) PageHandler(kind string, pageNum int) mvc.Result {

	fmt.Println("The pagenum is" + string(pageNum))
	fmt.Println(pageNum)
	service := new(services.ProductService)
	page := models.Page{PageNum: pageNum, PageSize: 12, Keyword: kind, Desc: false}
	data := service.GetPageProducts(page)
	fmt.Println("first page handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": strings.ToLower(kind),
		},
	}
}
func (p *ProductController) CategoryHandler(kind string) mvc.Result {
	service := new(services.ProductCategoryService)
	data := service.GetChildrenByName(kind)
	fmt.Println("first page handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": strings.ToLower(kind),
		},
	}
}
