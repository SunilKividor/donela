package repository

import (
	"context"

	"github.com/SunilKividor/donela/internal/models"
	"github.com/google/uuid"
)

type AlbumRepository struct{}

func NewAlbumRepository() *AlbumRepository {
	return &AlbumRepository{}
}

func (a *AlbumRepository) CreateAlbum(ctx context.Context, db DBTX, album *models.Album) (string, error) {

	smt := `INSERT INTO albums(artist_id,title,description,genre,type,release_date) VALUES($1,$2,$3,$4,$5,$6) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(ctx, smt, album.ArtistID, album.Title, album.Description, album.Genre, album.Type, album.ReleaseDate).Scan(&id)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
