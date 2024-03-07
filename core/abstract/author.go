package abstract

import (
	"context"
	"go-gin-bookstore/models"
)

type Author interface {
	AddAuthor(ctx context.Context, author models.Author) (*models.Author, error)
}
