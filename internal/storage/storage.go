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
	Upload(ctx context.Context, bucket, key, fullPath string) error
	Download(ctx context.Context, buket, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, bucket, key string) error

	Exists(ctx context.Context, bucket, key string) (bool, error)
}

// type Song struct {
// 	Id           string `json:"id"`
// 	UserId       string `json:"user_id"`
// 	Title        string `json:"title"`
// 	Artist       string `json:"artist"`
// 	Album        string `json:"album"`
// 	Genre        string `json:"genre"`
// 	AudioFileURL string `json:"audio_file_url"`
// 	CoverFileURL string `json:"cover_file_url"`
// }
