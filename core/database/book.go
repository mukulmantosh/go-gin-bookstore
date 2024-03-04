package database

import (
	"context"
	"go-gin-bookstore/models"
)

func (c Client) AddBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	var maxID int
	if result := c.db.Model(&models.Book{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID); result.Error != nil {
		panic("failed to get max id")
	}
	book.ID = uint(maxID) + 1
	result := c.db.WithContext(ctx).Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}
