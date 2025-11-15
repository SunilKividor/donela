package util

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateNewAWSObjectKey(prefix string, userId string, contentExt string) string {
	key := uuid.New().String()

	awsS3BucketKey := fmt.Sprintf("%s/%s/%s%s", prefix, userId, key, contentExt)

	return awsS3BucketKey
}
