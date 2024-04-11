package model

import (
	"context"
	"slices"

	"github.com/songhuangcn/admin-template/internal/common/locale"
	"github.com/vorlif/spreak/localize"
)

type Permission struct {
	Controller     string `json:"controller"`
	Action         string `json:"action"`
	Name           string `json:"name"`
	ControllerI18n string `json:"controller_i18n"`
	ActionI18n     string `json:"action_i18n"`
	NameI18n       string `json:"name_i18n"`
}

func NewPermission(ctx context.Context, controller, action string) *Permission {
	permission := &Permission{
		Controller: controller,
		Action:     action,
	}
	permission.setDynamicFields(ctx)

	return permission
}

var (
	// 这些 controller 粒度太小，不单独校验权限
	UnauthzControllers = []string{
		"session",
		"permission",
	}

	// 列出来方便 gettext 自动提取 .pot 翻译键
	ControllerMsgids = []localize.MsgID{
		"controller|user",
		"controller|role",

		// 以下资源在 UnauthzControllers 数组中，不需要校验，也就不需要翻译
		// "controller|session",
		// "controller|permission",
	}

	ActionMsgids = []localize.MsgID{
		"action|index",
		"action|create",
		"action|update",
		"action|delete",
	}
)

func (p *Permission) setDynamicFields(ctx context.Context) {
	controllerMsgid := "controller|" + p.Controller // xspreak: ignore
	if !slices.Contains(ControllerMsgids, controllerMsgid) {
		panic("Need add new controller to the msg list")
	}
	actionMsgid := "action|" + p.Action // xspreak: ignore
	if !slices.Contains(ActionMsgids, actionMsgid) {
		panic("Need add new action to the msg list")
	}
	p.ControllerI18n = locale.T(ctx, controllerMsgid)
	p.ActionI18n = locale.T(ctx, actionMsgid)
	p.Name = p.Controller + "#" + p.Action
	p.NameI18n = p.ControllerI18n + p.ActionI18n
}
