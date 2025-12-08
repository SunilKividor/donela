package di

import (
	"context"

	"github.com/SunilKividor/donela/internal/api"
	"github.com/SunilKividor/donela/internal/authentication/http/middleware"
	"github.com/SunilKividor/donela/internal/authentication/jwt"
	"github.com/SunilKividor/donela/internal/config"
	pg "github.com/SunilKividor/donela/internal/db/pg"
	redisdb "github.com/SunilKividor/donela/internal/db/redis"
	"github.com/SunilKividor/donela/internal/db/repository"
	"github.com/SunilKividor/donela/internal/handler"
	"github.com/SunilKividor/donela/internal/service"
	"github.com/SunilKividor/donela/internal/storage"
)

func InitializeApp() (*api.Server, error) {
	ctx := context.Background()

	cfg := config.Load()

	pgConn := pg.NewConnection(cfg.PostgresConfig.ConnectionString)
	pool, err := pgConn.Connect()
	if err != nil {
		return nil, err
	}

	redisConn := redisdb.NewConnection(cfg.RedisConfig.ConnectionString)
	redisClient, err := redisConn.Connect()
	if err != nil {
		return nil, err
	}

	s3Client, err := config.NewS3Client(ctx, cfg)
	if err != nil {
		return nil, err
	}

	jwtRepo := repository.NewAuthRepository(pool, redisClient)
	jwtAuth := jwt.NewJWTAuthenticationClient(jwtRepo, cfg.JWTConfig.Secret)

	s3Storage := storage.NewS3StorageClient(s3Client)

	userRepo := repository.NewUserRepository(pool)

	songRepo := repository.NewSongRepository()
	albumRepo := repository.NewAlbumRepository()
	songService := service.NewSongService(cfg, pool, *songRepo, *albumRepo, s3Storage)

	handlers := &handler.Handlers{
		Authentication: handler.NewAuthenticationHandler(jwtAuth),
		Storage:        handler.NewStorageHandler(cfg, s3Storage),
		User:           handler.NewUserHandler(userRepo),
		SongHandler:    handler.NewSongHandler(songService),
	}

	server := api.NewServer(cfg)

	middleware := middleware.JWTMiddleware()

	api.RegisterRoutes(server.Engine, cfg, handlers, middleware)

	return server, nil
}
