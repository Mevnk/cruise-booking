package domain

import (
	"github.com/google/uuid"
	"time"
)

type Cruise struct {
	Id            uuid.UUID `bun:",pk"`
	ShipID        uuid.UUID
	DepartureDate time.Time
	Price         int
}

func NewCruise(shipID uuid.UUID, departureDate time.Time, price int) *Cruise {
	return &Cruise{
		Id:            uuid.New(),
		ShipID:        shipID,
		DepartureDate: departureDate,
		Price:         price,
	}
}
