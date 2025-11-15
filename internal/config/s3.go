package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(ctx context.Context) (*s3.Client, error) {
	accessKey := os.Getenv("S3AccessKey")
	accessSecret := os.Getenv("S3AccessSecret")
	region := os.Getenv("S3Region")
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, accessSecret, "")),
	)

	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}
