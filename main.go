package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/database"
	"log"
)

func main() {
	// establish connection with db
	db, err := database.NewClient()
	if err != nil {

	}
	err = db.DBMigrate()
	if err != nil {
		log.Fatal("Database Migration Failed!")
		return
	}

	router := gin.Default()
	log.Fatal(router.Run(":8080"))

}
