package util

import (
	"fmt"
)

func GenerateRawUploadKey(songID string, artistID string, fileExtension string) string {

	prefix := "raw/"

	key := fmt.Sprintf("%s%s/%s%s",
		prefix,
		artistID,
		songID,
		fileExtension)

	return key
}
