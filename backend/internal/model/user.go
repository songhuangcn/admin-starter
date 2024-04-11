package model

import (
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username       string `json:"username"`
	Name           string `json:"name"`
	PasswordDigest string `json:"-"`
	IsAdmin        bool   `json:"is_admin"`
	Roles          []Role `json:"roles"            gorm:"many2many:users_roles"`

	// 以下为动态属性，通过回调自动设置值
	RoleNames       []string `json:"role_names"       gorm:"-"`
	PermissionNames []string `json:"permission_names" gorm:"-"`
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name"     binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	RoleIds  []uint `json:"role_ids" binding:"required"`
}

type UserUpdateRequest struct {
	Username string `json:"username" binding:""`
	Name     string `json:"name"     binding:""`
	Password string `json:"password" binding:"omitempty,min=6"`
	RoleIds  []uint `json:"role_ids" binding:""`
}

func NewUser() *User {
	return &User{}
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.setDynamicFields()

	return nil
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	u.setDynamicFields()

	return nil
}

func (u *User) setDynamicFields() {
	u.RoleNames = Pluck[Role, string](u.Roles, "Name")
	// u.RoleNames = Map(u.Roles, func(role Role) string { return role.Name })

	u.PermissionNames = []string{}
	added := make(map[string]bool)
	for _, r := range u.Roles {
		for _, rp := range r.RolesPermissions {
			if added[rp.PermissionName] {
				continue
			}
			added[rp.PermissionName] = true
			u.PermissionNames = append(u.PermissionNames, rp.PermissionName)
		}
	}
}
