package service

type Factory interface {
	User() UserService
	Role() RoleService
	Session() SessionService
	Permission() PermissionService
}
