package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"log/slog"
	"net/http"
	"strconv"
)

func getAndValidateCustomer(c *gin.Context, customer interface{}) (int64, bool) {
	customerID, ok := parseCustomerId(c)
	if !ok {
		return 0, false
	}
	if !bindCustomerData(c, customer) {
		return 0, false
	}
	return customerID, true
}

func parseCustomerId(c *gin.Context) (int64, bool) {
	customerId := c.Param("id")
	parseCusId, err := strconv.ParseInt(customerId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customerId is invalid"})
	}
	return parseCusId, err == nil
}

func bindCustomerData(c *gin.Context, customer interface{}) bool {
	if err := c.ShouldBindJSON(customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Check the data you are passing."})
		return false
	}
	return true
}

func (s *Server) CreateCustomer(c *gin.Context) {
	customer := models.Customer{}

	if !bindCustomerData(c, &customer) {
		return
	}
	addCustomer, err := s.db.AddCustomer(c, customer)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	c.JSON(http.StatusOK, addCustomer)
}

func (s *Server) UpdateCustomer(c *gin.Context) {
	customer := models.CustomerParams{}
	customerID, ok := getAndValidateCustomer(c, &customer)
	if !ok {
		return
	}

	updateCustomer, err := s.db.UpdateCustomer(c, customer, customerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updateCustomer {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Customer Information Updated!"})
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong."})
}

func (s *Server) DeleteCustomer(c *gin.Context) {
	customerID, ok := parseCustomerId(c)
	if !ok {
		return
	}

	if err := s.db.DeleteCustomer(c, customerID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "customer deleted"})
}
