package database

import (
	"context"
	"errors"
	"go-gin-bookstore/core"
	"go-gin-bookstore/models"
	"gorm.io/gorm"
	"net/http"
)

func (c Client) AddReview(ctx context.Context, revParams models.ReviewParams) (bool, error) {
	var maxID int64
	if result := c.db.Model(&models.Review{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID); result.Error != nil {
		return false, errors.New("something went wrong")
	}

	var customer models.Customer
	var book models.Book
	if err := c.db.Where("id = ?", revParams.CustomerID).Take(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &core.NotFoundError{Code: http.StatusNotFound,
				Message: "Invalid Customer ID"}
		} else {
			return false, err
		}
	}

	if err := c.db.Where("id = ?", revParams.BookID).Take(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &core.NotFoundError{Code: http.StatusNotFound,
				Message: "Invalid Book ID"}
		} else {
			return false, err
		}
	}

	var Review models.Review
	Review.ID = uint(maxID + 1)
	Review.Rating = revParams.Rating
	Review.Comment = revParams.Comment
	Review.CustomerID = revParams.CustomerID
	Review.BookID = revParams.BookID

	c.db.Create(&Review)
	return true, nil

}
