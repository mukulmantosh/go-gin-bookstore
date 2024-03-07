package database

import (
	"context"
	"errors"
	"go-gin-bookstore/models"
	"log/slog"
)

func (c Client) AddAuthor(ctx context.Context, author models.Author) (*models.Author, error) {
	var maxID int
	if result := c.db.Model(&models.Author{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID); result.Error != nil {
		return nil, errors.New("something went wrong")
	}
	var AuthorModel models.Author
	AuthorModel.ID = uint(maxID) + 1
	AuthorModel.Name = author.Name
	result := c.db.WithContext(ctx).Create(&AuthorModel)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return nil, errors.New("unable to register author")
	}
	return &AuthorModel, nil
}

func (c Client) LinkAuthorBook(_ context.Context, params models.AuthorBook) (bool, error) {
	author := models.Author{Id: params.AuthorID}
	book := models.Book{Id: params.BookID}
	err := c.db.Model(&author).Association("Books").Append(&book)
	if err != nil {
		return false, err
	}
	return true, nil
}
