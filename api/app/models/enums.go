package models

type RoleEnum string

const (
	APIAdmin RoleEnum = "APIAdmin"
	Admin    RoleEnum = "Admin"
	Member   RoleEnum = "User"
)
