package srvimpl

import (
	"github.com/songhuangcn/admin-template/internal/common/auth"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/model"
	"github.com/songhuangcn/admin-template/internal/store"
)

type userService struct {
	store store.Factory
}

func (u *userService) List(page, pageSize int) ([]*model.User, int64) {
	users, total := u.store.User().List(page, pageSize)

	return users, total
}

func (u *userService) Get(id uint) *model.User {
	user := u.store.User().Get(id)

	return user
}

func (u *userService) Create(request *model.UserCreateRequest) *model.User {
	user := model.NewUser()
	u.assignRequest(user, request)
	roles := u.store.Role().ListByIDs(request.RoleIds)
	user.Roles = Map(roles, func(role *model.Role) model.Role { return *role })
	u.store.User().Create(user)

	return user
}

func (u *userService) Update(id uint, request *model.UserUpdateRequest) *model.User {
	user := u.store.User().Get(id)
	u.assignRequest(user, (*model.UserCreateRequest)(request))
	roles := u.store.Role().ListByIDs(request.RoleIds)
	user.Roles = Map(roles, func(role *model.Role) model.Role { return *role })
	u.store.User().Update(user)

	return user
}

func (u *userService) Delete(id uint) {
	u.store.User().Delete(id)
}

func (u *userService) assignRequest(user *model.User, request *model.UserCreateRequest) {
	user.Username = request.Username
	user.Name = request.Name
	if request.Password != "" {
		user.PasswordDigest = auth.BcryptHash(request.Password)
	}
}
