package srvimpl

import (
	"github.com/gin-gonic/gin"
	"github.com/songhuangcn/admin-template/internal/common/auth"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/common/locale"
	"github.com/songhuangcn/admin-template/internal/config"
	"github.com/songhuangcn/admin-template/internal/model"
	"github.com/songhuangcn/admin-template/internal/store"
)

type sessionService struct {
	store store.Factory
}

func (s *sessionService) Login(ctx *gin.Context, request *model.SessionLoginRequest) string {
	user := s.getAuthUser(ctx, request.Username)
	if !auth.BcryptCompare(request.Password, user.PasswordDigest) {
		PanicApiError(locale.T(ctx, "用户名或密码错误"), 401)
	}
	token, err := auth.JwtEncode(user.ID, config.Auth.JwtExpHours, config.App.Secret)
	if err != nil {
		panic(err)
	}

	return token
}

func (s *sessionService) GetCurrent(ctx *gin.Context, uid uint) *model.User {
	user := s.getCurrentUser(ctx, uid)

	return user
}

func (s *sessionService) getAuthUser(ctx *gin.Context, username string) *model.User {
	defer func() {
		err := recover()
		if _, ok := err.(*ApiError); !ok {
			return
		}
		// 找不到用户不能响应 404，需要当成用户名输入错误响应 401
		PanicApiError(locale.T(ctx, "用户名或密码错误"), 401)
	}()

	return s.store.User().GetBy("username", username)
}

func (s *sessionService) getCurrentUser(ctx *gin.Context, uid uint) *model.User {
	defer func() {
		err := recover()
		if _, ok := err.(*ApiError); !ok {
			return
		}
		// 找不到用户不能响应 404，需要当成认证问题响应 401
		PanicApiError(locale.T(ctx, "Token 已失效"), 401)
	}()

	return s.store.User().Get(uid)
}
