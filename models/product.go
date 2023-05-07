package models

type Product struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Category    string `json:"category" gorm:"size:60"`
	Name        string `json:"name" gorm:"size:80"`
	PackingInfo string `json:"packingInfo"`
	Size        string `json:"size" gorm:"size:20"`
	Cbm         string `json:"cbm" gorm:"size:20"`
	Gw          string `json:"gw" gorm:"size:20"`
	Nw          string `json:"nw" gorm:"size:20"`
	ImgUrl      string `json:"imgUrl" gorm:"size:100"`
}
