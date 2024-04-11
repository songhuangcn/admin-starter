package service

import (
	"context"

	"github.com/songhuangcn/admin-template/internal/model"
)

type PermissionService interface {
	List(ctx context.Context) []*model.Permission
	Authz(ctx context.Context, user *model.User, handlerName string)
}
