package mysql

import (
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"gorm.io/gorm"
)

type roleStore struct {
	db *gorm.DB
}

func (r *roleStore) List(page, pageSize int) ([]*model.Role, int64) {
	var roles []*model.Role
	var total int64
	result := r.db.Scopes(latest, paged(page, pageSize)).
		Preload("RolesPermissions").
		Find(&roles).
		Offset(-1).
		Limit(-1).
		Count(&total)
	HandleError(result.Error)

	return roles, total
}

func (r *roleStore) ListByIDs(ids []uint) []*model.Role {
	if len(ids) == 0 {
		return []*model.Role{}
	}

	var roles []*model.Role
	result := r.db.Find(&roles, ids)
	HandleError(result.Error)

	return roles
}

func (r *roleStore) Get(id uint) *model.Role {
	role := model.NewRole()
	result := r.db.Take(role, id)
	HandleError(result.Error)

	return role
}

func (r *roleStore) Create(role *model.Role) {
	result := r.db.Create(role)
	HandleError(result.Error)
}

func (r *roleStore) Update(role *model.Role) {
	r.db.Transaction(func(tx *gorm.DB) error {
		// Save 不会执行关联的删除操作，这里需要手动删除
		tx.Where("role_id = ?", role.ID).Delete(model.NewRolesPermission())
		result := tx.Select("Name", "RolesPermissions").Save(role)
		HandleError(result.Error)

		return nil
	})
}

func (u *roleStore) Delete(id uint) {
	result := u.db.Delete(&model.Role{}, int(id))
	HandleError(result.Error)
}
