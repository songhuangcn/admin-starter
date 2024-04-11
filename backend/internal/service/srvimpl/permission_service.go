package srvimpl

import (
	"context"
	"reflect"
	"regexp"
	"runtime"
	"slices"
	"strings"

	log "github.com/sirupsen/logrus"

	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/common/locale"
	"github.com/songhuangcn/admin-template/internal/global"
	"github.com/songhuangcn/admin-template/internal/model"
	"github.com/songhuangcn/admin-template/internal/store"
)

type permissionService struct {
	store store.Factory
}

func (p *permissionService) List(ctx context.Context) []*model.Permission {
	permissions := []*model.Permission{}
	for _, route := range global.Engine.Routes() {
		// 从路由路径中提取控制器名称和动作名称
		handlerName := runtime.FuncForPC(reflect.ValueOf(route.HandlerFunc).Pointer()).Name()
		controller, action := p.parseHandler(handlerName)
		if slices.Contains(model.UnauthzControllers, controller) {
			continue
		}

		permission := model.NewPermission(ctx, controller, action)
		permissions = append(permissions, permission)
	}

	return permissions
}

func (p *permissionService) Authz(ctx context.Context, user *model.User, handlerName string) {
	controller, action := p.parseHandler(handlerName)
	// 访问的资源不需要检验权限
	if slices.Contains(model.UnauthzControllers, controller) {
		return
	}

	currentRpName := controller + "#" + action
	authedRps := p.store.RolesPermission().ListAuthed(user)
	// Pluck 需要结构体，这里是结构体指针，所以用不了
	// authedRpNames := Pluck[model.RolesPermission, string](*authedRps, "PermissionName")
	authedRpNames := Map(authedRps, func(rp *model.RolesPermission) string {
		return rp.PermissionName
	})

	log.Debugf("authedRpNames: %#v\n", authedRpNames)
	log.Debugf("currentRpName: %#v\n", currentRpName)

	if !slices.Contains(authedRpNames, currentRpName) {
		PanicApiError(locale.T(ctx, "您没有权限"), 403)
	}
}

func (p *permissionService) parseHandler(handlerName string) (string, string) {
	pattern := regexp.MustCompile(`\(\*(\w+)Controller\)\.(\w+)-fm$`)
	matches := pattern.FindStringSubmatch(handlerName)
	var controller, action string
	if len(matches) == 3 {
		controller, action = strings.ToLower(matches[1]), strings.ToLower(matches[2])
	}

	return controller, action
}
