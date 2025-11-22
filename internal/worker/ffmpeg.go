package worker

import (
	"fmt"
	"os"
	"os/exec"
)

type FFmpegService interface {
	GenerateAACVariants(input, outputDir string) error
	GenerateMasterPlaylist(qualities []string) string
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

func (f FFmpeg) GenerateMasterPlaylist(qualities []string) string {
	playlist := "#EXTM3U\n#EXT-X-VERSION:3\n\n"

	for _, q := range qualities {
		var bw int
		var avg int

		switch q {
		case "aac_64k":
			bw = 80000
			avg = 65000
		case "aac_128k":
			bw = 140000
			avg = 130000
		case "aac_256k":
			bw = 280000
			avg = 260000
		}

		playlist += fmt.Sprintf(
			"#EXT-X-STREAM-INF:BANDWIDTH=%d,AVERAGE-BANDWIDTH=%d,CODECS=\"mp4a.40.2\"\n%s/playlist.m3u8\n\n",
			bw, avg, q,
		)
	}

	return playlist
}
