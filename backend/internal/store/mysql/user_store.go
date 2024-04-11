package mysql

import (
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func (u *userStore) List(page, pageSize int) ([]*model.User, int64) {
	var users []*model.User
	var total int64
	result := u.db.Scopes(latest, paged(page, pageSize)).
		Preload("Roles.RolesPermissions").
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&total)
	HandleError(result.Error)

	return users, total
}

func (u *userStore) Get(id uint) *model.User {
	user := model.NewUser()
	result := u.db.Preload("Roles.RolesPermissions").Take(user, int(id))
	HandleError(result.Error)

	return user
}

func (u *userStore) GetBy(field, value string) *model.User {
	var relation *gorm.DB
	switch field {
	case "username":
		relation = u.db.Where("username = ?", value)
	default:
		panic("Not implement")
	}
	user := model.NewUser()
	result := relation.First(user)
	HandleError(result.Error)

	return user
}

func (u *userStore) Create(user *model.User) {
	result := u.db.Select("Name", "Username", "PasswordDigest", "Roles").Save(user)
	HandleError(result.Error)
}

func (u *userStore) Update(user *model.User) {
	u.db.Transaction(func(tx *gorm.DB) error {
		// Save 不会删除旧的关联，这里先做关联删除
		roles := user.Roles
		tx.Model(user).Association("Roles").Clear()
		user.Roles = roles

		if user.PasswordDigest != "" {
			result := tx.Select("Name", "Username", "Roles").Save(user)
			HandleError(result.Error)
		} else {
			result := tx.Select("Name", "Username", "Roles", "PasswordDigest").Save(user)
			HandleError(result.Error)
		}

		return nil
	})
}

func (u *userStore) Delete(id uint) {
	result := u.db.Delete(&model.User{}, int(id))
	HandleError(result.Error)
}
