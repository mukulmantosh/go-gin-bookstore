package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"net/http"
)

func (s *Server) CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addAuthor, err := s.db.AddAuthor(c, author)
	if err != nil {
		panic("err")
	}
	c.JSON(http.StatusOK, addAuthor)
}
