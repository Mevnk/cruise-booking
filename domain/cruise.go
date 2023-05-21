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
	Route         string
	NofPorts      int
	Duration      int
}

func NewCruise(shipID uuid.UUID, departureDate time.Time, price int, route string, nOfSports, duration int) *Cruise {
	return &Cruise{
		Id:            uuid.New(),
		ShipID:        shipID,
		DepartureDate: departureDate,
		Price:         price,
		Route:         route,
		NofPorts:      nOfSports,
		Duration:      duration,
	}
}
