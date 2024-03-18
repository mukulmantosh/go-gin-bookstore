package database

import (
	"context"
	"errors"
	"go-gin-bookstore/models"
	"gorm.io/gorm"
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

func (c Client) DeleteCustomer(_ context.Context, customerId int64) error {
	var CustomerInfo = models.Customer{Id: customerId}

	if err := c.db.First(&CustomerInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("there is no customer associated with this ID")
		}
	}
	c.db.Delete(&CustomerInfo)
	return nil
}

func (c Client) UpdateCustomer(_ context.Context, updateCusParams models.CustomerParams, customerId int64) (bool, error) {
	var cusInfo = models.Customer{Id: customerId}
	if err := c.db.First(&cusInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("there is no customer associated with this ID")
		}
	}
	c.db.Model(&cusInfo).Updates(models.Customer{
		FirstName: updateCusParams.FirstName,
		LastName:  updateCusParams.LastName,
		Address:   updateCusParams.Address})
	return true, nil
}
