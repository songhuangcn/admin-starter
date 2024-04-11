package service

import (
	"github.com/gin-gonic/gin"
	"github.com/songhuangcn/admin-template/internal/model"
)

type SessionService interface {
	Login(ctx *gin.Context, request *model.SessionLoginRequest) string
	GetCurrent(ctx *gin.Context, uid uint) *model.User
}
