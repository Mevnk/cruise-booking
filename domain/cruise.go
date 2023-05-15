package domain

import (
	"github.com/google/uuid"
	"time"
)

type Cruise struct {
	Id            uuid.UUID `bun:",pk"`
	ShipID        int
	DepartureDate time.Time
	Price         int
	Excursions    string
}

func NewCruise(shipID int, departureDate time.Time, price int, excursions string) *Cruise {
	return &Cruise{
		Id:            uuid.New(),
		ShipID:        shipID,
		DepartureDate: departureDate,
		Price:         price,
		Excursions:    excursions,
	}
}
