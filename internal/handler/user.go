package handler

import (
	"context"
	"net/http"

	"github.com/SunilKividor/donela/internal/db/repository"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (uh *UserHandler) GetUserID(c *gin.Context) {
	uh.UserRepo.GetByID(context.Background(), "123")
	c.JSON(http.StatusOK, gin.H{"user id": "123"})
}
