package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"net/http"
)

func (s *Server) CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addAuthor, err := s.db.AddAuthor(c, author)
	if err != nil {
		panic("err")
	}
	c.JSON(http.StatusOK, addAuthor)
}

func (s *Server) LinkBook(c *gin.Context) {

	var authorBook models.AuthorBook

	if err := c.ShouldBindJSON(&authorBook); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	linkBook, err := s.db.LinkAuthorBook(c, authorBook)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if linkBook {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Book has been successfully linked!"})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong."})
	}

}
