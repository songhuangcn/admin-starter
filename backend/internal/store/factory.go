package store

type Factory interface {
	User() UserStore
	Role() RoleStore
	RolesPermission() RolesPermissionStore
}
