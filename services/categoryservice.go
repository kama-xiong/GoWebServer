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

func (s *ProductCategoryService) GetRoot() models.Category {
	var category models.Category
	datasource.Db.Where("level=?", 0).First(&category)
	return category
}

func (s *ProductCategoryService) GetById(Id int) models.Category {
	var catetory models.Category
	datasource.Db.First(&catetory, Id)
	return catetory
}
func (s *ProductCategoryService) GetByName(name string) models.Category {
	var catetory models.Category
	datasource.Db.Where("name=?", name).Limit(1).Find(&catetory)

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
	datasource.Db.Where("lft>? and rgt<? and level=?+1", kind.Lft, kind.Rgt, kind.Level).Find(&categories)
	return categories
}
func (s *ProductCategoryService) GetChildrenByName(name string) []models.Category {
	return s.GetChildren(s.GetByName(name))
}
func (s *ProductCategoryService) GetPageChildrenByName(page models.Page) []models.Category {
	return s.GetChildren(s.GetByName(page.Keyword))
}
