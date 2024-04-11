package mysql

import (
	log "github.com/sirupsen/logrus"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"gorm.io/gorm"
)

type rolesPermissionStore struct {
	db *gorm.DB
}

func (rp *rolesPermissionStore) ListAuthed(user *model.User) []*model.RolesPermission {
	var roles []model.Role
	err := rp.db.Model(user).Association("Roles").Find(&roles)
	log.Debugf("roles: %#v\n", roles)
	HandleError(err)

	roleIds := Pluck[model.Role, uint](roles, "ID")
	log.Debugf("roleIds: %#v\n", roleIds)

	var rolesPermissions []*model.RolesPermission
	err = rp.db.Model(roles).Association("RolesPermissions").Find(&rolesPermissions)
	log.Debugf("rolesPermissions: %#v\n", rolesPermissions)
	HandleError(err)

	return rolesPermissions
}
