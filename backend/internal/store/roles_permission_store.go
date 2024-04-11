package store

import (
	"github.com/songhuangcn/admin-template/internal/model"
)

type RolesPermissionStore interface {
	ListAuthed(user *model.User) []*model.RolesPermission
}
