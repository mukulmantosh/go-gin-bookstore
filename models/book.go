package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id              int64  `json:"ID" gorm:"primaryKey"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	Image           string `json:"image,omitempty"`
	PublicationDate string `json:"publication_date"`
}
