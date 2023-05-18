package domain

import "github.com/google/uuid"

type Ticket struct {
	Id             uuid.UUID `bun:",pk"`
	CruiseID       uuid.UUID `bun:",pk"`
	PassengerID    uuid.UUID
	PassengerClass string
	Bonuses        string
}

func NewTicket(cruiseID uuid.UUID, passengerID uuid.UUID, passengerClass string, bonuses string) *Ticket {
	return &Ticket{
		Id:             uuid.New(),
		CruiseID:       cruiseID,
		PassengerID:    passengerID,
		PassengerClass: passengerClass,
		Bonuses:        bonuses,
	}
}
