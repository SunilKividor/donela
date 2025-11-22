package config

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewR2Client(ctx context.Context, c *Config) (*s3.Client, error) {

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(c.R2Config.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.R2Config.AccessKey, c.R2Config.AccessSecret, "")),
	)

	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", c.R2Config.AccountId))
	}), nil
}
