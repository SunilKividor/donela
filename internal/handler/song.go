package handler

import (
	"net/http"

	"github.com/SunilKividor/donela/internal/models"
	"github.com/SunilKividor/donela/internal/service"
	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	SongService *service.SongService
}

func NewSongHandler(service *service.SongService) *SongHandler {
	return &SongHandler{
		SongService: service,
	}
}

func (s *SongHandler) CreateSongWithAlbum(c *gin.Context) {
	var req models.CreateSongWithAlbumReq

	if c.ShouldBindJSON(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request format"})
		return
	}

	ctx := c.Request.Context()
	artistID := c.GetString("id")

	song := &models.Song{
		ArtistID:    artistID,
		Title:       req.Title,
		Genre:       req.Genre,
		ReleaseDate: req.ReleaseDate,
	}

	res, err := s.SongService.CreateSongWithAlbum(ctx, song)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
