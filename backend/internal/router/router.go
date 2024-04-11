package router

import (
	"github.com/gin-gonic/gin"
	"github.com/songhuangcn/admin-template/internal/common/apierror"
	"github.com/songhuangcn/admin-template/internal/common/auth"
	"github.com/songhuangcn/admin-template/internal/common/cors"
	"github.com/songhuangcn/admin-template/internal/common/locale"
	"github.com/songhuangcn/admin-template/internal/controller"
	"github.com/songhuangcn/admin-template/internal/service/srvimpl"
	"github.com/songhuangcn/admin-template/internal/store/mysql"
)

var (
	store       = mysql.NewStore()
	service     = srvimpl.NewService(store)
	mdwLogger   = gin.Logger()
	mdwRecovery = apierror.Middleware()
	mdwCors     = cors.Middleware()
	mdwAuthn    = auth.MdwAuthn()
	mdwAuthz    = auth.MdwAuthz(service)
	mdwLocale   = locale.Middleware()
)

func Config(engine *gin.Engine) {
	engine.Use(
		mdwLogger,
		mdwRecovery,
		mdwCors,
		mdwLocale,
	)

	api := engine.Group("/api")
	{
		sessions := api.Group("")
		{
			ctl := controller.NewSessionController(service)

			sessions.POST("/login", ctl.Login)
			sessions.GET("/user", mdwAuthn, ctl.User)
		}

		users := api.Group("/users", mdwAuthz)
		{
			ctl := controller.NewUserController(service)

			users.GET("", ctl.Index)
			users.POST("", ctl.Create)
			users.PUT("/:id", ctl.Update)
			users.DELETE("/:id", ctl.Delete)
		}

		roles := api.Group("/roles", mdwAuthz)
		{
			ctl := controller.NewRoleController(service)

			roles.GET("", ctl.Index)
			roles.POST("", ctl.Create)
			roles.PUT("/:id", ctl.Update)
			roles.DELETE("/:id", ctl.Delete)
		}

		permissions := api.Group("/permissions", mdwAuthz)
		{
			ctl := controller.NewPermissionController(service)

			permissions.GET("", ctl.Index)
		}
	}
}
