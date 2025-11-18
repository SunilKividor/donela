package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewSQSClient(ctx context.Context, c *Config) (*sqs.Client, error) {

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(c.AwsIAMConfig.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.AwsIAMConfig.AccessKey, c.AwsIAMConfig.AccessSecret, "")),
	)

	if err != nil {
		return nil, err
	}

	return sqs.NewFromConfig(cfg), nil
}
