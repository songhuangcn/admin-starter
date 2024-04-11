package apierror

import (
	"github.com/gin-gonic/gin"
	. "github.com/songhuangcn/admin-template/internal/common/core"
)

func Middleware() gin.HandlerFunc {
	return gin.CustomRecovery(recoveryHandler)
}

func recoveryHandler(ctx *gin.Context, err any) {
	if obj, ok := err.(*ApiError); ok {
		ctx.AbortWithStatusJSON(obj.Status, Hash{"message": obj.Msg})
	} else {
		ctx.AbortWithStatus(500)
	}
}
