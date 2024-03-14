package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/core"
	"go-gin-bookstore/models"
	"net/http"
)

func (s *Server) CreateReview(c *gin.Context) {
	var review models.ReviewParams

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := s.db.AddReview(c, review)

	if err != nil {
		// Check if the error is of type notFoundError
		var notFoundErr *core.NotFoundError
		if errors.As(err, &notFoundErr) {
			c.AbortWithStatusJSON(http.StatusNotFound,
				gin.H{"status": notFoundErr.Code, "message": notFoundErr.Message})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Review successfully added."})
}
