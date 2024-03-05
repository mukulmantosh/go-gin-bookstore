package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/models"
	"net/http"
)

func (s *Server) AddBookToDB(c *gin.Context) {
	var book models.CreateBookParams

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
	c.JSON(http.StatusOK, gin.H{"data": addBook})
}

func (s *Server) UploadBookCover(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error while getting file: %s", err.Error()))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error opening file: %s", err.Error()))
		return
	}
	defer file.Close()
	// Create an uploader passing it the client
	uploader := manager.NewUploader(s.s3)

	uploadResult, err := uploader.Upload(c, &s3.PutObjectInput{
		Bucket: aws.String("go-bookstore"),
		Key:    aws.String(fileHeader.Filename),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(uploadResult.Location)
	c.JSON(200, gin.H{"status": true})

}
