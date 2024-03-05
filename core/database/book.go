package database

import (
	"context"
	"errors"
	"go-gin-bookstore/models"
	"log/slog"
)

func (c Client) AddBook(ctx context.Context, bookParams *models.CreateBookParams) (*models.CreateBookParams, error) {
	var maxID int
	if result := c.db.Model(&models.Book{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID); result.Error != nil {
		return nil, errors.New("something went wrong")
	}
	var Book models.Book

	Book.ID = uint(maxID) + 1
	Book.Title = bookParams.Title
	Book.ISBN = bookParams.ISBN
	Book.PublicationDate, _ = bookParams.ParsePublicationDate()

	bookParams.Id = int64(Book.ID)
	result := c.db.WithContext(ctx).Create(&Book)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return nil, errors.New("unable to register book")
	}
	return bookParams, nil
}
