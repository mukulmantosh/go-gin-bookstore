package main

import (
	_ "github.com/joho/godotenv/autoload"
	"go-gin-bookstore/core/controllers"
	"go-gin-bookstore/core/database"
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

	service := controllers.NewServer(db)
	log.Fatal(service.Start())
}
