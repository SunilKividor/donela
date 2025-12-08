package models

import "time"

type Playlist struct {
	Id                   string    `json:"id"`
	UserID               string    `json:"user_id"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	CoverImageURL        string    `json:"cover_image_url"`
	IsPublic             bool      `json:"is_public"`
	TotalTracks          int       `json:"total_tracks"`
	TotalDurationSeconds int64     `json:"total_duration_seconds"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
