package domain

import (
	"github.com/google/uuid"
	"time"
)

type Cruise struct {
	Id            uuid.UUID
	DepartureDate time.Time
	Price         int
	Route         string
	NofPorts      int
	Duration      int
}

func NewCruise(departureDate time.Time, price int, route string, nOfSports, duration int) *Cruise {
	return &Cruise{
		Id:            uuid.New(),
		DepartureDate: departureDate,
		Price:         price,
		Route:         route,
		NofPorts:      nOfSports,
		Duration:      duration,
	}
}
