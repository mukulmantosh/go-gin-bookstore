package abstract

import (
	"context"
	"go-gin-bookstore/models"
)

type Customer interface {
	AddCustomer(ctx context.Context, cusParams models.Customer) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, updateCusParams models.CustomerParams, customerId int64) (bool, error)
	DeleteCustomer(ctx context.Context, customerId int64) error
}
