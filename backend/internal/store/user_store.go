package store

import (
	"github.com/songhuangcn/admin-template/internal/model"
)

type UserStore interface {
	List(page, pageSize int) ([]*model.User, int64)
	Get(id uint) *model.User
	GetBy(field, value string) *model.User
	Create(user *model.User)
	Update(user *model.User)
	Delete(id uint)
}
