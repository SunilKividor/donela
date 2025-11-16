package api

import (
	"net/http"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/handler"
	"github.com/SunilKividor/donela/internal/storage"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config, storage storage.StorageService) {

	storageHandler := handler.NewStorageHandler(cfg, storage)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/upload/signed", storageHandler.UploadURL)
		v1.GET("/download/signed", storageHandler.DownloadURL)

		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"message": "server running",
				},
			)
		})
	}
}
