package di

import (
	"context"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/worker"
)

func InitializeWorker() (*worker.Worker, error) {
	ctx := context.Background()

	cfg := config.Load()

	sqsClient, err := config.NewSQSClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	w := worker.NewWorker(sqsClient, cfg.AwsSQSConfig.QueueURL)

	return w, nil
}
