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
	//println("before activation...")
	b.Handle("GET", "/{kind:string}", "BigCategoryHandler")
	b.Handle("GET", "/{kind:string}/{smallkind:string}", "SmallCategoryHandler")
	//b.Handle("GET", "/{bigkind:string}/{smallkind:string}/{pageNum:int}", "SmallCategoryPageHandler")

}
func (p *ProductController) BigCategoryHandler(kind string) mvc.Result {
	service := new(services.ProductCategoryService)
	pservice := new(services.ProductService)
	page := new(models.Page)
	var pdata []models.Product
	cdata := service.GetChildrenByName(kind)
	var curCategory models.Category
	if len(cdata) > 0 {
		page.PageNum = 1
		page.Keyword = cdata[0].Name
		page.PageSize = 12
		pdata = pservice.GetPageProducts(*page)
		curCategory = service.GetByName(page.Keyword)
	}
	//fmt.Println("big Category handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{
			"categoryData": cdata,
			"productData":  pdata,
			"curCategory":  curCategory,
			"curnav":       strings.ToLower(kind),
			"title":        kind,
		},
	}
}

func (p *ProductController) SmallCategoryHandler(kind string, smallkind string) mvc.Result {

	service := new(services.ProductCategoryService)
	pservice := new(services.ProductService)
	page := new(models.Page)
	var pdata []models.Product
	cdata := service.GetChildrenByName(kind)
	var curCategory models.Category
	//for k, v := range cdata {
	//	cdata[k].Route = strings.Replace(v.Route, kind, "", 1)
	//}
	if len(cdata) > 0 {
		page.PageNum = 1
		page.Keyword = smallkind
		page.PageSize = 12
		pdata = pservice.GetPageProducts(*page)
		curCategory = service.GetByName(page.Keyword)
	}
	//fmt.Println("small Category handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{
			"categoryData": cdata,
			"productData":  pdata,
			"curCategory":  curCategory,
			"curnav":       strings.ToLower(kind),
		},
	}
}

func (p *ProductController) SmallCategoryPageHandler(kind string, smallkind string, pagenum int) mvc.Result {
	service := new(services.ProductService)
	page := models.Page{PageNum: pagenum, PageSize: 12, Keyword: smallkind}
	data := service.GetPageProducts(page)
	fmt.Println("small Category page handler")
	return mvc.View{
		Name: "ProductCategory/productCategory.html",
		Data: iris.Map{"content": data,
			"curnav": strings.ToLower(kind),
		},
	}
}
