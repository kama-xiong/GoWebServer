package services

import (
	"GoWebServer/datasource"
	"GoWebServer/models"
)

type MenuService struct {
}

func (m *MenuService) InitialData() {
	menus := []models.Menu{
		{Name_en: "About Us", Name_ch: "关于我们", Order: 1, Action: "aboutus", Router: "aboutus"},
		{Name_en: "Paper Box", Name_ch: "纸盒", Order: 2, Action: "paperbox", Router: "/product/paperbox"},
		{Name_en: "Covered Box", Name_ch: "精装盒", Order: 3, Action: "coveredbox", Router: "/product/coveredbox"},
		{Name_en: "Paper Bag", Name_ch: "纸袋", Order: 4, Action: "paperbag", Router: "/product/paperbag"},
		{Name_en: "Pouch", Name_ch: "珠包", Order: 5, Action: "pouch", Router: "/product/pouch"},
		{Name_en: "Display", Name_ch: "摆件", Order: 6, Action: "display", Router: "/product/display"},
		{Name_en: "New Product", Name_ch: "新产品", Order: 7, Action: "newproducts", Router: "newproducts"},
		{Name_en: "News", Name_ch: "新闻", Order: 8, Action: "news", Router: "news"},
		{Name_en: "Contact Us", Name_ch: "联系我们", Order: 9, Action: "contactus", Router: "contactus"},
	}
	datasource.Db.Create(&menus)
}
func (m *MenuService) GetAll() []models.Menu {
	menus := []models.Menu{}
	datasource.Db.Find(&menus)
	return menus
}
