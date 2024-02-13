package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Id    string `json:"Id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Books []Book `json:"books" gorm:"foreignKey:AuthorId;references:Id"`
}
