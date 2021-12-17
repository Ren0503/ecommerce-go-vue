package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name        string
	Description string
}

type UserRole struct {
	UserId uint
	RoleId uint
}

func (UserRole) TableName() string {
	return "users_roles"
}

func Any(roles []Role, f func(Role) bool) bool {
	for _, role := range roles {
		if f(role) {
			return true
		}
	}
	return false
}
