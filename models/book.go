package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Id              int64     `json:"ID" gorm:"primaryKey"`
	Title           string    `json:"title"`
	ISBN            string    `json:"isbn"`
	Image           string    `json:"image,omitempty"`
	PublicationDate time.Time `json:"publication_date"`
}

type CreateBookParams struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	PublicationDate string `json:"publication_date"`
}

func (params CreateBookParams) ParsePublicationDate() (time.Time, error) {
	standardFormat := "2006-01-02" // This layout represents the date format "YYYY-MM-DD"
	fmt.Println(params.PublicationDate)
	date, err := time.Parse(standardFormat, params.PublicationDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
