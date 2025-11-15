package api

import (
	"net/http"

	"github.com/SunilKividor/donela/internal/handler"
	"github.com/SunilKividor/donela/internal/storage"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, storage storage.StorageService) {

	storageHandler := handler.NewStorageHandler(storage)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"message": "server running",
				},
			)
		})

		v1.GET("/upload/signed", storageHandler.UploadURL)
	}
}
