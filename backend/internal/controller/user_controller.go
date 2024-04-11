package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"github.com/songhuangcn/admin-template/internal/service"
)

type userController struct {
	service service.Factory
}

func NewUserController(service service.Factory) *userController {
	return &userController{
		service: service,
	}
}

func (u *userController) Index(ctx *gin.Context) {
	pageRequest := model.NewPageRequest()
	panicBind(ctx, pageRequest)
	users, total := u.service.User().List(pageRequest.Page, pageRequest.PageSize)
	setPagination(ctx, pageRequest.Page, pageRequest.PageSize, total)

	renderSuccess(ctx, Hash{"data": users})
}

func (u *userController) Create(ctx *gin.Context) {
	request := &model.UserCreateRequest{}
	panicBind(ctx, request)
	userController := u.service.User().Create(request)

	renderSuccess(ctx, Hash{"data": userController})
}

func (u *userController) Update(ctx *gin.Context) {
	id := getParamID(ctx)
	request := &model.UserUpdateRequest{}
	panicBind(ctx, request)
	userController := u.service.User().Update(id, request)

	renderSuccess(ctx, Hash{"data": userController})
}

func (u *userController) Delete(ctx *gin.Context) {
	id := getParamID(ctx)
	u.service.User().Get(id)
	u.service.User().Delete(id)

	renderSuccess(ctx)
}
