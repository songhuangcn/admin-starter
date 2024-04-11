package srvimpl

import (
	"github.com/songhuangcn/admin-template/internal/service"
	"github.com/songhuangcn/admin-template/internal/store"
)

type factory struct {
	store store.Factory
}

func NewService(store store.Factory) service.Factory {
	return &factory{
		store: store,
	}
}

func (f *factory) User() service.UserService {
	return &userService{
		store: f.store,
	}
}

func (f *factory) Role() service.RoleService {
	return &roleService{
		store: f.store,
	}
}

func (f *factory) Session() service.SessionService {
	return &sessionService{
		store: f.store,
	}
}

func (f *factory) Permission() service.PermissionService {
	return &permissionService{
		store: f.store,
	}
}
