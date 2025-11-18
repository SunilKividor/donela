package di

import (
	"context"

	"github.com/SunilKividor/donela/internal/api"
	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/storage"
)

func InitializeApp() (*api.Server, error) {
	ctx := context.Background()

	cfg := config.Load()

	s3Client, err := config.NewS3Client(ctx, cfg)
	if err != nil {
		return nil, err
	}

	storage := storage.NewS3StorageClient(s3Client)

	server := api.NewServer(cfg, storage)

	return server, nil
}
