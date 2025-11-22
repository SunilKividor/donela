package worker

import (
	"fmt"
	"os"
)

func CreateTempDir(trackID string) (string, error) {
	path := fmt.Sprintf("/tmp/%s", trackID)

	err := os.MkdirAll(path, 0755)
	return path, err
}

func CleanTempDir(path string) error {
	return os.RemoveAll(path)
}
