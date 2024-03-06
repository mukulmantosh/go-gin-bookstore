package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"net/http"
	"os"
	"strconv"
)

func (s *Server) CreateBook(c *gin.Context) {
	var book models.BookParams

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := book.ParsePublicationDate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PublicationDate"})
		return
	}

	addBook, err := s.db.AddBook(c, &book)
	if err != nil {
		panic("err")
	}
	c.JSON(http.StatusOK, addBook)
}

func (s *Server) ListBook(c *gin.Context) {
	listBooks, err := s.db.ListBooks(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, listBooks)
	return
}

func (s *Server) UpdateBook(c *gin.Context) {
	var book models.UpdateBookParams
	bookId := c.Param("id")
	parseBookId, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookId is invalid"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Check the data you are passing."})
		return
	}

	if _, err := book.ParsePublicationDate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid PublicationDate"})
		return
	}

	updateBook, err := s.db.UpdateBook(c, book, parseBookId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updateBook == true {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Book Updated!"})
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong."})
	return
}

func (s *Server) DeleteBook(c *gin.Context) {
	bookId := c.Param("id")
	parseBookId, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookId is invalid"})
		return
	}

	if err = s.db.DeleteBook(c, parseBookId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "book deleted"})
	return
}

func (s *Server) UploadBookCover(c *gin.Context) {
	bookId := c.Param("id")
	parseBookId, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bookId is invalid"})
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error while getting file: %s", err.Error()))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("Error opening file: %s", err.Error()))
		return
	}
	defer file.Close()
	// Create an uploader passing it the client
	uploader := manager.NewUploader(s.s3)

	uploadResult, err := uploader.Upload(c, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(fileHeader.Filename),
		Body:   file,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Update URL in DB
	_, err = s.db.UpdateBookCover(c, parseBookId, uploadResult.Location)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
	}

	c.JSON(200, gin.H{"status": true, "message": "Image Updated Successfully!"})
	return
}
