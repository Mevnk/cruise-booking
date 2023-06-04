package dto

import "github.com/google/uuid"

type ExcursionDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
}

type AddExcursionDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type UpdateExcursionDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
