package worker

import (
	"fmt"
	"os"
	"os/exec"
)

type FFmpegService interface {
	GenerateAACVariants(input, outputDir string) error
}

type FFmpeg struct{}

func NewFFmpegService() FFmpegService { return FFmpeg{} }

func (f FFmpeg) GenerateAACVariants(inputPath, tempDir string) error {
	bitrates := []string{"64k", "128k", "256k"}

	for _, br := range bitrates {
		fmt.Printf("[FFMPEG] Processing : %s\n", br)
		outDir := fmt.Sprintf("%s/aac_%s", tempDir, br)
		os.MkdirAll(outDir, 0755)

		playlist := fmt.Sprintf("%s/playlist.m3u8", outDir)
		segment := fmt.Sprintf("%s/segment_%%03d.ts", outDir)

		cmd := exec.Command(
			"ffmpeg",
			"-i", inputPath,
			"-map", "0:a:0",
			"-vn",
			"-c:a", "aac",
			"-b:a", br,
			"-f", "hls",
			"-hls_time", "5",
			"-hls_playlist_type", "vod",
			"-hls_segment_filename", segment,
			playlist,
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func GenerateMasterPlaylist(uuid string, qualities []string) string {
	template := "#EXTM3U\n#EXT-X-VERSION:3\n\n"

	for _, q := range qualities {
		switch q {
		case "aac_64k":
			template += "#EXT-X-STREAM-INF:BANDWIDTH=64000,CODECS=\"mp4a.40.2\"\n"
		case "aac_128k":
			template += "#EXT-X-STREAM-INF:BANDWIDTH=128000,CODECS=\"mp4a.40.2\"\n"
		case "aac_256k":
			template += "#EXT-X-STREAM-INF:BANDWIDTH=256000,CODECS=\"mp4a.40.2\"\n"
		case "opus_96k":
			template += "#EXT-X-STREAM-INF:BANDWIDTH=96000,CODECS=\"opus\"\n"
		case "opus_160k":
			template += "#EXT-X-STREAM-INF:BANDWIDTH=160000,CODECS=\"opus\"\n"
		}
		template += fmt.Sprintf("%s/playlist.m3u8\n\n", q)
	}

	return template
}
