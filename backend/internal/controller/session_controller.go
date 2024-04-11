package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/songhuangcn/admin-template/internal/common/auth"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"github.com/songhuangcn/admin-template/internal/service"
)

type sessionController struct {
	service service.Factory
}

func NewSessionController(service service.Factory) *sessionController {
	return &sessionController{
		service: service,
	}
}

func (s *sessionController) Login(ctx *gin.Context) {
	loginRequest := &model.SessionLoginRequest{}
	panicBind(ctx, loginRequest)
	token := s.service.Session().Login(ctx, loginRequest)

	renderSuccess(ctx, Hash{"data": Hash{"token": token}})
}

func (s *sessionController) User(ctx *gin.Context) {
	uid := auth.CurrentUid(ctx)
	user := s.service.Session().GetCurrent(ctx, uid)

	renderSuccess(ctx, Hash{"data": user})
}
