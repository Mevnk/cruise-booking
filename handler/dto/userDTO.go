package dto

import "github.com/google/uuid"

type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type SignInResponseDTO struct {
	Token string `json:"token,omitempty"`
}

type SignUpRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type RoleDTO struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserDTO struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	IsActive bool      `json:"is_active"`
	Role     *RoleDTO  `json:"role"`
}

type UpdateRoleDTO struct {
	Role string `json:"role" validate:"required"`
}
