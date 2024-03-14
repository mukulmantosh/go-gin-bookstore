package database

import (
	"fmt"
	"go-gin-bookstore/core/abstract"
	"go-gin-bookstore/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type DBClient interface {
	DBMigrate() error
	Ready() error
	abstract.Book
	abstract.Author
	abstract.Customer
	abstract.Review
}

type Client struct {
	db *gorm.DB
}

func NewClient() (DBClient, error) {
	databaseHost := os.Getenv("DB_HOST")
	databaseUsername := os.Getenv("DB_USERNAME")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	databasePort := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(databasePort)
	if err != nil {
		log.Fatal("Invalid DB Port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		databaseHost, databaseUsername, databasePassword, databaseName, dbPort, "disable")

	dbInfo, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	client := Client{db: dbInfo}
	return client, nil
}

func (c Client) DBMigrate() error {
	err := c.db.AutoMigrate(&models.Author{},
		&models.Book{},
		&models.Customer{},
		&models.Review{},
	)
	if err != nil {
		return err
	}
	return nil

}

func (c Client) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (c Client) CloseDBConnection() {
	db, err := c.db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	db.Close()
}
