package handler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/SunilKividor/donela/internal/storage"
	"github.com/SunilKividor/donela/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StorageHandler struct {
	storage storage.StorageService
}

func NewStorageHandler(store storage.StorageService) *StorageHandler {
	return &StorageHandler{
		storage: store,
	}
}

func (sh *StorageHandler) UploadURL(c *gin.Context) {
	ctx := context.Background()

	bucket := os.Getenv("S3Bucket")
	prefix := "raw"
	userId := uuid.New().String() //get this from the gin context - user id
	contentExt := ".flac"

	key := util.GenerateNewAWSObjectKey(prefix, userId, contentExt)

	url, err := sh.storage.UploadURL(ctx, bucket, key, 5*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
