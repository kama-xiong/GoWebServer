package services

import (
	"GoWebServer/datasource"
	"GoWebServer/models"
	"fmt"
	"github.com/kataras/iris/v12"
)

type ProductService struct {
	ctx iris.Context
}

func (p *ProductService) GetProducts(limit int) []models.Product {
	var products []models.Product
	datasource.Db.Limit(limit).Find(&products)
	return products
}
func (p *ProductService) GetPageProducts(page models.Page) []models.Product {
	var products []models.Product
	var category models.Category
	if err := datasource.Db.Where("name=?", page.Keyword).Find(&category).Error; err != nil {
		p.ctx.View("err_msg", err.Error())
	} else {
		if err := datasource.Db.Preload("Category", "name=?", page.Keyword).Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&products).Error; err != nil {
			p.ctx.ViewData("err_msg", err.Error())
		}
	}
	fmt.Printf("products is %s", products)
	return products
}
