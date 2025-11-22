package worker

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/storage"
)

type Job struct {
	Key     string `json:"key"`
	TrackID string
}

type Processor struct {
	uploadStorage   storage.StorageService
	downloadStorage storage.StorageService
	ffmpeg          FFmpegService
	config          config.Config
}

func NewProcessor(uploadStorage, downloadStorage storage.StorageService, ffmpeg FFmpegService, config config.Config) *Processor {
	return &Processor{
		uploadStorage:   uploadStorage,
		downloadStorage: downloadStorage,
		ffmpeg:          ffmpeg,
		config:          config,
	}
}

func (p *Processor) Process(ctx context.Context, job *Job) error {
	fmt.Println("[PROCESSOR] Starting job:", job.TrackID)

	tempDir, err := CreateTempDir(job.TrackID)
	if err != nil {
		return nil
	}
	defer CleanTempDir(tempDir)

	inputPath := filepath.Join(tempDir, "input.flac")
	reader, err := p.downloadStorage.Download(ctx, p.config.AwsS3Config.Bucket, job.Key)
	if err != nil {
		return err
	}
	data, _ := io.ReadAll(reader)
	err = os.WriteFile(inputPath, data, 0644)
	if err != nil {
		return err
	}
	err = reader.Close()
	if err != nil {
		return err
	}

	if err := p.ffmpeg.GenerateAACVariants(inputPath, tempDir); err != nil {
		return err
	}

	if err := p.uploadProcessed(ctx, tempDir, job.TrackID); err != nil {
		return err
	}

	fmt.Println("[PROCESSOR] Processed files")

	return nil
}

func (p *Processor) uploadProcessed(ctx context.Context, tempDir, trackID string) error {
	bitrates := []string{"aac_64k", "aac_128k", "aac_256k"}

	for _, br := range bitrates {
		dir := tempDir + "/" + br

		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, f := range files {
			fullPath := dir + "/" + f.Name()
			key := fmt.Sprintf("tracks/%s/%s/%s", trackID, br, f.Name())

			err := p.uploadStorage.Upload(ctx, p.config.R2Config.Bucket, key, fullPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
