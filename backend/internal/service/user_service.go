package service

import (
	"github.com/songhuangcn/admin-template/internal/model"
)

type UserService interface {
	List(page, pageSize int) ([]*model.User, int64)
	Get(id uint) *model.User
	Create(request *model.UserCreateRequest) *model.User
	Update(id uint, request *model.UserUpdateRequest) *model.User
	Delete(id uint)
}
