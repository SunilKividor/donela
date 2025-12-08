package api

import (
	"net/http"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config, handlers *handler.Handlers, middleware gin.HandlerFunc) {

	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", handlers.Authentication.SignUp)
		v1.POST("/login", handlers.Authentication.Login)
		v1.POST("/refresh", handlers.Authentication.Refresh)

		v1.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "server running"}) })
	}

	authRequired := v1.Group("/")
	authRequired.Use(middleware)
	{
		authRequired.POST("/songs", handlers.SongHandler.CreateSongWithAlbum)
		authRequired.POST("/logout", handlers.Authentication.Logout)
	}
}
