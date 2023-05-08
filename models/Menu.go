package models

type Menu struct {
	ID      int64  `json:"id" gorm:"primaryKey"`
	Name_ch string `json:"name" gorm:"size:20"`
	Name_en string `json:"name" gorm:"size:20"`
	Router  string `json:"name" gorm:"size:20"`
	Action  string `json:"fullName" gorm:"size:20"` //paperbox,coveredbox,pouch,display,paperbag
	Order   int    `json:"order"`
}
