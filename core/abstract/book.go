package abstract

import (
	"context"
	"go-gin-bookstore/models"
)

type Book interface {
	AddBook(ctx context.Context, book *models.CreateBookParams) (*models.CreateBookParams, error)
}
