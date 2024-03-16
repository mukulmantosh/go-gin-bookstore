package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Id          int64  `json:"ID" gorm:"primaryKey"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type Review struct {
	gorm.Model
	CustomerID int64    `json:"customer_id" binding:"required"`
	BookID     int64    `json:"book_id" binding:"required"`
	Rating     int      `json:"rating" binding:"required"`
	Comment    string   `json:"comment,omitempty"`
	Customer   Customer `gorm:"foreignKey:CustomerID" json:"-"`
	Book       Book     `gorm:"foreignKey:BookID" json:"-"`
}

type ReviewParams struct {
	CustomerID int64  `json:"customer_id" binding:"required"`
	BookID     int64  `json:"book_id" binding:"required"`
	Rating     int    `json:"rating" binding:"required" validate:"min=1,max=5"`
	Comment    string `json:"comment,omitempty" binding:"required"`
}

type ReviewList struct {
	Id      int64  `json:"id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
