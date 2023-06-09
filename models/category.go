package models

import "time"

type Category struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:20"`
	FullName  string    `json:"fullName" gorm:"size:60"`
	Lft       int64     `json:"lft"`
	Rgt       int64     `json:"rgt"`
	Level     int64     `json:"level"`
	ImgUrl    string    `json:"imgUrl" gorm:"size:100"`
	Route     string    `json:"route" gorm:"size:40"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
