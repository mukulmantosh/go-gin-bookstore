package abstract

import (
	"context"
	"go-gin-bookstore/models"
)

type Book interface {
	AddBook(ctx context.Context, bookParams models.DateParser) (*models.BookParams, error)
	ListBooks(ctx context.Context) ([]models.Book, error)
	UpdateBook(ctx context.Context, updateBookParams models.DateParser, userId int64) (bool, error)
}
