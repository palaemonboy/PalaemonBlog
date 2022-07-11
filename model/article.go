package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Category Category `gorm:"foreignkey:Cid"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200);not null" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(1000)" json:"img"`
}
