package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addCustomer, err := s.db.AddCustomer(c, customer)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	c.JSON(http.StatusOK, addCustomer)
}

func (s *Server) DeleteCustomer(c *gin.Context) {
	customerId := c.Param("id")
	parseCustomerId, err := strconv.ParseInt(customerId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customerId is invalid"})
		return
	}

	if err = s.db.DeleteCustomer(c, parseCustomerId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "customer deleted"})
}
