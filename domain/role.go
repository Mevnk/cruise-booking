package domain

import "github.com/google/uuid"

type Role struct {
	Id   uuid.UUID
	Name string
}

func NewRole(name string) *Role {
	return &Role{Id: uuid.New(), Name: name}
}

const (
	USER    = "user"
	MANAGER = "manager"
	ADMIN   = "admin"
)
