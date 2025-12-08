package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/storage"
	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	Config  *config.Config
	Storage storage.StorageService
}

func NewStorageHandler(cfg *config.Config, store storage.StorageService) *StorageHandler {
	return &StorageHandler{
		Config:  cfg,
		Storage: store,
	}
}

func (sh *StorageHandler) UploadURL(c *gin.Context) {
	// ctx := context.Background()

	// bucket := sh.Config.AwsS3Config.Bucket
	// prefix := "raw"
	// userId := uuid.New().String()
	// contentExt := ".flac"

	// key := util.GenerateNewAWSObjectKey(prefix, userId, contentExt)

	// url, err := sh.Storage.UploadURL(ctx, bucket, key, 5*time.Minute)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"url": url})
}

func (sh *StorageHandler) DownloadURL(c *gin.Context) {
	ctx := context.Background()

	bucket := sh.Config.AwsS3Config.Bucket
	key := "" //get from the request context

	url, err := sh.Storage.DownloadURL(ctx, bucket, key, 5*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
