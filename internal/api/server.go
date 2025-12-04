package api

import (
	"fmt"
	"log"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Engine *gin.Engine
}

func NewServer(cfg *config.Config) *Server {

	engine := gin.New()
	engine.Use(gin.Logger())

	s := &Server{
		Port:   cfg.ServerConfig.Port,
		Engine: engine,
	}

	return s
}

func (s *Server) Serve(port string) error {

	if port == "" {
		port = "3000"
		log.Printf("[INFO] No port provided. Using default %s\n", port)
	}

	if err := s.Engine.Run(":" + s.Port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
