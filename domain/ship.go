package domain

import "github.com/google/uuid"

type Ship struct {
	Id       uuid.UUID `bun:",pk"`
	Name     string
	PassCap  int
	Route    string
	NofPorts int
	Duration int
	Staff    string
}

func NewShip(name string, passCap int, route string, nofports int, staff string, duration int) *Ship {
	return &Ship{
		Id:       uuid.New(),
		Name:     name,
		PassCap:  passCap,
		Route:    route,
		NofPorts: nofports,
		Staff:    staff,
		Duration: duration,
	}
}
