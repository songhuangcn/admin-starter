package srvimpl

import (
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"github.com/songhuangcn/admin-template/internal/store"
)

type roleService struct {
	store store.Factory
}

func (r *roleService) List(page, pageSize int) ([]*model.Role, int64) {
	roles, total := r.store.Role().List(page, pageSize)

	return roles, total
}

func (r *roleService) Create(request *model.RoleCreateRequest) *model.Role {
	role := model.NewRole()
	role.Name = request.Name
	role.RolesPermissions = Map(request.PermissionNames, func(permissionName string) model.RolesPermission {
		return model.RolesPermission{
			PermissionName: permissionName,
		}
	})
	r.store.Role().Create(role)

	return role
}

func (r *roleService) Update(id uint, request *model.RoleUpdateRequest) *model.Role {
	role := r.store.Role().Get(id)
	role.Name = request.Name

	role.RolesPermissions = Map(request.PermissionNames, func(permissionName string) model.RolesPermission {
		return model.RolesPermission{
			PermissionName: permissionName,
		}
	})
	r.store.Role().Update(role)

	return role
}

func (r *roleService) Delete(id uint) {
	r.store.Role().Get(id)
	r.store.Role().Delete(id)
}
