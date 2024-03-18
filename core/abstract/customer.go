package abstract

import (
	"context"
	"go-gin-bookstore/models"
)

type Customer interface {
	AddCustomer(ctx context.Context, cusParams models.Customer) (*models.Customer, error)
	//UpdateCustomer(ctx context.Context, updateBookParams models.DateParser, bookId int64) (bool, error)
	DeleteCustomer(ctx context.Context, customerId int64) error
}
