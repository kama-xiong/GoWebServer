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
	ctx iris.Context
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
	b.Handle("GET", "/{kind:string}", "BigCategoryHandler")
	b.Handle("GET", "/{kind:string}/{smallkind:string}", "SmallCategoryHandler")
	b.Handle("GET", "/{bigkind:string}/{smallkind:string}/{pageNum:int}", "SmallCategoryPageHandler")

}
func (p *ProductController) BigCategoryHandler(kind string) mvc.Result {
	service := new(services.ProductCategoryService)
	data := service.GetChildrenByName(kind)
	fmt.Println("big Category handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": strings.ToLower(kind),
		},
	}
}

func (p *ProductController) SmallCategoryHandler(kind string, smallkind string) mvc.Result {
	service := new(services.ProductCategoryService)
	data := service.GetChildrenByName(smallkind)
	fmt.Println("small Category handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": strings.ToLower(kind),
		},
	}
}

func (p *ProductController) SmallCategoryPageHandler(kind string, smallkind string, pagenum int) mvc.Result {
	service := new(services.ProductCategoryService)
	page := models.Page{PageNum: pagenum, PageSize: 12, Keyword: smallkind}
	data := service.GetPageChildrenByName(page)
	fmt.Println("small Category page handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": strings.ToLower(kind),
		},
	}
}
