package database

import (
	"context"
	"go-gin-bookstore/models"
)

func (c Client) AddBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	result := c.db.WithContext(ctx).Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}
