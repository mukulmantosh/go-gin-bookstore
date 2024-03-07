package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Id    int64  `json:"Id" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Books []Book `gorm:"many2many:author_books;"`
}

type AuthorBook struct {
	AuthorID int64 `json:"author_id" binding:"required"`
	BookID   int64 `json:"book_id" binding:"required"`
}
