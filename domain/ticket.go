package domain

import "github.com/google/uuid"

type Ticket struct {
	Id             uuid.UUID `bun:",pk"`
	CruiseID       int
	PassengerID    int
	PassengerClass string
	Bonuses        string
}

func NewTicket(cruiseID int, passengerID int, passengerClass string, bonuses string) *Ticket {
	return &Ticket{
		Id:             uuid.New(),
		CruiseID:       cruiseID,
		PassengerID:    passengerID,
		PassengerClass: passengerClass,
		Bonuses:        bonuses,
	}
}
