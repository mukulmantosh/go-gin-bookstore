package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-bookstore/core/database"
	"log"
	"log/slog"
)

type Server struct {
	gin *gin.Engine
	db  database.DBClient
}

func NewServer(db database.DBClient) *Server {

	server := &Server{
		gin: gin.Default(),
		db:  db,
	}
	server.endpoints()
	return server
}

func (s *Server) Start() error {
	slog.Info("serving at port 8080")
	err := s.gin.Run(":8080")
	if err != nil {
		log.Fatalf("Server Issue: %s", err)
		return err
	}
	return nil
}
