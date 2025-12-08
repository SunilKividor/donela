package models

import "time"

type Album struct {
	Id                   string    `json:"id"`
	ArtistID             string    `json:"artist_id"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	Genre                string    `json:"genre"`
	Type                 string    `json:"type"`
	CoverImageURL        string    `json:"cover_image_url"`
	ReleaseDate          time.Time `json:"release_date"`
	TotalTracks          int       `json:"total_tracks"`
	TotalDurationSeconds int64     `json:"total_duration_seconds"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
