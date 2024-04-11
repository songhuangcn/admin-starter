package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/service"
)

const (
	RouteRegexp = `\(\*(\w+)Controller\)\.(\w+)-fm$`
)

type permissionController struct {
	service service.Factory
}

func NewPermissionController(service service.Factory) *permissionController {
	return &permissionController{
		service: service,
	}
}

func (p *permissionController) Index(ctx *gin.Context) {
	permissions := p.service.Permission().List(ctx)

	renderSuccess(ctx, Hash{"data": permissions})
}
