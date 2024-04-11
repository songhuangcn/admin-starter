package model

type RolesPermission struct {
	Model
	Role           Role
	RoleID         uint   `json:"role_id"`
	PermissionName string `json:"permission_name"`
}

func NewRolesPermission() *RolesPermission {
	return &RolesPermission{}
}
