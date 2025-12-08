package repository

import (
	"context"

	"github.com/SunilKividor/donela/internal/models"
	"github.com/google/uuid"
)

type SongRepository struct {
}

func NewSongRepository() *SongRepository {
	return &SongRepository{}
}

func (s *SongRepository) CreateSong(ctx context.Context, db DBTX, song *models.Song) (string, error) {

	smt := `INSERT INTO songs(artist_id,album_id,title,genre,status,release_date) VALUES($1,$2,$3,$4,$5,$6) RETURNING id`

	var id uuid.UUID
	status := "pending_upload"
	err := db.QueryRow(ctx, smt, song.ArtistID, song.AlbumID, song.Title, song.Genre, status, song.ReleaseDate).Scan(&id)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
