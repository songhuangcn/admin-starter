package service

import (
	"github.com/songhuangcn/admin-template/internal/model"
)

type RoleService interface {
	List(page, pageSize int) ([]*model.Role, int64)
	Create(request *model.RoleCreateRequest) *model.Role
	Update(id uint, request *model.RoleUpdateRequest) *model.Role
	Delete(id uint)
}
