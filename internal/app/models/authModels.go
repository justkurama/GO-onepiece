package models

type User struct {
	ID        uint
	Login     string
	Password  string
	UserRoles []UserRole
}
type Permission struct {
	ID              uint
	Permission      string
	RolePermissions []RolePermission
}
type Role struct {
	ID              uint
	Role            string
	UserRoles       []UserRole
	RolePermissions []RolePermission
}
type UserRole struct {
	ID     uint
	UserID uint
	RoleID uint
	User   User
	Role   Role
}
type RolePermission struct {
	ID           uint
	RoleID       uint
	PermissionID uint
	Role         Role
	Permission   Permission
}
