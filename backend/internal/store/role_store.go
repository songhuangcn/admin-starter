package store

import (
	"github.com/songhuangcn/admin-template/internal/model"
)

type RoleStore interface {
	List(page, pageSize int) ([]*model.Role, int64)
	ListByIDs(ids []uint) []*model.Role
	Get(id uint) *model.Role
	Create(role *model.Role)
	Update(role *model.Role)
	Delete(id uint)
}
