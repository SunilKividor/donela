package main

import (
	"github.com/SunilKividor/donela/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	server := api.NewServer(engine, "3000")

	api.RegisterRoutes(engine)

	if err := server.Serve(); err != nil {
		panic(err)
	}

}
