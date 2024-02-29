package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"net/http"
)

func (s *Server) AddBookToDB(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	addBook, err := s.db.AddBook(c, &book)
	if err != nil {
		panic("err")
	}
	c.JSON(http.StatusOK, gin.H{"data": addBook})
}
