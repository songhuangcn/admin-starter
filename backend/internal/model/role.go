package model

import (
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"gorm.io/gorm"
)

type Role struct {
	Model
	Users            []User            `json:"-" gorm:"many2many:users_roles"`
	Name             string            `json:"name"`
	RolesPermissions []RolesPermission `json:"-"`

	// 以下为动态属性，通过回调自动设置值
	PermissionNames []string `json:"permission_names" gorm:"-"`
	PermissionCount int      `json:"permission_count" gorm:"-"`
}

func NewRole() *Role {
	return &Role{}
}

func (r *Role) AfterFind(tx *gorm.DB) (err error) {
	r.setDynamicFields()

	return nil
}

func (r *Role) AfterSave(tx *gorm.DB) (err error) {
	r.setDynamicFields()

	return nil
}

func (r *Role) setDynamicFields() {
	r.PermissionNames = Pluck[RolesPermission, string](r.RolesPermissions, "PermissionName")
	r.PermissionCount = len(r.RolesPermissions)
}

type RoleCreateRequest struct {
	Name            string   `json:"name" binding:"required"`
	PermissionNames []string `json:"permission_names" binding:"permission_names"`
}

func NewRoleCreateRequest() *RoleCreateRequest {
	return &RoleCreateRequest{}
}

type RoleUpdateRequest struct {
	Name            string   `json:"name"`
	PermissionNames []string `json:"permission_names" binding:"permission_names"`
}

func NewRoleUpdateRequest() *RoleUpdateRequest {
	return &RoleUpdateRequest{}
}
