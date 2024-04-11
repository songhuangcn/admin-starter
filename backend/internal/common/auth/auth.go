package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/common/enum"
	"github.com/songhuangcn/admin-template/internal/common/locale"
	"github.com/songhuangcn/admin-template/internal/config"
	"github.com/songhuangcn/admin-template/internal/service"
)

func MdwAuthn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		uid, err := JwtDecode(tokenString, config.App.Secret)
		if err != nil {
			PanicApiError(err.Error(), 401)
		}

		ctx.Set(enum.CurrentUid, uid)
	}
}

func MdwAuthz(service service.Factory) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 权限中间件已经加上了认证逻辑，路由中不需要在添加认证中间件
		MdwAuthn()(ctx)

		uid := CurrentUid(ctx)
		user := service.Session().GetCurrent(ctx, uid)
		if user.IsAdmin {
			return
		}
		handerName := ctx.HandlerName()
		service.Permission().Authz(ctx, user, handerName)
	}
}

func CurrentUid(ctx *gin.Context) uint {
	obj, ok := ctx.Get(enum.CurrentUid)
	if !ok {
		PanicApiError(locale.T(ctx, "登录已失效"), 401)
	}
	uid, ok := obj.(uint)
	if !ok {
		PanicApiError(locale.T(ctx, "登录已失效"), 401)
	}

	return uid
}
