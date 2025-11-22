package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type R2Storage struct {
	client *s3.Client
}

func NewR2StorageClient(client *s3.Client) *R2Storage {
	return &R2Storage{
		client: client,
	}
}

func (c *R2Storage) UploadURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error) {
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
		return "", fmt.Errorf("error generating pre-signed PUT url from s3")
	}
	return presignReq.URL, nil
}

func (c *R2Storage) DownloadURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(c.client)

	presignReq, err := presignClient.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
		s3.WithPresignExpires(expiry),
	)
	if err != nil {
		return "", fmt.Errorf("error generating pre-signed GET url from S3")
	}
	return presignReq.URL, nil
}

func (c *R2Storage) Upload(ctx context.Context, bucket, key, fullPath string) error {
	fmt.Println("[R2] Uploading:", fullPath, "→", key)

	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("failed to open file for upload: %w", err)
	}
	defer file.Close()

	_, err = c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file to R2: %w", err)
	}

	fmt.Println("[R2] Uploaded:", key)

	return nil
}

func (c *R2Storage) Download(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	return nil, errors.New("direct download not supported for S3; use signed URLs")
}

func (c *R2Storage) Delete(ctx context.Context, bucket, key string) error {
	return nil
}

func (c *R2Storage) Exists(ctx context.Context, bucket, key string) (bool, error) {
	if key == "" {
		return false, errors.New("key is required")
	}

	_, err := c.client.HeadObject(
		ctx,
		&s3.HeadObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
	)

	if err != nil {
		var nsk *types.NoSuchKey
		var nf *types.NotFound

		if errors.As(err, &nsk) || errors.As(err, &nf) {
			return false, fmt.Errorf("verification : object %s not found,", key)
		}

		return false, fmt.Errorf("error Verifying object %s:%v", key, err)
	}
	return true, nil
}
