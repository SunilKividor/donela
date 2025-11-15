package di

import (
	"context"

	"github.com/SunilKividor/donela/internal/api"
	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/storage"
)

func Initialize() (*api.Server, error) {
	ctx := context.Background()

	s3Client, err := config.NewS3Client(ctx)
	if err != nil {
		return nil, err
	}

	storage := storage.NewS3StorageClient(s3Client)

	server := api.NewServer(storage)

	return server, nil
}
