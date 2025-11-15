package storage

import (
	"context"
	"io"
	"time"
)

type StorageService interface {
	//cloud
	UploadURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error)
	DownloadURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error)

	//local processing
	Upload(ctx context.Context, bucket, key string, data io.Reader) error
	Download(ctx context.Context, buket, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, bucket, key string) error

	Exists(ctx context.Context, bucket, key string) (bool, error)
}
