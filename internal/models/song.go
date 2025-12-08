package models

import "time"

type Song struct {
	Id              string    `json:"id"`
	ArtistID        string    `json:"artist_id"`
	AlbumID         string    `json:"album_id"`
	Title           string    `json:"title"`
	Genre           string    `json:"genre"`
	AudioFileURL    string    `json:"audio_file_url"`
	CoverImageURL   string    `json:"cover_image_url"`
	Status          string    `json:"status"`
	DurationSeconds int64     `json:"duration_seconds"`
	Bitrate         int64     `json:"bitrate"`
	FileFormat      string    `json:"file_format"`
	FileSizeBytes   int64     `json:"file_size_bytes"`
	ReleaseDate     time.Time `json:"release_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateSongWithAlbumReq struct {
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
}

type CreateSongWithAlbumRes struct {
	SongID     string    `json:"song_id"`
	UploadURI  string    `json:"upload_uri"`
	Expiration time.Time `json:"exp"`
}
