package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Name     string
	Password string
	Email    string
	RoleId   uuid.UUID
	Role     *Role
}

func NewUser(name string, roleId uuid.UUID, email string, password string) *User {
	return &User{
		Id:       uuid.New(),
		Name:     name,
		RoleId:   roleId,
		Email:    email,
		Password: password,
	}
}
