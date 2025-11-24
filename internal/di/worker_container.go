package di

import (
	"context"

	"github.com/SunilKividor/donela/internal/config"
	queue "github.com/SunilKividor/donela/internal/queue"
	"github.com/SunilKividor/donela/internal/storage"
	"github.com/SunilKividor/donela/internal/worker"
)

func InitializeWorker() (*worker.Worker, error) {
	ctx := context.Background()

	cfg := config.Load()

	s3Client, err := config.NewS3Client(ctx, cfg)
	if err != nil {
		return nil, err
	}
	s3Storage := storage.NewS3StorageClient(s3Client)

	r2Client, err := config.NewR2Client(ctx, cfg)
	if err != nil {
		return nil, err
	}
	r2Storage := storage.NewR2StorageClient(r2Client)

	ffmpeg := worker.NewFFmpegService()

	processor := worker.NewProcessor(r2Storage, s3Storage, ffmpeg, *cfg)

	sqsClient, err := config.NewSQSClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	q := queue.NewSQSQueue(sqsClient, cfg.AwsSQSConfig.QueueURL)

	worker := worker.NewWorker(q, processor)

	return worker, nil
}
