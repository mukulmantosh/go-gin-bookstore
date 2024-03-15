package abstract

import (
	"context"
	"go-gin-bookstore/models"
)

type Review interface {
	AddReview(ctx context.Context, revParams models.ReviewParams) (bool, error)
	ListReview(ctx context.Context, bookId int64) ([]models.ReviewInfo, error)
}
