package database

import (
	"context"
	"errors"
	"go-gin-bookstore/models"
)

func (c Client) AddCustomer(ctx context.Context, cusParams models.Customer) (*models.Customer, error) {
	var maxID int64
	if result := c.db.Model(&models.Customer{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID); result.Error != nil {
		return nil, errors.New("something went wrong")
	}
	var Customer models.Customer
	Customer.Id = maxID + 1
	Customer.FirstName = cusParams.FirstName
	Customer.LastName = cusParams.LastName
	Customer.Email = cusParams.Email
	Customer.PhoneNumber = cusParams.PhoneNumber
	Customer.Address = cusParams.Address

	c.db.Create(&Customer)
	return &Customer, nil
}
