package dto

import "github.com/google/uuid"

type GetByIdDTO struct {
	Id string `params:"id" validate:"required,uuid4"`
}

type ReturnIdDTO struct {
	Id uuid.UUID `json:"id"`
}
