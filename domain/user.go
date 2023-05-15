package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `bun:",pk"`
	Name     string
	Password string
	Email    string
	RoleId   uuid.UUID
	Role     *Role `bun:"rel:belongs-to"`
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
