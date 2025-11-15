package api

import (
	"fmt"
	"log"

	"github.com/SunilKividor/donela/internal/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine  *gin.Engine
	storage storage.StorageService
}

func NewServer(storage storage.StorageService) *Server {

	engine := gin.New()
	engine.Use(gin.Logger())

	s := &Server{
		engine:  engine,
		storage: storage,
	}

	RegisterRoutes(engine, storage)

	return s
}

func (s *Server) Serve(port string) error {

	if port == "" {
		port = "3000"
		log.Printf("[INFO] No port provided. Using default %s\n", port)
	}

	if err := s.engine.Run(":" + port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
