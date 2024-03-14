package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"net/http"
)

func (s *Server) CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//if _, err := book.ParsePublicationDate(); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PublicationDate"})
	//	return
	//}

	addBook, err := s.db.AddCustomer(c, customer)
	if err != nil {
		fmt.Println("err", err.Error())
		panic("err")
	}
	c.JSON(http.StatusOK, addBook)
}
