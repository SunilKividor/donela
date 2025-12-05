package api

import (
	"net/http"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config, handlers *handler.Handlers) {

	v1 := r.Group("/api/v1")
	{
		//auth
		v1.POST("/signup", handlers.Authentication.SignUp)
		v1.POST("/login", handlers.Authentication.Login)
		v1.POST("/refresh", handlers.Authentication.Refresh)

		v1.GET("/upload/signed", handlers.Storage.UploadURL)
		v1.GET("/download/signed", handlers.Storage.DownloadURL)

		v1.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "server running"}) })
	}
}
