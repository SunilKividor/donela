package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	client *s3.Client
}

func NewS3StorageClient(client *s3.Client) *S3Storage {
	return &S3Storage{
		client: client,
	}
}

func (c *S3Storage) UploadURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(c.client)

	presignReq, err := presignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
		s3.WithPresignExpires(expiry),
	)

	if err != nil {
		return "", fmt.Errorf("error generating pre-signed put url from s3")
	}
	return presignReq.URL, nil
}

func (c *S3Storage) DownloadURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error) {
	return "", nil
}

func (c *S3Storage) Upload(ctx context.Context, bucket, key string, data io.Reader) error {
	return errors.New("direct upload not supported for S3; use signed URLs")
}

func (c *S3Storage) Download(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	return nil, errors.New("direct download not supported for S3; use signed URLs")
}

func (c *S3Storage) Delete(ctx context.Context, bucket, key string) error {
	return nil
}

func (c *S3Storage) Exists(ctx context.Context, bucket, key string) (bool, error) {
	return false, nil
}
