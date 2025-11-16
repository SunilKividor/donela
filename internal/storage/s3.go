package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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
		return "", fmt.Errorf("error generating pre-signed PUT url from s3")
	}
	return presignReq.URL, nil
}

func (c *S3Storage) DownloadURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error) {
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
