package database

import (
	"context"
	"errors"
	"go-gin-bookstore/models"
	"gorm.io/gorm"
	"log/slog"
)

func (c Client) AddAuthor(ctx context.Context, author models.Author) (*models.Author, error) {
	var AuthorModel models.Author
	AuthorModel.Name = author.Name
	result := c.db.WithContext(ctx).Create(&AuthorModel)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return nil, errors.New("unable to register author")
	}
	return &AuthorModel, nil
}

func (c Client) LinkAuthorBook(_ context.Context, params models.AuthorBook) (bool, error) {
	author := models.Author{Model: gorm.Model{ID: uint(params.AuthorID)}}
	book := models.Book{Model: gorm.Model{ID: uint(params.BookID)}}
	err := c.db.Model(&author).Association("Books").Append(&book)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c Client) ListAuthors(_ context.Context) ([]models.Author, error) {
	var authors []models.Author
	result := c.db.Preload("Books").Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}
