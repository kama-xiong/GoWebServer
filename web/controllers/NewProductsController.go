package controllers

import "github.com/kataras/iris/v12/mvc"

type NewProductsController struct {
}

func (p *NewProductsController) Get() mvc.Result {
	return mvc.View{
		Name: "NewProducts/NewProducts.html",
	}
}
