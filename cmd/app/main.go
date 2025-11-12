package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	err := os.MkdirAll("transcoded", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder")
		return
	}

	// cmd := exec.Command(
	// 	"ffprobe",
	// 	"-v", "error",
	// 	"-show_format",
	// 	"-show_streams",
	// 	"-print_format", "json",
	// 	"transcoded/Flash_in_the_Pan.webm",
	// )

	cmd := exec.Command(
		"ffmpeg", "-i",
		"flac_audio_samples/Flash_in_the_Pan.flac",
		"-vn",
		"-c:a", "libopus",
		"-b:a", "128k",
		"-ar", "48000",
		"-ac", "2",
		"transcoded/Flash_in_the_Pan.webm",
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		fmt.Println("Error running ffprobe:", err)
		fmt.Println("ffproble error: ", stderr.String())
		return
	}

}
