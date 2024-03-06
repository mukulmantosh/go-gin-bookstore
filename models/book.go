package models

import (
	"gorm.io/gorm"
	"time"
)

// DateParser interface wraps the ParsePublicationDate method
type DateParser interface {
	ParsePublicationDate() (time.Time, error)
}

type Book struct {
	gorm.Model
	Id              int64     `json:"ID" gorm:"primaryKey"`
	Title           string    `json:"title" binding:"required"`
	ISBN            string    `json:"isbn" binding:"required"`
	Image           string    `json:"image,omitempty"`
	PublicationDate time.Time `json:"publication_date" binding:"required"`
}

type BookParams struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	PublicationDate string `json:"publication_date"`
}

// ParsePublicationDate Implementing the DateParser interface
func (params BookParams) ParsePublicationDate() (time.Time, error) {
	return ValidateDate(params.PublicationDate)
}

type UpdateBookParams struct {
	Title           string `json:"title" binding:"required"`
	ISBN            string `json:"isbn" binding:"required"`
	PublicationDate string `json:"publication_date" binding:"required"`
}

// ParsePublicationDate Implementing the DateParser interface
func (params UpdateBookParams) ParsePublicationDate() (time.Time, error) {
	return ValidateDate(params.PublicationDate)
}

func ValidateDate(pubDate string) (time.Time, error) {
	standardFormat := "2006-01-02" // This layout represents the date format "YYYY-MM-DD"
	date, err := time.Parse(standardFormat, pubDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
