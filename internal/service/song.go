package service

import (
	"context"
	"log"
	"time"

	"github.com/SunilKividor/donela/internal/config"
	"github.com/SunilKividor/donela/internal/db/repository"
	"github.com/SunilKividor/donela/internal/models"
	"github.com/SunilKividor/donela/internal/storage"
	"github.com/SunilKividor/donela/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SongService struct {
	Config    *config.Config
	DB        *pgxpool.Pool
	SongRepo  repository.SongRepository
	AlbumRepo repository.AlbumRepository
	Storage   storage.StorageService
}

func NewSongService(config *config.Config, db *pgxpool.Pool, s repository.SongRepository, a repository.AlbumRepository, store storage.StorageService) *SongService {
	return &SongService{
		DB:        db,
		SongRepo:  s,
		AlbumRepo: a,
		Storage:   store,
		Config:    config,
	}
}

func (s *SongService) CreateSongWithAlbum(ctx context.Context, song *models.Song) (*models.CreateSongWithAlbumRes, error) {
	txn, err := s.DB.Begin(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = txn.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = txn.Rollback(ctx)
		}
	}()

	var album models.Album
	album.ArtistID = song.ArtistID
	album.Title = song.Title
	album.Genre = song.Genre
	album.Type = "single"
	album.ReleaseDate = song.ReleaseDate
	album.TotalTracks = 1

	log.Println("ArtistID =", song.ArtistID)
	log.Println("Title =", song.Title)
	log.Println("ReleaseDate =", song.ReleaseDate)

	var albumID string
	albumID, err = s.AlbumRepo.CreateAlbum(ctx, txn, &album)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("albumID =", albumID)

	var songID string
	song.AlbumID = albumID
	songID, err = s.SongRepo.CreateSong(ctx, txn, song)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("songID =", songID)

	duration := 5 * time.Minute
	expirationTime := time.Now().Add(duration)

	bucket := s.Config.AwsS3Config.Bucket
	key := util.GenerateRawUploadKey(songID, song.ArtistID, ".flac")
	var uri string
	uri, err = s.Storage.UploadURL(ctx, bucket, key, duration)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("uri =", uri)

	res := &models.CreateSongWithAlbumRes{
		SongID:     songID,
		UploadURI:  uri,
		Expiration: expirationTime,
	}

	if err = txn.Commit(ctx); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return res, nil
}
