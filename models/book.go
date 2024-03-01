package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id              string `json:"Id" gorm:"primaryKey"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	PublicationDate string `json:"publication_date"`
}
