package dto

import "github.com/google/uuid"

type TicketDTO struct {
	ID             uuid.UUID `json:"id"`
	CruiseID       uuid.UUID `json:"cruiseId"`
	PassengerID    uuid.UUID `json:"passengerId"`
	PassengerClass string    `json:"passengerClass"`
	Bonuses        string    `json:"bonuses"`
}

type AddTicketDTO struct {
	CruiseID       uuid.UUID `json:"cruiseId"`
	PassengerID    uuid.UUID `json:"passengerId"`
	PassengerClass string    `json:"passengerClass"`
	Bonuses        string    `json:"bonuses"`
}
