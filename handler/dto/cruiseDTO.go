package dto

import (
	"github.com/google/uuid"
	"time"
)

type CruiseDTO struct {
	ID            uuid.UUID `json:"id"`
	DepartureDate time.Time `json:"departureDate"`
	Price         int       `json:"price"`
	Route         string    `json:"route"`
	NumberOfPorts int       `json:"numberOfPorts"`
	Duration      int       `json:"duration"`
}

type AddCruiseDTO struct {
	DepartureDate time.Time `json:"departureDate"`
	Price         int       `json:"price"`
	Route         string    `json:"route"`
	NumberOfPorts int       `json:"numberOfPorts"`
	Duration      int       `json:"duration"`
}

type UpdateCruiseDTO struct {
	DepartureDate time.Time `json:"departureDate"`
	Price         int       `json:"price"`
	Route         string    `json:"route"`
	NumberOfPorts int       `json:"numberOfPorts"`
	Duration      int       `json:"duration"`
}
