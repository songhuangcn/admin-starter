package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"github.com/songhuangcn/admin-template/internal/service"
)

type roleController struct {
	service service.Factory
}

func NewRoleController(service service.Factory) *roleController {
	return &roleController{
		service: service,
	}
}

func (r *roleController) Index(ctx *gin.Context) {
	pageRequest := model.NewPageRequest()
	panicBind(ctx, pageRequest)
	roles, total := r.service.Role().List(pageRequest.Page, pageRequest.PageSize)
	setPagination(ctx, pageRequest.Page, pageRequest.PageSize, total)

	renderSuccess(ctx, Hash{"data": roles})
}

func (r *roleController) Create(ctx *gin.Context) {
	request := model.NewRoleCreateRequest()
	panicBind(ctx, request)
	role := r.service.Role().Create(request)

	renderSuccess(ctx, Hash{"data": role})
}

func (r *roleController) Update(ctx *gin.Context) {
	id := getParamID(ctx)
	request := model.NewRoleUpdateRequest()
	panicBind(ctx, request)
	role := r.service.Role().Update(id, request)

	renderSuccess(ctx, Hash{"data": role})
}

func (r *roleController) Delete(ctx *gin.Context) {
	id := getParamID(ctx)
	r.service.Role().Delete(id)

	renderSuccess(ctx)
}
