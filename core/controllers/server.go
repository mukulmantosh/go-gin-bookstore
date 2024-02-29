package controllers

import "go-gin-bookstore/core/database"

type Server struct {
	db database.Client
}
