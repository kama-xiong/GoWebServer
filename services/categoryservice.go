package services

import (
	"GoWebServer/datasource"
	"GoWebServer/models"
)

type CategoryService interface {
	GetById(Id int) []models.Category
	GetAll() []models.Category
	GetParent(kind models.Category) models.Category
	GetChildren(kind models.Category) []models.Category
}
type ProductCategoryService struct {
}

func (s *ProductCategoryService) GetById(Id int) models.Category {
	var catetory models.Category
	datasource.Db.First(&catetory, Id)
	return catetory
}

func (s *ProductCategoryService) GetAll() []models.Category {
	var categories []models.Category
	datasource.Db.Find(&categories)
	return categories

}
func (s *ProductCategoryService) GetParent(kind models.Category) models.Category {
	var category models.Category
	datasource.Db.Where("lft<? and rgt>? and level=?", kind.Lft, kind.Rgt, kind.Level-1).Find(&category)
	return category
}
func (s *ProductCategoryService) GetChildren(kind models.Category) []models.Category {
	var categories []models.Category
	datasource.Db.Where("lft>? and rgt<? and level>?", kind.Lft, kind.Rgt, kind.Level).Find(&categories)
	return categories
}
